# ContactAttributeChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OldValue** | Pointer to **interface{}** | The previous value of the attribute before the change. Can be a string, number, array, or boolean depending on the attribute type. This field is not included when the value is null. | [optional] 
**NewValue** | Pointer to **interface{}** | The new value of the attribute after the change. Can be a string, number, array, or boolean depending on the attribute type. This field is not included when the value is null. | [optional] 
**Extra** | Pointer to [**[]AttributeChangeAction**](AttributeChangeAction.md) | An array of change actions that describe what operations were performed on this attribute. | [optional] 

## Methods

### NewContactAttributeChange

`func NewContactAttributeChange() *ContactAttributeChange`

NewContactAttributeChange instantiates a new ContactAttributeChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactAttributeChangeWithDefaults

`func NewContactAttributeChangeWithDefaults() *ContactAttributeChange`

NewContactAttributeChangeWithDefaults instantiates a new ContactAttributeChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOldValue

`func (o *ContactAttributeChange) GetOldValue() interface{}`

GetOldValue returns the OldValue field if non-nil, zero value otherwise.

### GetOldValueOk

`func (o *ContactAttributeChange) GetOldValueOk() (*interface{}, bool)`

GetOldValueOk returns a tuple with the OldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValue

`func (o *ContactAttributeChange) SetOldValue(v interface{})`

SetOldValue sets OldValue field to given value.

### HasOldValue

`func (o *ContactAttributeChange) HasOldValue() bool`

HasOldValue returns a boolean if a field has been set.

### SetOldValueNil

`func (o *ContactAttributeChange) SetOldValueNil(b bool)`

 SetOldValueNil sets the value for OldValue to be an explicit nil

### UnsetOldValue
`func (o *ContactAttributeChange) UnsetOldValue()`

UnsetOldValue ensures that no value is present for OldValue, not even an explicit nil
### GetNewValue

`func (o *ContactAttributeChange) GetNewValue() interface{}`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *ContactAttributeChange) GetNewValueOk() (*interface{}, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *ContactAttributeChange) SetNewValue(v interface{})`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *ContactAttributeChange) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### SetNewValueNil

`func (o *ContactAttributeChange) SetNewValueNil(b bool)`

 SetNewValueNil sets the value for NewValue to be an explicit nil

### UnsetNewValue
`func (o *ContactAttributeChange) UnsetNewValue()`

UnsetNewValue ensures that no value is present for NewValue, not even an explicit nil
### GetExtra

`func (o *ContactAttributeChange) GetExtra() []AttributeChangeAction`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *ContactAttributeChange) GetExtraOk() (*[]AttributeChangeAction, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *ContactAttributeChange) SetExtra(v []AttributeChangeAction)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *ContactAttributeChange) HasExtra() bool`

HasExtra returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


