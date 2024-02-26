# WhatsappInboundMessageError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The error code. | [optional] 
**Title** | Pointer to **string** | The error title. | [optional] 
**Message** | Pointer to **string** | The error message. | [optional] 
**ErrorData** | Pointer to **map[string]interface{}** | An error data object with the following properties: - &#x60;details&#x60;: A string describing the reason for the error. Example: &#x60;Message type is currently not supported.&#x60;. | [optional] 

## Methods

### NewWhatsappInboundMessageError

`func NewWhatsappInboundMessageError() *WhatsappInboundMessageError`

NewWhatsappInboundMessageError instantiates a new WhatsappInboundMessageError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappInboundMessageErrorWithDefaults

`func NewWhatsappInboundMessageErrorWithDefaults() *WhatsappInboundMessageError`

NewWhatsappInboundMessageErrorWithDefaults instantiates a new WhatsappInboundMessageError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *WhatsappInboundMessageError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *WhatsappInboundMessageError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *WhatsappInboundMessageError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *WhatsappInboundMessageError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetTitle

`func (o *WhatsappInboundMessageError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WhatsappInboundMessageError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WhatsappInboundMessageError) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WhatsappInboundMessageError) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMessage

`func (o *WhatsappInboundMessageError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WhatsappInboundMessageError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WhatsappInboundMessageError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *WhatsappInboundMessageError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrorData

`func (o *WhatsappInboundMessageError) GetErrorData() map[string]interface{}`

GetErrorData returns the ErrorData field if non-nil, zero value otherwise.

### GetErrorDataOk

`func (o *WhatsappInboundMessageError) GetErrorDataOk() (*map[string]interface{}, bool)`

GetErrorDataOk returns a tuple with the ErrorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorData

`func (o *WhatsappInboundMessageError) SetErrorData(v map[string]interface{})`

SetErrorData sets ErrorData field to given value.

### HasErrorData

`func (o *WhatsappInboundMessageError) HasErrorData() bool`

HasErrorData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


