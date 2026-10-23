# PublicCreateConnectorToolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  |
**Description** | **string** |  |
**RequestDefinition** | [**ConnectorToolRequestDefinitionDTO**](ConnectorToolRequestDefinitionDTO.md) |  |
**UserAuthRequired** | **bool** |  |
**UserAuthActionConfig** | Pointer to [**ConnectorToolUserAuthActionConfigDTO**](ConnectorToolUserAuthActionConfigDTO.md) |  | [optional]
**TransformationSpec** | Pointer to [**NullableConnectorToolTransformationSpecDTO**](ConnectorToolTransformationSpecDTO.md) |  | [optional]

## Methods

### NewPublicCreateConnectorToolRequest

`func NewPublicCreateConnectorToolRequest(name string, description string, requestDefinition ConnectorToolRequestDefinitionDTO, userAuthRequired bool, ) *PublicCreateConnectorToolRequest`

NewPublicCreateConnectorToolRequest instantiates a new PublicCreateConnectorToolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicCreateConnectorToolRequestWithDefaults

`func NewPublicCreateConnectorToolRequestWithDefaults() *PublicCreateConnectorToolRequest`

NewPublicCreateConnectorToolRequestWithDefaults instantiates a new PublicCreateConnectorToolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PublicCreateConnectorToolRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicCreateConnectorToolRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicCreateConnectorToolRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PublicCreateConnectorToolRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PublicCreateConnectorToolRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PublicCreateConnectorToolRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRequestDefinition

`func (o *PublicCreateConnectorToolRequest) GetRequestDefinition() ConnectorToolRequestDefinitionDTO`

GetRequestDefinition returns the RequestDefinition field if non-nil, zero value otherwise.

### GetRequestDefinitionOk

`func (o *PublicCreateConnectorToolRequest) GetRequestDefinitionOk() (*ConnectorToolRequestDefinitionDTO, bool)`

GetRequestDefinitionOk returns a tuple with the RequestDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDefinition

`func (o *PublicCreateConnectorToolRequest) SetRequestDefinition(v ConnectorToolRequestDefinitionDTO)`

SetRequestDefinition sets RequestDefinition field to given value.


### GetUserAuthRequired

`func (o *PublicCreateConnectorToolRequest) GetUserAuthRequired() bool`

GetUserAuthRequired returns the UserAuthRequired field if non-nil, zero value otherwise.

### GetUserAuthRequiredOk

`func (o *PublicCreateConnectorToolRequest) GetUserAuthRequiredOk() (*bool, bool)`

GetUserAuthRequiredOk returns a tuple with the UserAuthRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAuthRequired

`func (o *PublicCreateConnectorToolRequest) SetUserAuthRequired(v bool)`

SetUserAuthRequired sets UserAuthRequired field to given value.


### GetUserAuthActionConfig

`func (o *PublicCreateConnectorToolRequest) GetUserAuthActionConfig() ConnectorToolUserAuthActionConfigDTO`

GetUserAuthActionConfig returns the UserAuthActionConfig field if non-nil, zero value otherwise.

### GetUserAuthActionConfigOk

`func (o *PublicCreateConnectorToolRequest) GetUserAuthActionConfigOk() (*ConnectorToolUserAuthActionConfigDTO, bool)`

GetUserAuthActionConfigOk returns a tuple with the UserAuthActionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAuthActionConfig

`func (o *PublicCreateConnectorToolRequest) SetUserAuthActionConfig(v ConnectorToolUserAuthActionConfigDTO)`

SetUserAuthActionConfig sets UserAuthActionConfig field to given value.

### HasUserAuthActionConfig

`func (o *PublicCreateConnectorToolRequest) HasUserAuthActionConfig() bool`

HasUserAuthActionConfig returns a boolean if a field has been set.

### GetTransformationSpec

`func (o *PublicCreateConnectorToolRequest) GetTransformationSpec() ConnectorToolTransformationSpecDTO`

GetTransformationSpec returns the TransformationSpec field if non-nil, zero value otherwise.

### GetTransformationSpecOk

`func (o *PublicCreateConnectorToolRequest) GetTransformationSpecOk() (*ConnectorToolTransformationSpecDTO, bool)`

GetTransformationSpecOk returns a tuple with the TransformationSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformationSpec

`func (o *PublicCreateConnectorToolRequest) SetTransformationSpec(v ConnectorToolTransformationSpecDTO)`

SetTransformationSpec sets TransformationSpec field to given value.

### HasTransformationSpec

`func (o *PublicCreateConnectorToolRequest) HasTransformationSpec() bool`

HasTransformationSpec returns a boolean if a field has been set.

### SetTransformationSpecNil

`func (o *PublicCreateConnectorToolRequest) SetTransformationSpecNil(b bool)`

 SetTransformationSpecNil sets the value for TransformationSpec to be an explicit nil

### UnsetTransformationSpec
`func (o *PublicCreateConnectorToolRequest) UnsetTransformationSpec()`

UnsetTransformationSpec ensures that no value is present for TransformationSpec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
