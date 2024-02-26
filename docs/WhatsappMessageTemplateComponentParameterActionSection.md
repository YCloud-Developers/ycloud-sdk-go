# WhatsappMessageTemplateComponentParameterActionSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Section title text. Maximum 24 characters. Markdown is not supported. | [optional] 
**ProductItems** | Pointer to [**[]WhatsappMessageTemplateComponentParameterActionSectionProductItem**](WhatsappMessageTemplateComponentParameterActionSectionProductItem.md) | Array of product SKU numbers. There is a minimum of 1 product per section and a maximum of 30 products across all sections. | [optional] 

## Methods

### NewWhatsappMessageTemplateComponentParameterActionSection

`func NewWhatsappMessageTemplateComponentParameterActionSection() *WhatsappMessageTemplateComponentParameterActionSection`

NewWhatsappMessageTemplateComponentParameterActionSection instantiates a new WhatsappMessageTemplateComponentParameterActionSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageTemplateComponentParameterActionSectionWithDefaults

`func NewWhatsappMessageTemplateComponentParameterActionSectionWithDefaults() *WhatsappMessageTemplateComponentParameterActionSection`

NewWhatsappMessageTemplateComponentParameterActionSectionWithDefaults instantiates a new WhatsappMessageTemplateComponentParameterActionSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *WhatsappMessageTemplateComponentParameterActionSection) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WhatsappMessageTemplateComponentParameterActionSection) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WhatsappMessageTemplateComponentParameterActionSection) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WhatsappMessageTemplateComponentParameterActionSection) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetProductItems

`func (o *WhatsappMessageTemplateComponentParameterActionSection) GetProductItems() []WhatsappMessageTemplateComponentParameterActionSectionProductItem`

GetProductItems returns the ProductItems field if non-nil, zero value otherwise.

### GetProductItemsOk

`func (o *WhatsappMessageTemplateComponentParameterActionSection) GetProductItemsOk() (*[]WhatsappMessageTemplateComponentParameterActionSectionProductItem, bool)`

GetProductItemsOk returns a tuple with the ProductItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductItems

`func (o *WhatsappMessageTemplateComponentParameterActionSection) SetProductItems(v []WhatsappMessageTemplateComponentParameterActionSectionProductItem)`

SetProductItems sets ProductItems field to given value.

### HasProductItems

`func (o *WhatsappMessageTemplateComponentParameterActionSection) HasProductItems() bool`

HasProductItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


