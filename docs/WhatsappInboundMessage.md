# WhatsappInboundMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID for the object. | 
**Wamid** | Pointer to **string** | The native WhatsApp message ID. | [optional] 
**WabaId** | Pointer to **string** | WhatsApp Business Account ID. | [optional] 
**From** | Pointer to **string** | The customer&#39;s phone number who sent the message to the business, formatted in [E.164](https://en.wikipedia.org/wiki/E.164) format. | [optional] 
**CustomerProfile** | Pointer to [**WhatsappProfile**](WhatsappProfile.md) |  | [optional] 
**To** | Pointer to **string** | The recipient&#39;s phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | [optional] 
**SendTime** | Pointer to **time.Time** | The time at which this message is sent, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional] 
**Type** | Pointer to [**WhatsappInboundMessageType**](WhatsappInboundMessageType.md) |  | [optional] 
**Text** | Pointer to [**WhatsappInboundMessageText**](WhatsappInboundMessageText.md) |  | [optional] 
**Image** | Pointer to [**WhatsappInboundMessageMedia**](WhatsappInboundMessageMedia.md) |  | [optional] 
**Video** | Pointer to [**WhatsappInboundMessageMedia**](WhatsappInboundMessageMedia.md) |  | [optional] 
**Audio** | Pointer to [**WhatsappInboundMessageMedia**](WhatsappInboundMessageMedia.md) |  | [optional] 
**Document** | Pointer to [**WhatsappInboundMessageMedia**](WhatsappInboundMessageMedia.md) |  | [optional] 
**Sticker** | Pointer to [**WhatsappInboundMessageMedia**](WhatsappInboundMessageMedia.md) |  | [optional] 
**Interactive** | Pointer to [**WhatsappInboundMessageInteractive**](WhatsappInboundMessageInteractive.md) |  | [optional] 
**Location** | Pointer to [**WhatsappInboundMessageLocation**](WhatsappInboundMessageLocation.md) |  | [optional] 
**Button** | Pointer to [**WhatsappInboundMessageButton**](WhatsappInboundMessageButton.md) |  | [optional] 
**Contacts** | Pointer to [**[]WhatsappMessageContact**](WhatsappMessageContact.md) |  | [optional] 
**Reaction** | Pointer to [**WhatsappInboundMessageReaction**](WhatsappInboundMessageReaction.md) |  | [optional] 
**Order** | Pointer to [**WhatsappInboundMessageOrder**](WhatsappInboundMessageOrder.md) |  | [optional] 
**System** | Pointer to [**WhatsappInboundMessageSystem**](WhatsappInboundMessageSystem.md) |  | [optional] 
**Errors** | Pointer to [**[]WhatsappInboundMessageError**](WhatsappInboundMessageError.md) |  | [optional] 
**Context** | Pointer to [**WhatsappInboundMessageContext**](WhatsappInboundMessageContext.md) |  | [optional] 
**Referral** | Pointer to [**WhatsappInboundMessageReferral**](WhatsappInboundMessageReferral.md) |  | [optional] 

## Methods

### NewWhatsappInboundMessage

`func NewWhatsappInboundMessage(id string, ) *WhatsappInboundMessage`

NewWhatsappInboundMessage instantiates a new WhatsappInboundMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappInboundMessageWithDefaults

`func NewWhatsappInboundMessageWithDefaults() *WhatsappInboundMessage`

NewWhatsappInboundMessageWithDefaults instantiates a new WhatsappInboundMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WhatsappInboundMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WhatsappInboundMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WhatsappInboundMessage) SetId(v string)`

SetId sets Id field to given value.


### GetWamid

`func (o *WhatsappInboundMessage) GetWamid() string`

GetWamid returns the Wamid field if non-nil, zero value otherwise.

### GetWamidOk

`func (o *WhatsappInboundMessage) GetWamidOk() (*string, bool)`

GetWamidOk returns a tuple with the Wamid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWamid

`func (o *WhatsappInboundMessage) SetWamid(v string)`

SetWamid sets Wamid field to given value.

### HasWamid

`func (o *WhatsappInboundMessage) HasWamid() bool`

HasWamid returns a boolean if a field has been set.

### GetWabaId

`func (o *WhatsappInboundMessage) GetWabaId() string`

GetWabaId returns the WabaId field if non-nil, zero value otherwise.

### GetWabaIdOk

`func (o *WhatsappInboundMessage) GetWabaIdOk() (*string, bool)`

GetWabaIdOk returns a tuple with the WabaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWabaId

`func (o *WhatsappInboundMessage) SetWabaId(v string)`

SetWabaId sets WabaId field to given value.

### HasWabaId

`func (o *WhatsappInboundMessage) HasWabaId() bool`

HasWabaId returns a boolean if a field has been set.

### GetFrom

`func (o *WhatsappInboundMessage) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *WhatsappInboundMessage) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *WhatsappInboundMessage) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *WhatsappInboundMessage) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetCustomerProfile

