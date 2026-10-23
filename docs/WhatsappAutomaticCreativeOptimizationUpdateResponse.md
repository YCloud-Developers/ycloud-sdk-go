# WhatsappAutomaticCreativeOptimizationUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WabaId** | Pointer to **string** | WhatsApp Business Account ID returned by Meta. | [optional]
**CreativeOptimizationFeatures** | Pointer to [**map[string]WhatsappAutomaticCreativeOptimizationEnrollStatus**](WhatsappAutomaticCreativeOptimizationEnrollStatus.md) | The submitted feature enrollment values for PATCH responses, or current feature enrollment values for GET responses. | [optional]

## Methods

### NewWhatsappAutomaticCreativeOptimizationUpdateResponse

`func NewWhatsappAutomaticCreativeOptimizationUpdateResponse() *WhatsappAutomaticCreativeOptimizationUpdateResponse`

NewWhatsappAutomaticCreativeOptimizationUpdateResponse instantiates a new WhatsappAutomaticCreativeOptimizationUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappAutomaticCreativeOptimizationUpdateResponseWithDefaults

`func NewWhatsappAutomaticCreativeOptimizationUpdateResponseWithDefaults() *WhatsappAutomaticCreativeOptimizationUpdateResponse`

NewWhatsappAutomaticCreativeOptimizationUpdateResponseWithDefaults instantiates a new WhatsappAutomaticCreativeOptimizationUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWabaId

`func (o *WhatsappAutomaticCreativeOptimizationUpdateResponse) GetWabaId() string`

GetWabaId returns the WabaId field if non-nil, zero value otherwise.

### GetWabaIdOk

`func (o *WhatsappAutomaticCreativeOptimizationUpdateResponse) GetWabaIdOk() (*string, bool)`

GetWabaIdOk returns a tuple with the WabaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWabaId

`func (o *WhatsappAutomaticCreativeOptimizationUpdateResponse) SetWabaId(v string)`

SetWabaId sets WabaId field to given value.

### HasWabaId

`func (o *WhatsappAutomaticCreativeOptimizationUpdateResponse) HasWabaId() bool`

HasWabaId returns a boolean if a field has been set.

### GetCreativeOptimizationFeatures

`func (o *WhatsappAutomaticCreativeOptimizationUpdateResponse) GetCreativeOptimizationFeatures() map[string]WhatsappAutomaticCreativeOptimizationEnrollStatus`

GetCreativeOptimizationFeatures returns the CreativeOptimizationFeatures field if non-nil, zero value otherwise.

### GetCreativeOptimizationFeaturesOk

`func (o *WhatsappAutomaticCreativeOptimizationUpdateResponse) GetCreativeOptimizationFeaturesOk() (*map[string]WhatsappAutomaticCreativeOptimizationEnrollStatus, bool)`

GetCreativeOptimizationFeaturesOk returns a tuple with the CreativeOptimizationFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreativeOptimizationFeatures

`func (o *WhatsappAutomaticCreativeOptimizationUpdateResponse) SetCreativeOptimizationFeatures(v map[string]WhatsappAutomaticCreativeOptimizationEnrollStatus)`

SetCreativeOptimizationFeatures sets CreativeOptimizationFeatures field to given value.

### HasCreativeOptimizationFeatures

`func (o *WhatsappAutomaticCreativeOptimizationUpdateResponse) HasCreativeOptimizationFeatures() bool`

HasCreativeOptimizationFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
