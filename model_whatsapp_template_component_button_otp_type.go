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
	"fmt"
)

// WhatsappTemplateComponentButtonOtpType Indicates button OTP type. Set to `COPY_CODE` if you want the template to use a copy code button, `ONE_TAP` to have it use a one-tap autofill button, or `ZERO_TAP` to have no button at all.
type WhatsappTemplateComponentButtonOtpType string

// List of WhatsappTemplateComponentButtonOtpType
const (
	WHATSAPPTEMPLATECOMPONENTBUTTONOTPTYPE_COPY_CODE WhatsappTemplateComponentButtonOtpType = "COPY_CODE"
	WHATSAPPTEMPLATECOMPONENTBUTTONOTPTYPE_ONE_TAP WhatsappTemplateComponentButtonOtpType = "ONE_TAP"
	WHATSAPPTEMPLATECOMPONENTBUTTONOTPTYPE_ZERO_TAP WhatsappTemplateComponentButtonOtpType = "ZERO_TAP"
)

// All allowed values of WhatsappTemplateComponentButtonOtpType enum
var AllowedWhatsappTemplateComponentButtonOtpTypeEnumValues = []WhatsappTemplateComponentButtonOtpType{
	"COPY_CODE",
	"ONE_TAP",
	"ZERO_TAP",
}

func (v *WhatsappTemplateComponentButtonOtpType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WhatsappTemplateComponentButtonOtpType(value)
	*v = enumTypeValue
	return nil
}

// NewWhatsappTemplateComponentButtonOtpTypeFromValue returns a pointer to a valid WhatsappTemplateComponentButtonOtpType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWhatsappTemplateComponentButtonOtpTypeFromValue(v string) (*WhatsappTemplateComponentButtonOtpType, error) {
	ev := WhatsappTemplateComponentButtonOtpType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WhatsappTemplateComponentButtonOtpType: valid values are %v", v, AllowedWhatsappTemplateComponentButtonOtpTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WhatsappTemplateComponentButtonOtpType) IsValid() bool {
	for _, existing := range AllowedWhatsappTemplateComponentButtonOtpTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WhatsappTemplateComponentButtonOtpType value
func (v WhatsappTemplateComponentButtonOtpType) Ptr() *WhatsappTemplateComponentButtonOtpType {
	return &v
}

type NullableWhatsappTemplateComponentButtonOtpType struct {
	value *WhatsappTemplateComponentButtonOtpType
	isSet bool
}

func (v NullableWhatsappTemplateComponentButtonOtpType) Get() *WhatsappTemplateComponentButtonOtpType {
	return v.value
}

func (v *NullableWhatsappTemplateComponentButtonOtpType) Set(val *WhatsappTemplateComponentButtonOtpType) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappTemplateComponentButtonOtpType) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappTemplateComponentButtonOtpType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappTemplateComponentButtonOtpType(val *WhatsappTemplateComponentButtonOtpType) *NullableWhatsappTemplateComponentButtonOtpType {
	return &NullableWhatsappTemplateComponentButtonOtpType{value: val, isSet: true}
}

func (v NullableWhatsappTemplateComponentButtonOtpType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappTemplateComponentButtonOtpType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

