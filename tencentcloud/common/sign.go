package common

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"sort"
	"strings"

	tchttp "github.com/qyzhaoxun/tencentcloud-sdk-go/tencentcloud/common/http"
)

const (
	SHA256 = "HmacSHA256"
	SHA1   = "HmacSHA1"
)

func Sign(s, secretKey, method string) string {
	hashed := hmac.New(sha1.New, []byte(secretKey))
	if method == SHA256 {
		hashed = hmac.New(sha256.New, []byte(secretKey))
	}
	hashed.Write([]byte(s))

	return base64.StdEncoding.EncodeToString(hashed.Sum(nil))
}

func signRequest(request tchttp.Request, credential *Credential, method string) (err error) {
	if method != SHA256 {
		method = SHA1
	}
	checkAuthParams(request, credential, method)
	s := getStringToSign(request)
	signature := Sign(s, credential.SecretKey, method)
	request.GetParams()["Signature"] = signature
	return
}

func checkAuthParams(request tchttp.Request, credential *Credential, method string) {
	params := request.GetParams()
	credentialParams := credential.GetCredentialParams()
	for key, value := range credentialParams {
		params[key] = value
	}
	params["SignatureMethod"] = method
	delete(params, "Signature")
}

func getStringToSign(request tchttp.Request) string {
	method := request.GetHttpMethod()
	domain := request.GetDomain()
	path := request.GetPath()

	text := method + domain + path + "?"

	params := request.GetParams()
	// sort params
	keys := make([]string, 0, len(params))
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i := range keys {
		k := keys[i]
		if params[k] == "" {
			continue
		}
		text += fmt.Sprintf("%v=%v&", strings.Replace(k, "_", ".", -1), params[k])
	}
	text = text[:len(text)-1]
	return text
}
