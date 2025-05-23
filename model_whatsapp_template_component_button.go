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

// WhatsappTemplateComponentButton struct for WhatsappTemplateComponentButton
type WhatsappTemplateComponentButton struct {
	Type WhatsappTemplateComponentButtonType `json:"type"`
	// **Required for button type `PHONE_NUMBER` or `URL`.** Button text. For `CODE_CODE` buttons, the text is a pre-set value and cannot be customized. For `OTP` buttons, if omitted, the text will default to a pre-set value localized to the template's language. For example, `Copy Code` for English (US). If your template is using a one-tap autofill button and you supply this value, the authentication template message will display a copy code button with this text if we are unable to validate your [handshake](https://developers.facebook.com/docs/whatsapp/business-management-api/authentication-templates/autofill-button-authentication-templates#handshake). Maximum 25 characters.
	Text *string `json:"text,omitempty"`
	// **Required for button type `URL`.** URL of website. There can be at most 1 variable at the end of the URL. Example: `https://www.luckyshrub.com/shop?promo={{1}}`. 2000 characters maximum.
	Url *string `json:"url,omitempty"`
	// **Required for button type `PHONE_NUMBER`.** Alphanumeric string. Business phone number to be (display phone number) called when the user taps the button. 20 characters maximum.
	PhoneNumber *string `json:"phone_number,omitempty"`
	OtpType *WhatsappTemplateComponentButtonOtpType `json:"otp_type,omitempty"`
	// **One-tap and zero-tap buttons only.** One-tap button text. Maximum 25 characters.
	AutofillText *string `json:"autofill_text,omitempty"`
	// **One-tap and zero-tap buttons only.** Your Android app's package name.
	PackageName *string `json:"package_name,omitempty"`
	// **One-tap and zero-tap buttons only.** Your app signing key hash. See [App Signing Key Hash](https://developers.facebook.com/docs/whatsapp/business-management-api/authentication-templates/zero-tap-authentication-templates#app-signing-key-hash).
	SignatureHash *string `json:"signature_hash,omitempty"`
	// **Zero-tap buttons only.** Set to `true` to indicate that you understand that your use of zero-tap authentication is subject to the WhatsApp Business Terms of Service, and that it's your responsibility to ensure your customers expect that the code will be automatically filled in on their behalf when they choose to receive the zero-tap code through WhatsApp. If set to `false`, the template will not be created as you need to accept zero-tap terms before creating zero-tap enabled message templates.
	ZeroTapTermsAccepted *bool `json:"zero_tap_terms_accepted,omitempty"`
	// Sample full URL for a `URL` button with a variable.
	Example []string `json:"example,omitempty"`
	// **Conditionally required for button type `FLOW`.** The unique ID of the Flow. Cannot be used if `flow_name` or `flow_json` parameters are provided. Only one of these parameters is allowed.
	FlowId *string `json:"flow_id,omitempty"`
	// **Conditionally required for button type `FLOW`.** The name of the Flow. Cannot be used if `flow_id` or `flow_json` parameters are provided. Only one of these parameters is allowed. The Flow ID is stored in the message template, not the name, so changing the Flow name will not affect existing message templates.
	FlowName *string `json:"flow_name,omitempty"`
	// **Conditionally required for button type `FLOW`.** The Flow JSON encoded as string with escaping. The Flow JSON specifies the content of the Flow. Cannot be used if `flow_id` or `flow_name` parameters are provided. Only one of these parameters is allowed.
	FlowJson *string `json:"flow_json,omitempty"`
	// **Use for button type `FLOW`.** Either `navigate` or `data_exchange`. Defaults to `navigate`.
	FlowAction *string `json:"flow_action,omitempty"`
	// **Required if `flow_action` is `navigate`.** The unique ID of the Screen in the Flow.
	NavigateScreen *string `json:"navigate_screen,omitempty"`
}

// NewWhatsappTemplateComponentButton instantiates a new WhatsappTemplateComponentButton object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappTemplateComponentButton(type_ WhatsappTemplateComponentButtonType) *WhatsappTemplateComponentButton {
	this := WhatsappTemplateComponentButton{}
	this.Type = type_
	return &this
}

// NewWhatsappTemplateComponentButtonWithDefaults instantiates a new WhatsappTemplateComponentButton object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappTemplateComponentButtonWithDefaults() *WhatsappTemplateComponentButton {
	this := WhatsappTemplateComponentButton{}
	return &this
}

