# \WebhookEndpointsApi

All URIs are relative to *https://api.ycloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](WebhookEndpointsApi.md#Create) | **Post** /webhookEndpoints | Create a webhook endpoint
[**Delete**](WebhookEndpointsApi.md#Delete) | **Delete** /webhookEndpoints/{id} | Delete a webhook endpoint
[**List**](WebhookEndpointsApi.md#List) | **Get** /webhookEndpoints | List webhook endpoints
[**Retrieve**](WebhookEndpointsApi.md#Retrieve) | **Get** /webhookEndpoints/{id} | Retrieve a webhook endpoint
[**RotateSecret**](WebhookEndpointsApi.md#RotateSecret) | **Post** /webhookEndpoints/{id}/rotateSecret | Rotate a webhook endpoint secret
[**Update**](WebhookEndpointsApi.md#Update) | **Patch** /webhookEndpoints/{id} | Update a webhook endpoint



## Create

> WebhookEndpoint Create(ctx).WebhookEndpointCreateRequest(webhookEndpointCreateRequest).Execute()

Create a webhook endpoint



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
    webhookEndpointCreateRequest := *ycloud.NewWebhookEndpointCreateRequest("https://httpbin.org/anything?tag=api", []ycloud.EventType{ycloud.EventType("email.delivery.updated")}) // WebhookEndpointCreateRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookEndpointsApi.Create(context.Background()).WebhookEndpointCreateRequest(webhookEndpointCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookEndpointsApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: WebhookEndpoint
    fmt.Fprintf(os.Stdout, "Response from `WebhookEndpointsApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEndpointCreateRequest** | [**WebhookEndpointCreateRequest**](WebhookEndpointCreateRequest.md) |  | 

### Return type

[**WebhookEndpoint**](WebhookEndpoint.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> WebhookEndpoint Delete(ctx, id).Execute()

Delete a webhook endpoint



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
    id := "wh627c8640675de8fc689ab9d9" // string | ID of the webhook endpoint.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookEndpointsApi.Delete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookEndpointsApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: WebhookEndpoint
    fmt.Fprintf(os.Stdout, "Response from `WebhookEndpointsApi.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the webhook endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookEndpoint**](WebhookEndpoint.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> WebhookEndpointPage List(ctx).Page(page).Limit(limit).IncludeTotal(includeTotal).Execute()

List webhook endpoints



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

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookEndpointsApi.List(context.Background()).Page(page).Limit(limit).IncludeTotal(includeTotal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookEndpointsApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: WebhookEndpointPage
    fmt.Fprintf(os.Stdout, "Response from `WebhookEndpointsApi.List`: %v\n", resp)
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

### Return type

[**WebhookEndpointPage**](WebhookEndpointPage.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Retrieve

> WebhookEndpoint Retrieve(ctx, id).Execute()

Retrieve a webhook endpoint



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
    id := "wh627c8640675de8fc689ab9d9" // string | ID of the webhook endpoint.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookEndpointsApi.Retrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookEndpointsApi.Retrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Retrieve`: WebhookEndpoint
    fmt.Fprintf(os.Stdout, "Response from `WebhookEndpointsApi.Retrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the webhook endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookEndpoint**](WebhookEndpoint.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateSecret

> WebhookEndpoint RotateSecret(ctx, id).Execute()

Rotate a webhook endpoint secret



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
    id := "wh627c8640675de8fc689ab9d9" // string | ID of the webhook endpoint.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookEndpointsApi.RotateSecret(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookEndpointsApi.RotateSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RotateSecret`: WebhookEndpoint
    fmt.Fprintf(os.Stdout, "Response from `WebhookEndpointsApi.RotateSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the webhook endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookEndpoint**](WebhookEndpoint.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> WebhookEndpoint Update(ctx, id).WebhookEndpointUpdateRequest(webhookEndpointUpdateRequest).Execute()

Update a webhook endpoint



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
    id := "wh627c8640675de8fc689ab9d9" // string | ID of the webhook endpoint.
    webhookEndpointUpdateRequest := *ycloud.NewWebhookEndpointUpdateRequest() // WebhookEndpointUpdateRequest |  (optional)

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookEndpointsApi.Update(context.Background(), id).WebhookEndpointUpdateRequest(webhookEndpointUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookEndpointsApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: WebhookEndpoint
    fmt.Fprintf(os.Stdout, "Response from `WebhookEndpointsApi.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the webhook endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webhookEndpointUpdateRequest** | [**WebhookEndpointUpdateRequest**](WebhookEndpointUpdateRequest.md) |  | 

### Return type

[**WebhookEndpoint**](WebhookEndpoint.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

