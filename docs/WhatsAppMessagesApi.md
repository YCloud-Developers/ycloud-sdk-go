# WhatsappMessagesApi

All URIs are relative to *https://api.ycloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Send**](WhatsappMessagesApi.md#Send) | **Post** /whatsapp/messages | Send a WhatsApp message



## Send

> WhatsappMessage Send(ctx).WhatsappMessageSendRequest(whatsappMessageSendRequest).Execute()

Send a WhatsApp message



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
    whatsappMessageSendRequest := *ycloud.NewWhatsappMessageSendRequest("+447901614024", "+447901614024", ycloud.WhatsappMessageType("template")) // WhatsappMessageSendRequest |  (optional)

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

