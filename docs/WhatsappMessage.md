# WhatsappMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID for the object. | 
**Wamid** | Pointer to **string** | The original message ID on WhatsApp&#39;s platform. | [optional] 
**WabaId** | **string** | WhatsApp Business Account ID. | 
**From** | **string** | The sender&#39;s phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | 
**To** | **string** | The recipient&#39;s phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. | 
**Conversation** | Pointer to [**WhatsappConversation**](WhatsappConversation.md) |  | [optional] 
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
**Contacts** | Pointer to [**[]WhatsappMessageContact**](WhatsappMessageContact.md) |  | [optional] 
**Reaction** | Pointer to [**WhatsappMessageReaction**](WhatsappMessageReaction.md) |  | [optional] 
**Context** | Pointer to [**WhatsappMessageContext**](WhatsappMessageContext.md) |  | [optional] 
**ExternalId** | Pointer to **string** | A unique string to reference the object. This can be an order number or similar, and can be used to reconcile the object with your internal systems. | [optional] 
**Status** | Pointer to [**WhatsappMessageStatus**](WhatsappMessageStatus.md) |  | [optional] 
**ErrorCode** | Pointer to **string** | Error code when the message status is &#x60;failed&#x60;. | [optional] 
**ErrorMessage** | Pointer to **string** | Error message when the message status is &#x60;failed&#x60;. | [optional] 
**CreateTime** | Pointer to **time.Time** | The time at which this message is created, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional] 
**UpdateTime** | Pointer to **time.Time** | The time at which this message is updated, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional] 
**TotalPrice** | Pointer to **float64** | Total price of this message. **Note: It&#39;s only an estimated price when the &#x60;status&#x60; is &#x60;accepted&#x60; or &#x60;sent&#x60;. It becomes the final price after the message is delivered, i.e., the &#x60;status&#x60; is &#x60;delivered&#x60; or &#x60;read&#x60;.** | [optional] 
**Currency** | Pointer to **string** | Price currency. [ISO 4217 currency code](https://en.wikipedia.org/wiki/ISO_4217). | [optional] 
**WhatsappApiError** | Pointer to [**WhatsappApiError**](WhatsappApiError.md) |  | [optional] 
**BizType** | Pointer to **string** | This can be either empty or one of &#x60;whatsapp&#x60;, or &#x60;verify&#x60;. Defaults to &#x60;whatsapp&#x60;. - &#x60;whatsapp&#x60;: Indicates that the message is sent via [WhatsApp](https://www.ycloud.com/whatsapp) product. - &#x60;verify&#x60;: Indicates that the message is sent via [Verify](https://www.ycloud.com/verify) product. | [optional] 
**VerificationId** | Pointer to **string** | The verification ID. Included only when &#x60;bizType&#x60; is &#x60;verify&#x60;. | [optional] 

## Methods

### NewWhatsappMessage

`func NewWhatsappMessage(id string, wabaId string, from string, to string, type_ WhatsappMessageType, ) *WhatsappMessage`

NewWhatsappMessage instantiates a new WhatsappMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageWithDefaults

`func NewWhatsappMessageWithDefaults() *WhatsappMessage`

NewWhatsappMessageWithDefaults instantiates a new WhatsappMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WhatsappMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WhatsappMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WhatsappMessage) SetId(v string)`

SetId sets Id field to given value.


### GetWamid

`func (o *WhatsappMessage) GetWamid() string`

GetWamid returns the Wamid field if non-nil, zero value otherwise.

### GetWamidOk

`func (o *WhatsappMessage) GetWamidOk() (*string, bool)`

GetWamidOk returns a tuple with the Wamid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWamid

`func (o *WhatsappMessage) SetWamid(v string)`

SetWamid sets Wamid field to given value.

### HasWamid

`func (o *WhatsappMessage) HasWamid() bool`

HasWamid returns a boolean if a field has been set.

### GetWabaId

`func (o *WhatsappMessage) GetWabaId() string`

GetWabaId returns the WabaId field if non-nil, zero value otherwise.

### GetWabaIdOk

`func (o *WhatsappMessage) GetWabaIdOk() (*string, bool)`

GetWabaIdOk returns a tuple with the WabaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWabaId

`func (o *WhatsappMessage) SetWabaId(v string)`

SetWabaId sets WabaId field to given value.


### GetFrom

`func (o *WhatsappMessage) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *WhatsappMessage) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *WhatsappMessage) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *WhatsappMessage) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *WhatsappMessage) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *WhatsappMessage) SetTo(v string)`

SetTo sets To field to given value.


### GetConversation

`func (o *WhatsappMessage) GetConversation() WhatsappConversation`

GetConversation returns the Conversation field if non-nil, zero value otherwise.

### GetConversationOk

`func (o *WhatsappMessage) GetConversationOk() (*WhatsappConversation, bool)`

GetConversationOk returns a tuple with the Conversation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversation

`func (o *WhatsappMessage) SetConversation(v WhatsappConversation)`

SetConversation sets Conversation field to given value.

### HasConversation

`func (o *WhatsappMessage) HasConversation() bool`

HasConversation returns a boolean if a field has been set.

### GetType

`func (o *WhatsappMessage) GetType() WhatsappMessageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappMessage) GetTypeOk() (*WhatsappMessageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappMessage) SetType(v WhatsappMessageType)`

SetType sets Type field to given value.


### GetTemplate

`func (o *WhatsappMessage) GetTemplate() WhatsappMessageTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WhatsappMessage) GetTemplateOk() (*WhatsappMessageTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WhatsappMessage) SetTemplate(v WhatsappMessageTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *WhatsappMessage) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetText

