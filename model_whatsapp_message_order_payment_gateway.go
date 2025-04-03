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

// WhatsappMessageOrderPaymentGateway An object that describes payment account information.
type WhatsappMessageOrderPaymentGateway struct {
	// Payment type. Must set this to `billdesk`, `razorpay`, `payu`, or `zaakpay`, if you have linked your BillDesk, Razorpay, PayU, or Zaakpay payment gateway to accept payments.
	Type string `json:"type"`
	// The name of the pre-configured payment configuration to use for this order and must not exceed 60 characters. This value must match with a payment configuration set up on the WhatsApp Business Manager.
	ConfigurationName string `json:"configuration_name"`
	Billdesk *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk `json:"billdesk,omitempty"`
	Payu *WhatsappMessageOrderPaymentSettingPaymentGatewayPayu `json:"payu,omitempty"`
	Razorpay *WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay `json:"razorpay,omitempty"`
	Zaakpay *WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay `json:"zaakpay,omitempty"`
}

// NewWhatsappMessageOrderPaymentGateway instantiates a new WhatsappMessageOrderPaymentGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappMessageOrderPaymentGateway(type_ string, configurationName string) *WhatsappMessageOrderPaymentGateway {
	this := WhatsappMessageOrderPaymentGateway{}
	this.Type = type_
	this.ConfigurationName = configurationName
	return &this
}

// NewWhatsappMessageOrderPaymentGatewayWithDefaults instantiates a new WhatsappMessageOrderPaymentGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappMessageOrderPaymentGatewayWithDefaults() *WhatsappMessageOrderPaymentGateway {
	this := WhatsappMessageOrderPaymentGateway{}
	return &this
}

// GetType returns the Type field value
func (o *WhatsappMessageOrderPaymentGateway) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderPaymentGateway) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WhatsappMessageOrderPaymentGateway) SetType(v string) {
	o.Type = v
}

// GetConfigurationName returns the ConfigurationName field value
func (o *WhatsappMessageOrderPaymentGateway) GetConfigurationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigurationName
}

// GetConfigurationNameOk returns a tuple with the ConfigurationName field value
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderPaymentGateway) GetConfigurationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigurationName, true
}

// SetConfigurationName sets field value
func (o *WhatsappMessageOrderPaymentGateway) SetConfigurationName(v string) {
	o.ConfigurationName = v
}

// GetBilldesk returns the Billdesk field value if set, zero value otherwise.
func (o *WhatsappMessageOrderPaymentGateway) GetBilldesk() WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk {
	if o == nil || o.Billdesk == nil {
		var ret WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk
		return ret
	}
	return *o.Billdesk
}

// GetBilldeskOk returns a tuple with the Billdesk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderPaymentGateway) GetBilldeskOk() (*WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk, bool) {
	if o == nil || o.Billdesk == nil {
		return nil, false
	}
	return o.Billdesk, true
}

// HasBilldesk returns a boolean if a field has been set.
func (o *WhatsappMessageOrderPaymentGateway) HasBilldesk() bool {
	if o != nil && o.Billdesk != nil {
		return true
	}

	return false
}

// SetBilldesk gets a reference to the given WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk and assigns it to the Billdesk field.
func (o *WhatsappMessageOrderPaymentGateway) SetBilldesk(v WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) {
	o.Billdesk = &v
}

// GetPayu returns the Payu field value if set, zero value otherwise.
func (o *WhatsappMessageOrderPaymentGateway) GetPayu() WhatsappMessageOrderPaymentSettingPaymentGatewayPayu {
	if o == nil || o.Payu == nil {
		var ret WhatsappMessageOrderPaymentSettingPaymentGatewayPayu
		return ret
	}
	return *o.Payu
}

// GetPayuOk returns a tuple with the Payu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderPaymentGateway) GetPayuOk() (*WhatsappMessageOrderPaymentSettingPaymentGatewayPayu, bool) {
	if o == nil || o.Payu == nil {
		return nil, false
	}
	return o.Payu, true
}

