# WhatsappMessageInteractiveActionCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardIndex** | Pointer to **float32** | Card index. Unique index for each card (0-9). | [optional]
**Type** | Pointer to **string** | Must be \&quot;cta_url\&quot;. | [optional]
**Header** | Pointer to [**WhatsappMessageInteractiveActionCardHeader**](WhatsappMessageInteractiveActionCardHeader.md) |  | [optional]
**Body** | Pointer to [**WhatsappMessageInteractiveActionCardBody**](WhatsappMessageInteractiveActionCardBody.md) |  | [optional]
**Action** | Pointer to [**WhatsappMessageInteractiveActionCardAction**](WhatsappMessageInteractiveActionCardAction.md) |  | [optional]

## Methods

### NewWhatsappMessageInteractiveActionCard

`func NewWhatsappMessageInteractiveActionCard() *WhatsappMessageInteractiveActionCard`

NewWhatsappMessageInteractiveActionCard instantiates a new WhatsappMessageInteractiveActionCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageInteractiveActionCardWithDefaults

`func NewWhatsappMessageInteractiveActionCardWithDefaults() *WhatsappMessageInteractiveActionCard`

NewWhatsappMessageInteractiveActionCardWithDefaults instantiates a new WhatsappMessageInteractiveActionCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardIndex

`func (o *WhatsappMessageInteractiveActionCard) GetCardIndex() float32`

GetCardIndex returns the CardIndex field if non-nil, zero value otherwise.

### GetCardIndexOk

`func (o *WhatsappMessageInteractiveActionCard) GetCardIndexOk() (*float32, bool)`

GetCardIndexOk returns a tuple with the CardIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIndex

`func (o *WhatsappMessageInteractiveActionCard) SetCardIndex(v float32)`

SetCardIndex sets CardIndex field to given value.

### HasCardIndex

`func (o *WhatsappMessageInteractiveActionCard) HasCardIndex() bool`

HasCardIndex returns a boolean if a field has been set.

### GetType

`func (o *WhatsappMessageInteractiveActionCard) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappMessageInteractiveActionCard) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappMessageInteractiveActionCard) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WhatsappMessageInteractiveActionCard) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHeader

`func (o *WhatsappMessageInteractiveActionCard) GetHeader() WhatsappMessageInteractiveActionCardHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *WhatsappMessageInteractiveActionCard) GetHeaderOk() (*WhatsappMessageInteractiveActionCardHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *WhatsappMessageInteractiveActionCard) SetHeader(v WhatsappMessageInteractiveActionCardHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *WhatsappMessageInteractiveActionCard) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetBody

`func (o *WhatsappMessageInteractiveActionCard) GetBody() WhatsappMessageInteractiveActionCardBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *WhatsappMessageInteractiveActionCard) GetBodyOk() (*WhatsappMessageInteractiveActionCardBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *WhatsappMessageInteractiveActionCard) SetBody(v WhatsappMessageInteractiveActionCardBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *WhatsappMessageInteractiveActionCard) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetAction

`func (o *WhatsappMessageInteractiveActionCard) GetAction() WhatsappMessageInteractiveActionCardAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WhatsappMessageInteractiveActionCard) GetActionOk() (*WhatsappMessageInteractiveActionCardAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WhatsappMessageInteractiveActionCard) SetAction(v WhatsappMessageInteractiveActionCardAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *WhatsappMessageInteractiveActionCard) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


