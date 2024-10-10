# WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Receipt** | Pointer to **string** | Receipt number that corresponds to this order, set for your internal reference.  Maximum length of 40 characters supported with minimum length greater than 0 characters. | [optional] 
**Notes** | Pointer to **map[string]string** | The object can be key value pairs with maximum 15 keys and each value limits to 256 characters. | [optional] 

## Methods

### NewWhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay

`func NewWhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay() *WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay`

NewWhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay instantiates a new WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageOrderPaymentSettingPaymentGatewayRazorpayWithDefaults

`func NewWhatsappMessageOrderPaymentSettingPaymentGatewayRazorpayWithDefaults() *WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay`

NewWhatsappMessageOrderPaymentSettingPaymentGatewayRazorpayWithDefaults instantiates a new WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceipt

`func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay) GetReceipt() string`

GetReceipt returns the Receipt field if non-nil, zero value otherwise.

### GetReceiptOk

`func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay) GetReceiptOk() (*string, bool)`

GetReceiptOk returns a tuple with the Receipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceipt

`func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay) SetReceipt(v string)`

SetReceipt sets Receipt field to given value.

### HasReceipt

`func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay) HasReceipt() bool`

HasReceipt returns a boolean if a field has been set.

### GetNotes

`func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay) GetNotes() map[string]string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay) GetNotesOk() (*map[string]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay) SetNotes(v map[string]string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


