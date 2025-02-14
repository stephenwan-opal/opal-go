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

// Resource # Resource Object ### Description The `Resource` object is used to represent a resource.  ### Usage Example Update from the `UPDATE Resources` endpoint.
type Resource struct {
	// The ID of the resource.
	ResourceId string `json:"resource_id"`
	// The ID of the app.
	AppId *string `json:"app_id,omitempty"`
	// The name of the resource.
	Name *string `json:"name,omitempty"`
	// A description of the resource.
	Description *string `json:"description,omitempty"`
	// The ID of the owner of the resource.
	AdminOwnerId *string `json:"admin_owner_id,omitempty"`
	// The ID of the resource on the remote system.
	RemoteResourceId *string `json:"remote_resource_id,omitempty"`
	// The name of the resource on the remote system.
	RemoteResourceName *string `json:"remote_resource_name,omitempty"`
	ResourceType *ResourceTypeEnum `json:"resource_type,omitempty"`
	// The maximum duration for which the resource can be requested (in minutes).
	MaxDuration *int32 `json:"max_duration,omitempty"`
	// A bool representing whether or not access requests to the resource require manager approval.
	RequireManagerApproval *bool `json:"require_manager_approval,omitempty"`
	// A bool representing whether or not access requests to the resource require a support ticket.
	RequireSupportTicket *bool `json:"require_support_ticket,omitempty"`
	// The ID of the folder that the resource is located in.
	FolderId *string `json:"folder_id,omitempty"`
	// A bool representing whether or not to require MFA for reviewers to approve requests for this resource.
	RequireMfaToApprove *bool `json:"require_mfa_to_approve,omitempty"`
	// A bool representing whether or not to automatically approve requests to this resource.
	AutoApproval *bool `json:"auto_approval,omitempty"`
	// The ID of the associated request template.
	RequestTemplateId *string `json:"request_template_id,omitempty"`
}

// NewResource instantiates a new Resource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResource(resourceId string) *Resource {
	this := Resource{}
	this.ResourceId = resourceId
	return &this
}

// NewResourceWithDefaults instantiates a new Resource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceWithDefaults() *Resource {
	this := Resource{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *Resource) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *Resource) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *Resource) SetResourceId(v string) {
	o.ResourceId = v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *Resource) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *Resource) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *Resource) SetAppId(v string) {
	o.AppId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Resource) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Resource) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Resource) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Resource) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Resource) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Resource) SetDescription(v string) {
	o.Description = &v
}

// GetAdminOwnerId returns the AdminOwnerId field value if set, zero value otherwise.
func (o *Resource) GetAdminOwnerId() string {
	if o == nil || o.AdminOwnerId == nil {
		var ret string
		return ret
	}
	return *o.AdminOwnerId
}

// GetAdminOwnerIdOk returns a tuple with the AdminOwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetAdminOwnerIdOk() (*string, bool) {
	if o == nil || o.AdminOwnerId == nil {
		return nil, false
	}
	return o.AdminOwnerId, true
}

// HasAdminOwnerId returns a boolean if a field has been set.
func (o *Resource) HasAdminOwnerId() bool {
	if o != nil && o.AdminOwnerId != nil {
		return true
	}

	return false
}

// SetAdminOwnerId gets a reference to the given string and assigns it to the AdminOwnerId field.
func (o *Resource) SetAdminOwnerId(v string) {
	o.AdminOwnerId = &v
}

// GetRemoteResourceId returns the RemoteResourceId field value if set, zero value otherwise.
func (o *Resource) GetRemoteResourceId() string {
	if o == nil || o.RemoteResourceId == nil {
		var ret string
		return ret
	}
	return *o.RemoteResourceId
}

// GetRemoteResourceIdOk returns a tuple with the RemoteResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetRemoteResourceIdOk() (*string, bool) {
	if o == nil || o.RemoteResourceId == nil {
		return nil, false
	}
	return o.RemoteResourceId, true
}

// HasRemoteResourceId returns a boolean if a field has been set.
func (o *Resource) HasRemoteResourceId() bool {
	if o != nil && o.RemoteResourceId != nil {
		return true
	}

	return false
}

// SetRemoteResourceId gets a reference to the given string and assigns it to the RemoteResourceId field.
func (o *Resource) SetRemoteResourceId(v string) {
	o.RemoteResourceId = &v
}

// GetRemoteResourceName returns the RemoteResourceName field value if set, zero value otherwise.
func (o *Resource) GetRemoteResourceName() string {
	if o == nil || o.RemoteResourceName == nil {
		var ret string
		return ret
	}
	return *o.RemoteResourceName
}

