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

// CustomEventDefinitionPropertyUpdateRequest Contains the properties of the event property definition to be updated.
type CustomEventDefinitionPropertyUpdateRequest struct {
	// The label of the event property definition.
	Label *string `json:"label,omitempty"`
	// The description of the event property definition.
	Description *string `json:"description,omitempty"`
}

// NewCustomEventDefinitionPropertyUpdateRequest instantiates a new CustomEventDefinitionPropertyUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomEventDefinitionPropertyUpdateRequest() *CustomEventDefinitionPropertyUpdateRequest {
	this := CustomEventDefinitionPropertyUpdateRequest{}
	return &this
}

// NewCustomEventDefinitionPropertyUpdateRequestWithDefaults instantiates a new CustomEventDefinitionPropertyUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomEventDefinitionPropertyUpdateRequestWithDefaults() *CustomEventDefinitionPropertyUpdateRequest {
	this := CustomEventDefinitionPropertyUpdateRequest{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CustomEventDefinitionPropertyUpdateRequest) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEventDefinitionPropertyUpdateRequest) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CustomEventDefinitionPropertyUpdateRequest) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *CustomEventDefinitionPropertyUpdateRequest) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CustomEventDefinitionPropertyUpdateRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEventDefinitionPropertyUpdateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomEventDefinitionPropertyUpdateRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CustomEventDefinitionPropertyUpdateRequest) SetDescription(v string) {
	o.Description = &v
}

func (o CustomEventDefinitionPropertyUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableCustomEventDefinitionPropertyUpdateRequest struct {
	value *CustomEventDefinitionPropertyUpdateRequest
	isSet bool
}

func (v NullableCustomEventDefinitionPropertyUpdateRequest) Get() *CustomEventDefinitionPropertyUpdateRequest {
	return v.value
}

func (v *NullableCustomEventDefinitionPropertyUpdateRequest) Set(val *CustomEventDefinitionPropertyUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomEventDefinitionPropertyUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomEventDefinitionPropertyUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomEventDefinitionPropertyUpdateRequest(val *CustomEventDefinitionPropertyUpdateRequest) *NullableCustomEventDefinitionPropertyUpdateRequest {
	return &NullableCustomEventDefinitionPropertyUpdateRequest{value: val, isSet: true}
}

func (v NullableCustomEventDefinitionPropertyUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomEventDefinitionPropertyUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}