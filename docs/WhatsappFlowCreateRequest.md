# WhatsappFlowCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WabaId** | **string** | WhatsApp Business Account ID. | 
**Name** | **string** | Flow name. | 
**Categories** | [**[]WhatsappFlowCategory**](WhatsappFlowCategory.md) | Flow categories. | 
**FlowJson** | Pointer to **string** | JSON string of the Flow structure. | [optional] 
**Publish** | Pointer to **bool** | If true, the Flow will be created in PUBLISHED state. | [optional] [default to false]
**CloneFlowId** | Pointer to **string** | ID of source Flow to clone. You must have permission to access the specified Flow. | [optional] 

## Methods

### NewWhatsappFlowCreateRequest

`func NewWhatsappFlowCreateRequest(wabaId string, name string, categories []WhatsappFlowCategory, ) *WhatsappFlowCreateRequest`

NewWhatsappFlowCreateRequest instantiates a new WhatsappFlowCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappFlowCreateRequestWithDefaults

`func NewWhatsappFlowCreateRequestWithDefaults() *WhatsappFlowCreateRequest`

NewWhatsappFlowCreateRequestWithDefaults instantiates a new WhatsappFlowCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWabaId

`func (o *WhatsappFlowCreateRequest) GetWabaId() string`

GetWabaId returns the WabaId field if non-nil, zero value otherwise.

### GetWabaIdOk

`func (o *WhatsappFlowCreateRequest) GetWabaIdOk() (*string, bool)`

GetWabaIdOk returns a tuple with the WabaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWabaId

`func (o *WhatsappFlowCreateRequest) SetWabaId(v string)`

SetWabaId sets WabaId field to given value.


### GetName

`func (o *WhatsappFlowCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WhatsappFlowCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WhatsappFlowCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCategories

`func (o *WhatsappFlowCreateRequest) GetCategories() []WhatsappFlowCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *WhatsappFlowCreateRequest) GetCategoriesOk() (*[]WhatsappFlowCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *WhatsappFlowCreateRequest) SetCategories(v []WhatsappFlowCategory)`

SetCategories sets Categories field to given value.


### GetFlowJson

`func (o *WhatsappFlowCreateRequest) GetFlowJson() string`

GetFlowJson returns the FlowJson field if non-nil, zero value otherwise.

### GetFlowJsonOk

`func (o *WhatsappFlowCreateRequest) GetFlowJsonOk() (*string, bool)`

GetFlowJsonOk returns a tuple with the FlowJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowJson

`func (o *WhatsappFlowCreateRequest) SetFlowJson(v string)`

SetFlowJson sets FlowJson field to given value.

### HasFlowJson

`func (o *WhatsappFlowCreateRequest) HasFlowJson() bool`

HasFlowJson returns a boolean if a field has been set.

### GetPublish

`func (o *WhatsappFlowCreateRequest) GetPublish() bool`

GetPublish returns the Publish field if non-nil, zero value otherwise.

### GetPublishOk

`func (o *WhatsappFlowCreateRequest) GetPublishOk() (*bool, bool)`

GetPublishOk returns a tuple with the Publish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublish

`func (o *WhatsappFlowCreateRequest) SetPublish(v bool)`

SetPublish sets Publish field to given value.

### HasPublish

`func (o *WhatsappFlowCreateRequest) HasPublish() bool`

HasPublish returns a boolean if a field has been set.

### GetCloneFlowId

`func (o *WhatsappFlowCreateRequest) GetCloneFlowId() string`

GetCloneFlowId returns the CloneFlowId field if non-nil, zero value otherwise.

### GetCloneFlowIdOk

`func (o *WhatsappFlowCreateRequest) GetCloneFlowIdOk() (*string, bool)`

GetCloneFlowIdOk returns a tuple with the CloneFlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneFlowId

`func (o *WhatsappFlowCreateRequest) SetCloneFlowId(v string)`

SetCloneFlowId sets CloneFlowId field to given value.

### HasCloneFlowId

`func (o *WhatsappFlowCreateRequest) HasCloneFlowId() bool`

HasCloneFlowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


