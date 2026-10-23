# StatsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartCount** | Pointer to **int32** |  | [optional]
**SuccessCount** | Pointer to **int32** |  | [optional]
**ExceptionCount** | Pointer to **int32** |  | [optional]
**SuccessRate** | Pointer to **float32** |  | [optional]
**AvgLatencyS** | Pointer to **float32** |  | [optional]
**P95LatencyS** | Pointer to **float32** |  | [optional]
**P99LatencyS** | Pointer to **float32** |  | [optional]
**TimeWindowSeconds** | Pointer to **int32** |  | [optional]

## Methods

### NewStatsDTO

`func NewStatsDTO() *StatsDTO`

NewStatsDTO instantiates a new StatsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsDTOWithDefaults

`func NewStatsDTOWithDefaults() *StatsDTO`

NewStatsDTOWithDefaults instantiates a new StatsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartCount

`func (o *StatsDTO) GetStartCount() int32`

GetStartCount returns the StartCount field if non-nil, zero value otherwise.

### GetStartCountOk

`func (o *StatsDTO) GetStartCountOk() (*int32, bool)`

GetStartCountOk returns a tuple with the StartCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCount

`func (o *StatsDTO) SetStartCount(v int32)`

SetStartCount sets StartCount field to given value.

### HasStartCount

`func (o *StatsDTO) HasStartCount() bool`

HasStartCount returns a boolean if a field has been set.

### GetSuccessCount

`func (o *StatsDTO) GetSuccessCount() int32`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *StatsDTO) GetSuccessCountOk() (*int32, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *StatsDTO) SetSuccessCount(v int32)`

SetSuccessCount sets SuccessCount field to given value.

### HasSuccessCount

`func (o *StatsDTO) HasSuccessCount() bool`

HasSuccessCount returns a boolean if a field has been set.

### GetExceptionCount

`func (o *StatsDTO) GetExceptionCount() int32`

GetExceptionCount returns the ExceptionCount field if non-nil, zero value otherwise.

### GetExceptionCountOk

`func (o *StatsDTO) GetExceptionCountOk() (*int32, bool)`

GetExceptionCountOk returns a tuple with the ExceptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionCount

`func (o *StatsDTO) SetExceptionCount(v int32)`

SetExceptionCount sets ExceptionCount field to given value.

### HasExceptionCount

`func (o *StatsDTO) HasExceptionCount() bool`

HasExceptionCount returns a boolean if a field has been set.

### GetSuccessRate

`func (o *StatsDTO) GetSuccessRate() float32`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *StatsDTO) GetSuccessRateOk() (*float32, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *StatsDTO) SetSuccessRate(v float32)`

SetSuccessRate sets SuccessRate field to given value.

### HasSuccessRate

`func (o *StatsDTO) HasSuccessRate() bool`

HasSuccessRate returns a boolean if a field has been set.

### GetAvgLatencyS

`func (o *StatsDTO) GetAvgLatencyS() float32`

GetAvgLatencyS returns the AvgLatencyS field if non-nil, zero value otherwise.

### GetAvgLatencySOk

`func (o *StatsDTO) GetAvgLatencySOk() (*float32, bool)`

GetAvgLatencySOk returns a tuple with the AvgLatencyS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgLatencyS

`func (o *StatsDTO) SetAvgLatencyS(v float32)`

SetAvgLatencyS sets AvgLatencyS field to given value.

### HasAvgLatencyS

`func (o *StatsDTO) HasAvgLatencyS() bool`

HasAvgLatencyS returns a boolean if a field has been set.

### GetP95LatencyS

`func (o *StatsDTO) GetP95LatencyS() float32`

GetP95LatencyS returns the P95LatencyS field if non-nil, zero value otherwise.

### GetP95LatencySOk

`func (o *StatsDTO) GetP95LatencySOk() (*float32, bool)`

GetP95LatencySOk returns a tuple with the P95LatencyS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP95LatencyS

`func (o *StatsDTO) SetP95LatencyS(v float32)`

SetP95LatencyS sets P95LatencyS field to given value.

### HasP95LatencyS

`func (o *StatsDTO) HasP95LatencyS() bool`

HasP95LatencyS returns a boolean if a field has been set.

### GetP99LatencyS

`func (o *StatsDTO) GetP99LatencyS() float32`

GetP99LatencyS returns the P99LatencyS field if non-nil, zero value otherwise.

### GetP99LatencySOk

`func (o *StatsDTO) GetP99LatencySOk() (*float32, bool)`

GetP99LatencySOk returns a tuple with the P99LatencyS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP99LatencyS

`func (o *StatsDTO) SetP99LatencyS(v float32)`

SetP99LatencyS sets P99LatencyS field to given value.

### HasP99LatencyS

`func (o *StatsDTO) HasP99LatencyS() bool`

HasP99LatencyS returns a boolean if a field has been set.

### GetTimeWindowSeconds

`func (o *StatsDTO) GetTimeWindowSeconds() int32`

GetTimeWindowSeconds returns the TimeWindowSeconds field if non-nil, zero value otherwise.

### GetTimeWindowSecondsOk

`func (o *StatsDTO) GetTimeWindowSecondsOk() (*int32, bool)`

GetTimeWindowSecondsOk returns a tuple with the TimeWindowSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindowSeconds

`func (o *StatsDTO) SetTimeWindowSeconds(v int32)`

SetTimeWindowSeconds sets TimeWindowSeconds field to given value.

### HasTimeWindowSeconds

`func (o *StatsDTO) HasTimeWindowSeconds() bool`

HasTimeWindowSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
