# WhatsappFlowsApi

All URIs are relative to *https://api.ycloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](WhatsappFlowsApi.md#Create) | **Post** /whatsapp/flows | Create a flow
[**Delete**](WhatsappFlowsApi.md#Delete) | **Delete** /whatsapp/flows/{flowId} | Delete a flow
[**Deprecate**](WhatsappFlowsApi.md#Deprecate) | **Post** /whatsapp/flows/{flowId}/deprecate | Deprecate a flow
[**List**](WhatsappFlowsApi.md#List) | **Get** /whatsapp/flows | List flows
[**Preview**](WhatsappFlowsApi.md#Preview) | **Get** /whatsapp/flows/{flowId}/preview | generate a web preview URL with this flow.
[**Publish**](WhatsappFlowsApi.md#Publish) | **Post** /whatsapp/flows/{flowId}/publish | Publish a flow
[**Retrieve**](WhatsappFlowsApi.md#Retrieve) | **Get** /whatsapp/flows/{flowId} | Retrieve a flow
[**UpdateMetadata**](WhatsappFlowsApi.md#UpdateMetadata) | **Patch** /whatsapp/flows/{flowId}/metadata | Update flow metadata
[**UpdateStructure**](WhatsappFlowsApi.md#UpdateStructure) | **Patch** /whatsapp/flows/{flowId}/assets | Update flow structure



## Create

> WhatsappFlowCreate200Response Create(ctx).WhatsappFlowCreateRequest(whatsappFlowCreateRequest).Execute()

Create a flow



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
    whatsappFlowCreateRequest := *ycloud.NewWhatsappFlowCreateRequest("whatsapp-business-account-id", "My first flow", []ycloud.WhatsappFlowCategory{ycloud.WhatsappFlowCategory("SIGN_UP")}) // WhatsappFlowCreateRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappFlowsApi.Create(context.Background()).WhatsappFlowCreateRequest(whatsappFlowCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappFlowsApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: WhatsappFlowCreate200Response
    fmt.Fprintf(os.Stdout, "Response from `WhatsappFlowsApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **whatsappFlowCreateRequest** | [**WhatsappFlowCreateRequest**](WhatsappFlowCreateRequest.md) |  | 

### Return type

[**WhatsappFlowCreate200Response**](WhatsappFlowCreate200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> WhatsappFlowUpdateMetadata200Response Delete(ctx, flowId).Execute()

Delete a flow



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
    flowId := "flow-1" // string | Flow ID.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappFlowsApi.Delete(context.Background(), flowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappFlowsApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: WhatsappFlowUpdateMetadata200Response
    fmt.Fprintf(os.Stdout, "Response from `WhatsappFlowsApi.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | Flow ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WhatsappFlowUpdateMetadata200Response**](WhatsappFlowUpdateMetadata200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Deprecate

> WhatsappFlowUpdateMetadata200Response Deprecate(ctx, flowId).Execute()

Deprecate a flow



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
    flowId := "flow-1" // string | Flow ID.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappFlowsApi.Deprecate(context.Background(), flowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappFlowsApi.Deprecate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Deprecate`: WhatsappFlowUpdateMetadata200Response
    fmt.Fprintf(os.Stdout, "Response from `WhatsappFlowsApi.Deprecate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | Flow ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeprecateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WhatsappFlowUpdateMetadata200Response**](WhatsappFlowUpdateMetadata200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> WhatsappFlowList200Response List(ctx).WabaId(wabaId).Execute()

List flows



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
    wabaId := "whatsapp-business-account-id" // string | WhatsApp Business Account ID.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappFlowsApi.List(context.Background()).WabaId(wabaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappFlowsApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: WhatsappFlowList200Response
    fmt.Fprintf(os.Stdout, "Response from `WhatsappFlowsApi.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wabaId** | **string** | WhatsApp Business Account ID. | 

### Return type

[**WhatsappFlowList200Response**](WhatsappFlowList200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Preview

> WhatsappFlowPreviewUrl Preview(ctx, flowId).Invalidate(invalidate).Execute()

generate a web preview URL with this flow.



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
    flowId := "flow-1" // string | Flow ID.
    invalidate := false // bool | the link will expire in 30 days in default, or if you set with invalidate=true which will generate a new link. (optional)

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappFlowsApi.Preview(context.Background(), flowId).Invalidate(invalidate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappFlowsApi.Preview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Preview`: WhatsappFlowPreviewUrl
    fmt.Fprintf(os.Stdout, "Response from `WhatsappFlowsApi.Preview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | Flow ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invalidate** | **bool** | the link will expire in 30 days in default, or if you set with invalidate&#x3D;true which will generate a new link. | 

### Return type

[**WhatsappFlowPreviewUrl**](WhatsappFlowPreviewUrl.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Publish

> WhatsappFlowUpdateMetadata200Response Publish(ctx, flowId).Execute()

Publish a flow



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
    flowId := "flow-1" // string | Flow ID.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappFlowsApi.Publish(context.Background(), flowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappFlowsApi.Publish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Publish`: WhatsappFlowUpdateMetadata200Response
    fmt.Fprintf(os.Stdout, "Response from `WhatsappFlowsApi.Publish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | Flow ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WhatsappFlowUpdateMetadata200Response**](WhatsappFlowUpdateMetadata200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Retrieve

> WhatsappFlow Retrieve(ctx, flowId).Execute()

Retrieve a flow



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
    flowId := "flow-1" // string | Flow ID.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappFlowsApi.Retrieve(context.Background(), flowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappFlowsApi.Retrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Retrieve`: WhatsappFlow
    fmt.Fprintf(os.Stdout, "Response from `WhatsappFlowsApi.Retrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | Flow ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WhatsappFlow**](WhatsappFlow.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMetadata

> WhatsappFlowUpdateMetadata200Response UpdateMetadata(ctx, flowId).WhatsappFlowUpdateMetadataRequest(whatsappFlowUpdateMetadataRequest).Execute()

Update flow metadata



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
    flowId := "flow-1" // string | Flow ID.
    whatsappFlowUpdateMetadataRequest := *ycloud.NewWhatsappFlowUpdateMetadataRequest() // WhatsappFlowUpdateMetadataRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappFlowsApi.UpdateMetadata(context.Background(), flowId).WhatsappFlowUpdateMetadataRequest(whatsappFlowUpdateMetadataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappFlowsApi.UpdateMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMetadata`: WhatsappFlowUpdateMetadata200Response
    fmt.Fprintf(os.Stdout, "Response from `WhatsappFlowsApi.UpdateMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | Flow ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **whatsappFlowUpdateMetadataRequest** | [**WhatsappFlowUpdateMetadataRequest**](WhatsappFlowUpdateMetadataRequest.md) |  | 

### Return type

[**WhatsappFlowUpdateMetadata200Response**](WhatsappFlowUpdateMetadata200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStructure

> WhatsappFlowUpdateMetadata200Response UpdateStructure(ctx, flowId).FlowJson(flowJson).Execute()

Update flow structure



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
    flowId := "flow-1" // string | Flow ID.
    flowJson := os.NewFile(1234, "some_file") // *os.File | JSON file containing the Flow structure.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappFlowsApi.UpdateStructure(context.Background(), flowId).FlowJson(flowJson).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappFlowsApi.UpdateStructure``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStructure`: WhatsappFlowUpdateMetadata200Response
    fmt.Fprintf(os.Stdout, "Response from `WhatsappFlowsApi.UpdateStructure`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** | Flow ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStructureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **flowJson** | ***os.File** | JSON file containing the Flow structure. | 

### Return type

[**WhatsappFlowUpdateMetadata200Response**](WhatsappFlowUpdateMetadata200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

