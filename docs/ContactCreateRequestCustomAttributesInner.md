# ContactCreateRequestCustomAttributesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the attribute that you&#39;ve previously defined. | [optional] 
**Value** | Pointer to **map[string]interface{}** | Value of the attribute. | [optional] 

## Methods

### NewContactCreateRequestCustomAttributesInner

`func NewContactCreateRequestCustomAttributesInner() *ContactCreateRequestCustomAttributesInner`

NewContactCreateRequestCustomAttributesInner instantiates a new ContactCreateRequestCustomAttributesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactCreateRequestCustomAttributesInnerWithDefaults

`func NewContactCreateRequestCustomAttributesInnerWithDefaults() *ContactCreateRequestCustomAttributesInner`

NewContactCreateRequestCustomAttributesInnerWithDefaults instantiates a new ContactCreateRequestCustomAttributesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContactCreateRequestCustomAttributesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContactCreateRequestCustomAttributesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContactCreateRequestCustomAttributesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContactCreateRequestCustomAttributesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *ContactCreateRequestCustomAttributesInner) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContactCreateRequestCustomAttributesInner) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContactCreateRequestCustomAttributesInner) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ContactCreateRequestCustomAttributesInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


