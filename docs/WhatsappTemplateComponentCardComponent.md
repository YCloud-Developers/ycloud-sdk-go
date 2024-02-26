# WhatsappTemplateComponentCardComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | **Required.** Card component type. - &#x60;BODY&#x60;: Body components are text-only components. Cards must have body text. - &#x60;HEADER&#x60;: Cards must have a media header (image or video). - &#x60;BUTTONS&#x60;: Buttons are interactive components that perform specific actions when tapped. Cards must have at least one button, up to 2 buttons. | [optional] 
**Format** | Pointer to **string** | **Required for type &#x60;HEADER&#x60;.** Cards must have a media header (image or video). | [optional] 
**Text** | Pointer to **string** | **Required for type &#x60;BODY&#x60;.** Card body text supports variables. Maximum 160 characters. | [optional] 
**Buttons** | Pointer to [**[]WhatsappTemplateComponentButton**](WhatsappTemplateComponentButton.md) | **Required for type &#x60;BUTTONS&#x60;.** Cards must have at least one button. Supports 2 buttons. Buttons can be the same or a mix of quick reply buttons, phone number buttons, or URL buttons. | [optional] 
**Example** | Pointer to [**WhatsappTemplateComponentExample**](WhatsappTemplateComponentExample.md) |  | [optional] 

## Methods

### NewWhatsappTemplateComponentCardComponent

`func NewWhatsappTemplateComponentCardComponent() *WhatsappTemplateComponentCardComponent`

NewWhatsappTemplateComponentCardComponent instantiates a new WhatsappTemplateComponentCardComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappTemplateComponentCardComponentWithDefaults

`func NewWhatsappTemplateComponentCardComponentWithDefaults() *WhatsappTemplateComponentCardComponent`

NewWhatsappTemplateComponentCardComponentWithDefaults instantiates a new WhatsappTemplateComponentCardComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WhatsappTemplateComponentCardComponent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappTemplateComponentCardComponent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappTemplateComponentCardComponent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WhatsappTemplateComponentCardComponent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFormat

`func (o *WhatsappTemplateComponentCardComponent) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *WhatsappTemplateComponentCardComponent) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *WhatsappTemplateComponentCardComponent) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *WhatsappTemplateComponentCardComponent) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetText

`func (o *WhatsappTemplateComponentCardComponent) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WhatsappTemplateComponentCardComponent) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WhatsappTemplateComponentCardComponent) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *WhatsappTemplateComponentCardComponent) HasText() bool`

HasText returns a boolean if a field has been set.

### GetButtons

`func (o *WhatsappTemplateComponentCardComponent) GetButtons() []WhatsappTemplateComponentButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *WhatsappTemplateComponentCardComponent) GetButtonsOk() (*[]WhatsappTemplateComponentButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *WhatsappTemplateComponentCardComponent) SetButtons(v []WhatsappTemplateComponentButton)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *WhatsappTemplateComponentCardComponent) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### GetExample

`func (o *WhatsappTemplateComponentCardComponent) GetExample() WhatsappTemplateComponentExample`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *WhatsappTemplateComponentCardComponent) GetExampleOk() (*WhatsappTemplateComponentExample, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *WhatsappTemplateComponentCardComponent) SetExample(v WhatsappTemplateComponentExample)`

SetExample sets Example field to given value.

### HasExample

`func (o *WhatsappTemplateComponentCardComponent) HasExample() bool`

HasExample returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


