# CallingCaptureSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordingEnabled** | **bool** | Whether recording capture is enabled. |
**TranscriptionEnabled** | **bool** | Whether transcription capture is enabled. |
**Purpose** | Pointer to **string** | Customer-defined capture purpose passed to the Calling provider configuration. Required when either capture switch is enabled. | [optional]
**AnnouncementLanguage** | Pointer to **string** | Meta-supported announcement language. Required when either capture switch is enabled. | [optional]

## Methods

### NewCallingCaptureSettings

`func NewCallingCaptureSettings(recordingEnabled bool, transcriptionEnabled bool, ) *CallingCaptureSettings`

NewCallingCaptureSettings instantiates a new CallingCaptureSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallingCaptureSettingsWithDefaults

`func NewCallingCaptureSettingsWithDefaults() *CallingCaptureSettings`

NewCallingCaptureSettingsWithDefaults instantiates a new CallingCaptureSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordingEnabled

`func (o *CallingCaptureSettings) GetRecordingEnabled() bool`

GetRecordingEnabled returns the RecordingEnabled field if non-nil, zero value otherwise.

### GetRecordingEnabledOk

`func (o *CallingCaptureSettings) GetRecordingEnabledOk() (*bool, bool)`

GetRecordingEnabledOk returns a tuple with the RecordingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingEnabled

`func (o *CallingCaptureSettings) SetRecordingEnabled(v bool)`

SetRecordingEnabled sets RecordingEnabled field to given value.


### GetTranscriptionEnabled

`func (o *CallingCaptureSettings) GetTranscriptionEnabled() bool`

GetTranscriptionEnabled returns the TranscriptionEnabled field if non-nil, zero value otherwise.

### GetTranscriptionEnabledOk

`func (o *CallingCaptureSettings) GetTranscriptionEnabledOk() (*bool, bool)`

GetTranscriptionEnabledOk returns a tuple with the TranscriptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionEnabled

`func (o *CallingCaptureSettings) SetTranscriptionEnabled(v bool)`

SetTranscriptionEnabled sets TranscriptionEnabled field to given value.


### GetPurpose

`func (o *CallingCaptureSettings) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *CallingCaptureSettings) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *CallingCaptureSettings) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *CallingCaptureSettings) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetAnnouncementLanguage

`func (o *CallingCaptureSettings) GetAnnouncementLanguage() string`

GetAnnouncementLanguage returns the AnnouncementLanguage field if non-nil, zero value otherwise.

### GetAnnouncementLanguageOk

`func (o *CallingCaptureSettings) GetAnnouncementLanguageOk() (*string, bool)`

GetAnnouncementLanguageOk returns a tuple with the AnnouncementLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementLanguage

`func (o *CallingCaptureSettings) SetAnnouncementLanguage(v string)`

SetAnnouncementLanguage sets AnnouncementLanguage field to given value.

### HasAnnouncementLanguage

`func (o *CallingCaptureSettings) HasAnnouncementLanguage() bool`

HasAnnouncementLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
