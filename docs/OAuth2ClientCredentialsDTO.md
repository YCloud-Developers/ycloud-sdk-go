# OAuth2ClientCredentialsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenUrl** | **string** |  |
**ScopesToRequest** | **[]string** |  |
**TokenRequestContentType** | Pointer to **string** |  | [optional]
**ClientId** | **string** |  |
**ClientSecret** | **string** |  |

## Methods

### NewOAuth2ClientCredentialsDTO

`func NewOAuth2ClientCredentialsDTO(tokenUrl string, scopesToRequest []string, clientId string, clientSecret string, ) *OAuth2ClientCredentialsDTO`

NewOAuth2ClientCredentialsDTO instantiates a new OAuth2ClientCredentialsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ClientCredentialsDTOWithDefaults

`func NewOAuth2ClientCredentialsDTOWithDefaults() *OAuth2ClientCredentialsDTO`

NewOAuth2ClientCredentialsDTOWithDefaults instantiates a new OAuth2ClientCredentialsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenUrl

`func (o *OAuth2ClientCredentialsDTO) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *OAuth2ClientCredentialsDTO) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *OAuth2ClientCredentialsDTO) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.


### GetScopesToRequest

`func (o *OAuth2ClientCredentialsDTO) GetScopesToRequest() []string`

GetScopesToRequest returns the ScopesToRequest field if non-nil, zero value otherwise.

### GetScopesToRequestOk

`func (o *OAuth2ClientCredentialsDTO) GetScopesToRequestOk() (*[]string, bool)`

GetScopesToRequestOk returns a tuple with the ScopesToRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopesToRequest

`func (o *OAuth2ClientCredentialsDTO) SetScopesToRequest(v []string)`

SetScopesToRequest sets ScopesToRequest field to given value.


### GetTokenRequestContentType

`func (o *OAuth2ClientCredentialsDTO) GetTokenRequestContentType() string`

GetTokenRequestContentType returns the TokenRequestContentType field if non-nil, zero value otherwise.

### GetTokenRequestContentTypeOk

`func (o *OAuth2ClientCredentialsDTO) GetTokenRequestContentTypeOk() (*string, bool)`

GetTokenRequestContentTypeOk returns a tuple with the TokenRequestContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRequestContentType

`func (o *OAuth2ClientCredentialsDTO) SetTokenRequestContentType(v string)`

SetTokenRequestContentType sets TokenRequestContentType field to given value.

### HasTokenRequestContentType

`func (o *OAuth2ClientCredentialsDTO) HasTokenRequestContentType() bool`

HasTokenRequestContentType returns a boolean if a field has been set.

### GetClientId

`func (o *OAuth2ClientCredentialsDTO) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuth2ClientCredentialsDTO) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuth2ClientCredentialsDTO) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *OAuth2ClientCredentialsDTO) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuth2ClientCredentialsDTO) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuth2ClientCredentialsDTO) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
