# PublicCreateAllowlistEntryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerPhoneNumber** | Pointer to **NullableString** | Consumer phone number in E.164 format. | [optional]
**ConsumerBsuid** | Pointer to **NullableString** | Ordinary WhatsApp consumer BSUID. Provide exactly one non-null identity. | [optional]

## Methods

### NewPublicCreateAllowlistEntryRequest

`func NewPublicCreateAllowlistEntryRequest() *PublicCreateAllowlistEntryRequest`

NewPublicCreateAllowlistEntryRequest instantiates a new PublicCreateAllowlistEntryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicCreateAllowlistEntryRequestWithDefaults

`func NewPublicCreateAllowlistEntryRequestWithDefaults() *PublicCreateAllowlistEntryRequest`

NewPublicCreateAllowlistEntryRequestWithDefaults instantiates a new PublicCreateAllowlistEntryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumerPhoneNumber

`func (o *PublicCreateAllowlistEntryRequest) GetConsumerPhoneNumber() string`

GetConsumerPhoneNumber returns the ConsumerPhoneNumber field if non-nil, zero value otherwise.

### GetConsumerPhoneNumberOk

`func (o *PublicCreateAllowlistEntryRequest) GetConsumerPhoneNumberOk() (*string, bool)`

GetConsumerPhoneNumberOk returns a tuple with the ConsumerPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerPhoneNumber

`func (o *PublicCreateAllowlistEntryRequest) SetConsumerPhoneNumber(v string)`

SetConsumerPhoneNumber sets ConsumerPhoneNumber field to given value.

### HasConsumerPhoneNumber

`func (o *PublicCreateAllowlistEntryRequest) HasConsumerPhoneNumber() bool`

HasConsumerPhoneNumber returns a boolean if a field has been set.

### SetConsumerPhoneNumberNil

`func (o *PublicCreateAllowlistEntryRequest) SetConsumerPhoneNumberNil(b bool)`

 SetConsumerPhoneNumberNil sets the value for ConsumerPhoneNumber to be an explicit nil

### UnsetConsumerPhoneNumber
`func (o *PublicCreateAllowlistEntryRequest) UnsetConsumerPhoneNumber()`

UnsetConsumerPhoneNumber ensures that no value is present for ConsumerPhoneNumber, not even an explicit nil
### GetConsumerBsuid

`func (o *PublicCreateAllowlistEntryRequest) GetConsumerBsuid() string`

GetConsumerBsuid returns the ConsumerBsuid field if non-nil, zero value otherwise.

### GetConsumerBsuidOk

`func (o *PublicCreateAllowlistEntryRequest) GetConsumerBsuidOk() (*string, bool)`

GetConsumerBsuidOk returns a tuple with the ConsumerBsuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerBsuid

`func (o *PublicCreateAllowlistEntryRequest) SetConsumerBsuid(v string)`

SetConsumerBsuid sets ConsumerBsuid field to given value.

### HasConsumerBsuid

`func (o *PublicCreateAllowlistEntryRequest) HasConsumerBsuid() bool`

HasConsumerBsuid returns a boolean if a field has been set.

### SetConsumerBsuidNil

`func (o *PublicCreateAllowlistEntryRequest) SetConsumerBsuidNil(b bool)`

 SetConsumerBsuidNil sets the value for ConsumerBsuid to be an explicit nil

### UnsetConsumerBsuid
`func (o *PublicCreateAllowlistEntryRequest) UnsetConsumerBsuid()`

UnsetConsumerBsuid ensures that no value is present for ConsumerBsuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
