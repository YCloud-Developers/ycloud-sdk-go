# VerificationFallback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supported** | Pointer to **bool** | Whether this fallback you requested is supported. If &#x60;false&#x60; is returned, it means that there are errors for this fallback, and this fallback will not be triggered. | [optional] 
**UnsupportedReason** | Pointer to **string** | The reason why the fallback is unsupported, e.g, &#x60;PARAM_INVALID&#x60;, &#x60;SMS_SIGNATURE_UNAVAILABLE&#x60;, &#x60;SENDER_ID_UNAVAILABLE&#x60;, or &#x60;MESSAGING_REGION_UNSUPPORTED&#x60;. | [optional] 
**UnsupportedDetail** | Pointer to **string** | The detail message why the fallback is unsupported. | [optional] 

## Methods

### NewVerificationFallback

`func NewVerificationFallback() *VerificationFallback`

NewVerificationFallback instantiates a new VerificationFallback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationFallbackWithDefaults

`func NewVerificationFallbackWithDefaults() *VerificationFallback`

NewVerificationFallbackWithDefaults instantiates a new VerificationFallback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupported

`func (o *VerificationFallback) GetSupported() bool`

GetSupported returns the Supported field if non-nil, zero value otherwise.

### GetSupportedOk

`func (o *VerificationFallback) GetSupportedOk() (*bool, bool)`

GetSupportedOk returns a tuple with the Supported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupported

`func (o *VerificationFallback) SetSupported(v bool)`

SetSupported sets Supported field to given value.

### HasSupported

`func (o *VerificationFallback) HasSupported() bool`

HasSupported returns a boolean if a field has been set.

### GetUnsupportedReason

`func (o *VerificationFallback) GetUnsupportedReason() string`

GetUnsupportedReason returns the UnsupportedReason field if non-nil, zero value otherwise.

### GetUnsupportedReasonOk

`func (o *VerificationFallback) GetUnsupportedReasonOk() (*string, bool)`

GetUnsupportedReasonOk returns a tuple with the UnsupportedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsupportedReason

`func (o *VerificationFallback) SetUnsupportedReason(v string)`

SetUnsupportedReason sets UnsupportedReason field to given value.

### HasUnsupportedReason

`func (o *VerificationFallback) HasUnsupportedReason() bool`

HasUnsupportedReason returns a boolean if a field has been set.

### GetUnsupportedDetail

`func (o *VerificationFallback) GetUnsupportedDetail() string`

GetUnsupportedDetail returns the UnsupportedDetail field if non-nil, zero value otherwise.

### GetUnsupportedDetailOk

`func (o *VerificationFallback) GetUnsupportedDetailOk() (*string, bool)`

GetUnsupportedDetailOk returns a tuple with the UnsupportedDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsupportedDetail

`func (o *VerificationFallback) SetUnsupportedDetail(v string)`

SetUnsupportedDetail sets UnsupportedDetail field to given value.

### HasUnsupportedDetail

`func (o *VerificationFallback) HasUnsupportedDetail() bool`

HasUnsupportedDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
