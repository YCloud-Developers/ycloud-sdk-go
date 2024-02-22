# WhatsappInboundMessageSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Describes the system message event. Supported use cases are: - Phone number update: for when a user changes from an old number to a new number. - Identity update: for when a user identity has changed. | [optional] 
**NewWaId** | Pointer to **string** | **Added to Webhooks for phone number updates.**  New WhatsApp ID of the customer. | [optional] 
**Type** | Pointer to **string** | Supported types are: - &#x60;user_changed_number&#x60;: for a user changed number notification. - &#x60;user_identity_changed&#x60;: for user identity changed notification. | [optional] 
**User** | Pointer to **string** | **Added to Webhooks for identity updates.**  The new WhatsApp user ID of the customer. | [optional] 

## Methods

### NewWhatsappInboundMessageSystem

`func NewWhatsappInboundMessageSystem() *WhatsappInboundMessageSystem`

NewWhatsappInboundMessageSystem instantiates a new WhatsappInboundMessageSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappInboundMessageSystemWithDefaults

`func NewWhatsappInboundMessageSystemWithDefaults() *WhatsappInboundMessageSystem`

NewWhatsappInboundMessageSystemWithDefaults instantiates a new WhatsappInboundMessageSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *WhatsappInboundMessageSystem) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *WhatsappInboundMessageSystem) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *WhatsappInboundMessageSystem) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *WhatsappInboundMessageSystem) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetNewWaId

`func (o *WhatsappInboundMessageSystem) GetNewWaId() string`

GetNewWaId returns the NewWaId field if non-nil, zero value otherwise.

### GetNewWaIdOk

`func (o *WhatsappInboundMessageSystem) GetNewWaIdOk() (*string, bool)`

GetNewWaIdOk returns a tuple with the NewWaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewWaId

`func (o *WhatsappInboundMessageSystem) SetNewWaId(v string)`

SetNewWaId sets NewWaId field to given value.

### HasNewWaId

`func (o *WhatsappInboundMessageSystem) HasNewWaId() bool`

HasNewWaId returns a boolean if a field has been set.

### GetType

`func (o *WhatsappInboundMessageSystem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappInboundMessageSystem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappInboundMessageSystem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WhatsappInboundMessageSystem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUser

`func (o *WhatsappInboundMessageSystem) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *WhatsappInboundMessageSystem) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *WhatsappInboundMessageSystem) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *WhatsappInboundMessageSystem) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
