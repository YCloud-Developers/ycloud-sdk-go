# WhatsappMessageTemplateComponentsInnerParametersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Describes the parameter type. | [optional] 
**Text** | Pointer to **string** | Required when &#x60;type&#x60; &#x3D; &#x60;text&#x60;. The message&#39;s text. For the header component, the character limit is 60 characters. For the body component, the character limit is 1024 characters. For url buttons, it indicates the developer-provided suffix that is appended to the predefined prefix URL in the template. | [optional] 
**Payload** | Pointer to **string** | Required for &#x60;quick_reply&#x60; buttons. Developer-defined payload that is returned when the button is clicked in addition to the display text on the button. | [optional] 
**Image** | Pointer to [**WhatsappMessageMedia**](WhatsappMessageMedia.md) |  | [optional] 
**Video** | Pointer to [**WhatsappMessageMedia**](WhatsappMessageMedia.md) |  | [optional] 
**Document** | Pointer to [**WhatsappMessageMedia**](WhatsappMessageMedia.md) |  | [optional] 

## Methods

### NewWhatsappMessageTemplateComponentsInnerParametersInner

`func NewWhatsappMessageTemplateComponentsInnerParametersInner() *WhatsappMessageTemplateComponentsInnerParametersInner`

NewWhatsappMessageTemplateComponentsInnerParametersInner instantiates a new WhatsappMessageTemplateComponentsInnerParametersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageTemplateComponentsInnerParametersInnerWithDefaults

`func NewWhatsappMessageTemplateComponentsInnerParametersInnerWithDefaults() *WhatsappMessageTemplateComponentsInnerParametersInner`

NewWhatsappMessageTemplateComponentsInnerParametersInnerWithDefaults instantiates a new WhatsappMessageTemplateComponentsInnerParametersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WhatsappMessageTemplateComponentsInnerParametersInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappMessageTemplateComponentsInnerParametersInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappMessageTemplateComponentsInnerParametersInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WhatsappMessageTemplateComponentsInnerParametersInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetText

`func (o *WhatsappMessageTemplateComponentsInnerParametersInner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WhatsappMessageTemplateComponentsInnerParametersInner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WhatsappMessageTemplateComponentsInnerParametersInner) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *WhatsappMessageTemplateComponentsInnerParametersInner) HasText() bool`

HasText returns a boolean if a field has been set.

### GetPayload

`func (o *WhatsappMessageTemplateComponentsInnerParametersInner) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *WhatsappMessageTemplateComponentsInnerParametersInner) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *WhatsappMessageTemplateComponentsInnerParametersInner) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *WhatsappMessageTemplateComponentsInnerParametersInner) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetImage

`func (o *WhatsappMessageTemplateComponentsInnerParametersInner) GetImage() WhatsappMessageMedia`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *WhatsappMessageTemplateComponentsInnerParametersInner) GetImageOk() (*WhatsappMessageMedia, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *WhatsappMessageTemplateComponentsInnerParametersInner) SetImage(v WhatsappMessageMedia)`

SetImage sets Image field to given value.

### HasImage

`func (o *WhatsappMessageTemplateComponentsInnerParametersInner) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetVideo

`func (o *WhatsappMessageTemplateComponentsInnerParametersInner) GetVideo() WhatsappMessageMedia`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *WhatsappMessageTemplateComponentsInnerParametersInner) GetVideoOk() (*WhatsappMessageMedia, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *WhatsappMessageTemplateComponentsInnerParametersInner) SetVideo(v WhatsappMessageMedia)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *WhatsappMessageTemplateComponentsInnerParametersInner) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetDocument

`func (o *WhatsappMessageTemplateComponentsInnerParametersInner) GetDocument() WhatsappMessageMedia`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *WhatsappMessageTemplateComponentsInnerParametersInner) GetDocumentOk() (*WhatsappMessageMedia, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *WhatsappMessageTemplateComponentsInnerParametersInner) SetDocument(v WhatsappMessageMedia)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *WhatsappMessageTemplateComponentsInnerParametersInner) HasDocument() bool`

HasDocument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


