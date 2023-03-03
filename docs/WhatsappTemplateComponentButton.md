# WhatsappTemplateComponentButton

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**WhatsappTemplateComponentButtonType**](WhatsappTemplateComponentButtonType.md) |  | 
**Text** | **string** | **Required.** Button text. | 
**Url** | Pointer to **string** | **Required for button type &#x60;URL&#x60;.**  There can be at most 1 variable at the end of the URL. | [optional] 
**PhoneNumber** | Pointer to **string** | **Required for button type &#x60;PHONE_NUMBER&#x60;.** | [optional] 
**Example** | Pointer to **[]string** | Sample full URL for a &#x60;URL&#x60; button with a variable. | [optional] 

## Methods

### NewWhatsappTemplateComponentButton

`func NewWhatsappTemplateComponentButton(type_ WhatsappTemplateComponentButtonType, text string, ) *WhatsappTemplateComponentButton`

NewWhatsappTemplateComponentButton instantiates a new WhatsappTemplateComponentButton object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappTemplateComponentButtonWithDefaults

`func NewWhatsappTemplateComponentButtonWithDefaults() *WhatsappTemplateComponentButton`

NewWhatsappTemplateComponentButtonWithDefaults instantiates a new WhatsappTemplateComponentButton object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WhatsappTemplateComponentButton) GetType() WhatsappTemplateComponentButtonType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappTemplateComponentButton) GetTypeOk() (*WhatsappTemplateComponentButtonType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappTemplateComponentButton) SetType(v WhatsappTemplateComponentButtonType)`

SetType sets Type field to given value.


### GetText

`func (o *WhatsappTemplateComponentButton) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WhatsappTemplateComponentButton) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WhatsappTemplateComponentButton) SetText(v string)`

SetText sets Text field to given value.


### GetUrl

`func (o *WhatsappTemplateComponentButton) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WhatsappTemplateComponentButton) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WhatsappTemplateComponentButton) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WhatsappTemplateComponentButton) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *WhatsappTemplateComponentButton) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *WhatsappTemplateComponentButton) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *WhatsappTemplateComponentButton) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *WhatsappTemplateComponentButton) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetExample

`func (o *WhatsappTemplateComponentButton) GetExample() []string`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *WhatsappTemplateComponentButton) GetExampleOk() (*[]string, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *WhatsappTemplateComponentButton) SetExample(v []string)`

SetExample sets Example field to given value.

### HasExample

`func (o *WhatsappTemplateComponentButton) HasExample() bool`

HasExample returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


