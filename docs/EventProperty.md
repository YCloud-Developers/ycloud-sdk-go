# EventProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** | The event type for which properties are configured. This field accepts any valid event type that supports property configuration. | 
**Properties** | **[]string** | A list of property names that should be included in the webhook payload for the specified event type. The available properties depend on the specific event type configured. | 

## Methods

### NewEventProperty

`func NewEventProperty(event string, properties []string, ) *EventProperty`

NewEventProperty instantiates a new EventProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventPropertyWithDefaults

`func NewEventPropertyWithDefaults() *EventProperty`

NewEventPropertyWithDefaults instantiates a new EventProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *EventProperty) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventProperty) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventProperty) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetProperties

`func (o *EventProperty) GetProperties() []string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *EventProperty) GetPropertiesOk() (*[]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *EventProperty) SetProperties(v []string)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


