# VerifyApi

All URIs are relative to *https://api.ycloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Check**](VerifyApi.md#Check) | **Post** /verify/verificationChecks | Check a verification
[**Send**](VerifyApi.md#Send) | **Post** /verify/verifications | Start a verification



## Check

> VerificationCheck Check(ctx).VerificationCheckRequest(verificationCheckRequest).Execute()

Check a verification



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
    verificationCheckRequest := *ycloud.NewVerificationCheckRequest() // VerificationCheckRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.VerifyApi.Check(context.Background()).VerificationCheckRequest(verificationCheckRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VerifyApi.Check``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Check`: VerificationCheck
    fmt.Fprintf(os.Stdout, "Response from `VerifyApi.Check`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verificationCheckRequest** | [**VerificationCheckRequest**](VerificationCheckRequest.md) |  | 

### Return type

[**VerificationCheck**](VerificationCheck.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Send

> Verification Send(ctx).VerificationSendRequest(verificationSendRequest).Execute()

Start a verification



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
    verificationSendRequest := *ycloud.NewVerificationSendRequest(ycloud.VerificationChannel("sms"), "+447901614024") // VerificationSendRequest | Verification request that needs to be sent.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.VerifyApi.Send(context.Background()).VerificationSendRequest(verificationSendRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VerifyApi.Send``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Send`: Verification
    fmt.Fprintf(os.Stdout, "Response from `VerifyApi.Send`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verificationSendRequest** | [**VerificationSendRequest**](VerificationSendRequest.md) | Verification request that needs to be sent. | 

### Return type

[**Verification**](Verification.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

