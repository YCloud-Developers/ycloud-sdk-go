# WhatsappInboundMessagesApi

All URIs are relative to *https://api.ycloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarkAsRead**](WhatsappInboundMessagesApi.md#MarkAsRead) | **Post** /whatsapp/inboundMessages/{id}/markAsRead | Mark message as read



## MarkAsRead

> MarkAsRead(ctx, id).Execute()

Mark message as read



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
    id := "627c8640675de8fc689ab9d9" // string | ID of the message.  A wamid (i.e., the original message ID on WhatsApp's platform) is also acceptable.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappInboundMessagesApi.MarkAsRead(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappInboundMessagesApi.MarkAsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the message.  A wamid (i.e., the original message ID on WhatsApp&#39;s platform) is also acceptable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkAsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
