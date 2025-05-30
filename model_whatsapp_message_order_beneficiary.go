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

// WhatsappMessageOrderBeneficiary A beneficiary is an intended recipient for shipping the physical goods in the order. Beneficiary information isn't shown to users but is needed for legal and compliance reasons.
type WhatsappMessageOrderBeneficiary struct {
	// Name of the individual or business receiving the physical goods. Cannot exceed 200 characters.
	Name string `json:"name"`
	// Shipping address (Door/Tower Number, Street Name etc.). Cannot exceed 100 characters.
	AddressLine1 string `json:"address_line1"`
	// Shipping address (Landmark, Area, etc.). Cannot exceed 100 characters.
	AddressLine2 *string `json:"address_line2,omitempty"`
	// Name of the city.
	City string `json:"city"`
	// Name of the state.
	State string `json:"state"`
	// Name of the country. Currently the only supported value is `India`.
	Country string `json:"country"`
	// 6-digit zipcode of shipping address.
	PostalCode string `json:"postal_code"`
}

// NewWhatsappMessageOrderBeneficiary instantiates a new WhatsappMessageOrderBeneficiary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappMessageOrderBeneficiary(name string, addressLine1 string, city string, state string, country string, postalCode string) *WhatsappMessageOrderBeneficiary {
	this := WhatsappMessageOrderBeneficiary{}
	this.Name = name
	this.AddressLine1 = addressLine1
	this.City = city
	this.State = state
	this.Country = country
	this.PostalCode = postalCode
	return &this
}

// NewWhatsappMessageOrderBeneficiaryWithDefaults instantiates a new WhatsappMessageOrderBeneficiary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappMessageOrderBeneficiaryWithDefaults() *WhatsappMessageOrderBeneficiary {
	this := WhatsappMessageOrderBeneficiary{}
	return &this
}

// GetName returns the Name field value
func (o *WhatsappMessageOrderBeneficiary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderBeneficiary) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WhatsappMessageOrderBeneficiary) SetName(v string) {
	o.Name = v
}

// GetAddressLine1 returns the AddressLine1 field value
func (o *WhatsappMessageOrderBeneficiary) GetAddressLine1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderBeneficiary) GetAddressLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressLine1, true
}

// SetAddressLine1 sets field value
func (o *WhatsappMessageOrderBeneficiary) SetAddressLine1(v string) {
	o.AddressLine1 = v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *WhatsappMessageOrderBeneficiary) GetAddressLine2() string {
	if o == nil || o.AddressLine2 == nil {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderBeneficiary) GetAddressLine2Ok() (*string, bool) {
	if o == nil || o.AddressLine2 == nil {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *WhatsappMessageOrderBeneficiary) HasAddressLine2() bool {
	if o != nil && o.AddressLine2 != nil {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *WhatsappMessageOrderBeneficiary) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetCity returns the City field value
func (o *WhatsappMessageOrderBeneficiary) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderBeneficiary) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *WhatsappMessageOrderBeneficiary) SetCity(v string) {
	o.City = v
}

// GetState returns the State field value
func (o *WhatsappMessageOrderBeneficiary) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderBeneficiary) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *WhatsappMessageOrderBeneficiary) SetState(v string) {
	o.State = v
}

// GetCountry returns the Country field value
func (o *WhatsappMessageOrderBeneficiary) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderBeneficiary) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *WhatsappMessageOrderBeneficiary) SetCountry(v string) {
	o.Country = v
}

// GetPostalCode returns the PostalCode field value
func (o *WhatsappMessageOrderBeneficiary) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *WhatsappMessageOrderBeneficiary) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *WhatsappMessageOrderBeneficiary) SetPostalCode(v string) {
	o.PostalCode = v
}

func (o WhatsappMessageOrderBeneficiary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["address_line1"] = o.AddressLine1
	}
	if o.AddressLine2 != nil {
		toSerialize["address_line2"] = o.AddressLine2
	}
	if true {
		toSerialize["city"] = o.City
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["country"] = o.Country
	}
	if true {
		toSerialize["postal_code"] = o.PostalCode
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappMessageOrderBeneficiary struct {
	value *WhatsappMessageOrderBeneficiary
	isSet bool
}

func (v NullableWhatsappMessageOrderBeneficiary) Get() *WhatsappMessageOrderBeneficiary {
	return v.value
}

func (v *NullableWhatsappMessageOrderBeneficiary) Set(val *WhatsappMessageOrderBeneficiary) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappMessageOrderBeneficiary) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappMessageOrderBeneficiary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappMessageOrderBeneficiary(val *WhatsappMessageOrderBeneficiary) *NullableWhatsappMessageOrderBeneficiary {
	return &NullableWhatsappMessageOrderBeneficiary{value: val, isSet: true}
}

func (v NullableWhatsappMessageOrderBeneficiary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappMessageOrderBeneficiary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


