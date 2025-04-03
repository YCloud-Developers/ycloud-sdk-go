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

// WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay Additional info for Zaakpay. User-defined fields (extra) are used to store any information corresponding to a particular order. Each extra field has a maximum character limit of 180.
type WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay struct {
	Extra1 *string `json:"extra1,omitempty"`
	Extra2 *string `json:"extra2,omitempty"`
}

// NewWhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay instantiates a new WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay() *WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay {
	this := WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay{}
	return &this
}

// NewWhatsappMessageOrderPaymentSettingPaymentGatewayZaakpayWithDefaults instantiates a new WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappMessageOrderPaymentSettingPaymentGatewayZaakpayWithDefaults() *WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay {
	this := WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay{}
	return &this
}

// GetExtra1 returns the Extra1 field value if set, zero value otherwise.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay) GetExtra1() string {
	if o == nil || o.Extra1 == nil {
		var ret string
		return ret
	}
	return *o.Extra1
}

// GetExtra1Ok returns a tuple with the Extra1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay) GetExtra1Ok() (*string, bool) {
	if o == nil || o.Extra1 == nil {
		return nil, false
	}
	return o.Extra1, true
}

// HasExtra1 returns a boolean if a field has been set.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay) HasExtra1() bool {
	if o != nil && o.Extra1 != nil {
		return true
	}

	return false
}

// SetExtra1 gets a reference to the given string and assigns it to the Extra1 field.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay) SetExtra1(v string) {
	o.Extra1 = &v
}

// GetExtra2 returns the Extra2 field value if set, zero value otherwise.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay) GetExtra2() string {
	if o == nil || o.Extra2 == nil {
		var ret string
		return ret
	}
	return *o.Extra2
}

// GetExtra2Ok returns a tuple with the Extra2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay) GetExtra2Ok() (*string, bool) {
	if o == nil || o.Extra2 == nil {
		return nil, false
	}
	return o.Extra2, true
}

// HasExtra2 returns a boolean if a field has been set.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay) HasExtra2() bool {
	if o != nil && o.Extra2 != nil {
		return true
	}

	return false
}

// SetExtra2 gets a reference to the given string and assigns it to the Extra2 field.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay) SetExtra2(v string) {
	o.Extra2 = &v
}

func (o WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Extra1 != nil {
		toSerialize["extra1"] = o.Extra1
	}
	if o.Extra2 != nil {
		toSerialize["extra2"] = o.Extra2
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay struct {
	value *WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay
	isSet bool
}

func (v NullableWhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay) Get() *WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay {
	return v.value
}

func (v *NullableWhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay) Set(val *WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay(val *WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay) *NullableWhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay {
	return &NullableWhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay{value: val, isSet: true}
}

func (v NullableWhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


