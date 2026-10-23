# PublicCreateConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  |
**Description** | **string** |  |
**BaseUrl** | **string** |  |
**AuthType** | **string** |  |
**AuthConfig** | Pointer to [**ConnectorAuthConfigDTO**](ConnectorAuthConfigDTO.md) |  | [optional]
**UserAuthInjectionConfig** | Pointer to [**UserAuthInjectionConfigDTO**](UserAuthInjectionConfigDTO.md) |  | [optional]
**RequiresCertificate** | Pointer to **bool** |  | [optional]
**ConnectorProtocol** | Pointer to **string** | Connector protocol. Omit on creation to use Meta’s HTTP default. The protocol cannot change after creation; omit on update to preserve it. | [optional]

## Methods

### NewPublicCreateConnectorRequest

`func NewPublicCreateConnectorRequest(name string, description string, baseUrl string, authType string, ) *PublicCreateConnectorRequest`

NewPublicCreateConnectorRequest instantiates a new PublicCreateConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicCreateConnectorRequestWithDefaults

`func NewPublicCreateConnectorRequestWithDefaults() *PublicCreateConnectorRequest`

NewPublicCreateConnectorRequestWithDefaults instantiates a new PublicCreateConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PublicCreateConnectorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicCreateConnectorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicCreateConnectorRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PublicCreateConnectorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PublicCreateConnectorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PublicCreateConnectorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetBaseUrl

`func (o *PublicCreateConnectorRequest) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *PublicCreateConnectorRequest) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *PublicCreateConnectorRequest) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetAuthType

`func (o *PublicCreateConnectorRequest) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *PublicCreateConnectorRequest) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *PublicCreateConnectorRequest) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### GetAuthConfig

`func (o *PublicCreateConnectorRequest) GetAuthConfig() ConnectorAuthConfigDTO`

GetAuthConfig returns the AuthConfig field if non-nil, zero value otherwise.

### GetAuthConfigOk

`func (o *PublicCreateConnectorRequest) GetAuthConfigOk() (*ConnectorAuthConfigDTO, bool)`

GetAuthConfigOk returns a tuple with the AuthConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthConfig

`func (o *PublicCreateConnectorRequest) SetAuthConfig(v ConnectorAuthConfigDTO)`

SetAuthConfig sets AuthConfig field to given value.

### HasAuthConfig

`func (o *PublicCreateConnectorRequest) HasAuthConfig() bool`

HasAuthConfig returns a boolean if a field has been set.

### GetUserAuthInjectionConfig

`func (o *PublicCreateConnectorRequest) GetUserAuthInjectionConfig() UserAuthInjectionConfigDTO`

GetUserAuthInjectionConfig returns the UserAuthInjectionConfig field if non-nil, zero value otherwise.

### GetUserAuthInjectionConfigOk

`func (o *PublicCreateConnectorRequest) GetUserAuthInjectionConfigOk() (*UserAuthInjectionConfigDTO, bool)`

GetUserAuthInjectionConfigOk returns a tuple with the UserAuthInjectionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAuthInjectionConfig

`func (o *PublicCreateConnectorRequest) SetUserAuthInjectionConfig(v UserAuthInjectionConfigDTO)`

SetUserAuthInjectionConfig sets UserAuthInjectionConfig field to given value.

### HasUserAuthInjectionConfig

`func (o *PublicCreateConnectorRequest) HasUserAuthInjectionConfig() bool`

HasUserAuthInjectionConfig returns a boolean if a field has been set.

### GetRequiresCertificate

`func (o *PublicCreateConnectorRequest) GetRequiresCertificate() bool`

GetRequiresCertificate returns the RequiresCertificate field if non-nil, zero value otherwise.

### GetRequiresCertificateOk

`func (o *PublicCreateConnectorRequest) GetRequiresCertificateOk() (*bool, bool)`

GetRequiresCertificateOk returns a tuple with the RequiresCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresCertificate

`func (o *PublicCreateConnectorRequest) SetRequiresCertificate(v bool)`

SetRequiresCertificate sets RequiresCertificate field to given value.

### HasRequiresCertificate

`func (o *PublicCreateConnectorRequest) HasRequiresCertificate() bool`

HasRequiresCertificate returns a boolean if a field has been set.

### GetConnectorProtocol

`func (o *PublicCreateConnectorRequest) GetConnectorProtocol() string`

GetConnectorProtocol returns the ConnectorProtocol field if non-nil, zero value otherwise.

### GetConnectorProtocolOk

`func (o *PublicCreateConnectorRequest) GetConnectorProtocolOk() (*string, bool)`

GetConnectorProtocolOk returns a tuple with the ConnectorProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorProtocol

`func (o *PublicCreateConnectorRequest) SetConnectorProtocol(v string)`

SetConnectorProtocol sets ConnectorProtocol field to given value.

### HasConnectorProtocol

`func (o *PublicCreateConnectorRequest) HasConnectorProtocol() bool`

HasConnectorProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
