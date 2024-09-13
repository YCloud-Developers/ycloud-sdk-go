# CustomEventDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the custom event definition. | [optional] 
**Label** | Pointer to **string** | The label of the event definition, used for display purposes. | [optional] 
**Description** | Pointer to **string** | The description of the event definition. | [optional] 
**ObjectType** | Pointer to **string** | Type of the object that the event will be associated with. - &#x60;CONTACT&#x60;: Indicates that the object is a &#x60;contact&#x60;. | [optional] 
**CreateTime** | Pointer to **time.Time** | The time at which this object is created, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional] 
**Properties** | Pointer to [**[]CustomEventDefinitionProperty**](CustomEventDefinitionProperty.md) | The list of property definitions for the event definition. | [optional] 

## Methods

### NewCustomEventDefinition

`func NewCustomEventDefinition() *CustomEventDefinition`

NewCustomEventDefinition instantiates a new CustomEventDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEventDefinitionWithDefaults

`func NewCustomEventDefinitionWithDefaults() *CustomEventDefinition`

NewCustomEventDefinitionWithDefaults instantiates a new CustomEventDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomEventDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomEventDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomEventDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomEventDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *CustomEventDefinition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CustomEventDefinition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CustomEventDefinition) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CustomEventDefinition) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *CustomEventDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomEventDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomEventDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomEventDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetObjectType

`func (o *CustomEventDefinition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CustomEventDefinition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CustomEventDefinition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *CustomEventDefinition) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetCreateTime

`func (o *CustomEventDefinition) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *CustomEventDefinition) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *CustomEventDefinition) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *CustomEventDefinition) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetProperties

`func (o *CustomEventDefinition) GetProperties() []CustomEventDefinitionProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CustomEventDefinition) GetPropertiesOk() (*[]CustomEventDefinitionProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CustomEventDefinition) SetProperties(v []CustomEventDefinitionProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CustomEventDefinition) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


