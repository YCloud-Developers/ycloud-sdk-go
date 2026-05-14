# WhatsappUserPreference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WabaId** | Pointer to **string** | WhatsApp Business Account ID. | [optional]
**BusinessPhoneNumber** | Pointer to **string** | Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | [optional]
**BusinessPhoneId** | Pointer to **string** | Phone number ID. | [optional]
**ContactName** | Pointer to **string** | WhatsApp user name. | [optional]
**ContactPhoneNumber** | Pointer to **string** | WhatsApp user phone number. Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | [optional]
**Detail** | Pointer to **string** | Description of marketing message preference. | [optional]
**Category** | Pointer to **string** |  | [optional]
**Value** | Pointer to **string** | Marketing message preference. | [optional]
**Timestamp** | Pointer to **string** | Unix timestamp indicating when the webhook was triggered. | [optional]

## Methods

### NewWhatsappUserPreference

`func NewWhatsappUserPreference() *WhatsappUserPreference`

NewWhatsappUserPreference instantiates a new WhatsappUserPreference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappUserPreferenceWithDefaults

`func NewWhatsappUserPreferenceWithDefaults() *WhatsappUserPreference`

NewWhatsappUserPreferenceWithDefaults instantiates a new WhatsappUserPreference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWabaId

`func (o *WhatsappUserPreference) GetWabaId() string`

GetWabaId returns the WabaId field if non-nil, zero value otherwise.

### GetWabaIdOk

`func (o *WhatsappUserPreference) GetWabaIdOk() (*string, bool)`

GetWabaIdOk returns a tuple with the WabaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWabaId

`func (o *WhatsappUserPreference) SetWabaId(v string)`

SetWabaId sets WabaId field to given value.

### HasWabaId

`func (o *WhatsappUserPreference) HasWabaId() bool`

HasWabaId returns a boolean if a field has been set.

### GetBusinessPhoneNumber

`func (o *WhatsappUserPreference) GetBusinessPhoneNumber() string`

GetBusinessPhoneNumber returns the BusinessPhoneNumber field if non-nil, zero value otherwise.

### GetBusinessPhoneNumberOk

`func (o *WhatsappUserPreference) GetBusinessPhoneNumberOk() (*string, bool)`

GetBusinessPhoneNumberOk returns a tuple with the BusinessPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPhoneNumber

`func (o *WhatsappUserPreference) SetBusinessPhoneNumber(v string)`

SetBusinessPhoneNumber sets BusinessPhoneNumber field to given value.

### HasBusinessPhoneNumber

`func (o *WhatsappUserPreference) HasBusinessPhoneNumber() bool`

HasBusinessPhoneNumber returns a boolean if a field has been set.

### GetBusinessPhoneId

`func (o *WhatsappUserPreference) GetBusinessPhoneId() string`

GetBusinessPhoneId returns the BusinessPhoneId field if non-nil, zero value otherwise.

### GetBusinessPhoneIdOk

`func (o *WhatsappUserPreference) GetBusinessPhoneIdOk() (*string, bool)`

GetBusinessPhoneIdOk returns a tuple with the BusinessPhoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPhoneId

`func (o *WhatsappUserPreference) SetBusinessPhoneId(v string)`

SetBusinessPhoneId sets BusinessPhoneId field to given value.

### HasBusinessPhoneId

`func (o *WhatsappUserPreference) HasBusinessPhoneId() bool`

HasBusinessPhoneId returns a boolean if a field has been set.

### GetContactName

`func (o *WhatsappUserPreference) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *WhatsappUserPreference) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *WhatsappUserPreference) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *WhatsappUserPreference) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### GetContactPhoneNumber

`func (o *WhatsappUserPreference) GetContactPhoneNumber() string`

GetContactPhoneNumber returns the ContactPhoneNumber field if non-nil, zero value otherwise.

### GetContactPhoneNumberOk

`func (o *WhatsappUserPreference) GetContactPhoneNumberOk() (*string, bool)`

GetContactPhoneNumberOk returns a tuple with the ContactPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPhoneNumber

`func (o *WhatsappUserPreference) SetContactPhoneNumber(v string)`

SetContactPhoneNumber sets ContactPhoneNumber field to given value.

### HasContactPhoneNumber

`func (o *WhatsappUserPreference) HasContactPhoneNumber() bool`

HasContactPhoneNumber returns a boolean if a field has been set.

### GetDetail

`func (o *WhatsappUserPreference) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *WhatsappUserPreference) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *WhatsappUserPreference) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *WhatsappUserPreference) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetCategory

`func (o *WhatsappUserPreference) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *WhatsappUserPreference) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *WhatsappUserPreference) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *WhatsappUserPreference) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetValue

`func (o *WhatsappUserPreference) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WhatsappUserPreference) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WhatsappUserPreference) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *WhatsappUserPreference) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetTimestamp

`func (o *WhatsappUserPreference) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WhatsappUserPreference) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WhatsappUserPreference) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WhatsappUserPreference) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


