/*
Saviynt SAV Roles API

Saviynt SAV Roles API

API version: 1.0
Contact: https://github.com/saviynt
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package savroles

import (
	"encoding/json"
)

// checks if the GetSAVRoleUsersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSAVRoleUsersResponse{}

// GetSAVRoleUsersResponse struct for GetSAVRoleUsersResponse
type GetSAVRoleUsersResponse struct {
	Users []User `json:"users,omitempty"`
}

// NewGetSAVRoleUsersResponse instantiates a new GetSAVRoleUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSAVRoleUsersResponse() *GetSAVRoleUsersResponse {
	this := GetSAVRoleUsersResponse{}
	return &this
}

// NewGetSAVRoleUsersResponseWithDefaults instantiates a new GetSAVRoleUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSAVRoleUsersResponseWithDefaults() *GetSAVRoleUsersResponse {
	this := GetSAVRoleUsersResponse{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *GetSAVRoleUsersResponse) GetUsers() []User {
	if o == nil || IsNil(o.Users) {
		var ret []User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSAVRoleUsersResponse) GetUsersOk() ([]User, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *GetSAVRoleUsersResponse) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []User and assigns it to the Users field.
func (o *GetSAVRoleUsersResponse) SetUsers(v []User) {
	o.Users = v
}

func (o GetSAVRoleUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSAVRoleUsersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableGetSAVRoleUsersResponse struct {
	value *GetSAVRoleUsersResponse
	isSet bool
}

func (v NullableGetSAVRoleUsersResponse) Get() *GetSAVRoleUsersResponse {
	return v.value
}

func (v *NullableGetSAVRoleUsersResponse) Set(val *GetSAVRoleUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSAVRoleUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSAVRoleUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSAVRoleUsersResponse(val *GetSAVRoleUsersResponse) *NullableGetSAVRoleUsersResponse {
	return &NullableGetSAVRoleUsersResponse{value: val, isSet: true}
}

func (v NullableGetSAVRoleUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSAVRoleUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


