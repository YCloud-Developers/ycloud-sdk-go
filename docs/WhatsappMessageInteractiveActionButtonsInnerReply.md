# WhatsappMessageInteractiveActionButtonsInnerReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Button title. It cannot be an empty string and must be unique within the message. Emojis are supported, markdown is not. | [optional] 
**Id** | Pointer to **string** | Unique identifier for your button. This ID is returned in the webhook when the button is clicked by the user. | [optional] 

## Methods

### NewWhatsappMessageInteractiveActionButtonsInnerReply

`func NewWhatsappMessageInteractiveActionButtonsInnerReply() *WhatsappMessageInteractiveActionButtonsInnerReply`

NewWhatsappMessageInteractiveActionButtonsInnerReply instantiates a new WhatsappMessageInteractiveActionButtonsInnerReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageInteractiveActionButtonsInnerReplyWithDefaults

`func NewWhatsappMessageInteractiveActionButtonsInnerReplyWithDefaults() *WhatsappMessageInteractiveActionButtonsInnerReply`

NewWhatsappMessageInteractiveActionButtonsInnerReplyWithDefaults instantiates a new WhatsappMessageInteractiveActionButtonsInnerReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *WhatsappMessageInteractiveActionButtonsInnerReply) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WhatsappMessageInteractiveActionButtonsInnerReply) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WhatsappMessageInteractiveActionButtonsInnerReply) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WhatsappMessageInteractiveActionButtonsInnerReply) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetId

`func (o *WhatsappMessageInteractiveActionButtonsInnerReply) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WhatsappMessageInteractiveActionButtonsInnerReply) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WhatsappMessageInteractiveActionButtonsInnerReply) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WhatsappMessageInteractiveActionButtonsInnerReply) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


