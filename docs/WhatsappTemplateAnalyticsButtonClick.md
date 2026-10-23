# WhatsappTemplateAnalyticsButtonClick

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of button associated with the clicks. Values may include quick_reply_button, unique_url_button, or url_button. |
**ButtonContent** | **string** | Button content associated with the clicks. |
**Count** | **int64** | Number of clicks for this button entry. |

## Methods

### NewWhatsappTemplateAnalyticsButtonClick

`func NewWhatsappTemplateAnalyticsButtonClick(type_ string, buttonContent string, count int64, ) *WhatsappTemplateAnalyticsButtonClick`

NewWhatsappTemplateAnalyticsButtonClick instantiates a new WhatsappTemplateAnalyticsButtonClick object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappTemplateAnalyticsButtonClickWithDefaults

`func NewWhatsappTemplateAnalyticsButtonClickWithDefaults() *WhatsappTemplateAnalyticsButtonClick`

NewWhatsappTemplateAnalyticsButtonClickWithDefaults instantiates a new WhatsappTemplateAnalyticsButtonClick object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WhatsappTemplateAnalyticsButtonClick) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappTemplateAnalyticsButtonClick) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappTemplateAnalyticsButtonClick) SetType(v string)`

SetType sets Type field to given value.


### GetButtonContent

`func (o *WhatsappTemplateAnalyticsButtonClick) GetButtonContent() string`

GetButtonContent returns the ButtonContent field if non-nil, zero value otherwise.

### GetButtonContentOk

`func (o *WhatsappTemplateAnalyticsButtonClick) GetButtonContentOk() (*string, bool)`

GetButtonContentOk returns a tuple with the ButtonContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonContent

`func (o *WhatsappTemplateAnalyticsButtonClick) SetButtonContent(v string)`

SetButtonContent sets ButtonContent field to given value.


### GetCount

`func (o *WhatsappTemplateAnalyticsButtonClick) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *WhatsappTemplateAnalyticsButtonClick) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *WhatsappTemplateAnalyticsButtonClick) SetCount(v int64)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
