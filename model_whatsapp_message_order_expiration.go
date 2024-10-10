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

// WhatsappMessageOrderExpiration Expiration for this order.
type WhatsappMessageOrderExpiration struct {
	// A string of UTC timestamp in seconds of time when order should expire. Minimum threshold is 300 seconds.
	Timestamp string `json:"timestamp"`
	// Text explanation for expiration.
	Description *string `json:"description,omitempty"`
}

// NewWhatsappMessageOrderExpiration instantiates a new WhatsappMessageOrderExpiration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappMessageOrderExpiration(timestamp string) *WhatsappMessageOrderExpiration {
	this := WhatsappMessageOrderExpiration{}
	this.Timestamp = timestamp
	return &this
}

// NewWhatsappMessageOrderExpirationWithDefaults instantiates a new WhatsappMessageOrderExpiration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappMessageOrderExpirationWithDefaults() *WhatsappMessageOrderExpiration {
	this := WhatsappMessageOrderExpiration{}
	return &this
}

// GetTimestamp returns the Timestamp field value
func (o *WhatsappMessageOrderExpiration) GetTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderExpiration) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *WhatsappMessageOrderExpiration) SetTimestamp(v string) {
	o.Timestamp = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WhatsappMessageOrderExpiration) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderExpiration) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WhatsappMessageOrderExpiration) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WhatsappMessageOrderExpiration) SetDescription(v string) {
	o.Description = &v
}

func (o WhatsappMessageOrderExpiration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappMessageOrderExpiration struct {
	value *WhatsappMessageOrderExpiration
	isSet bool
}

func (v NullableWhatsappMessageOrderExpiration) Get() *WhatsappMessageOrderExpiration {
	return v.value
}

func (v *NullableWhatsappMessageOrderExpiration) Set(val *WhatsappMessageOrderExpiration) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappMessageOrderExpiration) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappMessageOrderExpiration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappMessageOrderExpiration(val *WhatsappMessageOrderExpiration) *NullableWhatsappMessageOrderExpiration {
	return &NullableWhatsappMessageOrderExpiration{value: val, isSet: true}
}

func (v NullableWhatsappMessageOrderExpiration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappMessageOrderExpiration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