// GetType returns the Type field value
func (o *WhatsappTemplateComponentButton) GetType() WhatsappTemplateComponentButtonType {
	if o == nil {
		var ret WhatsappTemplateComponentButtonType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponentButton) GetTypeOk() (*WhatsappTemplateComponentButtonType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WhatsappTemplateComponentButton) SetType(v WhatsappTemplateComponentButtonType) {
	o.Type = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *WhatsappTemplateComponentButton) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponentButton) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *WhatsappTemplateComponentButton) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *WhatsappTemplateComponentButton) SetText(v string) {
	o.Text = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WhatsappTemplateComponentButton) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponentButton) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WhatsappTemplateComponentButton) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WhatsappTemplateComponentButton) SetUrl(v string) {
	o.Url = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *WhatsappTemplateComponentButton) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponentButton) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *WhatsappTemplateComponentButton) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *WhatsappTemplateComponentButton) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetOtpType returns the OtpType field value if set, zero value otherwise.
func (o *WhatsappTemplateComponentButton) GetOtpType() WhatsappTemplateComponentButtonOtpType {
	if o == nil || o.OtpType == nil {
		var ret WhatsappTemplateComponentButtonOtpType
		return ret
	}
	return *o.OtpType
}

// GetOtpTypeOk returns a tuple with the OtpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponentButton) GetOtpTypeOk() (*WhatsappTemplateComponentButtonOtpType, bool) {
	if o == nil || o.OtpType == nil {
		return nil, false
	}
	return o.OtpType, true
}

// HasOtpType returns a boolean if a field has been set.
func (o *WhatsappTemplateComponentButton) HasOtpType() bool {
	if o != nil && o.OtpType != nil {
		return true
	}

	return false
}

// SetOtpType gets a reference to the given WhatsappTemplateComponentButtonOtpType and assigns it to the OtpType field.
func (o *WhatsappTemplateComponentButton) SetOtpType(v WhatsappTemplateComponentButtonOtpType) {
	o.OtpType = &v
}

// GetAutofillText returns the AutofillText field value if set, zero value otherwise.
func (o *WhatsappTemplateComponentButton) GetAutofillText() string {
	if o == nil || o.AutofillText == nil {
		var ret string
		return ret
	}
	return *o.AutofillText
}

// GetAutofillTextOk returns a tuple with the AutofillText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponentButton) GetAutofillTextOk() (*string, bool) {
	if o == nil || o.AutofillText == nil {
		return nil, false
	}
	return o.AutofillText, true
}

// HasAutofillText returns a boolean if a field has been set.
func (o *WhatsappTemplateComponentButton) HasAutofillText() bool {
	if o != nil && o.AutofillText != nil {
		return true
	}

	return false
}

// SetAutofillText gets a reference to the given string and assigns it to the AutofillText field.
func (o *WhatsappTemplateComponentButton) SetAutofillText(v string) {
	o.AutofillText = &v
}

// GetPackageName returns the PackageName field value if set, zero value otherwise.
func (o *WhatsappTemplateComponentButton) GetPackageName() string {
	if o == nil || o.PackageName == nil {
		var ret string
		return ret
	}
	return *o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponentButton) GetPackageNameOk() (*string, bool) {
	if o == nil || o.PackageName == nil {
		return nil, false
	}
	return o.PackageName, true
}

// HasPackageName returns a boolean if a field has been set.
func (o *WhatsappTemplateComponentButton) HasPackageName() bool {
	if o != nil && o.PackageName != nil {
		return true
	}

	return false
}

// SetPackageName gets a reference to the given string and assigns it to the PackageName field.
func (o *WhatsappTemplateComponentButton) SetPackageName(v string) {
	o.PackageName = &v
}

// GetSignatureHash returns the SignatureHash field value if set, zero value otherwise.
func (o *WhatsappTemplateComponentButton) GetSignatureHash() string {
	if o == nil || o.SignatureHash == nil {
		var ret string
		return ret
	}
	return *o.SignatureHash
}

// GetSignatureHashOk returns a tuple with the SignatureHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponentButton) GetSignatureHashOk() (*string, bool) {
	if o == nil || o.SignatureHash == nil {
		return nil, false
	}
	return o.SignatureHash, true
}

// HasSignatureHash returns a boolean if a field has been set.
func (o *WhatsappTemplateComponentButton) HasSignatureHash() bool {
	if o != nil && o.SignatureHash != nil {
		return true
	}

	return false
}

