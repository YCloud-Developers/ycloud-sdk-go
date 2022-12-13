# WhatsappInboundMessageReaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | Pointer to **string** | Specifies the &#x60;wamid&#x60; of the message received that contained the reaction. | [optional] 
**Emoji** | Pointer to **string** | This field is included when a user reacts to messages with an emoji. Otherwise, it indicates a user removed the emoji. | [optional] 

## Methods

### NewWhatsappInboundMessageReaction

`func NewWhatsappInboundMessageReaction() *WhatsappInboundMessageReaction`

NewWhatsappInboundMessageReaction instantiates a new WhatsappInboundMessageReaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappInboundMessageReactionWithDefaults

`func NewWhatsappInboundMessageReactionWithDefaults() *WhatsappInboundMessageReaction`

NewWhatsappInboundMessageReactionWithDefaults instantiates a new WhatsappInboundMessageReaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *WhatsappInboundMessageReaction) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *WhatsappInboundMessageReaction) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *WhatsappInboundMessageReaction) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *WhatsappInboundMessageReaction) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetEmoji

`func (o *WhatsappInboundMessageReaction) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *WhatsappInboundMessageReaction) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *WhatsappInboundMessageReaction) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *WhatsappInboundMessageReaction) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


