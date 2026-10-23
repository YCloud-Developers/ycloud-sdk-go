# ConnectorDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional]
**Name** | Pointer to **string** |  | [optional]
**Description** | Pointer to **string** |  | [optional]
**BaseUrl** | Pointer to **string** |  | [optional]
**AuthType** | Pointer to **string** |  | [optional]
**AuthConfig** | Pointer to [**ConnectorAuthConfigDTO**](ConnectorAuthConfigDTO.md) |  | [optional]
**MtlsConfig** | Pointer to [**MtlsConfigDTO**](MtlsConfigDTO.md) |  | [optional]
**ConnectionStatus** | Pointer to [**ConnectionStatusDTO**](ConnectionStatusDTO.md) |  | [optional]
**UserAuthInjectionConfig** | Pointer to [**UserAuthInjectionConfigDTO**](UserAuthInjectionConfigDTO.md) |  | [optional]
**ConnectorProtocol** | Pointer to **NullableString** | Connector protocol. Omit on creation to use Meta’s HTTP default. The protocol cannot change after creation; omit on update to preserve it. | [optional]
**McpToolSync** | Pointer to [**NullableConnectorMcpToolSyncDTO**](ConnectorMcpToolSyncDTO.md) |  | [optional]

## Methods

### NewConnectorDTO

`func NewConnectorDTO() *ConnectorDTO`

NewConnectorDTO instantiates a new ConnectorDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorDTOWithDefaults

`func NewConnectorDTOWithDefaults() *ConnectorDTO`

NewConnectorDTOWithDefaults instantiates a new ConnectorDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectorDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectorDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectorDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectorDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ConnectorDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectorDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ConnectorDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectorDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectorDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConnectorDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBaseUrl

`func (o *ConnectorDTO) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *ConnectorDTO) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *ConnectorDTO) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *ConnectorDTO) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetAuthType

`func (o *ConnectorDTO) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *ConnectorDTO) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *ConnectorDTO) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *ConnectorDTO) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetAuthConfig

`func (o *ConnectorDTO) GetAuthConfig() ConnectorAuthConfigDTO`

GetAuthConfig returns the AuthConfig field if non-nil, zero value otherwise.

### GetAuthConfigOk

`func (o *ConnectorDTO) GetAuthConfigOk() (*ConnectorAuthConfigDTO, bool)`

GetAuthConfigOk returns a tuple with the AuthConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthConfig

`func (o *ConnectorDTO) SetAuthConfig(v ConnectorAuthConfigDTO)`

SetAuthConfig sets AuthConfig field to given value.

### HasAuthConfig

`func (o *ConnectorDTO) HasAuthConfig() bool`

HasAuthConfig returns a boolean if a field has been set.

### GetMtlsConfig

`func (o *ConnectorDTO) GetMtlsConfig() MtlsConfigDTO`

GetMtlsConfig returns the MtlsConfig field if non-nil, zero value otherwise.

### GetMtlsConfigOk

`func (o *ConnectorDTO) GetMtlsConfigOk() (*MtlsConfigDTO, bool)`

GetMtlsConfigOk returns a tuple with the MtlsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtlsConfig

`func (o *ConnectorDTO) SetMtlsConfig(v MtlsConfigDTO)`

SetMtlsConfig sets MtlsConfig field to given value.

### HasMtlsConfig

`func (o *ConnectorDTO) HasMtlsConfig() bool`

HasMtlsConfig returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *ConnectorDTO) GetConnectionStatus() ConnectionStatusDTO`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *ConnectorDTO) GetConnectionStatusOk() (*ConnectionStatusDTO, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *ConnectorDTO) SetConnectionStatus(v ConnectionStatusDTO)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *ConnectorDTO) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetUserAuthInjectionConfig

`func (o *ConnectorDTO) GetUserAuthInjectionConfig() UserAuthInjectionConfigDTO`

GetUserAuthInjectionConfig returns the UserAuthInjectionConfig field if non-nil, zero value otherwise.

### GetUserAuthInjectionConfigOk

`func (o *ConnectorDTO) GetUserAuthInjectionConfigOk() (*UserAuthInjectionConfigDTO, bool)`

GetUserAuthInjectionConfigOk returns a tuple with the UserAuthInjectionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAuthInjectionConfig

`func (o *ConnectorDTO) SetUserAuthInjectionConfig(v UserAuthInjectionConfigDTO)`

SetUserAuthInjectionConfig sets UserAuthInjectionConfig field to given value.

### HasUserAuthInjectionConfig

`func (o *ConnectorDTO) HasUserAuthInjectionConfig() bool`

HasUserAuthInjectionConfig returns a boolean if a field has been set.

### GetConnectorProtocol

`func (o *ConnectorDTO) GetConnectorProtocol() string`

GetConnectorProtocol returns the ConnectorProtocol field if non-nil, zero value otherwise.

### GetConnectorProtocolOk

`func (o *ConnectorDTO) GetConnectorProtocolOk() (*string, bool)`

GetConnectorProtocolOk returns a tuple with the ConnectorProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorProtocol

`func (o *ConnectorDTO) SetConnectorProtocol(v string)`

SetConnectorProtocol sets ConnectorProtocol field to given value.

### HasConnectorProtocol

`func (o *ConnectorDTO) HasConnectorProtocol() bool`

HasConnectorProtocol returns a boolean if a field has been set.

### SetConnectorProtocolNil

`func (o *ConnectorDTO) SetConnectorProtocolNil(b bool)`

 SetConnectorProtocolNil sets the value for ConnectorProtocol to be an explicit nil

### UnsetConnectorProtocol
`func (o *ConnectorDTO) UnsetConnectorProtocol()`

UnsetConnectorProtocol ensures that no value is present for ConnectorProtocol, not even an explicit nil
### GetMcpToolSync

`func (o *ConnectorDTO) GetMcpToolSync() ConnectorMcpToolSyncDTO`

GetMcpToolSync returns the McpToolSync field if non-nil, zero value otherwise.

### GetMcpToolSyncOk

`func (o *ConnectorDTO) GetMcpToolSyncOk() (*ConnectorMcpToolSyncDTO, bool)`

GetMcpToolSyncOk returns a tuple with the McpToolSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpToolSync

`func (o *ConnectorDTO) SetMcpToolSync(v ConnectorMcpToolSyncDTO)`

SetMcpToolSync sets McpToolSync field to given value.

### HasMcpToolSync

`func (o *ConnectorDTO) HasMcpToolSync() bool`

HasMcpToolSync returns a boolean if a field has been set.

### SetMcpToolSyncNil

`func (o *ConnectorDTO) SetMcpToolSyncNil(b bool)`

 SetMcpToolSyncNil sets the value for McpToolSync to be an explicit nil

### UnsetMcpToolSync
`func (o *ConnectorDTO) UnsetMcpToolSync()`

UnsetMcpToolSync ensures that no value is present for McpToolSync, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
