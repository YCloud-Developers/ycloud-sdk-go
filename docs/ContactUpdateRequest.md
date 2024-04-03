# ContactUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nickname** | Pointer to **string** | Contact&#39;s nickname. Maximum length: 250 characters. | [optional] 
**PhoneNumber** | Pointer to **string** | Unique Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | [optional] 
**CountryCode** | Pointer to **string** | Two-letter country abbreviation. See [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). | [optional] 
**Email** | Pointer to **string** | The contact&#39;s email address. If present, the email address must be unique. | [optional] 
**Tags** | Pointer to **[]string** | Contact&#39;s tags. Maximum items: 50. | [optional] 
**CustomAttributes** | Pointer to [**[]ContactCustomAttribute**](ContactCustomAttribute.md) | Contact&#39;s custom attributes. If present (i.e., not &#x60;null&#x60;), all previous attributes of this contact will be replaced. | [optional] 
**OwnerEmail** | Pointer to **string** | The email address of the contact&#39;s owner. | [optional] 

## Methods

### NewContactUpdateRequest

`func NewContactUpdateRequest() *ContactUpdateRequest`

NewContactUpdateRequest instantiates a new ContactUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactUpdateRequestWithDefaults

`func NewContactUpdateRequestWithDefaults() *ContactUpdateRequest`

NewContactUpdateRequestWithDefaults instantiates a new ContactUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNickname

`func (o *ContactUpdateRequest) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *ContactUpdateRequest) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *ContactUpdateRequest) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *ContactUpdateRequest) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *ContactUpdateRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ContactUpdateRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ContactUpdateRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ContactUpdateRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetCountryCode

`func (o *ContactUpdateRequest) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ContactUpdateRequest) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ContactUpdateRequest) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *ContactUpdateRequest) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetEmail

`func (o *ContactUpdateRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContactUpdateRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContactUpdateRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ContactUpdateRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetTags

`func (o *ContactUpdateRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ContactUpdateRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ContactUpdateRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ContactUpdateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomAttributes

`func (o *ContactUpdateRequest) GetCustomAttributes() []ContactCustomAttribute`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *ContactUpdateRequest) GetCustomAttributesOk() (*[]ContactCustomAttribute, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *ContactUpdateRequest) SetCustomAttributes(v []ContactCustomAttribute)`

SetCustomAttributes sets CustomAttributes field to given value.

### HasCustomAttributes

`func (o *ContactUpdateRequest) HasCustomAttributes() bool`

HasCustomAttributes returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *ContactUpdateRequest) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *ContactUpdateRequest) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *ContactUpdateRequest) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *ContactUpdateRequest) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


