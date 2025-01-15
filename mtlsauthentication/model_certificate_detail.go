/*
Saviynt mTLS Authentication

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mtlsauthentication

import (
	"encoding/json"
)

// checks if the CertificateDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateDetail{}

// CertificateDetail struct for CertificateDetail
type CertificateDetail struct {
	Alias *string `json:"alias,omitempty"`
	CertificateExpiry *string `json:"certificateExpiry,omitempty"`
	CertificateIssuerName *string `json:"certificateIssuerName,omitempty"`
	CertificateStatus *string `json:"certificateStatus,omitempty"`
	CertificateSubjectName *string `json:"certificateSubjectName,omitempty"`
	Thumbprints []CertificateThumbprint `json:"thumbprints,omitempty"`
}

// NewCertificateDetail instantiates a new CertificateDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateDetail() *CertificateDetail {
	this := CertificateDetail{}
	return &this
}

// NewCertificateDetailWithDefaults instantiates a new CertificateDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateDetailWithDefaults() *CertificateDetail {
	this := CertificateDetail{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *CertificateDetail) GetAlias() string {
	if o == nil || IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDetail) GetAliasOk() (*string, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *CertificateDetail) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *CertificateDetail) SetAlias(v string) {
	o.Alias = &v
}

// GetCertificateExpiry returns the CertificateExpiry field value if set, zero value otherwise.
func (o *CertificateDetail) GetCertificateExpiry() string {
	if o == nil || IsNil(o.CertificateExpiry) {
		var ret string
		return ret
	}
	return *o.CertificateExpiry
}

// GetCertificateExpiryOk returns a tuple with the CertificateExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDetail) GetCertificateExpiryOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateExpiry) {
		return nil, false
	}
	return o.CertificateExpiry, true
}

// HasCertificateExpiry returns a boolean if a field has been set.
func (o *CertificateDetail) HasCertificateExpiry() bool {
	if o != nil && !IsNil(o.CertificateExpiry) {
		return true
	}

	return false
}

// SetCertificateExpiry gets a reference to the given string and assigns it to the CertificateExpiry field.
func (o *CertificateDetail) SetCertificateExpiry(v string) {
	o.CertificateExpiry = &v
}

// GetCertificateIssuerName returns the CertificateIssuerName field value if set, zero value otherwise.
func (o *CertificateDetail) GetCertificateIssuerName() string {
	if o == nil || IsNil(o.CertificateIssuerName) {
		var ret string
		return ret
	}
	return *o.CertificateIssuerName
}

// GetCertificateIssuerNameOk returns a tuple with the CertificateIssuerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDetail) GetCertificateIssuerNameOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateIssuerName) {
		return nil, false
	}
	return o.CertificateIssuerName, true
}

// HasCertificateIssuerName returns a boolean if a field has been set.
func (o *CertificateDetail) HasCertificateIssuerName() bool {
	if o != nil && !IsNil(o.CertificateIssuerName) {
		return true
	}

	return false
}

// SetCertificateIssuerName gets a reference to the given string and assigns it to the CertificateIssuerName field.
func (o *CertificateDetail) SetCertificateIssuerName(v string) {
	o.CertificateIssuerName = &v
}

// GetCertificateStatus returns the CertificateStatus field value if set, zero value otherwise.
func (o *CertificateDetail) GetCertificateStatus() string {
	if o == nil || IsNil(o.CertificateStatus) {
		var ret string
		return ret
	}
	return *o.CertificateStatus
}

// GetCertificateStatusOk returns a tuple with the CertificateStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDetail) GetCertificateStatusOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateStatus) {
		return nil, false
	}
	return o.CertificateStatus, true
}

// HasCertificateStatus returns a boolean if a field has been set.
func (o *CertificateDetail) HasCertificateStatus() bool {
	if o != nil && !IsNil(o.CertificateStatus) {
		return true
	}

	return false
}

// SetCertificateStatus gets a reference to the given string and assigns it to the CertificateStatus field.
func (o *CertificateDetail) SetCertificateStatus(v string) {
	o.CertificateStatus = &v
}

// GetCertificateSubjectName returns the CertificateSubjectName field value if set, zero value otherwise.
func (o *CertificateDetail) GetCertificateSubjectName() string {
	if o == nil || IsNil(o.CertificateSubjectName) {
		var ret string
		return ret
	}
	return *o.CertificateSubjectName
}

// GetCertificateSubjectNameOk returns a tuple with the CertificateSubjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDetail) GetCertificateSubjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateSubjectName) {
		return nil, false
	}
	return o.CertificateSubjectName, true
}

// HasCertificateSubjectName returns a boolean if a field has been set.
func (o *CertificateDetail) HasCertificateSubjectName() bool {
	if o != nil && !IsNil(o.CertificateSubjectName) {
		return true
	}

	return false
}

// SetCertificateSubjectName gets a reference to the given string and assigns it to the CertificateSubjectName field.
func (o *CertificateDetail) SetCertificateSubjectName(v string) {
	o.CertificateSubjectName = &v
}

// GetThumbprints returns the Thumbprints field value if set, zero value otherwise.
func (o *CertificateDetail) GetThumbprints() []CertificateThumbprint {
	if o == nil || IsNil(o.Thumbprints) {
		var ret []CertificateThumbprint
		return ret
	}
	return o.Thumbprints
}

// GetThumbprintsOk returns a tuple with the Thumbprints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateDetail) GetThumbprintsOk() ([]CertificateThumbprint, bool) {
	if o == nil || IsNil(o.Thumbprints) {
		return nil, false
	}
	return o.Thumbprints, true
}

// HasThumbprints returns a boolean if a field has been set.
func (o *CertificateDetail) HasThumbprints() bool {
	if o != nil && !IsNil(o.Thumbprints) {
		return true
	}

	return false
}

// SetThumbprints gets a reference to the given []CertificateThumbprint and assigns it to the Thumbprints field.
func (o *CertificateDetail) SetThumbprints(v []CertificateThumbprint) {
	o.Thumbprints = v
}

func (o CertificateDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !IsNil(o.CertificateExpiry) {
		toSerialize["certificateExpiry"] = o.CertificateExpiry
	}
	if !IsNil(o.CertificateIssuerName) {
		toSerialize["certificateIssuerName"] = o.CertificateIssuerName
	}
	if !IsNil(o.CertificateStatus) {
		toSerialize["certificateStatus"] = o.CertificateStatus
	}
	if !IsNil(o.CertificateSubjectName) {
		toSerialize["certificateSubjectName"] = o.CertificateSubjectName
	}
	if !IsNil(o.Thumbprints) {
		toSerialize["thumbprints"] = o.Thumbprints
	}
	return toSerialize, nil
}

type NullableCertificateDetail struct {
	value *CertificateDetail
	isSet bool
}

func (v NullableCertificateDetail) Get() *CertificateDetail {
	return v.value
}

func (v *NullableCertificateDetail) Set(val *CertificateDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateDetail(val *CertificateDetail) *NullableCertificateDetail {
	return &NullableCertificateDetail{value: val, isSet: true}
}

func (v NullableCertificateDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


