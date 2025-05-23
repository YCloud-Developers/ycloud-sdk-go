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

// WhatsappTemplateComponentCardComponent struct for WhatsappTemplateComponentCardComponent
type WhatsappTemplateComponentCardComponent struct {
	// **Required.** Card component type. - `BODY`: Body components are text-only components. Cards must have body text. - `HEADER`: Cards must have a media header (image or video). - `BUTTONS`: Buttons are interactive components that perform specific actions when tapped. Cards must have at least one button, up to 2 buttons.
	Type *string `json:"type,omitempty"`
	// **Required for type `HEADER`.** Cards must have a media header (image or video).
	Format *string `json:"format,omitempty"`
	// **Required for type `BODY`.** Card body text supports variables. Maximum 160 characters.
	Text *string `json:"text,omitempty"`
	// **Required for type `BUTTONS`.** Cards must have at least one button. Supports 2 buttons. Buttons can be the same or a mix of quick reply buttons, phone number buttons, or URL buttons.
	Buttons []WhatsappTemplateComponentButton `json:"buttons,omitempty"`
	Example *WhatsappTemplateComponentExample `json:"example,omitempty"`
}

// NewWhatsappTemplateComponentCardComponent instantiates a new WhatsappTemplateComponentCardComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappTemplateComponentCardComponent() *WhatsappTemplateComponentCardComponent {
	this := WhatsappTemplateComponentCardComponent{}
	return &this
}

// NewWhatsappTemplateComponentCardComponentWithDefaults instantiates a new WhatsappTemplateComponentCardComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappTemplateComponentCardComponentWithDefaults() *WhatsappTemplateComponentCardComponent {
	this := WhatsappTemplateComponentCardComponent{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WhatsappTemplateComponentCardComponent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponentCardComponent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WhatsappTemplateComponentCardComponent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WhatsappTemplateComponentCardComponent) SetType(v string) {
	o.Type = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *WhatsappTemplateComponentCardComponent) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponentCardComponent) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *WhatsappTemplateComponentCardComponent) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *WhatsappTemplateComponentCardComponent) SetFormat(v string) {
	o.Format = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *WhatsappTemplateComponentCardComponent) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponentCardComponent) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *WhatsappTemplateComponentCardComponent) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *WhatsappTemplateComponentCardComponent) SetText(v string) {
	o.Text = &v
}

// GetButtons returns the Buttons field value if set, zero value otherwise.
func (o *WhatsappTemplateComponentCardComponent) GetButtons() []WhatsappTemplateComponentButton {
	if o == nil || o.Buttons == nil {
		var ret []WhatsappTemplateComponentButton
		return ret
	}
	return o.Buttons
}

// GetButtonsOk returns a tuple with the Buttons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponentCardComponent) GetButtonsOk() ([]WhatsappTemplateComponentButton, bool) {
	if o == nil || o.Buttons == nil {
		return nil, false
	}
	return o.Buttons, true
}

// HasButtons returns a boolean if a field has been set.
func (o *WhatsappTemplateComponentCardComponent) HasButtons() bool {
	if o != nil && o.Buttons != nil {
		return true
	}

	return false
}

// SetButtons gets a reference to the given []WhatsappTemplateComponentButton and assigns it to the Buttons field.
func (o *WhatsappTemplateComponentCardComponent) SetButtons(v []WhatsappTemplateComponentButton) {
	o.Buttons = v
}

// GetExample returns the Example field value if set, zero value otherwise.
func (o *WhatsappTemplateComponentCardComponent) GetExample() WhatsappTemplateComponentExample {
	if o == nil || o.Example == nil {
		var ret WhatsappTemplateComponentExample
		return ret
	}
	return *o.Example
}

// GetExampleOk returns a tuple with the Example field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponentCardComponent) GetExampleOk() (*WhatsappTemplateComponentExample, bool) {
	if o == nil || o.Example == nil {
		return nil, false
	}
	return o.Example, true
}

// HasExample returns a boolean if a field has been set.
func (o *WhatsappTemplateComponentCardComponent) HasExample() bool {
	if o != nil && o.Example != nil {
		return true
	}

	return false
}

// SetExample gets a reference to the given WhatsappTemplateComponentExample and assigns it to the Example field.
func (o *WhatsappTemplateComponentCardComponent) SetExample(v WhatsappTemplateComponentExample) {
	o.Example = &v
}

func (o WhatsappTemplateComponentCardComponent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.Buttons != nil {
		toSerialize["buttons"] = o.Buttons
	}
	if o.Example != nil {
		toSerialize["example"] = o.Example
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappTemplateComponentCardComponent struct {
	value *WhatsappTemplateComponentCardComponent
	isSet bool
}

func (v NullableWhatsappTemplateComponentCardComponent) Get() *WhatsappTemplateComponentCardComponent {
	return v.value
}

func (v *NullableWhatsappTemplateComponentCardComponent) Set(val *WhatsappTemplateComponentCardComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappTemplateComponentCardComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappTemplateComponentCardComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappTemplateComponentCardComponent(val *WhatsappTemplateComponentCardComponent) *NullableWhatsappTemplateComponentCardComponent {
	return &NullableWhatsappTemplateComponentCardComponent{value: val, isSet: true}
}

func (v NullableWhatsappTemplateComponentCardComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappTemplateComponentCardComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


