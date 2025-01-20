/*
Saviynt Job Control API

Saviynt Job Control API Spec

API version: 1.0
Contact: https://github.com/saviynt
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jobcontrol

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ResumePauseJobsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResumePauseJobsRequest{}

// ResumePauseJobsRequest struct for ResumePauseJobsRequest
type ResumePauseJobsRequest struct {
	Action string `json:"action"`
	Triggername *string `json:"triggername,omitempty"`
	Jobname *string `json:"jobname,omitempty"`
}

type _ResumePauseJobsRequest ResumePauseJobsRequest

// NewResumePauseJobsRequest instantiates a new ResumePauseJobsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResumePauseJobsRequest(action string) *ResumePauseJobsRequest {
	this := ResumePauseJobsRequest{}
	this.Action = action
	return &this
}

// NewResumePauseJobsRequestWithDefaults instantiates a new ResumePauseJobsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResumePauseJobsRequestWithDefaults() *ResumePauseJobsRequest {
	this := ResumePauseJobsRequest{}
	return &this
}

// GetAction returns the Action field value
func (o *ResumePauseJobsRequest) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ResumePauseJobsRequest) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ResumePauseJobsRequest) SetAction(v string) {
	o.Action = v
}

// GetTriggername returns the Triggername field value if set, zero value otherwise.
func (o *ResumePauseJobsRequest) GetTriggername() string {
	if o == nil || IsNil(o.Triggername) {
		var ret string
		return ret
	}
	return *o.Triggername
}

// GetTriggernameOk returns a tuple with the Triggername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumePauseJobsRequest) GetTriggernameOk() (*string, bool) {
	if o == nil || IsNil(o.Triggername) {
		return nil, false
	}
	return o.Triggername, true
}

// HasTriggername returns a boolean if a field has been set.
func (o *ResumePauseJobsRequest) HasTriggername() bool {
	if o != nil && !IsNil(o.Triggername) {
		return true
	}

	return false
}

// SetTriggername gets a reference to the given string and assigns it to the Triggername field.
func (o *ResumePauseJobsRequest) SetTriggername(v string) {
	o.Triggername = &v
}

// GetJobname returns the Jobname field value if set, zero value otherwise.
func (o *ResumePauseJobsRequest) GetJobname() string {
	if o == nil || IsNil(o.Jobname) {
		var ret string
		return ret
	}
	return *o.Jobname
}

// GetJobnameOk returns a tuple with the Jobname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumePauseJobsRequest) GetJobnameOk() (*string, bool) {
	if o == nil || IsNil(o.Jobname) {
		return nil, false
	}
	return o.Jobname, true
}

// HasJobname returns a boolean if a field has been set.
func (o *ResumePauseJobsRequest) HasJobname() bool {
	if o != nil && !IsNil(o.Jobname) {
		return true
	}

	return false
}

// SetJobname gets a reference to the given string and assigns it to the Jobname field.
func (o *ResumePauseJobsRequest) SetJobname(v string) {
	o.Jobname = &v
}

func (o ResumePauseJobsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResumePauseJobsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	if !IsNil(o.Triggername) {
		toSerialize["triggername"] = o.Triggername
	}
	if !IsNil(o.Jobname) {
		toSerialize["jobname"] = o.Jobname
	}
	return toSerialize, nil
}

func (o *ResumePauseJobsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varResumePauseJobsRequest := _ResumePauseJobsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResumePauseJobsRequest)

	if err != nil {
		return err
	}

	*o = ResumePauseJobsRequest(varResumePauseJobsRequest)

	return err
}

type NullableResumePauseJobsRequest struct {
	value *ResumePauseJobsRequest
	isSet bool
}

func (v NullableResumePauseJobsRequest) Get() *ResumePauseJobsRequest {
	return v.value
}

func (v *NullableResumePauseJobsRequest) Set(val *ResumePauseJobsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableResumePauseJobsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableResumePauseJobsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResumePauseJobsRequest(val *ResumePauseJobsRequest) *NullableResumePauseJobsRequest {
	return &NullableResumePauseJobsRequest{value: val, isSet: true}
}

func (v NullableResumePauseJobsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResumePauseJobsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


