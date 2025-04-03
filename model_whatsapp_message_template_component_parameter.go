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

// WhatsappMessageTemplateComponentParameter struct for WhatsappMessageTemplateComponentParameter
type WhatsappMessageTemplateComponentParameter struct {
	// **Required.** Component parameter type. - `text`: Used when the template component type is `BODY`, or the `HEADER` component format is `TEXT`. - `image`: Used when the template `HEADER` component is `IMAGE`. - `video`: Used when the template `HEADER` component is `VIDEO`. - `document`: Used when the template `HEADER` component is `DOCUMENT`. - `payload`: Used when the template component button type is `QUICK_REPLY`. - `coupon_code`: Used when the template component button type is `COPY_CODE`. - `limited_time_offer`: Used when the template component type is `LIMITED_TIME_OFFER`. - `action`: Used when the template component button type is `CATALOG`, `MPM`, `FLOW`, or `ORDER_DETAILS`. - `order_status`: Used when the template subcategory is `ORDER_STATUS`. - `location`: Used when the template `HEADER` component is `LOCATION`.
	Type *string `json:"type,omitempty"`
	// **Required when `type` = `text`.** The message's text. For the header component, the character limit is 60 characters. For the body component, the character limit is 1024 characters. For url buttons, it indicates the developer-provided suffix that is appended to the predefined prefix URL in the template.
	Text *string `json:"text,omitempty"`
	// Required for `quick_reply` buttons. Developer-defined payload that is returned when the button is clicked in addition to the display text on the button.
	Payload *string `json:"payload,omitempty"`
	// **Required when `type` = `coupon_code`.** The coupon code to be copied when the customer taps the button.
	CouponCode *string `json:"coupon_code,omitempty"`
	Image *WhatsappMessageMedia `json:"image,omitempty"`
	Video *WhatsappMessageMedia `json:"video,omitempty"`
	Document *WhatsappMessageMedia `json:"document,omitempty"`
	LimitedTimeOffer *WhatsappMessageTemplateComponentParameterLimitedTimeOffer `json:"limited_time_offer,omitempty"`
	Action *WhatsappMessageTemplateComponentParameterAction `json:"action,omitempty"`
	OrderStatus *WhatsappMessageOrderStatus `json:"order_status,omitempty"`
	Location *WhatsappMessageLocation `json:"location,omitempty"`
}

// NewWhatsappMessageTemplateComponentParameter instantiates a new WhatsappMessageTemplateComponentParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsappMessageTemplateComponentParameter() *WhatsappMessageTemplateComponentParameter {
	this := WhatsappMessageTemplateComponentParameter{}
	return &this
}