// SetSignatureHash gets a reference to the given string and assigns it to the SignatureHash field.
func (o *WhatsappTemplateComponentButton) SetSignatureHash(v string) {
	o.SignatureHash = &v
}

// GetZeroTapTermsAccepted returns the ZeroTapTermsAccepted field value if set, zero value otherwise.
func (o *WhatsappTemplateComponentButton) GetZeroTapTermsAccepted() bool {
	if o == nil || o.ZeroTapTermsAccepted == nil {
		var ret bool
		return ret
	}
	return *o.ZeroTapTermsAccepted
}

// GetZeroTapTermsAcceptedOk returns a tuple with the ZeroTapTermsAccepted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponentButton) GetZeroTapTermsAcceptedOk() (*bool, bool) {
	if o == nil || o.ZeroTapTermsAccepted == nil {
		return nil, false
	}
	return o.ZeroTapTermsAccepted, true
}

// HasZeroTapTermsAccepted returns a boolean if a field has been set.
func (o *WhatsappTemplateComponentButton) HasZeroTapTermsAccepted() bool {
	if o != nil && o.ZeroTapTermsAccepted != nil {
		return true
	}

	return false
}

// SetZeroTapTermsAccepted gets a reference to the given bool and assigns it to the ZeroTapTermsAccepted field.
func (o *WhatsappTemplateComponentButton) SetZeroTapTermsAccepted(v bool) {
	o.ZeroTapTermsAccepted = &v
}

// GetExample returns the Example field value if set, zero value otherwise.
func (o *WhatsappTemplateComponentButton) GetExample() []string {
	if o == nil || o.Example == nil {
		var ret []string
		return ret
	}
	return o.Example
}

// GetExampleOk returns a tuple with the Example field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponentButton) GetExampleOk() ([]string, bool) {
	if o == nil || o.Example == nil {
		return nil, false
	}
	return o.Example, true
}

// HasExample returns a boolean if a field has been set.
func (o *WhatsappTemplateComponentButton) HasExample() bool {
	if o != nil && o.Example != nil {
		return true
	}

	return false
}

// SetExample gets a reference to the given []string and assigns it to the Example field.
func (o *WhatsappTemplateComponentButton) SetExample(v []string) {
	o.Example = v
}

// GetFlowId returns the FlowId field value if set, zero value otherwise.
func (o *WhatsappTemplateComponentButton) GetFlowId() string {
	if o == nil || o.FlowId == nil {
		var ret string
		return ret
	}
	return *o.FlowId
}

// GetFlowIdOk returns a tuple with the FlowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponentButton) GetFlowIdOk() (*string, bool) {
	if o == nil || o.FlowId == nil {
		return nil, false
	}
	return o.FlowId, true
}

// HasFlowId returns a boolean if a field has been set.
func (o *WhatsappTemplateComponentButton) HasFlowId() bool {
	if o != nil && o.FlowId != nil {
		return true
	}

	return false
}

// SetFlowId gets a reference to the given string and assigns it to the FlowId field.
func (o *WhatsappTemplateComponentButton) SetFlowId(v string) {
	o.FlowId = &v
}

// GetFlowName returns the FlowName field value if set, zero value otherwise.
func (o *WhatsappTemplateComponentButton) GetFlowName() string {
	if o == nil || o.FlowName == nil {
		var ret string
		return ret
	}
	return *o.FlowName
}

// GetFlowNameOk returns a tuple with the FlowName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponentButton) GetFlowNameOk() (*string, bool) {
	if o == nil || o.FlowName == nil {
		return nil, false
	}
	return o.FlowName, true
}

// HasFlowName returns a boolean if a field has been set.
func (o *WhatsappTemplateComponentButton) HasFlowName() bool {
	if o != nil && o.FlowName != nil {
		return true
	}

	return false
}

// SetFlowName gets a reference to the given string and assigns it to the FlowName field.
func (o *WhatsappTemplateComponentButton) SetFlowName(v string) {
	o.FlowName = &v
}

// GetFlowJson returns the FlowJson field value if set, zero value otherwise.
func (o *WhatsappTemplateComponentButton) GetFlowJson() string {
	if o == nil || o.FlowJson == nil {
		var ret string
		return ret
	}
	return *o.FlowJson
}

