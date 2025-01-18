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

// checks if the ObjectsToExport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectsToExport{}

// ObjectsToExport struct for ObjectsToExport
type ObjectsToExport struct {
	SavRoles []string `json:"savRoles,omitempty"`
	EmailTemplate []string `json:"emailTemplate,omitempty"`
}

// NewObjectsToExport instantiates a new ObjectsToExport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectsToExport() *ObjectsToExport {
	this := ObjectsToExport{}
	return &this
}

// NewObjectsToExportWithDefaults instantiates a new ObjectsToExport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectsToExportWithDefaults() *ObjectsToExport {
	this := ObjectsToExport{}
	return &this
}

// GetSavRoles returns the SavRoles field value if set, zero value otherwise.
func (o *ObjectsToExport) GetSavRoles() []string {
	if o == nil || IsNil(o.SavRoles) {
		var ret []string
		return ret
	}
	return o.SavRoles
}

// GetSavRolesOk returns a tuple with the SavRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectsToExport) GetSavRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.SavRoles) {
		return nil, false
	}
	return o.SavRoles, true
}

// HasSavRoles returns a boolean if a field has been set.
func (o *ObjectsToExport) HasSavRoles() bool {
	if o != nil && !IsNil(o.SavRoles) {
		return true
	}

	return false
}

// SetSavRoles gets a reference to the given []string and assigns it to the SavRoles field.
func (o *ObjectsToExport) SetSavRoles(v []string) {
	o.SavRoles = v
}

// GetEmailTemplate returns the EmailTemplate field value if set, zero value otherwise.
func (o *ObjectsToExport) GetEmailTemplate() []string {
	if o == nil || IsNil(o.EmailTemplate) {
		var ret []string
		return ret
	}
	return o.EmailTemplate
}

// GetEmailTemplateOk returns a tuple with the EmailTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectsToExport) GetEmailTemplateOk() ([]string, bool) {
	if o == nil || IsNil(o.EmailTemplate) {
		return nil, false
	}
	return o.EmailTemplate, true
}

// HasEmailTemplate returns a boolean if a field has been set.
func (o *ObjectsToExport) HasEmailTemplate() bool {
	if o != nil && !IsNil(o.EmailTemplate) {
		return true
	}

	return false
}

// SetEmailTemplate gets a reference to the given []string and assigns it to the EmailTemplate field.
func (o *ObjectsToExport) SetEmailTemplate(v []string) {
	o.EmailTemplate = v
}

func (o ObjectsToExport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectsToExport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SavRoles) {
		toSerialize["savRoles"] = o.SavRoles
	}
	if !IsNil(o.EmailTemplate) {
		toSerialize["emailTemplate"] = o.EmailTemplate
	}
	return toSerialize, nil
}

type NullableObjectsToExport struct {
	value *ObjectsToExport
	isSet bool
}

func (v NullableObjectsToExport) Get() *ObjectsToExport {
	return v.value
}

func (v *NullableObjectsToExport) Set(val *ObjectsToExport) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectsToExport) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectsToExport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectsToExport(val *ObjectsToExport) *NullableObjectsToExport {
	return &NullableObjectsToExport{value: val, isSet: true}
}

func (v NullableObjectsToExport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectsToExport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


