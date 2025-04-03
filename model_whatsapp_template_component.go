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

// WhatsappTemplateComponent struct for WhatsappTemplateComponent
type WhatsappTemplateComponent struct {
	// **Required.** Template component type. - `BODY`: Body components are text-only components and are required by all templates. Templates are limited to one body component. - `HEADER`: Headers are optional components that appear at the top of template messages. Headers support text, media (images, videos, documents). Templates are limited to one header component. - `FOOTER`: Footers are optional text-only components that appear immediately after the body component. Templates are limited to one footer component. - `BUTTONS`: Buttons are optional interactive components that perform specific actions when tapped. - `LIMITED_TIME_OFFER`: Use for limited-time offer templates. The delivered message can display an offer expiration details section with a heading, an optional expiration timer, and the offer code itself. - `CAROUSEL`: Carousel templates allow you to send a single text message (1), accompanied by a set of up to 10 carousel cards (2) in a horizontally scrollable view.
	Type *string `json:"type,omitempty"`
	// **Required for type `HEADER`.**
	Format *string `json:"format,omitempty"`
	// For body text (type = `BODY`), maximum 1024 characters. For header text (type = `HEADER`, format = `TEXT`), maximum 60 characters. For footer text (type = `FOOTER`), maximum 60 characters. For card body text (`CAROUSEL` card component type = `BODY`), maximum 160 characters.
	Text *string `json:"text,omitempty"`
	// **Required for type `BUTTONS`.** Buttons are optional interactive components that perform specific actions when tapped. Templates can have a mixture of up to 10 button components total, although there are limits to individual buttons of the same type as well as combination limits. If a template has more than three buttons, two buttons will appear in the delivered message and the remaining buttons will be replaced with a **See all options** button. Tapping the **See all options** button reveals the remaining buttons.
	Buttons []WhatsappTemplateComponentButton `json:"buttons,omitempty"`
	// **Optional. Only applicable in the `BODY` component of an AUTHENTICATION template.** Set to `true` if you want the template to include the string, *For your security, do not share this code.* Set to `false` to exclude the string.
	AddSecurityRecommendation *bool `json:"add_security_recommendation,omitempty"`
	// **Optional. Only applicable in the `FOOTER` component of an AUTHENTICATION template.** Indicates number of minutes the password or code is valid. If omitted, the code expiration warning will not be displayed in the delivered message. Minimum 1, maximum 90.
	CodeExpirationMinutes *int32 `json:"code_expiration_minutes,omitempty"`
	LimitedTimeOffer *WhatsappTemplateComponentLimitedTimeOffer `json:"limited_time_offer,omitempty"`
	Example *WhatsappTemplateComponentExample `json:"example,omitempty"`
	// **Required for type `CAROUSEL`.** Carousel templates support up to 10 carousel cards.
	Cards []WhatsappTemplateComponentCard `json:"cards,omitempty"`
}

// NewWhatsappTemplateComponent instantiates a new WhatsappTemplateComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappTemplateComponent() *WhatsappTemplateComponent {
	this := WhatsappTemplateComponent{}
	return &this
}

// NewWhatsappTemplateComponentWithDefaults instantiates a new WhatsappTemplateComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappTemplateComponentWithDefaults() *WhatsappTemplateComponent {
	this := WhatsappTemplateComponent{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WhatsappTemplateComponent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WhatsappTemplateComponent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WhatsappTemplateComponent) SetType(v string) {
	o.Type = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *WhatsappTemplateComponent) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponent) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *WhatsappTemplateComponent) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *WhatsappTemplateComponent) SetFormat(v string) {
	o.Format = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *WhatsappTemplateComponent) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponent) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *WhatsappTemplateComponent) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *WhatsappTemplateComponent) SetText(v string) {
	o.Text = &v
}

// GetButtons returns the Buttons field value if set, zero value otherwise.
func (o *WhatsappTemplateComponent) GetButtons() []WhatsappTemplateComponentButton {
	if o == nil || o.Buttons == nil {
		var ret []WhatsappTemplateComponentButton
		return ret
	}
	return o.Buttons
}

// GetButtonsOk returns a tuple with the Buttons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponent) GetButtonsOk() ([]WhatsappTemplateComponentButton, bool) {
	if o == nil || o.Buttons == nil {
		return nil, false
	}
	return o.Buttons, true
}

// HasButtons returns a boolean if a field has been set.
func (o *WhatsappTemplateComponent) HasButtons() bool {
	if o != nil && o.Buttons != nil {
		return true
	}

	return false
}

// SetButtons gets a reference to the given []WhatsappTemplateComponentButton and assigns it to the Buttons field.
func (o *WhatsappTemplateComponent) SetButtons(v []WhatsappTemplateComponentButton) {
	o.Buttons = v
}

// GetAddSecurityRecommendation returns the AddSecurityRecommendation field value if set, zero value otherwise.
func (o *WhatsappTemplateComponent) GetAddSecurityRecommendation() bool {
	if o == nil || o.AddSecurityRecommendation == nil {
		var ret bool
		return ret
	}
	return *o.AddSecurityRecommendation
}

// GetAddSecurityRecommendationOk returns a tuple with the AddSecurityRecommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponent) GetAddSecurityRecommendationOk() (*bool, bool) {
	if o == nil || o.AddSecurityRecommendation == nil {
		return nil, false
	}
	return o.AddSecurityRecommendation, true
}

// HasAddSecurityRecommendation returns a boolean if a field has been set.
func (o *WhatsappTemplateComponent) HasAddSecurityRecommendation() bool {
	if o != nil && o.AddSecurityRecommendation != nil {
		return true
	}

	return false
}

