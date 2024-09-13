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

// WhatsappTemplateEditRequest The request body to edit a WhatsApp template.
type WhatsappTemplateEditRequest struct {
	Components []WhatsappTemplateComponent `json:"components"`
	// **Use only for template category is `AUTHENTICATION` or `UTILITY`.** If we are unable to deliver a message for an amount of time that exceeds its time-to-live, we will stop retrying and drop the message. By default, messages that use an authentication template have a default TTL of **10 minutes**, and messages that use a utility template have a default TTL of **30 days**. Set its value between `60` and `600` seconds (i.e., 1 to 10 minutes) for authentication templates, or `60` and `3600` seconds (i.e., 1 to 60 minutes) for utility templates. Alternatively, you can set this value to `-1`, which will set a custom TTL of 30 days for either type of template. We encourage you to set a time-to-live for all of your authentication templates, preferably equal to or less than your code expiration time, to ensure your customers only get a message when a code is still usable. Authentication templates created before October 23, 2024, have a default TTL of 30 days.
	MessageSendTtlSeconds *int32 `json:"messageSendTtlSeconds,omitempty"`
}

// NewWhatsappTemplateEditRequest instantiates a new WhatsappTemplateEditRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappTemplateEditRequest(components []WhatsappTemplateComponent) *WhatsappTemplateEditRequest {
	this := WhatsappTemplateEditRequest{}
	this.Components = components
	return &this
}

// NewWhatsappTemplateEditRequestWithDefaults instantiates a new WhatsappTemplateEditRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappTemplateEditRequestWithDefaults() *WhatsappTemplateEditRequest {
	this := WhatsappTemplateEditRequest{}
	return &this
}

// GetComponents returns the Components field value
func (o *WhatsappTemplateEditRequest) GetComponents() []WhatsappTemplateComponent {
	if o == nil {
		var ret []WhatsappTemplateComponent
		return ret
	}

	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateEditRequest) GetComponentsOk() ([]WhatsappTemplateComponent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Components, true
}

// SetComponents sets field value
func (o *WhatsappTemplateEditRequest) SetComponents(v []WhatsappTemplateComponent) {
	o.Components = v
}

// GetMessageSendTtlSeconds returns the MessageSendTtlSeconds field value if set, zero value otherwise.
func (o *WhatsappTemplateEditRequest) GetMessageSendTtlSeconds() int32 {
	if o == nil || o.MessageSendTtlSeconds == nil {
		var ret int32
		return ret
	}
	return *o.MessageSendTtlSeconds
}

// GetMessageSendTtlSecondsOk returns a tuple with the MessageSendTtlSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappTemplateEditRequest) GetMessageSendTtlSecondsOk() (*int32, bool) {
	if o == nil || o.MessageSendTtlSeconds == nil {
		return nil, false
	}
	return o.MessageSendTtlSeconds, true
}

// HasMessageSendTtlSeconds returns a boolean if a field has been set.
func (o *WhatsappTemplateEditRequest) HasMessageSendTtlSeconds() bool {
	if o != nil && o.MessageSendTtlSeconds != nil {
		return true
	}

	return false
}

// SetMessageSendTtlSeconds gets a reference to the given int32 and assigns it to the MessageSendTtlSeconds field.
func (o *WhatsappTemplateEditRequest) SetMessageSendTtlSeconds(v int32) {
	o.MessageSendTtlSeconds = &v
}

func (o WhatsappTemplateEditRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["components"] = o.Components
	}
	if o.MessageSendTtlSeconds != nil {
		toSerialize["messageSendTtlSeconds"] = o.MessageSendTtlSeconds
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappTemplateEditRequest struct {
	value *WhatsappTemplateEditRequest
	isSet bool
}

func (v NullableWhatsappTemplateEditRequest) Get() *WhatsappTemplateEditRequest {
	return v.value
}

func (v *NullableWhatsappTemplateEditRequest) Set(val *WhatsappTemplateEditRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappTemplateEditRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappTemplateEditRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappTemplateEditRequest(val *WhatsappTemplateEditRequest) *NullableWhatsappTemplateEditRequest {
	return &NullableWhatsappTemplateEditRequest{value: val, isSet: true}
}

func (v NullableWhatsappTemplateEditRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappTemplateEditRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
