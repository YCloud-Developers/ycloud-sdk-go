# WhatsappGroupWebhookMessageStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | WhatsApp message ID. | [optional] 
**Status** | Pointer to **string** | Message status. | [optional] 
**Timestamp** | Pointer to **int64** | Unix timestamp indicating when the message status was updated. | [optional] 
**RecipientId** | Pointer to **string** | Recipient group ID. | [optional] 
**RecipientType** | Pointer to **string** | Recipient type. | [optional] 
**RecipientParticipantId** | Pointer to **string** | WhatsApp user ID of the recipient participant. | [optional] 
**RecipientUserId** | Pointer to **string** | Business-scoped user ID of the recipient participant. | [optional] 
**ParentRecipientUserId** | Pointer to **string** | Parent business-scoped user ID of the recipient participant. | [optional] 
**Conversation** | Pointer to [**WhatsappGroupWebhookConversation**](WhatsappGroupWebhookConversation.md) |  | [optional] 
**Pricing** | Pointer to [**WhatsappGroupWebhookPricing**](WhatsappGroupWebhookPricing.md) |  | [optional] 
**Errors** | Pointer to **[]map[string]interface{}** | Errors returned by WhatsApp. | [optional] 

## Methods

### NewWhatsappGroupWebhookMessageStatus

`func NewWhatsappGroupWebhookMessageStatus() *WhatsappGroupWebhookMessageStatus`

NewWhatsappGroupWebhookMessageStatus instantiates a new WhatsappGroupWebhookMessageStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappGroupWebhookMessageStatusWithDefaults

`func NewWhatsappGroupWebhookMessageStatusWithDefaults() *WhatsappGroupWebhookMessageStatus`

NewWhatsappGroupWebhookMessageStatusWithDefaults instantiates a new WhatsappGroupWebhookMessageStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WhatsappGroupWebhookMessageStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WhatsappGroupWebhookMessageStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WhatsappGroupWebhookMessageStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WhatsappGroupWebhookMessageStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *WhatsappGroupWebhookMessageStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WhatsappGroupWebhookMessageStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WhatsappGroupWebhookMessageStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WhatsappGroupWebhookMessageStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimestamp

`func (o *WhatsappGroupWebhookMessageStatus) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WhatsappGroupWebhookMessageStatus) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WhatsappGroupWebhookMessageStatus) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WhatsappGroupWebhookMessageStatus) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetRecipientId

`func (o *WhatsappGroupWebhookMessageStatus) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *WhatsappGroupWebhookMessageStatus) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *WhatsappGroupWebhookMessageStatus) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.

### HasRecipientId

`func (o *WhatsappGroupWebhookMessageStatus) HasRecipientId() bool`

HasRecipientId returns a boolean if a field has been set.

### GetRecipientType

`func (o *WhatsappGroupWebhookMessageStatus) GetRecipientType() string`

GetRecipientType returns the RecipientType field if non-nil, zero value otherwise.

### GetRecipientTypeOk

`func (o *WhatsappGroupWebhookMessageStatus) GetRecipientTypeOk() (*string, bool)`

GetRecipientTypeOk returns a tuple with the RecipientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientType

`func (o *WhatsappGroupWebhookMessageStatus) SetRecipientType(v string)`

SetRecipientType sets RecipientType field to given value.

### HasRecipientType

`func (o *WhatsappGroupWebhookMessageStatus) HasRecipientType() bool`

HasRecipientType returns a boolean if a field has been set.

### GetRecipientParticipantId

`func (o *WhatsappGroupWebhookMessageStatus) GetRecipientParticipantId() string`

GetRecipientParticipantId returns the RecipientParticipantId field if non-nil, zero value otherwise.

### GetRecipientParticipantIdOk

`func (o *WhatsappGroupWebhookMessageStatus) GetRecipientParticipantIdOk() (*string, bool)`

GetRecipientParticipantIdOk returns a tuple with the RecipientParticipantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientParticipantId

`func (o *WhatsappGroupWebhookMessageStatus) SetRecipientParticipantId(v string)`

SetRecipientParticipantId sets RecipientParticipantId field to given value.

### HasRecipientParticipantId

`func (o *WhatsappGroupWebhookMessageStatus) HasRecipientParticipantId() bool`

HasRecipientParticipantId returns a boolean if a field has been set.

### GetRecipientUserId

`func (o *WhatsappGroupWebhookMessageStatus) GetRecipientUserId() string`

GetRecipientUserId returns the RecipientUserId field if non-nil, zero value otherwise.

### GetRecipientUserIdOk

`func (o *WhatsappGroupWebhookMessageStatus) GetRecipientUserIdOk() (*string, bool)`

GetRecipientUserIdOk returns a tuple with the RecipientUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientUserId

`func (o *WhatsappGroupWebhookMessageStatus) SetRecipientUserId(v string)`

SetRecipientUserId sets RecipientUserId field to given value.

### HasRecipientUserId

`func (o *WhatsappGroupWebhookMessageStatus) HasRecipientUserId() bool`

HasRecipientUserId returns a boolean if a field has been set.

### GetParentRecipientUserId

`func (o *WhatsappGroupWebhookMessageStatus) GetParentRecipientUserId() string`

GetParentRecipientUserId returns the ParentRecipientUserId field if non-nil, zero value otherwise.

### GetParentRecipientUserIdOk

`func (o *WhatsappGroupWebhookMessageStatus) GetParentRecipientUserIdOk() (*string, bool)`

GetParentRecipientUserIdOk returns a tuple with the ParentRecipientUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRecipientUserId

`func (o *WhatsappGroupWebhookMessageStatus) SetParentRecipientUserId(v string)`

SetParentRecipientUserId sets ParentRecipientUserId field to given value.

### HasParentRecipientUserId

`func (o *WhatsappGroupWebhookMessageStatus) HasParentRecipientUserId() bool`

HasParentRecipientUserId returns a boolean if a field has been set.

### GetConversation

`func (o *WhatsappGroupWebhookMessageStatus) GetConversation() WhatsappGroupWebhookConversation`

GetConversation returns the Conversation field if non-nil, zero value otherwise.

### GetConversationOk

`func (o *WhatsappGroupWebhookMessageStatus) GetConversationOk() (*WhatsappGroupWebhookConversation, bool)`

GetConversationOk returns a tuple with the Conversation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversation

`func (o *WhatsappGroupWebhookMessageStatus) SetConversation(v WhatsappGroupWebhookConversation)`

SetConversation sets Conversation field to given value.

### HasConversation

`func (o *WhatsappGroupWebhookMessageStatus) HasConversation() bool`

HasConversation returns a boolean if a field has been set.

### GetPricing

`func (o *WhatsappGroupWebhookMessageStatus) GetPricing() WhatsappGroupWebhookPricing`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *WhatsappGroupWebhookMessageStatus) GetPricingOk() (*WhatsappGroupWebhookPricing, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *WhatsappGroupWebhookMessageStatus) SetPricing(v WhatsappGroupWebhookPricing)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *WhatsappGroupWebhookMessageStatus) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### GetErrors

`func (o *WhatsappGroupWebhookMessageStatus) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *WhatsappGroupWebhookMessageStatus) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *WhatsappGroupWebhookMessageStatus) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *WhatsappGroupWebhookMessageStatus) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


