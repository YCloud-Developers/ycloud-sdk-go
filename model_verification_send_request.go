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

// VerificationSendRequest struct for VerificationSendRequest
type VerificationSendRequest struct {
	Channel VerificationChannel `json:"channel"`
	// The recipient's phone number or email address depending on `channel`. - Phone number: In [E.164](https://en.wikipedia.org/wiki/E.164) format. Applicable when `channel` is `sms` or `voice`. - Email address: For example, `tom@example.com`. Applicable when `channel` is `email_code`.
	To string `json:"to"`
	// Verification code to be sent. This field is optional. If not provided, we will automatically generate a code.
	Code *string `json:"code,omitempty"`
	// [Sender ID](https://help.ycloud.com/en/articles/3080386) to be used.
	SenderId *string `json:"senderId,omitempty"`
	// This parameter is only required for Chinese mainland SMS messages. You must specify an approved signature such as `Brand`. It will be added to the beginning of SMS body and wrapped with `【】`, e.g. `【Brand】Your verification code is 123456`.
	Signature *string `json:"signature,omitempty"`
	// [ISO 639 Language Code](https://www.iso.org/iso-639-language-codes.html). If not specified, language will be set as `en` by default. Notably, in certain countries or regions, language will be automatically set as the local language due to the regional restrictions. Applicable languages: `ar`: Arabic `de`: German `en`: English `es`: Spanish `fr`: French `id`: Indonesian `it`: Italian `pt_BR`: Portuguese `ru`: Russian `tr`: Turkish `vi`: Vietnamese `zh_CN`: Simplified Chinese `zh_HK`: Traditional Chinese
	Language *string `json:"language,omitempty"`
	// A unique (recommended) string to reference the object. This can be an order number or similar, and can be used to reconcile the object with your internal systems. If present, this value will also be attached to the `externalId` of message objects.
	ExternalId *string `json:"externalId,omitempty"`
}

// NewVerificationSendRequest instantiates a new VerificationSendRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationSendRequest(channel VerificationChannel, to string) *VerificationSendRequest {
	this := VerificationSendRequest{}
	this.Channel = channel
	this.To = to
	return &this
}

// NewVerificationSendRequestWithDefaults instantiates a new VerificationSendRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationSendRequestWithDefaults() *VerificationSendRequest {
	this := VerificationSendRequest{}
	return &this
}

// GetChannel returns the Channel field value
func (o *VerificationSendRequest) GetChannel() VerificationChannel {
	if o == nil {
		var ret VerificationChannel
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *VerificationSendRequest) GetChannelOk() (*VerificationChannel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *VerificationSendRequest) SetChannel(v VerificationChannel) {
	o.Channel = v
}

// GetTo returns the To field value
func (o *VerificationSendRequest) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *VerificationSendRequest) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *VerificationSendRequest) SetTo(v string) {
	o.To = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *VerificationSendRequest) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationSendRequest) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *VerificationSendRequest) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *VerificationSendRequest) SetCode(v string) {
	o.Code = &v
}

// GetSenderId returns the SenderId field value if set, zero value otherwise.
func (o *VerificationSendRequest) GetSenderId() string {
	if o == nil || o.SenderId == nil {
		var ret string
		return ret
	}
	return *o.SenderId
}

// GetSenderIdOk returns a tuple with the SenderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationSendRequest) GetSenderIdOk() (*string, bool) {
	if o == nil || o.SenderId == nil {
		return nil, false
	}
	return o.SenderId, true
}

// HasSenderId returns a boolean if a field has been set.
func (o *VerificationSendRequest) HasSenderId() bool {
	if o != nil && o.SenderId != nil {
		return true
	}

	return false
}

// SetSenderId gets a reference to the given string and assigns it to the SenderId field.
func (o *VerificationSendRequest) SetSenderId(v string) {
	o.SenderId = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *VerificationSendRequest) GetSignature() string {
	if o == nil || o.Signature == nil {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationSendRequest) GetSignatureOk() (*string, bool) {
	if o == nil || o.Signature == nil {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *VerificationSendRequest) HasSignature() bool {
	if o != nil && o.Signature != nil {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *VerificationSendRequest) SetSignature(v string) {
	o.Signature = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *VerificationSendRequest) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationSendRequest) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *VerificationSendRequest) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *VerificationSendRequest) SetLanguage(v string) {
	o.Language = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *VerificationSendRequest) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationSendRequest) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *VerificationSendRequest) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *VerificationSendRequest) SetExternalId(v string) {
	o.ExternalId = &v
}

func (o VerificationSendRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["channel"] = o.Channel
	}
	if true {
		toSerialize["to"] = o.To
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.SenderId != nil {
		toSerialize["senderId"] = o.SenderId
	}
	if o.Signature != nil {
		toSerialize["signature"] = o.Signature
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}
	return json.Marshal(toSerialize)
}

type NullableVerificationSendRequest struct {
	value *VerificationSendRequest
	isSet bool
}

func (v NullableVerificationSendRequest) Get() *VerificationSendRequest {
	return v.value
}

func (v *NullableVerificationSendRequest) Set(val *VerificationSendRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationSendRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationSendRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationSendRequest(val *VerificationSendRequest) *NullableVerificationSendRequest {
	return &NullableVerificationSendRequest{value: val, isSet: true}
}

func (v NullableVerificationSendRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationSendRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


