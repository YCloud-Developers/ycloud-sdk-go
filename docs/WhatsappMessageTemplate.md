# WhatsappMessageTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the template. | 
**Language** | [**WhatsappMessageTemplateLanguage**](WhatsappMessageTemplateLanguage.md) |  | 
**Components** | Pointer to [**[]WhatsappMessageTemplateComponentsInner**](WhatsappMessageTemplateComponentsInner.md) | Array of components objects containing the parameters of the message. | [optional] 

## Methods

### NewWhatsappMessageTemplate

`func NewWhatsappMessageTemplate(name string, language WhatsappMessageTemplateLanguage, ) *WhatsappMessageTemplate`

NewWhatsappMessageTemplate instantiates a new WhatsappMessageTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageTemplateWithDefaults

`func NewWhatsappMessageTemplateWithDefaults() *WhatsappMessageTemplate`

NewWhatsappMessageTemplateWithDefaults instantiates a new WhatsappMessageTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WhatsappMessageTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WhatsappMessageTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WhatsappMessageTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetLanguage

`func (o *WhatsappMessageTemplate) GetLanguage() WhatsappMessageTemplateLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *WhatsappMessageTemplate) GetLanguageOk() (*WhatsappMessageTemplateLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *WhatsappMessageTemplate) SetLanguage(v WhatsappMessageTemplateLanguage)`

SetLanguage sets Language field to given value.


### GetComponents

`func (o *WhatsappMessageTemplate) GetComponents() []WhatsappMessageTemplateComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *WhatsappMessageTemplate) GetComponentsOk() (*[]WhatsappMessageTemplateComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *WhatsappMessageTemplate) SetComponents(v []WhatsappMessageTemplateComponentsInner)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *WhatsappMessageTemplate) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


