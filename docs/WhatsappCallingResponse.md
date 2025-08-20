# WhatsappCallingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wacid** | Pointer to **string** | The WhatsApp call ID associated with this calling operation. | [optional] 
**Success** | **bool** | Indicates whether the calling operation was successful. | 

## Methods

### NewWhatsappCallingResponse

`func NewWhatsappCallingResponse(success bool, ) *WhatsappCallingResponse`

NewWhatsappCallingResponse instantiates a new WhatsappCallingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappCallingResponseWithDefaults

`func NewWhatsappCallingResponseWithDefaults() *WhatsappCallingResponse`

NewWhatsappCallingResponseWithDefaults instantiates a new WhatsappCallingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWacid

`func (o *WhatsappCallingResponse) GetWacid() string`

GetWacid returns the Wacid field if non-nil, zero value otherwise.

### GetWacidOk

`func (o *WhatsappCallingResponse) GetWacidOk() (*string, bool)`

GetWacidOk returns a tuple with the Wacid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWacid

`func (o *WhatsappCallingResponse) SetWacid(v string)`

SetWacid sets Wacid field to given value.

### HasWacid

`func (o *WhatsappCallingResponse) HasWacid() bool`

HasWacid returns a boolean if a field has been set.

### GetSuccess

`func (o *WhatsappCallingResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *WhatsappCallingResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *WhatsappCallingResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


