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
	"time"
)

// Session # Session Object ### Description The `Session` object is used to represent an access session. Some resources can be accessed temporarily via a time-bounded session.  ### Usage Example Fetch from the `LIST Sessions` endpoint.
type Session struct {
	// The ID of the connection.
	ConnectionId string `json:"connection_id"`
	// The ID of the user.
	UserId string `json:"user_id"`
	// The ID of the resource.
	ResourceId string `json:"resource_id"`
	AccessLevel ResourceAccessLevel `json:"access_level"`
	// The day and time the user's access will expire.
	ExpirationDate time.Time `json:"expiration_date"`
}

// NewSession instantiates a new Session object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSession(connectionId string, userId string, resourceId string, accessLevel ResourceAccessLevel, expirationDate time.Time) *Session {
	this := Session{}
	this.ConnectionId = connectionId
	this.UserId = userId
	this.ResourceId = resourceId
	this.AccessLevel = accessLevel
	this.ExpirationDate = expirationDate
	return &this
}

// NewSessionWithDefaults instantiates a new Session object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionWithDefaults() *Session {
	this := Session{}
	return &this
}

// GetConnectionId returns the ConnectionId field value
func (o *Session) GetConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *Session) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value
func (o *Session) SetConnectionId(v string) {
	o.ConnectionId = v
}

// GetUserId returns the UserId field value
func (o *Session) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *Session) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *Session) SetUserId(v string) {
	o.UserId = v
}

// GetResourceId returns the ResourceId field value
func (o *Session) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *Session) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *Session) SetResourceId(v string) {
	o.ResourceId = v
}

// GetAccessLevel returns the AccessLevel field value
func (o *Session) GetAccessLevel() ResourceAccessLevel {
	if o == nil {
		var ret ResourceAccessLevel
		return ret
	}

	return o.AccessLevel
}

// GetAccessLevelOk returns a tuple with the AccessLevel field value
// and a boolean to check if the value has been set.
func (o *Session) GetAccessLevelOk() (*ResourceAccessLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessLevel, true
}

// SetAccessLevel sets field value
func (o *Session) SetAccessLevel(v ResourceAccessLevel) {
	o.AccessLevel = v
}

// GetExpirationDate returns the ExpirationDate field value
func (o *Session) GetExpirationDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value
// and a boolean to check if the value has been set.
func (o *Session) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationDate, true
}

// SetExpirationDate sets field value
func (o *Session) SetExpirationDate(v time.Time) {
	o.ExpirationDate = v
}

func (o Session) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["connection_id"] = o.ConnectionId
	}
	if true {
		toSerialize["user_id"] = o.UserId
	}
	if true {
		toSerialize["resource_id"] = o.ResourceId
	}
	if true {
		toSerialize["access_level"] = o.AccessLevel
	}
	if true {
		toSerialize["expiration_date"] = o.ExpirationDate
	}
	return json.Marshal(toSerialize)
}

type NullableSession struct {
	value *Session
	isSet bool
}

func (v NullableSession) Get() *Session {
	return v.value
}

func (v *NullableSession) Set(val *Session) {
	v.value = val
	v.isSet = true
}

func (v NullableSession) IsSet() bool {
	return v.isSet
}

func (v *NullableSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSession(val *Session) *NullableSession {
	return &NullableSession{value: val, isSet: true}
}

func (v NullableSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


