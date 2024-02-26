# WhatsappMessageText

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **string** | Required for text messages. The text of the text message which can contain URLs which begin with http:// or https:// and formatting. See available formatting options here. If you include URLs in your text and want to include a preview box in text messages (preview_url: true), make sure the URL starts with http:// or https:// — https:// URLs are preferred. You must include a hostname, since IP addresses will not be matched. Maximum length: 4096 characters. | 
**PreviewUrl** | Pointer to **bool** | By default, WhatsApp recognizes URLs and makes them clickable, but you can also include a preview box with more information about the link. Set this field to true if you want to include a URL preview box. The majority of the time, the receiver will see a URL they can click on when you send an URL, set preview_url to true, and provide a body object with a http or https link. URL previews are only rendered after one of the following has happened: - The business has sent a message template to the user. - The user initiates a conversation with a \&quot;click to chat\&quot; link. - The user adds the business phone number to their address book and initiates a conversation. Default: &#x60;false&#x60;. | [optional] 

## Methods

### NewWhatsappMessageText

`func NewWhatsappMessageText(body string, ) *WhatsappMessageText`

NewWhatsappMessageText instantiates a new WhatsappMessageText object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageTextWithDefaults

`func NewWhatsappMessageTextWithDefaults() *WhatsappMessageText`

NewWhatsappMessageTextWithDefaults instantiates a new WhatsappMessageText object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *WhatsappMessageText) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *WhatsappMessageText) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *WhatsappMessageText) SetBody(v string)`

SetBody sets Body field to given value.


### GetPreviewUrl

`func (o *WhatsappMessageText) GetPreviewUrl() bool`

GetPreviewUrl returns the PreviewUrl field if non-nil, zero value otherwise.

### GetPreviewUrlOk

`func (o *WhatsappMessageText) GetPreviewUrlOk() (*bool, bool)`

GetPreviewUrlOk returns a tuple with the PreviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewUrl

`func (o *WhatsappMessageText) SetPreviewUrl(v bool)`

SetPreviewUrl sets PreviewUrl field to given value.

### HasPreviewUrl

`func (o *WhatsappMessageText) HasPreviewUrl() bool`

HasPreviewUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


