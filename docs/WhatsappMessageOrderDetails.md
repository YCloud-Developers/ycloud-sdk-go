# WhatsappMessageOrderDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | The currency for this order. Currently the only supported value is &#x60;INR&#x60;. | 
**Order** | [**WhatsappMessageOrderInfo**](WhatsappMessageOrderInfo.md) |  | 
**ReferenceId** | **string** | Unique identifier for the order provided by the business. It is case sensitive and cannot be an empty string and can only contain English letters, numbers, underscores, dashes, or dots, and should not exceed 35 characters.  The &#x60;reference_id&#x60; must be unique for each order_details message for a given business. If there is a need to send multiple order_details messages for the same order, it is recommended to include a sequence number in the reference_id (for example, \&quot;BM345A-12\&quot;) to ensure reference_id uniqueness. | 
**TotalAmount** | [**WhatsappMessageOrderAmount**](WhatsappMessageOrderAmount.md) |  | 
**Type** | **string** | The type of goods being paid for in this order. Current supported options are &#x60;digital-goods&#x60; and &#x60;physical-goods&#x60;. | 
**PaymentSettings** | [**[]WhatsappMessageOrderPaymentSetting**](WhatsappMessageOrderPaymentSetting.md) | Payment settings for the order. | 

## Methods

### NewWhatsappMessageOrderDetails

`func NewWhatsappMessageOrderDetails(currency string, order WhatsappMessageOrderInfo, referenceId string, totalAmount WhatsappMessageOrderAmount, type_ string, paymentSettings []WhatsappMessageOrderPaymentSetting, ) *WhatsappMessageOrderDetails`

NewWhatsappMessageOrderDetails instantiates a new WhatsappMessageOrderDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageOrderDetailsWithDefaults

`func NewWhatsappMessageOrderDetailsWithDefaults() *WhatsappMessageOrderDetails`

NewWhatsappMessageOrderDetailsWithDefaults instantiates a new WhatsappMessageOrderDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *WhatsappMessageOrderDetails) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *WhatsappMessageOrderDetails) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *WhatsappMessageOrderDetails) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetOrder

`func (o *WhatsappMessageOrderDetails) GetOrder() WhatsappMessageOrderInfo`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *WhatsappMessageOrderDetails) GetOrderOk() (*WhatsappMessageOrderInfo, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *WhatsappMessageOrderDetails) SetOrder(v WhatsappMessageOrderInfo)`

SetOrder sets Order field to given value.


### GetReferenceId

`func (o *WhatsappMessageOrderDetails) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *WhatsappMessageOrderDetails) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *WhatsappMessageOrderDetails) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetTotalAmount

`func (o *WhatsappMessageOrderDetails) GetTotalAmount() WhatsappMessageOrderAmount`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *WhatsappMessageOrderDetails) GetTotalAmountOk() (*WhatsappMessageOrderAmount, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *WhatsappMessageOrderDetails) SetTotalAmount(v WhatsappMessageOrderAmount)`

SetTotalAmount sets TotalAmount field to given value.


### GetType

`func (o *WhatsappMessageOrderDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappMessageOrderDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappMessageOrderDetails) SetType(v string)`

SetType sets Type field to given value.


### GetPaymentSettings

`func (o *WhatsappMessageOrderDetails) GetPaymentSettings() []WhatsappMessageOrderPaymentSetting`

GetPaymentSettings returns the PaymentSettings field if non-nil, zero value otherwise.

### GetPaymentSettingsOk

`func (o *WhatsappMessageOrderDetails) GetPaymentSettingsOk() (*[]WhatsappMessageOrderPaymentSetting, bool)`

GetPaymentSettingsOk returns a tuple with the PaymentSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSettings

`func (o *WhatsappMessageOrderDetails) SetPaymentSettings(v []WhatsappMessageOrderPaymentSetting)`

SetPaymentSettings sets PaymentSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


