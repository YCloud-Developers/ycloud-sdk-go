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

// WhatsappMessageTemplateComponentCardComponent Card component object containing the parameters of the message.
type WhatsappMessageTemplateComponentCardComponent struct {
	// Component type.
	Type string `json:"type"`
	// **Required when type is `button`.** Type of button. - `quick_reply`: Refers to a previously created quick reply button that allows for the customer to return a predefined message. - `url`: Refers to a previously created url button that allows the customer to visit the URL generated by appending the text parameter to the predefined prefix URL in the template.
	SubType *string `json:"sub_type,omitempty"`
	// **Required when `type` = `button`. Not used for the other types.** Indicates order in which button should appear, if the template uses multiple buttons. Buttons are zero-indexed, so setting value to 0 will cause the button to appear first, and another button with an index of 1 will appear next, etc.
	Index *int32 `json:"index,omitempty"`
	// **Required when `type` = `button`, or there are variables in the corresponding template component, or the card component `HEADER` format is media (`IMAGE`, `VIDEO`).** Array of parameter objects with the content of the message. For components of `type` = `button`, see the [button parameter object](https://developers.facebook.com/docs/whatsapp/cloud-api/reference/messages#button-parameter-object).
	Parameters []WhatsappMessageTemplateComponentParameter `json:"parameters,omitempty"`
}

// NewWhatsappMessageTemplateComponentCardComponent instantiates a new WhatsappMessageTemplateComponentCardComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappMessageTemplateComponentCardComponent(type_ string) *WhatsappMessageTemplateComponentCardComponent {
	this := WhatsappMessageTemplateComponentCardComponent{}
	this.Type = type_
	return &this
}

// NewWhatsappMessageTemplateComponentCardComponentWithDefaults instantiates a new WhatsappMessageTemplateComponentCardComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappMessageTemplateComponentCardComponentWithDefaults() *WhatsappMessageTemplateComponentCardComponent {
	this := WhatsappMessageTemplateComponentCardComponent{}
	return &this
}

// GetType returns the Type field value
func (o *WhatsappMessageTemplateComponentCardComponent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WhatsappMessageTemplateComponentCardComponent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WhatsappMessageTemplateComponentCardComponent) SetType(v string) {
	o.Type = v
}

// GetSubType returns the SubType field value if set, zero value otherwise.
func (o *WhatsappMessageTemplateComponentCardComponent) GetSubType() string {
	if o == nil || o.SubType == nil {
		var ret string
		return ret
	}
	return *o.SubType
}

// GetSubTypeOk returns a tuple with the SubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageTemplateComponentCardComponent) GetSubTypeOk() (*string, bool) {
	if o == nil || o.SubType == nil {
		return nil, false
	}
	return o.SubType, true
}

// HasSubType returns a boolean if a field has been set.
func (o *WhatsappMessageTemplateComponentCardComponent) HasSubType() bool {
	if o != nil && o.SubType != nil {
		return true
	}

	return false
}

// SetSubType gets a reference to the given string and assigns it to the SubType field.
func (o *WhatsappMessageTemplateComponentCardComponent) SetSubType(v string) {
	o.SubType = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *WhatsappMessageTemplateComponentCardComponent) GetIndex() int32 {
	if o == nil || o.Index == nil {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageTemplateComponentCardComponent) GetIndexOk() (*int32, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *WhatsappMessageTemplateComponentCardComponent) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *WhatsappMessageTemplateComponentCardComponent) SetIndex(v int32) {
	o.Index = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *WhatsappMessageTemplateComponentCardComponent) GetParameters() []WhatsappMessageTemplateComponentParameter {
	if o == nil || o.Parameters == nil {
		var ret []WhatsappMessageTemplateComponentParameter
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageTemplateComponentCardComponent) GetParametersOk() ([]WhatsappMessageTemplateComponentParameter, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *WhatsappMessageTemplateComponentCardComponent) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []WhatsappMessageTemplateComponentParameter and assigns it to the Parameters field.
func (o *WhatsappMessageTemplateComponentCardComponent) SetParameters(v []WhatsappMessageTemplateComponentParameter) {
	o.Parameters = v
}

func (o WhatsappMessageTemplateComponentCardComponent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.SubType != nil {
		toSerialize["sub_type"] = o.SubType
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappMessageTemplateComponentCardComponent struct {
	value *WhatsappMessageTemplateComponentCardComponent
	isSet bool
}

func (v NullableWhatsappMessageTemplateComponentCardComponent) Get() *WhatsappMessageTemplateComponentCardComponent {
	return v.value
}

func (v *NullableWhatsappMessageTemplateComponentCardComponent) Set(val *WhatsappMessageTemplateComponentCardComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappMessageTemplateComponentCardComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappMessageTemplateComponentCardComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappMessageTemplateComponentCardComponent(val *WhatsappMessageTemplateComponentCardComponent) *NullableWhatsappMessageTemplateComponentCardComponent {
	return &NullableWhatsappMessageTemplateComponentCardComponent{value: val, isSet: true}
}

func (v NullableWhatsappMessageTemplateComponentCardComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappMessageTemplateComponentCardComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}