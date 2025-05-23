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

// UnsubscriberPage Represents a given page of unsubscriber objects.
type UnsubscriberPage struct {
	// An array containing unsubscriber objects.
	Items []Unsubscriber `json:"items,omitempty"`
	Cursor *PageCursor `json:"cursor,omitempty"`
	// The position of the item this page starts from, zero-based. e.g., the 11th item is at offset 10.
	Offset int32 `json:"offset"`
	// A limit on the number of items to be returned, between 1 and 100, defaults to 10.
	Limit int32 `json:"limit"`
	// The actual number of items in the page.
	Length int32 `json:"length"`
	// The total number of items. This field is returned only when the request parameter `includeTotal` is set to `true`.
	Total *int32 `json:"total,omitempty"`
}

// NewUnsubscriberPage instantiates a new UnsubscriberPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnsubscriberPage(offset int32, limit int32, length int32) *UnsubscriberPage {
	this := UnsubscriberPage{}
	this.Offset = offset
	this.Limit = limit
	this.Length = length
	return &this
}

// NewUnsubscriberPageWithDefaults instantiates a new UnsubscriberPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnsubscriberPageWithDefaults() *UnsubscriberPage {
	this := UnsubscriberPage{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *UnsubscriberPage) GetItems() []Unsubscriber {
	if o == nil || o.Items == nil {
		var ret []Unsubscriber
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnsubscriberPage) GetItemsOk() ([]Unsubscriber, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *UnsubscriberPage) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Unsubscriber and assigns it to the Items field.
func (o *UnsubscriberPage) SetItems(v []Unsubscriber) {
	o.Items = v
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *UnsubscriberPage) GetCursor() PageCursor {
	if o == nil || o.Cursor == nil {
		var ret PageCursor
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnsubscriberPage) GetCursorOk() (*PageCursor, bool) {
	if o == nil || o.Cursor == nil {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *UnsubscriberPage) HasCursor() bool {
	if o != nil && o.Cursor != nil {
		return true
	}

	return false
}

// SetCursor gets a reference to the given PageCursor and assigns it to the Cursor field.
func (o *UnsubscriberPage) SetCursor(v PageCursor) {
	o.Cursor = &v
}

// GetOffset returns the Offset field value
func (o *UnsubscriberPage) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *UnsubscriberPage) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *UnsubscriberPage) SetOffset(v int32) {
	o.Offset = v
}

// GetLimit returns the Limit field value
func (o *UnsubscriberPage) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *UnsubscriberPage) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *UnsubscriberPage) SetLimit(v int32) {
	o.Limit = v
}

// GetLength returns the Length field value
func (o *UnsubscriberPage) GetLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Length
}

// GetLengthOk returns a tuple with the Length field value
// and a boolean to check if the value has been set.
func (o *UnsubscriberPage) GetLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Length, true
}

// SetLength sets field value
func (o *UnsubscriberPage) SetLength(v int32) {
	o.Length = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *UnsubscriberPage) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnsubscriberPage) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *UnsubscriberPage) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *UnsubscriberPage) SetTotal(v int32) {
	o.Total = &v
}

func (o UnsubscriberPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Cursor != nil {
		toSerialize["cursor"] = o.Cursor
	}
	if true {
		toSerialize["offset"] = o.Offset
	}
	if true {
		toSerialize["limit"] = o.Limit
	}
	if true {
		toSerialize["length"] = o.Length
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableUnsubscriberPage struct {
	value *UnsubscriberPage
	isSet bool
}

func (v NullableUnsubscriberPage) Get() *UnsubscriberPage {
	return v.value
}

func (v *NullableUnsubscriberPage) Set(val *UnsubscriberPage) {
	v.value = val
	v.isSet = true
}

func (v NullableUnsubscriberPage) IsSet() bool {
	return v.isSet
}

func (v *NullableUnsubscriberPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnsubscriberPage(val *UnsubscriberPage) *NullableUnsubscriberPage {
	return &NullableUnsubscriberPage{value: val, isSet: true}
}

func (v NullableUnsubscriberPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnsubscriberPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


