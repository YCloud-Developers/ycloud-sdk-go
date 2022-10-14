# WhatsappMessageMedia

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | **string** | The protocol and URL of the media to be sent. Use only with HTTP/HTTPS URLs. | 
**Caption** | Pointer to **string** | Describes the specified &#x60;video&#x60; or &#x60;image&#x60; media. | [optional] 
**Filename** | Pointer to **string** | Describes the filename for the specific document. Use only with &#x60;document&#x60; media. | [optional] 

## Methods

### NewWhatsappMessageMedia

`func NewWhatsappMessageMedia(link string, ) *WhatsappMessageMedia`

NewWhatsappMessageMedia instantiates a new WhatsappMessageMedia object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageMediaWithDefaults

`func NewWhatsappMessageMediaWithDefaults() *WhatsappMessageMedia`

NewWhatsappMessageMediaWithDefaults instantiates a new WhatsappMessageMedia object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLink

`func (o *WhatsappMessageMedia) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *WhatsappMessageMedia) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *WhatsappMessageMedia) SetLink(v string)`

SetLink sets Link field to given value.


### GetCaption

`func (o *WhatsappMessageMedia) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *WhatsappMessageMedia) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *WhatsappMessageMedia) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *WhatsappMessageMedia) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetFilename

`func (o *WhatsappMessageMedia) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *WhatsappMessageMedia) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *WhatsappMessageMedia) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *WhatsappMessageMedia) HasFilename() bool`

HasFilename returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


