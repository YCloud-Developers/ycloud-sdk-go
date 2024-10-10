# WhatsappMessageOrderExpiration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **string** | A string of UTC timestamp in seconds of time when order should expire. Minimum threshold is 300 seconds. | 
**Description** | Pointer to **string** | Text explanation for expiration. | [optional] 

## Methods

### NewWhatsappMessageOrderExpiration

`func NewWhatsappMessageOrderExpiration(timestamp string, ) *WhatsappMessageOrderExpiration`

NewWhatsappMessageOrderExpiration instantiates a new WhatsappMessageOrderExpiration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageOrderExpirationWithDefaults

`func NewWhatsappMessageOrderExpirationWithDefaults() *WhatsappMessageOrderExpiration`

NewWhatsappMessageOrderExpirationWithDefaults instantiates a new WhatsappMessageOrderExpiration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *WhatsappMessageOrderExpiration) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WhatsappMessageOrderExpiration) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WhatsappMessageOrderExpiration) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetDescription

`func (o *WhatsappMessageOrderExpiration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WhatsappMessageOrderExpiration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WhatsappMessageOrderExpiration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WhatsappMessageOrderExpiration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


