# PublicUpsertConnectorApiKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyConfig** | [**ApiKeyDTO**](ApiKeyDTO.md) |  |

## Methods

### NewPublicUpsertConnectorApiKeyRequest

`func NewPublicUpsertConnectorApiKeyRequest(apiKeyConfig ApiKeyDTO, ) *PublicUpsertConnectorApiKeyRequest`

NewPublicUpsertConnectorApiKeyRequest instantiates a new PublicUpsertConnectorApiKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicUpsertConnectorApiKeyRequestWithDefaults

`func NewPublicUpsertConnectorApiKeyRequestWithDefaults() *PublicUpsertConnectorApiKeyRequest`

NewPublicUpsertConnectorApiKeyRequestWithDefaults instantiates a new PublicUpsertConnectorApiKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyConfig

`func (o *PublicUpsertConnectorApiKeyRequest) GetApiKeyConfig() ApiKeyDTO`

GetApiKeyConfig returns the ApiKeyConfig field if non-nil, zero value otherwise.

### GetApiKeyConfigOk

`func (o *PublicUpsertConnectorApiKeyRequest) GetApiKeyConfigOk() (*ApiKeyDTO, bool)`

GetApiKeyConfigOk returns a tuple with the ApiKeyConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyConfig

`func (o *PublicUpsertConnectorApiKeyRequest) SetApiKeyConfig(v ApiKeyDTO)`

SetApiKeyConfig sets ApiKeyConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
