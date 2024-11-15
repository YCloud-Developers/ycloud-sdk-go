# WhatsappMessageTemplateComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Component type. | 
**SubType** | Pointer to **string** | **Required when type is &#x60;button&#x60;.** Type of button. - &#x60;quick_reply&#x60;: Refers to a previously created quick reply button that allows for the customer to return a predefined message. - &#x60;url&#x60;: Refers to a previously created url button that allows the customer to visit the URL generated by appending the text parameter to the predefined prefix URL in the template. - &#x60;copy_code&#x60;: Refers to a previously created copy code button that allows the customer to copy a text string (defined when the template is sent in a template message) to the device&#39;s clipboard when tapped by the app user. - &#x60;catalog&#x60;: Refers to a previously created catalog button that allows the customer to view your product catalog. - &#x60;mpm&#x60;: Refers to a previously created MPM (multi-product message) button that allows the customer to browser products and sections. - &#x60;flow&#x60;: Refers to a previously created flow button that allows the customer to interact with a [flow](https://developers.facebook.com/docs/whatsapp/flows). - &#x60;order_details&#x60;: Refers to a previously created order details button that allows the customer to view the details of an order. | [optional] 
**Index** | Pointer to **int32** | **Required when &#x60;type&#x60; &#x3D; &#x60;button&#x60;. Not used for the other types.** Indicates order in which button should appear, if the template uses multiple buttons. Buttons are zero-indexed, so setting value to 0 will cause the button to appear first, and another button with an index of 1 will appear next, etc. | [optional] 
**Parameters** | Pointer to [**[]WhatsappMessageTemplateComponentParameter**](WhatsappMessageTemplateComponentParameter.md) | **Required when &#x60;type&#x60; &#x3D; &#x60;button&#x60;, or there are variables in the corresponding template component, or the template &#x60;HEADER&#x60; format is media (&#x60;IMAGE&#x60;, &#x60;VIDEO&#x60;, or &#x60;DOCUMENT&#x60;).** Array of parameter objects with the content of the message. | [optional] 
**Cards** | Pointer to [**[]WhatsappMessageTemplateComponentCard**](WhatsappMessageTemplateComponentCard.md) | Use for &#x60;carousel&#x60; components. Provides card components containing the parameters of the message. | [optional] 

## Methods

### NewWhatsappMessageTemplateComponent

`func NewWhatsappMessageTemplateComponent(type_ string, ) *WhatsappMessageTemplateComponent`

NewWhatsappMessageTemplateComponent instantiates a new WhatsappMessageTemplateComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageTemplateComponentWithDefaults

`func NewWhatsappMessageTemplateComponentWithDefaults() *WhatsappMessageTemplateComponent`

NewWhatsappMessageTemplateComponentWithDefaults instantiates a new WhatsappMessageTemplateComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WhatsappMessageTemplateComponent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappMessageTemplateComponent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappMessageTemplateComponent) SetType(v string)`

SetType sets Type field to given value.


### GetSubType

`func (o *WhatsappMessageTemplateComponent) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *WhatsappMessageTemplateComponent) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *WhatsappMessageTemplateComponent) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *WhatsappMessageTemplateComponent) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetIndex

`func (o *WhatsappMessageTemplateComponent) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *WhatsappMessageTemplateComponent) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *WhatsappMessageTemplateComponent) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *WhatsappMessageTemplateComponent) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetParameters

`func (o *WhatsappMessageTemplateComponent) GetParameters() []WhatsappMessageTemplateComponentParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *WhatsappMessageTemplateComponent) GetParametersOk() (*[]WhatsappMessageTemplateComponentParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *WhatsappMessageTemplateComponent) SetParameters(v []WhatsappMessageTemplateComponentParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *WhatsappMessageTemplateComponent) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetCards

`func (o *WhatsappMessageTemplateComponent) GetCards() []WhatsappMessageTemplateComponentCard`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *WhatsappMessageTemplateComponent) GetCardsOk() (*[]WhatsappMessageTemplateComponentCard, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *WhatsappMessageTemplateComponent) SetCards(v []WhatsappMessageTemplateComponentCard)`

SetCards sets Cards field to given value.

### HasCards

`func (o *WhatsappMessageTemplateComponent) HasCards() bool`

HasCards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


