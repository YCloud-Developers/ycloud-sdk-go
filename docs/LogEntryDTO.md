# LogEntryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventTime** | Pointer to **string** |  | [optional]
**FailureCodeName** | Pointer to **string** |  | [optional]
**ErrorMessage** | Pointer to **string** |  | [optional]
**ToolName** | Pointer to **string** |  | [optional]
**Occurrences** | Pointer to **int32** |  | [optional]
**LastSeen** | Pointer to **string** |  | [optional]

## Methods

### NewLogEntryDTO

`func NewLogEntryDTO() *LogEntryDTO`

NewLogEntryDTO instantiates a new LogEntryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogEntryDTOWithDefaults

`func NewLogEntryDTOWithDefaults() *LogEntryDTO`

NewLogEntryDTOWithDefaults instantiates a new LogEntryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventTime

`func (o *LogEntryDTO) GetEventTime() string`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *LogEntryDTO) GetEventTimeOk() (*string, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *LogEntryDTO) SetEventTime(v string)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *LogEntryDTO) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetFailureCodeName

`func (o *LogEntryDTO) GetFailureCodeName() string`

GetFailureCodeName returns the FailureCodeName field if non-nil, zero value otherwise.

### GetFailureCodeNameOk

`func (o *LogEntryDTO) GetFailureCodeNameOk() (*string, bool)`

GetFailureCodeNameOk returns a tuple with the FailureCodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCodeName

`func (o *LogEntryDTO) SetFailureCodeName(v string)`

SetFailureCodeName sets FailureCodeName field to given value.

### HasFailureCodeName

`func (o *LogEntryDTO) HasFailureCodeName() bool`

HasFailureCodeName returns a boolean if a field has been set.

### GetErrorMessage

`func (o *LogEntryDTO) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *LogEntryDTO) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *LogEntryDTO) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *LogEntryDTO) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetToolName

`func (o *LogEntryDTO) GetToolName() string`

GetToolName returns the ToolName field if non-nil, zero value otherwise.

### GetToolNameOk

`func (o *LogEntryDTO) GetToolNameOk() (*string, bool)`

GetToolNameOk returns a tuple with the ToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolName

`func (o *LogEntryDTO) SetToolName(v string)`

SetToolName sets ToolName field to given value.

### HasToolName

`func (o *LogEntryDTO) HasToolName() bool`

HasToolName returns a boolean if a field has been set.

### GetOccurrences

`func (o *LogEntryDTO) GetOccurrences() int32`

GetOccurrences returns the Occurrences field if non-nil, zero value otherwise.

### GetOccurrencesOk

`func (o *LogEntryDTO) GetOccurrencesOk() (*int32, bool)`

GetOccurrencesOk returns a tuple with the Occurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrences

`func (o *LogEntryDTO) SetOccurrences(v int32)`

SetOccurrences sets Occurrences field to given value.

### HasOccurrences

`func (o *LogEntryDTO) HasOccurrences() bool`

HasOccurrences returns a boolean if a field has been set.

### GetLastSeen

`func (o *LogEntryDTO) GetLastSeen() string`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *LogEntryDTO) GetLastSeenOk() (*string, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *LogEntryDTO) SetLastSeen(v string)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *LogEntryDTO) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
