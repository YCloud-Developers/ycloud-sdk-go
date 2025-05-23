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
	"time"
)

// WhatsappAuthIntlRateEligibilityCountry Starting June 1, 2024, we are updating our authentication rate card and introducing a new authentication-international rate. This rate will apply in the the following countries: - June 1, 2024 – Indonesia (country calling code +62, country code `ID`) - July 1, 2024 – India (country calling code +91, country code `IN`)  See also [Authentication-International Rates](https://developers.facebook.com/docs/whatsapp/pricing/authentication-international-rates).
type WhatsappAuthIntlRateEligibilityCountry struct {
	// [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	CountryCode *string `json:"countryCode,omitempty"`
	// Date when newly-opened authentication conversations are subject to authentication-international rates, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., `2024-07-01T00:00:00.000Z`.
	StartTime *time.Time `json:"startTime,omitempty"`
}

// NewWhatsappAuthIntlRateEligibilityCountry instantiates a new WhatsappAuthIntlRateEligibilityCountry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappAuthIntlRateEligibilityCountry() *WhatsappAuthIntlRateEligibilityCountry {
	this := WhatsappAuthIntlRateEligibilityCountry{}
	return &this
}

// NewWhatsappAuthIntlRateEligibilityCountryWithDefaults instantiates a new WhatsappAuthIntlRateEligibilityCountry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappAuthIntlRateEligibilityCountryWithDefaults() *WhatsappAuthIntlRateEligibilityCountry {
	this := WhatsappAuthIntlRateEligibilityCountry{}
	return &this
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *WhatsappAuthIntlRateEligibilityCountry) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappAuthIntlRateEligibilityCountry) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *WhatsappAuthIntlRateEligibilityCountry) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *WhatsappAuthIntlRateEligibilityCountry) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *WhatsappAuthIntlRateEligibilityCountry) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappAuthIntlRateEligibilityCountry) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *WhatsappAuthIntlRateEligibilityCountry) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *WhatsappAuthIntlRateEligibilityCountry) SetStartTime(v time.Time) {
	o.StartTime = &v
}

func (o WhatsappAuthIntlRateEligibilityCountry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CountryCode != nil {
		toSerialize["countryCode"] = o.CountryCode
	}
	if o.StartTime != nil {
		toSerialize["startTime"] = o.StartTime
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappAuthIntlRateEligibilityCountry struct {
	value *WhatsappAuthIntlRateEligibilityCountry
	isSet bool
}

func (v NullableWhatsappAuthIntlRateEligibilityCountry) Get() *WhatsappAuthIntlRateEligibilityCountry {
	return v.value
}

func (v *NullableWhatsappAuthIntlRateEligibilityCountry) Set(val *WhatsappAuthIntlRateEligibilityCountry) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappAuthIntlRateEligibilityCountry) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappAuthIntlRateEligibilityCountry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappAuthIntlRateEligibilityCountry(val *WhatsappAuthIntlRateEligibilityCountry) *NullableWhatsappAuthIntlRateEligibilityCountry {
	return &NullableWhatsappAuthIntlRateEligibilityCountry{value: val, isSet: true}
}

func (v NullableWhatsappAuthIntlRateEligibilityCountry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappAuthIntlRateEligibilityCountry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


