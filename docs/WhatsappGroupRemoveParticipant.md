# WhatsappGroupRemoveParticipant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **string** | WhatsApp user ID to remove. | [optional] 
**FromUserId** | Pointer to **string** | Business-scoped user ID to remove. Also accepts &#x60;userId&#x60; as an alias. | [optional] 

## Methods

### NewWhatsappGroupRemoveParticipant

`func NewWhatsappGroupRemoveParticipant() *WhatsappGroupRemoveParticipant`

NewWhatsappGroupRemoveParticipant instantiates a new WhatsappGroupRemoveParticipant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappGroupRemoveParticipantWithDefaults

`func NewWhatsappGroupRemoveParticipantWithDefaults() *WhatsappGroupRemoveParticipant`

NewWhatsappGroupRemoveParticipantWithDefaults instantiates a new WhatsappGroupRemoveParticipant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *WhatsappGroupRemoveParticipant) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *WhatsappGroupRemoveParticipant) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *WhatsappGroupRemoveParticipant) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *WhatsappGroupRemoveParticipant) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetFromUserId

`func (o *WhatsappGroupRemoveParticipant) GetFromUserId() string`

GetFromUserId returns the FromUserId field if non-nil, zero value otherwise.

### GetFromUserIdOk

`func (o *WhatsappGroupRemoveParticipant) GetFromUserIdOk() (*string, bool)`

GetFromUserIdOk returns a tuple with the FromUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserId

`func (o *WhatsappGroupRemoveParticipant) SetFromUserId(v string)`

SetFromUserId sets FromUserId field to given value.

### HasFromUserId

`func (o *WhatsappGroupRemoveParticipant) HasFromUserId() bool`

HasFromUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