// SetAddSecurityRecommendation gets a reference to the given bool and assigns it to the AddSecurityRecommendation field.
func (o *WhatsappTemplateComponent) SetAddSecurityRecommendation(v bool) {
	o.AddSecurityRecommendation = &v
}

// GetCodeExpirationMinutes returns the CodeExpirationMinutes field value if set, zero value otherwise.
func (o *WhatsappTemplateComponent) GetCodeExpirationMinutes() int32 {
	if o == nil || o.CodeExpirationMinutes == nil {
		var ret int32
		return ret
	}
	return *o.CodeExpirationMinutes
}

// GetCodeExpirationMinutesOk returns a tuple with the CodeExpirationMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponent) GetCodeExpirationMinutesOk() (*int32, bool) {
	if o == nil || o.CodeExpirationMinutes == nil {
		return nil, false
	}
	return o.CodeExpirationMinutes, true
}

// HasCodeExpirationMinutes returns a boolean if a field has been set.
func (o *WhatsappTemplateComponent) HasCodeExpirationMinutes() bool {
	if o != nil && o.CodeExpirationMinutes != nil {
		return true
	}

	return false
}

// SetCodeExpirationMinutes gets a reference to the given int32 and assigns it to the CodeExpirationMinutes field.
func (o *WhatsappTemplateComponent) SetCodeExpirationMinutes(v int32) {
	o.CodeExpirationMinutes = &v
}

// GetLimitedTimeOffer returns the LimitedTimeOffer field value if set, zero value otherwise.
func (o *WhatsappTemplateComponent) GetLimitedTimeOffer() WhatsappTemplateComponentLimitedTimeOffer {
	if o == nil || o.LimitedTimeOffer == nil {
		var ret WhatsappTemplateComponentLimitedTimeOffer
		return ret
	}
	return *o.LimitedTimeOffer
}

// GetLimitedTimeOfferOk returns a tuple with the LimitedTimeOffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponent) GetLimitedTimeOfferOk() (*WhatsappTemplateComponentLimitedTimeOffer, bool) {
	if o == nil || o.LimitedTimeOffer == nil {
		return nil, false
	}
	return o.LimitedTimeOffer, true
}

// HasLimitedTimeOffer returns a boolean if a field has been set.
func (o *WhatsappTemplateComponent) HasLimitedTimeOffer() bool {
	if o != nil && o.LimitedTimeOffer != nil {
		return true
	}

	return false
}

// SetLimitedTimeOffer gets a reference to the given WhatsappTemplateComponentLimitedTimeOffer and assigns it to the LimitedTimeOffer field.
func (o *WhatsappTemplateComponent) SetLimitedTimeOffer(v WhatsappTemplateComponentLimitedTimeOffer) {
	o.LimitedTimeOffer = &v
}

// GetExample returns the Example field value if set, zero value otherwise.
func (o *WhatsappTemplateComponent) GetExample() WhatsappTemplateComponentExample {
	if o == nil || o.Example == nil {
		var ret WhatsappTemplateComponentExample
		return ret
	}
	return *o.Example
}

// GetExampleOk returns a tuple with the Example field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponent) GetExampleOk() (*WhatsappTemplateComponentExample, bool) {
	if o == nil || o.Example == nil {
		return nil, false
	}
	return o.Example, true
}

// HasExample returns a boolean if a field has been set.
func (o *WhatsappTemplateComponent) HasExample() bool {
	if o != nil && o.Example != nil {
		return true
	}

	return false
}

// SetExample gets a reference to the given WhatsappTemplateComponentExample and assigns it to the Example field.
func (o *WhatsappTemplateComponent) SetExample(v WhatsappTemplateComponentExample) {
	o.Example = &v
}

// GetCards returns the Cards field value if set, zero value otherwise.
func (o *WhatsappTemplateComponent) GetCards() []WhatsappTemplateComponentCard {
	if o == nil || o.Cards == nil {
		var ret []WhatsappTemplateComponentCard
		return ret
	}
	return o.Cards
}

// GetCardsOk returns a tuple with the Cards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponent) GetCardsOk() ([]WhatsappTemplateComponentCard, bool) {
	if o == nil || o.Cards == nil {
		return nil, false
	}
	return o.Cards, true
}

// HasCards returns a boolean if a field has been set.
func (o *WhatsappTemplateComponent) HasCards() bool {
	if o != nil && o.Cards != nil {
		return true
	}

	return false
}

// SetCards gets a reference to the given []WhatsappTemplateComponentCard and assigns it to the Cards field.
func (o *WhatsappTemplateComponent) SetCards(v []WhatsappTemplateComponentCard) {
	o.Cards = v
}

func (o WhatsappTemplateComponent) MarshalJSON() ([]byte, error) {
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
	if o.AddSecurityRecommendation != nil {
		toSerialize["add_security_recommendation"] = o.AddSecurityRecommendation
	}
	if o.CodeExpirationMinutes != nil {
		toSerialize["code_expiration_minutes"] = o.CodeExpirationMinutes
	}
	if o.LimitedTimeOffer != nil {
		toSerialize["limited_time_offer"] = o.LimitedTimeOffer
	}
	if o.Example != nil {
		toSerialize["example"] = o.Example
	}
	if o.Cards != nil {
		toSerialize["cards"] = o.Cards
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappTemplateComponent struct {
	value *WhatsappTemplateComponent
	isSet bool
}

func (v NullableWhatsappTemplateComponent) Get() *WhatsappTemplateComponent {
	return v.value
}

func (v *NullableWhatsappTemplateComponent) Set(val *WhatsappTemplateComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappTemplateComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappTemplateComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappTemplateComponent(val *WhatsappTemplateComponent) *NullableWhatsappTemplateComponent {
	return &NullableWhatsappTemplateComponent{value: val, isSet: true}
}

func (v NullableWhatsappTemplateComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappTemplateComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


