# WhatsappGroupWebhookSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Setting name. | [optional] 
**Text** | Pointer to **string** | Text value for subject or description updates. | [optional] 
**UpdateSuccessful** | Pointer to **bool** | Whether the setting update succeeded. | [optional] 
**MimeType** | Pointer to **string** | MIME type for profile picture updates. | [optional] 
**Sha256** | Pointer to **string** | SHA-256 hash for profile picture updates. | [optional] 
**Errors** | Pointer to **[]map[string]interface{}** | Errors returned by WhatsApp for this setting. | [optional] 

## Methods

### NewWhatsappGroupWebhookSetting

`func NewWhatsappGroupWebhookSetting() *WhatsappGroupWebhookSetting`

NewWhatsappGroupWebhookSetting instantiates a new WhatsappGroupWebhookSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappGroupWebhookSettingWithDefaults

`func NewWhatsappGroupWebhookSettingWithDefaults() *WhatsappGroupWebhookSetting`

NewWhatsappGroupWebhookSettingWithDefaults instantiates a new WhatsappGroupWebhookSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WhatsappGroupWebhookSetting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WhatsappGroupWebhookSetting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WhatsappGroupWebhookSetting) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WhatsappGroupWebhookSetting) HasName() bool`

HasName returns a boolean if a field has been set.

### GetText

`func (o *WhatsappGroupWebhookSetting) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WhatsappGroupWebhookSetting) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WhatsappGroupWebhookSetting) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *WhatsappGroupWebhookSetting) HasText() bool`

HasText returns a boolean if a field has been set.

### GetUpdateSuccessful

`func (o *WhatsappGroupWebhookSetting) GetUpdateSuccessful() bool`

GetUpdateSuccessful returns the UpdateSuccessful field if non-nil, zero value otherwise.

### GetUpdateSuccessfulOk

`func (o *WhatsappGroupWebhookSetting) GetUpdateSuccessfulOk() (*bool, bool)`

GetUpdateSuccessfulOk returns a tuple with the UpdateSuccessful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSuccessful

`func (o *WhatsappGroupWebhookSetting) SetUpdateSuccessful(v bool)`

SetUpdateSuccessful sets UpdateSuccessful field to given value.

### HasUpdateSuccessful

`func (o *WhatsappGroupWebhookSetting) HasUpdateSuccessful() bool`

HasUpdateSuccessful returns a boolean if a field has been set.

### GetMimeType

`func (o *WhatsappGroupWebhookSetting) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *WhatsappGroupWebhookSetting) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *WhatsappGroupWebhookSetting) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *WhatsappGroupWebhookSetting) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetSha256

`func (o *WhatsappGroupWebhookSetting) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *WhatsappGroupWebhookSetting) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *WhatsappGroupWebhookSetting) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *WhatsappGroupWebhookSetting) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetErrors

`func (o *WhatsappGroupWebhookSetting) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *WhatsappGroupWebhookSetting) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *WhatsappGroupWebhookSetting) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *WhatsappGroupWebhookSetting) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