`func (o *WhatsappMessage) GetText() WhatsappMessageText`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WhatsappMessage) GetTextOk() (*WhatsappMessageText, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WhatsappMessage) SetText(v WhatsappMessageText)`

SetText sets Text field to given value.

### HasText

`func (o *WhatsappMessage) HasText() bool`

HasText returns a boolean if a field has been set.

### GetImage

`func (o *WhatsappMessage) GetImage() WhatsappMessageMedia`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *WhatsappMessage) GetImageOk() (*WhatsappMessageMedia, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *WhatsappMessage) SetImage(v WhatsappMessageMedia)`

SetImage sets Image field to given value.

### HasImage

`func (o *WhatsappMessage) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetVideo

`func (o *WhatsappMessage) GetVideo() WhatsappMessageMedia`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *WhatsappMessage) GetVideoOk() (*WhatsappMessageMedia, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *WhatsappMessage) SetVideo(v WhatsappMessageMedia)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *WhatsappMessage) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetAudio

`func (o *WhatsappMessage) GetAudio() WhatsappMessageMedia`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *WhatsappMessage) GetAudioOk() (*WhatsappMessageMedia, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *WhatsappMessage) SetAudio(v WhatsappMessageMedia)`

SetAudio sets Audio field to given value.

### HasAudio

`func (o *WhatsappMessage) HasAudio() bool`

HasAudio returns a boolean if a field has been set.

### GetDocument

`func (o *WhatsappMessage) GetDocument() WhatsappMessageMedia`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *WhatsappMessage) GetDocumentOk() (*WhatsappMessageMedia, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *WhatsappMessage) SetDocument(v WhatsappMessageMedia)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *WhatsappMessage) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetSticker

`func (o *WhatsappMessage) GetSticker() WhatsappMessageMedia`

GetSticker returns the Sticker field if non-nil, zero value otherwise.

### GetStickerOk

`func (o *WhatsappMessage) GetStickerOk() (*WhatsappMessageMedia, bool)`

GetStickerOk returns a tuple with the Sticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSticker

`func (o *WhatsappMessage) SetSticker(v WhatsappMessageMedia)`

SetSticker sets Sticker field to given value.

### HasSticker

`func (o *WhatsappMessage) HasSticker() bool`

HasSticker returns a boolean if a field has been set.

### GetLocation

`func (o *WhatsappMessage) GetLocation() WhatsappMessageLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *WhatsappMessage) GetLocationOk() (*WhatsappMessageLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *WhatsappMessage) SetLocation(v WhatsappMessageLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *WhatsappMessage) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetInteractive

`func (o *WhatsappMessage) GetInteractive() WhatsappMessageInteractive`

GetInteractive returns the Interactive field if non-nil, zero value otherwise.

### GetInteractiveOk

`func (o *WhatsappMessage) GetInteractiveOk() (*WhatsappMessageInteractive, bool)`

GetInteractiveOk returns a tuple with the Interactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractive

`func (o *WhatsappMessage) SetInteractive(v WhatsappMessageInteractive)`

SetInteractive sets Interactive field to given value.

### HasInteractive

`func (o *WhatsappMessage) HasInteractive() bool`

HasInteractive returns a boolean if a field has been set.

### GetContacts

`func (o *WhatsappMessage) GetContacts() []WhatsappMessageContact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *WhatsappMessage) GetContactsOk() (*[]WhatsappMessageContact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *WhatsappMessage) SetContacts(v []WhatsappMessageContact)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *WhatsappMessage) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetReaction

