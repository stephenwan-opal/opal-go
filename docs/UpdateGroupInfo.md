# UpdateGroupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | The ID of the group. | 
**Name** | Pointer to **string** | The name of the group. | [optional] 
**Description** | Pointer to **string** | A description of the group. | [optional] 
**AdminOwnerId** | Pointer to **string** | The ID of the owner of the group. | [optional] 
**MaxDuration** | Pointer to **int32** | The maximum duration for which the group can be requested (in minutes). Use 0 to set to indefinite. | [optional] 
**RequireManagerApproval** | Pointer to **bool** | A bool representing whether or not access requests to the group require manager approval. | [optional] 
**RequireSupportTicket** | Pointer to **bool** | A bool representing whether or not access requests to the group require a support ticket. | [optional] 
**FolderId** | Pointer to **string** | The ID of the folder that the group is located in. | [optional] 
**RequireMfaToApprove** | Pointer to **bool** | A bool representing whether or not to require MFA for reviewers to approve requests for this group. | [optional] 
**AutoApproval** | Pointer to **bool** | A bool representing whether or not to automatically approve requests to this group. | [optional] 
**RequestTemplateId** | Pointer to **string** | The ID of the associated request template. | [optional] 

## Methods

### NewUpdateGroupInfo

`func NewUpdateGroupInfo(groupId string, ) *UpdateGroupInfo`

NewUpdateGroupInfo instantiates a new UpdateGroupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGroupInfoWithDefaults

`func NewUpdateGroupInfoWithDefaults() *UpdateGroupInfo`

NewUpdateGroupInfoWithDefaults instantiates a new UpdateGroupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *UpdateGroupInfo) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UpdateGroupInfo) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UpdateGroupInfo) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetName

`func (o *UpdateGroupInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGroupInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGroupInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateGroupInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateGroupInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateGroupInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateGroupInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateGroupInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAdminOwnerId

`func (o *UpdateGroupInfo) GetAdminOwnerId() string`

GetAdminOwnerId returns the AdminOwnerId field if non-nil, zero value otherwise.

### GetAdminOwnerIdOk

`func (o *UpdateGroupInfo) GetAdminOwnerIdOk() (*string, bool)`

GetAdminOwnerIdOk returns a tuple with the AdminOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminOwnerId

`func (o *UpdateGroupInfo) SetAdminOwnerId(v string)`

SetAdminOwnerId sets AdminOwnerId field to given value.

### HasAdminOwnerId

`func (o *UpdateGroupInfo) HasAdminOwnerId() bool`

HasAdminOwnerId returns a boolean if a field has been set.

### GetMaxDuration

`func (o *UpdateGroupInfo) GetMaxDuration() int32`

GetMaxDuration returns the MaxDuration field if non-nil, zero value otherwise.

### GetMaxDurationOk

`func (o *UpdateGroupInfo) GetMaxDurationOk() (*int32, bool)`

GetMaxDurationOk returns a tuple with the MaxDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDuration

`func (o *UpdateGroupInfo) SetMaxDuration(v int32)`

SetMaxDuration sets MaxDuration field to given value.

### HasMaxDuration

`func (o *UpdateGroupInfo) HasMaxDuration() bool`

HasMaxDuration returns a boolean if a field has been set.

### GetRequireManagerApproval

`func (o *UpdateGroupInfo) GetRequireManagerApproval() bool`

GetRequireManagerApproval returns the RequireManagerApproval field if non-nil, zero value otherwise.

### GetRequireManagerApprovalOk

`func (o *UpdateGroupInfo) GetRequireManagerApprovalOk() (*bool, bool)`

GetRequireManagerApprovalOk returns a tuple with the RequireManagerApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireManagerApproval

`func (o *UpdateGroupInfo) SetRequireManagerApproval(v bool)`

SetRequireManagerApproval sets RequireManagerApproval field to given value.

### HasRequireManagerApproval

`func (o *UpdateGroupInfo) HasRequireManagerApproval() bool`

HasRequireManagerApproval returns a boolean if a field has been set.

### GetRequireSupportTicket

`func (o *UpdateGroupInfo) GetRequireSupportTicket() bool`

GetRequireSupportTicket returns the RequireSupportTicket field if non-nil, zero value otherwise.

### GetRequireSupportTicketOk

`func (o *UpdateGroupInfo) GetRequireSupportTicketOk() (*bool, bool)`

GetRequireSupportTicketOk returns a tuple with the RequireSupportTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSupportTicket

`func (o *UpdateGroupInfo) SetRequireSupportTicket(v bool)`

SetRequireSupportTicket sets RequireSupportTicket field to given value.

### HasRequireSupportTicket

`func (o *UpdateGroupInfo) HasRequireSupportTicket() bool`

HasRequireSupportTicket returns a boolean if a field has been set.

### GetFolderId

`func (o *UpdateGroupInfo) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *UpdateGroupInfo) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *UpdateGroupInfo) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *UpdateGroupInfo) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetRequireMfaToApprove

`func (o *UpdateGroupInfo) GetRequireMfaToApprove() bool`

GetRequireMfaToApprove returns the RequireMfaToApprove field if non-nil, zero value otherwise.

### GetRequireMfaToApproveOk

`func (o *UpdateGroupInfo) GetRequireMfaToApproveOk() (*bool, bool)`

GetRequireMfaToApproveOk returns a tuple with the RequireMfaToApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMfaToApprove

`func (o *UpdateGroupInfo) SetRequireMfaToApprove(v bool)`

SetRequireMfaToApprove sets RequireMfaToApprove field to given value.

### HasRequireMfaToApprove

`func (o *UpdateGroupInfo) HasRequireMfaToApprove() bool`

HasRequireMfaToApprove returns a boolean if a field has been set.

### GetAutoApproval

`func (o *UpdateGroupInfo) GetAutoApproval() bool`

GetAutoApproval returns the AutoApproval field if non-nil, zero value otherwise.

### GetAutoApprovalOk

`func (o *UpdateGroupInfo) GetAutoApprovalOk() (*bool, bool)`

GetAutoApprovalOk returns a tuple with the AutoApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproval

`func (o *UpdateGroupInfo) SetAutoApproval(v bool)`

SetAutoApproval sets AutoApproval field to given value.

### HasAutoApproval

`func (o *UpdateGroupInfo) HasAutoApproval() bool`

HasAutoApproval returns a boolean if a field has been set.

### GetRequestTemplateId

`func (o *UpdateGroupInfo) GetRequestTemplateId() string`

GetRequestTemplateId returns the RequestTemplateId field if non-nil, zero value otherwise.

### GetRequestTemplateIdOk

`func (o *UpdateGroupInfo) GetRequestTemplateIdOk() (*string, bool)`

GetRequestTemplateIdOk returns a tuple with the RequestTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTemplateId

`func (o *UpdateGroupInfo) SetRequestTemplateId(v string)`

SetRequestTemplateId sets RequestTemplateId field to given value.

### HasRequestTemplateId

`func (o *UpdateGroupInfo) HasRequestTemplateId() bool`

HasRequestTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


