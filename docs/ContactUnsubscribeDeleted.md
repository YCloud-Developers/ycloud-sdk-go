# ContactUnsubscribeDeleted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | Pointer to **string** | Unique Customer Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | [optional] 
**Source** | Pointer to **string** | The source from which a customer resumed their subscription - &#x60;Whatsapp&#x60;: The customer resumed their subscription on the whatsapp client - &#x60;API&#x60;: You remove the customer from the unsubscribe list through the OpenAPI of YCloud - &#x60;Manual&#x60;: You remove the customer from the unsubscribe list on the Contact page of YCloud. | [optional] 
**UpdateTime** | Pointer to **time.Time** | The time when customers cancel unsubscribe, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional] 

## Methods

### NewContactUnsubscribeDeleted

`func NewContactUnsubscribeDeleted() *ContactUnsubscribeDeleted`

NewContactUnsubscribeDeleted instantiates a new ContactUnsubscribeDeleted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactUnsubscribeDeletedWithDefaults

`func NewContactUnsubscribeDeletedWithDefaults() *ContactUnsubscribeDeleted`

NewContactUnsubscribeDeletedWithDefaults instantiates a new ContactUnsubscribeDeleted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *ContactUnsubscribeDeleted) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ContactUnsubscribeDeleted) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ContactUnsubscribeDeleted) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ContactUnsubscribeDeleted) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetSource

`func (o *ContactUnsubscribeDeleted) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ContactUnsubscribeDeleted) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ContactUnsubscribeDeleted) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ContactUnsubscribeDeleted) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUpdateTime

`func (o *ContactUnsubscribeDeleted) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *ContactUnsubscribeDeleted) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *ContactUnsubscribeDeleted) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *ContactUnsubscribeDeleted) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


