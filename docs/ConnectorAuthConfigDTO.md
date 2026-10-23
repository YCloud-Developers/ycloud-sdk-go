# ConnectorAuthConfigDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Oauth2ClientCredentials** | Pointer to [**OAuth2ClientCredentialsDTO**](OAuth2ClientCredentialsDTO.md) |  | [optional]
**ApiKey** | Pointer to [**ApiKeyDTO**](ApiKeyDTO.md) |  | [optional]

## Methods

### NewConnectorAuthConfigDTO

`func NewConnectorAuthConfigDTO() *ConnectorAuthConfigDTO`

NewConnectorAuthConfigDTO instantiates a new ConnectorAuthConfigDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorAuthConfigDTOWithDefaults

`func NewConnectorAuthConfigDTOWithDefaults() *ConnectorAuthConfigDTO`

NewConnectorAuthConfigDTOWithDefaults instantiates a new ConnectorAuthConfigDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOauth2ClientCredentials

`func (o *ConnectorAuthConfigDTO) GetOauth2ClientCredentials() OAuth2ClientCredentialsDTO`

GetOauth2ClientCredentials returns the Oauth2ClientCredentials field if non-nil, zero value otherwise.

### GetOauth2ClientCredentialsOk

`func (o *ConnectorAuthConfigDTO) GetOauth2ClientCredentialsOk() (*OAuth2ClientCredentialsDTO, bool)`

GetOauth2ClientCredentialsOk returns a tuple with the Oauth2ClientCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ClientCredentials

`func (o *ConnectorAuthConfigDTO) SetOauth2ClientCredentials(v OAuth2ClientCredentialsDTO)`

SetOauth2ClientCredentials sets Oauth2ClientCredentials field to given value.

### HasOauth2ClientCredentials

`func (o *ConnectorAuthConfigDTO) HasOauth2ClientCredentials() bool`

HasOauth2ClientCredentials returns a boolean if a field has been set.

### GetApiKey

`func (o *ConnectorAuthConfigDTO) GetApiKey() ApiKeyDTO`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *ConnectorAuthConfigDTO) GetApiKeyOk() (*ApiKeyDTO, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *ConnectorAuthConfigDTO) SetApiKey(v ApiKeyDTO)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *ConnectorAuthConfigDTO) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
