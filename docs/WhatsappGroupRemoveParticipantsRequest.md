# WhatsappGroupRemoveParticipantsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Participants** | [**[]WhatsappGroupRemoveParticipant**](WhatsappGroupRemoveParticipant.md) | Participants to remove. Up to 8 participants are supported in one request. | 

## Methods

### NewWhatsappGroupRemoveParticipantsRequest

`func NewWhatsappGroupRemoveParticipantsRequest(participants []WhatsappGroupRemoveParticipant, ) *WhatsappGroupRemoveParticipantsRequest`

NewWhatsappGroupRemoveParticipantsRequest instantiates a new WhatsappGroupRemoveParticipantsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappGroupRemoveParticipantsRequestWithDefaults

`func NewWhatsappGroupRemoveParticipantsRequestWithDefaults() *WhatsappGroupRemoveParticipantsRequest`

NewWhatsappGroupRemoveParticipantsRequestWithDefaults instantiates a new WhatsappGroupRemoveParticipantsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParticipants

`func (o *WhatsappGroupRemoveParticipantsRequest) GetParticipants() []WhatsappGroupRemoveParticipant`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *WhatsappGroupRemoveParticipantsRequest) GetParticipantsOk() (*[]WhatsappGroupRemoveParticipant, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *WhatsappGroupRemoveParticipantsRequest) SetParticipants(v []WhatsappGroupRemoveParticipant)`

SetParticipants sets Participants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


