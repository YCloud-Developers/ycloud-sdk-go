# CustomEventDefinitionCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name of the custom event. | 
**Label** | **string** | The label of the custom event. | 
**Description** | Pointer to **string** | The description of the event. | [optional] 
**ObjectType** | **string** | Type of the object that the event will be associated with. - &#x60;CONTACT&#x60;: Indicates that the object is a &#x60;contact&#x60;. | 
**Properties** | Pointer to [**[]CustomEventDefinitionPropertyCreateRequest**](CustomEventDefinitionPropertyCreateRequest.md) | A list of property definitions for the event. | [optional] 

## Methods

### NewCustomEventDefinitionCreateRequest

`func NewCustomEventDefinitionCreateRequest(name string, label string, objectType string, ) *CustomEventDefinitionCreateRequest`

NewCustomEventDefinitionCreateRequest instantiates a new CustomEventDefinitionCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEventDefinitionCreateRequestWithDefaults

`func NewCustomEventDefinitionCreateRequestWithDefaults() *CustomEventDefinitionCreateRequest`

NewCustomEventDefinitionCreateRequestWithDefaults instantiates a new CustomEventDefinitionCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomEventDefinitionCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomEventDefinitionCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomEventDefinitionCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *CustomEventDefinitionCreateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CustomEventDefinitionCreateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CustomEventDefinitionCreateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *CustomEventDefinitionCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomEventDefinitionCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomEventDefinitionCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomEventDefinitionCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetObjectType

`func (o *CustomEventDefinitionCreateRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CustomEventDefinitionCreateRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CustomEventDefinitionCreateRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetProperties

`func (o *CustomEventDefinitionCreateRequest) GetProperties() []CustomEventDefinitionPropertyCreateRequest`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CustomEventDefinitionCreateRequest) GetPropertiesOk() (*[]CustomEventDefinitionPropertyCreateRequest, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CustomEventDefinitionCreateRequest) SetProperties(v []CustomEventDefinitionPropertyCreateRequest)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CustomEventDefinitionCreateRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


