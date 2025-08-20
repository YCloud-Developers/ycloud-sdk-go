# ContactCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID for the object. | 
**NickName** | Pointer to **string** | Contact&#39;s nickname. | [optional] 
**RealName** | Pointer to **string** | Contact&#39;s real name. | [optional] 
**PhoneNumber** | Pointer to **string** | Unique Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | [optional] 
**CountryCode** | Pointer to **string** | Two-letter country abbreviation. See [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). | [optional] 
**CountryName** | Pointer to **string** | Full country name. | [optional] 
**Email** | Pointer to **string** | The contact&#39;s email address. If present, the email address must be unique. | [optional] 
**SourceType** | Pointer to **string** | The source type where the contact was created. | [optional] 
**SourceId** | Pointer to **string** | The source ID where the contact was created. | [optional] 
**SourceUrl** | Pointer to **string** | The source URL where the contact was created. | [optional] 
**LastSeen** | Pointer to **time.Time** | The time at which the contact last sent a message to your business, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional] 
**LastConnectedNumber** | Pointer to **string** | The business phone number that the contact last connected to. | [optional] 
**OwnerEmail** | Pointer to **string** | The email address of the contact&#39;s owner. | [optional] 
**Tags** | Pointer to **[]string** | Contact&#39;s tags. | [optional] 
**CreateTime** | Pointer to **time.Time** | The time at which the contact was created, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional] 
**UpdateTime** | Pointer to **time.Time** | The time at which the contact was last updated, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional] 
**Blocked** | Pointer to **bool** | Whether the contact is blocked. | [optional] 
**CustomAttributes** | Pointer to **map[string]map[string]interface{}** | Contact&#39;s custom attributes as key-value pairs. | [optional] 

## Methods

### NewContactCreated

`func NewContactCreated(id string, ) *ContactCreated`

NewContactCreated instantiates a new ContactCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactCreatedWithDefaults

`func NewContactCreatedWithDefaults() *ContactCreated`

NewContactCreatedWithDefaults instantiates a new ContactCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactCreated) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactCreated) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactCreated) SetId(v string)`

SetId sets Id field to given value.


### GetNickName

`func (o *ContactCreated) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *ContactCreated) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *ContactCreated) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *ContactCreated) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### GetRealName

`func (o *ContactCreated) GetRealName() string`

GetRealName returns the RealName field if non-nil, zero value otherwise.

### GetRealNameOk

`func (o *ContactCreated) GetRealNameOk() (*string, bool)`

GetRealNameOk returns a tuple with the RealName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealName

`func (o *ContactCreated) SetRealName(v string)`

SetRealName sets RealName field to given value.

### HasRealName

`func (o *ContactCreated) HasRealName() bool`

HasRealName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *ContactCreated) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ContactCreated) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ContactCreated) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ContactCreated) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetCountryCode

`func (o *ContactCreated) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ContactCreated) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ContactCreated) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *ContactCreated) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCountryName

`func (o *ContactCreated) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *ContactCreated) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *ContactCreated) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *ContactCreated) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### GetEmail

`func (o *ContactCreated) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContactCreated) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContactCreated) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ContactCreated) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSourceType

`func (o *ContactCreated) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ContactCreated) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ContactCreated) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *ContactCreated) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetSourceId

`func (o *ContactCreated) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ContactCreated) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ContactCreated) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ContactCreated) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceUrl

`func (o *ContactCreated) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *ContactCreated) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *ContactCreated) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *ContactCreated) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### GetLastSeen

`func (o *ContactCreated) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ContactCreated) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ContactCreated) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ContactCreated) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLastConnectedNumber

`func (o *ContactCreated) GetLastConnectedNumber() string`

GetLastConnectedNumber returns the LastConnectedNumber field if non-nil, zero value otherwise.

### GetLastConnectedNumberOk

`func (o *ContactCreated) GetLastConnectedNumberOk() (*string, bool)`

GetLastConnectedNumberOk returns a tuple with the LastConnectedNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectedNumber

`func (o *ContactCreated) SetLastConnectedNumber(v string)`

SetLastConnectedNumber sets LastConnectedNumber field to given value.

### HasLastConnectedNumber

`func (o *ContactCreated) HasLastConnectedNumber() bool`

HasLastConnectedNumber returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *ContactCreated) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *ContactCreated) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *ContactCreated) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *ContactCreated) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.

### GetTags

`func (o *ContactCreated) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ContactCreated) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ContactCreated) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ContactCreated) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreateTime

`func (o *ContactCreated) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ContactCreated) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ContactCreated) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ContactCreated) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *ContactCreated) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *ContactCreated) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *ContactCreated) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *ContactCreated) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetBlocked

`func (o *ContactCreated) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *ContactCreated) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *ContactCreated) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *ContactCreated) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetCustomAttributes

`func (o *ContactCreated) GetCustomAttributes() map[string]map[string]interface{}`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *ContactCreated) GetCustomAttributesOk() (*map[string]map[string]interface{}, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *ContactCreated) SetCustomAttributes(v map[string]map[string]interface{})`

SetCustomAttributes sets CustomAttributes field to given value.

### HasCustomAttributes

`func (o *ContactCreated) HasCustomAttributes() bool`

HasCustomAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


