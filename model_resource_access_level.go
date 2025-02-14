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

// ResourceAccessLevel # Access Level Object ### Description The `ResourceAccessLevel` object is used to represent the level of access that a user has to a resource or a resource has to a group. The \"default\" access level is a `ResourceAccessLevel` object whose fields are all empty strings.  ### Usage Example View the `ResourceAccessLevel` of a resource/user or resource/group pair to see the level of access granted to the resource.
type ResourceAccessLevel struct {
	// The human-readable name of the access level.
	AccessLevelName string `json:"access_level_name"`
	// The machine-readable identifier of the access level.
	AccessLevelRemoteId string `json:"access_level_remote_id"`
}

// NewResourceAccessLevel instantiates a new ResourceAccessLevel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceAccessLevel(accessLevelName string, accessLevelRemoteId string) *ResourceAccessLevel {
	this := ResourceAccessLevel{}
	this.AccessLevelName = accessLevelName
	this.AccessLevelRemoteId = accessLevelRemoteId
	return &this
}

// NewResourceAccessLevelWithDefaults instantiates a new ResourceAccessLevel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceAccessLevelWithDefaults() *ResourceAccessLevel {
	this := ResourceAccessLevel{}
	return &this
}

// GetAccessLevelName returns the AccessLevelName field value
func (o *ResourceAccessLevel) GetAccessLevelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessLevelName
}

// GetAccessLevelNameOk returns a tuple with the AccessLevelName field value
// and a boolean to check if the value has been set.
func (o *ResourceAccessLevel) GetAccessLevelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessLevelName, true
}

// SetAccessLevelName sets field value
func (o *ResourceAccessLevel) SetAccessLevelName(v string) {
	o.AccessLevelName = v
}

// GetAccessLevelRemoteId returns the AccessLevelRemoteId field value
func (o *ResourceAccessLevel) GetAccessLevelRemoteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessLevelRemoteId
}

// GetAccessLevelRemoteIdOk returns a tuple with the AccessLevelRemoteId field value
// and a boolean to check if the value has been set.
func (o *ResourceAccessLevel) GetAccessLevelRemoteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessLevelRemoteId, true
}

// SetAccessLevelRemoteId sets field value
func (o *ResourceAccessLevel) SetAccessLevelRemoteId(v string) {
	o.AccessLevelRemoteId = v
}

func (o ResourceAccessLevel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["access_level_name"] = o.AccessLevelName
	}
	if true {
		toSerialize["access_level_remote_id"] = o.AccessLevelRemoteId
	}
	return json.Marshal(toSerialize)
}

type NullableResourceAccessLevel struct {
	value *ResourceAccessLevel
	isSet bool
}

func (v NullableResourceAccessLevel) Get() *ResourceAccessLevel {
	return v.value
}

func (v *NullableResourceAccessLevel) Set(val *ResourceAccessLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceAccessLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceAccessLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceAccessLevel(val *ResourceAccessLevel) *NullableResourceAccessLevel {
	return &NullableResourceAccessLevel{value: val, isSet: true}
}

func (v NullableResourceAccessLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceAccessLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


