# Contact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID for the object. | 
**Nickname** | Pointer to **string** | Contact&#39;s nickname. | [optional] 
**CountryCode** | Pointer to **string** | Two-letter country abbreviation. See [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). | [optional] 
**CountryName** | Pointer to **string** | Full country name. | [optional] 
**PhoneNumber** | Pointer to **string** | Unique Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | [optional] 
**Email** | Pointer to **string** | The contact&#39;s email address. If present, the email address must be unique. | [optional] 
**LastSeen** | Pointer to **time.Time** | The time at which the contact last sent a message to your business, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional] 
**LastMessageToPhoneNumber** | Pointer to **string** | The business phone number that the contact last sent a message to. | [optional] 
**Tags** | Pointer to **[]string** | Contact&#39;s tags. | [optional] 
**CreateTime** | Pointer to **time.Time** | The time at which the contact was created, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional] 
**CustomAttributes** | Pointer to [**[]ContactCustomAttribute**](ContactCustomAttribute.md) | Contact&#39;s custom attributes. | [optional] 
**OwnerEmail** | Pointer to **string** | The email address of the contact&#39;s owner. | [optional] 
**SourceType** | Pointer to [**ContactSourceType**](ContactSourceType.md) |  | [optional] 
**SourceId** | Pointer to **string** | Source identifier. A unique identifier related to the contact creation source. | [optional] 
**SourceUrl** | Pointer to **string** | Source URL. The source link address where the contact was created. | [optional] 

## Methods

### NewContact

`func NewContact(id string, ) *Contact`

NewContact instantiates a new Contact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactWithDefaults

`func NewContactWithDefaults() *Contact`

NewContactWithDefaults instantiates a new Contact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Contact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Contact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Contact) SetId(v string)`

SetId sets Id field to given value.


### GetNickname

`func (o *Contact) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *Contact) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *Contact) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *Contact) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetCountryCode

`func (o *Contact) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Contact) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Contact) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Contact) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCountryName

`func (o *Contact) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *Contact) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *Contact) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *Contact) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *Contact) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *Contact) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *Contact) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *Contact) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetEmail

`func (o *Contact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Contact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Contact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Contact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLastSeen

`func (o *Contact) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *Contact) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *Contact) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *Contact) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLastMessageToPhoneNumber

`func (o *Contact) GetLastMessageToPhoneNumber() string`

GetLastMessageToPhoneNumber returns the LastMessageToPhoneNumber field if non-nil, zero value otherwise.

### GetLastMessageToPhoneNumberOk

`func (o *Contact) GetLastMessageToPhoneNumberOk() (*string, bool)`

GetLastMessageToPhoneNumberOk returns a tuple with the LastMessageToPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageToPhoneNumber

`func (o *Contact) SetLastMessageToPhoneNumber(v string)`

SetLastMessageToPhoneNumber sets LastMessageToPhoneNumber field to given value.

### HasLastMessageToPhoneNumber

`func (o *Contact) HasLastMessageToPhoneNumber() bool`

HasLastMessageToPhoneNumber returns a boolean if a field has been set.

### GetTags

`func (o *Contact) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Contact) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Contact) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Contact) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCreateTime

`func (o *Contact) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *Contact) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *Contact) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *Contact) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCustomAttributes

`func (o *Contact) GetCustomAttributes() []ContactCustomAttribute`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *Contact) GetCustomAttributesOk() (*[]ContactCustomAttribute, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *Contact) SetCustomAttributes(v []ContactCustomAttribute)`

SetCustomAttributes sets CustomAttributes field to given value.

### HasCustomAttributes

`func (o *Contact) HasCustomAttributes() bool`

HasCustomAttributes returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *Contact) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *Contact) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *Contact) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *Contact) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.

### GetSourceType

`func (o *Contact) GetSourceType() ContactSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *Contact) GetSourceTypeOk() (*ContactSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *Contact) SetSourceType(v ContactSourceType)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *Contact) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetSourceId

`func (o *Contact) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Contact) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Contact) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *Contact) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceUrl

`func (o *Contact) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *Contact) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *Contact) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *Contact) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


