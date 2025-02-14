# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | The ID of the group. | 
**AppId** | Pointer to **string** | The ID of the group&#39;s app. | [optional] 
**Name** | Pointer to **string** | The name of the group. | [optional] 
**Description** | Pointer to **string** | A description of the group. | [optional] 
**AdminOwnerId** | Pointer to **string** | The ID of the owner of the group. | [optional] 
**RemoteId** | Pointer to **string** | The ID of the remote. | [optional] 
**RemoteName** | Pointer to **string** | The name of the remote. | [optional] 
**GroupFunction** | Pointer to [**GroupFunctionEnum**](GroupFunctionEnum.md) |  | [optional] 
**GroupType** | Pointer to [**GroupTypeEnum**](GroupTypeEnum.md) |  | [optional] 
**MaxDuration** | Pointer to **int32** | The maximum duration for which the group can be requested (in minutes). | [optional] 
**RequireManagerApproval** | Pointer to **bool** | A bool representing whether or not access requests to the group require manager approval. | [optional] 
**RequireSupportTicket** | Pointer to **bool** | A bool representing whether or not access requests to the group require a support ticket. | [optional] 
**FolderId** | Pointer to **string** | The ID of the folder that the group is located in. | [optional] 
**RequireMfaToApprove** | Pointer to **bool** | A bool representing whether or not to require MFA for reviewers to approve requests for this group. | [optional] 
**AutoApproval** | Pointer to **bool** | A bool representing whether or not to automatically approve requests to this group. | [optional] 
**RequestTemplateId** | Pointer to **string** | The ID of the associated request template. | [optional] 

## Methods

### NewGroup

`func NewGroup(groupId string, ) *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *Group) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Group) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Group) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetAppId

`func (o *Group) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *Group) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *Group) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *Group) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Group) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Group) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Group) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Group) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Group) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAdminOwnerId

`func (o *Group) GetAdminOwnerId() string`

GetAdminOwnerId returns the AdminOwnerId field if non-nil, zero value otherwise.

### GetAdminOwnerIdOk

`func (o *Group) GetAdminOwnerIdOk() (*string, bool)`

GetAdminOwnerIdOk returns a tuple with the AdminOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminOwnerId

`func (o *Group) SetAdminOwnerId(v string)`

SetAdminOwnerId sets AdminOwnerId field to given value.

### HasAdminOwnerId

`func (o *Group) HasAdminOwnerId() bool`

HasAdminOwnerId returns a boolean if a field has been set.

### GetRemoteId

