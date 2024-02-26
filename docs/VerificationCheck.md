# VerificationCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of this verification check. | 
**Valid** | **bool** | Whether the verification code is valid for this check. | 
**Status** | Pointer to [**VerificationStatus**](VerificationStatus.md) |  | [optional] 
**To** | Pointer to **string** | The recipient&#39;s phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format or email address. | [optional] 
**Channel** | Pointer to [**VerificationChannel**](VerificationChannel.md) |  | [optional] 

## Methods

### NewVerificationCheck

`func NewVerificationCheck(id string, valid bool, ) *VerificationCheck`

NewVerificationCheck instantiates a new VerificationCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationCheckWithDefaults

`func NewVerificationCheckWithDefaults() *VerificationCheck`

NewVerificationCheckWithDefaults instantiates a new VerificationCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VerificationCheck) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VerificationCheck) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VerificationCheck) SetId(v string)`

SetId sets Id field to given value.


### GetValid

`func (o *VerificationCheck) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *VerificationCheck) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *VerificationCheck) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetStatus

`func (o *VerificationCheck) GetStatus() VerificationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VerificationCheck) GetStatusOk() (*VerificationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VerificationCheck) SetStatus(v VerificationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VerificationCheck) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTo

`func (o *VerificationCheck) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *VerificationCheck) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *VerificationCheck) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *VerificationCheck) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetChannel

`func (o *VerificationCheck) GetChannel() VerificationChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *VerificationCheck) GetChannelOk() (*VerificationChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *VerificationCheck) SetChannel(v VerificationChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *VerificationCheck) HasChannel() bool`

HasChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


