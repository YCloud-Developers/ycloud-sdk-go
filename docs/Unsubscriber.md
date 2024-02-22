# Unsubscriber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**UnsubscriberType**](UnsubscriberType.md) |  | [optional] 
**Customer** | Pointer to **string** | The customer who has opted out. For &#x60;type&#x3D;PHONE_NUMBER&#x60;, it should be a phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | [optional] 
**Channel** | Pointer to [**UnsubscriberChannel**](UnsubscriberChannel.md) |  | [optional] 
**RegionCode** | Pointer to **string** | The customer&#39;s region code, formatted in [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). | [optional] 
**CreateTime** | Pointer to **time.Time** | The time at which this object was created, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional] 

## Methods

### NewUnsubscriber

`func NewUnsubscriber() *Unsubscriber`

NewUnsubscriber instantiates a new Unsubscriber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnsubscriberWithDefaults

`func NewUnsubscriberWithDefaults() *Unsubscriber`

NewUnsubscriberWithDefaults instantiates a new Unsubscriber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Unsubscriber) GetType() UnsubscriberType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Unsubscriber) GetTypeOk() (*UnsubscriberType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Unsubscriber) SetType(v UnsubscriberType)`

SetType sets Type field to given value.

### HasType

`func (o *Unsubscriber) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCustomer

`func (o *Unsubscriber) GetCustomer() string`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *Unsubscriber) GetCustomerOk() (*string, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *Unsubscriber) SetCustomer(v string)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *Unsubscriber) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetChannel

`func (o *Unsubscriber) GetChannel() UnsubscriberChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *Unsubscriber) GetChannelOk() (*UnsubscriberChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *Unsubscriber) SetChannel(v UnsubscriberChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *Unsubscriber) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetRegionCode

`func (o *Unsubscriber) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *Unsubscriber) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *Unsubscriber) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *Unsubscriber) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetCreateTime

`func (o *Unsubscriber) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *Unsubscriber) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *Unsubscriber) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *Unsubscriber) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
