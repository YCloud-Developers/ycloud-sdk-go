# ApiKeyParamDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | **string** |  |
**Value** | **string** |  |
**Prefix** | Pointer to **string** |  | [optional]

## Methods

### NewApiKeyParamDTO

`func NewApiKeyParamDTO(fieldName string, value string, ) *ApiKeyParamDTO`

NewApiKeyParamDTO instantiates a new ApiKeyParamDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyParamDTOWithDefaults

`func NewApiKeyParamDTOWithDefaults() *ApiKeyParamDTO`

NewApiKeyParamDTOWithDefaults instantiates a new ApiKeyParamDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *ApiKeyParamDTO) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *ApiKeyParamDTO) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *ApiKeyParamDTO) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetValue

`func (o *ApiKeyParamDTO) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApiKeyParamDTO) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApiKeyParamDTO) SetValue(v string)`

SetValue sets Value field to given value.


### GetPrefix

`func (o *ApiKeyParamDTO) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ApiKeyParamDTO) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ApiKeyParamDTO) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ApiKeyParamDTO) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
