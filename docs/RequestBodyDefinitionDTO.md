# RequestBodyDefinitionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | **string** |  |
**Params** | [**map[string]BodyNodeDTO**](BodyNodeDTO.md) |  |
**Required** | Pointer to **[]string** |  | [optional]

## Methods

### NewRequestBodyDefinitionDTO

`func NewRequestBodyDefinitionDTO(contentType string, params map[string]BodyNodeDTO, ) *RequestBodyDefinitionDTO`

NewRequestBodyDefinitionDTO instantiates a new RequestBodyDefinitionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestBodyDefinitionDTOWithDefaults

`func NewRequestBodyDefinitionDTOWithDefaults() *RequestBodyDefinitionDTO`

NewRequestBodyDefinitionDTOWithDefaults instantiates a new RequestBodyDefinitionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *RequestBodyDefinitionDTO) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *RequestBodyDefinitionDTO) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *RequestBodyDefinitionDTO) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetParams

`func (o *RequestBodyDefinitionDTO) GetParams() map[string]BodyNodeDTO`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *RequestBodyDefinitionDTO) GetParamsOk() (*map[string]BodyNodeDTO, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *RequestBodyDefinitionDTO) SetParams(v map[string]BodyNodeDTO)`

SetParams sets Params field to given value.


### GetRequired

`func (o *RequestBodyDefinitionDTO) GetRequired() []string`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *RequestBodyDefinitionDTO) GetRequiredOk() (*[]string, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *RequestBodyDefinitionDTO) SetRequired(v []string)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *RequestBodyDefinitionDTO) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
