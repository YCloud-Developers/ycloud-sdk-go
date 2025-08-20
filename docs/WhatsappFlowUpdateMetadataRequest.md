# WhatsappFlowUpdateMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Flow name. | [optional] 
**Categories** | Pointer to [**[]WhatsappFlowCategory**](WhatsappFlowCategory.md) | Flow categories. | [optional] 

## Methods

### NewWhatsappFlowUpdateMetadataRequest

`func NewWhatsappFlowUpdateMetadataRequest() *WhatsappFlowUpdateMetadataRequest`

NewWhatsappFlowUpdateMetadataRequest instantiates a new WhatsappFlowUpdateMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappFlowUpdateMetadataRequestWithDefaults

`func NewWhatsappFlowUpdateMetadataRequestWithDefaults() *WhatsappFlowUpdateMetadataRequest`

NewWhatsappFlowUpdateMetadataRequestWithDefaults instantiates a new WhatsappFlowUpdateMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WhatsappFlowUpdateMetadataRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WhatsappFlowUpdateMetadataRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WhatsappFlowUpdateMetadataRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WhatsappFlowUpdateMetadataRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategories

`func (o *WhatsappFlowUpdateMetadataRequest) GetCategories() []WhatsappFlowCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *WhatsappFlowUpdateMetadataRequest) GetCategoriesOk() (*[]WhatsappFlowCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *WhatsappFlowUpdateMetadataRequest) SetCategories(v []WhatsappFlowCategory)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *WhatsappFlowUpdateMetadataRequest) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


