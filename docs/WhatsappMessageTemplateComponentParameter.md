# WhatsappMessageTemplateComponentParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | **Required.** Describes the parameter type. - &#x60;text&#x60;: Used when the template component type is &#x60;BODY&#x60;, or the &#x60;HEADER&#x60; component format is &#x60;TEXT&#x60;. - &#x60;image&#x60;: Used when the template &#x60;HEADER&#x60; component is &#x60;IMAGE&#x60;. - &#x60;video&#x60;: Used when the template &#x60;HEADER&#x60; component is &#x60;VIDEO&#x60;. - &#x60;document&#x60;: Used when the template &#x60;HEADER&#x60; component is &#x60;DOCUMENT&#x60;. - &#x60;payload&#x60;: Used when the template component button type is &#x60;QUICK_REPLY&#x60;. - &#x60;coupon_code&#x60;: Used when the template component button type is &#x60;COPY_CODE&#x60;. - &#x60;limited_time_offer&#x60;: Used when the template component type is &#x60;LIMITED_TIME_OFFER&#x60;. | [optional] 
**Text** | Pointer to **string** | **Required when &#x60;type&#x60; &#x3D; &#x60;text&#x60;.** The message&#39;s text. For the header component, the character limit is 60 characters. For the body component, the character limit is 1024 characters. For url buttons, it indicates the developer-provided suffix that is appended to the predefined prefix URL in the template. | [optional] 
**Payload** | Pointer to **string** | Required for &#x60;quick_reply&#x60; buttons. Developer-defined payload that is returned when the button is clicked in addition to the display text on the button. | [optional] 
**CouponCode** | Pointer to **string** | **Required when &#x60;type&#x60; &#x3D; &#x60;coupon_code&#x60;.** The coupon code to be copied when the customer taps the button. | [optional] 
**Image** | Pointer to [**WhatsappMessageMedia**](WhatsappMessageMedia.md) |  | [optional] 
**Video** | Pointer to [**WhatsappMessageMedia**](WhatsappMessageMedia.md) |  | [optional] 
**Document** | Pointer to [**WhatsappMessageMedia**](WhatsappMessageMedia.md) |  | [optional] 
**LimitedTimeOffer** | Pointer to [**WhatsappMessageTemplateComponentParameterLimitedTimeOffer**](WhatsappMessageTemplateComponentParameterLimitedTimeOffer.md) |  | [optional] 

## Methods

### NewWhatsappMessageTemplateComponentParameter

`func NewWhatsappMessageTemplateComponentParameter() *WhatsappMessageTemplateComponentParameter`

NewWhatsappMessageTemplateComponentParameter instantiates a new WhatsappMessageTemplateComponentParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageTemplateComponentParameterWithDefaults

`func NewWhatsappMessageTemplateComponentParameterWithDefaults() *WhatsappMessageTemplateComponentParameter`

NewWhatsappMessageTemplateComponentParameterWithDefaults instantiates a new WhatsappMessageTemplateComponentParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WhatsappMessageTemplateComponentParameter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappMessageTemplateComponentParameter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappMessageTemplateComponentParameter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WhatsappMessageTemplateComponentParameter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetText

`func (o *WhatsappMessageTemplateComponentParameter) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WhatsappMessageTemplateComponentParameter) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WhatsappMessageTemplateComponentParameter) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *WhatsappMessageTemplateComponentParameter) HasText() bool`

HasText returns a boolean if a field has been set.

### GetPayload

`func (o *WhatsappMessageTemplateComponentParameter) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *WhatsappMessageTemplateComponentParameter) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *WhatsappMessageTemplateComponentParameter) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *WhatsappMessageTemplateComponentParameter) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetCouponCode

`func (o *WhatsappMessageTemplateComponentParameter) GetCouponCode() string`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *WhatsappMessageTemplateComponentParameter) GetCouponCodeOk() (*string, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *WhatsappMessageTemplateComponentParameter) SetCouponCode(v string)`

SetCouponCode sets CouponCode field to given value.

### HasCouponCode

`func (o *WhatsappMessageTemplateComponentParameter) HasCouponCode() bool`

HasCouponCode returns a boolean if a field has been set.

### GetImage

`func (o *WhatsappMessageTemplateComponentParameter) GetImage() WhatsappMessageMedia`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *WhatsappMessageTemplateComponentParameter) GetImageOk() (*WhatsappMessageMedia, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *WhatsappMessageTemplateComponentParameter) SetImage(v WhatsappMessageMedia)`

SetImage sets Image field to given value.

### HasImage

`func (o *WhatsappMessageTemplateComponentParameter) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetVideo

`func (o *WhatsappMessageTemplateComponentParameter) GetVideo() WhatsappMessageMedia`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *WhatsappMessageTemplateComponentParameter) GetVideoOk() (*WhatsappMessageMedia, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *WhatsappMessageTemplateComponentParameter) SetVideo(v WhatsappMessageMedia)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *WhatsappMessageTemplateComponentParameter) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetDocument

`func (o *WhatsappMessageTemplateComponentParameter) GetDocument() WhatsappMessageMedia`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *WhatsappMessageTemplateComponentParameter) GetDocumentOk() (*WhatsappMessageMedia, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *WhatsappMessageTemplateComponentParameter) SetDocument(v WhatsappMessageMedia)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *WhatsappMessageTemplateComponentParameter) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetLimitedTimeOffer

`func (o *WhatsappMessageTemplateComponentParameter) GetLimitedTimeOffer() WhatsappMessageTemplateComponentParameterLimitedTimeOffer`

GetLimitedTimeOffer returns the LimitedTimeOffer field if non-nil, zero value otherwise.

### GetLimitedTimeOfferOk

`func (o *WhatsappMessageTemplateComponentParameter) GetLimitedTimeOfferOk() (*WhatsappMessageTemplateComponentParameterLimitedTimeOffer, bool)`

GetLimitedTimeOfferOk returns a tuple with the LimitedTimeOffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitedTimeOffer

`func (o *WhatsappMessageTemplateComponentParameter) SetLimitedTimeOffer(v WhatsappMessageTemplateComponentParameterLimitedTimeOffer)`

SetLimitedTimeOffer sets LimitedTimeOffer field to given value.

### HasLimitedTimeOffer

`func (o *WhatsappMessageTemplateComponentParameter) HasLimitedTimeOffer() bool`

HasLimitedTimeOffer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


