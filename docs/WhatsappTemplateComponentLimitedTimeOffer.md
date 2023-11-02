# WhatsappTemplateComponentLimitedTimeOffer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** | **Required.** Offer details text. Maximum 16 characters. | [optional] 
**HasExpiration** | Pointer to **bool** | **Optional.** Set to &#x60;true&#x60; to have the [offer expiration details](https://developers.facebook.com/docs/whatsapp/business-management-api/message-templates/limited-time-offer-templates#offer-expiration-details) appear in the delivered message. If set to &#x60;true&#x60;, the copy code button component must be included in the &#x60;buttons&#x60; array, and must appear first in the array. If set to &#x60;false&#x60;, offer expiration details will not appear in the delivered message and the copy code button component is optional. If including the copy code button, it must appear first in the &#x60;buttons&#x60; array. | [optional] 

## Methods

### NewWhatsappTemplateComponentLimitedTimeOffer

`func NewWhatsappTemplateComponentLimitedTimeOffer() *WhatsappTemplateComponentLimitedTimeOffer`

NewWhatsappTemplateComponentLimitedTimeOffer instantiates a new WhatsappTemplateComponentLimitedTimeOffer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappTemplateComponentLimitedTimeOfferWithDefaults

`func NewWhatsappTemplateComponentLimitedTimeOfferWithDefaults() *WhatsappTemplateComponentLimitedTimeOffer`

NewWhatsappTemplateComponentLimitedTimeOfferWithDefaults instantiates a new WhatsappTemplateComponentLimitedTimeOffer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *WhatsappTemplateComponentLimitedTimeOffer) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WhatsappTemplateComponentLimitedTimeOffer) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WhatsappTemplateComponentLimitedTimeOffer) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *WhatsappTemplateComponentLimitedTimeOffer) HasText() bool`

HasText returns a boolean if a field has been set.

### GetHasExpiration

`func (o *WhatsappTemplateComponentLimitedTimeOffer) GetHasExpiration() bool`

GetHasExpiration returns the HasExpiration field if non-nil, zero value otherwise.

### GetHasExpirationOk

`func (o *WhatsappTemplateComponentLimitedTimeOffer) GetHasExpirationOk() (*bool, bool)`

GetHasExpirationOk returns a tuple with the HasExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasExpiration

`func (o *WhatsappTemplateComponentLimitedTimeOffer) SetHasExpiration(v bool)`

SetHasExpiration sets HasExpiration field to given value.

### HasHasExpiration

`func (o *WhatsappTemplateComponentLimitedTimeOffer) HasHasExpiration() bool`

HasHasExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


