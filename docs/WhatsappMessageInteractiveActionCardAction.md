# WhatsappMessageInteractiveActionCardAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Required when card action is url button. Must be \&quot;cta_url\&quot;. | [optional]
**Parameters** | Pointer to [**WhatsappMessageInteractiveActionCardActionParameters**](WhatsappMessageInteractiveActionCardActionParameters.md) |  | [optional]
**Buttons** | Pointer to [**[]WhatsappMessageInteractiveActionCardActionButton**](WhatsappMessageInteractiveActionCardActionButton.md) | Required when card action is quick reply button. | [optional]

## Methods

### NewWhatsappMessageInteractiveActionCardAction

`func NewWhatsappMessageInteractiveActionCardAction() *WhatsappMessageInteractiveActionCardAction`

NewWhatsappMessageInteractiveActionCardAction instantiates a new WhatsappMessageInteractiveActionCardAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageInteractiveActionCardActionWithDefaults

`func NewWhatsappMessageInteractiveActionCardActionWithDefaults() *WhatsappMessageInteractiveActionCardAction`

NewWhatsappMessageInteractiveActionCardActionWithDefaults instantiates a new WhatsappMessageInteractiveActionCardAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WhatsappMessageInteractiveActionCardAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WhatsappMessageInteractiveActionCardAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WhatsappMessageInteractiveActionCardAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WhatsappMessageInteractiveActionCardAction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *WhatsappMessageInteractiveActionCardAction) GetParameters() WhatsappMessageInteractiveActionCardActionParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *WhatsappMessageInteractiveActionCardAction) GetParametersOk() (*WhatsappMessageInteractiveActionCardActionParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *WhatsappMessageInteractiveActionCardAction) SetParameters(v WhatsappMessageInteractiveActionCardActionParameters)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *WhatsappMessageInteractiveActionCardAction) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetButtons

`func (o *WhatsappMessageInteractiveActionCardAction) GetButtons() []WhatsappMessageInteractiveActionCardActionButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *WhatsappMessageInteractiveActionCardAction) GetButtonsOk() (*[]WhatsappMessageInteractiveActionCardActionButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *WhatsappMessageInteractiveActionCardAction) SetButtons(v []WhatsappMessageInteractiveActionCardActionButton)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *WhatsappMessageInteractiveActionCardAction) HasButtons() bool`

HasButtons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


