# WhatsappFlowCreate400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | Whether the operation was successful. | [optional] 
**ValidationErrors** | Pointer to [**[]WhatsappFlowValidationError**](WhatsappFlowValidationError.md) | List of validation errors. | [optional] 

## Methods

### NewWhatsappFlowCreate400Response

`func NewWhatsappFlowCreate400Response() *WhatsappFlowCreate400Response`

NewWhatsappFlowCreate400Response instantiates a new WhatsappFlowCreate400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappFlowCreate400ResponseWithDefaults

`func NewWhatsappFlowCreate400ResponseWithDefaults() *WhatsappFlowCreate400Response`

NewWhatsappFlowCreate400ResponseWithDefaults instantiates a new WhatsappFlowCreate400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *WhatsappFlowCreate400Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *WhatsappFlowCreate400Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *WhatsappFlowCreate400Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *WhatsappFlowCreate400Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetValidationErrors

`func (o *WhatsappFlowCreate400Response) GetValidationErrors() []WhatsappFlowValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *WhatsappFlowCreate400Response) GetValidationErrorsOk() (*[]WhatsappFlowValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *WhatsappFlowCreate400Response) SetValidationErrors(v []WhatsappFlowValidationError)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *WhatsappFlowCreate400Response) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


