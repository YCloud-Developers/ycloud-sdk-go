# WhatsappTemplatesApi

All URIs are relative to *https://api.ycloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](WhatsappTemplatesApi.md#Create) | **Post** /whatsapp/templates | Create a WhatsApp template
[**DeleteByName**](WhatsappTemplatesApi.md#DeleteByName) | **Delete** /whatsapp/templates/{wabaId}/{name} | Delete WhatsApp templates by name
[**EditByNameAndLanguage**](WhatsappTemplatesApi.md#EditByNameAndLanguage) | **Patch** /whatsapp/templates/{wabaId}/{name}/{language} | Edit a WhatsApp template
[**List**](WhatsappTemplatesApi.md#List) | **Get** /whatsapp/templates | List WhatsApp templates
[**RetrieveByNameAndLanguage**](WhatsappTemplatesApi.md#RetrieveByNameAndLanguage) | **Get** /whatsapp/templates/{wabaId}/{name}/{language} | Retrieve a WhatsApp template



## Create

> WhatsappTemplate Create(ctx).WhatsappTemplateCreateRequest(whatsappTemplateCreateRequest).Execute()

Create a WhatsApp template



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
    whatsappTemplateCreateRequest := *ycloud.NewWhatsappTemplateCreateRequest("whatsapp-business-account-id", "sample_whatsapp_template", "en", ycloud.WhatsappTemplateCategory("AUTHENTICATION"), []ycloud.WhatsappTemplateComponent{*ycloud.NewWhatsappTemplateComponent()}) // WhatsappTemplateCreateRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappTemplatesApi.Create(context.Background()).WhatsappTemplateCreateRequest(whatsappTemplateCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappTemplatesApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: WhatsappTemplate
    fmt.Fprintf(os.Stdout, "Response from `WhatsappTemplatesApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **whatsappTemplateCreateRequest** | [**WhatsappTemplateCreateRequest**](WhatsappTemplateCreateRequest.md) |  | 

### Return type

[**WhatsappTemplate**](WhatsappTemplate.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteByName

> []WhatsappTemplate DeleteByName(ctx, wabaId, name).Execute()

Delete WhatsApp templates by name



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
    name := "sample_whatsapp_template" // string | Name of the template.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappTemplatesApi.DeleteByName(context.Background(), wabaId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappTemplatesApi.DeleteByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteByName`: []WhatsappTemplate
    fmt.Fprintf(os.Stdout, "Response from `WhatsappTemplatesApi.DeleteByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wabaId** | **string** | WhatsApp Business Account ID. | 
**name** | **string** | Name of the template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]WhatsappTemplate**](WhatsappTemplate.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditByNameAndLanguage

> WhatsappTemplate EditByNameAndLanguage(ctx, wabaId, name, language).WhatsappTemplateEditRequest(whatsappTemplateEditRequest).Execute()

Edit a WhatsApp template



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
    name := "sample_whatsapp_template" // string | Name of the template.
    language := "en" // string | Language code of the template. See [Supported Languages](https://developers.facebook.com/docs/whatsapp/api/messages/message-templates#supported-languages) for all codes.
    whatsappTemplateEditRequest := *ycloud.NewWhatsappTemplateEditRequest([]ycloud.WhatsappTemplateComponent{*ycloud.NewWhatsappTemplateComponent()}) // WhatsappTemplateEditRequest |  (optional)

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappTemplatesApi.EditByNameAndLanguage(context.Background(), wabaId, name, language).WhatsappTemplateEditRequest(whatsappTemplateEditRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappTemplatesApi.EditByNameAndLanguage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditByNameAndLanguage`: WhatsappTemplate
    fmt.Fprintf(os.Stdout, "Response from `WhatsappTemplatesApi.EditByNameAndLanguage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wabaId** | **string** | WhatsApp Business Account ID. | 
**name** | **string** | Name of the template. | 
**language** | **string** | Language code of the template. See [Supported Languages](https://developers.facebook.com/docs/whatsapp/api/messages/message-templates#supported-languages) for all codes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditByNameAndLanguageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **whatsappTemplateEditRequest** | [**WhatsappTemplateEditRequest**](WhatsappTemplateEditRequest.md) |  | 

### Return type

[**WhatsappTemplate**](WhatsappTemplate.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> WhatsappTemplatePage List(ctx).Page(page).Limit(limit).IncludeTotal(includeTotal).FilterWabaId(filterWabaId).FilterName(filterName).FilterLanguage(filterLanguage).Execute()

List WhatsApp templates



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
    filterName := "sample_whatsapp_template" // string | Name of the template. (optional)
    filterLanguage := "en" // string | Language of the template. (optional)

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappTemplatesApi.List(context.Background()).Page(page).Limit(limit).IncludeTotal(includeTotal).FilterWabaId(filterWabaId).FilterName(filterName).FilterLanguage(filterLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappTemplatesApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: WhatsappTemplatePage
    fmt.Fprintf(os.Stdout, "Response from `WhatsappTemplatesApi.List`: %v\n", resp)
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
 **filterName** | **string** | Name of the template. | 
 **filterLanguage** | **string** | Language of the template. | 

### Return type

[**WhatsappTemplatePage**](WhatsappTemplatePage.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveByNameAndLanguage

> WhatsappTemplate RetrieveByNameAndLanguage(ctx, wabaId, name, language).Execute()

Retrieve a WhatsApp template



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
    name := "sample_whatsapp_template" // string | Name of the template.
    language := "en" // string | Language code of the template. See [Supported Languages](https://developers.facebook.com/docs/whatsapp/api/messages/message-templates#supported-languages) for all codes.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappTemplatesApi.RetrieveByNameAndLanguage(context.Background(), wabaId, name, language).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappTemplatesApi.RetrieveByNameAndLanguage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveByNameAndLanguage`: WhatsappTemplate
    fmt.Fprintf(os.Stdout, "Response from `WhatsappTemplatesApi.RetrieveByNameAndLanguage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wabaId** | **string** | WhatsApp Business Account ID. | 
**name** | **string** | Name of the template. | 
**language** | **string** | Language code of the template. See [Supported Languages](https://developers.facebook.com/docs/whatsapp/api/messages/message-templates#supported-languages) for all codes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveByNameAndLanguageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WhatsappTemplate**](WhatsappTemplate.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

