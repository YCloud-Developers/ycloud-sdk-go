# WhatsappMessageInteractiveHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The header type you would like to use. - &#x60;text&#x60;: Used for List Messages and Reply Buttons. - &#x60;video&#x60;: Used for Reply Buttons. - &#x60;image&#x60;: Used for Reply Buttons. - &#x60;document&#x60;: Used for Reply Buttons. | 
**Text** | Pointer to **string** | Text for the header. Formatting allows emojis, but not markdown. | [optional] 
**Image** | Pointer to [**WhatsappMessageMedia**](WhatsappMessageMedia.md) |  | [optional] 
**Video** | Pointer to [**WhatsappMessageMedia**](WhatsappMessageMedia.md) |  | [optional] 
**Document** | Pointer to [**WhatsappMessageMedia**](WhatsappMessageMedia.md) |  | [optional] 

## Methods

### NewWhatsappMessageInteractiveHeader

`func NewWhatsappMessageInteractiveHeader(type_ string, ) *WhatsappMessageInteractiveHeader`

NewWhatsappMessageInteractiveHeader instantiates a new WhatsappMessageInteractiveHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageInteractiveHeaderWithDefaults

`func NewWhatsappMessageInteractiveHeaderWithDefaults() *WhatsappMessageInteractiveHeader`

NewWhatsappMessageInteractiveHeaderWithDefaults instantiates a new WhatsappMessageInteractiveHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WhatsappMessageInteractiveHeader) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappMessageInteractiveHeader) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappMessageInteractiveHeader) SetType(v string)`

SetType sets Type field to given value.


### GetText

`func (o *WhatsappMessageInteractiveHeader) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WhatsappMessageInteractiveHeader) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WhatsappMessageInteractiveHeader) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *WhatsappMessageInteractiveHeader) HasText() bool`

HasText returns a boolean if a field has been set.

### GetImage

`func (o *WhatsappMessageInteractiveHeader) GetImage() WhatsappMessageMedia`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *WhatsappMessageInteractiveHeader) GetImageOk() (*WhatsappMessageMedia, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *WhatsappMessageInteractiveHeader) SetImage(v WhatsappMessageMedia)`

SetImage sets Image field to given value.

### HasImage

`func (o *WhatsappMessageInteractiveHeader) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetVideo

`func (o *WhatsappMessageInteractiveHeader) GetVideo() WhatsappMessageMedia`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *WhatsappMessageInteractiveHeader) GetVideoOk() (*WhatsappMessageMedia, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *WhatsappMessageInteractiveHeader) SetVideo(v WhatsappMessageMedia)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *WhatsappMessageInteractiveHeader) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetDocument

`func (o *WhatsappMessageInteractiveHeader) GetDocument() WhatsappMessageMedia`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *WhatsappMessageInteractiveHeader) GetDocumentOk() (*WhatsappMessageMedia, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *WhatsappMessageInteractiveHeader) SetDocument(v WhatsappMessageMedia)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *WhatsappMessageInteractiveHeader) HasDocument() bool`

HasDocument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


