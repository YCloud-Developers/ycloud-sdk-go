# WhatsappCallingTerminateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneId** | **string** | The WhatsApp Business phone number ID. | 
**Wacid** | **string** | The WhatsApp call ID. Required for terminate operations. This ID is received from the Call Connect webhook when a WhatsApp user initiates the call. | 

## Methods

### NewWhatsappCallingTerminateRequest

`func NewWhatsappCallingTerminateRequest(phoneId string, wacid string, ) *WhatsappCallingTerminateRequest`

NewWhatsappCallingTerminateRequest instantiates a new WhatsappCallingTerminateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappCallingTerminateRequestWithDefaults

`func NewWhatsappCallingTerminateRequestWithDefaults() *WhatsappCallingTerminateRequest`

NewWhatsappCallingTerminateRequestWithDefaults instantiates a new WhatsappCallingTerminateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneId

`func (o *WhatsappCallingTerminateRequest) GetPhoneId() string`

GetPhoneId returns the PhoneId field if non-nil, zero value otherwise.

### GetPhoneIdOk

`func (o *WhatsappCallingTerminateRequest) GetPhoneIdOk() (*string, bool)`

GetPhoneIdOk returns a tuple with the PhoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneId

`func (o *WhatsappCallingTerminateRequest) SetPhoneId(v string)`

SetPhoneId sets PhoneId field to given value.


### GetWacid

`func (o *WhatsappCallingTerminateRequest) GetWacid() string`

GetWacid returns the Wacid field if non-nil, zero value otherwise.

### GetWacidOk

`func (o *WhatsappCallingTerminateRequest) GetWacidOk() (*string, bool)`

GetWacidOk returns a tuple with the Wacid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWacid

`func (o *WhatsappCallingTerminateRequest) SetWacid(v string)`

SetWacid sets Wacid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


