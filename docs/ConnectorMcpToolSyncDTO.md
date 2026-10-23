# ConnectorMcpToolSyncDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Tool discovery state. ERROR indicates discovery failure even when the connector operation succeeds. |
**LastAttemptedAt** | Pointer to **NullableInt64** | Last discovery attempt, in Unix seconds. | [optional]
**LastSuccessfulAt** | Pointer to **NullableInt64** | Last successful discovery, in Unix seconds. | [optional]
**Fingerprint** | Pointer to **NullableString** |  | [optional]
**ToolCount** | Pointer to **NullableInt32** |  | [optional]

## Methods

### NewConnectorMcpToolSyncDTO

`func NewConnectorMcpToolSyncDTO(status string, ) *ConnectorMcpToolSyncDTO`

NewConnectorMcpToolSyncDTO instantiates a new ConnectorMcpToolSyncDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorMcpToolSyncDTOWithDefaults

`func NewConnectorMcpToolSyncDTOWithDefaults() *ConnectorMcpToolSyncDTO`

NewConnectorMcpToolSyncDTOWithDefaults instantiates a new ConnectorMcpToolSyncDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ConnectorMcpToolSyncDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectorMcpToolSyncDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectorMcpToolSyncDTO) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLastAttemptedAt

`func (o *ConnectorMcpToolSyncDTO) GetLastAttemptedAt() int64`

GetLastAttemptedAt returns the LastAttemptedAt field if non-nil, zero value otherwise.

### GetLastAttemptedAtOk

`func (o *ConnectorMcpToolSyncDTO) GetLastAttemptedAtOk() (*int64, bool)`

GetLastAttemptedAtOk returns a tuple with the LastAttemptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAttemptedAt

`func (o *ConnectorMcpToolSyncDTO) SetLastAttemptedAt(v int64)`

SetLastAttemptedAt sets LastAttemptedAt field to given value.

### HasLastAttemptedAt

`func (o *ConnectorMcpToolSyncDTO) HasLastAttemptedAt() bool`

HasLastAttemptedAt returns a boolean if a field has been set.

### SetLastAttemptedAtNil

`func (o *ConnectorMcpToolSyncDTO) SetLastAttemptedAtNil(b bool)`

 SetLastAttemptedAtNil sets the value for LastAttemptedAt to be an explicit nil

### UnsetLastAttemptedAt
`func (o *ConnectorMcpToolSyncDTO) UnsetLastAttemptedAt()`

UnsetLastAttemptedAt ensures that no value is present for LastAttemptedAt, not even an explicit nil
### GetLastSuccessfulAt

`func (o *ConnectorMcpToolSyncDTO) GetLastSuccessfulAt() int64`

GetLastSuccessfulAt returns the LastSuccessfulAt field if non-nil, zero value otherwise.

### GetLastSuccessfulAtOk

`func (o *ConnectorMcpToolSyncDTO) GetLastSuccessfulAtOk() (*int64, bool)`

GetLastSuccessfulAtOk returns a tuple with the LastSuccessfulAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulAt

`func (o *ConnectorMcpToolSyncDTO) SetLastSuccessfulAt(v int64)`

SetLastSuccessfulAt sets LastSuccessfulAt field to given value.

### HasLastSuccessfulAt

`func (o *ConnectorMcpToolSyncDTO) HasLastSuccessfulAt() bool`

HasLastSuccessfulAt returns a boolean if a field has been set.

### SetLastSuccessfulAtNil

`func (o *ConnectorMcpToolSyncDTO) SetLastSuccessfulAtNil(b bool)`

 SetLastSuccessfulAtNil sets the value for LastSuccessfulAt to be an explicit nil

### UnsetLastSuccessfulAt
`func (o *ConnectorMcpToolSyncDTO) UnsetLastSuccessfulAt()`

UnsetLastSuccessfulAt ensures that no value is present for LastSuccessfulAt, not even an explicit nil
### GetFingerprint

`func (o *ConnectorMcpToolSyncDTO) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *ConnectorMcpToolSyncDTO) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *ConnectorMcpToolSyncDTO) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *ConnectorMcpToolSyncDTO) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *ConnectorMcpToolSyncDTO) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *ConnectorMcpToolSyncDTO) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetToolCount

`func (o *ConnectorMcpToolSyncDTO) GetToolCount() int32`

GetToolCount returns the ToolCount field if non-nil, zero value otherwise.

### GetToolCountOk

`func (o *ConnectorMcpToolSyncDTO) GetToolCountOk() (*int32, bool)`

GetToolCountOk returns a tuple with the ToolCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCount

`func (o *ConnectorMcpToolSyncDTO) SetToolCount(v int32)`

SetToolCount sets ToolCount field to given value.

### HasToolCount

`func (o *ConnectorMcpToolSyncDTO) HasToolCount() bool`

HasToolCount returns a boolean if a field has been set.

### SetToolCountNil

`func (o *ConnectorMcpToolSyncDTO) SetToolCountNil(b bool)`

 SetToolCountNil sets the value for ToolCount to be an explicit nil

### UnsetToolCount
`func (o *ConnectorMcpToolSyncDTO) UnsetToolCount()`

UnsetToolCount ensures that no value is present for ToolCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
