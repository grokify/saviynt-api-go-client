/*
Saviynt Transport API

Transporting Packages: https://docs.saviyntcloud.com/bundle/EIC-Admin-AMS/page/Content/Chapter07-General-Administrator/Transporting-Packages.htm

API version: 1.0
Contact: https://github.com/saviynt
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transport

import (
	"encoding/json"
)

// checks if the TransportPackageStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransportPackageStatusResponse{}

// TransportPackageStatusResponse struct for TransportPackageStatusResponse
type TransportPackageStatusResponse struct {
	Msg *string `json:"msg,omitempty"`
	MsgDescription *string `json:"msgDescription,omitempty"`
	ErrorCode *int32 `json:"errorCode,omitempty"`
}

// NewTransportPackageStatusResponse instantiates a new TransportPackageStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportPackageStatusResponse() *TransportPackageStatusResponse {
	this := TransportPackageStatusResponse{}
	return &this
}

// NewTransportPackageStatusResponseWithDefaults instantiates a new TransportPackageStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportPackageStatusResponseWithDefaults() *TransportPackageStatusResponse {
	this := TransportPackageStatusResponse{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *TransportPackageStatusResponse) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportPackageStatusResponse) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *TransportPackageStatusResponse) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *TransportPackageStatusResponse) SetMsg(v string) {
	o.Msg = &v
}

// GetMsgDescription returns the MsgDescription field value if set, zero value otherwise.
func (o *TransportPackageStatusResponse) GetMsgDescription() string {
	if o == nil || IsNil(o.MsgDescription) {
		var ret string
		return ret
	}
	return *o.MsgDescription
}

// GetMsgDescriptionOk returns a tuple with the MsgDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportPackageStatusResponse) GetMsgDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.MsgDescription) {
		return nil, false
	}
	return o.MsgDescription, true
}

// HasMsgDescription returns a boolean if a field has been set.
func (o *TransportPackageStatusResponse) HasMsgDescription() bool {
	if o != nil && !IsNil(o.MsgDescription) {
		return true
	}

	return false
}

// SetMsgDescription gets a reference to the given string and assigns it to the MsgDescription field.
func (o *TransportPackageStatusResponse) SetMsgDescription(v string) {
	o.MsgDescription = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *TransportPackageStatusResponse) GetErrorCode() int32 {
	if o == nil || IsNil(o.ErrorCode) {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportPackageStatusResponse) GetErrorCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *TransportPackageStatusResponse) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *TransportPackageStatusResponse) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

func (o TransportPackageStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransportPackageStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !IsNil(o.MsgDescription) {
		toSerialize["msgDescription"] = o.MsgDescription
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	return toSerialize, nil
}

type NullableTransportPackageStatusResponse struct {
	value *TransportPackageStatusResponse
	isSet bool
}

func (v NullableTransportPackageStatusResponse) Get() *TransportPackageStatusResponse {
	return v.value
}

func (v *NullableTransportPackageStatusResponse) Set(val *TransportPackageStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportPackageStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportPackageStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportPackageStatusResponse(val *TransportPackageStatusResponse) *NullableTransportPackageStatusResponse {
	return &NullableTransportPackageStatusResponse{value: val, isSet: true}
}

func (v NullableTransportPackageStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportPackageStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


