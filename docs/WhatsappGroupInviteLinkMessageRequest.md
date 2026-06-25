# WhatsappGroupInviteLinkMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | Pointer to **string** | The recipient&#39;s phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. Required when &#x60;recipient&#x60; is not provided. | [optional]
**Recipient** | Pointer to **string** | The recipient&#39;s WhatsApp Business-scoped user ID (BSUID) or parent BSUID. Required when &#x60;to&#x60; is not provided. | [optional]
**TemplateName** | **string** | The name of the approved WhatsApp template. |
**LanguageCode** | **string** | The template language code. |
**Parameters** | [**[]WhatsappMessageTemplateComponentParameter**](WhatsappMessageTemplateComponentParameter.md) | Template body parameters in template variable order. Must include one group invite link parameter with &#x60;type&#x3D;group_id&#x60; and &#x60;group_id&#x3D;&lt;groupId&gt;&#x60;. |

## Methods

### NewWhatsappGroupInviteLinkMessageRequest

`func NewWhatsappGroupInviteLinkMessageRequest(templateName string, languageCode string, parameters []WhatsappMessageTemplateComponentParameter, ) *WhatsappGroupInviteLinkMessageRequest`

NewWhatsappGroupInviteLinkMessageRequest instantiates a new WhatsappGroupInviteLinkMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappGroupInviteLinkMessageRequestWithDefaults

`func NewWhatsappGroupInviteLinkMessageRequestWithDefaults() *WhatsappGroupInviteLinkMessageRequest`

NewWhatsappGroupInviteLinkMessageRequestWithDefaults instantiates a new WhatsappGroupInviteLinkMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *WhatsappGroupInviteLinkMessageRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *WhatsappGroupInviteLinkMessageRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *WhatsappGroupInviteLinkMessageRequest) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *WhatsappGroupInviteLinkMessageRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetRecipient

`func (o *WhatsappGroupInviteLinkMessageRequest) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *WhatsappGroupInviteLinkMessageRequest) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *WhatsappGroupInviteLinkMessageRequest) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *WhatsappGroupInviteLinkMessageRequest) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetTemplateName

`func (o *WhatsappGroupInviteLinkMessageRequest) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *WhatsappGroupInviteLinkMessageRequest) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *WhatsappGroupInviteLinkMessageRequest) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.


### GetLanguageCode

`func (o *WhatsappGroupInviteLinkMessageRequest) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *WhatsappGroupInviteLinkMessageRequest) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *WhatsappGroupInviteLinkMessageRequest) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.


### GetParameters

`func (o *WhatsappGroupInviteLinkMessageRequest) GetParameters() []WhatsappMessageTemplateComponentParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *WhatsappGroupInviteLinkMessageRequest) GetParametersOk() (*[]WhatsappMessageTemplateComponentParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *WhatsappGroupInviteLinkMessageRequest) SetParameters(v []WhatsappMessageTemplateComponentParameter)`

SetParameters sets Parameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


