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

// WhatsappBusinessAccountUpdateEventEnum Indicates the update event type of the WABA when a notification is sent to you to report a [policy violation](https://developers.facebook.com/docs/whatsapp/overview/policy-enforcement), a WABA has been banned and more. - `DISABLED_UPDATE`: WhatsApp Business Account Banned. - `ACCOUNT_RESTRICTION`: WhatsApp Business Account Restricted Due To Policy Violations. - `ACCOUNT_VIOLATION`: WhatsApp Business Account Violates Policy. - `AUTH_INTL_PRICE_ELIGIBILITY_UPDATE`: WhatsApp Business Account is eligible for the [authentication-international rate](https://developers.facebook.com/docs/whatsapp/pricing/authentication-international-rates). - `BUSINESS_PRIMARY_LOCATION_COUNTRY_UPDATE`: Business's [primary business location](https://developers.facebook.com/docs/whatsapp/pricing/authentication-international-rates#primary-business-location) is set.
type WhatsappBusinessAccountUpdateEventEnum string

// List of WhatsappBusinessAccountUpdateEventEnum
const (
	WHATSAPPBUSINESSACCOUNTUPDATEEVENTENUM_DISABLED_UPDATE WhatsappBusinessAccountUpdateEventEnum = "DISABLED_UPDATE"
	WHATSAPPBUSINESSACCOUNTUPDATEEVENTENUM_ACCOUNT_RESTRICTION WhatsappBusinessAccountUpdateEventEnum = "ACCOUNT_RESTRICTION"
	WHATSAPPBUSINESSACCOUNTUPDATEEVENTENUM_ACCOUNT_VIOLATION WhatsappBusinessAccountUpdateEventEnum = "ACCOUNT_VIOLATION"
	WHATSAPPBUSINESSACCOUNTUPDATEEVENTENUM_AUTH_INTL_PRICE_ELIGIBILITY_UPDATE WhatsappBusinessAccountUpdateEventEnum = "AUTH_INTL_PRICE_ELIGIBILITY_UPDATE"
	WHATSAPPBUSINESSACCOUNTUPDATEEVENTENUM_BUSINESS_PRIMARY_LOCATION_COUNTRY_UPDATE WhatsappBusinessAccountUpdateEventEnum = "BUSINESS_PRIMARY_LOCATION_COUNTRY_UPDATE"
)

// All allowed values of WhatsappBusinessAccountUpdateEventEnum enum
var AllowedWhatsappBusinessAccountUpdateEventEnumEnumValues = []WhatsappBusinessAccountUpdateEventEnum{
	"DISABLED_UPDATE",
	"ACCOUNT_RESTRICTION",
	"ACCOUNT_VIOLATION",
	"AUTH_INTL_PRICE_ELIGIBILITY_UPDATE",
	"BUSINESS_PRIMARY_LOCATION_COUNTRY_UPDATE",
}

func (v *WhatsappBusinessAccountUpdateEventEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WhatsappBusinessAccountUpdateEventEnum(value)
	*v = enumTypeValue
	return nil
}

// NewWhatsappBusinessAccountUpdateEventEnumFromValue returns a pointer to a valid WhatsappBusinessAccountUpdateEventEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWhatsappBusinessAccountUpdateEventEnumFromValue(v string) (*WhatsappBusinessAccountUpdateEventEnum, error) {
	ev := WhatsappBusinessAccountUpdateEventEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WhatsappBusinessAccountUpdateEventEnum: valid values are %v", v, AllowedWhatsappBusinessAccountUpdateEventEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WhatsappBusinessAccountUpdateEventEnum) IsValid() bool {
	for _, existing := range AllowedWhatsappBusinessAccountUpdateEventEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WhatsappBusinessAccountUpdateEventEnum value
func (v WhatsappBusinessAccountUpdateEventEnum) Ptr() *WhatsappBusinessAccountUpdateEventEnum {
	return &v
}

type NullableWhatsappBusinessAccountUpdateEventEnum struct {
	value *WhatsappBusinessAccountUpdateEventEnum
	isSet bool
}

func (v NullableWhatsappBusinessAccountUpdateEventEnum) Get() *WhatsappBusinessAccountUpdateEventEnum {
	return v.value
}

func (v *NullableWhatsappBusinessAccountUpdateEventEnum) Set(val *WhatsappBusinessAccountUpdateEventEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappBusinessAccountUpdateEventEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappBusinessAccountUpdateEventEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappBusinessAccountUpdateEventEnum(val *WhatsappBusinessAccountUpdateEventEnum) *NullableWhatsappBusinessAccountUpdateEventEnum {
	return &NullableWhatsappBusinessAccountUpdateEventEnum{value: val, isSet: true}
}

func (v NullableWhatsappBusinessAccountUpdateEventEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappBusinessAccountUpdateEventEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

