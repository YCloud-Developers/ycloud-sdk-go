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

// CustomEventDefinitionPropertyCreateRequest Contains the properties of the custom event property definition to be created.
type CustomEventDefinitionPropertyCreateRequest struct {
	// The unique name of the custom property.
	Name string `json:"name"`
	// The label of the property.
	Label string `json:"label"`
	// The description of the property.
	Description *string `json:"description,omitempty"`
	// Type of the property. - `STRING`: Indicates a property that receives plain text strings. - `NUMBER`: Indicates a property that receives numeric values with up to one decimal. - `TIMESTAMP`: Indicates a property that receives epoch millisecond. - `URL`: Indicates a property that receives URLs, formatted as strings starting with `http://` or `https://`.
	Type string `json:"type"`
}

// NewCustomEventDefinitionPropertyCreateRequest instantiates a new CustomEventDefinitionPropertyCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomEventDefinitionPropertyCreateRequest(name string, label string, type_ string) *CustomEventDefinitionPropertyCreateRequest {
	this := CustomEventDefinitionPropertyCreateRequest{}
	this.Name = name
	this.Label = label
	this.Type = type_
	return &this
}

// NewCustomEventDefinitionPropertyCreateRequestWithDefaults instantiates a new CustomEventDefinitionPropertyCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomEventDefinitionPropertyCreateRequestWithDefaults() *CustomEventDefinitionPropertyCreateRequest {
	this := CustomEventDefinitionPropertyCreateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CustomEventDefinitionPropertyCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomEventDefinitionPropertyCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomEventDefinitionPropertyCreateRequest) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value
func (o *CustomEventDefinitionPropertyCreateRequest) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *CustomEventDefinitionPropertyCreateRequest) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *CustomEventDefinitionPropertyCreateRequest) SetLabel(v string) {
	o.Label = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CustomEventDefinitionPropertyCreateRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomEventDefinitionPropertyCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomEventDefinitionPropertyCreateRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CustomEventDefinitionPropertyCreateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *CustomEventDefinitionPropertyCreateRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomEventDefinitionPropertyCreateRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomEventDefinitionPropertyCreateRequest) SetType(v string) {
	o.Type = v
}

func (o CustomEventDefinitionPropertyCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["label"] = o.Label
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableCustomEventDefinitionPropertyCreateRequest struct {
	value *CustomEventDefinitionPropertyCreateRequest
	isSet bool
}

func (v NullableCustomEventDefinitionPropertyCreateRequest) Get() *CustomEventDefinitionPropertyCreateRequest {
	return v.value
}

func (v *NullableCustomEventDefinitionPropertyCreateRequest) Set(val *CustomEventDefinitionPropertyCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomEventDefinitionPropertyCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomEventDefinitionPropertyCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomEventDefinitionPropertyCreateRequest(val *CustomEventDefinitionPropertyCreateRequest) *NullableCustomEventDefinitionPropertyCreateRequest {
	return &NullableCustomEventDefinitionPropertyCreateRequest{value: val, isSet: true}
}

func (v NullableCustomEventDefinitionPropertyCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomEventDefinitionPropertyCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
