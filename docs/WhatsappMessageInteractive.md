# WhatsappMessageInteractive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | **Required.** The type of interactive message you want to send. - &#x60;button&#x60;: Use for Reply Buttons. - &#x60;list&#x60;: Use for List Messages. - &#x60;cta_url&#x60;: Use for Call-To-Action (CTA) URL Button Messages. - &#x60;product&#x60;: Use for Single Product Messages. - &#x60;product_list&#x60;: Use for Multi-Product Messages. - &#x60;catalog_message&#x60;: Use for Catalog Messages. | [optional] 
**Action** | Pointer to [**WhatsappMessageInteractiveAction**](WhatsappMessageInteractiveAction.md) |  | [optional] 
**Body** | Pointer to [**WhatsappMessageInteractiveBody**](WhatsappMessageInteractiveBody.md) |  | [optional] 
**Header** | Pointer to [**WhatsappMessageInteractiveHeader**](WhatsappMessageInteractiveHeader.md) |  | [optional] 
**Footer** | Pointer to [**WhatsappMessageInteractiveFooter**](WhatsappMessageInteractiveFooter.md) |  | [optional] 

## Methods

### NewWhatsappMessageInteractive

`func NewWhatsappMessageInteractive() *WhatsappMessageInteractive`

NewWhatsappMessageInteractive instantiates a new WhatsappMessageInteractive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageInteractiveWithDefaults

`func NewWhatsappMessageInteractiveWithDefaults() *WhatsappMessageInteractive`

NewWhatsappMessageInteractiveWithDefaults instantiates a new WhatsappMessageInteractive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WhatsappMessageInteractive) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsappMessageInteractive) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsappMessageInteractive) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WhatsappMessageInteractive) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAction

`func (o *WhatsappMessageInteractive) GetAction() WhatsappMessageInteractiveAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WhatsappMessageInteractive) GetActionOk() (*WhatsappMessageInteractiveAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WhatsappMessageInteractive) SetAction(v WhatsappMessageInteractiveAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *WhatsappMessageInteractive) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetBody

`func (o *WhatsappMessageInteractive) GetBody() WhatsappMessageInteractiveBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *WhatsappMessageInteractive) GetBodyOk() (*WhatsappMessageInteractiveBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *WhatsappMessageInteractive) SetBody(v WhatsappMessageInteractiveBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *WhatsappMessageInteractive) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetHeader

`func (o *WhatsappMessageInteractive) GetHeader() WhatsappMessageInteractiveHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *WhatsappMessageInteractive) GetHeaderOk() (*WhatsappMessageInteractiveHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *WhatsappMessageInteractive) SetHeader(v WhatsappMessageInteractiveHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *WhatsappMessageInteractive) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetFooter

`func (o *WhatsappMessageInteractive) GetFooter() WhatsappMessageInteractiveFooter`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *WhatsappMessageInteractive) GetFooterOk() (*WhatsappMessageInteractiveFooter, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *WhatsappMessageInteractive) SetFooter(v WhatsappMessageInteractiveFooter)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *WhatsappMessageInteractive) HasFooter() bool`

HasFooter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


