# WhatsappTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OfficialTemplateId** | Pointer to **string** | Official template ID assigned by WhatsApp. This ID is used to identify the template in WhatsApp&#39;s system. | [optional] 
**WabaId** | **string** | WhatsApp Business Account ID. | 
**Name** | **string** | Name of the template. | 
**Language** | **string** | Language code of the template. See [Supported Languages](https://developers.facebook.com/docs/whatsapp/api/messages/message-templates#supported-languages) for all codes. | 
**Category** | Pointer to [**WhatsappTemplateCategory**](WhatsappTemplateCategory.md) |  | [optional] 
**SubCategory** | Pointer to [**WhatsappTemplateSubCategory**](WhatsappTemplateSubCategory.md) |  | [optional] 
**PreviousCategory** | Pointer to **string** | This field indicates the template&#39;s previous category (or &#x60;null&#x60;, for newly created templates after April 1, 2023). Compare this value to the template&#39;s &#x60;category&#x60; field value, which indicates the template&#39;s current category. | [optional] 
**MessageSendTtlSeconds** | Pointer to **int32** | If we are unable to deliver a message for an amount of time that exceeds its time-to-live, we will stop retrying and drop the message. By default, messages that use an authentication template have a default TTL of **10 minutes**, and messages that use a utility or marketing template have a default TTL of **30 days**. Set its value between &#x60;30&#x60; and &#x60;900&#x60; seconds (i.e., 30 seconds to 15 minutes) for authentication templates, or &#x60;30&#x60; and &#x60;43200&#x60; seconds (i.e., 30 seconds to 12 hours) for utility templates, or &#x60;43200&#x60; and &#x60;2592000&#x60; seconds (i.e., 12 hours to 30 days) for marketing templates. Alternatively, you can set this value to &#x60;-1&#x60;, which will set a custom TTL of 30 days for either type of template. We encourage you to set a time-to-live for all of your authentication templates, preferably equal to or less than your code expiration time, to ensure your customers only get a message when a code is still usable. Authentication templates created before October 23, 2024, have a default TTL of 30 days. | [optional] 
**Components** | Pointer to [**[]WhatsappTemplateComponent**](WhatsappTemplateComponent.md) | Template components. A template consists of &#x60;HEADER&#x60;, &#x60;BODY&#x60;, &#x60;FOOTER&#x60;, and &#x60;BUTTONS&#x60; components. &#x60;BODY&#x60; component is required, the other types are optional. | [optional] 
**Status** | Pointer to [**WhatsappTemplateStatus**](WhatsappTemplateStatus.md) |  | [optional] 
**QualityRating** | Pointer to [**WhatsappTemplateQualityRating**](WhatsappTemplateQualityRating.md) |  | [optional] 
**Reason** | Pointer to **string** | The reason why the template is rejected. | [optional] 
**CreateTime** | Pointer to **time.Time** | The time at which this object is created, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional] 
**UpdateTime** | Pointer to **time.Time** | The time at which this object is updated, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional] 
**StatusUpdateEvent** | Pointer to [**WhatsappTemplateStatusUpdateEventEnum**](WhatsappTemplateStatusUpdateEventEnum.md) |  | [optional] 
**DisableDate** | Pointer to **string** | The date at which the template will be disabled. When a WhatsApp template &#x60;FLAGGED&#x60; event is received, this field is set. | [optional] 
**WhatsappApiError** | Pointer to [**WhatsappApiError**](WhatsappApiError.md) |  | [optional] 

## Methods

### NewWhatsappTemplate

`func NewWhatsappTemplate(wabaId string, name string, language string, ) *WhatsappTemplate`

NewWhatsappTemplate instantiates a new WhatsappTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappTemplateWithDefaults

`func NewWhatsappTemplateWithDefaults() *WhatsappTemplate`

NewWhatsappTemplateWithDefaults instantiates a new WhatsappTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOfficialTemplateId

`func (o *WhatsappTemplate) GetOfficialTemplateId() string`

GetOfficialTemplateId returns the OfficialTemplateId field if non-nil, zero value otherwise.

### GetOfficialTemplateIdOk

`func (o *WhatsappTemplate) GetOfficialTemplateIdOk() (*string, bool)`

GetOfficialTemplateIdOk returns a tuple with the OfficialTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialTemplateId

`func (o *WhatsappTemplate) SetOfficialTemplateId(v string)`

SetOfficialTemplateId sets OfficialTemplateId field to given value.

### HasOfficialTemplateId

`func (o *WhatsappTemplate) HasOfficialTemplateId() bool`

HasOfficialTemplateId returns a boolean if a field has been set.

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

### HasCategory

