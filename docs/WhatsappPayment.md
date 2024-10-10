# WhatsappPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WabaId** | **string** | WhatsApp Business Account ID. | 
**ReferenceId** | **string** | Unique identifier for the payment provided by the business. It is case sensitive and cannot be an empty string and can only contain English letters, numbers, underscores, dashes, or dots, and should not exceed 35 characters. | 
**Status** | [**WhatsappPaymentStatus**](WhatsappPaymentStatus.md) |  | 
**Transactions** | Pointer to [**[]WhatsappPaymentTransaction**](WhatsappPaymentTransaction.md) | Contains the latest transaction attempt for this payment. | [optional] 

## Methods

### NewWhatsappPayment

`func NewWhatsappPayment(wabaId string, referenceId string, status WhatsappPaymentStatus, ) *WhatsappPayment`

NewWhatsappPayment instantiates a new WhatsappPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappPaymentWithDefaults

`func NewWhatsappPaymentWithDefaults() *WhatsappPayment`

NewWhatsappPaymentWithDefaults instantiates a new WhatsappPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWabaId

`func (o *WhatsappPayment) GetWabaId() string`

GetWabaId returns the WabaId field if non-nil, zero value otherwise.

### GetWabaIdOk

`func (o *WhatsappPayment) GetWabaIdOk() (*string, bool)`

GetWabaIdOk returns a tuple with the WabaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWabaId

`func (o *WhatsappPayment) SetWabaId(v string)`

SetWabaId sets WabaId field to given value.


### GetReferenceId

`func (o *WhatsappPayment) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *WhatsappPayment) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *WhatsappPayment) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetStatus

`func (o *WhatsappPayment) GetStatus() WhatsappPaymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WhatsappPayment) GetStatusOk() (*WhatsappPaymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WhatsappPayment) SetStatus(v WhatsappPaymentStatus)`

SetStatus sets Status field to given value.


### GetTransactions

`func (o *WhatsappPayment) GetTransactions() []WhatsappPaymentTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *WhatsappPayment) GetTransactionsOk() (*[]WhatsappPaymentTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *WhatsappPayment) SetTransactions(v []WhatsappPaymentTransaction)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *WhatsappPayment) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


