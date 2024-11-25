# WhatsappTemplateCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WabaId** | **string** | WhatsApp Business Account ID. | 
**Name** | **string** | Name of the template. | 
**Language** | **string** | Language code of the template. See [Supported Languages](https://developers.facebook.com/docs/whatsapp/api/messages/message-templates#supported-languages) for all codes. | 
**Category** | [**WhatsappTemplateCategory**](WhatsappTemplateCategory.md) |  | 
**SubCategory** | Pointer to [**WhatsappTemplateSubCategory**](WhatsappTemplateSubCategory.md) |  | [optional] 
**MessageSendTtlSeconds** | Pointer to **int32** | If we are unable to deliver a message for an amount of time that exceeds its time-to-live, we will stop retrying and drop the message. By default, messages that use an authentication template have a default TTL of **10 minutes**, and messages that use a utility or marketing template have a default TTL of **30 days**. Set its value between &#x60;30&#x60; and &#x60;900&#x60; seconds (i.e., 30 seconds to 15 minutes) for authentication templates, or &#x60;30&#x60; and &#x60;43200&#x60; seconds (i.e., 30 seconds to 12 hours) for utility templates, or &#x60;43200&#x60; and &#x60;2592000&#x60; seconds (i.e., 12 hours to 30 days) for marketing templates. Alternatively, you can set this value to &#x60;-1&#x60;, which will set a custom TTL of 30 days for either type of template. We encourage you to set a time-to-live for all of your authentication templates, preferably equal to or less than your code expiration time, to ensure your customers only get a message when a code is still usable. Authentication templates created before October 23, 2024, have a default TTL of 30 days. | [optional] 
**Components** | [**[]WhatsappTemplateComponent**](WhatsappTemplateComponent.md) |  | 

## Methods

### NewWhatsappTemplateCreateRequest

`func NewWhatsappTemplateCreateRequest(wabaId string, name string, language string, category WhatsappTemplateCategory, components []WhatsappTemplateComponent, ) *WhatsappTemplateCreateRequest`

NewWhatsappTemplateCreateRequest instantiates a new WhatsappTemplateCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappTemplateCreateRequestWithDefaults

`func NewWhatsappTemplateCreateRequestWithDefaults() *WhatsappTemplateCreateRequest`

NewWhatsappTemplateCreateRequestWithDefaults instantiates a new WhatsappTemplateCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWabaId

`func (o *WhatsappTemplateCreateRequest) GetWabaId() string`

GetWabaId returns the WabaId field if non-nil, zero value otherwise.

### GetWabaIdOk

`func (o *WhatsappTemplateCreateRequest) GetWabaIdOk() (*string, bool)`

GetWabaIdOk returns a tuple with the WabaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWabaId

`func (o *WhatsappTemplateCreateRequest) SetWabaId(v string)`

SetWabaId sets WabaId field to given value.


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


### GetSubCategory

`func (o *WhatsappTemplateCreateRequest) GetSubCategory() WhatsappTemplateSubCategory`

GetSubCategory returns the SubCategory field if non-nil, zero value otherwise.

### GetSubCategoryOk

`func (o *WhatsappTemplateCreateRequest) GetSubCategoryOk() (*WhatsappTemplateSubCategory, bool)`

GetSubCategoryOk returns a tuple with the SubCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategory

`func (o *WhatsappTemplateCreateRequest) SetSubCategory(v WhatsappTemplateSubCategory)`

SetSubCategory sets SubCategory field to given value.

### HasSubCategory

`func (o *WhatsappTemplateCreateRequest) HasSubCategory() bool`

HasSubCategory returns a boolean if a field has been set.

### GetMessageSendTtlSeconds

`func (o *WhatsappTemplateCreateRequest) GetMessageSendTtlSeconds() int32`

GetMessageSendTtlSeconds returns the MessageSendTtlSeconds field if non-nil, zero value otherwise.

### GetMessageSendTtlSecondsOk

`func (o *WhatsappTemplateCreateRequest) GetMessageSendTtlSecondsOk() (*int32, bool)`

GetMessageSendTtlSecondsOk returns a tuple with the MessageSendTtlSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSendTtlSeconds

`func (o *WhatsappTemplateCreateRequest) SetMessageSendTtlSeconds(v int32)`

SetMessageSendTtlSeconds sets MessageSendTtlSeconds field to given value.

### HasMessageSendTtlSeconds

`func (o *WhatsappTemplateCreateRequest) HasMessageSendTtlSeconds() bool`

HasMessageSendTtlSeconds returns a boolean if a field has been set.

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


