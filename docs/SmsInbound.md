# SmsInbound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID of the message. | [optional] 
**From** | Pointer to **string** | The user&#39;s phone number who sent the message to your registered sender ID, formatted in [E.164](https://en.wikipedia.org/wiki/E.164) format. | [optional] 
**To** | Pointer to **string** | The receiver&#39;s phone number, which is one of your registered Sender IDs. | [optional] 
**Text** | Pointer to **string** | The text of this message. | [optional] 
**SendTime** | Pointer to **time.Time** | The time at which this message was sent, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional] 

## Methods

### NewSmsInbound

`func NewSmsInbound() *SmsInbound`

NewSmsInbound instantiates a new SmsInbound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsInboundWithDefaults

`func NewSmsInboundWithDefaults() *SmsInbound`

NewSmsInboundWithDefaults instantiates a new SmsInbound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SmsInbound) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SmsInbound) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SmsInbound) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SmsInbound) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFrom

`func (o *SmsInbound) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SmsInbound) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SmsInbound) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *SmsInbound) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *SmsInbound) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SmsInbound) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SmsInbound) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *SmsInbound) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetText

`func (o *SmsInbound) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SmsInbound) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SmsInbound) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *SmsInbound) HasText() bool`

HasText returns a boolean if a field has been set.

### GetSendTime

`func (o *SmsInbound) GetSendTime() time.Time`

GetSendTime returns the SendTime field if non-nil, zero value otherwise.

### GetSendTimeOk

`func (o *SmsInbound) GetSendTimeOk() (*time.Time, bool)`

GetSendTimeOk returns a tuple with the SendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendTime

`func (o *SmsInbound) SetSendTime(v time.Time)`

SetSendTime sets SendTime field to given value.

### HasSendTime

`func (o *SmsInbound) HasSendTime() bool`

HasSendTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


