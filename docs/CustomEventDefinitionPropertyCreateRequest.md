# CustomEventDefinitionPropertyCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name of the custom property. | 
**Label** | **string** | The label of the property. | 
**Description** | Pointer to **string** | The description of the property. | [optional] 
**Type** | **string** | Type of the property. - &#x60;STRING&#x60;: Indicates a property that receives plain text strings. - &#x60;NUMBER&#x60;: Indicates a property that receives numeric values with up to one decimal. - &#x60;TIMESTAMP&#x60;: Indicates a property that receives epoch millisecond. - &#x60;URL&#x60;: Indicates a property that receives URLs, formatted as strings starting with &#x60;http://&#x60; or &#x60;https://&#x60;. | 

## Methods

### NewCustomEventDefinitionPropertyCreateRequest

`func NewCustomEventDefinitionPropertyCreateRequest(name string, label string, type_ string, ) *CustomEventDefinitionPropertyCreateRequest`

NewCustomEventDefinitionPropertyCreateRequest instantiates a new CustomEventDefinitionPropertyCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEventDefinitionPropertyCreateRequestWithDefaults

`func NewCustomEventDefinitionPropertyCreateRequestWithDefaults() *CustomEventDefinitionPropertyCreateRequest`

NewCustomEventDefinitionPropertyCreateRequestWithDefaults instantiates a new CustomEventDefinitionPropertyCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomEventDefinitionPropertyCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomEventDefinitionPropertyCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomEventDefinitionPropertyCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *CustomEventDefinitionPropertyCreateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CustomEventDefinitionPropertyCreateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CustomEventDefinitionPropertyCreateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *CustomEventDefinitionPropertyCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomEventDefinitionPropertyCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomEventDefinitionPropertyCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomEventDefinitionPropertyCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *CustomEventDefinitionPropertyCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomEventDefinitionPropertyCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomEventDefinitionPropertyCreateRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


