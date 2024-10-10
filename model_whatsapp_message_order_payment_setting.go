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

// WhatsappMessageOrderPaymentSetting Payment settings for the order.
type WhatsappMessageOrderPaymentSetting struct {
	// Must be set to `payment_gateway`.
	Type           string                             `json:"type"`
	PaymentGateway WhatsappMessageOrderPaymentGateway `json:"payment_gateway"`
}

// NewWhatsappMessageOrderPaymentSetting instantiates a new WhatsappMessageOrderPaymentSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappMessageOrderPaymentSetting(type_ string, paymentGateway WhatsappMessageOrderPaymentGateway) *WhatsappMessageOrderPaymentSetting {
	this := WhatsappMessageOrderPaymentSetting{}
	this.Type = type_
	this.PaymentGateway = paymentGateway
	return &this
}

// NewWhatsappMessageOrderPaymentSettingWithDefaults instantiates a new WhatsappMessageOrderPaymentSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappMessageOrderPaymentSettingWithDefaults() *WhatsappMessageOrderPaymentSetting {
	this := WhatsappMessageOrderPaymentSetting{}
	return &this
}

// GetType returns the Type field value
func (o *WhatsappMessageOrderPaymentSetting) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderPaymentSetting) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WhatsappMessageOrderPaymentSetting) SetType(v string) {
	o.Type = v
}

// GetPaymentGateway returns the PaymentGateway field value
func (o *WhatsappMessageOrderPaymentSetting) GetPaymentGateway() WhatsappMessageOrderPaymentGateway {
	if o == nil {
		var ret WhatsappMessageOrderPaymentGateway
		return ret
	}

	return o.PaymentGateway
}

// GetPaymentGatewayOk returns a tuple with the PaymentGateway field value
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderPaymentSetting) GetPaymentGatewayOk() (*WhatsappMessageOrderPaymentGateway, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentGateway, true
}

// SetPaymentGateway sets field value
func (o *WhatsappMessageOrderPaymentSetting) SetPaymentGateway(v WhatsappMessageOrderPaymentGateway) {
	o.PaymentGateway = v
}

func (o WhatsappMessageOrderPaymentSetting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["payment_gateway"] = o.PaymentGateway
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappMessageOrderPaymentSetting struct {
	value *WhatsappMessageOrderPaymentSetting
	isSet bool
}

func (v NullableWhatsappMessageOrderPaymentSetting) Get() *WhatsappMessageOrderPaymentSetting {
	return v.value
}

func (v *NullableWhatsappMessageOrderPaymentSetting) Set(val *WhatsappMessageOrderPaymentSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappMessageOrderPaymentSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappMessageOrderPaymentSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappMessageOrderPaymentSetting(val *WhatsappMessageOrderPaymentSetting) *NullableWhatsappMessageOrderPaymentSetting {
	return &NullableWhatsappMessageOrderPaymentSetting{value: val, isSet: true}
}

func (v NullableWhatsappMessageOrderPaymentSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappMessageOrderPaymentSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
