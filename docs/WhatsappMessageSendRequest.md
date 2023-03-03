# WhatsappMessageSendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | **string** | The sender&#39;s phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | 
**To** | **string** | The recipient&#39;s phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | 
**Type** | [**WhatsappMessageType**](WhatsappMessageType.md) |  | 
**Template** | Pointer to [**WhatsappMessageTemplate**](WhatsappMessageTemplate.md) |  | [optional] 
**Text** | Pointer to [**WhatsappMessageText**](WhatsappMessageText.md) |  | [optional] 
**Image** | Pointer to [**WhatsappMessageMedia**](WhatsappMessageMedia.md) |  | [optional] 
**Video** | Pointer to [**WhatsappMessageMedia**](WhatsappMessageMedia.md) |  | [optional] 
**Audio** | Pointer to [**WhatsappMessageMedia**](WhatsappMessageMedia.md) |  | [optional] 
**Document** | Pointer to [**WhatsappMessageMedia**](WhatsappMessageMedia.md) |  | [optional] 
**Sticker** | Pointer to [**WhatsappMessageMedia**](WhatsappMessageMedia.md) |  | [optional] 
**Location** | Pointer to [**WhatsappMessageLocation**](WhatsappMessageLocation.md) |  | [optional] 
**Interactive** | Pointer to [**WhatsappMessageInteractive**](WhatsappMessageInteractive.md) |  | [optional] 
**Contacts** | Pointer to [**[]WhatsappMessageContact**](WhatsappMessageContact.md) | Required when &#x60;type&#x60; is &#x60;contacts&#x60;. | [optional] 
**Reaction** | Pointer to [**WhatsappMessageReaction**](WhatsappMessageReaction.md) |  | [optional] 
**Context** | Pointer to [**WhatsappMessageContext**](WhatsappMessageContext.md) |  | [optional] 
**ExternalId** | Pointer to **string** | A unique string to reference the object. This can be an order number or similar, and can be used to reconcile the object with your internal systems. | [optional] 

## Methods

### NewWhatsappMessageSendRequest

`func NewWhatsappMessageSendRequest(from string, to string, type_ WhatsappMessageType, ) *WhatsappMessageSendRequest`

NewWhatsappMessageSendRequest instantiates a new WhatsappMessageSendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageSendRequestWithDefaults

`func NewWhatsappMessageSendRequestWithDefaults() *WhatsappMessageSendRequest`

NewWhatsappMessageSendRequestWithDefaults instantiates a new WhatsappMessageSendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *WhatsappMessageSendRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *WhatsappMessageSendRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *WhatsappMessageSendRequest) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *WhatsappMessageSendRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *WhatsappMessageSendRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *WhatsappMessageSendRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetType

`func (o *WhatsappMessageSendRequest) GetType() WhatsappMessageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappMessageSendRequest) GetTypeOk() (*WhatsappMessageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappMessageSendRequest) SetType(v WhatsappMessageType)`

SetType sets Type field to given value.


### GetTemplate

`func (o *WhatsappMessageSendRequest) GetTemplate() WhatsappMessageTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WhatsappMessageSendRequest) GetTemplateOk() (*WhatsappMessageTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WhatsappMessageSendRequest) SetTemplate(v WhatsappMessageTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *WhatsappMessageSendRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetText

`func (o *WhatsappMessageSendRequest) GetText() WhatsappMessageText`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WhatsappMessageSendRequest) GetTextOk() (*WhatsappMessageText, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WhatsappMessageSendRequest) SetText(v WhatsappMessageText)`

SetText sets Text field to given value.

### HasText

`func (o *WhatsappMessageSendRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetImage

`func (o *WhatsappMessageSendRequest) GetImage() WhatsappMessageMedia`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *WhatsappMessageSendRequest) GetImageOk() (*WhatsappMessageMedia, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *WhatsappMessageSendRequest) SetImage(v WhatsappMessageMedia)`

SetImage sets Image field to given value.

### HasImage

