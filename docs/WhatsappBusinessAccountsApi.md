# WhatsappBusinessAccountsApi

All URIs are relative to *https://api.ycloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](WhatsappBusinessAccountsApi.md#List) | **Get** /whatsapp/businessAccounts | List WABAs
[**Retrieve**](WhatsappBusinessAccountsApi.md#Retrieve) | **Get** /whatsapp/businessAccounts/{id} | Retrieve a WABA



## List

> WhatsappBusinessAccount List(ctx).FilterAccountReviewStatus(filterAccountReviewStatus).Execute()

List WABAs



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
    filterAccountReviewStatus := "APPROVED" // string | WhatsApp Business Account review status. (optional)

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappBusinessAccountsApi.List(context.Background()).FilterAccountReviewStatus(filterAccountReviewStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappBusinessAccountsApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: WhatsappBusinessAccount
    fmt.Fprintf(os.Stdout, "Response from `WhatsappBusinessAccountsApi.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterAccountReviewStatus** | **string** | WhatsApp Business Account review status. | 

### Return type

[**WhatsappBusinessAccount**](WhatsappBusinessAccount.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Retrieve

> WhatsappBusinessAccount Retrieve(ctx, id).Execute()

Retrieve a WABA



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
    id := "whatsapp-business-account-id" // string | WhatsApp Business Account ID.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappBusinessAccountsApi.Retrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappBusinessAccountsApi.Retrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Retrieve`: WhatsappBusinessAccount
    fmt.Fprintf(os.Stdout, "Response from `WhatsappBusinessAccountsApi.Retrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | WhatsApp Business Account ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WhatsappBusinessAccount**](WhatsappBusinessAccount.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