`func (o *Group) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *Group) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *Group) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *Group) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetRemoteName

`func (o *Group) GetRemoteName() string`

GetRemoteName returns the RemoteName field if non-nil, zero value otherwise.

### GetRemoteNameOk

`func (o *Group) GetRemoteNameOk() (*string, bool)`

GetRemoteNameOk returns a tuple with the RemoteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteName

`func (o *Group) SetRemoteName(v string)`

SetRemoteName sets RemoteName field to given value.

### HasRemoteName

`func (o *Group) HasRemoteName() bool`

HasRemoteName returns a boolean if a field has been set.

### GetGroupFunction

`func (o *Group) GetGroupFunction() GroupFunctionEnum`

GetGroupFunction returns the GroupFunction field if non-nil, zero value otherwise.

### GetGroupFunctionOk

`func (o *Group) GetGroupFunctionOk() (*GroupFunctionEnum, bool)`

GetGroupFunctionOk returns a tuple with the GroupFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupFunction

`func (o *Group) SetGroupFunction(v GroupFunctionEnum)`

SetGroupFunction sets GroupFunction field to given value.

### HasGroupFunction

`func (o *Group) HasGroupFunction() bool`

HasGroupFunction returns a boolean if a field has been set.

### GetGroupType

`func (o *Group) GetGroupType() GroupTypeEnum`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *Group) GetGroupTypeOk() (*GroupTypeEnum, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *Group) SetGroupType(v GroupTypeEnum)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *Group) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### GetMaxDuration

`func (o *Group) GetMaxDuration() int32`

GetMaxDuration returns the MaxDuration field if non-nil, zero value otherwise.

### GetMaxDurationOk

`func (o *Group) GetMaxDurationOk() (*int32, bool)`

GetMaxDurationOk returns a tuple with the MaxDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDuration

`func (o *Group) SetMaxDuration(v int32)`

SetMaxDuration sets MaxDuration field to given value.

### HasMaxDuration

`func (o *Group) HasMaxDuration() bool`

HasMaxDuration returns a boolean if a field has been set.

### GetRequireManagerApproval

`func (o *Group) GetRequireManagerApproval() bool`

GetRequireManagerApproval returns the RequireManagerApproval field if non-nil, zero value otherwise.

### GetRequireManagerApprovalOk

`func (o *Group) GetRequireManagerApprovalOk() (*bool, bool)`

GetRequireManagerApprovalOk returns a tuple with the RequireManagerApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireManagerApproval

`func (o *Group) SetRequireManagerApproval(v bool)`

SetRequireManagerApproval sets RequireManagerApproval field to given value.

### HasRequireManagerApproval

`func (o *Group) HasRequireManagerApproval() bool`

HasRequireManagerApproval returns a boolean if a field has been set.

### GetRequireSupportTicket

`func (o *Group) GetRequireSupportTicket() bool`

GetRequireSupportTicket returns the RequireSupportTicket field if non-nil, zero value otherwise.

### GetRequireSupportTicketOk

`func (o *Group) GetRequireSupportTicketOk() (*bool, bool)`

GetRequireSupportTicketOk returns a tuple with the RequireSupportTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSupportTicket

`func (o *Group) SetRequireSupportTicket(v bool)`

SetRequireSupportTicket sets RequireSupportTicket field to given value.

### HasRequireSupportTicket

`func (o *Group) HasRequireSupportTicket() bool`

HasRequireSupportTicket returns a boolean if a field has been set.

### GetFolderId

`func (o *Group) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *Group) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *Group) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *Group) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetRequireMfaToApprove

`func (o *Group) GetRequireMfaToApprove() bool`

GetRequireMfaToApprove returns the RequireMfaToApprove field if non-nil, zero value otherwise.

### GetRequireMfaToApproveOk

`func (o *Group) GetRequireMfaToApproveOk() (*bool, bool)`

GetRequireMfaToApproveOk returns a tuple with the RequireMfaToApprove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMfaToApprove

`func (o *Group) SetRequireMfaToApprove(v bool)`

SetRequireMfaToApprove sets RequireMfaToApprove field to given value.

### HasRequireMfaToApprove

`func (o *Group) HasRequireMfaToApprove() bool`

HasRequireMfaToApprove returns a boolean if a field has been set.

### GetAutoApproval

`func (o *Group) GetAutoApproval() bool`

GetAutoApproval returns the AutoApproval field if non-nil, zero value otherwise.

### GetAutoApprovalOk

`func (o *Group) GetAutoApprovalOk() (*bool, bool)`

GetAutoApprovalOk returns a tuple with the AutoApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApproval

`func (o *Group) SetAutoApproval(v bool)`

SetAutoApproval sets AutoApproval field to given value.

### HasAutoApproval

`func (o *Group) HasAutoApproval() bool`

HasAutoApproval returns a boolean if a field has been set.

### GetRequestTemplateId

`func (o *Group) GetRequestTemplateId() string`

GetRequestTemplateId returns the RequestTemplateId field if non-nil, zero value otherwise.

### GetRequestTemplateIdOk

`func (o *Group) GetRequestTemplateIdOk() (*string, bool)`

GetRequestTemplateIdOk returns a tuple with the RequestTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTemplateId

`func (o *Group) SetRequestTemplateId(v string)`

SetRequestTemplateId sets RequestTemplateId field to given value.

### HasRequestTemplateId

`func (o *Group) HasRequestTemplateId() bool`

HasRequestTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


