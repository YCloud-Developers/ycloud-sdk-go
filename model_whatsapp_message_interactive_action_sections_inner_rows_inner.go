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

// WhatsappMessageInteractiveActionSectionsInnerRowsInner struct for WhatsappMessageInteractiveActionSectionsInnerRowsInner
type WhatsappMessageInteractiveActionSectionsInnerRowsInner struct {
	// Unique row ID.
	Id *string `json:"id,omitempty"`
	// Row title content.
	Title *string `json:"title,omitempty"`
	// Row description content.
	Description *string `json:"description,omitempty"`
}

// NewWhatsappMessageInteractiveActionSectionsInnerRowsInner instantiates a new WhatsappMessageInteractiveActionSectionsInnerRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappMessageInteractiveActionSectionsInnerRowsInner() *WhatsappMessageInteractiveActionSectionsInnerRowsInner {
	this := WhatsappMessageInteractiveActionSectionsInnerRowsInner{}
	return &this
}

// NewWhatsappMessageInteractiveActionSectionsInnerRowsInnerWithDefaults instantiates a new WhatsappMessageInteractiveActionSectionsInnerRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappMessageInteractiveActionSectionsInnerRowsInnerWithDefaults() *WhatsappMessageInteractiveActionSectionsInnerRowsInner {
	this := WhatsappMessageInteractiveActionSectionsInnerRowsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WhatsappMessageInteractiveActionSectionsInnerRowsInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageInteractiveActionSectionsInnerRowsInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WhatsappMessageInteractiveActionSectionsInnerRowsInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WhatsappMessageInteractiveActionSectionsInnerRowsInner) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *WhatsappMessageInteractiveActionSectionsInnerRowsInner) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageInteractiveActionSectionsInnerRowsInner) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *WhatsappMessageInteractiveActionSectionsInnerRowsInner) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *WhatsappMessageInteractiveActionSectionsInnerRowsInner) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WhatsappMessageInteractiveActionSectionsInnerRowsInner) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageInteractiveActionSectionsInnerRowsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WhatsappMessageInteractiveActionSectionsInnerRowsInner) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WhatsappMessageInteractiveActionSectionsInnerRowsInner) SetDescription(v string) {
	o.Description = &v
}

func (o WhatsappMessageInteractiveActionSectionsInnerRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappMessageInteractiveActionSectionsInnerRowsInner struct {
	value *WhatsappMessageInteractiveActionSectionsInnerRowsInner
	isSet bool
}

func (v NullableWhatsappMessageInteractiveActionSectionsInnerRowsInner) Get() *WhatsappMessageInteractiveActionSectionsInnerRowsInner {
	return v.value
}

func (v *NullableWhatsappMessageInteractiveActionSectionsInnerRowsInner) Set(val *WhatsappMessageInteractiveActionSectionsInnerRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappMessageInteractiveActionSectionsInnerRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappMessageInteractiveActionSectionsInnerRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappMessageInteractiveActionSectionsInnerRowsInner(val *WhatsappMessageInteractiveActionSectionsInnerRowsInner) *NullableWhatsappMessageInteractiveActionSectionsInnerRowsInner {
	return &NullableWhatsappMessageInteractiveActionSectionsInnerRowsInner{value: val, isSet: true}
}

func (v NullableWhatsappMessageInteractiveActionSectionsInnerRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappMessageInteractiveActionSectionsInnerRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

