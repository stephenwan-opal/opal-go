# \GroupsApi

All URIs are relative to *https://api.opal.dev/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroup**](GroupsApi.md#CreateGroup) | **Post** /groups | 
[**DeleteGroup**](GroupsApi.md#DeleteGroup) | **Delete** /groups/{group_id} | 
[**GetGroupMessageChannels**](GroupsApi.md#GetGroupMessageChannels) | **Get** /groups/{group_id}/message-channels | 
[**GetGroupReviewers**](GroupsApi.md#GetGroupReviewers) | **Get** /groups/{group_id}/reviewers | 
[**GetGroupTags**](GroupsApi.md#GetGroupTags) | **Get** /groups/{group_id}/tags | 
[**GetGroupVisibility**](GroupsApi.md#GetGroupVisibility) | **Get** /groups/{group_id}/visibility | 
[**GetGroups**](GroupsApi.md#GetGroups) | **Get** /groups | 
[**SetGroupMessageChannels**](GroupsApi.md#SetGroupMessageChannels) | **Put** /groups/{group_id}/message-channels | 
[**SetGroupReviewers**](GroupsApi.md#SetGroupReviewers) | **Put** /groups/{group_id}/reviewers | 
[**SetGroupVisibility**](GroupsApi.md#SetGroupVisibility) | **Put** /groups/{group_id}/visibility | 
[**UpdateGroups**](GroupsApi.md#UpdateGroups) | **Put** /groups | 



## CreateGroup

> Group CreateGroup(ctx).CreateGroupInfo(createGroupInfo).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createGroupInfo := *openapiclient.NewCreateGroupInfo("mongo-db-prod", openapiclient.GroupTypeEnum("ACTIVE_DIRECTORY_GROUP"), "f454d283-ca87-4a8a-bdbb-df212eca5353") // CreateGroupInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.CreateGroup(context.Background()).CreateGroupInfo(createGroupInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.CreateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroup`: Group
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.CreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createGroupInfo** | [**CreateGroupInfo**](CreateGroupInfo.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> DeleteGroup(ctx, groupId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.DeleteGroup(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.DeleteGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupMessageChannels

> MessageChannelList GetGroupMessageChannels(ctx, groupId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroupMessageChannels(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupMessageChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupMessageChannels`: MessageChannelList
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupMessageChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupMessageChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageChannelList**](MessageChannelList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupReviewers

> []string GetGroupReviewers(ctx, groupId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroupReviewers(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupReviewers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupReviewers`: []string
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupReviewers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupReviewersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupTags

> TagsList GetGroupTags(ctx, groupId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    groupId := "1b978423-db0a-4037-a4cf-f79c60cb67b3" // string | The ID of the group whose tags to return.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroupTags(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupTags`: TagsList
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group whose tags to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TagsList**](TagsList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupVisibility

> VisibilityInfo GetGroupVisibility(ctx, groupId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroupVisibility(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroupVisibility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupVisibility`: VisibilityInfo
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroupVisibility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupVisibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VisibilityInfo**](VisibilityInfo.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroups

> PaginatedGroupsList GetGroups(ctx).Cursor(cursor).PageSize(pageSize).GroupFunctionFilter(groupFunctionFilter).GroupTypeFilter(groupTypeFilter).GroupIds(groupIds).GroupName(groupName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    pageSize := int32(200) // int32 | Number of results to return per page. Default is 200. (optional)
    groupFunctionFilter := openapiclient.GroupFunctionEnum("ON_CALL") // GroupFunctionEnum | The group function to filter by. (optional)
    groupTypeFilter := openapiclient.GroupTypeEnum("ACTIVE_DIRECTORY_GROUP") // GroupTypeEnum | The group type to filter by. (optional)
    groupIds := []string{"1b978423-db0a-4037-a4cf-f79c60cb67b3"} // []string | The group ids to filter by. (optional)
    groupName := "example-name" // string | Group name. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetGroups(context.Background()).Cursor(cursor).PageSize(pageSize).GroupFunctionFilter(groupFunctionFilter).GroupTypeFilter(groupTypeFilter).GroupIds(groupIds).GroupName(groupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroups`: PaginatedGroupsList
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | The pagination cursor value. | 
 **pageSize** | **int32** | Number of results to return per page. Default is 200. | 
 **groupFunctionFilter** | [**GroupFunctionEnum**](GroupFunctionEnum.md) | The group function to filter by. | 
 **groupTypeFilter** | [**GroupTypeEnum**](GroupTypeEnum.md) | The group type to filter by. | 
 **groupIds** | **[]string** | The group ids to filter by. | 
 **groupName** | **string** | Group name. | 

### Return type

[**PaginatedGroupsList**](PaginatedGroupsList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGroupMessageChannels

> []string SetGroupMessageChannels(ctx, groupId).MessageChannelIDList(messageChannelIDList).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.
    messageChannelIDList := *openapiclient.NewMessageChannelIDList([]string{"MessageChannelIds_example"}) // MessageChannelIDList | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.SetGroupMessageChannels(context.Background(), groupId).MessageChannelIDList(messageChannelIDList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.SetGroupMessageChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetGroupMessageChannels`: []string
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.SetGroupMessageChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetGroupMessageChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **messageChannelIDList** | [**MessageChannelIDList**](MessageChannelIDList.md) |  | 

### Return type

**[]string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGroupReviewers

> []string SetGroupReviewers(ctx, groupId).ReviewerIDList(reviewerIDList).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.
    reviewerIDList := *openapiclient.NewReviewerIDList([]string{"ReviewerIds_example"}) // ReviewerIDList | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.SetGroupReviewers(context.Background(), groupId).ReviewerIDList(reviewerIDList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.SetGroupReviewers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetGroupReviewers`: []string
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.SetGroupReviewers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetGroupReviewersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reviewerIDList** | [**ReviewerIDList**](ReviewerIDList.md) |  | 

### Return type

**[]string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGroupVisibility

> VisibilityInfo SetGroupVisibility(ctx, groupId).VisibilityInfo(visibilityInfo).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    groupId := "4baf8423-db0a-4037-a4cf-f79c60cb67a5" // string | The ID of the group.
    visibilityInfo := *openapiclient.NewVisibilityInfo(openapiclient.VisibilityTypeEnum("GLOBAL")) // VisibilityInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.SetGroupVisibility(context.Background(), groupId).VisibilityInfo(visibilityInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.SetGroupVisibility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetGroupVisibility`: VisibilityInfo
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.SetGroupVisibility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetGroupVisibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **visibilityInfo** | [**VisibilityInfo**](VisibilityInfo.md) |  | 

### Return type

[**VisibilityInfo**](VisibilityInfo.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroups

> UpdateGroupInfoList UpdateGroups(ctx).UpdateGroupInfoList(updateGroupInfoList).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    updateGroupInfoList := *openapiclient.NewUpdateGroupInfoList([]openapiclient.UpdateGroupInfo{*openapiclient.NewUpdateGroupInfo("f454d283-ca87-4a8a-bdbb-df212eca5353")}) // UpdateGroupInfoList | Groups to be updated

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.UpdateGroups(context.Background()).UpdateGroupInfoList(updateGroupInfoList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.UpdateGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroups`: UpdateGroupInfoList
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.UpdateGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateGroupInfoList** | [**UpdateGroupInfoList**](UpdateGroupInfoList.md) | Groups to be updated | 

### Return type

[**UpdateGroupInfoList**](UpdateGroupInfoList.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