`func (o *WhatsappInboundMessage) GetCustomerProfile() WhatsappProfile`

GetCustomerProfile returns the CustomerProfile field if non-nil, zero value otherwise.

### GetCustomerProfileOk

`func (o *WhatsappInboundMessage) GetCustomerProfileOk() (*WhatsappProfile, bool)`

GetCustomerProfileOk returns a tuple with the CustomerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerProfile

`func (o *WhatsappInboundMessage) SetCustomerProfile(v WhatsappProfile)`

SetCustomerProfile sets CustomerProfile field to given value.

### HasCustomerProfile

`func (o *WhatsappInboundMessage) HasCustomerProfile() bool`

HasCustomerProfile returns a boolean if a field has been set.

### GetTo

`func (o *WhatsappInboundMessage) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *WhatsappInboundMessage) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *WhatsappInboundMessage) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *WhatsappInboundMessage) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetSendTime

`func (o *WhatsappInboundMessage) GetSendTime() time.Time`

GetSendTime returns the SendTime field if non-nil, zero value otherwise.

### GetSendTimeOk

`func (o *WhatsappInboundMessage) GetSendTimeOk() (*time.Time, bool)`

GetSendTimeOk returns a tuple with the SendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendTime

`func (o *WhatsappInboundMessage) SetSendTime(v time.Time)`

SetSendTime sets SendTime field to given value.

### HasSendTime

`func (o *WhatsappInboundMessage) HasSendTime() bool`

HasSendTime returns a boolean if a field has been set.

### GetType

`func (o *WhatsappInboundMessage) GetType() WhatsappInboundMessageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappInboundMessage) GetTypeOk() (*WhatsappInboundMessageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappInboundMessage) SetType(v WhatsappInboundMessageType)`

SetType sets Type field to given value.

### HasType

`func (o *WhatsappInboundMessage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetText

`func (o *WhatsappInboundMessage) GetText() WhatsappInboundMessageText`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WhatsappInboundMessage) GetTextOk() (*WhatsappInboundMessageText, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WhatsappInboundMessage) SetText(v WhatsappInboundMessageText)`

SetText sets Text field to given value.

### HasText

`func (o *WhatsappInboundMessage) HasText() bool`

HasText returns a boolean if a field has been set.

### GetImage

`func (o *WhatsappInboundMessage) GetImage() WhatsappInboundMessageMedia`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *WhatsappInboundMessage) GetImageOk() (*WhatsappInboundMessageMedia, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *WhatsappInboundMessage) SetImage(v WhatsappInboundMessageMedia)`

SetImage sets Image field to given value.

### HasImage

`func (o *WhatsappInboundMessage) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetVideo

`func (o *WhatsappInboundMessage) GetVideo() WhatsappInboundMessageMedia`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *WhatsappInboundMessage) GetVideoOk() (*WhatsappInboundMessageMedia, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *WhatsappInboundMessage) SetVideo(v WhatsappInboundMessageMedia)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *WhatsappInboundMessage) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetAudio

`func (o *WhatsappInboundMessage) GetAudio() WhatsappInboundMessageMedia`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *WhatsappInboundMessage) GetAudioOk() (*WhatsappInboundMessageMedia, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *WhatsappInboundMessage) SetAudio(v WhatsappInboundMessageMedia)`

SetAudio sets Audio field to given value.

### HasAudio

`func (o *WhatsappInboundMessage) HasAudio() bool`

HasAudio returns a boolean if a field has been set.

### GetDocument

`func (o *WhatsappInboundMessage) GetDocument() WhatsappInboundMessageMedia`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *WhatsappInboundMessage) GetDocumentOk() (*WhatsappInboundMessageMedia, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *WhatsappInboundMessage) SetDocument(v WhatsappInboundMessageMedia)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *WhatsappInboundMessage) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetSticker

`func (o *WhatsappInboundMessage) GetSticker() WhatsappInboundMessageMedia`

GetSticker returns the Sticker field if non-nil, zero value otherwise.

### GetStickerOk

`func (o *WhatsappInboundMessage) GetStickerOk() (*WhatsappInboundMessageMedia, bool)`

GetStickerOk returns a tuple with the Sticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSticker

`func (o *WhatsappInboundMessage) SetSticker(v WhatsappInboundMessageMedia)`

SetSticker sets Sticker field to given value.

### HasSticker

`func (o *WhatsappInboundMessage) HasSticker() bool`

