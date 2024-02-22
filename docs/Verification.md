# Verification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the verification. | 
**Status** | Pointer to [**VerificationStatus**](VerificationStatus.md) |  | [optional] 
**To** | Pointer to **string** | Recipient of the verification. | [optional] 
**Channel** | Pointer to [**VerificationChannel**](VerificationChannel.md) |  | [optional] 
**SendTime** | Pointer to **time.Time** | The time at which this verification was sent, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional] 
**TotalPrice** | Pointer to **float64** | Total price of this verification. | [optional] 
**Currency** | Pointer to **string** | Price currency. [ISO 4217 currency code](https://en.wikipedia.org/wiki/ISO_4217). | [optional] 
**SmsFallbackEnabled** | Pointer to **bool** | Whether sms fallback is enabled or not. Applicable when &#x60;channel&#x60; is &#x60;whatsapp&#x60;. If enabled, we will try to send the verification code via sms when the WhatsApp message is failed. | [optional] 
**SmsFallback** | Pointer to [**VerificationFallback**](VerificationFallback.md) |  | [optional] 
**ExternalId** | Pointer to **string** | A unique string to reference the object. This can be an order number or similar, and can be used to reconcile the object with your internal systems. | [optional] 

## Methods

### NewVerification

`func NewVerification(id string, ) *Verification`

NewVerification instantiates a new Verification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationWithDefaults

`func NewVerificationWithDefaults() *Verification`

NewVerificationWithDefaults instantiates a new Verification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Verification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Verification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Verification) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *Verification) GetStatus() VerificationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Verification) GetStatusOk() (*VerificationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Verification) SetStatus(v VerificationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Verification) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTo

`func (o *Verification) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Verification) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Verification) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *Verification) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetChannel

`func (o *Verification) GetChannel() VerificationChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *Verification) GetChannelOk() (*VerificationChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *Verification) SetChannel(v VerificationChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *Verification) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetSendTime

`func (o *Verification) GetSendTime() time.Time`

GetSendTime returns the SendTime field if non-nil, zero value otherwise.

### GetSendTimeOk

`func (o *Verification) GetSendTimeOk() (*time.Time, bool)`

GetSendTimeOk returns a tuple with the SendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendTime

`func (o *Verification) SetSendTime(v time.Time)`

SetSendTime sets SendTime field to given value.

### HasSendTime

`func (o *Verification) HasSendTime() bool`

HasSendTime returns a boolean if a field has been set.

### GetTotalPrice

`func (o *Verification) GetTotalPrice() float64`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *Verification) GetTotalPriceOk() (*float64, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *Verification) SetTotalPrice(v float64)`

SetTotalPrice sets TotalPrice field to given value.

### HasTotalPrice

`func (o *Verification) HasTotalPrice() bool`

HasTotalPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *Verification) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Verification) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Verification) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Verification) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetSmsFallbackEnabled

`func (o *Verification) GetSmsFallbackEnabled() bool`

GetSmsFallbackEnabled returns the SmsFallbackEnabled field if non-nil, zero value otherwise.

### GetSmsFallbackEnabledOk

`func (o *Verification) GetSmsFallbackEnabledOk() (*bool, bool)`

GetSmsFallbackEnabledOk returns a tuple with the SmsFallbackEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsFallbackEnabled

`func (o *Verification) SetSmsFallbackEnabled(v bool)`

SetSmsFallbackEnabled sets SmsFallbackEnabled field to given value.

### HasSmsFallbackEnabled

`func (o *Verification) HasSmsFallbackEnabled() bool`

HasSmsFallbackEnabled returns a boolean if a field has been set.

### GetSmsFallback

`func (o *Verification) GetSmsFallback() VerificationFallback`

GetSmsFallback returns the SmsFallback field if non-nil, zero value otherwise.

### GetSmsFallbackOk

`func (o *Verification) GetSmsFallbackOk() (*VerificationFallback, bool)`

GetSmsFallbackOk returns a tuple with the SmsFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsFallback

`func (o *Verification) SetSmsFallback(v VerificationFallback)`

SetSmsFallback sets SmsFallback field to given value.

### HasSmsFallback

`func (o *Verification) HasSmsFallback() bool`

HasSmsFallback returns a boolean if a field has been set.

### GetExternalId

`func (o *Verification) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Verification) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Verification) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Verification) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
