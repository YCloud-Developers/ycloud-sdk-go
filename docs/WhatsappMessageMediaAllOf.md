# WhatsappMessageMediaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Caption** | Pointer to **string** | Describes the specified &#x60;image&#x60;, &#x60;video&#x60;, or &#x60;document&#x60; media. Not applicable in the &#x60;header&#x60; of &#x60;template&#x60; or &#x60;interactive&#x60; messages. | [optional] 
**Filename** | Pointer to **string** | Describes the filename for the specific document. Use only with &#x60;document&#x60; media. | [optional] 

## Methods

### NewWhatsappMessageMediaAllOf

`func NewWhatsappMessageMediaAllOf() *WhatsappMessageMediaAllOf`

NewWhatsappMessageMediaAllOf instantiates a new WhatsappMessageMediaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageMediaAllOfWithDefaults

`func NewWhatsappMessageMediaAllOfWithDefaults() *WhatsappMessageMediaAllOf`

NewWhatsappMessageMediaAllOfWithDefaults instantiates a new WhatsappMessageMediaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaption

`func (o *WhatsappMessageMediaAllOf) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *WhatsappMessageMediaAllOf) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *WhatsappMessageMediaAllOf) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *WhatsappMessageMediaAllOf) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetFilename

`func (o *WhatsappMessageMediaAllOf) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *WhatsappMessageMediaAllOf) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *WhatsappMessageMediaAllOf) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *WhatsappMessageMediaAllOf) HasFilename() bool`

HasFilename returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


