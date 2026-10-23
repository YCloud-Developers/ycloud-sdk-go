# CallingMediaUpdated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wacid** | **string** | WhatsApp call ID. |
**PhoneId** | **string** | WhatsApp business phone number ID used for authorized-asset delivery. |
**MediaAssetId** | **string** | YCloud media asset ID used with the call media download API. |
**Status** | **string** |  |
**Error** | Pointer to [**CallingMediaUpdatedError**](CallingMediaUpdatedError.md) |  | [optional]

## Methods

### NewCallingMediaUpdated

`func NewCallingMediaUpdated(wacid string, phoneId string, mediaAssetId string, status string, ) *CallingMediaUpdated`

NewCallingMediaUpdated instantiates a new CallingMediaUpdated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallingMediaUpdatedWithDefaults

`func NewCallingMediaUpdatedWithDefaults() *CallingMediaUpdated`

NewCallingMediaUpdatedWithDefaults instantiates a new CallingMediaUpdated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWacid

`func (o *CallingMediaUpdated) GetWacid() string`

GetWacid returns the Wacid field if non-nil, zero value otherwise.

### GetWacidOk

`func (o *CallingMediaUpdated) GetWacidOk() (*string, bool)`

GetWacidOk returns a tuple with the Wacid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWacid

`func (o *CallingMediaUpdated) SetWacid(v string)`

SetWacid sets Wacid field to given value.


### GetPhoneId

`func (o *CallingMediaUpdated) GetPhoneId() string`

GetPhoneId returns the PhoneId field if non-nil, zero value otherwise.

### GetPhoneIdOk

`func (o *CallingMediaUpdated) GetPhoneIdOk() (*string, bool)`

GetPhoneIdOk returns a tuple with the PhoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneId

`func (o *CallingMediaUpdated) SetPhoneId(v string)`

SetPhoneId sets PhoneId field to given value.


### GetMediaAssetId

`func (o *CallingMediaUpdated) GetMediaAssetId() string`

GetMediaAssetId returns the MediaAssetId field if non-nil, zero value otherwise.

### GetMediaAssetIdOk

`func (o *CallingMediaUpdated) GetMediaAssetIdOk() (*string, bool)`

GetMediaAssetIdOk returns a tuple with the MediaAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaAssetId

`func (o *CallingMediaUpdated) SetMediaAssetId(v string)`

SetMediaAssetId sets MediaAssetId field to given value.


### GetStatus

`func (o *CallingMediaUpdated) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CallingMediaUpdated) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CallingMediaUpdated) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetError

`func (o *CallingMediaUpdated) GetError() CallingMediaUpdatedError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CallingMediaUpdated) GetErrorOk() (*CallingMediaUpdatedError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CallingMediaUpdated) SetError(v CallingMediaUpdatedError)`

SetError sets Error field to given value.

### HasError

`func (o *CallingMediaUpdated) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
