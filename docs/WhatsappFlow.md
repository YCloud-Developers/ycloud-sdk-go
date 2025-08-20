# WhatsappFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Flow ID. | [optional] 
**Name** | Pointer to **string** | Flow name. | [optional] 
**Status** | Pointer to [**WhatsappFlowStatus**](WhatsappFlowStatus.md) |  | [optional] 
**Categories** | Pointer to [**[]WhatsappFlowCategory**](WhatsappFlowCategory.md) | Flow categories. | [optional] 
**WhatsappBusinessAccount** | Pointer to [**WhatsappFlowWhatsappBusinessAccount**](WhatsappFlowWhatsappBusinessAccount.md) |  | [optional] 
**ValidationErrors** | Pointer to [**[]WhatsappFlowValidationError**](WhatsappFlowValidationError.md) | List of validation errors. | [optional] 
**JsonVersion** | Pointer to **string** | Version of the Flow JSON structure. | [optional] 
**DataApiVersion** | Pointer to **string** | Version of the Data API. | [optional] 

## Methods

### NewWhatsappFlow

`func NewWhatsappFlow() *WhatsappFlow`

NewWhatsappFlow instantiates a new WhatsappFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappFlowWithDefaults

`func NewWhatsappFlowWithDefaults() *WhatsappFlow`

NewWhatsappFlowWithDefaults instantiates a new WhatsappFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WhatsappFlow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WhatsappFlow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WhatsappFlow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WhatsappFlow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WhatsappFlow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WhatsappFlow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WhatsappFlow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WhatsappFlow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *WhatsappFlow) GetStatus() WhatsappFlowStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WhatsappFlow) GetStatusOk() (*WhatsappFlowStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WhatsappFlow) SetStatus(v WhatsappFlowStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WhatsappFlow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCategories

`func (o *WhatsappFlow) GetCategories() []WhatsappFlowCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *WhatsappFlow) GetCategoriesOk() (*[]WhatsappFlowCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *WhatsappFlow) SetCategories(v []WhatsappFlowCategory)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *WhatsappFlow) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetWhatsappBusinessAccount

`func (o *WhatsappFlow) GetWhatsappBusinessAccount() WhatsappFlowWhatsappBusinessAccount`

GetWhatsappBusinessAccount returns the WhatsappBusinessAccount field if non-nil, zero value otherwise.

### GetWhatsappBusinessAccountOk

`func (o *WhatsappFlow) GetWhatsappBusinessAccountOk() (*WhatsappFlowWhatsappBusinessAccount, bool)`

GetWhatsappBusinessAccountOk returns a tuple with the WhatsappBusinessAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhatsappBusinessAccount

`func (o *WhatsappFlow) SetWhatsappBusinessAccount(v WhatsappFlowWhatsappBusinessAccount)`

SetWhatsappBusinessAccount sets WhatsappBusinessAccount field to given value.

### HasWhatsappBusinessAccount

`func (o *WhatsappFlow) HasWhatsappBusinessAccount() bool`

HasWhatsappBusinessAccount returns a boolean if a field has been set.

### GetValidationErrors

`func (o *WhatsappFlow) GetValidationErrors() []WhatsappFlowValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *WhatsappFlow) GetValidationErrorsOk() (*[]WhatsappFlowValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *WhatsappFlow) SetValidationErrors(v []WhatsappFlowValidationError)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *WhatsappFlow) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.

### GetJsonVersion

`func (o *WhatsappFlow) GetJsonVersion() string`

GetJsonVersion returns the JsonVersion field if non-nil, zero value otherwise.

### GetJsonVersionOk

`func (o *WhatsappFlow) GetJsonVersionOk() (*string, bool)`

GetJsonVersionOk returns a tuple with the JsonVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonVersion

`func (o *WhatsappFlow) SetJsonVersion(v string)`

SetJsonVersion sets JsonVersion field to given value.

### HasJsonVersion

`func (o *WhatsappFlow) HasJsonVersion() bool`

HasJsonVersion returns a boolean if a field has been set.

### GetDataApiVersion

`func (o *WhatsappFlow) GetDataApiVersion() string`

GetDataApiVersion returns the DataApiVersion field if non-nil, zero value otherwise.

### GetDataApiVersionOk

`func (o *WhatsappFlow) GetDataApiVersionOk() (*string, bool)`

GetDataApiVersionOk returns a tuple with the DataApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataApiVersion

`func (o *WhatsappFlow) SetDataApiVersion(v string)`

SetDataApiVersion sets DataApiVersion field to given value.

### HasDataApiVersion

`func (o *WhatsappFlow) HasDataApiVersion() bool`

HasDataApiVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


