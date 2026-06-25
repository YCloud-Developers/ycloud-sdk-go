# WhatsappMessageInteractiveActionCardHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | **Required.** The header type you would like to use. - &#x60;video&#x60;: Used for Reply Buttons. - &#x60;image&#x60;: Used for Reply Buttons. | [optional] 
**Image** | Pointer to [**WhatsappMessageMedia**](WhatsappMessageMedia.md) |  | [optional] 
**Video** | Pointer to [**WhatsappMessageMedia**](WhatsappMessageMedia.md) |  | [optional] 

## Methods

### NewWhatsappMessageInteractiveActionCardHeader

`func NewWhatsappMessageInteractiveActionCardHeader() *WhatsappMessageInteractiveActionCardHeader`

NewWhatsappMessageInteractiveActionCardHeader instantiates a new WhatsappMessageInteractiveActionCardHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageInteractiveActionCardHeaderWithDefaults

`func NewWhatsappMessageInteractiveActionCardHeaderWithDefaults() *WhatsappMessageInteractiveActionCardHeader`

NewWhatsappMessageInteractiveActionCardHeaderWithDefaults instantiates a new WhatsappMessageInteractiveActionCardHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WhatsappMessageInteractiveActionCardHeader) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappMessageInteractiveActionCardHeader) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappMessageInteractiveActionCardHeader) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WhatsappMessageInteractiveActionCardHeader) HasType() bool`

HasType returns a boolean if a field has been set.

### GetImage

`func (o *WhatsappMessageInteractiveActionCardHeader) GetImage() WhatsappMessageMedia`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *WhatsappMessageInteractiveActionCardHeader) GetImageOk() (*WhatsappMessageMedia, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *WhatsappMessageInteractiveActionCardHeader) SetImage(v WhatsappMessageMedia)`

SetImage sets Image field to given value.

### HasImage

`func (o *WhatsappMessageInteractiveActionCardHeader) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetVideo

`func (o *WhatsappMessageInteractiveActionCardHeader) GetVideo() WhatsappMessageMedia`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *WhatsappMessageInteractiveActionCardHeader) GetVideoOk() (*WhatsappMessageMedia, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *WhatsappMessageInteractiveActionCardHeader) SetVideo(v WhatsappMessageMedia)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *WhatsappMessageInteractiveActionCardHeader) HasVideo() bool`

HasVideo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


