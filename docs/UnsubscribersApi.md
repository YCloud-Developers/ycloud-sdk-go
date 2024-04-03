# UnsubscribersApi

All URIs are relative to *https://api.ycloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](UnsubscribersApi.md#Create) | **Post** /unsubscribers | Create an unsubscriber
[**DeleteByCustomerAndChannel**](UnsubscribersApi.md#DeleteByCustomerAndChannel) | **Delete** /unsubscribers/{customer}/{channel} | Delete an unsubscriber
[**List**](UnsubscribersApi.md#List) | **Get** /unsubscribers | List unsubscribers
[**ListAllByCustomer**](UnsubscribersApi.md#ListAllByCustomer) | **Get** /unsubscribers/{customer} | List all unsubscribers by customer
[**RetrieveByCustomerAndChannel**](UnsubscribersApi.md#RetrieveByCustomerAndChannel) | **Get** /unsubscribers/{customer}/{channel} | Retrieve an unsubscriber



## Create

> Unsubscriber Create(ctx).UnsubscriberCreateRequest(unsubscriberCreateRequest).Execute()

Create an unsubscriber



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
    unsubscriberCreateRequest := *ycloud.NewUnsubscriberCreateRequest(ycloud.UnsubscriberType("PHONE_NUMBER"), "+16315551111", ycloud.UnsubscriberChannel("whatsapp")) // UnsubscriberCreateRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.UnsubscribersApi.Create(context.Background()).UnsubscriberCreateRequest(unsubscriberCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnsubscribersApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Unsubscriber
    fmt.Fprintf(os.Stdout, "Response from `UnsubscribersApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unsubscriberCreateRequest** | [**UnsubscriberCreateRequest**](UnsubscriberCreateRequest.md) |  | 

### Return type

[**Unsubscriber**](Unsubscriber.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteByCustomerAndChannel

> Unsubscriber DeleteByCustomerAndChannel(ctx, customer, channel).Execute()

Delete an unsubscriber



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
    customer := "+16315551111" // string | The customer who has opted out.
    channel := ycloud.UnsubscriberChannel("whatsapp") // UnsubscriberChannel | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.UnsubscribersApi.DeleteByCustomerAndChannel(context.Background(), customer, channel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnsubscribersApi.DeleteByCustomerAndChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteByCustomerAndChannel`: Unsubscriber
    fmt.Fprintf(os.Stdout, "Response from `UnsubscribersApi.DeleteByCustomerAndChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** | The customer who has opted out. | 
**channel** | [**UnsubscriberChannel**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteByCustomerAndChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Unsubscriber**](Unsubscriber.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> UnsubscriberPage List(ctx).Page(page).Limit(limit).IncludeTotal(includeTotal).PageAfter(pageAfter).FilterCustomer(filterCustomer).FilterChannel(filterChannel).FilterRegionCode(filterRegionCode).Execute()

List unsubscribers



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
    pageAfter := "id:foo" // string | A cursor to fetch the next page in cursor pagination. For example, if you make a list request, receive 100 objects and `cursor.after=id:foo`, your subsequent call can include `pageAfter=id:foo` in order to fetch the next page of the list. (optional)
    filterCustomer := "+16315551111" // string |  (optional)
    filterChannel := ycloud.UnsubscriberChannel("whatsapp") // UnsubscriberChannel |  (optional)
    filterRegionCode := "filterRegionCode_example" // string |  (optional)

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.UnsubscribersApi.List(context.Background()).Page(page).Limit(limit).IncludeTotal(includeTotal).PageAfter(pageAfter).FilterCustomer(filterCustomer).FilterChannel(filterChannel).FilterRegionCode(filterRegionCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnsubscribersApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: UnsubscriberPage
    fmt.Fprintf(os.Stdout, "Response from `UnsubscribersApi.List`: %v\n", resp)
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
 **pageAfter** | **string** | A cursor to fetch the next page in cursor pagination. For example, if you make a list request, receive 100 objects and &#x60;cursor.after&#x3D;id:foo&#x60;, your subsequent call can include &#x60;pageAfter&#x3D;id:foo&#x60; in order to fetch the next page of the list. | 
 **filterCustomer** | **string** |  | 
 **filterChannel** | [**UnsubscriberChannel**](UnsubscriberChannel.md) |  | 
 **filterRegionCode** | **string** |  | 

### Return type

[**UnsubscriberPage**](UnsubscriberPage.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllByCustomer

> []Unsubscriber ListAllByCustomer(ctx, customer).Execute()

List all unsubscribers by customer



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
    customer := "+16315551111" // string | The customer who has opted out.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.UnsubscribersApi.ListAllByCustomer(context.Background(), customer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnsubscribersApi.ListAllByCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllByCustomer`: []Unsubscriber
    fmt.Fprintf(os.Stdout, "Response from `UnsubscribersApi.ListAllByCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** | The customer who has opted out. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllByCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Unsubscriber**](Unsubscriber.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveByCustomerAndChannel

> Unsubscriber RetrieveByCustomerAndChannel(ctx, customer, channel).Execute()

Retrieve an unsubscriber



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
    customer := "+16315551111" // string | The customer who has opted out.
    channel := ycloud.UnsubscriberChannel("whatsapp") // UnsubscriberChannel | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.UnsubscribersApi.RetrieveByCustomerAndChannel(context.Background(), customer, channel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnsubscribersApi.RetrieveByCustomerAndChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveByCustomerAndChannel`: Unsubscriber
    fmt.Fprintf(os.Stdout, "Response from `UnsubscribersApi.RetrieveByCustomerAndChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | **string** | The customer who has opted out. | 
**channel** | [**UnsubscriberChannel**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveByCustomerAndChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Unsubscriber**](Unsubscriber.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