// GetFlowJsonOk returns a tuple with the FlowJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponentButton) GetFlowJsonOk() (*string, bool) {
	if o == nil || o.FlowJson == nil {
		return nil, false
	}
	return o.FlowJson, true
}

// HasFlowJson returns a boolean if a field has been set.
func (o *WhatsappTemplateComponentButton) HasFlowJson() bool {
	if o != nil && o.FlowJson != nil {
		return true
	}

	return false
}

// SetFlowJson gets a reference to the given string and assigns it to the FlowJson field.
func (o *WhatsappTemplateComponentButton) SetFlowJson(v string) {
	o.FlowJson = &v
}

// GetFlowAction returns the FlowAction field value if set, zero value otherwise.
func (o *WhatsappTemplateComponentButton) GetFlowAction() string {
	if o == nil || o.FlowAction == nil {
		var ret string
		return ret
	}
	return *o.FlowAction
}

// GetFlowActionOk returns a tuple with the FlowAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponentButton) GetFlowActionOk() (*string, bool) {
	if o == nil || o.FlowAction == nil {
		return nil, false
	}
	return o.FlowAction, true
}

// HasFlowAction returns a boolean if a field has been set.
func (o *WhatsappTemplateComponentButton) HasFlowAction() bool {
	if o != nil && o.FlowAction != nil {
		return true
	}

	return false
}

// SetFlowAction gets a reference to the given string and assigns it to the FlowAction field.
func (o *WhatsappTemplateComponentButton) SetFlowAction(v string) {
	o.FlowAction = &v
}

// GetNavigateScreen returns the NavigateScreen field value if set, zero value otherwise.
func (o *WhatsappTemplateComponentButton) GetNavigateScreen() string {
	if o == nil || o.NavigateScreen == nil {
		var ret string
		return ret
	}
	return *o.NavigateScreen
}

// GetNavigateScreenOk returns a tuple with the NavigateScreen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateComponentButton) GetNavigateScreenOk() (*string, bool) {
	if o == nil || o.NavigateScreen == nil {
		return nil, false
	}
	return o.NavigateScreen, true
}

// HasNavigateScreen returns a boolean if a field has been set.
func (o *WhatsappTemplateComponentButton) HasNavigateScreen() bool {
	if o != nil && o.NavigateScreen != nil {
		return true
	}

	return false
}

// SetNavigateScreen gets a reference to the given string and assigns it to the NavigateScreen field.
func (o *WhatsappTemplateComponentButton) SetNavigateScreen(v string) {
	o.NavigateScreen = &v
}

func (o WhatsappTemplateComponentButton) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.PhoneNumber != nil {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if o.OtpType != nil {
		toSerialize["otp_type"] = o.OtpType
	}
	if o.AutofillText != nil {
		toSerialize["autofill_text"] = o.AutofillText
	}
	if o.PackageName != nil {
		toSerialize["package_name"] = o.PackageName
	}
	if o.SignatureHash != nil {
		toSerialize["signature_hash"] = o.SignatureHash
	}
	if o.ZeroTapTermsAccepted != nil {
		toSerialize["zero_tap_terms_accepted"] = o.ZeroTapTermsAccepted
	}
	if o.Example != nil {
		toSerialize["example"] = o.Example
	}
	if o.FlowId != nil {
		toSerialize["flow_id"] = o.FlowId
	}
	if o.FlowName != nil {
		toSerialize["flow_name"] = o.FlowName
	}
	if o.FlowJson != nil {
		toSerialize["flow_json"] = o.FlowJson
	}
	if o.FlowAction != nil {
		toSerialize["flow_action"] = o.FlowAction
	}
	if o.NavigateScreen != nil {
		toSerialize["navigate_screen"] = o.NavigateScreen
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappTemplateComponentButton struct {
	value *WhatsappTemplateComponentButton
	isSet bool
}

func (v NullableWhatsappTemplateComponentButton) Get() *WhatsappTemplateComponentButton {
	return v.value
}

func (v *NullableWhatsappTemplateComponentButton) Set(val *WhatsappTemplateComponentButton) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappTemplateComponentButton) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappTemplateComponentButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappTemplateComponentButton(val *WhatsappTemplateComponentButton) *NullableWhatsappTemplateComponentButton {
	return &NullableWhatsappTemplateComponentButton{value: val, isSet: true}
}

func (v NullableWhatsappTemplateComponentButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappTemplateComponentButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


