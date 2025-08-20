# WhatsappMessageMedia

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Required when using media that has been uploaded to WhatsApp servers.  Provide the media object ID obtained from WhatsApp media upload API (https://docs.ycloud.com/update/reference/whatsapp_media-upload#/). | 
**Link** | **string** | Required when sending media directly from your server.  The protocol and URL of the media to be sent. Use only with HTTP/HTTPS URLs. Note: WhatsApp Cloud API caches media resources for 10 minutes. To ensure latest content, add random query strings to the URL. | 
**Caption** | Pointer to **string** | Describes the specified &#x60;image&#x60;, &#x60;video&#x60;, or &#x60;document&#x60; media. Not applicable in the &#x60;header&#x60; of &#x60;template&#x60; or &#x60;interactive&#x60; messages. | [optional] 
**Filename** | Pointer to **string** | Describes the filename for the specific document. Use only with &#x60;document&#x60; media. | [optional] 

## Methods

### NewWhatsappMessageMedia

`func NewWhatsappMessageMedia(id string, link string, ) *WhatsappMessageMedia`

NewWhatsappMessageMedia instantiates a new WhatsappMessageMedia object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageMediaWithDefaults

`func NewWhatsappMessageMediaWithDefaults() *WhatsappMessageMedia`

NewWhatsappMessageMediaWithDefaults instantiates a new WhatsappMessageMedia object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WhatsappMessageMedia) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WhatsappMessageMedia) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WhatsappMessageMedia) SetId(v string)`

SetId sets Id field to given value.


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


