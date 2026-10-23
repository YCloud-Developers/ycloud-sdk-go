# ApiKeyDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | Pointer to [**[]ApiKeyParamDTO**](ApiKeyParamDTO.md) |  | [optional]
**QueryParams** | Pointer to [**[]ApiKeyParamDTO**](ApiKeyParamDTO.md) |  | [optional]
**BodyParams** | Pointer to [**[]ApiKeyParamDTO**](ApiKeyParamDTO.md) |  | [optional]

## Methods

### NewApiKeyDTO

`func NewApiKeyDTO() *ApiKeyDTO`

NewApiKeyDTO instantiates a new ApiKeyDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyDTOWithDefaults

`func NewApiKeyDTOWithDefaults() *ApiKeyDTO`

NewApiKeyDTOWithDefaults instantiates a new ApiKeyDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *ApiKeyDTO) GetHeaders() []ApiKeyParamDTO`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ApiKeyDTO) GetHeadersOk() (*[]ApiKeyParamDTO, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ApiKeyDTO) SetHeaders(v []ApiKeyParamDTO)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ApiKeyDTO) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetQueryParams

`func (o *ApiKeyDTO) GetQueryParams() []ApiKeyParamDTO`

GetQueryParams returns the QueryParams field if non-nil, zero value otherwise.

### GetQueryParamsOk

`func (o *ApiKeyDTO) GetQueryParamsOk() (*[]ApiKeyParamDTO, bool)`

GetQueryParamsOk returns a tuple with the QueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParams

`func (o *ApiKeyDTO) SetQueryParams(v []ApiKeyParamDTO)`

SetQueryParams sets QueryParams field to given value.

### HasQueryParams

`func (o *ApiKeyDTO) HasQueryParams() bool`

HasQueryParams returns a boolean if a field has been set.

### GetBodyParams

`func (o *ApiKeyDTO) GetBodyParams() []ApiKeyParamDTO`

GetBodyParams returns the BodyParams field if non-nil, zero value otherwise.

### GetBodyParamsOk

`func (o *ApiKeyDTO) GetBodyParamsOk() (*[]ApiKeyParamDTO, bool)`

GetBodyParamsOk returns a tuple with the BodyParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyParams

`func (o *ApiKeyDTO) SetBodyParams(v []ApiKeyParamDTO)`

SetBodyParams sets BodyParams field to given value.

### HasBodyParams

`func (o *ApiKeyDTO) HasBodyParams() bool`

HasBodyParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
