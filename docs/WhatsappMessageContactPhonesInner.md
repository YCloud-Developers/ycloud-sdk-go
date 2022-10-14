# WhatsappMessageContactPhonesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | Pointer to **string** | Automatically populated with the &#x60;wa_id&#x60; value as a formatted phone number. | [optional] 
**Type** | Pointer to **string** | Standard Values are &#x60;CELL&#x60;, &#x60;MAIN&#x60;, &#x60;IPHONE&#x60;, &#x60;HOME&#x60;, and &#x60;WORK&#x60;. | [optional] 
**WaId** | Pointer to **string** | WhatsApp ID. | [optional] 

## Methods

### NewWhatsappMessageContactPhonesInner

`func NewWhatsappMessageContactPhonesInner() *WhatsappMessageContactPhonesInner`

NewWhatsappMessageContactPhonesInner instantiates a new WhatsappMessageContactPhonesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageContactPhonesInnerWithDefaults

`func NewWhatsappMessageContactPhonesInnerWithDefaults() *WhatsappMessageContactPhonesInner`

NewWhatsappMessageContactPhonesInnerWithDefaults instantiates a new WhatsappMessageContactPhonesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *WhatsappMessageContactPhonesInner) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *WhatsappMessageContactPhonesInner) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *WhatsappMessageContactPhonesInner) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *WhatsappMessageContactPhonesInner) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetType

`func (o *WhatsappMessageContactPhonesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappMessageContactPhonesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappMessageContactPhonesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WhatsappMessageContactPhonesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWaId

`func (o *WhatsappMessageContactPhonesInner) GetWaId() string`

GetWaId returns the WaId field if non-nil, zero value otherwise.

### GetWaIdOk

`func (o *WhatsappMessageContactPhonesInner) GetWaIdOk() (*string, bool)`

GetWaIdOk returns a tuple with the WaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaId

`func (o *WhatsappMessageContactPhonesInner) SetWaId(v string)`

SetWaId sets WaId field to given value.

### HasWaId

`func (o *WhatsappMessageContactPhonesInner) HasWaId() bool`

HasWaId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


