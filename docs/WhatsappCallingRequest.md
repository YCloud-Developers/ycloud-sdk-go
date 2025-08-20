# WhatsappCallingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneId** | Pointer to **string** | The WhatsApp Business phone number ID. | [optional] 
**From** | Pointer to **string** | The caller&#39;s phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. Required for connect operations when phoneId is empty. | [optional] 
**To** | Pointer to **string** | The callee&#39;s phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. Required for outbound call connections. | [optional] 
**Wacid** | Pointer to **string** | The WhatsApp call ID. Required for inbound call operations. This ID is received from the Call Connect webhook when a WhatsApp user initiates the call. | [optional] 
**SdpType** | Pointer to **string** | The SDP type. For pre-accept and accept operations, must be \&quot;answer\&quot;. | [optional] 
**Sdp** | Pointer to **string** | The Session Description Protocol (SDP) information compliant with [RFC 8866](https://datatracker.ietf.org/doc/html/rfc8866). Required for pre-accept and accept operations. Contains media session parameters for the WebRTC connection. | [optional] 

## Methods

### NewWhatsappCallingRequest

`func NewWhatsappCallingRequest() *WhatsappCallingRequest`

NewWhatsappCallingRequest instantiates a new WhatsappCallingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappCallingRequestWithDefaults

`func NewWhatsappCallingRequestWithDefaults() *WhatsappCallingRequest`

NewWhatsappCallingRequestWithDefaults instantiates a new WhatsappCallingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneId

`func (o *WhatsappCallingRequest) GetPhoneId() string`

GetPhoneId returns the PhoneId field if non-nil, zero value otherwise.

### GetPhoneIdOk

`func (o *WhatsappCallingRequest) GetPhoneIdOk() (*string, bool)`

GetPhoneIdOk returns a tuple with the PhoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneId

`func (o *WhatsappCallingRequest) SetPhoneId(v string)`

SetPhoneId sets PhoneId field to given value.

### HasPhoneId

`func (o *WhatsappCallingRequest) HasPhoneId() bool`

HasPhoneId returns a boolean if a field has been set.

### GetFrom

`func (o *WhatsappCallingRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *WhatsappCallingRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *WhatsappCallingRequest) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *WhatsappCallingRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *WhatsappCallingRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *WhatsappCallingRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *WhatsappCallingRequest) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *WhatsappCallingRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetWacid

`func (o *WhatsappCallingRequest) GetWacid() string`

GetWacid returns the Wacid field if non-nil, zero value otherwise.

### GetWacidOk

`func (o *WhatsappCallingRequest) GetWacidOk() (*string, bool)`

GetWacidOk returns a tuple with the Wacid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWacid

`func (o *WhatsappCallingRequest) SetWacid(v string)`

SetWacid sets Wacid field to given value.

### HasWacid

`func (o *WhatsappCallingRequest) HasWacid() bool`

HasWacid returns a boolean if a field has been set.

### GetSdpType

`func (o *WhatsappCallingRequest) GetSdpType() string`

GetSdpType returns the SdpType field if non-nil, zero value otherwise.

### GetSdpTypeOk

`func (o *WhatsappCallingRequest) GetSdpTypeOk() (*string, bool)`

GetSdpTypeOk returns a tuple with the SdpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdpType

`func (o *WhatsappCallingRequest) SetSdpType(v string)`

SetSdpType sets SdpType field to given value.

### HasSdpType

`func (o *WhatsappCallingRequest) HasSdpType() bool`

HasSdpType returns a boolean if a field has been set.

### GetSdp

`func (o *WhatsappCallingRequest) GetSdp() string`

GetSdp returns the Sdp field if non-nil, zero value otherwise.

### GetSdpOk

`func (o *WhatsappCallingRequest) GetSdpOk() (*string, bool)`

GetSdpOk returns a tuple with the Sdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdp

`func (o *WhatsappCallingRequest) SetSdp(v string)`

SetSdp sets Sdp field to given value.

### HasSdp

`func (o *WhatsappCallingRequest) HasSdp() bool`

HasSdp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


