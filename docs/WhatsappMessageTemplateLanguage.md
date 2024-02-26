# WhatsappMessageTemplateLanguage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | The code of the language or locale to use. Accepts both language and language_locale formats (e.g., en and en_US). See [Supported Languages](https://developers.facebook.com/docs/whatsapp/api/messages/message-templates#supported-languages) for all codes. | 
**Policy** | Pointer to **string** | The language policy the message should follow. Default (and only supported option): &#x60;deterministic&#x60;, which means that WhatsApp delivers the message template in exactly the language and locale asked for. | [optional] 

## Methods

### NewWhatsappMessageTemplateLanguage

`func NewWhatsappMessageTemplateLanguage(code string, ) *WhatsappMessageTemplateLanguage`

NewWhatsappMessageTemplateLanguage instantiates a new WhatsappMessageTemplateLanguage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageTemplateLanguageWithDefaults

`func NewWhatsappMessageTemplateLanguageWithDefaults() *WhatsappMessageTemplateLanguage`

NewWhatsappMessageTemplateLanguageWithDefaults instantiates a new WhatsappMessageTemplateLanguage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *WhatsappMessageTemplateLanguage) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *WhatsappMessageTemplateLanguage) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *WhatsappMessageTemplateLanguage) SetCode(v string)`

SetCode sets Code field to given value.


### GetPolicy

`func (o *WhatsappMessageTemplateLanguage) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *WhatsappMessageTemplateLanguage) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *WhatsappMessageTemplateLanguage) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *WhatsappMessageTemplateLanguage) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