`func (o *WhatsappTemplate) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetSubCategory

`func (o *WhatsappTemplate) GetSubCategory() WhatsappTemplateSubCategory`

GetSubCategory returns the SubCategory field if non-nil, zero value otherwise.

### GetSubCategoryOk

`func (o *WhatsappTemplate) GetSubCategoryOk() (*WhatsappTemplateSubCategory, bool)`

GetSubCategoryOk returns a tuple with the SubCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategory

`func (o *WhatsappTemplate) SetSubCategory(v WhatsappTemplateSubCategory)`

SetSubCategory sets SubCategory field to given value.

### HasSubCategory

`func (o *WhatsappTemplate) HasSubCategory() bool`

HasSubCategory returns a boolean if a field has been set.

### GetPreviousCategory

`func (o *WhatsappTemplate) GetPreviousCategory() string`

GetPreviousCategory returns the PreviousCategory field if non-nil, zero value otherwise.

### GetPreviousCategoryOk

`func (o *WhatsappTemplate) GetPreviousCategoryOk() (*string, bool)`

GetPreviousCategoryOk returns a tuple with the PreviousCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousCategory

`func (o *WhatsappTemplate) SetPreviousCategory(v string)`

SetPreviousCategory sets PreviousCategory field to given value.

### HasPreviousCategory

`func (o *WhatsappTemplate) HasPreviousCategory() bool`

HasPreviousCategory returns a boolean if a field has been set.

### GetMessageSendTtlSeconds

`func (o *WhatsappTemplate) GetMessageSendTtlSeconds() int32`

GetMessageSendTtlSeconds returns the MessageSendTtlSeconds field if non-nil, zero value otherwise.

### GetMessageSendTtlSecondsOk

`func (o *WhatsappTemplate) GetMessageSendTtlSecondsOk() (*int32, bool)`

GetMessageSendTtlSecondsOk returns a tuple with the MessageSendTtlSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSendTtlSeconds

`func (o *WhatsappTemplate) SetMessageSendTtlSeconds(v int32)`

SetMessageSendTtlSeconds sets MessageSendTtlSeconds field to given value.

### HasMessageSendTtlSeconds

`func (o *WhatsappTemplate) HasMessageSendTtlSeconds() bool`

HasMessageSendTtlSeconds returns a boolean if a field has been set.

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

### HasComponents

`func (o *WhatsappTemplate) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

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

### GetQualityRating

`func (o *WhatsappTemplate) GetQualityRating() WhatsappTemplateQualityRating`

GetQualityRating returns the QualityRating field if non-nil, zero value otherwise.

### GetQualityRatingOk

`func (o *WhatsappTemplate) GetQualityRatingOk() (*WhatsappTemplateQualityRating, bool)`

GetQualityRatingOk returns a tuple with the QualityRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityRating

`func (o *WhatsappTemplate) SetQualityRating(v WhatsappTemplateQualityRating)`

SetQualityRating sets QualityRating field to given value.

### HasQualityRating

`func (o *WhatsappTemplate) HasQualityRating() bool`

HasQualityRating returns a boolean if a field has been set.

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

### GetStatusUpdateEvent

`func (o *WhatsappTemplate) GetStatusUpdateEvent() WhatsappTemplateStatusUpdateEventEnum`

GetStatusUpdateEvent returns the StatusUpdateEvent field if non-nil, zero value otherwise.

### GetStatusUpdateEventOk

`func (o *WhatsappTemplate) GetStatusUpdateEventOk() (*WhatsappTemplateStatusUpdateEventEnum, bool)`

GetStatusUpdateEventOk returns a tuple with the StatusUpdateEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUpdateEvent

`func (o *WhatsappTemplate) SetStatusUpdateEvent(v WhatsappTemplateStatusUpdateEventEnum)`

SetStatusUpdateEvent sets StatusUpdateEvent field to given value.

### HasStatusUpdateEvent

`func (o *WhatsappTemplate) HasStatusUpdateEvent() bool`

HasStatusUpdateEvent returns a boolean if a field has been set.

### GetDisableDate

`func (o *WhatsappTemplate) GetDisableDate() string`

GetDisableDate returns the DisableDate field if non-nil, zero value otherwise.

### GetDisableDateOk

`func (o *WhatsappTemplate) GetDisableDateOk() (*string, bool)`

GetDisableDateOk returns a tuple with the DisableDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDate

`func (o *WhatsappTemplate) SetDisableDate(v string)`

SetDisableDate sets DisableDate field to given value.

### HasDisableDate

`func (o *WhatsappTemplate) HasDisableDate() bool`

HasDisableDate returns a boolean if a field has been set.

### GetWhatsappApiError

`func (o *WhatsappTemplate) GetWhatsappApiError() WhatsappApiError`

GetWhatsappApiError returns the WhatsappApiError field if non-nil, zero value otherwise.

### GetWhatsappApiErrorOk

`func (o *WhatsappTemplate) GetWhatsappApiErrorOk() (*WhatsappApiError, bool)`

GetWhatsappApiErrorOk returns a tuple with the WhatsappApiError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhatsappApiError

`func (o *WhatsappTemplate) SetWhatsappApiError(v WhatsappApiError)`

SetWhatsappApiError sets WhatsappApiError field to given value.

### HasWhatsappApiError

`func (o *WhatsappTemplate) HasWhatsappApiError() bool`

HasWhatsappApiError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


