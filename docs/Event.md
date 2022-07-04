# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID for the object. | 
**Type** | **string** | Type of this event. | 
**ApiVersion** | **string** | The API version used to render this event. | 
**CreateTime** | **time.Time** | The time at which this event was created, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | 
**EmailDelivery** | Pointer to [**EmailDelivery**](EmailDelivery.md) |  | [optional] 
**Sms** | Pointer to [**Sms**](Sms.md) |  | [optional] 
**Voice** | Pointer to [**Voice**](Voice.md) |  | [optional] 

## Methods

### NewEvent

`func NewEvent(id string, type_ string, apiVersion string, createTime time.Time, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Event) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Event) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Event) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *Event) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Event) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Event) SetType(v string)`

SetType sets Type field to given value.


### GetApiVersion

`func (o *Event) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Event) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Event) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetCreateTime

`func (o *Event) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *Event) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *Event) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.


### GetEmailDelivery

`func (o *Event) GetEmailDelivery() EmailDelivery`

GetEmailDelivery returns the EmailDelivery field if non-nil, zero value otherwise.

### GetEmailDeliveryOk

`func (o *Event) GetEmailDeliveryOk() (*EmailDelivery, bool)`

GetEmailDeliveryOk returns a tuple with the EmailDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDelivery

`func (o *Event) SetEmailDelivery(v EmailDelivery)`

SetEmailDelivery sets EmailDelivery field to given value.

### HasEmailDelivery

`func (o *Event) HasEmailDelivery() bool`

HasEmailDelivery returns a boolean if a field has been set.

### GetSms

`func (o *Event) GetSms() Sms`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *Event) GetSmsOk() (*Sms, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *Event) SetSms(v Sms)`

SetSms sets Sms field to given value.

### HasSms

`func (o *Event) HasSms() bool`

HasSms returns a boolean if a field has been set.

### GetVoice

`func (o *Event) GetVoice() Voice`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *Event) GetVoiceOk() (*Voice, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *Event) SetVoice(v Voice)`

SetVoice sets Voice field to given value.

### HasVoice

`func (o *Event) HasVoice() bool`

HasVoice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


