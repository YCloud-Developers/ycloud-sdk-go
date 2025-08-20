# WhatsappFlowValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Error code. | [optional] 
**ErrorType** | Pointer to **string** | Error type. | [optional] 
**Message** | Pointer to **string** | Error message. | [optional] 
**LineStart** | Pointer to **int32** | Start line of the error. | [optional] 
**LineEnd** | Pointer to **int32** | End line of the error. | [optional] 
**ColumnStart** | Pointer to **int32** | Start column of the error. | [optional] 
**ColumnEnd** | Pointer to **int32** | End column of the error. | [optional] 
**Pointers** | Pointer to [**[]WhatsappFlowValidationErrorPointersInner**](WhatsappFlowValidationErrorPointersInner.md) | List of pointers to the error location. | [optional] 

## Methods

### NewWhatsappFlowValidationError

`func NewWhatsappFlowValidationError() *WhatsappFlowValidationError`

NewWhatsappFlowValidationError instantiates a new WhatsappFlowValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappFlowValidationErrorWithDefaults

`func NewWhatsappFlowValidationErrorWithDefaults() *WhatsappFlowValidationError`

NewWhatsappFlowValidationErrorWithDefaults instantiates a new WhatsappFlowValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *WhatsappFlowValidationError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WhatsappFlowValidationError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WhatsappFlowValidationError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *WhatsappFlowValidationError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorType

`func (o *WhatsappFlowValidationError) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *WhatsappFlowValidationError) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *WhatsappFlowValidationError) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *WhatsappFlowValidationError) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### GetMessage

`func (o *WhatsappFlowValidationError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WhatsappFlowValidationError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WhatsappFlowValidationError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *WhatsappFlowValidationError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetLineStart

`func (o *WhatsappFlowValidationError) GetLineStart() int32`

GetLineStart returns the LineStart field if non-nil, zero value otherwise.

### GetLineStartOk

`func (o *WhatsappFlowValidationError) GetLineStartOk() (*int32, bool)`

GetLineStartOk returns a tuple with the LineStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineStart

`func (o *WhatsappFlowValidationError) SetLineStart(v int32)`

SetLineStart sets LineStart field to given value.

### HasLineStart

`func (o *WhatsappFlowValidationError) HasLineStart() bool`

HasLineStart returns a boolean if a field has been set.

### GetLineEnd

`func (o *WhatsappFlowValidationError) GetLineEnd() int32`

GetLineEnd returns the LineEnd field if non-nil, zero value otherwise.

### GetLineEndOk

`func (o *WhatsappFlowValidationError) GetLineEndOk() (*int32, bool)`

GetLineEndOk returns a tuple with the LineEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineEnd

`func (o *WhatsappFlowValidationError) SetLineEnd(v int32)`

SetLineEnd sets LineEnd field to given value.

### HasLineEnd

`func (o *WhatsappFlowValidationError) HasLineEnd() bool`

HasLineEnd returns a boolean if a field has been set.

### GetColumnStart

`func (o *WhatsappFlowValidationError) GetColumnStart() int32`

GetColumnStart returns the ColumnStart field if non-nil, zero value otherwise.

### GetColumnStartOk

`func (o *WhatsappFlowValidationError) GetColumnStartOk() (*int32, bool)`

GetColumnStartOk returns a tuple with the ColumnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnStart

`func (o *WhatsappFlowValidationError) SetColumnStart(v int32)`

SetColumnStart sets ColumnStart field to given value.

### HasColumnStart

`func (o *WhatsappFlowValidationError) HasColumnStart() bool`

HasColumnStart returns a boolean if a field has been set.

### GetColumnEnd

`func (o *WhatsappFlowValidationError) GetColumnEnd() int32`

GetColumnEnd returns the ColumnEnd field if non-nil, zero value otherwise.

### GetColumnEndOk

`func (o *WhatsappFlowValidationError) GetColumnEndOk() (*int32, bool)`

GetColumnEndOk returns a tuple with the ColumnEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnEnd

`func (o *WhatsappFlowValidationError) SetColumnEnd(v int32)`

SetColumnEnd sets ColumnEnd field to given value.

### HasColumnEnd

`func (o *WhatsappFlowValidationError) HasColumnEnd() bool`

HasColumnEnd returns a boolean if a field has been set.

### GetPointers

`func (o *WhatsappFlowValidationError) GetPointers() []WhatsappFlowValidationErrorPointersInner`

GetPointers returns the Pointers field if non-nil, zero value otherwise.

### GetPointersOk

`func (o *WhatsappFlowValidationError) GetPointersOk() (*[]WhatsappFlowValidationErrorPointersInner, bool)`

GetPointersOk returns a tuple with the Pointers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointers

`func (o *WhatsappFlowValidationError) SetPointers(v []WhatsappFlowValidationErrorPointersInner)`

SetPointers sets Pointers field to given value.

### HasPointers

`func (o *WhatsappFlowValidationError) HasPointers() bool`

HasPointers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


