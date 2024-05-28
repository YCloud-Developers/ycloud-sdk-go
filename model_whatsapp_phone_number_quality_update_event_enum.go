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

// WhatsappPhoneNumberQualityUpdateEventEnum Indicates the update event type of WhatsApp phone number quality when a notification is sent to you. - `ONBOARDING`: Typically when the messaging limit changes from `TIER_NOT_SET` to another tier.  - `UPGRADE`: Messaging limit tier upgraded. - `DOWNGRADE`: Messaging limit tier downgraded. - `FLAGGED`: Flagged status occurs when the quality rating reaches a low state. If the message quality improves to a high or medium state and maintains this for 7 days, your status will return to Connected. If the quality rating doesn't improve, your status will still return to Connected, but you'll be placed in a lower messaging limit tier. Learn more on [Phone Number Quality Rating](https://www.facebook.com/business/help/896873687365001) docs. - `UNFLAGGED`: Phone number status changes from `FLAGGED` to `CONNECTED`.
type WhatsappPhoneNumberQualityUpdateEventEnum string

// List of WhatsappPhoneNumberQualityUpdateEventEnum
const (
	WHATSAPPPHONENUMBERQUALITYUPDATEEVENTENUM_ONBOARDING WhatsappPhoneNumberQualityUpdateEventEnum = "ONBOARDING"
	WHATSAPPPHONENUMBERQUALITYUPDATEEVENTENUM_UPGRADE    WhatsappPhoneNumberQualityUpdateEventEnum = "UPGRADE"
	WHATSAPPPHONENUMBERQUALITYUPDATEEVENTENUM_DOWNGRADE  WhatsappPhoneNumberQualityUpdateEventEnum = "DOWNGRADE"
	WHATSAPPPHONENUMBERQUALITYUPDATEEVENTENUM_FLAGGED    WhatsappPhoneNumberQualityUpdateEventEnum = "FLAGGED"
	WHATSAPPPHONENUMBERQUALITYUPDATEEVENTENUM_UNFLAGGED  WhatsappPhoneNumberQualityUpdateEventEnum = "UNFLAGGED"
)

// All allowed values of WhatsappPhoneNumberQualityUpdateEventEnum enum
var AllowedWhatsappPhoneNumberQualityUpdateEventEnumEnumValues = []WhatsappPhoneNumberQualityUpdateEventEnum{
	"ONBOARDING",
	"UPGRADE",
	"DOWNGRADE",
	"FLAGGED",
	"UNFLAGGED",
}

func (v *WhatsappPhoneNumberQualityUpdateEventEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WhatsappPhoneNumberQualityUpdateEventEnum(value)
	*v = enumTypeValue
	return nil
}

// NewWhatsappPhoneNumberQualityUpdateEventEnumFromValue returns a pointer to a valid WhatsappPhoneNumberQualityUpdateEventEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWhatsappPhoneNumberQualityUpdateEventEnumFromValue(v string) (*WhatsappPhoneNumberQualityUpdateEventEnum, error) {
	ev := WhatsappPhoneNumberQualityUpdateEventEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WhatsappPhoneNumberQualityUpdateEventEnum: valid values are %v", v, AllowedWhatsappPhoneNumberQualityUpdateEventEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WhatsappPhoneNumberQualityUpdateEventEnum) IsValid() bool {
	for _, existing := range AllowedWhatsappPhoneNumberQualityUpdateEventEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WhatsappPhoneNumberQualityUpdateEventEnum value
func (v WhatsappPhoneNumberQualityUpdateEventEnum) Ptr() *WhatsappPhoneNumberQualityUpdateEventEnum {
	return &v
}

type NullableWhatsappPhoneNumberQualityUpdateEventEnum struct {
	value *WhatsappPhoneNumberQualityUpdateEventEnum
	isSet bool
}

func (v NullableWhatsappPhoneNumberQualityUpdateEventEnum) Get() *WhatsappPhoneNumberQualityUpdateEventEnum {
	return v.value
}

func (v *NullableWhatsappPhoneNumberQualityUpdateEventEnum) Set(val *WhatsappPhoneNumberQualityUpdateEventEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappPhoneNumberQualityUpdateEventEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappPhoneNumberQualityUpdateEventEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappPhoneNumberQualityUpdateEventEnum(val *WhatsappPhoneNumberQualityUpdateEventEnum) *NullableWhatsappPhoneNumberQualityUpdateEventEnum {
	return &NullableWhatsappPhoneNumberQualityUpdateEventEnum{value: val, isSet: true}
}

func (v NullableWhatsappPhoneNumberQualityUpdateEventEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappPhoneNumberQualityUpdateEventEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
