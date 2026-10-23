# WhatsappTemplateAnalyticsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WabaId** | **string** |  |
**OfficialTemplateId** | Pointer to **NullableString** | Official WhatsApp/Meta template ID exposed by the existing template REST APIs. | [optional]
**TemplateName** | Pointer to **NullableString** | Exact template name. Must be provided together with language when officialTemplateId is absent. | [optional]
**Language** | Pointer to **NullableString** | Template language code. Must be provided together with templateName when officialTemplateId is absent. | [optional]
**StartDate** | **string** | Inclusive start date interpreted in the WABA timezone. |
**EndDate** | **string** | Inclusive end date interpreted in the WABA timezone. The inclusive range must not exceed 90 calendar days. |

## Methods

### NewWhatsappTemplateAnalyticsRequest

`func NewWhatsappTemplateAnalyticsRequest(wabaId string, startDate string, endDate string, ) *WhatsappTemplateAnalyticsRequest`

NewWhatsappTemplateAnalyticsRequest instantiates a new WhatsappTemplateAnalyticsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappTemplateAnalyticsRequestWithDefaults

`func NewWhatsappTemplateAnalyticsRequestWithDefaults() *WhatsappTemplateAnalyticsRequest`

NewWhatsappTemplateAnalyticsRequestWithDefaults instantiates a new WhatsappTemplateAnalyticsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWabaId

`func (o *WhatsappTemplateAnalyticsRequest) GetWabaId() string`

GetWabaId returns the WabaId field if non-nil, zero value otherwise.

### GetWabaIdOk

`func (o *WhatsappTemplateAnalyticsRequest) GetWabaIdOk() (*string, bool)`

GetWabaIdOk returns a tuple with the WabaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWabaId

`func (o *WhatsappTemplateAnalyticsRequest) SetWabaId(v string)`

SetWabaId sets WabaId field to given value.


### GetOfficialTemplateId

`func (o *WhatsappTemplateAnalyticsRequest) GetOfficialTemplateId() string`

GetOfficialTemplateId returns the OfficialTemplateId field if non-nil, zero value otherwise.

### GetOfficialTemplateIdOk

`func (o *WhatsappTemplateAnalyticsRequest) GetOfficialTemplateIdOk() (*string, bool)`

GetOfficialTemplateIdOk returns a tuple with the OfficialTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialTemplateId

`func (o *WhatsappTemplateAnalyticsRequest) SetOfficialTemplateId(v string)`

SetOfficialTemplateId sets OfficialTemplateId field to given value.

### HasOfficialTemplateId

`func (o *WhatsappTemplateAnalyticsRequest) HasOfficialTemplateId() bool`

HasOfficialTemplateId returns a boolean if a field has been set.

### SetOfficialTemplateIdNil

`func (o *WhatsappTemplateAnalyticsRequest) SetOfficialTemplateIdNil(b bool)`

 SetOfficialTemplateIdNil sets the value for OfficialTemplateId to be an explicit nil

### UnsetOfficialTemplateId
`func (o *WhatsappTemplateAnalyticsRequest) UnsetOfficialTemplateId()`

UnsetOfficialTemplateId ensures that no value is present for OfficialTemplateId, not even an explicit nil
### GetTemplateName

`func (o *WhatsappTemplateAnalyticsRequest) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *WhatsappTemplateAnalyticsRequest) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *WhatsappTemplateAnalyticsRequest) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *WhatsappTemplateAnalyticsRequest) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### SetTemplateNameNil

`func (o *WhatsappTemplateAnalyticsRequest) SetTemplateNameNil(b bool)`

 SetTemplateNameNil sets the value for TemplateName to be an explicit nil

### UnsetTemplateName
`func (o *WhatsappTemplateAnalyticsRequest) UnsetTemplateName()`

UnsetTemplateName ensures that no value is present for TemplateName, not even an explicit nil
### GetLanguage

`func (o *WhatsappTemplateAnalyticsRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *WhatsappTemplateAnalyticsRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *WhatsappTemplateAnalyticsRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *WhatsappTemplateAnalyticsRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *WhatsappTemplateAnalyticsRequest) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *WhatsappTemplateAnalyticsRequest) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetStartDate

`func (o *WhatsappTemplateAnalyticsRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *WhatsappTemplateAnalyticsRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *WhatsappTemplateAnalyticsRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *WhatsappTemplateAnalyticsRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *WhatsappTemplateAnalyticsRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *WhatsappTemplateAnalyticsRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
