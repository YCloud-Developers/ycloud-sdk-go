# WhatsappGroupsApi

All URIs are relative to *https://api.ycloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveJoinRequests**](WhatsappGroupsApi.md#ApproveJoinRequests) | **Post** /whatsapp/{businessPhoneNumber}/groups/{groupId}/joinRequests/approve | Approve group join requests
[**Create**](WhatsappGroupsApi.md#Create) | **Post** /whatsapp/{businessPhoneNumber}/groups | Create a group
[**Delete**](WhatsappGroupsApi.md#Delete) | **Delete** /whatsapp/{businessPhoneNumber}/groups/{groupId} | Delete a group
[**List**](WhatsappGroupsApi.md#List) | **Get** /whatsapp/{businessPhoneNumber}/groups | List groups
[**ListJoinRequests**](WhatsappGroupsApi.md#ListJoinRequests) | **Get** /whatsapp/{businessPhoneNumber}/groups/{groupId}/joinRequests | List group join requests
[**RejectJoinRequests**](WhatsappGroupsApi.md#RejectJoinRequests) | **Post** /whatsapp/{businessPhoneNumber}/groups/{groupId}/joinRequests/reject | Reject group join requests
[**RemoveParticipants**](WhatsappGroupsApi.md#RemoveParticipants) | **Post** /whatsapp/groups/{groupId}/participants/remove | Remove group participants
[**ResetInviteLink**](WhatsappGroupsApi.md#ResetInviteLink) | **Post** /whatsapp/{businessPhoneNumber}/groups/{groupId}/inviteLink/reset | Reset a group invite link
[**Retrieve**](WhatsappGroupsApi.md#Retrieve) | **Get** /whatsapp/{businessPhoneNumber}/groups/{groupId} | Retrieve a group
[**RetrieveInviteLink**](WhatsappGroupsApi.md#RetrieveInviteLink) | **Get** /whatsapp/{businessPhoneNumber}/groups/{groupId}/inviteLink | Retrieve a group invite link
[**SendInviteLinkMessage**](WhatsappGroupsApi.md#SendInviteLinkMessage) | **Post** /whatsapp/{businessPhoneNumber}/groups/inviteLink/messages | Send a group invite link message
[**UpdateSettings**](WhatsappGroupsApi.md#UpdateSettings) | **Patch** /whatsapp/{businessPhoneNumber}/groups/{groupId}/settings | Update group settings



## ApproveJoinRequests

> WhatsappGroupJoinRequestActionResponse ApproveJoinRequests(ctx, businessPhoneNumber, groupId).WhatsappGroupJoinRequestActionRequest(whatsappGroupJoinRequestActionRequest).Execute()

Approve group join requests



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    businessPhoneNumber := "+16315551111" // string | The WhatsApp business phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
    groupId := "120363345678901234@g.us" // string | WhatsApp group ID.
    whatsappGroupJoinRequestActionRequest := *ycloud.NewWhatsappGroupJoinRequestActionRequest([]string{"JoinRequests_example"}) // WhatsappGroupJoinRequestActionRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappGroupsApi.ApproveJoinRequests(context.Background(), businessPhoneNumber, groupId).WhatsappGroupJoinRequestActionRequest(whatsappGroupJoinRequestActionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappGroupsApi.ApproveJoinRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApproveJoinRequests`: WhatsappGroupJoinRequestActionResponse
    fmt.Fprintf(os.Stdout, "Response from `WhatsappGroupsApi.ApproveJoinRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessPhoneNumber** | **string** | The WhatsApp business phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | 
**groupId** | **string** | WhatsApp group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveJoinRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **whatsappGroupJoinRequestActionRequest** | [**WhatsappGroupJoinRequestActionRequest**](WhatsappGroupJoinRequestActionRequest.md) |  | 

### Return type

[**WhatsappGroupJoinRequestActionResponse**](WhatsappGroupJoinRequestActionResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> WhatsappGroupAsyncResponse Create(ctx, businessPhoneNumber).WhatsappGroupCreateRequest(whatsappGroupCreateRequest).Execute()

Create a group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    businessPhoneNumber := "+16315551111" // string | The WhatsApp business phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
    whatsappGroupCreateRequest := *ycloud.NewWhatsappGroupCreateRequest("New Purchase Inquiry") // WhatsappGroupCreateRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappGroupsApi.Create(context.Background(), businessPhoneNumber).WhatsappGroupCreateRequest(whatsappGroupCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappGroupsApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: WhatsappGroupAsyncResponse
    fmt.Fprintf(os.Stdout, "Response from `WhatsappGroupsApi.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessPhoneNumber** | **string** | The WhatsApp business phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **whatsappGroupCreateRequest** | [**WhatsappGroupCreateRequest**](WhatsappGroupCreateRequest.md) |  | 

### Return type

[**WhatsappGroupAsyncResponse**](WhatsappGroupAsyncResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> WhatsappGroupAsyncResponse Delete(ctx, businessPhoneNumber, groupId).Execute()

Delete a group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    businessPhoneNumber := "+16315551111" // string | The WhatsApp business phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
    groupId := "120363345678901234@g.us" // string | WhatsApp group ID.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappGroupsApi.Delete(context.Background(), businessPhoneNumber, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappGroupsApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: WhatsappGroupAsyncResponse
    fmt.Fprintf(os.Stdout, "Response from `WhatsappGroupsApi.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessPhoneNumber** | **string** | The WhatsApp business phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | 
**groupId** | **string** | WhatsApp group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WhatsappGroupAsyncResponse**](WhatsappGroupAsyncResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> WhatsappGroupListResponse List(ctx, businessPhoneNumber).Limit(limit).Before(before).After(after).Execute()

List groups



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    businessPhoneNumber := "+16315551111" // string | The WhatsApp business phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
    limit := int32(56) // int32 | A limit on the number of results to be returned, between 1 and 1024. Defaults to 25. (optional) (default to 25)
    before := "eyJvIjoiYmVmb3JlIn0" // string | A cursor to fetch the previous page. (optional)
    after := "eyJvIjoiYWZ0ZXIifQ" // string | A cursor to fetch the next page. (optional)

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappGroupsApi.List(context.Background(), businessPhoneNumber).Limit(limit).Before(before).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappGroupsApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: WhatsappGroupListResponse
    fmt.Fprintf(os.Stdout, "Response from `WhatsappGroupsApi.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessPhoneNumber** | **string** | The WhatsApp business phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | A limit on the number of results to be returned, between 1 and 1024. Defaults to 25. | [default to 25]
 **before** | **string** | A cursor to fetch the previous page. | 
 **after** | **string** | A cursor to fetch the next page. | 

### Return type

[**WhatsappGroupListResponse**](WhatsappGroupListResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListJoinRequests

> WhatsappGroupJoinRequestListResponse ListJoinRequests(ctx, businessPhoneNumber, groupId).Limit(limit).Before(before).After(after).Execute()

List group join requests



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    businessPhoneNumber := "+16315551111" // string | The WhatsApp business phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
    groupId := "120363345678901234@g.us" // string | WhatsApp group ID.
    limit := int32(56) // int32 | A limit on the number of results to be returned, between 1 and 1024. Defaults to 25. (optional) (default to 25)
    before := "eyJvIjoiYmVmb3JlIn0" // string | A cursor to fetch the previous page. (optional)
    after := "eyJvIjoiYWZ0ZXIifQ" // string | A cursor to fetch the next page. (optional)

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappGroupsApi.ListJoinRequests(context.Background(), businessPhoneNumber, groupId).Limit(limit).Before(before).After(after).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappGroupsApi.ListJoinRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListJoinRequests`: WhatsappGroupJoinRequestListResponse
    fmt.Fprintf(os.Stdout, "Response from `WhatsappGroupsApi.ListJoinRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessPhoneNumber** | **string** | The WhatsApp business phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | 
**groupId** | **string** | WhatsApp group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListJoinRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | A limit on the number of results to be returned, between 1 and 1024. Defaults to 25. | [default to 25]
 **before** | **string** | A cursor to fetch the previous page. | 
 **after** | **string** | A cursor to fetch the next page. | 

### Return type

[**WhatsappGroupJoinRequestListResponse**](WhatsappGroupJoinRequestListResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectJoinRequests

> WhatsappGroupJoinRequestActionResponse RejectJoinRequests(ctx, businessPhoneNumber, groupId).WhatsappGroupJoinRequestActionRequest(whatsappGroupJoinRequestActionRequest).Execute()

Reject group join requests



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    businessPhoneNumber := "+16315551111" // string | The WhatsApp business phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
    groupId := "120363345678901234@g.us" // string | WhatsApp group ID.
    whatsappGroupJoinRequestActionRequest := *ycloud.NewWhatsappGroupJoinRequestActionRequest([]string{"JoinRequests_example"}) // WhatsappGroupJoinRequestActionRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappGroupsApi.RejectJoinRequests(context.Background(), businessPhoneNumber, groupId).WhatsappGroupJoinRequestActionRequest(whatsappGroupJoinRequestActionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappGroupsApi.RejectJoinRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RejectJoinRequests`: WhatsappGroupJoinRequestActionResponse
    fmt.Fprintf(os.Stdout, "Response from `WhatsappGroupsApi.RejectJoinRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessPhoneNumber** | **string** | The WhatsApp business phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | 
**groupId** | **string** | WhatsApp group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectJoinRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **whatsappGroupJoinRequestActionRequest** | [**WhatsappGroupJoinRequestActionRequest**](WhatsappGroupJoinRequestActionRequest.md) |  | 

### Return type

[**WhatsappGroupJoinRequestActionResponse**](WhatsappGroupJoinRequestActionResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveParticipants

> WhatsappGroupAsyncResponse RemoveParticipants(ctx, groupId).WhatsappGroupRemoveParticipantsRequest(whatsappGroupRemoveParticipantsRequest).Execute()

Remove group participants



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    groupId := "120363345678901234@g.us" // string | WhatsApp group ID.
    whatsappGroupRemoveParticipantsRequest := *ycloud.NewWhatsappGroupRemoveParticipantsRequest([]ycloud.WhatsappGroupRemoveParticipant{*ycloud.NewWhatsappGroupRemoveParticipant()}) // WhatsappGroupRemoveParticipantsRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappGroupsApi.RemoveParticipants(context.Background(), groupId).WhatsappGroupRemoveParticipantsRequest(whatsappGroupRemoveParticipantsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappGroupsApi.RemoveParticipants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveParticipants`: WhatsappGroupAsyncResponse
    fmt.Fprintf(os.Stdout, "Response from `WhatsappGroupsApi.RemoveParticipants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | WhatsApp group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveParticipantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **whatsappGroupRemoveParticipantsRequest** | [**WhatsappGroupRemoveParticipantsRequest**](WhatsappGroupRemoveParticipantsRequest.md) |  | 

### Return type

[**WhatsappGroupAsyncResponse**](WhatsappGroupAsyncResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetInviteLink

> WhatsappGroupInviteLink ResetInviteLink(ctx, businessPhoneNumber, groupId).Execute()

Reset a group invite link



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    businessPhoneNumber := "+16315551111" // string | The WhatsApp business phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
    groupId := "120363345678901234@g.us" // string | WhatsApp group ID.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappGroupsApi.ResetInviteLink(context.Background(), businessPhoneNumber, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappGroupsApi.ResetInviteLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetInviteLink`: WhatsappGroupInviteLink
    fmt.Fprintf(os.Stdout, "Response from `WhatsappGroupsApi.ResetInviteLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessPhoneNumber** | **string** | The WhatsApp business phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | 
**groupId** | **string** | WhatsApp group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetInviteLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WhatsappGroupInviteLink**](WhatsappGroupInviteLink.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Retrieve

> WhatsappGroup Retrieve(ctx, businessPhoneNumber, groupId).Execute()

Retrieve a group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    businessPhoneNumber := "+16315551111" // string | The WhatsApp business phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
    groupId := "120363345678901234@g.us" // string | WhatsApp group ID.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappGroupsApi.Retrieve(context.Background(), businessPhoneNumber, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappGroupsApi.Retrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Retrieve`: WhatsappGroup
    fmt.Fprintf(os.Stdout, "Response from `WhatsappGroupsApi.Retrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessPhoneNumber** | **string** | The WhatsApp business phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | 
**groupId** | **string** | WhatsApp group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WhatsappGroup**](WhatsappGroup.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveInviteLink

> WhatsappGroupInviteLink RetrieveInviteLink(ctx, businessPhoneNumber, groupId).Execute()

Retrieve a group invite link



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    businessPhoneNumber := "+16315551111" // string | The WhatsApp business phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
    groupId := "120363345678901234@g.us" // string | WhatsApp group ID.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappGroupsApi.RetrieveInviteLink(context.Background(), businessPhoneNumber, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappGroupsApi.RetrieveInviteLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveInviteLink`: WhatsappGroupInviteLink
    fmt.Fprintf(os.Stdout, "Response from `WhatsappGroupsApi.RetrieveInviteLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessPhoneNumber** | **string** | The WhatsApp business phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | 
**groupId** | **string** | WhatsApp group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveInviteLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WhatsappGroupInviteLink**](WhatsappGroupInviteLink.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendInviteLinkMessage

> WhatsappMessage SendInviteLinkMessage(ctx, businessPhoneNumber).WhatsappGroupInviteLinkMessageRequest(whatsappGroupInviteLinkMessageRequest).Execute()

Send a group invite link message



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    businessPhoneNumber := "+16315551111" // string | The WhatsApp business phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
    whatsappGroupInviteLinkMessageRequest := *ycloud.NewWhatsappGroupInviteLinkMessageRequest("+16315551111", "group_invite_link", "en_US", []ycloud.WhatsappMessageTemplateComponentParameter{*ycloud.NewWhatsappMessageTemplateComponentParameter()}) // WhatsappGroupInviteLinkMessageRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappGroupsApi.SendInviteLinkMessage(context.Background(), businessPhoneNumber).WhatsappGroupInviteLinkMessageRequest(whatsappGroupInviteLinkMessageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappGroupsApi.SendInviteLinkMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendInviteLinkMessage`: WhatsappMessage
    fmt.Fprintf(os.Stdout, "Response from `WhatsappGroupsApi.SendInviteLinkMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessPhoneNumber** | **string** | The WhatsApp business phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendInviteLinkMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **whatsappGroupInviteLinkMessageRequest** | [**WhatsappGroupInviteLinkMessageRequest**](WhatsappGroupInviteLinkMessageRequest.md) |  | 

### Return type

[**WhatsappMessage**](WhatsappMessage.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSettings

> WhatsappGroupAsyncResponse UpdateSettings(ctx, businessPhoneNumber, groupId).WhatsappGroupUpdateSettingsRequest(whatsappGroupUpdateSettingsRequest).Execute()

Update group settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    businessPhoneNumber := "+16315551111" // string | The WhatsApp business phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
    groupId := "120363345678901234@g.us" // string | WhatsApp group ID.
    whatsappGroupUpdateSettingsRequest := *ycloud.NewWhatsappGroupUpdateSettingsRequest() // WhatsappGroupUpdateSettingsRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappGroupsApi.UpdateSettings(context.Background(), businessPhoneNumber, groupId).WhatsappGroupUpdateSettingsRequest(whatsappGroupUpdateSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappGroupsApi.UpdateSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSettings`: WhatsappGroupAsyncResponse
    fmt.Fprintf(os.Stdout, "Response from `WhatsappGroupsApi.UpdateSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**businessPhoneNumber** | **string** | The WhatsApp business phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | 
**groupId** | **string** | WhatsApp group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **whatsappGroupUpdateSettingsRequest** | [**WhatsappGroupUpdateSettingsRequest**](WhatsappGroupUpdateSettingsRequest.md) |  | 

### Return type

[**WhatsappGroupAsyncResponse**](WhatsappGroupAsyncResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

