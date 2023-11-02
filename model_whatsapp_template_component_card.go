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

// WhatsappTemplateComponentCard Carousel templates support up to 10 carousel cards. Cards must have a media header (image or video) and can optionally include body text and up to 2 quick reply buttons, phone number buttons, or URL buttons (button types can be mixed).
type WhatsappTemplateComponentCard struct {
	// **Required.** Card components.
	Components []WhatsappTemplateComponent `json:"components,omitempty"`
}

// NewWhatsappTemplateComponentCard instantiates a new WhatsappTemplateComponentCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappTemplateComponentCard() *WhatsappTemplateComponentCard {
	this := WhatsappTemplateComponentCard{}
	return &this
}

// NewWhatsappTemplateComponentCardWithDefaults instantiates a new WhatsappTemplateComponentCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappTemplateComponentCardWithDefaults() *WhatsappTemplateComponentCard {
	this := WhatsappTemplateComponentCard{}
	return &this
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *WhatsappTemplateComponentCard) GetComponents() []WhatsappTemplateComponent {
	if o == nil || o.Components == nil {
		var ret []WhatsappTemplateComponent
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponentCard) GetComponentsOk() ([]WhatsappTemplateComponent, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *WhatsappTemplateComponentCard) HasComponents() bool {
	if o != nil && o.Components != nil {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []WhatsappTemplateComponent and assigns it to the Components field.
func (o *WhatsappTemplateComponentCard) SetComponents(v []WhatsappTemplateComponent) {
	o.Components = v
}

func (o WhatsappTemplateComponentCard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappTemplateComponentCard struct {
	value *WhatsappTemplateComponentCard
	isSet bool
}

func (v NullableWhatsappTemplateComponentCard) Get() *WhatsappTemplateComponentCard {
	return v.value
}

func (v *NullableWhatsappTemplateComponentCard) Set(val *WhatsappTemplateComponentCard) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappTemplateComponentCard) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappTemplateComponentCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappTemplateComponentCard(val *WhatsappTemplateComponentCard) *NullableWhatsappTemplateComponentCard {
	return &NullableWhatsappTemplateComponentCard{value: val, isSet: true}
}

func (v NullableWhatsappTemplateComponentCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappTemplateComponentCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
