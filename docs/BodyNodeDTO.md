# BodyNodeDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional]
**Description** | Pointer to **string** |  | [optional]
**Required** | Pointer to **[]string** |  | [optional]
**Properties** | Pointer to **map[string]string** |  | [optional]
**Items** | Pointer to **string** |  | [optional]
**Binding** | Pointer to [**ParameterBindingDTO**](ParameterBindingDTO.md) |  | [optional]

## Methods

### NewBodyNodeDTO

`func NewBodyNodeDTO() *BodyNodeDTO`

NewBodyNodeDTO instantiates a new BodyNodeDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBodyNodeDTOWithDefaults

`func NewBodyNodeDTOWithDefaults() *BodyNodeDTO`

NewBodyNodeDTOWithDefaults instantiates a new BodyNodeDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BodyNodeDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BodyNodeDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BodyNodeDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BodyNodeDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *BodyNodeDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BodyNodeDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BodyNodeDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BodyNodeDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequired

`func (o *BodyNodeDTO) GetRequired() []string`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *BodyNodeDTO) GetRequiredOk() (*[]string, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *BodyNodeDTO) SetRequired(v []string)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *BodyNodeDTO) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetProperties

`func (o *BodyNodeDTO) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BodyNodeDTO) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BodyNodeDTO) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BodyNodeDTO) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetItems

`func (o *BodyNodeDTO) GetItems() string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BodyNodeDTO) GetItemsOk() (*string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BodyNodeDTO) SetItems(v string)`

SetItems sets Items field to given value.

### HasItems

`func (o *BodyNodeDTO) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetBinding

`func (o *BodyNodeDTO) GetBinding() ParameterBindingDTO`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *BodyNodeDTO) GetBindingOk() (*ParameterBindingDTO, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *BodyNodeDTO) SetBinding(v ParameterBindingDTO)`

SetBinding sets Binding field to given value.

### HasBinding

`func (o *BodyNodeDTO) HasBinding() bool`

HasBinding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
