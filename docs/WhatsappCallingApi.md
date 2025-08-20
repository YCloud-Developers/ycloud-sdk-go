# WhatsappCallingApi

All URIs are relative to *https://api.ycloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Accept**](WhatsappCallingApi.md#Accept) | **Post** /whatsapp/calls/accept | Accept a call
[**Connect**](WhatsappCallingApi.md#Connect) | **Post** /whatsapp/calls/connect | Connect a call
[**PreAccept**](WhatsappCallingApi.md#PreAccept) | **Post** /whatsapp/calls/preAccept | Pre-accept a call
[**Reject**](WhatsappCallingApi.md#Reject) | **Post** /whatsapp/calls/reject | Reject a call
[**Terminate**](WhatsappCallingApi.md#Terminate) | **Post** /whatsapp/calls/terminate | Terminate a call



## Accept

> WhatsappCallingResponse Accept(ctx).WhatsappCallingPreAcceptRequest(whatsappCallingPreAcceptRequest).Execute()

Accept a call



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
    whatsappCallingPreAcceptRequest := *ycloud.NewWhatsappCallingPreAcceptRequest("461269257068832", "wacid.HBgNNjI4MTM2MTkwNTEzMxUCABIYIDhEMjc4NEY2QUU4NTA2MTgxNTBGNzQ3N0M4QTBDMTU5HBgNNjI4MzEzODIwNTE1MBUCABUeAA==", "answer", "v=0
o=- 2239925877841361960 2 IN IP4 127.0.0.1
s=-
t=0 0
a=group:BUNDLE audio
a=msid-semantic: WMS 0ed5100f-da68-4193-8865-146c1ac7a087
m=audio 9 UDP/TLS/RTP/SAVPF 111 126
c=IN IP4 0.0.0.0
a=rtcp:9 IN IP4 0.0.0.0
a=ice-ufrag:NCH6
a=ice-pwd:ogUYxqJDPNn0C5RFif6UlLz6
a=ice-options:trickle
a=fingerprint:sha-256 5B:95:C4:E4:8B:2B:06:B6:DB:FB:2C:08:2F:FD:3B:C7:9C:8D:84:4C:97:8D:84:AC:B2:93:32:B8:20:5C:3C:85
a=setup:active
a=mid:audio
a=sendrecv
a=msid:0ed5100f-da68-4193-8865-146c1ac7a087 bc955459-4bae-4504-99c8-348944c12b6f
a=rtcp-mux
a=rtpmap:111 opus/48000/2
a=rtcp-fb:111 transport-cc
a=fmtp:111 minptime=10;useinbandfec=1
a=rtpmap:126 telephone-event/8000
a=ssrc:436995058 cname:BIzAP4IgR06SrZ1S
") // WhatsappCallingPreAcceptRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappCallingApi.Accept(context.Background()).WhatsappCallingPreAcceptRequest(whatsappCallingPreAcceptRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappCallingApi.Accept``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Accept`: WhatsappCallingResponse
    fmt.Fprintf(os.Stdout, "Response from `WhatsappCallingApi.Accept`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **whatsappCallingPreAcceptRequest** | [**WhatsappCallingPreAcceptRequest**](WhatsappCallingPreAcceptRequest.md) |  | 

### Return type

[**WhatsappCallingResponse**](WhatsappCallingResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Connect

> WhatsappCallingResponse Connect(ctx).WhatsappCallingConnectRequest(whatsappCallingConnectRequest).Execute()

Connect a call



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
    whatsappCallingConnectRequest := *ycloud.NewWhatsappCallingConnectRequest("+6283138205150", "+6281361905133", "offer", "v=0
o=- 4054442297240208280 2 IN IP4 127.0.0.1
s=-
t=0 0
a=group:BUNDLE 0
a=extmap-allow-mixed
a=msid-semantic: WMS 6c364341-f90d-48e6-b497-76e047e3c31a
m=audio 9 UDP/TLS/RTP/SAVPF 111 63 9 0 8 13 110 126
c=IN IP4 0.0.0.0
a=rtcp:9 IN IP4 0.0.0.0
a=ice-ufrag:Dpsa
a=ice-pwd:oVuOd7HKhA8aWTvspLYACWJe
a=ice-options:trickle
a=fingerprint:sha-256 0A:A9:64:82:AD:D5:31:08:38:71:1C:C0:08:AA:CE:93:22:F4:17:2C:B6:F1:8F:F1:20:71:38:16:37:18:3F:FA
a=setup:actpass
a=mid:0
a=extmap:1 urn:ietf:params:rtp-hdrext:ssrc-audio-level
a=extmap:2 http://www.webrtc.org/experiments/rtp-hdrext/abs-send-time
a=extmap:3 http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01
a=extmap:4 urn:ietf:params:rtp-hdrext:sdes:mid
a=sendrecv
a=msid:6c364341-f90d-48e6-b497-76e047e3c31a fa42cdbe-8696-4ee8-bce0-0919d86a223f
a=rtcp-mux
a=rtcp-rsize
a=rtpmap:111 opus/48000/2
a=rtcp-fb:111 transport-cc
a=fmtp:111 minptime=10;useinbandfec=1
a=rtpmap:63 red/48000/2
a=fmtp:63 111/111
a=rtpmap:9 G722/8000
a=rtpmap:0 PCMU/8000
a=rtpmap:8 PCMA/8000
a=rtpmap:13 CN/8000
a=rtpmap:110 telephone-event/48000
a=rtpmap:126 telephone-event/8000
a=ssrc:3208712354 cname:bg/Ix8uTnsTsiMoe
a=ssrc:3208712354 msid:6c364341-f90d-48e6-b497-76e047e3c31a fa42cdbe-8696-4ee8-bce0-0919d86a223f
") // WhatsappCallingConnectRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappCallingApi.Connect(context.Background()).WhatsappCallingConnectRequest(whatsappCallingConnectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappCallingApi.Connect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Connect`: WhatsappCallingResponse
    fmt.Fprintf(os.Stdout, "Response from `WhatsappCallingApi.Connect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **whatsappCallingConnectRequest** | [**WhatsappCallingConnectRequest**](WhatsappCallingConnectRequest.md) |  | 

### Return type

[**WhatsappCallingResponse**](WhatsappCallingResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreAccept

> WhatsappCallingResponse PreAccept(ctx).WhatsappCallingPreAcceptRequest(whatsappCallingPreAcceptRequest).Execute()

Pre-accept a call



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
    whatsappCallingPreAcceptRequest := *ycloud.NewWhatsappCallingPreAcceptRequest("461269257068832", "wacid.HBgNNjI4MTM2MTkwNTEzMxUCABIYIDhEMjc4NEY2QUU4NTA2MTgxNTBGNzQ3N0M4QTBDMTU5HBgNNjI4MzEzODIwNTE1MBUCABUeAA==", "answer", "v=0
o=- 2239925877841361960 2 IN IP4 127.0.0.1
s=-
t=0 0
a=group:BUNDLE audio
a=msid-semantic: WMS 0ed5100f-da68-4193-8865-146c1ac7a087
m=audio 9 UDP/TLS/RTP/SAVPF 111 126
c=IN IP4 0.0.0.0
a=rtcp:9 IN IP4 0.0.0.0
a=ice-ufrag:NCH6
a=ice-pwd:ogUYxqJDPNn0C5RFif6UlLz6
a=ice-options:trickle
a=fingerprint:sha-256 5B:95:C4:E4:8B:2B:06:B6:DB:FB:2C:08:2F:FD:3B:C7:9C:8D:84:4C:97:8D:84:AC:B2:93:32:B8:20:5C:3C:85
a=setup:active
a=mid:audio
a=sendrecv
a=msid:0ed5100f-da68-4193-8865-146c1ac7a087 bc955459-4bae-4504-99c8-348944c12b6f
a=rtcp-mux
a=rtpmap:111 opus/48000/2
a=rtcp-fb:111 transport-cc
a=fmtp:111 minptime=10;useinbandfec=1
a=rtpmap:126 telephone-event/8000
a=ssrc:436995058 cname:BIzAP4IgR06SrZ1S
") // WhatsappCallingPreAcceptRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappCallingApi.PreAccept(context.Background()).WhatsappCallingPreAcceptRequest(whatsappCallingPreAcceptRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappCallingApi.PreAccept``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreAccept`: WhatsappCallingResponse
    fmt.Fprintf(os.Stdout, "Response from `WhatsappCallingApi.PreAccept`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **whatsappCallingPreAcceptRequest** | [**WhatsappCallingPreAcceptRequest**](WhatsappCallingPreAcceptRequest.md) |  | 

### Return type

[**WhatsappCallingResponse**](WhatsappCallingResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Reject

> WhatsappCallingResponse Reject(ctx).WhatsappCallingTerminateRequest(whatsappCallingTerminateRequest).Execute()

Reject a call



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
    whatsappCallingTerminateRequest := *ycloud.NewWhatsappCallingTerminateRequest("461269257068832", "wacid.HBgNNjI4MTM2MTkwNTEzMxUCABEYIDNENjg2OEMzNTFFRDkwRkUxRUE1RTgxNjY1NjJCQUJBHBgNNjI4MzEzODIwNTE1MBUCABUeAA==") // WhatsappCallingTerminateRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappCallingApi.Reject(context.Background()).WhatsappCallingTerminateRequest(whatsappCallingTerminateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappCallingApi.Reject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Reject`: WhatsappCallingResponse
    fmt.Fprintf(os.Stdout, "Response from `WhatsappCallingApi.Reject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRejectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **whatsappCallingTerminateRequest** | [**WhatsappCallingTerminateRequest**](WhatsappCallingTerminateRequest.md) |  | 

### Return type

[**WhatsappCallingResponse**](WhatsappCallingResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Terminate

> WhatsappCallingResponse Terminate(ctx).WhatsappCallingTerminateRequest(whatsappCallingTerminateRequest).Execute()

Terminate a call



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
    whatsappCallingTerminateRequest := *ycloud.NewWhatsappCallingTerminateRequest("461269257068832", "wacid.HBgNNjI4MTM2MTkwNTEzMxUCABEYIDNENjg2OEMzNTFFRDkwRkUxRUE1RTgxNjY1NjJCQUJBHBgNNjI4MzEzODIwNTE1MBUCABUeAA==") // WhatsappCallingTerminateRequest | 

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.WhatsappCallingApi.Terminate(context.Background()).WhatsappCallingTerminateRequest(whatsappCallingTerminateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhatsappCallingApi.Terminate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Terminate`: WhatsappCallingResponse
    fmt.Fprintf(os.Stdout, "Response from `WhatsappCallingApi.Terminate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTerminateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **whatsappCallingTerminateRequest** | [**WhatsappCallingTerminateRequest**](WhatsappCallingTerminateRequest.md) |  | 

### Return type

[**WhatsappCallingResponse**](WhatsappCallingResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

