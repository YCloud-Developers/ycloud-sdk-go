# WhatsappTemplateAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WabaId** | **string** | WhatsApp Business Account ID. |
**OfficialTemplateId** | **NullableString** | Resolved official WhatsApp/Meta template ID. Successful queries resolve a current template; the property remains nullable for contract compatibility. |
**TemplateName** | **string** | Template name used for the statistics query. |
**Language** | **string** | Template language used for the statistics query. |
**Timezone** | **string** | IANA timezone resolved from the WABA configuration and used to interpret the date range. |
**AnalyticsStatus** | **string** | Meta Template Insights status only; it does not describe YCloud message metrics. NO_DATA means the Meta query completed without data for the requested date range. ERROR indicates that Template Insights are unavailable because of an error after the current template was successfully resolved and validated. |
**StartDate** | **string** | Inclusive start date from the request. |
**EndDate** | **string** | Inclusive end date from the request. |
**DataPoints** | [**[]WhatsappTemplateAnalyticsDataPoint**](WhatsappTemplateAnalyticsDataPoint.md) | One item for every date in the requested range, ordered by date ascending. Missing metrics are returned as zero and missing button details as an empty array. |

## Methods

### NewWhatsappTemplateAnalytics

`func NewWhatsappTemplateAnalytics(wabaId string, officialTemplateId NullableString, templateName string, language string, timezone string, analyticsStatus string, startDate string, endDate string, dataPoints []WhatsappTemplateAnalyticsDataPoint, ) *WhatsappTemplateAnalytics`

NewWhatsappTemplateAnalytics instantiates a new WhatsappTemplateAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappTemplateAnalyticsWithDefaults

`func NewWhatsappTemplateAnalyticsWithDefaults() *WhatsappTemplateAnalytics`

NewWhatsappTemplateAnalyticsWithDefaults instantiates a new WhatsappTemplateAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWabaId

`func (o *WhatsappTemplateAnalytics) GetWabaId() string`

GetWabaId returns the WabaId field if non-nil, zero value otherwise.

### GetWabaIdOk

`func (o *WhatsappTemplateAnalytics) GetWabaIdOk() (*string, bool)`

GetWabaIdOk returns a tuple with the WabaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWabaId

`func (o *WhatsappTemplateAnalytics) SetWabaId(v string)`

SetWabaId sets WabaId field to given value.


### GetOfficialTemplateId

`func (o *WhatsappTemplateAnalytics) GetOfficialTemplateId() string`

GetOfficialTemplateId returns the OfficialTemplateId field if non-nil, zero value otherwise.

### GetOfficialTemplateIdOk

`func (o *WhatsappTemplateAnalytics) GetOfficialTemplateIdOk() (*string, bool)`

GetOfficialTemplateIdOk returns a tuple with the OfficialTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialTemplateId

`func (o *WhatsappTemplateAnalytics) SetOfficialTemplateId(v string)`

SetOfficialTemplateId sets OfficialTemplateId field to given value.


### SetOfficialTemplateIdNil

`func (o *WhatsappTemplateAnalytics) SetOfficialTemplateIdNil(b bool)`

 SetOfficialTemplateIdNil sets the value for OfficialTemplateId to be an explicit nil

### UnsetOfficialTemplateId
`func (o *WhatsappTemplateAnalytics) UnsetOfficialTemplateId()`

UnsetOfficialTemplateId ensures that no value is present for OfficialTemplateId, not even an explicit nil
### GetTemplateName

`func (o *WhatsappTemplateAnalytics) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *WhatsappTemplateAnalytics) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *WhatsappTemplateAnalytics) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.


### GetLanguage

`func (o *WhatsappTemplateAnalytics) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *WhatsappTemplateAnalytics) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *WhatsappTemplateAnalytics) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetTimezone

`func (o *WhatsappTemplateAnalytics) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *WhatsappTemplateAnalytics) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *WhatsappTemplateAnalytics) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetAnalyticsStatus

`func (o *WhatsappTemplateAnalytics) GetAnalyticsStatus() string`

GetAnalyticsStatus returns the AnalyticsStatus field if non-nil, zero value otherwise.

### GetAnalyticsStatusOk

`func (o *WhatsappTemplateAnalytics) GetAnalyticsStatusOk() (*string, bool)`

GetAnalyticsStatusOk returns a tuple with the AnalyticsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsStatus

`func (o *WhatsappTemplateAnalytics) SetAnalyticsStatus(v string)`

SetAnalyticsStatus sets AnalyticsStatus field to given value.


### GetStartDate

`func (o *WhatsappTemplateAnalytics) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *WhatsappTemplateAnalytics) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *WhatsappTemplateAnalytics) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *WhatsappTemplateAnalytics) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *WhatsappTemplateAnalytics) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *WhatsappTemplateAnalytics) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetDataPoints

`func (o *WhatsappTemplateAnalytics) GetDataPoints() []WhatsappTemplateAnalyticsDataPoint`

GetDataPoints returns the DataPoints field if non-nil, zero value otherwise.

### GetDataPointsOk

`func (o *WhatsappTemplateAnalytics) GetDataPointsOk() (*[]WhatsappTemplateAnalyticsDataPoint, bool)`

GetDataPointsOk returns a tuple with the DataPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPoints

`func (o *WhatsappTemplateAnalytics) SetDataPoints(v []WhatsappTemplateAnalyticsDataPoint)`

SetDataPoints sets DataPoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
