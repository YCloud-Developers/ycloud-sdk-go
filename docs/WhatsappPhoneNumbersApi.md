# WhatsappPhoneNumbersApi

All URIs are relative to *https://api.ycloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](WhatsappPhoneNumbersApi.md#List) | **Get** /whatsapp/phoneNumbers | List WhatsApp phone numbers
[**Retrieve**](WhatsappPhoneNumbersApi.md#Retrieve) | **Get** /whatsapp/phoneNumbers/{wabaId}/{phoneNumber} | Retrieve a WhatsApp phone number



## List

> WhatsappPhoneNumberPage List(ctx).Page(page).Limit(limit).IncludeTotal(includeTotal).FilterWabaId(filterWabaId).Execute()

List WhatsApp phone numbers



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
    page := int32(56) // int32 | Page number of the results to be returned, 1-based. (optional) (default to 1)
    limit := int32(56) // int32 | A limit on the number of results to be returned, or number of results per page, between 1 and 100, defaults to 10. (optional) (default to 10)
    includeTotal := true // bool | Return results inside an object that contains the total result count or not. (optional) (default to false)
    filterWabaId := "whatsapp-business-account-id" // string | WhatsApp Business Account ID. (optional)

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappPhoneNumbersApi.List(context.Background()).Page(page).Limit(limit).IncludeTotal(includeTotal).FilterWabaId(filterWabaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappPhoneNumbersApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: WhatsappPhoneNumberPage
    fmt.Fprintf(os.Stdout, "Response from `WhatsappPhoneNumbersApi.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number of the results to be returned, 1-based. | [default to 1]
 **limit** | **int32** | A limit on the number of results to be returned, or number of results per page, between 1 and 100, defaults to 10. | [default to 10]
 **includeTotal** | **bool** | Return results inside an object that contains the total result count or not. | [default to false]
 **filterWabaId** | **string** | WhatsApp Business Account ID. | 

### Return type

[**WhatsappPhoneNumberPage**](WhatsappPhoneNumberPage.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Retrieve

> WhatsappPhoneNumber Retrieve(ctx, wabaId, phoneNumber).Execute()

Retrieve a WhatsApp phone number



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
    wabaId := "whatsapp-business-account-id" // string | WhatsApp Business Account ID.
    phoneNumber := "+447901614024" // string | Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappPhoneNumbersApi.Retrieve(context.Background(), wabaId, phoneNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappPhoneNumbersApi.Retrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Retrieve`: WhatsappPhoneNumber
    fmt.Fprintf(os.Stdout, "Response from `WhatsappPhoneNumbersApi.Retrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wabaId** | **string** | WhatsApp Business Account ID. | 
**phoneNumber** | **string** | Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WhatsappPhoneNumber**](WhatsappPhoneNumber.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

