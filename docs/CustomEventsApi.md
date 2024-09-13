# CustomEventsApi

All URIs are relative to *https://api.ycloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDefinition**](CustomEventsApi.md#CreateDefinition) | **Post** /event/definitions | Create an event definition
[**CreatePropertyDefinition**](CustomEventsApi.md#CreatePropertyDefinition) | **Post** /event/definitions/{name}/properties | Create an event property definition
[**PropertyDefinition**](CustomEventsApi.md#PropertyDefinition) | **Delete** /event/definitions/{name}/properties/{propertyName} | Delete an event property definition
[**PropertyDefinition_0**](CustomEventsApi.md#PropertyDefinition_0) | **Patch** /event/definitions/{name}/properties/{propertyName} | Update an event property definition
[**RetrieveDefinition**](CustomEventsApi.md#RetrieveDefinition) | **Get** /event/definitions/{name} | Retrieve an event definition
[**SendEvent**](CustomEventsApi.md#SendEvent) | **Post** /event/events | Send an event
[**UpdateDefinition**](CustomEventsApi.md#UpdateDefinition) | **Patch** /event/definitions/{name} | Update an event definition



## CreateDefinition

> CustomEventDefinition CreateDefinition(ctx).CustomEventDefinitionCreateRequest(customEventDefinitionCreateRequest).Execute()

Create an event definition



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    customEventDefinitionCreateRequest := *ycloud.NewCustomEventDefinitionCreateRequest("unique_event_name", "My event label", "CONTACT") // CustomEventDefinitionCreateRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomEventsApi.CreateDefinition(context.Background()).CustomEventDefinitionCreateRequest(customEventDefinitionCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomEventsApi.CreateDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDefinition`: CustomEventDefinition
    fmt.Fprintf(os.Stdout, "Response from `CustomEventsApi.CreateDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customEventDefinitionCreateRequest** | [**CustomEventDefinitionCreateRequest**](CustomEventDefinitionCreateRequest.md) |  | 

### Return type

[**CustomEventDefinition**](CustomEventDefinition.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePropertyDefinition

> CustomEventDefinitionProperty CreatePropertyDefinition(ctx, name).CustomEventDefinitionPropertyCreateRequest(customEventDefinitionPropertyCreateRequest).Execute()

Create an event property definition



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    name := "unique_event_name" // string | Name of the custom event.
    customEventDefinitionPropertyCreateRequest := *ycloud.NewCustomEventDefinitionPropertyCreateRequest("unique_property_name", "Property Label", "STRING") // CustomEventDefinitionPropertyCreateRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomEventsApi.CreatePropertyDefinition(context.Background(), name).CustomEventDefinitionPropertyCreateRequest(customEventDefinitionPropertyCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomEventsApi.CreatePropertyDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePropertyDefinition`: CustomEventDefinitionProperty
    fmt.Fprintf(os.Stdout, "Response from `CustomEventsApi.CreatePropertyDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the custom event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePropertyDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customEventDefinitionPropertyCreateRequest** | [**CustomEventDefinitionPropertyCreateRequest**](CustomEventDefinitionPropertyCreateRequest.md) |  | 

### Return type

[**CustomEventDefinitionProperty**](CustomEventDefinitionProperty.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PropertyDefinition

> PropertyDefinition(ctx, name, propertyName).Execute()

Delete an event property definition



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    name := "unique_event_name" // string | Name of the custom event.
    propertyName := "unique_property_name" // string | Name of the custom event property.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomEventsApi.PropertyDefinition(context.Background(), name, propertyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomEventsApi.PropertyDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the custom event. | 
**propertyName** | **string** | Name of the custom event property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPropertyDefinitionRequest struct via the builder pattern


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


## PropertyDefinition_0

> CustomEventDefinitionProperty PropertyDefinition_0(ctx, name, propertyName).CustomEventDefinitionPropertyUpdateRequest(customEventDefinitionPropertyUpdateRequest).Execute()

Update an event property definition



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    name := "unique_event_name" // string | Name of the custom event.
    propertyName := "unique_property_name" // string | Name of the custom event property.
    customEventDefinitionPropertyUpdateRequest := *ycloud.NewCustomEventDefinitionPropertyUpdateRequest() // CustomEventDefinitionPropertyUpdateRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomEventsApi.PropertyDefinition_0(context.Background(), name, propertyName).CustomEventDefinitionPropertyUpdateRequest(customEventDefinitionPropertyUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomEventsApi.PropertyDefinition_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PropertyDefinition_0`: CustomEventDefinitionProperty
    fmt.Fprintf(os.Stdout, "Response from `CustomEventsApi.PropertyDefinition_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the custom event. | 
**propertyName** | **string** | Name of the custom event property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPropertyDefinition_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **customEventDefinitionPropertyUpdateRequest** | [**CustomEventDefinitionPropertyUpdateRequest**](CustomEventDefinitionPropertyUpdateRequest.md) |  | 

### Return type

[**CustomEventDefinitionProperty**](CustomEventDefinitionProperty.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDefinition

> CustomEventDefinition RetrieveDefinition(ctx, name).Execute()

Retrieve an event definition



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    name := "unique_event_name" // string | Name of the custom event.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomEventsApi.RetrieveDefinition(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomEventsApi.RetrieveDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveDefinition`: CustomEventDefinition
    fmt.Fprintf(os.Stdout, "Response from `CustomEventsApi.RetrieveDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the custom event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomEventDefinition**](CustomEventDefinition.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendEvent

> SendEvent(ctx).CustomEventSendRequest(customEventSendRequest).Execute()

Send an event



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    customEventSendRequest := *ycloud.NewCustomEventSendRequest("unique_event_name") // CustomEventSendRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomEventsApi.SendEvent(context.Background()).CustomEventSendRequest(customEventSendRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomEventsApi.SendEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customEventSendRequest** | [**CustomEventSendRequest**](CustomEventSendRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDefinition

> CustomEventDefinition UpdateDefinition(ctx, name).CustomEventDefinitionUpdateRequest(customEventDefinitionUpdateRequest).Execute()

Update an event definition



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    name := "unique_event_name" // string | Name of the custom event.
    customEventDefinitionUpdateRequest := *ycloud.NewCustomEventDefinitionUpdateRequest() // CustomEventDefinitionUpdateRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomEventsApi.UpdateDefinition(context.Background(), name).CustomEventDefinitionUpdateRequest(customEventDefinitionUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomEventsApi.UpdateDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDefinition`: CustomEventDefinition
    fmt.Fprintf(os.Stdout, "Response from `CustomEventsApi.UpdateDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the custom event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customEventDefinitionUpdateRequest** | [**CustomEventDefinitionUpdateRequest**](CustomEventDefinitionUpdateRequest.md) |  | 

### Return type

[**CustomEventDefinition**](CustomEventDefinition.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

