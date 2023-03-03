# WhatsappMessageReaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | **string** | Specifies the &#x60;wamid&#x60; of the message received that contained the reaction. | 
**Emoji** | Pointer to **string** | **Required** when you send a &#x60;reaction&#x60; message. Set it to &#x60;\&quot;\&quot;&#x60; if you want to remove the emoji. **Optional** when you received a message from a user. This field is included when a user reacts to messages with an emoji. Otherwise, it indicates a user removed the emoji. | [optional] 

## Methods

### NewWhatsappMessageReaction

`func NewWhatsappMessageReaction(messageId string, ) *WhatsappMessageReaction`

NewWhatsappMessageReaction instantiates a new WhatsappMessageReaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageReactionWithDefaults

`func NewWhatsappMessageReactionWithDefaults() *WhatsappMessageReaction`

NewWhatsappMessageReactionWithDefaults instantiates a new WhatsappMessageReaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *WhatsappMessageReaction) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *WhatsappMessageReaction) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *WhatsappMessageReaction) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.


### GetEmoji

`func (o *WhatsappMessageReaction) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *WhatsappMessageReaction) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *WhatsappMessageReaction) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *WhatsappMessageReaction) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


