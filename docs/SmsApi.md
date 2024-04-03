# SmsApi

All URIs are relative to *https://api.ycloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](SmsApi.md#List) | **Get** /sms | List SMS records
[**Send**](SmsApi.md#Send) | **Post** /sms | Send an SMS



## List

> SmsPage List(ctx).Page(page).Limit(limit).IncludeTotal(includeTotal).FilterCreateTimeGte(filterCreateTimeGte).FilterCreateTimeLte(filterCreateTimeLte).FilterId(filterId).Execute()

List SMS records



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    ycloud "github.com/ycloud-cpaas/ycloud-sdk-go"
)

func main() {
    page := int32(56) // int32 | Page number of the results to be returned, 1-based. (optional) (default to 1)
    limit := int32(56) // int32 | A limit on the number of results to be returned, or number of results per page, between 1 and 100, defaults to 10. (optional) (default to 10)
    includeTotal := true // bool | Return results inside an object that contains the total result count or not. (optional) (default to false)
    filterCreateTimeGte := time.Now() // time.Time | Return results where the `createTime` field is greater than or equal to this value. Default: One day ago from now. (optional)
    filterCreateTimeLte := time.Now() // time.Time | Return results where the `createTime` field is less than or equal to this value. (optional)
    filterId := "filterId_example" // string | Unique object ID on our side. Other filter parameters will be ignored if this parameter is present. (optional)

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.SmsApi.List(context.Background()).Page(page).Limit(limit).IncludeTotal(includeTotal).FilterCreateTimeGte(filterCreateTimeGte).FilterCreateTimeLte(filterCreateTimeLte).FilterId(filterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmsApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: SmsPage
    fmt.Fprintf(os.Stdout, "Response from `SmsApi.List`: %v\n", resp)
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
 **filterCreateTimeGte** | **time.Time** | Return results where the &#x60;createTime&#x60; field is greater than or equal to this value. Default: One day ago from now. | 
 **filterCreateTimeLte** | **time.Time** | Return results where the &#x60;createTime&#x60; field is less than or equal to this value. | 
 **filterId** | **string** | Unique object ID on our side. Other filter parameters will be ignored if this parameter is present. | 

### Return type

[**SmsPage**](SmsPage.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Send

> Sms Send(ctx).SmsSendRequest(smsSendRequest).Execute()

Send an SMS



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
    smsSendRequest := *ycloud.NewSmsSendRequest("+16315551111", "Your verification code is 123456.") // SmsSendRequest | SMS request that needs to be sent.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.SmsApi.Send(context.Background()).SmsSendRequest(smsSendRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmsApi.Send``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Send`: Sms
    fmt.Fprintf(os.Stdout, "Response from `SmsApi.Send`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smsSendRequest** | [**SmsSendRequest**](SmsSendRequest.md) | SMS request that needs to be sent. | 

### Return type

[**Sms**](Sms.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

