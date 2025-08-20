# ContactAttributesChanged

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the contact whose attributes were changed. | 
**UpdateTime** | **time.Time** | The time at which the contact attributes were updated, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | 
**ChangedAttributes** | [**map[string]ContactAttributeChange**](ContactAttributeChange.md) | An object containing the changed attributes. Each key represents the name of the changed attribute, and the value contains the old value, new value, and change actions. | 

## Methods

### NewContactAttributesChanged

`func NewContactAttributesChanged(id string, updateTime time.Time, changedAttributes map[string]ContactAttributeChange, ) *ContactAttributesChanged`

NewContactAttributesChanged instantiates a new ContactAttributesChanged object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactAttributesChangedWithDefaults

`func NewContactAttributesChangedWithDefaults() *ContactAttributesChanged`

NewContactAttributesChangedWithDefaults instantiates a new ContactAttributesChanged object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactAttributesChanged) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactAttributesChanged) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactAttributesChanged) SetId(v string)`

SetId sets Id field to given value.


### GetUpdateTime

`func (o *ContactAttributesChanged) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *ContactAttributesChanged) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *ContactAttributesChanged) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.


### GetChangedAttributes

`func (o *ContactAttributesChanged) GetChangedAttributes() map[string]ContactAttributeChange`

GetChangedAttributes returns the ChangedAttributes field if non-nil, zero value otherwise.

### GetChangedAttributesOk

`func (o *ContactAttributesChanged) GetChangedAttributesOk() (*map[string]ContactAttributeChange, bool)`

GetChangedAttributesOk returns a tuple with the ChangedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedAttributes

`func (o *ContactAttributesChanged) SetChangedAttributes(v map[string]ContactAttributeChange)`

SetChangedAttributes sets ChangedAttributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


