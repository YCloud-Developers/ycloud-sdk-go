# WhatsappTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WabaId** | **string** | WhatsApp Business Account ID. | 
**Name** | **string** | Name of the template. | 
**Language** | **string** | Language code of the template. See [Supported Languages](https://developers.facebook.com/docs/whatsapp/api/messages/message-templates#supported-languages-) for all codes. | 
**Category** | [**WhatsappTemplateCategory**](WhatsappTemplateCategory.md) |  | 
**Components** | [**[]WhatsappTemplateComponent**](WhatsappTemplateComponent.md) |  | 
**Status** | Pointer to [**WhatsappTemplateStatus**](WhatsappTemplateStatus.md) |  | [optional] 
**Reason** | Pointer to **string** | The reason why the template is rejected. | [optional] 
**CreateTime** | Pointer to **time.Time** | The time at which this object is created, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional] 
**UpdateTime** | Pointer to **time.Time** | The time at which this object is updated, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional] 

## Methods

### NewWhatsappTemplate

`func NewWhatsappTemplate(wabaId string, name string, language string, category WhatsappTemplateCategory, components []WhatsappTemplateComponent, ) *WhatsappTemplate`

NewWhatsappTemplate instantiates a new WhatsappTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappTemplateWithDefaults

`func NewWhatsappTemplateWithDefaults() *WhatsappTemplate`

NewWhatsappTemplateWithDefaults instantiates a new WhatsappTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWabaId

`func (o *WhatsappTemplate) GetWabaId() string`

GetWabaId returns the WabaId field if non-nil, zero value otherwise.

### GetWabaIdOk

`func (o *WhatsappTemplate) GetWabaIdOk() (*string, bool)`

GetWabaIdOk returns a tuple with the WabaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWabaId

`func (o *WhatsappTemplate) SetWabaId(v string)`

SetWabaId sets WabaId field to given value.


### GetName

`func (o *WhatsappTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WhatsappTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WhatsappTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetLanguage

`func (o *WhatsappTemplate) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *WhatsappTemplate) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *WhatsappTemplate) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetCategory

`func (o *WhatsappTemplate) GetCategory() WhatsappTemplateCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *WhatsappTemplate) GetCategoryOk() (*WhatsappTemplateCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *WhatsappTemplate) SetCategory(v WhatsappTemplateCategory)`

SetCategory sets Category field to given value.


### GetComponents

`func (o *WhatsappTemplate) GetComponents() []WhatsappTemplateComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *WhatsappTemplate) GetComponentsOk() (*[]WhatsappTemplateComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *WhatsappTemplate) SetComponents(v []WhatsappTemplateComponent)`

SetComponents sets Components field to given value.


### GetStatus

`func (o *WhatsappTemplate) GetStatus() WhatsappTemplateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WhatsappTemplate) GetStatusOk() (*WhatsappTemplateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WhatsappTemplate) SetStatus(v WhatsappTemplateStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WhatsappTemplate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *WhatsappTemplate) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *WhatsappTemplate) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *WhatsappTemplate) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *WhatsappTemplate) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetCreateTime

`func (o *WhatsappTemplate) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *WhatsappTemplate) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *WhatsappTemplate) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *WhatsappTemplate) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *WhatsappTemplate) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *WhatsappTemplate) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *WhatsappTemplate) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *WhatsappTemplate) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


