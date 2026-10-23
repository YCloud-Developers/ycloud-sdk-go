# ConnectorToolDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional]
**Name** | Pointer to **string** |  | [optional]
**Description** | Pointer to **string** |  | [optional]
**RequestDefinition** | Pointer to [**ConnectorToolRequestDefinitionDTO**](ConnectorToolRequestDefinitionDTO.md) |  | [optional]
**UserAuthRequired** | Pointer to **bool** |  | [optional]
**UserAuthActionConfig** | Pointer to [**ConnectorToolUserAuthActionConfigDTO**](ConnectorToolUserAuthActionConfigDTO.md) |  | [optional]
**TransformationSpec** | Pointer to [**NullableConnectorToolTransformationSpecDTO**](ConnectorToolTransformationSpecDTO.md) |  | [optional]

## Methods

### NewConnectorToolDTO

`func NewConnectorToolDTO() *ConnectorToolDTO`

NewConnectorToolDTO instantiates a new ConnectorToolDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorToolDTOWithDefaults

`func NewConnectorToolDTOWithDefaults() *ConnectorToolDTO`

NewConnectorToolDTOWithDefaults instantiates a new ConnectorToolDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectorToolDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectorToolDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectorToolDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectorToolDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ConnectorToolDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorToolDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorToolDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectorToolDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ConnectorToolDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectorToolDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectorToolDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConnectorToolDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequestDefinition

`func (o *ConnectorToolDTO) GetRequestDefinition() ConnectorToolRequestDefinitionDTO`

GetRequestDefinition returns the RequestDefinition field if non-nil, zero value otherwise.

### GetRequestDefinitionOk

`func (o *ConnectorToolDTO) GetRequestDefinitionOk() (*ConnectorToolRequestDefinitionDTO, bool)`

GetRequestDefinitionOk returns a tuple with the RequestDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDefinition

`func (o *ConnectorToolDTO) SetRequestDefinition(v ConnectorToolRequestDefinitionDTO)`

SetRequestDefinition sets RequestDefinition field to given value.

### HasRequestDefinition

`func (o *ConnectorToolDTO) HasRequestDefinition() bool`

HasRequestDefinition returns a boolean if a field has been set.

### GetUserAuthRequired

`func (o *ConnectorToolDTO) GetUserAuthRequired() bool`

GetUserAuthRequired returns the UserAuthRequired field if non-nil, zero value otherwise.

### GetUserAuthRequiredOk

`func (o *ConnectorToolDTO) GetUserAuthRequiredOk() (*bool, bool)`

GetUserAuthRequiredOk returns a tuple with the UserAuthRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAuthRequired

`func (o *ConnectorToolDTO) SetUserAuthRequired(v bool)`

SetUserAuthRequired sets UserAuthRequired field to given value.

### HasUserAuthRequired

`func (o *ConnectorToolDTO) HasUserAuthRequired() bool`

HasUserAuthRequired returns a boolean if a field has been set.

### GetUserAuthActionConfig

`func (o *ConnectorToolDTO) GetUserAuthActionConfig() ConnectorToolUserAuthActionConfigDTO`

GetUserAuthActionConfig returns the UserAuthActionConfig field if non-nil, zero value otherwise.

### GetUserAuthActionConfigOk

`func (o *ConnectorToolDTO) GetUserAuthActionConfigOk() (*ConnectorToolUserAuthActionConfigDTO, bool)`

GetUserAuthActionConfigOk returns a tuple with the UserAuthActionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAuthActionConfig

`func (o *ConnectorToolDTO) SetUserAuthActionConfig(v ConnectorToolUserAuthActionConfigDTO)`

SetUserAuthActionConfig sets UserAuthActionConfig field to given value.

### HasUserAuthActionConfig

`func (o *ConnectorToolDTO) HasUserAuthActionConfig() bool`

HasUserAuthActionConfig returns a boolean if a field has been set.

### GetTransformationSpec

`func (o *ConnectorToolDTO) GetTransformationSpec() ConnectorToolTransformationSpecDTO`

GetTransformationSpec returns the TransformationSpec field if non-nil, zero value otherwise.

### GetTransformationSpecOk

`func (o *ConnectorToolDTO) GetTransformationSpecOk() (*ConnectorToolTransformationSpecDTO, bool)`

GetTransformationSpecOk returns a tuple with the TransformationSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformationSpec

`func (o *ConnectorToolDTO) SetTransformationSpec(v ConnectorToolTransformationSpecDTO)`

SetTransformationSpec sets TransformationSpec field to given value.

### HasTransformationSpec

`func (o *ConnectorToolDTO) HasTransformationSpec() bool`

HasTransformationSpec returns a boolean if a field has been set.

### SetTransformationSpecNil

`func (o *ConnectorToolDTO) SetTransformationSpecNil(b bool)`

 SetTransformationSpecNil sets the value for TransformationSpec to be an explicit nil

### UnsetTransformationSpec
`func (o *ConnectorToolDTO) UnsetTransformationSpec()`

UnsetTransformationSpec ensures that no value is present for TransformationSpec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
