# UnsubscriberCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**UnsubscriberType**](UnsubscriberType.md) |  | 
**Customer** | **string** | The customer who has opted out. For &#x60;type&#x3D;PHONE_NUMBER&#x60;, it should be a phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | 
**Channel** | [**UnsubscriberChannel**](UnsubscriberChannel.md) |  | 
**RegionCode** | Pointer to **string** | The customer&#39;s region code, formatted in [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). | [optional] 

## Methods

### NewUnsubscriberCreateRequest

`func NewUnsubscriberCreateRequest(type_ UnsubscriberType, customer string, channel UnsubscriberChannel, ) *UnsubscriberCreateRequest`

NewUnsubscriberCreateRequest instantiates a new UnsubscriberCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnsubscriberCreateRequestWithDefaults

`func NewUnsubscriberCreateRequestWithDefaults() *UnsubscriberCreateRequest`

NewUnsubscriberCreateRequestWithDefaults instantiates a new UnsubscriberCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UnsubscriberCreateRequest) GetType() UnsubscriberType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnsubscriberCreateRequest) GetTypeOk() (*UnsubscriberType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnsubscriberCreateRequest) SetType(v UnsubscriberType)`

SetType sets Type field to given value.


### GetCustomer

`func (o *UnsubscriberCreateRequest) GetCustomer() string`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *UnsubscriberCreateRequest) GetCustomerOk() (*string, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *UnsubscriberCreateRequest) SetCustomer(v string)`

SetCustomer sets Customer field to given value.


### GetChannel

`func (o *UnsubscriberCreateRequest) GetChannel() UnsubscriberChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *UnsubscriberCreateRequest) GetChannelOk() (*UnsubscriberChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *UnsubscriberCreateRequest) SetChannel(v UnsubscriberChannel)`

SetChannel sets Channel field to given value.


### GetRegionCode

`func (o *UnsubscriberCreateRequest) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *UnsubscriberCreateRequest) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *UnsubscriberCreateRequest) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *UnsubscriberCreateRequest) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


