# WhatsappTemplateCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the template. | 
**Language** | **string** | Language code of the template. See [Supported Languages](https://developers.facebook.com/docs/whatsapp/api/messages/message-templates#supported-languages) for all codes. | 
**Category** | [**WhatsappTemplateCategory**](WhatsappTemplateCategory.md) |  | 
**Components** | [**[]WhatsappTemplateComponent**](WhatsappTemplateComponent.md) |  | 

## Methods

### NewWhatsappTemplateCreateRequest

`func NewWhatsappTemplateCreateRequest(name string, language string, category WhatsappTemplateCategory, components []WhatsappTemplateComponent, ) *WhatsappTemplateCreateRequest`

NewWhatsappTemplateCreateRequest instantiates a new WhatsappTemplateCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappTemplateCreateRequestWithDefaults

`func NewWhatsappTemplateCreateRequestWithDefaults() *WhatsappTemplateCreateRequest`

NewWhatsappTemplateCreateRequestWithDefaults instantiates a new WhatsappTemplateCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WhatsappTemplateCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WhatsappTemplateCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WhatsappTemplateCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLanguage

`func (o *WhatsappTemplateCreateRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *WhatsappTemplateCreateRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *WhatsappTemplateCreateRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetCategory

`func (o *WhatsappTemplateCreateRequest) GetCategory() WhatsappTemplateCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *WhatsappTemplateCreateRequest) GetCategoryOk() (*WhatsappTemplateCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *WhatsappTemplateCreateRequest) SetCategory(v WhatsappTemplateCategory)`

SetCategory sets Category field to given value.


### GetComponents

`func (o *WhatsappTemplateCreateRequest) GetComponents() []WhatsappTemplateComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *WhatsappTemplateCreateRequest) GetComponentsOk() (*[]WhatsappTemplateComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *WhatsappTemplateCreateRequest) SetComponents(v []WhatsappTemplateComponent)`

SetComponents sets Components field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


