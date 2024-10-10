# WhatsappPaymentTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Transaction ID. | 
**Type** | **string** | The payment type for this transactions. One of &#x60;billdesk&#x60;, &#x60;razorpay&#x60;, &#x60;payu&#x60;, or &#x60;zaakpay&#x60;. | 
**Status** | **string** | The status of the transaction. One of &#x60;pending&#x60;, &#x60;success&#x60; or &#x60;failed&#x60;. | 
**CreatedTimestamp** | **int64** | Time when transaction was created in epoch milliseconds. | 
**UpdatedTimestamp** | **int64** | Time when transaction was last updated in epoch milliseconds. | 
**Amount** | [**WhatsappMessageOrderAmount**](WhatsappMessageOrderAmount.md) |  | 
**Currency** | **string** | The currency for this payment. Currently the only supported value is &#x60;INR&#x60;. | 
**MethodType** | Pointer to **string** | Describes the type of payment method used by consumer to pay for the order. Can be one of &#x60;upi&#x60;, &#x60;card&#x60;, &#x60;wallet&#x60;, or &#x60;netbanking&#x60;. The payment method information might not be available for failed payments. | [optional] 
**Error** | Pointer to [**WhatsappPaymentTransactionError**](WhatsappPaymentTransactionError.md) |  | [optional] 

## Methods

### NewWhatsappPaymentTransaction

`func NewWhatsappPaymentTransaction(id string, type_ string, status string, createdTimestamp int64, updatedTimestamp int64, amount WhatsappMessageOrderAmount, currency string, ) *WhatsappPaymentTransaction`

NewWhatsappPaymentTransaction instantiates a new WhatsappPaymentTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappPaymentTransactionWithDefaults

`func NewWhatsappPaymentTransactionWithDefaults() *WhatsappPaymentTransaction`

NewWhatsappPaymentTransactionWithDefaults instantiates a new WhatsappPaymentTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WhatsappPaymentTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WhatsappPaymentTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WhatsappPaymentTransaction) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *WhatsappPaymentTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappPaymentTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappPaymentTransaction) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *WhatsappPaymentTransaction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WhatsappPaymentTransaction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WhatsappPaymentTransaction) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedTimestamp

`func (o *WhatsappPaymentTransaction) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *WhatsappPaymentTransaction) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *WhatsappPaymentTransaction) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *WhatsappPaymentTransaction) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *WhatsappPaymentTransaction) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *WhatsappPaymentTransaction) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### GetAmount

`func (o *WhatsappPaymentTransaction) GetAmount() WhatsappMessageOrderAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WhatsappPaymentTransaction) GetAmountOk() (*WhatsappMessageOrderAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WhatsappPaymentTransaction) SetAmount(v WhatsappMessageOrderAmount)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *WhatsappPaymentTransaction) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *WhatsappPaymentTransaction) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *WhatsappPaymentTransaction) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetMethodType

`func (o *WhatsappPaymentTransaction) GetMethodType() string`

GetMethodType returns the MethodType field if non-nil, zero value otherwise.

### GetMethodTypeOk

`func (o *WhatsappPaymentTransaction) GetMethodTypeOk() (*string, bool)`

GetMethodTypeOk returns a tuple with the MethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodType

`func (o *WhatsappPaymentTransaction) SetMethodType(v string)`

SetMethodType sets MethodType field to given value.

### HasMethodType

`func (o *WhatsappPaymentTransaction) HasMethodType() bool`

HasMethodType returns a boolean if a field has been set.

### GetError

`func (o *WhatsappPaymentTransaction) GetError() WhatsappPaymentTransactionError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WhatsappPaymentTransaction) GetErrorOk() (*WhatsappPaymentTransactionError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WhatsappPaymentTransaction) SetError(v WhatsappPaymentTransactionError)`

SetError sets Error field to given value.

### HasError

`func (o *WhatsappPaymentTransaction) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


