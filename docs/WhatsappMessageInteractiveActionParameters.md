# WhatsappMessageInteractiveActionParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayText** | Pointer to **string** | Text of the CTA URL button. Maximum length: 20 bytes. | [optional] 
**Url** | Pointer to **string** | URL of the CTA URL button. | [optional] 
**ThumbnailProductRetailerId** | Pointer to **string** | Item SKU number. Labeled as **Content ID** in the [Commerce Manager](https://business.facebook.com/commerce/). The thumbnail of this item will be used as the message&#39;s header image. | [optional] 

## Methods

### NewWhatsappMessageInteractiveActionParameters

`func NewWhatsappMessageInteractiveActionParameters() *WhatsappMessageInteractiveActionParameters`

NewWhatsappMessageInteractiveActionParameters instantiates a new WhatsappMessageInteractiveActionParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageInteractiveActionParametersWithDefaults

`func NewWhatsappMessageInteractiveActionParametersWithDefaults() *WhatsappMessageInteractiveActionParameters`

NewWhatsappMessageInteractiveActionParametersWithDefaults instantiates a new WhatsappMessageInteractiveActionParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayText

`func (o *WhatsappMessageInteractiveActionParameters) GetDisplayText() string`

GetDisplayText returns the DisplayText field if non-nil, zero value otherwise.

### GetDisplayTextOk

`func (o *WhatsappMessageInteractiveActionParameters) GetDisplayTextOk() (*string, bool)`

GetDisplayTextOk returns a tuple with the DisplayText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayText

`func (o *WhatsappMessageInteractiveActionParameters) SetDisplayText(v string)`

SetDisplayText sets DisplayText field to given value.

### HasDisplayText

`func (o *WhatsappMessageInteractiveActionParameters) HasDisplayText() bool`

HasDisplayText returns a boolean if a field has been set.

### GetUrl

`func (o *WhatsappMessageInteractiveActionParameters) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WhatsappMessageInteractiveActionParameters) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WhatsappMessageInteractiveActionParameters) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WhatsappMessageInteractiveActionParameters) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetThumbnailProductRetailerId

`func (o *WhatsappMessageInteractiveActionParameters) GetThumbnailProductRetailerId() string`

GetThumbnailProductRetailerId returns the ThumbnailProductRetailerId field if non-nil, zero value otherwise.

### GetThumbnailProductRetailerIdOk

`func (o *WhatsappMessageInteractiveActionParameters) GetThumbnailProductRetailerIdOk() (*string, bool)`

GetThumbnailProductRetailerIdOk returns a tuple with the ThumbnailProductRetailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailProductRetailerId

`func (o *WhatsappMessageInteractiveActionParameters) SetThumbnailProductRetailerId(v string)`

SetThumbnailProductRetailerId sets ThumbnailProductRetailerId field to given value.

### HasThumbnailProductRetailerId

`func (o *WhatsappMessageInteractiveActionParameters) HasThumbnailProductRetailerId() bool`

HasThumbnailProductRetailerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
