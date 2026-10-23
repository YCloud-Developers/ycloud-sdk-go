# WhatsappAutomaticCreativeOptimizationRetrieveResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WabaId** | Pointer to **string** | WhatsApp Business Account ID returned by Meta. | [optional]
**CreativeOptimizationFeatures** | Pointer to **map[string]string** | Current feature enrollment string values returned by Meta. REST returns fields and original string values from &#x60;creative_features_spec[0]&#x60; without feature-key filtering, status-value validation, or case normalization. | [optional]

## Methods

### NewWhatsappAutomaticCreativeOptimizationRetrieveResponse

`func NewWhatsappAutomaticCreativeOptimizationRetrieveResponse() *WhatsappAutomaticCreativeOptimizationRetrieveResponse`

NewWhatsappAutomaticCreativeOptimizationRetrieveResponse instantiates a new WhatsappAutomaticCreativeOptimizationRetrieveResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappAutomaticCreativeOptimizationRetrieveResponseWithDefaults

`func NewWhatsappAutomaticCreativeOptimizationRetrieveResponseWithDefaults() *WhatsappAutomaticCreativeOptimizationRetrieveResponse`

NewWhatsappAutomaticCreativeOptimizationRetrieveResponseWithDefaults instantiates a new WhatsappAutomaticCreativeOptimizationRetrieveResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWabaId

`func (o *WhatsappAutomaticCreativeOptimizationRetrieveResponse) GetWabaId() string`

GetWabaId returns the WabaId field if non-nil, zero value otherwise.

### GetWabaIdOk

`func (o *WhatsappAutomaticCreativeOptimizationRetrieveResponse) GetWabaIdOk() (*string, bool)`

GetWabaIdOk returns a tuple with the WabaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWabaId

`func (o *WhatsappAutomaticCreativeOptimizationRetrieveResponse) SetWabaId(v string)`

SetWabaId sets WabaId field to given value.

### HasWabaId

`func (o *WhatsappAutomaticCreativeOptimizationRetrieveResponse) HasWabaId() bool`

HasWabaId returns a boolean if a field has been set.

### GetCreativeOptimizationFeatures

`func (o *WhatsappAutomaticCreativeOptimizationRetrieveResponse) GetCreativeOptimizationFeatures() map[string]string`

GetCreativeOptimizationFeatures returns the CreativeOptimizationFeatures field if non-nil, zero value otherwise.

### GetCreativeOptimizationFeaturesOk

`func (o *WhatsappAutomaticCreativeOptimizationRetrieveResponse) GetCreativeOptimizationFeaturesOk() (*map[string]string, bool)`

GetCreativeOptimizationFeaturesOk returns a tuple with the CreativeOptimizationFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreativeOptimizationFeatures

`func (o *WhatsappAutomaticCreativeOptimizationRetrieveResponse) SetCreativeOptimizationFeatures(v map[string]string)`

SetCreativeOptimizationFeatures sets CreativeOptimizationFeatures field to given value.

### HasCreativeOptimizationFeatures

`func (o *WhatsappAutomaticCreativeOptimizationRetrieveResponse) HasCreativeOptimizationFeatures() bool`

HasCreativeOptimizationFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
