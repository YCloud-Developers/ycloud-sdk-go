# WhatsappMessageInteractiveActionButton

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Only supported type is &#x60;reply&#x60; (for Reply Button). | [optional] 
**Reply** | Pointer to [**WhatsappMessageInteractiveActionButtonReply**](WhatsappMessageInteractiveActionButtonReply.md) |  | [optional] 

## Methods

### NewWhatsappMessageInteractiveActionButton

`func NewWhatsappMessageInteractiveActionButton() *WhatsappMessageInteractiveActionButton`

NewWhatsappMessageInteractiveActionButton instantiates a new WhatsappMessageInteractiveActionButton object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageInteractiveActionButtonWithDefaults

`func NewWhatsappMessageInteractiveActionButtonWithDefaults() *WhatsappMessageInteractiveActionButton`

NewWhatsappMessageInteractiveActionButtonWithDefaults instantiates a new WhatsappMessageInteractiveActionButton object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WhatsappMessageInteractiveActionButton) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappMessageInteractiveActionButton) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappMessageInteractiveActionButton) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WhatsappMessageInteractiveActionButton) HasType() bool`

HasType returns a boolean if a field has been set.

### GetReply

`func (o *WhatsappMessageInteractiveActionButton) GetReply() WhatsappMessageInteractiveActionButtonReply`

GetReply returns the Reply field if non-nil, zero value otherwise.

### GetReplyOk

`func (o *WhatsappMessageInteractiveActionButton) GetReplyOk() (*WhatsappMessageInteractiveActionButtonReply, bool)`

GetReplyOk returns a tuple with the Reply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReply

`func (o *WhatsappMessageInteractiveActionButton) SetReply(v WhatsappMessageInteractiveActionButtonReply)`

SetReply sets Reply field to given value.

### HasReply

`func (o *WhatsappMessageInteractiveActionButton) HasReply() bool`

HasReply returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


