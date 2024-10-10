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

// WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk Additional info for BillDesk. User-defined fields (extra) are used to store any information corresponding to a particular order. Each extra field has a maximum character limit of 120.
type WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk struct {
	AdditionalInfo1 *string `json:"additional_info1,omitempty"`
	AdditionalInfo2 *string `json:"additional_info2,omitempty"`
	AdditionalInfo3 *string `json:"additional_info3,omitempty"`
	AdditionalInfo4 *string `json:"additional_info4,omitempty"`
	AdditionalInfo5 *string `json:"additional_info5,omitempty"`
	AdditionalInfo6 *string `json:"additional_info6,omitempty"`
	AdditionalInfo7 *string `json:"additional_info7,omitempty"`
}

// NewWhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk instantiates a new WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk() *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk {
	this := WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk{}
	return &this
}

// NewWhatsappMessageOrderPaymentSettingPaymentGatewayBilldeskWithDefaults instantiates a new WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappMessageOrderPaymentSettingPaymentGatewayBilldeskWithDefaults() *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk {
	this := WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk{}
	return &this
}

// GetAdditionalInfo1 returns the AdditionalInfo1 field value if set, zero value otherwise.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) GetAdditionalInfo1() string {
	if o == nil || o.AdditionalInfo1 == nil {
		var ret string
		return ret
	}
	return *o.AdditionalInfo1
}

// GetAdditionalInfo1Ok returns a tuple with the AdditionalInfo1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) GetAdditionalInfo1Ok() (*string, bool) {
	if o == nil || o.AdditionalInfo1 == nil {
		return nil, false
	}
	return o.AdditionalInfo1, true
}

// HasAdditionalInfo1 returns a boolean if a field has been set.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) HasAdditionalInfo1() bool {
	if o != nil && o.AdditionalInfo1 != nil {
		return true
	}

	return false
}

// SetAdditionalInfo1 gets a reference to the given string and assigns it to the AdditionalInfo1 field.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) SetAdditionalInfo1(v string) {
	o.AdditionalInfo1 = &v
}

// GetAdditionalInfo2 returns the AdditionalInfo2 field value if set, zero value otherwise.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) GetAdditionalInfo2() string {
	if o == nil || o.AdditionalInfo2 == nil {
		var ret string
		return ret
	}
	return *o.AdditionalInfo2
}

// GetAdditionalInfo2Ok returns a tuple with the AdditionalInfo2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) GetAdditionalInfo2Ok() (*string, bool) {
	if o == nil || o.AdditionalInfo2 == nil {
		return nil, false
	}
	return o.AdditionalInfo2, true
}

// HasAdditionalInfo2 returns a boolean if a field has been set.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) HasAdditionalInfo2() bool {
	if o != nil && o.AdditionalInfo2 != nil {
		return true
	}

	return false
}

// SetAdditionalInfo2 gets a reference to the given string and assigns it to the AdditionalInfo2 field.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) SetAdditionalInfo2(v string) {
	o.AdditionalInfo2 = &v
}

// GetAdditionalInfo3 returns the AdditionalInfo3 field value if set, zero value otherwise.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) GetAdditionalInfo3() string {
	if o == nil || o.AdditionalInfo3 == nil {
		var ret string
		return ret
	}
	return *o.AdditionalInfo3
}

// GetAdditionalInfo3Ok returns a tuple with the AdditionalInfo3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) GetAdditionalInfo3Ok() (*string, bool) {
	if o == nil || o.AdditionalInfo3 == nil {
		return nil, false
	}
	return o.AdditionalInfo3, true
}

// HasAdditionalInfo3 returns a boolean if a field has been set.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) HasAdditionalInfo3() bool {
	if o != nil && o.AdditionalInfo3 != nil {
		return true
	}

	return false
}

// SetAdditionalInfo3 gets a reference to the given string and assigns it to the AdditionalInfo3 field.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) SetAdditionalInfo3(v string) {
	o.AdditionalInfo3 = &v
}

// GetAdditionalInfo4 returns the AdditionalInfo4 field value if set, zero value otherwise.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) GetAdditionalInfo4() string {
	if o == nil || o.AdditionalInfo4 == nil {
		var ret string
		return ret
	}
	return *o.AdditionalInfo4
}

// GetAdditionalInfo4Ok returns a tuple with the AdditionalInfo4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) GetAdditionalInfo4Ok() (*string, bool) {
	if o == nil || o.AdditionalInfo4 == nil {
		return nil, false
	}
	return o.AdditionalInfo4, true
}

