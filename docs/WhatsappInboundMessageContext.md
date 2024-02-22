# WhatsappInboundMessageContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Forwarded** | Pointer to **bool** | **Added to Webhooks if message was forwarded.**  Set to &#x60;true&#x60; if the received message has been forwarded. | [optional] 
**FrequentlyForwarded** | Pointer to **bool** | **Added to Webhooks if message has been frequently forwarded.**  Set to &#x60;true&#x60; if the received message has been forwarded more than five times. | [optional] 
**From** | Pointer to **string** | **Added to Webhooks if message is an inbound reply to a sent message.**  The WhatsApp ID (a phone number without the &#39;+&#39; prefix) of the sender of the sent message. | [optional] 
**Id** | Pointer to **string** | **Optional.**  The &#x60;wamid&#x60; for the sent message for an inbound reply. &#x60;wamid&#x60; is the original message ID on WhatsApp’s platform. | [optional] 
**ReferredProduct** | Pointer to [**WhatsappInboundMessageReferredProduct**](WhatsappInboundMessageReferredProduct.md) |  | [optional] 

## Methods

### NewWhatsappInboundMessageContext

`func NewWhatsappInboundMessageContext() *WhatsappInboundMessageContext`

NewWhatsappInboundMessageContext instantiates a new WhatsappInboundMessageContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappInboundMessageContextWithDefaults

`func NewWhatsappInboundMessageContextWithDefaults() *WhatsappInboundMessageContext`

NewWhatsappInboundMessageContextWithDefaults instantiates a new WhatsappInboundMessageContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForwarded

`func (o *WhatsappInboundMessageContext) GetForwarded() bool`

GetForwarded returns the Forwarded field if non-nil, zero value otherwise.

### GetForwardedOk

`func (o *WhatsappInboundMessageContext) GetForwardedOk() (*bool, bool)`

GetForwardedOk returns a tuple with the Forwarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwarded

`func (o *WhatsappInboundMessageContext) SetForwarded(v bool)`

SetForwarded sets Forwarded field to given value.

### HasForwarded

`func (o *WhatsappInboundMessageContext) HasForwarded() bool`

HasForwarded returns a boolean if a field has been set.

### GetFrequentlyForwarded

`func (o *WhatsappInboundMessageContext) GetFrequentlyForwarded() bool`

GetFrequentlyForwarded returns the FrequentlyForwarded field if non-nil, zero value otherwise.

### GetFrequentlyForwardedOk

`func (o *WhatsappInboundMessageContext) GetFrequentlyForwardedOk() (*bool, bool)`

GetFrequentlyForwardedOk returns a tuple with the FrequentlyForwarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequentlyForwarded

`func (o *WhatsappInboundMessageContext) SetFrequentlyForwarded(v bool)`

SetFrequentlyForwarded sets FrequentlyForwarded field to given value.

### HasFrequentlyForwarded

`func (o *WhatsappInboundMessageContext) HasFrequentlyForwarded() bool`

HasFrequentlyForwarded returns a boolean if a field has been set.

### GetFrom

`func (o *WhatsappInboundMessageContext) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *WhatsappInboundMessageContext) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *WhatsappInboundMessageContext) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *WhatsappInboundMessageContext) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetId

`func (o *WhatsappInboundMessageContext) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WhatsappInboundMessageContext) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WhatsappInboundMessageContext) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WhatsappInboundMessageContext) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReferredProduct

`func (o *WhatsappInboundMessageContext) GetReferredProduct() WhatsappInboundMessageReferredProduct`

GetReferredProduct returns the ReferredProduct field if non-nil, zero value otherwise.

### GetReferredProductOk

`func (o *WhatsappInboundMessageContext) GetReferredProductOk() (*WhatsappInboundMessageReferredProduct, bool)`

GetReferredProductOk returns a tuple with the ReferredProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredProduct

`func (o *WhatsappInboundMessageContext) SetReferredProduct(v WhatsappInboundMessageReferredProduct)`

SetReferredProduct sets ReferredProduct field to given value.

### HasReferredProduct

`func (o *WhatsappInboundMessageContext) HasReferredProduct() bool`

HasReferredProduct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
