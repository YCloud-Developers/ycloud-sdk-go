/*
YCloud API

The [YCloud](https://ycloud.com) API is organized around [REST](https://en.wikipedia.org/wiki/Representational_state_transfer). Our API is designed to have predictable, resource-oriented URLs, return [JSON](https://www.json.org) responses, and use standard HTTP response codes and verbs.

API version: v2
Contact: service@ycloud.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ycloud

import (
	"encoding/json"
)

// WhatsappPaymentTransaction Represents a transaction attempt for a payment.
type WhatsappPaymentTransaction struct {
	// Transaction ID.
	Id string `json:"id"`
	// The payment type for this transactions. One of `billdesk`, `razorpay`, `payu`, or `zaakpay`.
	Type string `json:"type"`
	// The status of the transaction. One of `pending`, `success` or `failed`.
	Status string `json:"status"`
	// Time when transaction was created in epoch milliseconds.
	CreatedTimestamp int64 `json:"createdTimestamp"`
	// Time when transaction was last updated in epoch milliseconds.
	UpdatedTimestamp int64 `json:"updatedTimestamp"`
	Amount WhatsappMessageOrderAmount `json:"amount"`
	// The currency for this payment. Currently the only supported value is `INR`.
	Currency string `json:"currency"`
	// Describes the type of payment method used by consumer to pay for the order. Can be one of `upi`, `card`, `wallet`, or `netbanking`. The payment method information might not be available for failed payments.
	MethodType *string `json:"methodType,omitempty"`
	Error *WhatsappPaymentTransactionError `json:"error,omitempty"`
}

// NewWhatsappPaymentTransaction instantiates a new WhatsappPaymentTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappPaymentTransaction(id string, type_ string, status string, createdTimestamp int64, updatedTimestamp int64, amount WhatsappMessageOrderAmount, currency string) *WhatsappPaymentTransaction {
	this := WhatsappPaymentTransaction{}
	this.Id = id
	this.Type = type_
	this.Status = status
	this.CreatedTimestamp = createdTimestamp
	this.UpdatedTimestamp = updatedTimestamp
	this.Amount = amount
	this.Currency = currency
	return &this
}

// NewWhatsappPaymentTransactionWithDefaults instantiates a new WhatsappPaymentTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappPaymentTransactionWithDefaults() *WhatsappPaymentTransaction {
	this := WhatsappPaymentTransaction{}
	return &this
}

// GetId returns the Id field value
func (o *WhatsappPaymentTransaction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WhatsappPaymentTransaction) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WhatsappPaymentTransaction) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *WhatsappPaymentTransaction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WhatsappPaymentTransaction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WhatsappPaymentTransaction) SetType(v string) {
	o.Type = v
}

// GetStatus returns the Status field value
func (o *WhatsappPaymentTransaction) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WhatsappPaymentTransaction) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *WhatsappPaymentTransaction) SetStatus(v string) {
	o.Status = v
}

// GetCreatedTimestamp returns the CreatedTimestamp field value
func (o *WhatsappPaymentTransaction) GetCreatedTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *WhatsappPaymentTransaction) GetCreatedTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTimestamp, true
}

// SetCreatedTimestamp sets field value
func (o *WhatsappPaymentTransaction) SetCreatedTimestamp(v int64) {
	o.CreatedTimestamp = v
}

// GetUpdatedTimestamp returns the UpdatedTimestamp field value
func (o *WhatsappPaymentTransaction) GetUpdatedTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedTimestamp
}

// GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field value
// and a boolean to check if the value has been set.
func (o *WhatsappPaymentTransaction) GetUpdatedTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedTimestamp, true
}

// SetUpdatedTimestamp sets field value
func (o *WhatsappPaymentTransaction) SetUpdatedTimestamp(v int64) {
	o.UpdatedTimestamp = v
}

// GetAmount returns the Amount field value
func (o *WhatsappPaymentTransaction) GetAmount() WhatsappMessageOrderAmount {
	if o == nil {
		var ret WhatsappMessageOrderAmount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *WhatsappPaymentTransaction) GetAmountOk() (*WhatsappMessageOrderAmount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *WhatsappPaymentTransaction) SetAmount(v WhatsappMessageOrderAmount) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *WhatsappPaymentTransaction) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *WhatsappPaymentTransaction) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *WhatsappPaymentTransaction) SetCurrency(v string) {
	o.Currency = v
}

// GetMethodType returns the MethodType field value if set, zero value otherwise.
func (o *WhatsappPaymentTransaction) GetMethodType() string {
	if o == nil || o.MethodType == nil {
		var ret string
		return ret
	}
	return *o.MethodType
}

// GetMethodTypeOk returns a tuple with the MethodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappPaymentTransaction) GetMethodTypeOk() (*string, bool) {
	if o == nil || o.MethodType == nil {
		return nil, false
	}
	return o.MethodType, true
}

// HasMethodType returns a boolean if a field has been set.
func (o *WhatsappPaymentTransaction) HasMethodType() bool {
	if o != nil && o.MethodType != nil {
		return true
	}

	return false
}

// SetMethodType gets a reference to the given string and assigns it to the MethodType field.
func (o *WhatsappPaymentTransaction) SetMethodType(v string) {
	o.MethodType = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *WhatsappPaymentTransaction) GetError() WhatsappPaymentTransactionError {
	if o == nil || o.Error == nil {
		var ret WhatsappPaymentTransactionError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappPaymentTransaction) GetErrorOk() (*WhatsappPaymentTransactionError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *WhatsappPaymentTransaction) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given WhatsappPaymentTransactionError and assigns it to the Error field.
func (o *WhatsappPaymentTransaction) SetError(v WhatsappPaymentTransactionError) {
	o.Error = &v
}

func (o WhatsappPaymentTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["createdTimestamp"] = o.CreatedTimestamp
	}
	if true {
		toSerialize["updatedTimestamp"] = o.UpdatedTimestamp
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if o.MethodType != nil {
		toSerialize["methodType"] = o.MethodType
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappPaymentTransaction struct {
	value *WhatsappPaymentTransaction
	isSet bool
}

func (v NullableWhatsappPaymentTransaction) Get() *WhatsappPaymentTransaction {
	return v.value
}

func (v *NullableWhatsappPaymentTransaction) Set(val *WhatsappPaymentTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappPaymentTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappPaymentTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappPaymentTransaction(val *WhatsappPaymentTransaction) *NullableWhatsappPaymentTransaction {
	return &NullableWhatsappPaymentTransaction{value: val, isSet: true}
}

func (v NullableWhatsappPaymentTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappPaymentTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


