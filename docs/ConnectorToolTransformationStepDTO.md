# ConnectorToolTransformationStepDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  |
**Target** | Pointer to **NullableString** |  | [optional]
**Params** | Pointer to **NullableString** | Kind-specific parameters as a JSON object encoded in a string. A math expression supports at most four nested levels. Parameters are forwarded without evaluation or rewriting. | [optional]

## Methods

### NewConnectorToolTransformationStepDTO

`func NewConnectorToolTransformationStepDTO(kind string, ) *ConnectorToolTransformationStepDTO`

NewConnectorToolTransformationStepDTO instantiates a new ConnectorToolTransformationStepDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorToolTransformationStepDTOWithDefaults

`func NewConnectorToolTransformationStepDTOWithDefaults() *ConnectorToolTransformationStepDTO`

NewConnectorToolTransformationStepDTOWithDefaults instantiates a new ConnectorToolTransformationStepDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ConnectorToolTransformationStepDTO) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConnectorToolTransformationStepDTO) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConnectorToolTransformationStepDTO) SetKind(v string)`

SetKind sets Kind field to given value.


### GetTarget

`func (o *ConnectorToolTransformationStepDTO) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ConnectorToolTransformationStepDTO) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ConnectorToolTransformationStepDTO) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ConnectorToolTransformationStepDTO) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *ConnectorToolTransformationStepDTO) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *ConnectorToolTransformationStepDTO) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetParams

`func (o *ConnectorToolTransformationStepDTO) GetParams() string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ConnectorToolTransformationStepDTO) GetParamsOk() (*string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ConnectorToolTransformationStepDTO) SetParams(v string)`

SetParams sets Params field to given value.

### HasParams

`func (o *ConnectorToolTransformationStepDTO) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParamsNil

`func (o *ConnectorToolTransformationStepDTO) SetParamsNil(b bool)`

 SetParamsNil sets the value for Params to be an explicit nil

### UnsetParams
`func (o *ConnectorToolTransformationStepDTO) UnsetParams()`

UnsetParams ensures that no value is present for Params, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
