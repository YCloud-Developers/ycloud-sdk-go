# WhatsappPaymentTransactionError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Describes the payment failure reason that generated by payment gateway and Meta transmits this to partners. | 
**Reason** | **string** | Describes the payment failure reason in plain text that is generated by payment gateway and Meta transmits this to partners. | 

## Methods

### NewWhatsappPaymentTransactionError

`func NewWhatsappPaymentTransactionError(code string, reason string, ) *WhatsappPaymentTransactionError`

NewWhatsappPaymentTransactionError instantiates a new WhatsappPaymentTransactionError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappPaymentTransactionErrorWithDefaults

`func NewWhatsappPaymentTransactionErrorWithDefaults() *WhatsappPaymentTransactionError`

NewWhatsappPaymentTransactionErrorWithDefaults instantiates a new WhatsappPaymentTransactionError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *WhatsappPaymentTransactionError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *WhatsappPaymentTransactionError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *WhatsappPaymentTransactionError) SetCode(v string)`

SetCode sets Code field to given value.


### GetReason

`func (o *WhatsappPaymentTransactionError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *WhatsappPaymentTransactionError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *WhatsappPaymentTransactionError) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


