# ParameterNodeDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional]
**Description** | Pointer to **string** |  | [optional]
**Required** | Pointer to **bool** |  | [optional]
**Binding** | Pointer to [**ParameterBindingDTO**](ParameterBindingDTO.md) |  | [optional]

## Methods

### NewParameterNodeDTO

`func NewParameterNodeDTO() *ParameterNodeDTO`

NewParameterNodeDTO instantiates a new ParameterNodeDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterNodeDTOWithDefaults

`func NewParameterNodeDTOWithDefaults() *ParameterNodeDTO`

NewParameterNodeDTOWithDefaults instantiates a new ParameterNodeDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ParameterNodeDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ParameterNodeDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ParameterNodeDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ParameterNodeDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *ParameterNodeDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ParameterNodeDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ParameterNodeDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ParameterNodeDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequired

`func (o *ParameterNodeDTO) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ParameterNodeDTO) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ParameterNodeDTO) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ParameterNodeDTO) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetBinding

`func (o *ParameterNodeDTO) GetBinding() ParameterBindingDTO`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *ParameterNodeDTO) GetBindingOk() (*ParameterBindingDTO, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *ParameterNodeDTO) SetBinding(v ParameterBindingDTO)`

SetBinding sets Binding field to given value.

### HasBinding

`func (o *ParameterNodeDTO) HasBinding() bool`

HasBinding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
