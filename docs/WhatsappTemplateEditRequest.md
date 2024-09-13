# WhatsappTemplateEditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | [**[]WhatsappTemplateComponent**](WhatsappTemplateComponent.md) |  | 
**MessageSendTtlSeconds** | Pointer to **int32** | **Use only for template category is &#x60;AUTHENTICATION&#x60; or &#x60;UTILITY&#x60;.** If we are unable to deliver a message for an amount of time that exceeds its time-to-live, we will stop retrying and drop the message. By default, messages that use an authentication template have a default TTL of **10 minutes**, and messages that use a utility template have a default TTL of **30 days**. Set its value between &#x60;60&#x60; and &#x60;600&#x60; seconds (i.e., 1 to 10 minutes) for authentication templates, or &#x60;60&#x60; and &#x60;3600&#x60; seconds (i.e., 1 to 60 minutes) for utility templates. Alternatively, you can set this value to &#x60;-1&#x60;, which will set a custom TTL of 30 days for either type of template. We encourage you to set a time-to-live for all of your authentication templates, preferably equal to or less than your code expiration time, to ensure your customers only get a message when a code is still usable. Authentication templates created before October 23, 2024, have a default TTL of 30 days. | [optional] 

## Methods

### NewWhatsappTemplateEditRequest

`func NewWhatsappTemplateEditRequest(components []WhatsappTemplateComponent, ) *WhatsappTemplateEditRequest`

NewWhatsappTemplateEditRequest instantiates a new WhatsappTemplateEditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappTemplateEditRequestWithDefaults

`func NewWhatsappTemplateEditRequestWithDefaults() *WhatsappTemplateEditRequest`

NewWhatsappTemplateEditRequestWithDefaults instantiates a new WhatsappTemplateEditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *WhatsappTemplateEditRequest) GetComponents() []WhatsappTemplateComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *WhatsappTemplateEditRequest) GetComponentsOk() (*[]WhatsappTemplateComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *WhatsappTemplateEditRequest) SetComponents(v []WhatsappTemplateComponent)`

SetComponents sets Components field to given value.


### GetMessageSendTtlSeconds

`func (o *WhatsappTemplateEditRequest) GetMessageSendTtlSeconds() int32`

GetMessageSendTtlSeconds returns the MessageSendTtlSeconds field if non-nil, zero value otherwise.

### GetMessageSendTtlSecondsOk

`func (o *WhatsappTemplateEditRequest) GetMessageSendTtlSecondsOk() (*int32, bool)`

GetMessageSendTtlSecondsOk returns a tuple with the MessageSendTtlSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSendTtlSeconds

`func (o *WhatsappTemplateEditRequest) SetMessageSendTtlSeconds(v int32)`

SetMessageSendTtlSeconds sets MessageSendTtlSeconds field to given value.

### HasMessageSendTtlSeconds

`func (o *WhatsappTemplateEditRequest) HasMessageSendTtlSeconds() bool`

HasMessageSendTtlSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


