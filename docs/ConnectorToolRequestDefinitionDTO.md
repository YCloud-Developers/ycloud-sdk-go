# ConnectorToolRequestDefinitionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** |  |
**Path** | **string** |  |
**PathParameters** | Pointer to [**map[string]ParameterNodeDTO**](ParameterNodeDTO.md) |  | [optional]
**QueryParameters** | Pointer to [**map[string]ParameterNodeDTO**](ParameterNodeDTO.md) |  | [optional]
**Headers** | Pointer to [**map[string]ParameterNodeDTO**](ParameterNodeDTO.md) |  | [optional]
**Body** | Pointer to [**RequestBodyDefinitionDTO**](RequestBodyDefinitionDTO.md) |  | [optional]

## Methods

### NewConnectorToolRequestDefinitionDTO

`func NewConnectorToolRequestDefinitionDTO(method string, path string, ) *ConnectorToolRequestDefinitionDTO`

NewConnectorToolRequestDefinitionDTO instantiates a new ConnectorToolRequestDefinitionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorToolRequestDefinitionDTOWithDefaults

`func NewConnectorToolRequestDefinitionDTOWithDefaults() *ConnectorToolRequestDefinitionDTO`

NewConnectorToolRequestDefinitionDTOWithDefaults instantiates a new ConnectorToolRequestDefinitionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *ConnectorToolRequestDefinitionDTO) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ConnectorToolRequestDefinitionDTO) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ConnectorToolRequestDefinitionDTO) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetPath

`func (o *ConnectorToolRequestDefinitionDTO) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ConnectorToolRequestDefinitionDTO) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ConnectorToolRequestDefinitionDTO) SetPath(v string)`

SetPath sets Path field to given value.


### GetPathParameters

`func (o *ConnectorToolRequestDefinitionDTO) GetPathParameters() map[string]ParameterNodeDTO`

GetPathParameters returns the PathParameters field if non-nil, zero value otherwise.

### GetPathParametersOk

`func (o *ConnectorToolRequestDefinitionDTO) GetPathParametersOk() (*map[string]ParameterNodeDTO, bool)`

GetPathParametersOk returns a tuple with the PathParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathParameters

`func (o *ConnectorToolRequestDefinitionDTO) SetPathParameters(v map[string]ParameterNodeDTO)`

SetPathParameters sets PathParameters field to given value.

### HasPathParameters

`func (o *ConnectorToolRequestDefinitionDTO) HasPathParameters() bool`

HasPathParameters returns a boolean if a field has been set.

### GetQueryParameters

`func (o *ConnectorToolRequestDefinitionDTO) GetQueryParameters() map[string]ParameterNodeDTO`

GetQueryParameters returns the QueryParameters field if non-nil, zero value otherwise.

### GetQueryParametersOk

`func (o *ConnectorToolRequestDefinitionDTO) GetQueryParametersOk() (*map[string]ParameterNodeDTO, bool)`

GetQueryParametersOk returns a tuple with the QueryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParameters

`func (o *ConnectorToolRequestDefinitionDTO) SetQueryParameters(v map[string]ParameterNodeDTO)`

SetQueryParameters sets QueryParameters field to given value.

### HasQueryParameters

`func (o *ConnectorToolRequestDefinitionDTO) HasQueryParameters() bool`

HasQueryParameters returns a boolean if a field has been set.

### GetHeaders

`func (o *ConnectorToolRequestDefinitionDTO) GetHeaders() map[string]ParameterNodeDTO`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ConnectorToolRequestDefinitionDTO) GetHeadersOk() (*map[string]ParameterNodeDTO, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ConnectorToolRequestDefinitionDTO) SetHeaders(v map[string]ParameterNodeDTO)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ConnectorToolRequestDefinitionDTO) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetBody

`func (o *ConnectorToolRequestDefinitionDTO) GetBody() RequestBodyDefinitionDTO`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ConnectorToolRequestDefinitionDTO) GetBodyOk() (*RequestBodyDefinitionDTO, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ConnectorToolRequestDefinitionDTO) SetBody(v RequestBodyDefinitionDTO)`

SetBody sets Body field to given value.

### HasBody

`func (o *ConnectorToolRequestDefinitionDTO) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
