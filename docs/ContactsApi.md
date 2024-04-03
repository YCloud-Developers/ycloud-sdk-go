# ContactsApi

All URIs are relative to *https://api.ycloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](ContactsApi.md#Create) | **Post** /contact/contacts | Create a contact
[**Delete**](ContactsApi.md#Delete) | **Delete** /contact/contacts/{id} | Delete a contact
[**List**](ContactsApi.md#List) | **Get** /contact/contacts | List contacts
[**Retrieve**](ContactsApi.md#Retrieve) | **Get** /contact/contacts/{id} | Retrieve a contact
[**Update**](ContactsApi.md#Update) | **Patch** /contact/contacts/{id} | Update a contact



## Create

> Contact Create(ctx).ContactCreateRequest(contactCreateRequest).Execute()

Create a contact



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
    contactCreateRequest := *ycloud.NewContactCreateRequest("+16315551111") // ContactCreateRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactsApi.Create(context.Background()).ContactCreateRequest(contactCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Contact
    fmt.Fprintf(os.Stdout, "Response from `ContactsApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactCreateRequest** | [**ContactCreateRequest**](ContactCreateRequest.md) |  | 

### Return type

[**Contact**](Contact.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Contact Delete(ctx, id).Execute()

Delete a contact



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
    id := "1693364594105000026" // string | ID of the contact.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactsApi.Delete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: Contact
    fmt.Fprintf(os.Stdout, "Response from `ContactsApi.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the contact. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Contact**](Contact.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ContactPage List(ctx).Page(page).Limit(limit).IncludeTotal(includeTotal).FilterTags(filterTags).FilterCountryCode(filterCountryCode).FilterPhoneNumber(filterPhoneNumber).FilterEmail(filterEmail).Execute()

List contacts



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
    filterTags := "tag1,tag2" // string | Comma-separated list of tags. (optional)
    filterCountryCode := "US" // string | Two-letter country abbreviation. See [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). (optional)
    filterPhoneNumber := "+16315551111" // string | Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. (optional)
    filterEmail := "support@example.com" // string | The contact's email address. (optional)

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactsApi.List(context.Background()).Page(page).Limit(limit).IncludeTotal(includeTotal).FilterTags(filterTags).FilterCountryCode(filterCountryCode).FilterPhoneNumber(filterPhoneNumber).FilterEmail(filterEmail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ContactPage
    fmt.Fprintf(os.Stdout, "Response from `ContactsApi.List`: %v\n", resp)
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
 **filterTags** | **string** | Comma-separated list of tags. | 
 **filterCountryCode** | **string** | Two-letter country abbreviation. See [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). | 
 **filterPhoneNumber** | **string** | Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | 
 **filterEmail** | **string** | The contact&#39;s email address. | 

### Return type

[**ContactPage**](ContactPage.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Retrieve

> Contact Retrieve(ctx, id).Execute()

Retrieve a contact



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
    id := "1693364594105000026" // string | ID of the contact.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactsApi.Retrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsApi.Retrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Retrieve`: Contact
    fmt.Fprintf(os.Stdout, "Response from `ContactsApi.Retrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the contact. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Contact**](Contact.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Contact Update(ctx, id).ContactUpdateRequest(contactUpdateRequest).Execute()

Update a contact



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
    id := "1693364594105000026" // string | ID of the contact.
    contactUpdateRequest := *ycloud.NewContactUpdateRequest() // ContactUpdateRequest |  (optional)

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactsApi.Update(context.Background(), id).ContactUpdateRequest(contactUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Contact
    fmt.Fprintf(os.Stdout, "Response from `ContactsApi.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the contact. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contactUpdateRequest** | [**ContactUpdateRequest**](ContactUpdateRequest.md) |  | 

### Return type

[**Contact**](Contact.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

