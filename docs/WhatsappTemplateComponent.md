# WhatsappTemplateComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | **Required.** Template component type. - &#x60;BODY&#x60;: Body components are text-only components and are required by all templates. Templates are limited to one body component. - &#x60;HEADER&#x60;: Headers are optional components that appear at the top of template messages. Headers support text, media (images, videos, documents). Templates are limited to one header component. - &#x60;FOOTER&#x60;: Footers are optional text-only components that appear immediately after the body component. Templates are limited to one footer component. - &#x60;BUTTONS&#x60;: Buttons are optional interactive components that perform specific actions when tapped. - &#x60;LIMITED_TIME_OFFER&#x60;: Use for limited-time offer templates. The delivered message can display an offer expiration details section with a heading, an optional expiration timer, and the offer code itself. - &#x60;CAROUSEL&#x60;: Carousel templates allow you to send a single text message (1), accompanied by a set of up to 10 carousel cards (2) in a horizontally scrollable view. | [optional] 
**Format** | Pointer to **string** | **Required for type &#x60;HEADER&#x60;.** | [optional] 
**Text** | Pointer to **string** | For body text (type &#x3D; &#x60;BODY&#x60;), maximum 1024 characters. For header text (type &#x3D; &#x60;HEADER&#x60;, format &#x3D; &#x60;TEXT&#x60;), maximum 60 characters. For footer text (type &#x3D; &#x60;FOOTER&#x60;), maximum 60 characters. For card body text (&#x60;CAROUSEL&#x60; card component type &#x3D; &#x60;BODY&#x60;), maximum 160 characters. | [optional] 
**Buttons** | Pointer to [**[]WhatsappTemplateComponentButton**](WhatsappTemplateComponentButton.md) | **Required for type &#x60;BUTTONS&#x60;.** Buttons are optional interactive components that perform specific actions when tapped. Templates can have a mixture of up to 10 button components total, although there are limits to individual buttons of the same type as well as combination limits. If a template has more than three buttons, two buttons will appear in the delivered message and the remaining buttons will be replaced with a **See all options** button. Tapping the **See all options** button reveals the remaining buttons. | [optional] 
**AddSecurityRecommendation** | Pointer to **bool** | **Optional. Only applicable in the &#x60;BODY&#x60; component of an AUTHENTICATION template.** Set to &#x60;true&#x60; if you want the template to include the string, *For your security, do not share this code.* Set to &#x60;false&#x60; to exclude the string. | [optional] 
**CodeExpirationMinutes** | Pointer to **int32** | **Optional. Only applicable in the &#x60;FOOTER&#x60; component of an AUTHENTICATION template.** Indicates number of minutes the password or code is valid. If omitted, the code expiration warning will not be displayed in the delivered message. Minimum 1, maximum 90. | [optional] 
**LimitedTimeOffer** | Pointer to [**WhatsappTemplateComponentLimitedTimeOffer**](WhatsappTemplateComponentLimitedTimeOffer.md) |  | [optional] 
**Example** | Pointer to [**WhatsappTemplateComponentExample**](WhatsappTemplateComponentExample.md) |  | [optional] 
**Cards** | Pointer to [**[]WhatsappTemplateComponentCard**](WhatsappTemplateComponentCard.md) | **Required for type &#x60;CAROUSEL&#x60;.** Carousel templates support up to 10 carousel cards. | [optional] 

## Methods

### NewWhatsappTemplateComponent

`func NewWhatsappTemplateComponent() *WhatsappTemplateComponent`

NewWhatsappTemplateComponent instantiates a new WhatsappTemplateComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappTemplateComponentWithDefaults

`func NewWhatsappTemplateComponentWithDefaults() *WhatsappTemplateComponent`

NewWhatsappTemplateComponentWithDefaults instantiates a new WhatsappTemplateComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WhatsappTemplateComponent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappTemplateComponent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappTemplateComponent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WhatsappTemplateComponent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFormat

`func (o *WhatsappTemplateComponent) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *WhatsappTemplateComponent) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *WhatsappTemplateComponent) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *WhatsappTemplateComponent) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetText

`func (o *WhatsappTemplateComponent) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WhatsappTemplateComponent) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WhatsappTemplateComponent) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *WhatsappTemplateComponent) HasText() bool`

HasText returns a boolean if a field has been set.

### GetButtons

`func (o *WhatsappTemplateComponent) GetButtons() []WhatsappTemplateComponentButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *WhatsappTemplateComponent) GetButtonsOk() (*[]WhatsappTemplateComponentButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *WhatsappTemplateComponent) SetButtons(v []WhatsappTemplateComponentButton)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *WhatsappTemplateComponent) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### GetAddSecurityRecommendation