HasSticker returns a boolean if a field has been set.

### GetInteractive

`func (o *WhatsappInboundMessage) GetInteractive() WhatsappInboundMessageInteractive`

GetInteractive returns the Interactive field if non-nil, zero value otherwise.

### GetInteractiveOk

`func (o *WhatsappInboundMessage) GetInteractiveOk() (*WhatsappInboundMessageInteractive, bool)`

GetInteractiveOk returns a tuple with the Interactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractive

`func (o *WhatsappInboundMessage) SetInteractive(v WhatsappInboundMessageInteractive)`

SetInteractive sets Interactive field to given value.

### HasInteractive

`func (o *WhatsappInboundMessage) HasInteractive() bool`

HasInteractive returns a boolean if a field has been set.

### GetLocation

`func (o *WhatsappInboundMessage) GetLocation() WhatsappInboundMessageLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *WhatsappInboundMessage) GetLocationOk() (*WhatsappInboundMessageLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *WhatsappInboundMessage) SetLocation(v WhatsappInboundMessageLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *WhatsappInboundMessage) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetButton

`func (o *WhatsappInboundMessage) GetButton() WhatsappInboundMessageButton`

GetButton returns the Button field if non-nil, zero value otherwise.

### GetButtonOk

`func (o *WhatsappInboundMessage) GetButtonOk() (*WhatsappInboundMessageButton, bool)`

GetButtonOk returns a tuple with the Button field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButton

`func (o *WhatsappInboundMessage) SetButton(v WhatsappInboundMessageButton)`

SetButton sets Button field to given value.

### HasButton

`func (o *WhatsappInboundMessage) HasButton() bool`

HasButton returns a boolean if a field has been set.

### GetContacts

`func (o *WhatsappInboundMessage) GetContacts() []WhatsappMessageContact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *WhatsappInboundMessage) GetContactsOk() (*[]WhatsappMessageContact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *WhatsappInboundMessage) SetContacts(v []WhatsappMessageContact)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *WhatsappInboundMessage) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetReaction

`func (o *WhatsappInboundMessage) GetReaction() WhatsappInboundMessageReaction`

GetReaction returns the Reaction field if non-nil, zero value otherwise.

### GetReactionOk

`func (o *WhatsappInboundMessage) GetReactionOk() (*WhatsappInboundMessageReaction, bool)`

GetReactionOk returns a tuple with the Reaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReaction

`func (o *WhatsappInboundMessage) SetReaction(v WhatsappInboundMessageReaction)`

SetReaction sets Reaction field to given value.

### HasReaction

`func (o *WhatsappInboundMessage) HasReaction() bool`

HasReaction returns a boolean if a field has been set.

### GetOrder

`func (o *WhatsappInboundMessage) GetOrder() WhatsappInboundMessageOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *WhatsappInboundMessage) GetOrderOk() (*WhatsappInboundMessageOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *WhatsappInboundMessage) SetOrder(v WhatsappInboundMessageOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *WhatsappInboundMessage) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetSystem

`func (o *WhatsappInboundMessage) GetSystem() WhatsappInboundMessageSystem`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *WhatsappInboundMessage) GetSystemOk() (*WhatsappInboundMessageSystem, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *WhatsappInboundMessage) SetSystem(v WhatsappInboundMessageSystem)`

SetSystem sets System field to given value.

### HasSystem

`func (o *WhatsappInboundMessage) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetErrors

`func (o *WhatsappInboundMessage) GetErrors() []WhatsappInboundMessageError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *WhatsappInboundMessage) GetErrorsOk() (*[]WhatsappInboundMessageError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *WhatsappInboundMessage) SetErrors(v []WhatsappInboundMessageError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *WhatsappInboundMessage) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetContext

`func (o *WhatsappInboundMessage) GetContext() WhatsappInboundMessageContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *WhatsappInboundMessage) GetContextOk() (*WhatsappInboundMessageContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *WhatsappInboundMessage) SetContext(v WhatsappInboundMessageContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *WhatsappInboundMessage) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetReferral

`func (o *WhatsappInboundMessage) GetReferral() WhatsappInboundMessageReferral`

GetReferral returns the Referral field if non-nil, zero value otherwise.

### GetReferralOk

`func (o *WhatsappInboundMessage) GetReferralOk() (*WhatsappInboundMessageReferral, bool)`

GetReferralOk returns a tuple with the Referral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferral

`func (o *WhatsappInboundMessage) SetReferral(v WhatsappInboundMessageReferral)`

SetReferral sets Referral field to given value.

### HasReferral

`func (o *WhatsappInboundMessage) HasReferral() bool`

HasReferral returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