// NewWhatsappMessageTemplateComponentParameterWithDefaults instantiates a new WhatsappMessageTemplateComponentParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsappMessageTemplateComponentParameterWithDefaults() *WhatsappMessageTemplateComponentParameter {
	this := WhatsappMessageTemplateComponentParameter{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WhatsappMessageTemplateComponentParameter) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageTemplateComponentParameter) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WhatsappMessageTemplateComponentParameter) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WhatsappMessageTemplateComponentParameter) SetType(v string) {
	o.Type = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *WhatsappMessageTemplateComponentParameter) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageTemplateComponentParameter) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *WhatsappMessageTemplateComponentParameter) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *WhatsappMessageTemplateComponentParameter) SetText(v string) {
	o.Text = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *WhatsappMessageTemplateComponentParameter) GetPayload() string {
	if o == nil || o.Payload == nil {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageTemplateComponentParameter) GetPayloadOk() (*string, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *WhatsappMessageTemplateComponentParameter) HasPayload() bool {
	if o != nil && o.Payload != nil {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *WhatsappMessageTemplateComponentParameter) SetPayload(v string) {
	o.Payload = &v
}

// GetCouponCode returns the CouponCode field value if set, zero value otherwise.
func (o *WhatsappMessageTemplateComponentParameter) GetCouponCode() string {
	if o == nil || o.CouponCode == nil {
		var ret string
		return ret
	}
	return *o.CouponCode
}

// GetCouponCodeOk returns a tuple with the CouponCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageTemplateComponentParameter) GetCouponCodeOk() (*string, bool) {
	if o == nil || o.CouponCode == nil {
		return nil, false
	}
	return o.CouponCode, true
}

// HasCouponCode returns a boolean if a field has been set.
func (o *WhatsappMessageTemplateComponentParameter) HasCouponCode() bool {
	if o != nil && o.CouponCode != nil {
		return true
	}

	return false
}

// SetCouponCode gets a reference to the given string and assigns it to the CouponCode field.
func (o *WhatsappMessageTemplateComponentParameter) SetCouponCode(v string) {
	o.CouponCode = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *WhatsappMessageTemplateComponentParameter) GetImage() WhatsappMessageMedia {
	if o == nil || o.Image == nil {
		var ret WhatsappMessageMedia
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageTemplateComponentParameter) GetImageOk() (*WhatsappMessageMedia, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *WhatsappMessageTemplateComponentParameter) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given WhatsappMessageMedia and assigns it to the Image field.
func (o *WhatsappMessageTemplateComponentParameter) SetImage(v WhatsappMessageMedia) {
	o.Image = &v
}

// GetVideo returns the Video field value if set, zero value otherwise.
func (o *WhatsappMessageTemplateComponentParameter) GetVideo() WhatsappMessageMedia {
	if o == nil || o.Video == nil {
		var ret WhatsappMessageMedia
		return ret
	}
	return *o.Video
}

// GetVideoOk returns a tuple with the Video field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageTemplateComponentParameter) GetVideoOk() (*WhatsappMessageMedia, bool) {
	if o == nil || o.Video == nil {
		return nil, false
	}
	return o.Video, true
}

// HasVideo returns a boolean if a field has been set.
func (o *WhatsappMessageTemplateComponentParameter) HasVideo() bool {
	if o != nil && o.Video != nil {
		return true
	}

	return false
}

// SetVideo gets a reference to the given WhatsappMessageMedia and assigns it to the Video field.
func (o *WhatsappMessageTemplateComponentParameter) SetVideo(v WhatsappMessageMedia) {
	o.Video = &v
}

// GetDocument returns the Document field value if set, zero value otherwise.
func (o *WhatsappMessageTemplateComponentParameter) GetDocument() WhatsappMessageMedia {
	if o == nil || o.Document == nil {
		var ret WhatsappMessageMedia
		return ret
	}
	return *o.Document
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageTemplateComponentParameter) GetDocumentOk() (*WhatsappMessageMedia, bool) {
	if o == nil || o.Document == nil {
		return nil, false
	}
	return o.Document, true
}

// HasDocument returns a boolean if a field has been set.
func (o *WhatsappMessageTemplateComponentParameter) HasDocument() bool {
	if o != nil && o.Document != nil {
		return true
	}

	return false
}

// SetDocument gets a reference to the given WhatsappMessageMedia and assigns it to the Document field.
func (o *WhatsappMessageTemplateComponentParameter) SetDocument(v WhatsappMessageMedia) {
	o.Document = &v
}

// GetLimitedTimeOffer returns the LimitedTimeOffer field value if set, zero value otherwise.
func (o *WhatsappMessageTemplateComponentParameter) GetLimitedTimeOffer() WhatsappMessageTemplateComponentParameterLimitedTimeOffer {
	if o == nil || o.LimitedTimeOffer == nil {
		var ret WhatsappMessageTemplateComponentParameterLimitedTimeOffer
		return ret
	}
	return *o.LimitedTimeOffer
}

// GetLimitedTimeOfferOk returns a tuple with the LimitedTimeOffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageTemplateComponentParameter) GetLimitedTimeOfferOk() (*WhatsappMessageTemplateComponentParameterLimitedTimeOffer, bool) {
	if o == nil || o.LimitedTimeOffer == nil {
		return nil, false
	}
	return o.LimitedTimeOffer, true
}

