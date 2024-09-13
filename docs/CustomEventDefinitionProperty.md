# CustomEventDefinitionProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the custom property. | [optional] 
**Label** | Pointer to **string** | The label of the property, used for display purposes. | [optional] 
**Description** | Pointer to **string** | The description of the property. | [optional] 
**CreateTime** | Pointer to **time.Time** | The time at which this object is created, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional] 
**Type** | Pointer to **string** | The data type of the property. - &#x60;STRING&#x60;: Indicates a property that receives plain text strings. - &#x60;NUMBER&#x60;: Indicates a property that receives numeric values with up to one decimal. - &#x60;TIMESTAMP&#x60;: Indicates a property that receives epoch millisecond. - &#x60;URL&#x60;: Indicates a property that receives URLs, formatted as strings starting with &#x60;http://&#x60; or &#x60;https://&#x60;. | [optional] 

## Methods

### NewCustomEventDefinitionProperty

`func NewCustomEventDefinitionProperty() *CustomEventDefinitionProperty`

NewCustomEventDefinitionProperty instantiates a new CustomEventDefinitionProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEventDefinitionPropertyWithDefaults

`func NewCustomEventDefinitionPropertyWithDefaults() *CustomEventDefinitionProperty`

NewCustomEventDefinitionPropertyWithDefaults instantiates a new CustomEventDefinitionProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomEventDefinitionProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomEventDefinitionProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomEventDefinitionProperty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomEventDefinitionProperty) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *CustomEventDefinitionProperty) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CustomEventDefinitionProperty) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CustomEventDefinitionProperty) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CustomEventDefinitionProperty) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *CustomEventDefinitionProperty) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomEventDefinitionProperty) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomEventDefinitionProperty) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomEventDefinitionProperty) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreateTime

`func (o *CustomEventDefinitionProperty) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *CustomEventDefinitionProperty) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *CustomEventDefinitionProperty) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *CustomEventDefinitionProperty) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetType

`func (o *CustomEventDefinitionProperty) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomEventDefinitionProperty) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomEventDefinitionProperty) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomEventDefinitionProperty) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


