# ContactCustomAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the attribute that you&#39;ve previously defined. | [optional] 
**Value** | Pointer to **map[string]interface{}** | Value of the attribute. Its data type depends on the format of the attribute you defined: For Text, the &#x60;value&#x60; is a string with a maximum length of 250.  For Array, the &#x60;value&#x60; is an array of strings with a maximum length of 250. For Number, the &#x60;value&#x60; is a signed decimal number. For Boolean, the &#x60;value&#x60; is either &#x60;true&#x60; or &#x60;false&#x60;. For Time, the &#x60;value&#x60; is a Unix timestamp in milliseconds. For Long Text, the &#x60;value&#x60; is a string with a maximum length of 5000. | [optional] 

## Methods

### NewContactCustomAttribute

`func NewContactCustomAttribute() *ContactCustomAttribute`

NewContactCustomAttribute instantiates a new ContactCustomAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactCustomAttributeWithDefaults

`func NewContactCustomAttributeWithDefaults() *ContactCustomAttribute`

NewContactCustomAttributeWithDefaults instantiates a new ContactCustomAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContactCustomAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContactCustomAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContactCustomAttribute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContactCustomAttribute) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *ContactCustomAttribute) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContactCustomAttribute) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContactCustomAttribute) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ContactCustomAttribute) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