`func (o *WhatsappTemplateComponent) GetAddSecurityRecommendation() bool`

GetAddSecurityRecommendation returns the AddSecurityRecommendation field if non-nil, zero value otherwise.

### GetAddSecurityRecommendationOk

`func (o *WhatsappTemplateComponent) GetAddSecurityRecommendationOk() (*bool, bool)`

GetAddSecurityRecommendationOk returns a tuple with the AddSecurityRecommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddSecurityRecommendation

`func (o *WhatsappTemplateComponent) SetAddSecurityRecommendation(v bool)`

SetAddSecurityRecommendation sets AddSecurityRecommendation field to given value.

### HasAddSecurityRecommendation

`func (o *WhatsappTemplateComponent) HasAddSecurityRecommendation() bool`

HasAddSecurityRecommendation returns a boolean if a field has been set.

### GetCodeExpirationMinutes

`func (o *WhatsappTemplateComponent) GetCodeExpirationMinutes() int32`

GetCodeExpirationMinutes returns the CodeExpirationMinutes field if non-nil, zero value otherwise.

### GetCodeExpirationMinutesOk

`func (o *WhatsappTemplateComponent) GetCodeExpirationMinutesOk() (*int32, bool)`

GetCodeExpirationMinutesOk returns a tuple with the CodeExpirationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeExpirationMinutes

`func (o *WhatsappTemplateComponent) SetCodeExpirationMinutes(v int32)`

SetCodeExpirationMinutes sets CodeExpirationMinutes field to given value.

### HasCodeExpirationMinutes

`func (o *WhatsappTemplateComponent) HasCodeExpirationMinutes() bool`

HasCodeExpirationMinutes returns a boolean if a field has been set.

### GetLimitedTimeOffer

`func (o *WhatsappTemplateComponent) GetLimitedTimeOffer() WhatsappTemplateComponentLimitedTimeOffer`

GetLimitedTimeOffer returns the LimitedTimeOffer field if non-nil, zero value otherwise.

### GetLimitedTimeOfferOk

`func (o *WhatsappTemplateComponent) GetLimitedTimeOfferOk() (*WhatsappTemplateComponentLimitedTimeOffer, bool)`

GetLimitedTimeOfferOk returns a tuple with the LimitedTimeOffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitedTimeOffer

`func (o *WhatsappTemplateComponent) SetLimitedTimeOffer(v WhatsappTemplateComponentLimitedTimeOffer)`

SetLimitedTimeOffer sets LimitedTimeOffer field to given value.

### HasLimitedTimeOffer

`func (o *WhatsappTemplateComponent) HasLimitedTimeOffer() bool`

HasLimitedTimeOffer returns a boolean if a field has been set.

### GetExample

`func (o *WhatsappTemplateComponent) GetExample() WhatsappTemplateComponentExample`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *WhatsappTemplateComponent) GetExampleOk() (*WhatsappTemplateComponentExample, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *WhatsappTemplateComponent) SetExample(v WhatsappTemplateComponentExample)`

SetExample sets Example field to given value.

### HasExample

`func (o *WhatsappTemplateComponent) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetCards

`func (o *WhatsappTemplateComponent) GetCards() []WhatsappTemplateComponentCard`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *WhatsappTemplateComponent) GetCardsOk() (*[]WhatsappTemplateComponentCard, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *WhatsappTemplateComponent) SetCards(v []WhatsappTemplateComponentCard)`

SetCards sets Cards field to given value.

### HasCards

`func (o *WhatsappTemplateComponent) HasCards() bool`

HasCards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


