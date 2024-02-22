# ContactCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nickname** | Pointer to **string** | Contact&#39;s nickname. Maximum length: 250 characters. | [optional] 
**PhoneNumber** | **string** | Unique Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | 
**CountryCode** | Pointer to **string** | Two-letter country abbreviation. See [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). | [optional] 
**Email** | Pointer to **string** | Contact&#39;s email address. If present, the email address must be unique. | [optional] 
**Tags** | Pointer to **[]string** | Contact&#39;s tags. Max items: 50. Max characters per tag: 50. | [optional] 
**CustomAttributes** | Pointer to [**[]ContactCustomAttribute**](ContactCustomAttribute.md) | Contact&#39;s custom attributes. | [optional] 

## Methods

### NewContactCreateRequest

`func NewContactCreateRequest(phoneNumber string, ) *ContactCreateRequest`

NewContactCreateRequest instantiates a new ContactCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactCreateRequestWithDefaults

`func NewContactCreateRequestWithDefaults() *ContactCreateRequest`

NewContactCreateRequestWithDefaults instantiates a new ContactCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNickname

`func (o *ContactCreateRequest) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *ContactCreateRequest) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *ContactCreateRequest) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *ContactCreateRequest) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *ContactCreateRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ContactCreateRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ContactCreateRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetCountryCode

`func (o *ContactCreateRequest) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ContactCreateRequest) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ContactCreateRequest) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *ContactCreateRequest) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetEmail

`func (o *ContactCreateRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContactCreateRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContactCreateRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ContactCreateRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetTags

`func (o *ContactCreateRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ContactCreateRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ContactCreateRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ContactCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomAttributes

`func (o *ContactCreateRequest) GetCustomAttributes() []ContactCustomAttribute`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *ContactCreateRequest) GetCustomAttributesOk() (*[]ContactCustomAttribute, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *ContactCreateRequest) SetCustomAttributes(v []ContactCustomAttribute)`

SetCustomAttributes sets CustomAttributes field to given value.

### HasCustomAttributes

`func (o *ContactCreateRequest) HasCustomAttributes() bool`

HasCustomAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
