# WhatsappTemplateComponentExample

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BodyText** | Pointer to **[][]string** | Sample values for variables in &#x60;text&#x60; of a &#x60;BODY&#x60; component. | [optional] 
**HeaderText** | Pointer to **[]string** | Sample value for the variable in &#x60;text&#x60; of a &#x60;HEADER&#x60; component. | [optional] 
**HeaderUrl** | Pointer to **[]string** | Sample media URL for a &#x60;HEADER&#x60; component whose format is one of &#x60;IMAGE&#x60;, &#x60;VIDEO&#x60;, or &#x60;DOCUMENT&#x60;. Supported types: - For &#x60;IMAGE&#x60;, the URL must end with one of &#x60;.jpg&#x60;, &#x60;.jpeg&#x60;, or &#x60;.png&#x60;, size limit is 5MB. - For &#x60;VIDEO&#x60;, the URL must end with &#x60;.mp4&#x60;, size limit is 16MB. - For &#x60;DOCUMENT&#x60;, the URL must end with &#x60;.pdf&#x60;, size limit is 100MB. | [optional] 

## Methods

### NewWhatsappTemplateComponentExample

`func NewWhatsappTemplateComponentExample() *WhatsappTemplateComponentExample`

NewWhatsappTemplateComponentExample instantiates a new WhatsappTemplateComponentExample object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappTemplateComponentExampleWithDefaults

`func NewWhatsappTemplateComponentExampleWithDefaults() *WhatsappTemplateComponentExample`

NewWhatsappTemplateComponentExampleWithDefaults instantiates a new WhatsappTemplateComponentExample object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBodyText

`func (o *WhatsappTemplateComponentExample) GetBodyText() [][]string`

GetBodyText returns the BodyText field if non-nil, zero value otherwise.

### GetBodyTextOk

`func (o *WhatsappTemplateComponentExample) GetBodyTextOk() (*[][]string, bool)`

GetBodyTextOk returns a tuple with the BodyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyText

`func (o *WhatsappTemplateComponentExample) SetBodyText(v [][]string)`

SetBodyText sets BodyText field to given value.

### HasBodyText

`func (o *WhatsappTemplateComponentExample) HasBodyText() bool`

HasBodyText returns a boolean if a field has been set.

### GetHeaderText

`func (o *WhatsappTemplateComponentExample) GetHeaderText() []string`

GetHeaderText returns the HeaderText field if non-nil, zero value otherwise.

### GetHeaderTextOk

`func (o *WhatsappTemplateComponentExample) GetHeaderTextOk() (*[]string, bool)`

GetHeaderTextOk returns a tuple with the HeaderText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderText

`func (o *WhatsappTemplateComponentExample) SetHeaderText(v []string)`

SetHeaderText sets HeaderText field to given value.

### HasHeaderText

`func (o *WhatsappTemplateComponentExample) HasHeaderText() bool`

HasHeaderText returns a boolean if a field has been set.

### GetHeaderUrl

`func (o *WhatsappTemplateComponentExample) GetHeaderUrl() []string`

GetHeaderUrl returns the HeaderUrl field if non-nil, zero value otherwise.

### GetHeaderUrlOk

`func (o *WhatsappTemplateComponentExample) GetHeaderUrlOk() (*[]string, bool)`

GetHeaderUrlOk returns a tuple with the HeaderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderUrl

`func (o *WhatsappTemplateComponentExample) SetHeaderUrl(v []string)`

SetHeaderUrl sets HeaderUrl field to given value.

### HasHeaderUrl

`func (o *WhatsappTemplateComponentExample) HasHeaderUrl() bool`

HasHeaderUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


