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

// VerificationFallback Contains information about verification fallback. For example, you can enable sms fallback for WhatsApp verification messages.
type VerificationFallback struct {
	// Whether this fallback you requested is supported. If `false` is returned, it means that there are errors for this fallback, and this fallback will not be triggered.
	Supported *bool `json:"supported,omitempty"`
	// The reason why the fallback is unsupported, e.g, `PARAM_INVALID`, `SMS_SIGNATURE_UNAVAILABLE`, `SENDER_ID_UNAVAILABLE`, or `MESSAGING_REGION_UNSUPPORTED`.
	UnsupportedReason *string `json:"unsupportedReason,omitempty"`
	// The detail message why the fallback is unsupported.
	UnsupportedDetail *string `json:"unsupportedDetail,omitempty"`
}

// NewVerificationFallback instantiates a new VerificationFallback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationFallback() *VerificationFallback {
	this := VerificationFallback{}
	return &this
}

// NewVerificationFallbackWithDefaults instantiates a new VerificationFallback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationFallbackWithDefaults() *VerificationFallback {
	this := VerificationFallback{}
	return &this
}

// GetSupported returns the Supported field value if set, zero value otherwise.
func (o *VerificationFallback) GetSupported() bool {
	if o == nil || o.Supported == nil {
		var ret bool
		return ret
	}
	return *o.Supported
}

// GetSupportedOk returns a tuple with the Supported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationFallback) GetSupportedOk() (*bool, bool) {
	if o == nil || o.Supported == nil {
		return nil, false
	}
	return o.Supported, true
}

// HasSupported returns a boolean if a field has been set.
func (o *VerificationFallback) HasSupported() bool {
	if o != nil && o.Supported != nil {
		return true
	}

	return false
}

// SetSupported gets a reference to the given bool and assigns it to the Supported field.
func (o *VerificationFallback) SetSupported(v bool) {
	o.Supported = &v
}

// GetUnsupportedReason returns the UnsupportedReason field value if set, zero value otherwise.
func (o *VerificationFallback) GetUnsupportedReason() string {
	if o == nil || o.UnsupportedReason == nil {
		var ret string
		return ret
	}
	return *o.UnsupportedReason
}

// GetUnsupportedReasonOk returns a tuple with the UnsupportedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationFallback) GetUnsupportedReasonOk() (*string, bool) {
	if o == nil || o.UnsupportedReason == nil {
		return nil, false
	}
	return o.UnsupportedReason, true
}

// HasUnsupportedReason returns a boolean if a field has been set.
func (o *VerificationFallback) HasUnsupportedReason() bool {
	if o != nil && o.UnsupportedReason != nil {
		return true
	}

	return false
}

// SetUnsupportedReason gets a reference to the given string and assigns it to the UnsupportedReason field.
func (o *VerificationFallback) SetUnsupportedReason(v string) {
	o.UnsupportedReason = &v
}

// GetUnsupportedDetail returns the UnsupportedDetail field value if set, zero value otherwise.
func (o *VerificationFallback) GetUnsupportedDetail() string {
	if o == nil || o.UnsupportedDetail == nil {
		var ret string
		return ret
	}
	return *o.UnsupportedDetail
}

// GetUnsupportedDetailOk returns a tuple with the UnsupportedDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationFallback) GetUnsupportedDetailOk() (*string, bool) {
	if o == nil || o.UnsupportedDetail == nil {
		return nil, false
	}
	return o.UnsupportedDetail, true
}

// HasUnsupportedDetail returns a boolean if a field has been set.
func (o *VerificationFallback) HasUnsupportedDetail() bool {
	if o != nil && o.UnsupportedDetail != nil {
		return true
	}

	return false
}

// SetUnsupportedDetail gets a reference to the given string and assigns it to the UnsupportedDetail field.
func (o *VerificationFallback) SetUnsupportedDetail(v string) {
	o.UnsupportedDetail = &v
}

func (o VerificationFallback) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Supported != nil {
		toSerialize["supported"] = o.Supported
	}
	if o.UnsupportedReason != nil {
		toSerialize["unsupportedReason"] = o.UnsupportedReason
	}
	if o.UnsupportedDetail != nil {
		toSerialize["unsupportedDetail"] = o.UnsupportedDetail
	}
	return json.Marshal(toSerialize)
}

type NullableVerificationFallback struct {
	value *VerificationFallback
	isSet bool
}

func (v NullableVerificationFallback) Get() *VerificationFallback {
	return v.value
}

func (v *NullableVerificationFallback) Set(val *VerificationFallback) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationFallback) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationFallback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationFallback(val *VerificationFallback) *NullableVerificationFallback {
	return &NullableVerificationFallback{value: val, isSet: true}
}

func (v NullableVerificationFallback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationFallback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

