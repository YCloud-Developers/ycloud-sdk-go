# WhatsappMessageInteractiveAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buttons** | Pointer to [**[]WhatsappMessageInteractiveActionButtonsInner**](WhatsappMessageInteractiveActionButtonsInner.md) | Required for Reply Buttons. | [optional] 
**Button** | Pointer to **string** | Required for List Messages. Button content. It cannot be an empty string and must be unique within the message. Emojis are supported, markdown is not. | [optional] 
**Sections** | Pointer to [**[]WhatsappMessageInteractiveActionSectionsInner**](WhatsappMessageInteractiveActionSectionsInner.md) | Required for List Messages. | [optional] 

## Methods

### NewWhatsappMessageInteractiveAction

`func NewWhatsappMessageInteractiveAction() *WhatsappMessageInteractiveAction`

NewWhatsappMessageInteractiveAction instantiates a new WhatsappMessageInteractiveAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageInteractiveActionWithDefaults

`func NewWhatsappMessageInteractiveActionWithDefaults() *WhatsappMessageInteractiveAction`

NewWhatsappMessageInteractiveActionWithDefaults instantiates a new WhatsappMessageInteractiveAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetButtons

`func (o *WhatsappMessageInteractiveAction) GetButtons() []WhatsappMessageInteractiveActionButtonsInner`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *WhatsappMessageInteractiveAction) GetButtonsOk() (*[]WhatsappMessageInteractiveActionButtonsInner, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *WhatsappMessageInteractiveAction) SetButtons(v []WhatsappMessageInteractiveActionButtonsInner)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *WhatsappMessageInteractiveAction) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### GetButton

`func (o *WhatsappMessageInteractiveAction) GetButton() string`

GetButton returns the Button field if non-nil, zero value otherwise.

### GetButtonOk

`func (o *WhatsappMessageInteractiveAction) GetButtonOk() (*string, bool)`

GetButtonOk returns a tuple with the Button field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButton

`func (o *WhatsappMessageInteractiveAction) SetButton(v string)`

SetButton sets Button field to given value.

### HasButton

`func (o *WhatsappMessageInteractiveAction) HasButton() bool`

HasButton returns a boolean if a field has been set.

### GetSections

`func (o *WhatsappMessageInteractiveAction) GetSections() []WhatsappMessageInteractiveActionSectionsInner`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *WhatsappMessageInteractiveAction) GetSectionsOk() (*[]WhatsappMessageInteractiveActionSectionsInner, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *WhatsappMessageInteractiveAction) SetSections(v []WhatsappMessageInteractiveActionSectionsInner)`

SetSections sets Sections field to given value.

### HasSections

`func (o *WhatsappMessageInteractiveAction) HasSections() bool`

HasSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