`func (o *WhatsappMessage) GetReaction() WhatsappMessageReaction`

GetReaction returns the Reaction field if non-nil, zero value otherwise.

### GetReactionOk

`func (o *WhatsappMessage) GetReactionOk() (*WhatsappMessageReaction, bool)`

GetReactionOk returns a tuple with the Reaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReaction

`func (o *WhatsappMessage) SetReaction(v WhatsappMessageReaction)`

SetReaction sets Reaction field to given value.

### HasReaction

`func (o *WhatsappMessage) HasReaction() bool`

HasReaction returns a boolean if a field has been set.

### GetContext

`func (o *WhatsappMessage) GetContext() WhatsappMessageContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *WhatsappMessage) GetContextOk() (*WhatsappMessageContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *WhatsappMessage) SetContext(v WhatsappMessageContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *WhatsappMessage) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetExternalId

`func (o *WhatsappMessage) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *WhatsappMessage) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *WhatsappMessage) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *WhatsappMessage) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetStatus

`func (o *WhatsappMessage) GetStatus() WhatsappMessageStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WhatsappMessage) GetStatusOk() (*WhatsappMessageStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WhatsappMessage) SetStatus(v WhatsappMessageStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WhatsappMessage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorCode

`func (o *WhatsappMessage) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *WhatsappMessage) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *WhatsappMessage) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *WhatsappMessage) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *WhatsappMessage) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *WhatsappMessage) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *WhatsappMessage) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *WhatsappMessage) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetCreateTime

`func (o *WhatsappMessage) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *WhatsappMessage) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *WhatsappMessage) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *WhatsappMessage) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *WhatsappMessage) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *WhatsappMessage) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *WhatsappMessage) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *WhatsappMessage) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetTotalPrice

`func (o *WhatsappMessage) GetTotalPrice() float64`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *WhatsappMessage) GetTotalPriceOk() (*float64, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *WhatsappMessage) SetTotalPrice(v float64)`

SetTotalPrice sets TotalPrice field to given value.

### HasTotalPrice

`func (o *WhatsappMessage) HasTotalPrice() bool`

HasTotalPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *WhatsappMessage) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *WhatsappMessage) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *WhatsappMessage) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *WhatsappMessage) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetWhatsappApiError

`func (o *WhatsappMessage) GetWhatsappApiError() WhatsappApiError`

GetWhatsappApiError returns the WhatsappApiError field if non-nil, zero value otherwise.

### GetWhatsappApiErrorOk

`func (o *WhatsappMessage) GetWhatsappApiErrorOk() (*WhatsappApiError, bool)`

GetWhatsappApiErrorOk returns a tuple with the WhatsappApiError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhatsappApiError

`func (o *WhatsappMessage) SetWhatsappApiError(v WhatsappApiError)`

SetWhatsappApiError sets WhatsappApiError field to given value.

### HasWhatsappApiError

`func (o *WhatsappMessage) HasWhatsappApiError() bool`

HasWhatsappApiError returns a boolean if a field has been set.

### GetBizType

`func (o *WhatsappMessage) GetBizType() string`

GetBizType returns the BizType field if non-nil, zero value otherwise.

### GetBizTypeOk

`func (o *WhatsappMessage) GetBizTypeOk() (*string, bool)`

GetBizTypeOk returns a tuple with the BizType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizType

`func (o *WhatsappMessage) SetBizType(v string)`

SetBizType sets BizType field to given value.

### HasBizType

`func (o *WhatsappMessage) HasBizType() bool`

HasBizType returns a boolean if a field has been set.

### GetVerificationId

`func (o *WhatsappMessage) GetVerificationId() string`

GetVerificationId returns the VerificationId field if non-nil, zero value otherwise.

### GetVerificationIdOk

`func (o *WhatsappMessage) GetVerificationIdOk() (*string, bool)`

GetVerificationIdOk returns a tuple with the VerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationId

`func (o *WhatsappMessage) SetVerificationId(v string)`

SetVerificationId sets VerificationId field to given value.

### HasVerificationId

`func (o *WhatsappMessage) HasVerificationId() bool`

HasVerificationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


