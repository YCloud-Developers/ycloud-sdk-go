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

// WhatsappInboundMessageOrderProductItem struct for WhatsappInboundMessageOrderProductItem
type WhatsappInboundMessageOrderProductItem struct {
	// The product SKU identifier.
	ProductRetailerId *string `json:"product_retailer_id,omitempty"`
	// Number of item.
	Quantity *int32 `json:"quantity,omitempty"`
	// Unitary price of item.
	ItemPrice *float64 `json:"item_price,omitempty"`
	// Price currency. [ISO 4217 currency code](https://en.wikipedia.org/wiki/ISO_4217).
	Currency *string `json:"currency,omitempty"`
}

// NewWhatsappInboundMessageOrderProductItem instantiates a new WhatsappInboundMessageOrderProductItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappInboundMessageOrderProductItem() *WhatsappInboundMessageOrderProductItem {
	this := WhatsappInboundMessageOrderProductItem{}
	return &this
}

// NewWhatsappInboundMessageOrderProductItemWithDefaults instantiates a new WhatsappInboundMessageOrderProductItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappInboundMessageOrderProductItemWithDefaults() *WhatsappInboundMessageOrderProductItem {
	this := WhatsappInboundMessageOrderProductItem{}
	return &this
}

// GetProductRetailerId returns the ProductRetailerId field value if set, zero value otherwise.
func (o *WhatsappInboundMessageOrderProductItem) GetProductRetailerId() string {
	if o == nil || o.ProductRetailerId == nil {
		var ret string
		return ret
	}
	return *o.ProductRetailerId
}

// GetProductRetailerIdOk returns a tuple with the ProductRetailerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappInboundMessageOrderProductItem) GetProductRetailerIdOk() (*string, bool) {
	if o == nil || o.ProductRetailerId == nil {
		return nil, false
	}
	return o.ProductRetailerId, true
}

// HasProductRetailerId returns a boolean if a field has been set.
func (o *WhatsappInboundMessageOrderProductItem) HasProductRetailerId() bool {
	if o != nil && o.ProductRetailerId != nil {
		return true
	}

	return false
}

// SetProductRetailerId gets a reference to the given string and assigns it to the ProductRetailerId field.
func (o *WhatsappInboundMessageOrderProductItem) SetProductRetailerId(v string) {
	o.ProductRetailerId = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *WhatsappInboundMessageOrderProductItem) GetQuantity() int32 {
	if o == nil || o.Quantity == nil {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappInboundMessageOrderProductItem) GetQuantityOk() (*int32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *WhatsappInboundMessageOrderProductItem) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *WhatsappInboundMessageOrderProductItem) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetItemPrice returns the ItemPrice field value if set, zero value otherwise.
func (o *WhatsappInboundMessageOrderProductItem) GetItemPrice() float64 {
	if o == nil || o.ItemPrice == nil {
		var ret float64
		return ret
	}
	return *o.ItemPrice
}

// GetItemPriceOk returns a tuple with the ItemPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappInboundMessageOrderProductItem) GetItemPriceOk() (*float64, bool) {
	if o == nil || o.ItemPrice == nil {
		return nil, false
	}
	return o.ItemPrice, true
}

// HasItemPrice returns a boolean if a field has been set.
func (o *WhatsappInboundMessageOrderProductItem) HasItemPrice() bool {
	if o != nil && o.ItemPrice != nil {
		return true
	}

	return false
}

// SetItemPrice gets a reference to the given float64 and assigns it to the ItemPrice field.
func (o *WhatsappInboundMessageOrderProductItem) SetItemPrice(v float64) {
	o.ItemPrice = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *WhatsappInboundMessageOrderProductItem) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappInboundMessageOrderProductItem) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *WhatsappInboundMessageOrderProductItem) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *WhatsappInboundMessageOrderProductItem) SetCurrency(v string) {
	o.Currency = &v
}

func (o WhatsappInboundMessageOrderProductItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProductRetailerId != nil {
		toSerialize["product_retailer_id"] = o.ProductRetailerId
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.ItemPrice != nil {
		toSerialize["item_price"] = o.ItemPrice
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappInboundMessageOrderProductItem struct {
	value *WhatsappInboundMessageOrderProductItem
	isSet bool
}

func (v NullableWhatsappInboundMessageOrderProductItem) Get() *WhatsappInboundMessageOrderProductItem {
	return v.value
}

func (v *NullableWhatsappInboundMessageOrderProductItem) Set(val *WhatsappInboundMessageOrderProductItem) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappInboundMessageOrderProductItem) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappInboundMessageOrderProductItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappInboundMessageOrderProductItem(val *WhatsappInboundMessageOrderProductItem) *NullableWhatsappInboundMessageOrderProductItem {
	return &NullableWhatsappInboundMessageOrderProductItem{value: val, isSet: true}
}

func (v NullableWhatsappInboundMessageOrderProductItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappInboundMessageOrderProductItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


