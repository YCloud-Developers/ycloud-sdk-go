# WhatsappTemplateAnalyticsDataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** |  |
**Sent** | **int64** | Number of messages created on this date for the selected template. |
**Delivered** | **int64** | Number of those messages that were delivered. |
**Failed** | **int64** | Number of those messages whose status is failed or expired. |
**Read** | **int64** | Number of those messages that were read. |
**Clicks** | **int64** | Total number of clicks recorded for this template on this date. Interpret zero together with analyticsStatus. |
**UniqueReplies** | **int64** | Number of unique replies recorded for this template on this date. This is mapped from Meta replied and is a Meta Template Insights metric. |
**ButtonClicks** | [**[]WhatsappTemplateAnalyticsButtonClick**](WhatsappTemplateAnalyticsButtonClick.md) | Button click breakdown for this template on this date. Empty when no details are available. |

## Methods

### NewWhatsappTemplateAnalyticsDataPoint

`func NewWhatsappTemplateAnalyticsDataPoint(date string, sent int64, delivered int64, failed int64, read int64, clicks int64, uniqueReplies int64, buttonClicks []WhatsappTemplateAnalyticsButtonClick, ) *WhatsappTemplateAnalyticsDataPoint`

NewWhatsappTemplateAnalyticsDataPoint instantiates a new WhatsappTemplateAnalyticsDataPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappTemplateAnalyticsDataPointWithDefaults

`func NewWhatsappTemplateAnalyticsDataPointWithDefaults() *WhatsappTemplateAnalyticsDataPoint`

NewWhatsappTemplateAnalyticsDataPointWithDefaults instantiates a new WhatsappTemplateAnalyticsDataPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *WhatsappTemplateAnalyticsDataPoint) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *WhatsappTemplateAnalyticsDataPoint) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *WhatsappTemplateAnalyticsDataPoint) SetDate(v string)`

SetDate sets Date field to given value.


### GetSent

`func (o *WhatsappTemplateAnalyticsDataPoint) GetSent() int64`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *WhatsappTemplateAnalyticsDataPoint) GetSentOk() (*int64, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *WhatsappTemplateAnalyticsDataPoint) SetSent(v int64)`

SetSent sets Sent field to given value.


### GetDelivered

`func (o *WhatsappTemplateAnalyticsDataPoint) GetDelivered() int64`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *WhatsappTemplateAnalyticsDataPoint) GetDeliveredOk() (*int64, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *WhatsappTemplateAnalyticsDataPoint) SetDelivered(v int64)`

SetDelivered sets Delivered field to given value.


### GetFailed

`func (o *WhatsappTemplateAnalyticsDataPoint) GetFailed() int64`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *WhatsappTemplateAnalyticsDataPoint) GetFailedOk() (*int64, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *WhatsappTemplateAnalyticsDataPoint) SetFailed(v int64)`

SetFailed sets Failed field to given value.


### GetRead

`func (o *WhatsappTemplateAnalyticsDataPoint) GetRead() int64`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *WhatsappTemplateAnalyticsDataPoint) GetReadOk() (*int64, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *WhatsappTemplateAnalyticsDataPoint) SetRead(v int64)`

SetRead sets Read field to given value.


### GetClicks

`func (o *WhatsappTemplateAnalyticsDataPoint) GetClicks() int64`

GetClicks returns the Clicks field if non-nil, zero value otherwise.

### GetClicksOk

`func (o *WhatsappTemplateAnalyticsDataPoint) GetClicksOk() (*int64, bool)`

GetClicksOk returns a tuple with the Clicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicks

`func (o *WhatsappTemplateAnalyticsDataPoint) SetClicks(v int64)`

SetClicks sets Clicks field to given value.


### GetUniqueReplies

`func (o *WhatsappTemplateAnalyticsDataPoint) GetUniqueReplies() int64`

GetUniqueReplies returns the UniqueReplies field if non-nil, zero value otherwise.

### GetUniqueRepliesOk

`func (o *WhatsappTemplateAnalyticsDataPoint) GetUniqueRepliesOk() (*int64, bool)`

GetUniqueRepliesOk returns a tuple with the UniqueReplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueReplies

`func (o *WhatsappTemplateAnalyticsDataPoint) SetUniqueReplies(v int64)`

SetUniqueReplies sets UniqueReplies field to given value.


### GetButtonClicks

`func (o *WhatsappTemplateAnalyticsDataPoint) GetButtonClicks() []WhatsappTemplateAnalyticsButtonClick`

GetButtonClicks returns the ButtonClicks field if non-nil, zero value otherwise.

### GetButtonClicksOk

`func (o *WhatsappTemplateAnalyticsDataPoint) GetButtonClicksOk() (*[]WhatsappTemplateAnalyticsButtonClick, bool)`

GetButtonClicksOk returns a tuple with the ButtonClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonClicks

`func (o *WhatsappTemplateAnalyticsDataPoint) SetButtonClicks(v []WhatsappTemplateAnalyticsButtonClick)`

SetButtonClicks sets ButtonClicks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
