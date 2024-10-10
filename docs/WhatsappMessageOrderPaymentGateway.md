# WhatsappMessageOrderPaymentGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Payment type. Must set this to &#x60;billdesk&#x60;, &#x60;razorpay&#x60;, &#x60;payu&#x60;, or &#x60;zaakpay&#x60;, if you have linked your BillDesk, Razorpay, PayU, or Zaakpay payment gateway to accept payments. | 
**ConfigurationName** | **string** | The name of the pre-configured payment configuration to use for this order and must not exceed 60 characters. This value must match with a payment configuration set up on the WhatsApp Business Manager. | 
**Billdesk** | Pointer to [**WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk**](WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk.md) |  | [optional] 
**Payu** | Pointer to [**WhatsappMessageOrderPaymentSettingPaymentGatewayPayu**](WhatsappMessageOrderPaymentSettingPaymentGatewayPayu.md) |  | [optional] 
**Razorpay** | Pointer to [**WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay**](WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay.md) |  | [optional] 
**Zaakpay** | Pointer to [**WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay**](WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay.md) |  | [optional] 

## Methods

### NewWhatsappMessageOrderPaymentGateway

`func NewWhatsappMessageOrderPaymentGateway(type_ string, configurationName string, ) *WhatsappMessageOrderPaymentGateway`

NewWhatsappMessageOrderPaymentGateway instantiates a new WhatsappMessageOrderPaymentGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageOrderPaymentGatewayWithDefaults

`func NewWhatsappMessageOrderPaymentGatewayWithDefaults() *WhatsappMessageOrderPaymentGateway`

NewWhatsappMessageOrderPaymentGatewayWithDefaults instantiates a new WhatsappMessageOrderPaymentGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WhatsappMessageOrderPaymentGateway) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappMessageOrderPaymentGateway) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappMessageOrderPaymentGateway) SetType(v string)`

SetType sets Type field to given value.


### GetConfigurationName

`func (o *WhatsappMessageOrderPaymentGateway) GetConfigurationName() string`

GetConfigurationName returns the ConfigurationName field if non-nil, zero value otherwise.

### GetConfigurationNameOk

`func (o *WhatsappMessageOrderPaymentGateway) GetConfigurationNameOk() (*string, bool)`

GetConfigurationNameOk returns a tuple with the ConfigurationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationName

`func (o *WhatsappMessageOrderPaymentGateway) SetConfigurationName(v string)`

SetConfigurationName sets ConfigurationName field to given value.


### GetBilldesk

`func (o *WhatsappMessageOrderPaymentGateway) GetBilldesk() WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk`

GetBilldesk returns the Billdesk field if non-nil, zero value otherwise.

### GetBilldeskOk

`func (o *WhatsappMessageOrderPaymentGateway) GetBilldeskOk() (*WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk, bool)`

GetBilldeskOk returns a tuple with the Billdesk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilldesk

`func (o *WhatsappMessageOrderPaymentGateway) SetBilldesk(v WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk)`

SetBilldesk sets Billdesk field to given value.

### HasBilldesk

`func (o *WhatsappMessageOrderPaymentGateway) HasBilldesk() bool`

HasBilldesk returns a boolean if a field has been set.

### GetPayu

`func (o *WhatsappMessageOrderPaymentGateway) GetPayu() WhatsappMessageOrderPaymentSettingPaymentGatewayPayu`

GetPayu returns the Payu field if non-nil, zero value otherwise.

### GetPayuOk

`func (o *WhatsappMessageOrderPaymentGateway) GetPayuOk() (*WhatsappMessageOrderPaymentSettingPaymentGatewayPayu, bool)`

GetPayuOk returns a tuple with the Payu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayu

`func (o *WhatsappMessageOrderPaymentGateway) SetPayu(v WhatsappMessageOrderPaymentSettingPaymentGatewayPayu)`

SetPayu sets Payu field to given value.

### HasPayu

`func (o *WhatsappMessageOrderPaymentGateway) HasPayu() bool`

HasPayu returns a boolean if a field has been set.

### GetRazorpay

`func (o *WhatsappMessageOrderPaymentGateway) GetRazorpay() WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay`

GetRazorpay returns the Razorpay field if non-nil, zero value otherwise.

### GetRazorpayOk

`func (o *WhatsappMessageOrderPaymentGateway) GetRazorpayOk() (*WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay, bool)`

GetRazorpayOk returns a tuple with the Razorpay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRazorpay

`func (o *WhatsappMessageOrderPaymentGateway) SetRazorpay(v WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay)`

SetRazorpay sets Razorpay field to given value.

### HasRazorpay

`func (o *WhatsappMessageOrderPaymentGateway) HasRazorpay() bool`

HasRazorpay returns a boolean if a field has been set.

### GetZaakpay

`func (o *WhatsappMessageOrderPaymentGateway) GetZaakpay() WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay`

GetZaakpay returns the Zaakpay field if non-nil, zero value otherwise.

### GetZaakpayOk

`func (o *WhatsappMessageOrderPaymentGateway) GetZaakpayOk() (*WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay, bool)`

GetZaakpayOk returns a tuple with the Zaakpay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZaakpay

`func (o *WhatsappMessageOrderPaymentGateway) SetZaakpay(v WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay)`

SetZaakpay sets Zaakpay field to given value.

### HasZaakpay

`func (o *WhatsappMessageOrderPaymentGateway) HasZaakpay() bool`

HasZaakpay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


