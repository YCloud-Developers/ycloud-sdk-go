# WhatsappMessageOrderPaymentSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Must be set to &#x60;payment_gateway&#x60;. | 
**PaymentGateway** | [**WhatsappMessageOrderPaymentGateway**](WhatsappMessageOrderPaymentGateway.md) |  | 

## Methods

### NewWhatsappMessageOrderPaymentSetting

`func NewWhatsappMessageOrderPaymentSetting(type_ string, paymentGateway WhatsappMessageOrderPaymentGateway, ) *WhatsappMessageOrderPaymentSetting`

NewWhatsappMessageOrderPaymentSetting instantiates a new WhatsappMessageOrderPaymentSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageOrderPaymentSettingWithDefaults

`func NewWhatsappMessageOrderPaymentSettingWithDefaults() *WhatsappMessageOrderPaymentSetting`

NewWhatsappMessageOrderPaymentSettingWithDefaults instantiates a new WhatsappMessageOrderPaymentSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WhatsappMessageOrderPaymentSetting) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappMessageOrderPaymentSetting) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappMessageOrderPaymentSetting) SetType(v string)`

SetType sets Type field to given value.


### GetPaymentGateway

`func (o *WhatsappMessageOrderPaymentSetting) GetPaymentGateway() WhatsappMessageOrderPaymentGateway`

GetPaymentGateway returns the PaymentGateway field if non-nil, zero value otherwise.

### GetPaymentGatewayOk

`func (o *WhatsappMessageOrderPaymentSetting) GetPaymentGatewayOk() (*WhatsappMessageOrderPaymentGateway, bool)`

GetPaymentGatewayOk returns a tuple with the PaymentGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGateway

`func (o *WhatsappMessageOrderPaymentSetting) SetPaymentGateway(v WhatsappMessageOrderPaymentGateway)`

SetPaymentGateway sets PaymentGateway field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