// GetRemoteResourceNameOk returns a tuple with the RemoteResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetRemoteResourceNameOk() (*string, bool) {
	if o == nil || o.RemoteResourceName == nil {
		return nil, false
	}
	return o.RemoteResourceName, true
}

// HasRemoteResourceName returns a boolean if a field has been set.
func (o *Resource) HasRemoteResourceName() bool {
	if o != nil && o.RemoteResourceName != nil {
		return true
	}

	return false
}

// SetRemoteResourceName gets a reference to the given string and assigns it to the RemoteResourceName field.
func (o *Resource) SetRemoteResourceName(v string) {
	o.RemoteResourceName = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *Resource) GetResourceType() ResourceTypeEnum {
	if o == nil || o.ResourceType == nil {
		var ret ResourceTypeEnum
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetResourceTypeOk() (*ResourceTypeEnum, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *Resource) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given ResourceTypeEnum and assigns it to the ResourceType field.
func (o *Resource) SetResourceType(v ResourceTypeEnum) {
	o.ResourceType = &v
}

// GetMaxDuration returns the MaxDuration field value if set, zero value otherwise.
func (o *Resource) GetMaxDuration() int32 {
	if o == nil || o.MaxDuration == nil {
		var ret int32
		return ret
	}
	return *o.MaxDuration
}

// GetMaxDurationOk returns a tuple with the MaxDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetMaxDurationOk() (*int32, bool) {
	if o == nil || o.MaxDuration == nil {
		return nil, false
	}
	return o.MaxDuration, true
}

// HasMaxDuration returns a boolean if a field has been set.
func (o *Resource) HasMaxDuration() bool {
	if o != nil && o.MaxDuration != nil {
		return true
	}

	return false
}

// SetMaxDuration gets a reference to the given int32 and assigns it to the MaxDuration field.
func (o *Resource) SetMaxDuration(v int32) {
	o.MaxDuration = &v
}

// GetRequireManagerApproval returns the RequireManagerApproval field value if set, zero value otherwise.
func (o *Resource) GetRequireManagerApproval() bool {
	if o == nil || o.RequireManagerApproval == nil {
		var ret bool
		return ret
	}
	return *o.RequireManagerApproval
}

// GetRequireManagerApprovalOk returns a tuple with the RequireManagerApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetRequireManagerApprovalOk() (*bool, bool) {
	if o == nil || o.RequireManagerApproval == nil {
		return nil, false
	}
	return o.RequireManagerApproval, true
}

// HasRequireManagerApproval returns a boolean if a field has been set.
func (o *Resource) HasRequireManagerApproval() bool {
	if o != nil && o.RequireManagerApproval != nil {
		return true
	}

	return false
}

// SetRequireManagerApproval gets a reference to the given bool and assigns it to the RequireManagerApproval field.
func (o *Resource) SetRequireManagerApproval(v bool) {
	o.RequireManagerApproval = &v
}

// GetRequireSupportTicket returns the RequireSupportTicket field value if set, zero value otherwise.
func (o *Resource) GetRequireSupportTicket() bool {
	if o == nil || o.RequireSupportTicket == nil {
		var ret bool
		return ret
	}
	return *o.RequireSupportTicket
}

// GetRequireSupportTicketOk returns a tuple with the RequireSupportTicket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetRequireSupportTicketOk() (*bool, bool) {
	if o == nil || o.RequireSupportTicket == nil {
		return nil, false
	}
	return o.RequireSupportTicket, true
}

// HasRequireSupportTicket returns a boolean if a field has been set.
func (o *Resource) HasRequireSupportTicket() bool {
	if o != nil && o.RequireSupportTicket != nil {
		return true
	}

	return false
}

