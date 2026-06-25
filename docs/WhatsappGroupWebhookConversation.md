# WhatsappGroupWebhookConversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Conversation ID. | [optional] 
**ExpirationTimestamp** | Pointer to **int64** | Unix timestamp indicating when the conversation expires. | [optional] 
**Origin** | Pointer to [**WhatsappGroupWebhookConversationOrigin**](WhatsappGroupWebhookConversationOrigin.md) |  | [optional] 

## Methods

### NewWhatsappGroupWebhookConversation

`func NewWhatsappGroupWebhookConversation() *WhatsappGroupWebhookConversation`

NewWhatsappGroupWebhookConversation instantiates a new WhatsappGroupWebhookConversation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappGroupWebhookConversationWithDefaults

`func NewWhatsappGroupWebhookConversationWithDefaults() *WhatsappGroupWebhookConversation`

NewWhatsappGroupWebhookConversationWithDefaults instantiates a new WhatsappGroupWebhookConversation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WhatsappGroupWebhookConversation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WhatsappGroupWebhookConversation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WhatsappGroupWebhookConversation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WhatsappGroupWebhookConversation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExpirationTimestamp

`func (o *WhatsappGroupWebhookConversation) GetExpirationTimestamp() int64`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *WhatsappGroupWebhookConversation) GetExpirationTimestampOk() (*int64, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *WhatsappGroupWebhookConversation) SetExpirationTimestamp(v int64)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.

### HasExpirationTimestamp

`func (o *WhatsappGroupWebhookConversation) HasExpirationTimestamp() bool`

HasExpirationTimestamp returns a boolean if a field has been set.

### GetOrigin

`func (o *WhatsappGroupWebhookConversation) GetOrigin() WhatsappGroupWebhookConversationOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *WhatsappGroupWebhookConversation) GetOriginOk() (*WhatsappGroupWebhookConversationOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *WhatsappGroupWebhookConversation) SetOrigin(v WhatsappGroupWebhookConversationOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *WhatsappGroupWebhookConversation) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


