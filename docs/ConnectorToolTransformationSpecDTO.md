# ConnectorToolTransformationSpecDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** |  |
**Steps** | [**[]ConnectorToolTransformationStepDTO**](ConnectorToolTransformationStepDTO.md) |  |

## Methods

### NewConnectorToolTransformationSpecDTO

`func NewConnectorToolTransformationSpecDTO(version int32, steps []ConnectorToolTransformationStepDTO, ) *ConnectorToolTransformationSpecDTO`

NewConnectorToolTransformationSpecDTO instantiates a new ConnectorToolTransformationSpecDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorToolTransformationSpecDTOWithDefaults

`func NewConnectorToolTransformationSpecDTOWithDefaults() *ConnectorToolTransformationSpecDTO`

NewConnectorToolTransformationSpecDTOWithDefaults instantiates a new ConnectorToolTransformationSpecDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ConnectorToolTransformationSpecDTO) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConnectorToolTransformationSpecDTO) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConnectorToolTransformationSpecDTO) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetSteps

`func (o *ConnectorToolTransformationSpecDTO) GetSteps() []ConnectorToolTransformationStepDTO`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ConnectorToolTransformationSpecDTO) GetStepsOk() (*[]ConnectorToolTransformationStepDTO, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ConnectorToolTransformationSpecDTO) SetSteps(v []ConnectorToolTransformationStepDTO)`

SetSteps sets Steps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
