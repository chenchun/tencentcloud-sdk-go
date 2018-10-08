// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20180123

import (
    "encoding/json"

    tchttp "github.com/qyzhaoxun/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ActivateRuleRequest struct {
	*tchttp.BaseRequest
	// 规则Id
	RuleId *string `json:"RuleId" name:"RuleId"`
}

func (r *ActivateRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ActivateRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ActivateRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ActivateRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ActivateRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddDeviceRequest struct {
	*tchttp.BaseRequest
	// 产品Id
	ProductId *string `json:"ProductId" name:"ProductId"`
	// 设备名称，唯一标识某产品下的一个设备
	DeviceName *string `json:"DeviceName" name:"DeviceName"`
}

func (r *AddDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 设备信息
		Device *Device `json:"Device" name:"Device"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddProductRequest struct {
	*tchttp.BaseRequest
	// 产品名称，同一区域产品名称需唯一，支持中文、英文字母、中划线和下划线，长度不超过31个字符，中文占两个字符
	Name *string `json:"Name" name:"Name"`
	// 产品描述
	Description *string `json:"Description" name:"Description"`
	// 鉴权模式（1：动态令牌，推荐使用动态令牌）
	AuthType *uint64 `json:"AuthType" name:"AuthType"`
	// 数据模版（json数组）
	DataTemplate []*string `json:"DataTemplate" name:"DataTemplate" list`
	// 数据协议（native表示自定义，template表示数据模板，默认值为template）
	DataProtocol *string `json:"DataProtocol" name:"DataProtocol"`
}

func (r *AddProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 产品信息
		Product *Product `json:"Product" name:"Product"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddRuleRequest struct {
	*tchttp.BaseRequest
	// 名称
	Name *string `json:"Name" name:"Name"`
	// 描述
	Description *string `json:"Description" name:"Description"`
	// 查询
	Query *RuleQuery `json:"Query" name:"Query"`
	// 转发
	Actions []*Object `json:"Actions" name:"Actions" list`
}

func (r *AddRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 规则
		Rule *Rule `json:"Rule" name:"Rule"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddTopicRequest struct {
	*tchttp.BaseRequest
	// 产品Id
	ProductId *string `json:"ProductId" name:"ProductId"`
	// Topic名称
	TopicName *string `json:"TopicName" name:"TopicName"`
}

func (r *AddTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddTopicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// Topic信息
		Topic *Topic `json:"Topic" name:"Topic"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddTopicResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddUserRequest struct {
	*tchttp.BaseRequest
}

func (r *AddUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddUserRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AddUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 用户信息
		User *User `json:"User" name:"User"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AddUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AppAddUserRequest struct {
	*tchttp.BaseRequest
	// 用户名
	UserName *string `json:"UserName" name:"UserName"`
	// 密码
	Password *string `json:"Password" name:"Password"`
}

func (r *AppAddUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AppAddUserRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AppAddUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 应用用户
		AppUser *AppUser `json:"AppUser" name:"AppUser"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *AppAddUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AppAddUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AppUser struct {
	// 应用Id
	ApplicationId *string `json:"ApplicationId" name:"ApplicationId"`
	// 用户名
	UserName *string `json:"UserName" name:"UserName"`
	// 昵称
	NickName *string `json:"NickName" name:"NickName"`
	// 绑定设备列表
	Devices []*Object `json:"Devices" name:"Devices" list`
	// 创建时间
	CreateTime *string `json:"CreateTime" name:"CreateTime"`
	// 修改时间
	UpdateTime *string `json:"UpdateTime" name:"UpdateTime"`
}

type DeactivateRuleRequest struct {
	*tchttp.BaseRequest
	// 规则Id
	RuleId *string `json:"RuleId" name:"RuleId"`
}

func (r *DeactivateRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeactivateRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeactivateRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeactivateRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeactivateRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDeviceRequest struct {
	*tchttp.BaseRequest
	// 产品Id
	ProductId *string `json:"ProductId" name:"ProductId"`
	// 设备名称
	DeviceName *string `json:"DeviceName" name:"DeviceName"`
}

func (r *DeleteDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteProductRequest struct {
	*tchttp.BaseRequest
	// 产品Id
	ProductId *string `json:"ProductId" name:"ProductId"`
}

func (r *DeleteProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRuleRequest struct {
	*tchttp.BaseRequest
	// 规则Id
	RuleId *string `json:"RuleId" name:"RuleId"`
}

func (r *DeleteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicRequest struct {
	*tchttp.BaseRequest
	// TopicId
	TopicId *string `json:"TopicId" name:"TopicId"`
	// 产品Id
	ProductId *string `json:"ProductId" name:"ProductId"`
}

func (r *DeleteTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTopicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTopicResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Device struct {
	// 产品Id
	ProductId *string `json:"ProductId" name:"ProductId"`
	// 设备名称
	DeviceName *string `json:"DeviceName" name:"DeviceName"`
	// 设备密钥
	DeviceSecret *string `json:"DeviceSecret" name:"DeviceSecret"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime" name:"UpdateTime"`
	// 创建时间
	CreateTime *string `json:"CreateTime" name:"CreateTime"`
	// 设备信息
	DeviceInfo *Object `json:"DeviceInfo" name:"DeviceInfo"`
}

type DeviceStatus struct {
	// 设备名称
	DeviceName *string `json:"DeviceName" name:"DeviceName"`
	// 设备状态（inactive, online, offline）
	Status *string `json:"Status" name:"Status"`
}

type GetDataHistoryRequest struct {
	*tchttp.BaseRequest
	// 产品Id
	ProductId *string `json:"ProductId" name:"ProductId"`
	// 设备名称列表，允许最多一次100台
	DeviceNames []*string `json:"DeviceNames" name:"DeviceNames" list`
	// 查询开始时间
	StartTime *string `json:"StartTime" name:"StartTime"`
	// 查询结束时间
	EndTime *string `json:"EndTime" name:"EndTime"`
	// 查询数据量
	Size []*uint64 `json:"Size" name:"Size" list`
	// 时间排序（desc/asc）
	Order *string `json:"Order" name:"Order"`
	// 查询游标
	ScrollId *string `json:"ScrollId" name:"ScrollId"`
}

func (r *GetDataHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDataHistoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDataHistoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 数据历史
		DataHistory []*Object `json:"DataHistory" name:"DataHistory" list`
		// 查询游标
		ScrollId *string `json:"ScrollId" name:"ScrollId"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDataHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDataHistoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDeviceDataRequest struct {
	*tchttp.BaseRequest
	// 产品Id
	ProductId *string `json:"ProductId" name:"ProductId"`
	// 设备名称
	DeviceName *string `json:"DeviceName" name:"DeviceName"`
}

func (r *GetDeviceDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDeviceDataRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDeviceDataResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 设备数据
		DeviceData *Object `json:"DeviceData" name:"DeviceData"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDeviceDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDeviceDataResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDeviceLogRequest struct {
	*tchttp.BaseRequest
	// 产品Id
	ProductId *string `json:"ProductId" name:"ProductId"`
	// 设备名称列表，最大支持100台
	DeviceNames []*string `json:"DeviceNames" name:"DeviceNames" list`
	// 查询开始时间
	StartTime *string `json:"StartTime" name:"StartTime"`
	// 查询结束时间
	EndTime *string `json:"EndTime" name:"EndTime"`
	// 查询数据量
	Size *uint64 `json:"Size" name:"Size"`
	// 时间排序（desc/asc）
	Order *string `json:"Order" name:"Order"`
	// 查询游标
	ScrollId *string `json:"ScrollId" name:"ScrollId"`
	// 日志类型（comm/status）
	Type *string `json:"Type" name:"Type"`
}

func (r *GetDeviceLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDeviceLogRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDeviceLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 设备日志
		DeviceLog []*Object `json:"DeviceLog" name:"DeviceLog" list`
		// 查询游标
		ScrollId []*string `json:"ScrollId" name:"ScrollId" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDeviceLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDeviceLogResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDeviceRequest struct {
	*tchttp.BaseRequest
	// 产品Id
	ProductId *string `json:"ProductId" name:"ProductId"`
	// 设备名称
	DeviceName *string `json:"DeviceName" name:"DeviceName"`
}

func (r *GetDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 设备信息
		Device *Device `json:"Device" name:"Device"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDeviceStatusesRequest struct {
	*tchttp.BaseRequest
	// 产品ID
	ProductId *string `json:"ProductId" name:"ProductId"`
	// 设备名称列表（单次限制1000个设备）
	DeviceNames []*string `json:"DeviceNames" name:"DeviceNames" list`
}

func (r *GetDeviceStatusesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDeviceStatusesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDeviceStatusesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 设备状态列表
		DeviceStatuses []*DeviceStatus `json:"DeviceStatuses" name:"DeviceStatuses" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDeviceStatusesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDeviceStatusesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDevicesRequest struct {
	*tchttp.BaseRequest
	// 产品Id
	ProductId *string `json:"ProductId" name:"ProductId"`
	// 偏移
	Offset *uint64 `json:"Offset" name:"Offset"`
	// 长度
	Length *uint64 `json:"Length" name:"Length"`
}

func (r *GetDevicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDevicesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetDevicesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 设备列表
		Devices []*Device `json:"Devices" name:"Devices" list`
		// 设备总数
		Total *uint64 `json:"Total" name:"Total"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDevicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetDevicesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetProductRequest struct {
	*tchttp.BaseRequest
	// 产品Id
	ProductId *string `json:"ProductId" name:"ProductId"`
}

func (r *GetProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 产品信息
		Product *Product `json:"Product" name:"Product"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetProductsRequest struct {
	*tchttp.BaseRequest
	// 偏移
	Offset *uint64 `json:"Offset" name:"Offset"`
	// 长度
	Length *uint64 `json:"Length" name:"Length"`
}

func (r *GetProductsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetProductsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetProductsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// Product列表
		Products []*Product `json:"Products" name:"Products" list`
		// Product总数
		Total *uint64 `json:"Total" name:"Total"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetProductsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetProductsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetRuleRequest struct {
	*tchttp.BaseRequest
	// 规则Id
	RuleId *string `json:"RuleId" name:"RuleId"`
}

func (r *GetRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 规则
		Rule *Rule `json:"Rule" name:"Rule"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetRulesRequest struct {
	*tchttp.BaseRequest
	// 偏移
	Offset *uint64 `json:"Offset" name:"Offset"`
	// 长度
	Length *uint64 `json:"Length" name:"Length"`
}

func (r *GetRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetRulesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetRulesResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 规则列表
		Rules []*Rule `json:"Rules" name:"Rules" list`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetRulesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopicRequest struct {
	*tchttp.BaseRequest
	// TopicId
	TopicId *string `json:"TopicId" name:"TopicId"`
	// 产品Id
	ProductId *string `json:"ProductId" name:"ProductId"`
}

func (r *GetTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// Topic信息
		Topic *Topic `json:"Topic" name:"Topic"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopicResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopicsRequest struct {
	*tchttp.BaseRequest
	// 产品Id
	ProductId *string `json:"ProductId" name:"ProductId"`
	// 偏移
	Offset *uint64 `json:"Offset" name:"Offset"`
	// 长度
	Length *uint64 `json:"Length" name:"Length"`
}

func (r *GetTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopicsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetTopicsResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// Topic列表
		Topics []*Topic `json:"Topics" name:"Topics" list`
		// Topic总数
		Total *uint64 `json:"Total" name:"Total"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetTopicsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUserRequest struct {
	*tchttp.BaseRequest
}

func (r *GetUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUserRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUserResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 用户信息
		User *User `json:"User" name:"User"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUserResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IssueDeviceControlRequest struct {
	*tchttp.BaseRequest
	// 产品Id
	ProductId *string `json:"ProductId" name:"ProductId"`
	// 设备名称
	DeviceName *string `json:"DeviceName" name:"DeviceName"`
	// 控制数据（json）
	ControlData *string `json:"ControlData" name:"ControlData"`
}

func (r *IssueDeviceControlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IssueDeviceControlRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IssueDeviceControlResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *IssueDeviceControlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IssueDeviceControlResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Object struct {
}

type Product struct {
	// 产品Id
	ProductId *string `json:"ProductId" name:"ProductId"`
	// 产品Key
	ProductKey *string `json:"ProductKey" name:"ProductKey"`
	// 产品直连密钥
	ProductSecret *string `json:"ProductSecret" name:"ProductSecret"`
	// AppId
	AppId *uint64 `json:"AppId" name:"AppId"`
	// 产品名称
	Name *string `json:"Name" name:"Name"`
	// 产品描述
	Description *string `json:"Description" name:"Description"`
	// 连接域名
	Domain *string `json:"Domain" name:"Domain"`
	// 产品规格
	Standard *uint64 `json:"Standard" name:"Standard"`
	// 鉴权类型（0：直连，1：Token）
	AuthType *uint64 `json:"AuthType" name:"AuthType"`
	// 删除（0未删除）
	Deleted *uint64 `json:"Deleted" name:"Deleted"`
	// 备注
	Message *string `json:"Message" name:"Message"`
	// 创建时间
	CreateTime *string `json:"CreateTime" name:"CreateTime"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime" name:"UpdateTime"`
	// 数据模版
	DataTemplate *Object `json:"DataTemplate" name:"DataTemplate"`
}

type PublishMsgRequest struct {
	*tchttp.BaseRequest
	// Topic
	Topic *string `json:"Topic" name:"Topic"`
	// 消息内容
	Message *string `json:"Message" name:"Message"`
	// Qos(目前QoS支持0与1)
	Qos *int64 `json:"Qos" name:"Qos"`
}

func (r *PublishMsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PublishMsgRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PublishMsgResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *PublishMsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PublishMsgResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetDeviceRequest struct {
	*tchttp.BaseRequest
	// 产品Id
	ProductId *string `json:"ProductId" name:"ProductId"`
	// 设备名称
	DeviceName *string `json:"DeviceName" name:"DeviceName"`
}

func (r *ResetDeviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetDeviceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetDeviceResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetDeviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetDeviceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Rule struct {
	// 规则Id
	RuleId *string `json:"RuleId" name:"RuleId"`
	// AppId
	AppId *uint64 `json:"AppId" name:"AppId"`
	// 名称
	Name *string `json:"Name" name:"Name"`
	// 描述
	Description *string `json:"Description" name:"Description"`
	// 查询
	Query *RuleQuery `json:"Query" name:"Query"`
	// 转发
	Actions []*Object `json:"Actions" name:"Actions" list`
	// 已启动
	Active *uint64 `json:"Active" name:"Active"`
	// 已删除
	Deleted *uint64 `json:"Deleted" name:"Deleted"`
	// 创建时间
	CreateTime *string `json:"CreateTime" name:"CreateTime"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime" name:"UpdateTime"`
}

type RuleQuery struct {
	// 字段
	Field *string `json:"Field" name:"Field"`
	// Topic
	Topic *string `json:"Topic" name:"Topic"`
	// 过滤规则
	Condition *string `json:"Condition" name:"Condition"`
}

type Topic struct {
	// TopicId
	TopicId *string `json:"TopicId" name:"TopicId"`
	// Topic名称
	TopicName *string `json:"TopicName" name:"TopicName"`
	// 产品Id
	ProductId *string `json:"ProductId" name:"ProductId"`
	// 消息最大生命周期
	MsgLife *uint64 `json:"MsgLife" name:"MsgLife"`
	// 消息最大大小
	MsgSize *uint64 `json:"MsgSize" name:"MsgSize"`
	// 消息最大数量
	MsgCount *uint64 `json:"MsgCount" name:"MsgCount"`
	// 已删除
	Deleted *uint64 `json:"Deleted" name:"Deleted"`
	// Topic完整路径
	Path *string `json:"Path" name:"Path"`
	// 创建时间
	CreateTime *string `json:"CreateTime" name:"CreateTime"`
	// 更新时间
	UpdateTime *string `json:"UpdateTime" name:"UpdateTime"`
}

type UpdateProductRequest struct {
	*tchttp.BaseRequest
	// 产品Id
	ProductId *string `json:"ProductId" name:"ProductId"`
	// 产品名称
	Name *string `json:"Name" name:"Name"`
	// 产品描述
	Description *string `json:"Description" name:"Description"`
	// 数据模版（json）
	DataTemplate *string `json:"DataTemplate" name:"DataTemplate"`
}

func (r *UpdateProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 更新后的产品信息
		Product *Product `json:"Product" name:"Product"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateRuleRequest struct {
	*tchttp.BaseRequest
	// 规则Id
	RuleId *string `json:"RuleId" name:"RuleId"`
	// 名称
	Name *string `json:"Name" name:"Name"`
	// 描述
	Description *string `json:"Description" name:"Description"`
	// 查询
	Query *RuleQuery `json:"Query" name:"Query"`
	// 转发
	Actions []*Object `json:"Actions" name:"Actions" list`
}

func (r *UpdateRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateRuleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateRuleResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateRuleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type User struct {
	// app_id
	AppId *uint64 `json:"AppId" name:"AppId"`
	// 用户类型（1：国内，2：国际）
	Area *uint64 `json:"Area" name:"Area"`
	// 计费类型（日结、月结）
	BillingType *string `json:"BillingType" name:"BillingType"`
	// 用户状态（0：正常，1：欠费，2：恶意）
	Status *uint64 `json:"Status" name:"Status"`
	// 备注信息
	Message *string `json:"Message" name:"Message"`
	// 创建时间
	CreateTime *string `json:"CreateTime" name:"CreateTime"`
	// 修改时间
	UpdateTime *string `json:"UpdateTime" name:"UpdateTime"`
}