// HasAdditionalInfo4 returns a boolean if a field has been set.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) HasAdditionalInfo4() bool {
	if o != nil && o.AdditionalInfo4 != nil {
		return true
	}

	return false
}

// SetAdditionalInfo4 gets a reference to the given string and assigns it to the AdditionalInfo4 field.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) SetAdditionalInfo4(v string) {
	o.AdditionalInfo4 = &v
}

// GetAdditionalInfo5 returns the AdditionalInfo5 field value if set, zero value otherwise.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) GetAdditionalInfo5() string {
	if o == nil || o.AdditionalInfo5 == nil {
		var ret string
		return ret
	}
	return *o.AdditionalInfo5
}

// GetAdditionalInfo5Ok returns a tuple with the AdditionalInfo5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) GetAdditionalInfo5Ok() (*string, bool) {
	if o == nil || o.AdditionalInfo5 == nil {
		return nil, false
	}
	return o.AdditionalInfo5, true
}

// HasAdditionalInfo5 returns a boolean if a field has been set.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) HasAdditionalInfo5() bool {
	if o != nil && o.AdditionalInfo5 != nil {
		return true
	}

	return false
}

// SetAdditionalInfo5 gets a reference to the given string and assigns it to the AdditionalInfo5 field.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) SetAdditionalInfo5(v string) {
	o.AdditionalInfo5 = &v
}

// GetAdditionalInfo6 returns the AdditionalInfo6 field value if set, zero value otherwise.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) GetAdditionalInfo6() string {
	if o == nil || o.AdditionalInfo6 == nil {
		var ret string
		return ret
	}
	return *o.AdditionalInfo6
}

// GetAdditionalInfo6Ok returns a tuple with the AdditionalInfo6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) GetAdditionalInfo6Ok() (*string, bool) {
	if o == nil || o.AdditionalInfo6 == nil {
		return nil, false
	}
	return o.AdditionalInfo6, true
}

// HasAdditionalInfo6 returns a boolean if a field has been set.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) HasAdditionalInfo6() bool {
	if o != nil && o.AdditionalInfo6 != nil {
		return true
	}

	return false
}

// SetAdditionalInfo6 gets a reference to the given string and assigns it to the AdditionalInfo6 field.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) SetAdditionalInfo6(v string) {
	o.AdditionalInfo6 = &v
}

// GetAdditionalInfo7 returns the AdditionalInfo7 field value if set, zero value otherwise.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) GetAdditionalInfo7() string {
	if o == nil || o.AdditionalInfo7 == nil {
		var ret string
		return ret
	}
	return *o.AdditionalInfo7
}

// GetAdditionalInfo7Ok returns a tuple with the AdditionalInfo7 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) GetAdditionalInfo7Ok() (*string, bool) {
	if o == nil || o.AdditionalInfo7 == nil {
		return nil, false
	}
	return o.AdditionalInfo7, true
}

// HasAdditionalInfo7 returns a boolean if a field has been set.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) HasAdditionalInfo7() bool {
	if o != nil && o.AdditionalInfo7 != nil {
		return true
	}

	return false
}

// SetAdditionalInfo7 gets a reference to the given string and assigns it to the AdditionalInfo7 field.
func (o *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) SetAdditionalInfo7(v string) {
	o.AdditionalInfo7 = &v
}

func (o WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalInfo1 != nil {
		toSerialize["additional_info1"] = o.AdditionalInfo1
	}
	if o.AdditionalInfo2 != nil {
		toSerialize["additional_info2"] = o.AdditionalInfo2
	}
	if o.AdditionalInfo3 != nil {
		toSerialize["additional_info3"] = o.AdditionalInfo3
	}
	if o.AdditionalInfo4 != nil {
		toSerialize["additional_info4"] = o.AdditionalInfo4
	}
	if o.AdditionalInfo5 != nil {
		toSerialize["additional_info5"] = o.AdditionalInfo5
	}
	if o.AdditionalInfo6 != nil {
		toSerialize["additional_info6"] = o.AdditionalInfo6
	}
	if o.AdditionalInfo7 != nil {
		toSerialize["additional_info7"] = o.AdditionalInfo7
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk struct {
	value *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk
	isSet bool
}

func (v NullableWhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) Get() *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk {
	return v.value
}

func (v *NullableWhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) Set(val *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk(val *WhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) *NullableWhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk {
	return &NullableWhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk{value: val, isSet: true}
}

func (v NullableWhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappMessageOrderPaymentSettingPaymentGatewayBilldesk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
