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

// WhatsappCommerceSettingsUpdateRequest struct for WhatsappCommerceSettingsUpdateRequest
type WhatsappCommerceSettingsUpdateRequest struct {
	// When enabled, cart-related buttons appear in the conversation, catalog, and product details views. When the cart is disabled, customers can see products and their details, but all cart related buttons will not appear in any view.
	IsCartEnabled *bool `json:"isCartEnabled,omitempty"`
	// When enabled, the catalog storefront icon and catalog-related buttons appear in conversation and business profile views. When the catalog is disabled, the storefront icon and catalog-related buttons will not appear in any views and the catalog preview with thumbnails will not appear in the business profile view.
	IsCatalogVisible *bool `json:"isCatalogVisible,omitempty"`
}

// NewWhatsappCommerceSettingsUpdateRequest instantiates a new WhatsappCommerceSettingsUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappCommerceSettingsUpdateRequest() *WhatsappCommerceSettingsUpdateRequest {
	this := WhatsappCommerceSettingsUpdateRequest{}
	return &this
}

// NewWhatsappCommerceSettingsUpdateRequestWithDefaults instantiates a new WhatsappCommerceSettingsUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappCommerceSettingsUpdateRequestWithDefaults() *WhatsappCommerceSettingsUpdateRequest {
	this := WhatsappCommerceSettingsUpdateRequest{}
	return &this
}

// GetIsCartEnabled returns the IsCartEnabled field value if set, zero value otherwise.
func (o *WhatsappCommerceSettingsUpdateRequest) GetIsCartEnabled() bool {
	if o == nil || o.IsCartEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsCartEnabled
}

// GetIsCartEnabledOk returns a tuple with the IsCartEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappCommerceSettingsUpdateRequest) GetIsCartEnabledOk() (*bool, bool) {
	if o == nil || o.IsCartEnabled == nil {
		return nil, false
	}
	return o.IsCartEnabled, true
}

// HasIsCartEnabled returns a boolean if a field has been set.
func (o *WhatsappCommerceSettingsUpdateRequest) HasIsCartEnabled() bool {
	if o != nil && o.IsCartEnabled != nil {
		return true
	}

	return false
}

// SetIsCartEnabled gets a reference to the given bool and assigns it to the IsCartEnabled field.
func (o *WhatsappCommerceSettingsUpdateRequest) SetIsCartEnabled(v bool) {
	o.IsCartEnabled = &v
}

// GetIsCatalogVisible returns the IsCatalogVisible field value if set, zero value otherwise.
func (o *WhatsappCommerceSettingsUpdateRequest) GetIsCatalogVisible() bool {
	if o == nil || o.IsCatalogVisible == nil {
		var ret bool
		return ret
	}
	return *o.IsCatalogVisible
}

// GetIsCatalogVisibleOk returns a tuple with the IsCatalogVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappCommerceSettingsUpdateRequest) GetIsCatalogVisibleOk() (*bool, bool) {
	if o == nil || o.IsCatalogVisible == nil {
		return nil, false
	}
	return o.IsCatalogVisible, true
}

// HasIsCatalogVisible returns a boolean if a field has been set.
func (o *WhatsappCommerceSettingsUpdateRequest) HasIsCatalogVisible() bool {
	if o != nil && o.IsCatalogVisible != nil {
		return true
	}

	return false
}

// SetIsCatalogVisible gets a reference to the given bool and assigns it to the IsCatalogVisible field.
func (o *WhatsappCommerceSettingsUpdateRequest) SetIsCatalogVisible(v bool) {
	o.IsCatalogVisible = &v
}

func (o WhatsappCommerceSettingsUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsCartEnabled != nil {
		toSerialize["isCartEnabled"] = o.IsCartEnabled
	}
	if o.IsCatalogVisible != nil {
		toSerialize["isCatalogVisible"] = o.IsCatalogVisible
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappCommerceSettingsUpdateRequest struct {
	value *WhatsappCommerceSettingsUpdateRequest
	isSet bool
}

func (v NullableWhatsappCommerceSettingsUpdateRequest) Get() *WhatsappCommerceSettingsUpdateRequest {
	return v.value
}

func (v *NullableWhatsappCommerceSettingsUpdateRequest) Set(val *WhatsappCommerceSettingsUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappCommerceSettingsUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappCommerceSettingsUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappCommerceSettingsUpdateRequest(val *WhatsappCommerceSettingsUpdateRequest) *NullableWhatsappCommerceSettingsUpdateRequest {
	return &NullableWhatsappCommerceSettingsUpdateRequest{value: val, isSet: true}
}

func (v NullableWhatsappCommerceSettingsUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappCommerceSettingsUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}