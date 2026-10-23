# PublicUpsertConnectorCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientCertificate** | **string** |  |
**ClientKey** | **string** |  |
**CaCertificate** | Pointer to **string** |  | [optional]

## Methods

### NewPublicUpsertConnectorCertificateRequest

`func NewPublicUpsertConnectorCertificateRequest(clientCertificate string, clientKey string, ) *PublicUpsertConnectorCertificateRequest`

NewPublicUpsertConnectorCertificateRequest instantiates a new PublicUpsertConnectorCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicUpsertConnectorCertificateRequestWithDefaults

`func NewPublicUpsertConnectorCertificateRequestWithDefaults() *PublicUpsertConnectorCertificateRequest`

NewPublicUpsertConnectorCertificateRequestWithDefaults instantiates a new PublicUpsertConnectorCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientCertificate

`func (o *PublicUpsertConnectorCertificateRequest) GetClientCertificate() string`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *PublicUpsertConnectorCertificateRequest) GetClientCertificateOk() (*string, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *PublicUpsertConnectorCertificateRequest) SetClientCertificate(v string)`

SetClientCertificate sets ClientCertificate field to given value.


### GetClientKey

`func (o *PublicUpsertConnectorCertificateRequest) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *PublicUpsertConnectorCertificateRequest) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *PublicUpsertConnectorCertificateRequest) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.


### GetCaCertificate

`func (o *PublicUpsertConnectorCertificateRequest) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *PublicUpsertConnectorCertificateRequest) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *PublicUpsertConnectorCertificateRequest) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *PublicUpsertConnectorCertificateRequest) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
