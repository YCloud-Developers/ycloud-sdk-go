# EmailsApi

All URIs are relative to *https://api.ycloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Send**](EmailsApi.md#Send) | **Post** /emails | Send an email



## Send

> Email Send(ctx).EmailSendRequest(emailSendRequest).Execute()

Send an email



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
    emailSendRequest := *ycloud.NewEmailSendRequest("SupportTeam<support@example.com>", "to1@example.com,Nick<to2@example.com>", "Subject_example", "This is a test message from #nick#.") // EmailSendRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailsApi.Send(context.Background()).EmailSendRequest(emailSendRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailsApi.Send``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Send`: Email
    fmt.Fprintf(os.Stdout, "Response from `EmailsApi.Send`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailSendRequest** | [**EmailSendRequest**](EmailSendRequest.md) |  | 

### Return type

[**Email**](Email.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
