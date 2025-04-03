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

// PageCursor A cursor object is returned only if the endpoint you requested supports cursor pagination.
type PageCursor struct {
	// A cursor to fetch the next page in cursor pagination. For example, if you make a list request, receive 100 objects and `cursor.after=id:foo`, your subsequent call can include `pageAfter=id:foo` in order to fetch the next page of the list. This field is returned only if there are more items in the list.
	After *string `json:"after,omitempty"`
}

// NewPageCursor instantiates a new PageCursor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageCursor() *PageCursor {
	this := PageCursor{}
	return &this
}

// NewPageCursorWithDefaults instantiates a new PageCursor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageCursorWithDefaults() *PageCursor {
	this := PageCursor{}
	return &this
}

// GetAfter returns the After field value if set, zero value otherwise.
func (o *PageCursor) GetAfter() string {
	if o == nil || o.After == nil {
		var ret string
		return ret
	}
	return *o.After
}

// GetAfterOk returns a tuple with the After field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageCursor) GetAfterOk() (*string, bool) {
	if o == nil || o.After == nil {
		return nil, false
	}
	return o.After, true
}

// HasAfter returns a boolean if a field has been set.
func (o *PageCursor) HasAfter() bool {
	if o != nil && o.After != nil {
		return true
	}

	return false
}

// SetAfter gets a reference to the given string and assigns it to the After field.
func (o *PageCursor) SetAfter(v string) {
	o.After = &v
}

func (o PageCursor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.After != nil {
		toSerialize["after"] = o.After
	}
	return json.Marshal(toSerialize)
}

type NullablePageCursor struct {
	value *PageCursor
	isSet bool
}

func (v NullablePageCursor) Get() *PageCursor {
	return v.value
}

func (v *NullablePageCursor) Set(val *PageCursor) {
	v.value = val
	v.isSet = true
}

func (v NullablePageCursor) IsSet() bool {
	return v.isSet
}

func (v *NullablePageCursor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageCursor(val *PageCursor) *NullablePageCursor {
	return &NullablePageCursor{value: val, isSet: true}
}

func (v NullablePageCursor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageCursor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


