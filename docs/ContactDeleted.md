# ContactDeleted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Contact ID | 
**NickName** | Pointer to **string** | Contact&#39;s nickname. | [optional] 
**PhoneNumber** | Pointer to **string** | Unique Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | [optional] 
**UpdateTime** | Pointer to **time.Time** | The time at which the contact was last updated, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional] 

## Methods

### NewContactDeleted

`func NewContactDeleted(id string, ) *ContactDeleted`

NewContactDeleted instantiates a new ContactDeleted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactDeletedWithDefaults

`func NewContactDeletedWithDefaults() *ContactDeleted`

NewContactDeletedWithDefaults instantiates a new ContactDeleted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactDeleted) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactDeleted) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactDeleted) SetId(v string)`

SetId sets Id field to given value.


### GetNickName

`func (o *ContactDeleted) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *ContactDeleted) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *ContactDeleted) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *ContactDeleted) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *ContactDeleted) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ContactDeleted) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ContactDeleted) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ContactDeleted) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetUpdateTime

`func (o *ContactDeleted) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *ContactDeleted) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *ContactDeleted) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *ContactDeleted) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


