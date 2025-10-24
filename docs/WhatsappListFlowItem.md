# WhatsappListFlowItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Flow ID. | [optional] 
**Name** | Pointer to **string** | Flow name. | [optional] 
**Status** | Pointer to [**WhatsappFlowStatus**](WhatsappFlowStatus.md) |  | [optional] 
**Categories** | Pointer to [**[]WhatsappFlowCategory**](WhatsappFlowCategory.md) | Flow categories. | [optional] 
**ValidationErrors** | Pointer to [**[]WhatsappFlowValidationError**](WhatsappFlowValidationError.md) | List of validation errors. | [optional] 

## Methods

### NewWhatsappListFlowItem

`func NewWhatsappListFlowItem() *WhatsappListFlowItem`

NewWhatsappListFlowItem instantiates a new WhatsappListFlowItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappListFlowItemWithDefaults

`func NewWhatsappListFlowItemWithDefaults() *WhatsappListFlowItem`

NewWhatsappListFlowItemWithDefaults instantiates a new WhatsappListFlowItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WhatsappListFlowItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WhatsappListFlowItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WhatsappListFlowItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WhatsappListFlowItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WhatsappListFlowItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WhatsappListFlowItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WhatsappListFlowItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WhatsappListFlowItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *WhatsappListFlowItem) GetStatus() WhatsappFlowStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WhatsappListFlowItem) GetStatusOk() (*WhatsappFlowStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WhatsappListFlowItem) SetStatus(v WhatsappFlowStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WhatsappListFlowItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCategories

`func (o *WhatsappListFlowItem) GetCategories() []WhatsappFlowCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *WhatsappListFlowItem) GetCategoriesOk() (*[]WhatsappFlowCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *WhatsappListFlowItem) SetCategories(v []WhatsappFlowCategory)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *WhatsappListFlowItem) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetValidationErrors

`func (o *WhatsappListFlowItem) GetValidationErrors() []WhatsappFlowValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *WhatsappListFlowItem) GetValidationErrorsOk() (*[]WhatsappFlowValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *WhatsappListFlowItem) SetValidationErrors(v []WhatsappFlowValidationError)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *WhatsappListFlowItem) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


