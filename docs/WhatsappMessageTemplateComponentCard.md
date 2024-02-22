# WhatsappMessageTemplateComponentCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardIndex** | Pointer to **int32** | **Required.** Zero-indexed order in which card appears within the card carousel. 0 indicates first card, 1 indicates second card, etc. | [optional] 
**Components** | Pointer to [**[]WhatsappMessageTemplateComponentCardComponent**](WhatsappMessageTemplateComponentCardComponent.md) | Card component. | [optional] 

## Methods

### NewWhatsappMessageTemplateComponentCard

`func NewWhatsappMessageTemplateComponentCard() *WhatsappMessageTemplateComponentCard`

NewWhatsappMessageTemplateComponentCard instantiates a new WhatsappMessageTemplateComponentCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageTemplateComponentCardWithDefaults

`func NewWhatsappMessageTemplateComponentCardWithDefaults() *WhatsappMessageTemplateComponentCard`

NewWhatsappMessageTemplateComponentCardWithDefaults instantiates a new WhatsappMessageTemplateComponentCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardIndex

`func (o *WhatsappMessageTemplateComponentCard) GetCardIndex() int32`

GetCardIndex returns the CardIndex field if non-nil, zero value otherwise.

### GetCardIndexOk

`func (o *WhatsappMessageTemplateComponentCard) GetCardIndexOk() (*int32, bool)`

GetCardIndexOk returns a tuple with the CardIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIndex

`func (o *WhatsappMessageTemplateComponentCard) SetCardIndex(v int32)`

SetCardIndex sets CardIndex field to given value.

### HasCardIndex

`func (o *WhatsappMessageTemplateComponentCard) HasCardIndex() bool`

HasCardIndex returns a boolean if a field has been set.

### GetComponents

`func (o *WhatsappMessageTemplateComponentCard) GetComponents() []WhatsappMessageTemplateComponentCardComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *WhatsappMessageTemplateComponentCard) GetComponentsOk() (*[]WhatsappMessageTemplateComponentCardComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *WhatsappMessageTemplateComponentCard) SetComponents(v []WhatsappMessageTemplateComponentCardComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *WhatsappMessageTemplateComponentCard) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
