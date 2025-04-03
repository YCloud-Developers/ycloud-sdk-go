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

// WhatsappReviewDecision Used if a decision about WhatsApp accounts or phone numbers has been made.
type WhatsappReviewDecision string

// List of WhatsappReviewDecision
const (
	WHATSAPPREVIEWDECISION_APPROVED WhatsappReviewDecision = "APPROVED"
	WHATSAPPREVIEWDECISION_REJECTED WhatsappReviewDecision = "REJECTED"
	WHATSAPPREVIEWDECISION_DEFERRED WhatsappReviewDecision = "DEFERRED"
)

// All allowed values of WhatsappReviewDecision enum
var AllowedWhatsappReviewDecisionEnumValues = []WhatsappReviewDecision{
	"APPROVED",
	"REJECTED",
	"DEFERRED",
}

func (v *WhatsappReviewDecision) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WhatsappReviewDecision(value)
	*v = enumTypeValue
	return nil
}

// NewWhatsappReviewDecisionFromValue returns a pointer to a valid WhatsappReviewDecision
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWhatsappReviewDecisionFromValue(v string) (*WhatsappReviewDecision, error) {
	ev := WhatsappReviewDecision(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WhatsappReviewDecision: valid values are %v", v, AllowedWhatsappReviewDecisionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WhatsappReviewDecision) IsValid() bool {
	for _, existing := range AllowedWhatsappReviewDecisionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WhatsappReviewDecision value
func (v WhatsappReviewDecision) Ptr() *WhatsappReviewDecision {
	return &v
}

type NullableWhatsappReviewDecision struct {
	value *WhatsappReviewDecision
	isSet bool
}

func (v NullableWhatsappReviewDecision) Get() *WhatsappReviewDecision {
	return v.value
}

func (v *NullableWhatsappReviewDecision) Set(val *WhatsappReviewDecision) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappReviewDecision) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappReviewDecision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappReviewDecision(val *WhatsappReviewDecision) *NullableWhatsappReviewDecision {
	return &NullableWhatsappReviewDecision{value: val, isSet: true}
}

func (v NullableWhatsappReviewDecision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappReviewDecision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

