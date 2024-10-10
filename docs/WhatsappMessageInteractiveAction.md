# WhatsappMessageInteractiveAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buttons** | Pointer to [**[]WhatsappMessageInteractiveActionButton**](WhatsappMessageInteractiveActionButton.md) | Required for Reply Buttons. You can have up to 3 buttons. | [optional] 
**Button** | Pointer to **string** | Required for List Messages. Button content. It cannot be an empty string and must be unique within the message. Emojis are supported, markdown is not. Maximum length: 20 characters. | [optional] 
**CatalogId** | Pointer to **string** | Required for Single Product Messages and Multi-Product Messages. Unique identifier of the Facebook catalog linked to your WhatsApp Business Account. This ID can be retrieved via the [Meta Commerce Manager](https://business.facebook.com/commerce). | [optional] 
**ProductRetailerId** | Pointer to **string** | Required for Single Product Messages and Multi-Product Messages. Unique identifier of the product in a catalog. | [optional] 
**Sections** | Pointer to [**[]WhatsappMessageInteractiveActionSection**](WhatsappMessageInteractiveActionSection.md) | Required for List Messages and Multi-Product Messages. Array of section objects. Minimum of 1, maximum of 10. | [optional] 
**Name** | Pointer to **string** | Action name. Required for Call-To-Action (CTA) buttons. - &#x60;cta_url&#x60;: Use for Call-To-Action (CTA) URL buttons. - &#x60;send_location&#x60;: Use for Location Request buttons. - &#x60;flow&#x60;: Use for Flow buttons. - &#x60;review_and_pay&#x60;: Use for Order Details buttons. - &#x60;review_order&#x60;: Use for Order Status buttons. | [optional] 
**Parameters** | Pointer to [**WhatsappMessageInteractiveActionParameters**](WhatsappMessageInteractiveActionParameters.md) |  | [optional] 

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

`func (o *WhatsappMessageInteractiveAction) GetButtons() []WhatsappMessageInteractiveActionButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *WhatsappMessageInteractiveAction) GetButtonsOk() (*[]WhatsappMessageInteractiveActionButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *WhatsappMessageInteractiveAction) SetButtons(v []WhatsappMessageInteractiveActionButton)`

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

### GetCatalogId

`func (o *WhatsappMessageInteractiveAction) GetCatalogId() string`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *WhatsappMessageInteractiveAction) GetCatalogIdOk() (*string, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *WhatsappMessageInteractiveAction) SetCatalogId(v string)`

SetCatalogId sets CatalogId field to given value.

### HasCatalogId

`func (o *WhatsappMessageInteractiveAction) HasCatalogId() bool`

HasCatalogId returns a boolean if a field has been set.

### GetProductRetailerId

`func (o *WhatsappMessageInteractiveAction) GetProductRetailerId() string`

GetProductRetailerId returns the ProductRetailerId field if non-nil, zero value otherwise.

### GetProductRetailerIdOk

`func (o *WhatsappMessageInteractiveAction) GetProductRetailerIdOk() (*string, bool)`

GetProductRetailerIdOk returns a tuple with the ProductRetailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductRetailerId

`func (o *WhatsappMessageInteractiveAction) SetProductRetailerId(v string)`

SetProductRetailerId sets ProductRetailerId field to given value.

### HasProductRetailerId

`func (o *WhatsappMessageInteractiveAction) HasProductRetailerId() bool`

HasProductRetailerId returns a boolean if a field has been set.

### GetSections

`func (o *WhatsappMessageInteractiveAction) GetSections() []WhatsappMessageInteractiveActionSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *WhatsappMessageInteractiveAction) GetSectionsOk() (*[]WhatsappMessageInteractiveActionSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *WhatsappMessageInteractiveAction) SetSections(v []WhatsappMessageInteractiveActionSection)`

SetSections sets Sections field to given value.

### HasSections

`func (o *WhatsappMessageInteractiveAction) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetName

`func (o *WhatsappMessageInteractiveAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WhatsappMessageInteractiveAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WhatsappMessageInteractiveAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WhatsappMessageInteractiveAction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *WhatsappMessageInteractiveAction) GetParameters() WhatsappMessageInteractiveActionParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *WhatsappMessageInteractiveAction) GetParametersOk() (*WhatsappMessageInteractiveActionParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *WhatsappMessageInteractiveAction) SetParameters(v WhatsappMessageInteractiveActionParameters)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *WhatsappMessageInteractiveAction) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


