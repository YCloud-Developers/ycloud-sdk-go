# WhatsappMessagesApi

All URIs are relative to *https://api.ycloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Retrieve**](WhatsappMessagesApi.md#Retrieve) | **Get** /whatsapp/messages/{id} | Retrieve a message
[**Send**](WhatsappMessagesApi.md#Send) | **Post** /whatsapp/messages | Enqueue a message
[**SendDirectly**](WhatsappMessagesApi.md#SendDirectly) | **Post** /whatsapp/messages/sendDirectly | Send a message directly



## Retrieve

> WhatsappMessage Retrieve(ctx, id).Execute()

Retrieve a message



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-cpaas/ycloud-sdk-go"
)

func main() {
    id := "627c8640675de8fc689ab9d9" // string | ID of the object.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappMessagesApi.Retrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappMessagesApi.Retrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Retrieve`: WhatsappMessage
    fmt.Fprintf(os.Stdout, "Response from `WhatsappMessagesApi.Retrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WhatsappMessage**](WhatsappMessage.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Send

> WhatsappMessage Send(ctx).WhatsappMessageSendRequest(whatsappMessageSendRequest).Execute()

Enqueue a message



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-cpaas/ycloud-sdk-go"
)

func main() {
    whatsappMessageSendRequest := *ycloud.NewWhatsappMessageSendRequest("+447901614024", "+447901614024", ycloud.WhatsappMessageType("template")) // WhatsappMessageSendRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappMessagesApi.Send(context.Background()).WhatsappMessageSendRequest(whatsappMessageSendRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappMessagesApi.Send``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Send`: WhatsappMessage
    fmt.Fprintf(os.Stdout, "Response from `WhatsappMessagesApi.Send`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **whatsappMessageSendRequest** | [**WhatsappMessageSendRequest**](WhatsappMessageSendRequest.md) |  | 

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


## SendDirectly

> WhatsappMessage SendDirectly(ctx).WhatsappMessageSendRequest(whatsappMessageSendRequest).Execute()

Send a message directly



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-cpaas/ycloud-sdk-go"
)

func main() {
    whatsappMessageSendRequest := *ycloud.NewWhatsappMessageSendRequest("+447901614024", "+447901614024", ycloud.WhatsappMessageType("template")) // WhatsappMessageSendRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappMessagesApi.SendDirectly(context.Background()).WhatsappMessageSendRequest(whatsappMessageSendRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappMessagesApi.SendDirectly``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendDirectly`: WhatsappMessage
    fmt.Fprintf(os.Stdout, "Response from `WhatsappMessagesApi.SendDirectly`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendDirectlyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **whatsappMessageSendRequest** | [**WhatsappMessageSendRequest**](WhatsappMessageSendRequest.md) |  | 

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
