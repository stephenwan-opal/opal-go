/*
Opal API

Your Home For Developer Resources.

API version: 1.0
Contact: hello@opal.dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opal

import (
	"encoding/json"
)

// VisibilityInfo Visibility infomation of an entity.
type VisibilityInfo struct {
	Visibility VisibilityTypeEnum `json:"visibility"`
	VisibilityGroupIds []string `json:"visibility_group_ids,omitempty"`
}

// NewVisibilityInfo instantiates a new VisibilityInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVisibilityInfo(visibility VisibilityTypeEnum) *VisibilityInfo {
	this := VisibilityInfo{}
	this.Visibility = visibility
	return &this
}

// NewVisibilityInfoWithDefaults instantiates a new VisibilityInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVisibilityInfoWithDefaults() *VisibilityInfo {
	this := VisibilityInfo{}
	return &this
}

// GetVisibility returns the Visibility field value
func (o *VisibilityInfo) GetVisibility() VisibilityTypeEnum {
	if o == nil {
		var ret VisibilityTypeEnum
		return ret
	}

	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value
// and a boolean to check if the value has been set.
func (o *VisibilityInfo) GetVisibilityOk() (*VisibilityTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visibility, true
}

// SetVisibility sets field value
func (o *VisibilityInfo) SetVisibility(v VisibilityTypeEnum) {
	o.Visibility = v
}

// GetVisibilityGroupIds returns the VisibilityGroupIds field value if set, zero value otherwise.
func (o *VisibilityInfo) GetVisibilityGroupIds() []string {
	if o == nil || o.VisibilityGroupIds == nil {
		var ret []string
		return ret
	}
	return o.VisibilityGroupIds
}

// GetVisibilityGroupIdsOk returns a tuple with the VisibilityGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisibilityInfo) GetVisibilityGroupIdsOk() ([]string, bool) {
	if o == nil || o.VisibilityGroupIds == nil {
		return nil, false
	}
	return o.VisibilityGroupIds, true
}

// HasVisibilityGroupIds returns a boolean if a field has been set.
func (o *VisibilityInfo) HasVisibilityGroupIds() bool {
	if o != nil && o.VisibilityGroupIds != nil {
		return true
	}

	return false
}

// SetVisibilityGroupIds gets a reference to the given []string and assigns it to the VisibilityGroupIds field.
func (o *VisibilityInfo) SetVisibilityGroupIds(v []string) {
	o.VisibilityGroupIds = v
}

func (o VisibilityInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["visibility"] = o.Visibility
	}
	if o.VisibilityGroupIds != nil {
		toSerialize["visibility_group_ids"] = o.VisibilityGroupIds
	}
	return json.Marshal(toSerialize)
}

type NullableVisibilityInfo struct {
	value *VisibilityInfo
	isSet bool
}

func (v NullableVisibilityInfo) Get() *VisibilityInfo {
	return v.value
}

func (v *NullableVisibilityInfo) Set(val *VisibilityInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVisibilityInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVisibilityInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVisibilityInfo(val *VisibilityInfo) *NullableVisibilityInfo {
	return &NullableVisibilityInfo{value: val, isSet: true}
}

func (v NullableVisibilityInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVisibilityInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