// SetRequireSupportTicket gets a reference to the given bool and assigns it to the RequireSupportTicket field.
func (o *Resource) SetRequireSupportTicket(v bool) {
	o.RequireSupportTicket = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *Resource) GetFolderId() string {
	if o == nil || o.FolderId == nil {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetFolderIdOk() (*string, bool) {
	if o == nil || o.FolderId == nil {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *Resource) HasFolderId() bool {
	if o != nil && o.FolderId != nil {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *Resource) SetFolderId(v string) {
	o.FolderId = &v
}

// GetRequireMfaToApprove returns the RequireMfaToApprove field value if set, zero value otherwise.
func (o *Resource) GetRequireMfaToApprove() bool {
	if o == nil || o.RequireMfaToApprove == nil {
		var ret bool
		return ret
	}
	return *o.RequireMfaToApprove
}

// GetRequireMfaToApproveOk returns a tuple with the RequireMfaToApprove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetRequireMfaToApproveOk() (*bool, bool) {
	if o == nil || o.RequireMfaToApprove == nil {
		return nil, false
	}
	return o.RequireMfaToApprove, true
}

// HasRequireMfaToApprove returns a boolean if a field has been set.
func (o *Resource) HasRequireMfaToApprove() bool {
	if o != nil && o.RequireMfaToApprove != nil {
		return true
	}

	return false
}

// SetRequireMfaToApprove gets a reference to the given bool and assigns it to the RequireMfaToApprove field.
func (o *Resource) SetRequireMfaToApprove(v bool) {
	o.RequireMfaToApprove = &v
}

// GetAutoApproval returns the AutoApproval field value if set, zero value otherwise.
func (o *Resource) GetAutoApproval() bool {
	if o == nil || o.AutoApproval == nil {
		var ret bool
		return ret
	}
	return *o.AutoApproval
}

// GetAutoApprovalOk returns a tuple with the AutoApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetAutoApprovalOk() (*bool, bool) {
	if o == nil || o.AutoApproval == nil {
		return nil, false
	}
	return o.AutoApproval, true
}

// HasAutoApproval returns a boolean if a field has been set.
func (o *Resource) HasAutoApproval() bool {
	if o != nil && o.AutoApproval != nil {
		return true
	}

	return false
}

// SetAutoApproval gets a reference to the given bool and assigns it to the AutoApproval field.
func (o *Resource) SetAutoApproval(v bool) {
	o.AutoApproval = &v
}

// GetRequestTemplateId returns the RequestTemplateId field value if set, zero value otherwise.
func (o *Resource) GetRequestTemplateId() string {
	if o == nil || o.RequestTemplateId == nil {
		var ret string
		return ret
	}
	return *o.RequestTemplateId
}

// GetRequestTemplateIdOk returns a tuple with the RequestTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetRequestTemplateIdOk() (*string, bool) {
	if o == nil || o.RequestTemplateId == nil {
		return nil, false
	}
	return o.RequestTemplateId, true
}

// HasRequestTemplateId returns a boolean if a field has been set.
func (o *Resource) HasRequestTemplateId() bool {
	if o != nil && o.RequestTemplateId != nil {
		return true
	}

	return false
}

// SetRequestTemplateId gets a reference to the given string and assigns it to the RequestTemplateId field.
func (o *Resource) SetRequestTemplateId(v string) {
	o.RequestTemplateId = &v
}

func (o Resource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resource_id"] = o.ResourceId
	}
	if o.AppId != nil {
		toSerialize["app_id"] = o.AppId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.AdminOwnerId != nil {
		toSerialize["admin_owner_id"] = o.AdminOwnerId
	}
	if o.RemoteResourceId != nil {
		toSerialize["remote_resource_id"] = o.RemoteResourceId
	}
	if o.RemoteResourceName != nil {
		toSerialize["remote_resource_name"] = o.RemoteResourceName
	}
	if o.ResourceType != nil {
		toSerialize["resource_type"] = o.ResourceType
	}
	if o.MaxDuration != nil {
		toSerialize["max_duration"] = o.MaxDuration
	}
	if o.RequireManagerApproval != nil {
		toSerialize["require_manager_approval"] = o.RequireManagerApproval
	}
	if o.RequireSupportTicket != nil {
		toSerialize["require_support_ticket"] = o.RequireSupportTicket
	}
	if o.FolderId != nil {
		toSerialize["folder_id"] = o.FolderId
	}
	if o.RequireMfaToApprove != nil {
		toSerialize["require_mfa_to_approve"] = o.RequireMfaToApprove
	}
	if o.AutoApproval != nil {
		toSerialize["auto_approval"] = o.AutoApproval
	}
	if o.RequestTemplateId != nil {
		toSerialize["request_template_id"] = o.RequestTemplateId
	}
	return json.Marshal(toSerialize)
}

type NullableResource struct {
	value *Resource
	isSet bool
}

func (v NullableResource) Get() *Resource {
	return v.value
}

func (v *NullableResource) Set(val *Resource) {
	v.value = val
	v.isSet = true
}

func (v NullableResource) IsSet() bool {
	return v.isSet
}

func (v *NullableResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResource(val *Resource) *NullableResource {
	return &NullableResource{value: val, isSet: true}
}

func (v NullableResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


