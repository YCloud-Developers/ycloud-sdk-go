# WhatsappGroupWebhookParticipant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | Pointer to **string** | The original participant input. | [optional] 
**WaId** | Pointer to **string** | WhatsApp user ID. | [optional] 
**RecipientUserId** | Pointer to **string** | Business-scoped user ID. | [optional] 
**ParentRecipientUserId** | Pointer to **string** | Parent business-scoped user ID. | [optional] 
**CustomerProfile** | Pointer to [**WhatsappGroupCustomerProfile**](WhatsappGroupCustomerProfile.md) |  | [optional] 
**Errors** | Pointer to **[]map[string]interface{}** | Errors returned by WhatsApp for this participant. | [optional] 

## Methods

### NewWhatsappGroupWebhookParticipant

`func NewWhatsappGroupWebhookParticipant() *WhatsappGroupWebhookParticipant`

NewWhatsappGroupWebhookParticipant instantiates a new WhatsappGroupWebhookParticipant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappGroupWebhookParticipantWithDefaults

`func NewWhatsappGroupWebhookParticipantWithDefaults() *WhatsappGroupWebhookParticipant`

NewWhatsappGroupWebhookParticipantWithDefaults instantiates a new WhatsappGroupWebhookParticipant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *WhatsappGroupWebhookParticipant) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WhatsappGroupWebhookParticipant) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WhatsappGroupWebhookParticipant) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *WhatsappGroupWebhookParticipant) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetWaId

`func (o *WhatsappGroupWebhookParticipant) GetWaId() string`

GetWaId returns the WaId field if non-nil, zero value otherwise.

### GetWaIdOk

`func (o *WhatsappGroupWebhookParticipant) GetWaIdOk() (*string, bool)`

GetWaIdOk returns a tuple with the WaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaId

`func (o *WhatsappGroupWebhookParticipant) SetWaId(v string)`

SetWaId sets WaId field to given value.

### HasWaId

`func (o *WhatsappGroupWebhookParticipant) HasWaId() bool`

HasWaId returns a boolean if a field has been set.

### GetRecipientUserId

`func (o *WhatsappGroupWebhookParticipant) GetRecipientUserId() string`

GetRecipientUserId returns the RecipientUserId field if non-nil, zero value otherwise.

### GetRecipientUserIdOk

`func (o *WhatsappGroupWebhookParticipant) GetRecipientUserIdOk() (*string, bool)`

GetRecipientUserIdOk returns a tuple with the RecipientUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientUserId

`func (o *WhatsappGroupWebhookParticipant) SetRecipientUserId(v string)`

SetRecipientUserId sets RecipientUserId field to given value.

### HasRecipientUserId

`func (o *WhatsappGroupWebhookParticipant) HasRecipientUserId() bool`

HasRecipientUserId returns a boolean if a field has been set.

### GetParentRecipientUserId

`func (o *WhatsappGroupWebhookParticipant) GetParentRecipientUserId() string`

GetParentRecipientUserId returns the ParentRecipientUserId field if non-nil, zero value otherwise.

### GetParentRecipientUserIdOk

`func (o *WhatsappGroupWebhookParticipant) GetParentRecipientUserIdOk() (*string, bool)`

GetParentRecipientUserIdOk returns a tuple with the ParentRecipientUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRecipientUserId

`func (o *WhatsappGroupWebhookParticipant) SetParentRecipientUserId(v string)`

SetParentRecipientUserId sets ParentRecipientUserId field to given value.

### HasParentRecipientUserId

`func (o *WhatsappGroupWebhookParticipant) HasParentRecipientUserId() bool`

HasParentRecipientUserId returns a boolean if a field has been set.

### GetCustomerProfile

`func (o *WhatsappGroupWebhookParticipant) GetCustomerProfile() WhatsappGroupCustomerProfile`

GetCustomerProfile returns the CustomerProfile field if non-nil, zero value otherwise.

### GetCustomerProfileOk

`func (o *WhatsappGroupWebhookParticipant) GetCustomerProfileOk() (*WhatsappGroupCustomerProfile, bool)`

GetCustomerProfileOk returns a tuple with the CustomerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerProfile

`func (o *WhatsappGroupWebhookParticipant) SetCustomerProfile(v WhatsappGroupCustomerProfile)`

SetCustomerProfile sets CustomerProfile field to given value.

### HasCustomerProfile

`func (o *WhatsappGroupWebhookParticipant) HasCustomerProfile() bool`

HasCustomerProfile returns a boolean if a field has been set.

### GetErrors

`func (o *WhatsappGroupWebhookParticipant) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *WhatsappGroupWebhookParticipant) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *WhatsappGroupWebhookParticipant) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *WhatsappGroupWebhookParticipant) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


