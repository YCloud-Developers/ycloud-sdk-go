# WhatsappInboundMessageMedia

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the media. Can be used to delete the media if stored locally on the client. | [optional] 
**Link** | Pointer to **string** | The url to download the media file. Note that This link can be directly accessed in a few minutes for the convenience of the consumer, but you should always include an &#x60;X-API-Key&#x60; header to download this file within a month. | [optional] 
**Caption** | Pointer to **string** | The provided caption for the media. Only present if specified. | [optional] 
**Filename** | Pointer to **string** | Filename on the sender&#39;s device. This will only be present in &#x60;document&#x60; media messages. | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Metadata pertaining to &#x60;sticker&#x60; media. | [optional] 
**MimeType** | Pointer to **string** | Mime type of the media. | [optional] 
**Sha256** | Pointer to **string** | Checksum. | [optional] 

## Methods

### NewWhatsappInboundMessageMedia

`func NewWhatsappInboundMessageMedia() *WhatsappInboundMessageMedia`

NewWhatsappInboundMessageMedia instantiates a new WhatsappInboundMessageMedia object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappInboundMessageMediaWithDefaults

`func NewWhatsappInboundMessageMediaWithDefaults() *WhatsappInboundMessageMedia`

NewWhatsappInboundMessageMediaWithDefaults instantiates a new WhatsappInboundMessageMedia object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WhatsappInboundMessageMedia) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WhatsappInboundMessageMedia) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WhatsappInboundMessageMedia) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WhatsappInboundMessageMedia) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLink

`func (o *WhatsappInboundMessageMedia) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *WhatsappInboundMessageMedia) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *WhatsappInboundMessageMedia) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *WhatsappInboundMessageMedia) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetCaption

`func (o *WhatsappInboundMessageMedia) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *WhatsappInboundMessageMedia) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *WhatsappInboundMessageMedia) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *WhatsappInboundMessageMedia) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetFilename

`func (o *WhatsappInboundMessageMedia) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *WhatsappInboundMessageMedia) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *WhatsappInboundMessageMedia) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *WhatsappInboundMessageMedia) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetMetadata

`func (o *WhatsappInboundMessageMedia) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WhatsappInboundMessageMedia) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WhatsappInboundMessageMedia) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WhatsappInboundMessageMedia) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMimeType

`func (o *WhatsappInboundMessageMedia) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *WhatsappInboundMessageMedia) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *WhatsappInboundMessageMedia) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *WhatsappInboundMessageMedia) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetSha256

`func (o *WhatsappInboundMessageMedia) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *WhatsappInboundMessageMedia) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *WhatsappInboundMessageMedia) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *WhatsappInboundMessageMedia) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
