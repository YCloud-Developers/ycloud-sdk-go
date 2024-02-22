# WhatsappMessageTemplateComponentParameterAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThumbnailProductRetailerId** | Pointer to **string** | **Optional.** Use for catalog and MPM template messages. Item SKU number. Labeled as Content ID in the Commerce Manager. The thumbnail of this item will be used as the message&#39;s header image. If the &#x60;parameters&#x60; object is omitted, the product image of the first item in your catalog will be used. | [optional] 
**Sections** | Pointer to [**[]WhatsappMessageTemplateComponentParameterActionSection**](WhatsappMessageTemplateComponentParameterActionSection.md) | Use for MPM templates. Product sections. You can define up to 10 sections. | [optional] 

## Methods

### NewWhatsappMessageTemplateComponentParameterAction

`func NewWhatsappMessageTemplateComponentParameterAction() *WhatsappMessageTemplateComponentParameterAction`

NewWhatsappMessageTemplateComponentParameterAction instantiates a new WhatsappMessageTemplateComponentParameterAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageTemplateComponentParameterActionWithDefaults

`func NewWhatsappMessageTemplateComponentParameterActionWithDefaults() *WhatsappMessageTemplateComponentParameterAction`

NewWhatsappMessageTemplateComponentParameterActionWithDefaults instantiates a new WhatsappMessageTemplateComponentParameterAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThumbnailProductRetailerId

`func (o *WhatsappMessageTemplateComponentParameterAction) GetThumbnailProductRetailerId() string`

GetThumbnailProductRetailerId returns the ThumbnailProductRetailerId field if non-nil, zero value otherwise.

### GetThumbnailProductRetailerIdOk

`func (o *WhatsappMessageTemplateComponentParameterAction) GetThumbnailProductRetailerIdOk() (*string, bool)`

GetThumbnailProductRetailerIdOk returns a tuple with the ThumbnailProductRetailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailProductRetailerId

`func (o *WhatsappMessageTemplateComponentParameterAction) SetThumbnailProductRetailerId(v string)`

SetThumbnailProductRetailerId sets ThumbnailProductRetailerId field to given value.

### HasThumbnailProductRetailerId

`func (o *WhatsappMessageTemplateComponentParameterAction) HasThumbnailProductRetailerId() bool`

HasThumbnailProductRetailerId returns a boolean if a field has been set.

### GetSections

`func (o *WhatsappMessageTemplateComponentParameterAction) GetSections() []WhatsappMessageTemplateComponentParameterActionSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *WhatsappMessageTemplateComponentParameterAction) GetSectionsOk() (*[]WhatsappMessageTemplateComponentParameterActionSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *WhatsappMessageTemplateComponentParameterAction) SetSections(v []WhatsappMessageTemplateComponentParameterActionSection)`

SetSections sets Sections field to given value.

### HasSections

`func (o *WhatsappMessageTemplateComponentParameterAction) HasSections() bool`

HasSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
