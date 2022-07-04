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

// EmailContentType The MIME type of the email content (`text/html` or `text/plain`). Be aware that We won't count click and open events for the type `text/plain`.
type EmailContentType string

// List of EmailContentType
const (
	TEXT_HTML EmailContentType = "text/html"
	TEXT_PLAIN EmailContentType = "text/plain"
)

// All allowed values of EmailContentType enum
var AllowedEmailContentTypeEnumValues = []EmailContentType{
	"text/html",
	"text/plain",
}

func (v *EmailContentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EmailContentType(value)
	*v = enumTypeValue
	return nil
}

// NewEmailContentTypeFromValue returns a pointer to a valid EmailContentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEmailContentTypeFromValue(v string) (*EmailContentType, error) {
	ev := EmailContentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EmailContentType: valid values are %v", v, AllowedEmailContentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EmailContentType) IsValid() bool {
	for _, existing := range AllowedEmailContentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EmailContentType value
func (v EmailContentType) Ptr() *EmailContentType {
	return &v
}

type NullableEmailContentType struct {
	value *EmailContentType
	isSet bool
}

func (v NullableEmailContentType) Get() *EmailContentType {
	return v.value
}

func (v *NullableEmailContentType) Set(val *EmailContentType) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailContentType) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailContentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailContentType(val *EmailContentType) *NullableEmailContentType {
	return &NullableEmailContentType{value: val, isSet: true}
}

func (v NullableEmailContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailContentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

