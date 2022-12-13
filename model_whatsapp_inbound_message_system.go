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

// WhatsappInboundMessageSystem When the message type is set to `system`, this field is included. This object is added to Webhooks if a user has changed their phone number and if a user’s identity has potentially changed on WhatsApp.
type WhatsappInboundMessageSystem struct {
	// Describes the system message event. Supported use cases are: - Phone number update: for when a user changes from an old number to a new number. - Identity update: for when a user identity has changed.
	Body *string `json:"body,omitempty"`
	// **Added to Webhooks for phone number updates.**  New WhatsApp ID of the customer.
	NewWaId *string `json:"new_wa_id,omitempty"`
	// **Added to Webhooks for identity updates.**  New WhatsApp ID of the customer.
	Identity *string `json:"identity,omitempty"`
	// Supported types are: - `user_changed_number`: for a user changed number notification. - `user_identity_changed`: for user identity changed notification.
	Type *string `json:"type,omitempty"`
	// **Added to Webhooks for identity updates.**  The new WhatsApp user ID of the customer.
	User *string `json:"user,omitempty"`
}

// NewWhatsappInboundMessageSystem instantiates a new WhatsappInboundMessageSystem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappInboundMessageSystem() *WhatsappInboundMessageSystem {
	this := WhatsappInboundMessageSystem{}
	return &this
}

// NewWhatsappInboundMessageSystemWithDefaults instantiates a new WhatsappInboundMessageSystem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappInboundMessageSystemWithDefaults() *WhatsappInboundMessageSystem {
	this := WhatsappInboundMessageSystem{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *WhatsappInboundMessageSystem) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappInboundMessageSystem) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *WhatsappInboundMessageSystem) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *WhatsappInboundMessageSystem) SetBody(v string) {
	o.Body = &v
}

// GetNewWaId returns the NewWaId field value if set, zero value otherwise.
func (o *WhatsappInboundMessageSystem) GetNewWaId() string {
	if o == nil || o.NewWaId == nil {
		var ret string
		return ret
	}
	return *o.NewWaId
}

// GetNewWaIdOk returns a tuple with the NewWaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappInboundMessageSystem) GetNewWaIdOk() (*string, bool) {
	if o == nil || o.NewWaId == nil {
		return nil, false
	}
	return o.NewWaId, true
}

// HasNewWaId returns a boolean if a field has been set.
func (o *WhatsappInboundMessageSystem) HasNewWaId() bool {
	if o != nil && o.NewWaId != nil {
		return true
	}

	return false
}

// SetNewWaId gets a reference to the given string and assigns it to the NewWaId field.
func (o *WhatsappInboundMessageSystem) SetNewWaId(v string) {
	o.NewWaId = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *WhatsappInboundMessageSystem) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappInboundMessageSystem) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *WhatsappInboundMessageSystem) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *WhatsappInboundMessageSystem) SetIdentity(v string) {
	o.Identity = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WhatsappInboundMessageSystem) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappInboundMessageSystem) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WhatsappInboundMessageSystem) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WhatsappInboundMessageSystem) SetType(v string) {
	o.Type = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *WhatsappInboundMessageSystem) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappInboundMessageSystem) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *WhatsappInboundMessageSystem) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *WhatsappInboundMessageSystem) SetUser(v string) {
	o.User = &v
}

func (o WhatsappInboundMessageSystem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.NewWaId != nil {
		toSerialize["new_wa_id"] = o.NewWaId
	}
	if o.Identity != nil {
		toSerialize["identity"] = o.Identity
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappInboundMessageSystem struct {
	value *WhatsappInboundMessageSystem
	isSet bool
}

func (v NullableWhatsappInboundMessageSystem) Get() *WhatsappInboundMessageSystem {
	return v.value
}

func (v *NullableWhatsappInboundMessageSystem) Set(val *WhatsappInboundMessageSystem) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappInboundMessageSystem) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappInboundMessageSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappInboundMessageSystem(val *WhatsappInboundMessageSystem) *NullableWhatsappInboundMessageSystem {
	return &NullableWhatsappInboundMessageSystem{value: val, isSet: true}
}

func (v NullableWhatsappInboundMessageSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappInboundMessageSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


