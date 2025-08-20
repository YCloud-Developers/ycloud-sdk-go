# WhatsappCallingConnectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **string** | The caller&#39;s phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | 
**To** | **string** | The callee&#39;s phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | 
**SdpType** | **string** | The SDP type, must be \&quot;offer\&quot; for connection requests. | 
**Sdp** | **string** | The Session Description Protocol (SDP) offer information compliant with [RFC 8866](https://datatracker.ietf.org/doc/html/rfc8866). Contains media session parameters for establishing the WebRTC connection. | 

## Methods

### NewWhatsappCallingConnectRequest

`func NewWhatsappCallingConnectRequest(from string, to string, sdpType string, sdp string, ) *WhatsappCallingConnectRequest`

NewWhatsappCallingConnectRequest instantiates a new WhatsappCallingConnectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappCallingConnectRequestWithDefaults

`func NewWhatsappCallingConnectRequestWithDefaults() *WhatsappCallingConnectRequest`

NewWhatsappCallingConnectRequestWithDefaults instantiates a new WhatsappCallingConnectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *WhatsappCallingConnectRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *WhatsappCallingConnectRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *WhatsappCallingConnectRequest) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *WhatsappCallingConnectRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *WhatsappCallingConnectRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *WhatsappCallingConnectRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetSdpType

`func (o *WhatsappCallingConnectRequest) GetSdpType() string`

GetSdpType returns the SdpType field if non-nil, zero value otherwise.

### GetSdpTypeOk

`func (o *WhatsappCallingConnectRequest) GetSdpTypeOk() (*string, bool)`

GetSdpTypeOk returns a tuple with the SdpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdpType

`func (o *WhatsappCallingConnectRequest) SetSdpType(v string)`

SetSdpType sets SdpType field to given value.


### GetSdp

`func (o *WhatsappCallingConnectRequest) GetSdp() string`

GetSdp returns the Sdp field if non-nil, zero value otherwise.

### GetSdpOk

`func (o *WhatsappCallingConnectRequest) GetSdpOk() (*string, bool)`

GetSdpOk returns a tuple with the Sdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdp

`func (o *WhatsappCallingConnectRequest) SetSdp(v string)`

SetSdp sets Sdp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


