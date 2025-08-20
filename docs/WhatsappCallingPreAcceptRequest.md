# WhatsappCallingPreAcceptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneId** | **string** | The WhatsApp Business phone number ID. | 
**Wacid** | **string** | The WhatsApp call ID. Required for inbound call operations. This ID is received from the Call Connect webhook when a WhatsApp user initiates the call. | 
**SdpType** | **string** | The SDP type for pre-accept operations. Must be \&quot;answer\&quot;. | 
**Sdp** | **string** | The Session Description Protocol (SDP) information compliant with [RFC 8866](https://datatracker.ietf.org/doc/html/rfc8866). Contains media session parameters for the WebRTC connection. | 

## Methods

### NewWhatsappCallingPreAcceptRequest

`func NewWhatsappCallingPreAcceptRequest(phoneId string, wacid string, sdpType string, sdp string, ) *WhatsappCallingPreAcceptRequest`

NewWhatsappCallingPreAcceptRequest instantiates a new WhatsappCallingPreAcceptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappCallingPreAcceptRequestWithDefaults

`func NewWhatsappCallingPreAcceptRequestWithDefaults() *WhatsappCallingPreAcceptRequest`

NewWhatsappCallingPreAcceptRequestWithDefaults instantiates a new WhatsappCallingPreAcceptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneId

`func (o *WhatsappCallingPreAcceptRequest) GetPhoneId() string`

GetPhoneId returns the PhoneId field if non-nil, zero value otherwise.

### GetPhoneIdOk

`func (o *WhatsappCallingPreAcceptRequest) GetPhoneIdOk() (*string, bool)`

GetPhoneIdOk returns a tuple with the PhoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneId

`func (o *WhatsappCallingPreAcceptRequest) SetPhoneId(v string)`

SetPhoneId sets PhoneId field to given value.


### GetWacid

`func (o *WhatsappCallingPreAcceptRequest) GetWacid() string`

GetWacid returns the Wacid field if non-nil, zero value otherwise.

### GetWacidOk

`func (o *WhatsappCallingPreAcceptRequest) GetWacidOk() (*string, bool)`

GetWacidOk returns a tuple with the Wacid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWacid

`func (o *WhatsappCallingPreAcceptRequest) SetWacid(v string)`

SetWacid sets Wacid field to given value.


### GetSdpType

`func (o *WhatsappCallingPreAcceptRequest) GetSdpType() string`

GetSdpType returns the SdpType field if non-nil, zero value otherwise.

### GetSdpTypeOk

`func (o *WhatsappCallingPreAcceptRequest) GetSdpTypeOk() (*string, bool)`

GetSdpTypeOk returns a tuple with the SdpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdpType

`func (o *WhatsappCallingPreAcceptRequest) SetSdpType(v string)`

SetSdpType sets SdpType field to given value.


### GetSdp

`func (o *WhatsappCallingPreAcceptRequest) GetSdp() string`

GetSdp returns the Sdp field if non-nil, zero value otherwise.

### GetSdpOk

`func (o *WhatsappCallingPreAcceptRequest) GetSdpOk() (*string, bool)`

GetSdpOk returns a tuple with the Sdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdp

`func (o *WhatsappCallingPreAcceptRequest) SetSdp(v string)`

SetSdp sets Sdp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


