# WhatsappMessageInteractiveActionSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | **Required if the message has more than one section.** Title of the section. Maximum length: 24 characters. | [optional] 
**Rows** | Pointer to [**[]WhatsappMessageInteractiveActionSectionRow**](WhatsappMessageInteractiveActionSectionRow.md) | Contains a list of rows. You can have a total of 10 rows across your sections. Each row must have a title (Maximum length: 24 characters) and an ID (Maximum length: 200 characters). You can add a description (Maximum length: 72 characters), but it is optional. | [optional] 
**ProductItems** | Pointer to [**[]WhatsappMessageInteractiveActionSectionProductItem**](WhatsappMessageInteractiveActionSectionProductItem.md) | Required for Multi-Product Messages. Array of product objects. There is a minimum of 1 product per section and a maximum of 30 products across all sections. | [optional] 

## Methods

### NewWhatsappMessageInteractiveActionSection

`func NewWhatsappMessageInteractiveActionSection() *WhatsappMessageInteractiveActionSection`

NewWhatsappMessageInteractiveActionSection instantiates a new WhatsappMessageInteractiveActionSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageInteractiveActionSectionWithDefaults

`func NewWhatsappMessageInteractiveActionSectionWithDefaults() *WhatsappMessageInteractiveActionSection`

NewWhatsappMessageInteractiveActionSectionWithDefaults instantiates a new WhatsappMessageInteractiveActionSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *WhatsappMessageInteractiveActionSection) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WhatsappMessageInteractiveActionSection) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WhatsappMessageInteractiveActionSection) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WhatsappMessageInteractiveActionSection) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetRows

`func (o *WhatsappMessageInteractiveActionSection) GetRows() []WhatsappMessageInteractiveActionSectionRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *WhatsappMessageInteractiveActionSection) GetRowsOk() (*[]WhatsappMessageInteractiveActionSectionRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *WhatsappMessageInteractiveActionSection) SetRows(v []WhatsappMessageInteractiveActionSectionRow)`

SetRows sets Rows field to given value.

### HasRows

`func (o *WhatsappMessageInteractiveActionSection) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetProductItems

`func (o *WhatsappMessageInteractiveActionSection) GetProductItems() []WhatsappMessageInteractiveActionSectionProductItem`

GetProductItems returns the ProductItems field if non-nil, zero value otherwise.

### GetProductItemsOk

`func (o *WhatsappMessageInteractiveActionSection) GetProductItemsOk() (*[]WhatsappMessageInteractiveActionSectionProductItem, bool)`

GetProductItemsOk returns a tuple with the ProductItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductItems

`func (o *WhatsappMessageInteractiveActionSection) SetProductItems(v []WhatsappMessageInteractiveActionSectionProductItem)`

SetProductItems sets ProductItems field to given value.

### HasProductItems

`func (o *WhatsappMessageInteractiveActionSection) HasProductItems() bool`

HasProductItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
