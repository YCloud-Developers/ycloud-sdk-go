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

// WhatsappMessageInteractiveBody Optional for type `product`. Required for other message types.
type WhatsappMessageInteractiveBody struct {
	// The body content of the message. Emojis and markdown are supported. Maximum length: 1024 characters.
	Text *string `json:"text,omitempty"`
}

// NewWhatsappMessageInteractiveBody instantiates a new WhatsappMessageInteractiveBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappMessageInteractiveBody() *WhatsappMessageInteractiveBody {
	this := WhatsappMessageInteractiveBody{}
	return &this
}

// NewWhatsappMessageInteractiveBodyWithDefaults instantiates a new WhatsappMessageInteractiveBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappMessageInteractiveBodyWithDefaults() *WhatsappMessageInteractiveBody {
	this := WhatsappMessageInteractiveBody{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *WhatsappMessageInteractiveBody) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageInteractiveBody) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *WhatsappMessageInteractiveBody) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *WhatsappMessageInteractiveBody) SetText(v string) {
	o.Text = &v
}

func (o WhatsappMessageInteractiveBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappMessageInteractiveBody struct {
	value *WhatsappMessageInteractiveBody
	isSet bool
}

func (v NullableWhatsappMessageInteractiveBody) Get() *WhatsappMessageInteractiveBody {
	return v.value
}

func (v *NullableWhatsappMessageInteractiveBody) Set(val *WhatsappMessageInteractiveBody) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappMessageInteractiveBody) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappMessageInteractiveBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappMessageInteractiveBody(val *WhatsappMessageInteractiveBody) *NullableWhatsappMessageInteractiveBody {
	return &NullableWhatsappMessageInteractiveBody{value: val, isSet: true}
}

func (v NullableWhatsappMessageInteractiveBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappMessageInteractiveBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


