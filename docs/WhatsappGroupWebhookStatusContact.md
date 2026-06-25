# WhatsappGroupWebhookStatusContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerProfile** | Pointer to [**WhatsappGroupCustomerProfile**](WhatsappGroupCustomerProfile.md) |  | [optional] 
**WaId** | Pointer to **string** | WhatsApp user ID. | [optional] 
**RecipientUserId** | Pointer to **string** | Business-scoped user ID. | [optional] 
**ParentRecipientUserId** | Pointer to **string** | Parent business-scoped user ID. | [optional] 

## Methods

### NewWhatsappGroupWebhookStatusContact

`func NewWhatsappGroupWebhookStatusContact() *WhatsappGroupWebhookStatusContact`

NewWhatsappGroupWebhookStatusContact instantiates a new WhatsappGroupWebhookStatusContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappGroupWebhookStatusContactWithDefaults

`func NewWhatsappGroupWebhookStatusContactWithDefaults() *WhatsappGroupWebhookStatusContact`

NewWhatsappGroupWebhookStatusContactWithDefaults instantiates a new WhatsappGroupWebhookStatusContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerProfile

`func (o *WhatsappGroupWebhookStatusContact) GetCustomerProfile() WhatsappGroupCustomerProfile`

GetCustomerProfile returns the CustomerProfile field if non-nil, zero value otherwise.

### GetCustomerProfileOk

`func (o *WhatsappGroupWebhookStatusContact) GetCustomerProfileOk() (*WhatsappGroupCustomerProfile, bool)`

GetCustomerProfileOk returns a tuple with the CustomerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerProfile

`func (o *WhatsappGroupWebhookStatusContact) SetCustomerProfile(v WhatsappGroupCustomerProfile)`

SetCustomerProfile sets CustomerProfile field to given value.

### HasCustomerProfile

`func (o *WhatsappGroupWebhookStatusContact) HasCustomerProfile() bool`

HasCustomerProfile returns a boolean if a field has been set.

### GetWaId

`func (o *WhatsappGroupWebhookStatusContact) GetWaId() string`

GetWaId returns the WaId field if non-nil, zero value otherwise.

### GetWaIdOk

`func (o *WhatsappGroupWebhookStatusContact) GetWaIdOk() (*string, bool)`

GetWaIdOk returns a tuple with the WaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaId

`func (o *WhatsappGroupWebhookStatusContact) SetWaId(v string)`

SetWaId sets WaId field to given value.

### HasWaId

`func (o *WhatsappGroupWebhookStatusContact) HasWaId() bool`

HasWaId returns a boolean if a field has been set.

### GetRecipientUserId

`func (o *WhatsappGroupWebhookStatusContact) GetRecipientUserId() string`

GetRecipientUserId returns the RecipientUserId field if non-nil, zero value otherwise.

### GetRecipientUserIdOk

`func (o *WhatsappGroupWebhookStatusContact) GetRecipientUserIdOk() (*string, bool)`

GetRecipientUserIdOk returns a tuple with the RecipientUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientUserId

`func (o *WhatsappGroupWebhookStatusContact) SetRecipientUserId(v string)`

SetRecipientUserId sets RecipientUserId field to given value.

### HasRecipientUserId

`func (o *WhatsappGroupWebhookStatusContact) HasRecipientUserId() bool`

HasRecipientUserId returns a boolean if a field has been set.

### GetParentRecipientUserId

`func (o *WhatsappGroupWebhookStatusContact) GetParentRecipientUserId() string`

GetParentRecipientUserId returns the ParentRecipientUserId field if non-nil, zero value otherwise.

### GetParentRecipientUserIdOk

`func (o *WhatsappGroupWebhookStatusContact) GetParentRecipientUserIdOk() (*string, bool)`

GetParentRecipientUserIdOk returns a tuple with the ParentRecipientUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRecipientUserId

`func (o *WhatsappGroupWebhookStatusContact) SetParentRecipientUserId(v string)`

SetParentRecipientUserId sets ParentRecipientUserId field to given value.

### HasParentRecipientUserId

`func (o *WhatsappGroupWebhookStatusContact) HasParentRecipientUserId() bool`

HasParentRecipientUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