// HasPayu returns a boolean if a field has been set.
func (o *WhatsappMessageOrderPaymentGateway) HasPayu() bool {
	if o != nil && o.Payu != nil {
		return true
	}

	return false
}

// SetPayu gets a reference to the given WhatsappMessageOrderPaymentSettingPaymentGatewayPayu and assigns it to the Payu field.
func (o *WhatsappMessageOrderPaymentGateway) SetPayu(v WhatsappMessageOrderPaymentSettingPaymentGatewayPayu) {
	o.Payu = &v
}

// GetRazorpay returns the Razorpay field value if set, zero value otherwise.
func (o *WhatsappMessageOrderPaymentGateway) GetRazorpay() WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay {
	if o == nil || o.Razorpay == nil {
		var ret WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay
		return ret
	}
	return *o.Razorpay
}

// GetRazorpayOk returns a tuple with the Razorpay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderPaymentGateway) GetRazorpayOk() (*WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay, bool) {
	if o == nil || o.Razorpay == nil {
		return nil, false
	}
	return o.Razorpay, true
}

// HasRazorpay returns a boolean if a field has been set.
func (o *WhatsappMessageOrderPaymentGateway) HasRazorpay() bool {
	if o != nil && o.Razorpay != nil {
		return true
	}

	return false
}

// SetRazorpay gets a reference to the given WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay and assigns it to the Razorpay field.
func (o *WhatsappMessageOrderPaymentGateway) SetRazorpay(v WhatsappMessageOrderPaymentSettingPaymentGatewayRazorpay) {
	o.Razorpay = &v
}

// GetZaakpay returns the Zaakpay field value if set, zero value otherwise.
func (o *WhatsappMessageOrderPaymentGateway) GetZaakpay() WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay {
	if o == nil || o.Zaakpay == nil {
		var ret WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay
		return ret
	}
	return *o.Zaakpay
}

// GetZaakpayOk returns a tuple with the Zaakpay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderPaymentGateway) GetZaakpayOk() (*WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay, bool) {
	if o == nil || o.Zaakpay == nil {
		return nil, false
	}
	return o.Zaakpay, true
}

// HasZaakpay returns a boolean if a field has been set.
func (o *WhatsappMessageOrderPaymentGateway) HasZaakpay() bool {
	if o != nil && o.Zaakpay != nil {
		return true
	}

	return false
}

// SetZaakpay gets a reference to the given WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay and assigns it to the Zaakpay field.
func (o *WhatsappMessageOrderPaymentGateway) SetZaakpay(v WhatsappMessageOrderPaymentSettingPaymentGatewayZaakpay) {
	o.Zaakpay = &v
}

func (o WhatsappMessageOrderPaymentGateway) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["configuration_name"] = o.ConfigurationName
	}
	if o.Billdesk != nil {
		toSerialize["billdesk"] = o.Billdesk
	}
	if o.Payu != nil {
		toSerialize["payu"] = o.Payu
	}
	if o.Razorpay != nil {
		toSerialize["razorpay"] = o.Razorpay
	}
	if o.Zaakpay != nil {
		toSerialize["zaakpay"] = o.Zaakpay
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappMessageOrderPaymentGateway struct {
	value *WhatsappMessageOrderPaymentGateway
	isSet bool
}

func (v NullableWhatsappMessageOrderPaymentGateway) Get() *WhatsappMessageOrderPaymentGateway {
	return v.value
}

func (v *NullableWhatsappMessageOrderPaymentGateway) Set(val *WhatsappMessageOrderPaymentGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappMessageOrderPaymentGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappMessageOrderPaymentGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappMessageOrderPaymentGateway(val *WhatsappMessageOrderPaymentGateway) *NullableWhatsappMessageOrderPaymentGateway {
	return &NullableWhatsappMessageOrderPaymentGateway{value: val, isSet: true}
}

func (v NullableWhatsappMessageOrderPaymentGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappMessageOrderPaymentGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


