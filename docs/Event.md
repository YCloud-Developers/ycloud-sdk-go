# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID for the event. | 
**Type** | [**EventType**](EventType.md) |  | 
**ApiVersion** | **string** | The API version used to render this event. | 
**CreateTime** | **time.Time** | The time at which this event was created, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | 
**EmailDelivery** | Pointer to [**EmailDelivery**](EmailDelivery.md) |  | [optional] 
**Sms** | Pointer to [**Sms**](Sms.md) |  | [optional] 
**SmsInbound** | Pointer to [**SmsInbound**](SmsInbound.md) |  | [optional] 
**Voice** | Pointer to [**Voice**](Voice.md) |  | [optional] 
**WhatsappBusinessAccount** | Pointer to [**WhatsappBusinessAccount**](WhatsappBusinessAccount.md) |  | [optional] 
**WhatsappInboundMessage** | Pointer to [**WhatsappInboundMessage**](WhatsappInboundMessage.md) |  | [optional] 
**WhatsappMessage** | Pointer to [**WhatsappMessage**](WhatsappMessage.md) |  | [optional] 
**WhatsappPhoneNumber** | Pointer to [**WhatsappPhoneNumber**](WhatsappPhoneNumber.md) |  | [optional] 
**WhatsappTemplate** | Pointer to [**WhatsappTemplate**](WhatsappTemplate.md) |  | [optional] 

## Methods

### NewEvent

`func NewEvent(id string, type_ EventType, apiVersion string, createTime time.Time, ) *Event`

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

`func (o *Event) GetType() EventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Event) GetTypeOk() (*EventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Event) SetType(v EventType)`

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

### GetSmsInbound

`func (o *Event) GetSmsInbound() SmsInbound`

GetSmsInbound returns the SmsInbound field if non-nil, zero value otherwise.

### GetSmsInboundOk

`func (o *Event) GetSmsInboundOk() (*SmsInbound, bool)`

GetSmsInboundOk returns a tuple with the SmsInbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsInbound

`func (o *Event) SetSmsInbound(v SmsInbound)`

SetSmsInbound sets SmsInbound field to given value.

### HasSmsInbound

`func (o *Event) HasSmsInbound() bool`

HasSmsInbound returns a boolean if a field has been set.

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

### GetWhatsappBusinessAccount

`func (o *Event) GetWhatsappBusinessAccount() WhatsappBusinessAccount`

GetWhatsappBusinessAccount returns the WhatsappBusinessAccount field if non-nil, zero value otherwise.

### GetWhatsappBusinessAccountOk

`func (o *Event) GetWhatsappBusinessAccountOk() (*WhatsappBusinessAccount, bool)`

GetWhatsappBusinessAccountOk returns a tuple with the WhatsappBusinessAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhatsappBusinessAccount

`func (o *Event) SetWhatsappBusinessAccount(v WhatsappBusinessAccount)`

SetWhatsappBusinessAccount sets WhatsappBusinessAccount field to given value.

### HasWhatsappBusinessAccount

`func (o *Event) HasWhatsappBusinessAccount() bool`

HasWhatsappBusinessAccount returns a boolean if a field has been set.

### GetWhatsappInboundMessage

`func (o *Event) GetWhatsappInboundMessage() WhatsappInboundMessage`

GetWhatsappInboundMessage returns the WhatsappInboundMessage field if non-nil, zero value otherwise.

### GetWhatsappInboundMessageOk

`func (o *Event) GetWhatsappInboundMessageOk() (*WhatsappInboundMessage, bool)`

GetWhatsappInboundMessageOk returns a tuple with the WhatsappInboundMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhatsappInboundMessage

`func (o *Event) SetWhatsappInboundMessage(v WhatsappInboundMessage)`

SetWhatsappInboundMessage sets WhatsappInboundMessage field to given value.

### HasWhatsappInboundMessage

`func (o *Event) HasWhatsappInboundMessage() bool`

HasWhatsappInboundMessage returns a boolean if a field has been set.

### GetWhatsappMessage

`func (o *Event) GetWhatsappMessage() WhatsappMessage`

GetWhatsappMessage returns the WhatsappMessage field if non-nil, zero value otherwise.

### GetWhatsappMessageOk

`func (o *Event) GetWhatsappMessageOk() (*WhatsappMessage, bool)`

GetWhatsappMessageOk returns a tuple with the WhatsappMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhatsappMessage

`func (o *Event) SetWhatsappMessage(v WhatsappMessage)`

SetWhatsappMessage sets WhatsappMessage field to given value.

### HasWhatsappMessage

`func (o *Event) HasWhatsappMessage() bool`

HasWhatsappMessage returns a boolean if a field has been set.

### GetWhatsappPhoneNumber

`func (o *Event) GetWhatsappPhoneNumber() WhatsappPhoneNumber`

GetWhatsappPhoneNumber returns the WhatsappPhoneNumber field if non-nil, zero value otherwise.

### GetWhatsappPhoneNumberOk

`func (o *Event) GetWhatsappPhoneNumberOk() (*WhatsappPhoneNumber, bool)`

GetWhatsappPhoneNumberOk returns a tuple with the WhatsappPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhatsappPhoneNumber

`func (o *Event) SetWhatsappPhoneNumber(v WhatsappPhoneNumber)`

SetWhatsappPhoneNumber sets WhatsappPhoneNumber field to given value.

### HasWhatsappPhoneNumber

`func (o *Event) HasWhatsappPhoneNumber() bool`

HasWhatsappPhoneNumber returns a boolean if a field has been set.

### GetWhatsappTemplate

`func (o *Event) GetWhatsappTemplate() WhatsappTemplate`

GetWhatsappTemplate returns the WhatsappTemplate field if non-nil, zero value otherwise.

### GetWhatsappTemplateOk

`func (o *Event) GetWhatsappTemplateOk() (*WhatsappTemplate, bool)`

GetWhatsappTemplateOk returns a tuple with the WhatsappTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhatsappTemplate

`func (o *Event) SetWhatsappTemplate(v WhatsappTemplate)`

SetWhatsappTemplate sets WhatsappTemplate field to given value.

### HasWhatsappTemplate

`func (o *Event) HasWhatsappTemplate() bool`

HasWhatsappTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


