# WhatsappBusinessAccountRestrictionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestrictionType** | Pointer to **string** | Restriction type. | [optional] 
**Expiration** | Pointer to **time.Time** | The time at which this restriction expires, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional] 

## Methods

### NewWhatsappBusinessAccountRestrictionInfo

`func NewWhatsappBusinessAccountRestrictionInfo() *WhatsappBusinessAccountRestrictionInfo`

NewWhatsappBusinessAccountRestrictionInfo instantiates a new WhatsappBusinessAccountRestrictionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappBusinessAccountRestrictionInfoWithDefaults

`func NewWhatsappBusinessAccountRestrictionInfoWithDefaults() *WhatsappBusinessAccountRestrictionInfo`

NewWhatsappBusinessAccountRestrictionInfoWithDefaults instantiates a new WhatsappBusinessAccountRestrictionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestrictionType

`func (o *WhatsappBusinessAccountRestrictionInfo) GetRestrictionType() string`

GetRestrictionType returns the RestrictionType field if non-nil, zero value otherwise.

### GetRestrictionTypeOk

`func (o *WhatsappBusinessAccountRestrictionInfo) GetRestrictionTypeOk() (*string, bool)`

GetRestrictionTypeOk returns a tuple with the RestrictionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionType

`func (o *WhatsappBusinessAccountRestrictionInfo) SetRestrictionType(v string)`

SetRestrictionType sets RestrictionType field to given value.

### HasRestrictionType

`func (o *WhatsappBusinessAccountRestrictionInfo) HasRestrictionType() bool`

HasRestrictionType returns a boolean if a field has been set.

### GetExpiration

`func (o *WhatsappBusinessAccountRestrictionInfo) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *WhatsappBusinessAccountRestrictionInfo) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *WhatsappBusinessAccountRestrictionInfo) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *WhatsappBusinessAccountRestrictionInfo) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