// HasLimitedTimeOffer returns a boolean if a field has been set.
func (o *WhatsappMessageTemplateComponentParameter) HasLimitedTimeOffer() bool {
	if o != nil && o.LimitedTimeOffer != nil {
		return true
	}

	return false
}

// SetLimitedTimeOffer gets a reference to the given WhatsappMessageTemplateComponentParameterLimitedTimeOffer and assigns it to the LimitedTimeOffer field.
func (o *WhatsappMessageTemplateComponentParameter) SetLimitedTimeOffer(v WhatsappMessageTemplateComponentParameterLimitedTimeOffer) {
	o.LimitedTimeOffer = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *WhatsappMessageTemplateComponentParameter) GetAction() WhatsappMessageTemplateComponentParameterAction {
	if o == nil || o.Action == nil {
		var ret WhatsappMessageTemplateComponentParameterAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageTemplateComponentParameter) GetActionOk() (*WhatsappMessageTemplateComponentParameterAction, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *WhatsappMessageTemplateComponentParameter) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given WhatsappMessageTemplateComponentParameterAction and assigns it to the Action field.
func (o *WhatsappMessageTemplateComponentParameter) SetAction(v WhatsappMessageTemplateComponentParameterAction) {
	o.Action = &v
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *WhatsappMessageTemplateComponentParameter) GetOrderStatus() WhatsappMessageOrderStatus {
	if o == nil || o.OrderStatus == nil {
		var ret WhatsappMessageOrderStatus
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageTemplateComponentParameter) GetOrderStatusOk() (*WhatsappMessageOrderStatus, bool) {
	if o == nil || o.OrderStatus == nil {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *WhatsappMessageTemplateComponentParameter) HasOrderStatus() bool {
	if o != nil && o.OrderStatus != nil {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given WhatsappMessageOrderStatus and assigns it to the OrderStatus field.
func (o *WhatsappMessageTemplateComponentParameter) SetOrderStatus(v WhatsappMessageOrderStatus) {
	o.OrderStatus = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *WhatsappMessageTemplateComponentParameter) GetLocation() WhatsappMessageLocation {
	if o == nil || o.Location == nil {
		var ret WhatsappMessageLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsappMessageTemplateComponentParameter) GetLocationOk() (*WhatsappMessageLocation, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *WhatsappMessageTemplateComponentParameter) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given WhatsappMessageLocation and assigns it to the Location field.
func (o *WhatsappMessageTemplateComponentParameter) SetLocation(v WhatsappMessageLocation) {
	o.Location = &v
}

func (o WhatsappMessageTemplateComponentParameter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	if o.CouponCode != nil {
		toSerialize["coupon_code"] = o.CouponCode
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.Video != nil {
		toSerialize["video"] = o.Video
	}
	if o.Document != nil {
		toSerialize["document"] = o.Document
	}
	if o.LimitedTimeOffer != nil {
		toSerialize["limited_time_offer"] = o.LimitedTimeOffer
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.OrderStatus != nil {
		toSerialize["order_status"] = o.OrderStatus
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsappMessageTemplateComponentParameter struct {
	value *WhatsappMessageTemplateComponentParameter
	isSet bool
}

func (v NullableWhatsappMessageTemplateComponentParameter) Get() *WhatsappMessageTemplateComponentParameter {
	return v.value
}

func (v *NullableWhatsappMessageTemplateComponentParameter) Set(val *WhatsappMessageTemplateComponentParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsappMessageTemplateComponentParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsappMessageTemplateComponentParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsappMessageTemplateComponentParameter(val *WhatsappMessageTemplateComponentParameter) *NullableWhatsappMessageTemplateComponentParameter {
	return &NullableWhatsappMessageTemplateComponentParameter{value: val, isSet: true}
}

func (v NullableWhatsappMessageTemplateComponentParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsappMessageTemplateComponentParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


