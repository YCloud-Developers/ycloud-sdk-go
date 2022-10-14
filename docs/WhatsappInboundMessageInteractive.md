# WhatsappInboundMessageInteractive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of interactive message received. - &#x60;button_reply&#x60;: Sent when a customer clicks a button. - &#x60;list_reply&#x60;: Sent when a customer selects an item from a list. | [optional] 
**ButtonReply** | Pointer to [**WhatsappInboundMessageInteractiveButtonReply**](WhatsappInboundMessageInteractiveButtonReply.md) |  | [optional] 
**ListReply** | Pointer to [**WhatsappInboundMessageInteractiveListReply**](WhatsappInboundMessageInteractiveListReply.md) |  | [optional] 

## Methods

### NewWhatsappInboundMessageInteractive

`func NewWhatsappInboundMessageInteractive() *WhatsappInboundMessageInteractive`

NewWhatsappInboundMessageInteractive instantiates a new WhatsappInboundMessageInteractive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappInboundMessageInteractiveWithDefaults

`func NewWhatsappInboundMessageInteractiveWithDefaults() *WhatsappInboundMessageInteractive`

NewWhatsappInboundMessageInteractiveWithDefaults instantiates a new WhatsappInboundMessageInteractive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WhatsappInboundMessageInteractive) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappInboundMessageInteractive) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappInboundMessageInteractive) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WhatsappInboundMessageInteractive) HasType() bool`

HasType returns a boolean if a field has been set.

### GetButtonReply

`func (o *WhatsappInboundMessageInteractive) GetButtonReply() WhatsappInboundMessageInteractiveButtonReply`

GetButtonReply returns the ButtonReply field if non-nil, zero value otherwise.

### GetButtonReplyOk

`func (o *WhatsappInboundMessageInteractive) GetButtonReplyOk() (*WhatsappInboundMessageInteractiveButtonReply, bool)`

GetButtonReplyOk returns a tuple with the ButtonReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonReply

`func (o *WhatsappInboundMessageInteractive) SetButtonReply(v WhatsappInboundMessageInteractiveButtonReply)`

SetButtonReply sets ButtonReply field to given value.

### HasButtonReply

`func (o *WhatsappInboundMessageInteractive) HasButtonReply() bool`

HasButtonReply returns a boolean if a field has been set.

### GetListReply

`func (o *WhatsappInboundMessageInteractive) GetListReply() WhatsappInboundMessageInteractiveListReply`

GetListReply returns the ListReply field if non-nil, zero value otherwise.

### GetListReplyOk

`func (o *WhatsappInboundMessageInteractive) GetListReplyOk() (*WhatsappInboundMessageInteractiveListReply, bool)`

GetListReplyOk returns a tuple with the ListReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListReply

`func (o *WhatsappInboundMessageInteractive) SetListReply(v WhatsappInboundMessageInteractiveListReply)`

SetListReply sets ListReply field to given value.

### HasListReply

`func (o *WhatsappInboundMessageInteractive) HasListReply() bool`

HasListReply returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