`func (o *WhatsappMessageSendRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetVideo

`func (o *WhatsappMessageSendRequest) GetVideo() WhatsappMessageMedia`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *WhatsappMessageSendRequest) GetVideoOk() (*WhatsappMessageMedia, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *WhatsappMessageSendRequest) SetVideo(v WhatsappMessageMedia)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *WhatsappMessageSendRequest) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetAudio

`func (o *WhatsappMessageSendRequest) GetAudio() WhatsappMessageMedia`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *WhatsappMessageSendRequest) GetAudioOk() (*WhatsappMessageMedia, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *WhatsappMessageSendRequest) SetAudio(v WhatsappMessageMedia)`

SetAudio sets Audio field to given value.

### HasAudio

`func (o *WhatsappMessageSendRequest) HasAudio() bool`

HasAudio returns a boolean if a field has been set.

### GetDocument

`func (o *WhatsappMessageSendRequest) GetDocument() WhatsappMessageMedia`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *WhatsappMessageSendRequest) GetDocumentOk() (*WhatsappMessageMedia, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *WhatsappMessageSendRequest) SetDocument(v WhatsappMessageMedia)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *WhatsappMessageSendRequest) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetSticker

`func (o *WhatsappMessageSendRequest) GetSticker() WhatsappMessageMedia`

GetSticker returns the Sticker field if non-nil, zero value otherwise.

### GetStickerOk

`func (o *WhatsappMessageSendRequest) GetStickerOk() (*WhatsappMessageMedia, bool)`

GetStickerOk returns a tuple with the Sticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSticker

`func (o *WhatsappMessageSendRequest) SetSticker(v WhatsappMessageMedia)`

SetSticker sets Sticker field to given value.

### HasSticker

`func (o *WhatsappMessageSendRequest) HasSticker() bool`

HasSticker returns a boolean if a field has been set.

### GetLocation

`func (o *WhatsappMessageSendRequest) GetLocation() WhatsappMessageLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *WhatsappMessageSendRequest) GetLocationOk() (*WhatsappMessageLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *WhatsappMessageSendRequest) SetLocation(v WhatsappMessageLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *WhatsappMessageSendRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetInteractive

`func (o *WhatsappMessageSendRequest) GetInteractive() WhatsappMessageInteractive`

GetInteractive returns the Interactive field if non-nil, zero value otherwise.

### GetInteractiveOk

`func (o *WhatsappMessageSendRequest) GetInteractiveOk() (*WhatsappMessageInteractive, bool)`

GetInteractiveOk returns a tuple with the Interactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractive

`func (o *WhatsappMessageSendRequest) SetInteractive(v WhatsappMessageInteractive)`

SetInteractive sets Interactive field to given value.

### HasInteractive

`func (o *WhatsappMessageSendRequest) HasInteractive() bool`

HasInteractive returns a boolean if a field has been set.

### GetContacts

`func (o *WhatsappMessageSendRequest) GetContacts() []WhatsappMessageContact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *WhatsappMessageSendRequest) GetContactsOk() (*[]WhatsappMessageContact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *WhatsappMessageSendRequest) SetContacts(v []WhatsappMessageContact)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *WhatsappMessageSendRequest) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetReaction

`func (o *WhatsappMessageSendRequest) GetReaction() WhatsappMessageReaction`

GetReaction returns the Reaction field if non-nil, zero value otherwise.

### GetReactionOk

`func (o *WhatsappMessageSendRequest) GetReactionOk() (*WhatsappMessageReaction, bool)`

GetReactionOk returns a tuple with the Reaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReaction

`func (o *WhatsappMessageSendRequest) SetReaction(v WhatsappMessageReaction)`

SetReaction sets Reaction field to given value.

### HasReaction

`func (o *WhatsappMessageSendRequest) HasReaction() bool`

HasReaction returns a boolean if a field has been set.

### GetContext

`func (o *WhatsappMessageSendRequest) GetContext() WhatsappMessageContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *WhatsappMessageSendRequest) GetContextOk() (*WhatsappMessageContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *WhatsappMessageSendRequest) SetContext(v WhatsappMessageContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *WhatsappMessageSendRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetExternalId

`func (o *WhatsappMessageSendRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *WhatsappMessageSendRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *WhatsappMessageSendRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *WhatsappMessageSendRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


