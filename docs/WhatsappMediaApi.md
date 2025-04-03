# WhatsappMediaApi

All URIs are relative to *https://api.ycloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Upload**](WhatsappMediaApi.md#Upload) | **Post** /whatsapp/media/{phoneNumber}/upload | Upload media



## Upload

> WhatsappMediaUpload200Response Upload(ctx, phoneNumber).File(file).Execute()

Upload media



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
    phoneNumber := "+16315551111" // string | Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format to use for the upload.
    file := os.NewFile(1234, "some_file") // *os.File | The media file to upload. Only one file is supported. If multiple files are uploaded, only the first file will be processed.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappMediaApi.Upload(context.Background(), phoneNumber).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappMediaApi.Upload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Upload`: WhatsappMediaUpload200Response
    fmt.Fprintf(os.Stdout, "Response from `WhatsappMediaApi.Upload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumber** | **string** | Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format to use for the upload. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | The media file to upload. Only one file is supported. If multiple files are uploaded, only the first file will be processed. | 

### Return type

[**WhatsappMediaUpload200Response**](WhatsappMediaUpload200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

