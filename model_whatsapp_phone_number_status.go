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

// WhatsappPhoneNumberStatus The status of a WhatsApp business phone number. - `PENDING`: Pending. Phone number is newly added. Verify and register this phone number so it can be connected to your account. - `UNVERIFIED`: Unverified. Verify this phone number to start sending messages. - `MANUAL_REVIEW`: Being reviewed. Phone number is currently being reviewed for connection to your account. - `DISCONNECTED`: Offline. Phone number is currently not reachable by WhatsApp servers. - `CONNECTED`: Connected. Phone number is associated with this account and working properly. - `FLAGGED`: Flagged. This phone number has been flagged due to low quality messages. - `WARNED`: Warned. A warning has been issued for this number, potentially due to spam reports. - `RATE_LIMITED`: Rate limited. The number of messages you can send from this phone number may be restricted. - `BANNED`: Banned. Phone number cannot be used with a WhatsApp account. - `RESTRICTED`: Restricted. This phone number has reached its 24-hour messaging limit and can no longer send messages to customers. Please wait until the messaging limit is reset to send messages. - `BLOCKED`: Message limit reached. The limit has been reached for this 24-hour period. - `MIGRATED`: Transferred. This phone number has been transferred to another WhatsApp Business account. - `UNKNOWN`: Unavailable. The status of this phone number can't be determined right now.
type WhatsappPhoneNumberStatus string

// List of WhatsappPhoneNumberStatus
const (
	WHATSAPPPHONENUMBERSTATUS_PENDING WhatsappPhoneNumberStatus = "PENDING"
	WHATSAPPPHONENUMBERSTATUS_UNVERIFIED WhatsappPhoneNumberStatus = "UNVERIFIED"
	WHATSAPPPHONENUMBERSTATUS_MANUAL_REVIEW WhatsappPhoneNumberStatus = "MANUAL_REVIEW"
	WHATSAPPPHONENUMBERSTATUS_DISCONNECTED WhatsappPhoneNumberStatus = "DISCONNECTED"
	WHATSAPPPHONENUMBERSTATUS_CONNECTED WhatsappPhoneNumberStatus = "CONNECTED"
	WHATSAPPPHONENUMBERSTATUS_FLAGGED WhatsappPhoneNumberStatus = "FLAGGED"
	WHATSAPPPHONENUMBERSTATUS_WARNED WhatsappPhoneNumberStatus = "WARNED"
	WHATSAPPPHONENUMBERSTATUS_RATE_LIMITED WhatsappPhoneNumberStatus = "RATE_LIMITED"
	WHATSAPPPHONENUMBERSTATUS_BANNED WhatsappPhoneNumberStatus = "BANNED"
	WHATSAPPPHONENUMBERSTATUS_RESTRICTED WhatsappPhoneNumberStatus = "RESTRICTED"
	WHATSAPPPHONENUMBERSTATUS_BLOCKED WhatsappPhoneNumberStatus = "BLOCKED"
	WHATSAPPPHONENUMBERSTATUS_MIGRATED WhatsappPhoneNumberStatus = "MIGRATED"
	WHATSAPPPHONENUMBERSTATUS_UNKNOWN WhatsappPhoneNumberStatus = "UNKNOWN"
)

// All allowed values of WhatsappPhoneNumberStatus enum
var AllowedWhatsappPhoneNumberStatusEnumValues = []WhatsappPhoneNumberStatus{
	"PENDING",
	"UNVERIFIED",
	"MANUAL_REVIEW",
	"DISCONNECTED",
	"CONNECTED",
	"FLAGGED",
	"WARNED",
	"RATE_LIMITED",
	"BANNED",
	"RESTRICTED",
	"BLOCKED",
	"MIGRATED",
	"UNKNOWN",
}

func (v *WhatsappPhoneNumberStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WhatsappPhoneNumberStatus(value)
	*v = enumTypeValue
	return nil
}

// NewWhatsappPhoneNumberStatusFromValue returns a pointer to a valid WhatsappPhoneNumberStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWhatsappPhoneNumberStatusFromValue(v string) (*WhatsappPhoneNumberStatus, error) {
	ev := WhatsappPhoneNumberStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WhatsappPhoneNumberStatus: valid values are %v", v, AllowedWhatsappPhoneNumberStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WhatsappPhoneNumberStatus) IsValid() bool {
	for _, existing := range AllowedWhatsappPhoneNumberStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WhatsappPhoneNumberStatus value
func (v WhatsappPhoneNumberStatus) Ptr() *WhatsappPhoneNumberStatus {
	return &v
}

type NullableWhatsappPhoneNumberStatus struct {
	value *WhatsappPhoneNumberStatus
	isSet bool
}

func (v NullableWhatsappPhoneNumberStatus) Get() *WhatsappPhoneNumberStatus {
	return v.value
}

func (v *NullableWhatsappPhoneNumberStatus) Set(val *WhatsappPhoneNumberStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappPhoneNumberStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappPhoneNumberStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappPhoneNumberStatus(val *WhatsappPhoneNumberStatus) *NullableWhatsappPhoneNumberStatus {
	return &NullableWhatsappPhoneNumberStatus{value: val, isSet: true}
}

func (v NullableWhatsappPhoneNumberStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappPhoneNumberStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

