# MtlsConfigDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasCertificate** | Pointer to **bool** |  | [optional]
**Fingerprint** | Pointer to **string** |  | [optional]
**ExpiresAt** | Pointer to **int64** |  | [optional]
**Subject** | Pointer to **string** |  | [optional]
**ClientCertificate** | Pointer to **string** |  | [optional]
**CaCertificate** | Pointer to **string** |  | [optional]

## Methods

### NewMtlsConfigDTO

`func NewMtlsConfigDTO() *MtlsConfigDTO`

NewMtlsConfigDTO instantiates a new MtlsConfigDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMtlsConfigDTOWithDefaults

`func NewMtlsConfigDTOWithDefaults() *MtlsConfigDTO`

NewMtlsConfigDTOWithDefaults instantiates a new MtlsConfigDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasCertificate

`func (o *MtlsConfigDTO) GetHasCertificate() bool`

GetHasCertificate returns the HasCertificate field if non-nil, zero value otherwise.

### GetHasCertificateOk

`func (o *MtlsConfigDTO) GetHasCertificateOk() (*bool, bool)`

GetHasCertificateOk returns a tuple with the HasCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCertificate

`func (o *MtlsConfigDTO) SetHasCertificate(v bool)`

SetHasCertificate sets HasCertificate field to given value.

### HasHasCertificate

`func (o *MtlsConfigDTO) HasHasCertificate() bool`

HasHasCertificate returns a boolean if a field has been set.

### GetFingerprint

`func (o *MtlsConfigDTO) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *MtlsConfigDTO) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *MtlsConfigDTO) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *MtlsConfigDTO) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetExpiresAt

`func (o *MtlsConfigDTO) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *MtlsConfigDTO) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *MtlsConfigDTO) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *MtlsConfigDTO) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetSubject

`func (o *MtlsConfigDTO) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MtlsConfigDTO) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MtlsConfigDTO) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *MtlsConfigDTO) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetClientCertificate

`func (o *MtlsConfigDTO) GetClientCertificate() string`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *MtlsConfigDTO) GetClientCertificateOk() (*string, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *MtlsConfigDTO) SetClientCertificate(v string)`

SetClientCertificate sets ClientCertificate field to given value.

### HasClientCertificate

`func (o *MtlsConfigDTO) HasClientCertificate() bool`

HasClientCertificate returns a boolean if a field has been set.

### GetCaCertificate

`func (o *MtlsConfigDTO) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *MtlsConfigDTO) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *MtlsConfigDTO) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *MtlsConfigDTO) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
