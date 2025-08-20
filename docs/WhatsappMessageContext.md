# WhatsappMessageContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | Pointer to **string** | Specifies the &#x60;wamid&#x60; of the message your are replying to. &#x60;wamid&#x60; is the original message ID on WhatsApp&#39;s platform. | [optional] 

## Methods

### NewWhatsappMessageContext

`func NewWhatsappMessageContext() *WhatsappMessageContext`

NewWhatsappMessageContext instantiates a new WhatsappMessageContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageContextWithDefaults

`func NewWhatsappMessageContextWithDefaults() *WhatsappMessageContext`

NewWhatsappMessageContextWithDefaults instantiates a new WhatsappMessageContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *WhatsappMessageContext) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *WhatsappMessageContext) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *WhatsappMessageContext) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *WhatsappMessageContext) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


