# MetaBusinessAgentBudgetInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetId** | Pointer to **string** | Required when replacing an existing Budget; omit for a new Budget. | [optional]
**UnitType** | **string** |  |
**TimeWindow** | **string** |  |
**MaxBudget** | **int64** |  |

## Methods

### NewMetaBusinessAgentBudgetInput

`func NewMetaBusinessAgentBudgetInput(unitType string, timeWindow string, maxBudget int64, ) *MetaBusinessAgentBudgetInput`

NewMetaBusinessAgentBudgetInput instantiates a new MetaBusinessAgentBudgetInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaBusinessAgentBudgetInputWithDefaults

`func NewMetaBusinessAgentBudgetInputWithDefaults() *MetaBusinessAgentBudgetInput`

NewMetaBusinessAgentBudgetInputWithDefaults instantiates a new MetaBusinessAgentBudgetInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetId

`func (o *MetaBusinessAgentBudgetInput) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *MetaBusinessAgentBudgetInput) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *MetaBusinessAgentBudgetInput) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.

### HasBudgetId

`func (o *MetaBusinessAgentBudgetInput) HasBudgetId() bool`

HasBudgetId returns a boolean if a field has been set.

### GetUnitType

`func (o *MetaBusinessAgentBudgetInput) GetUnitType() string`

GetUnitType returns the UnitType field if non-nil, zero value otherwise.

### GetUnitTypeOk

`func (o *MetaBusinessAgentBudgetInput) GetUnitTypeOk() (*string, bool)`

GetUnitTypeOk returns a tuple with the UnitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitType

`func (o *MetaBusinessAgentBudgetInput) SetUnitType(v string)`

SetUnitType sets UnitType field to given value.


### GetTimeWindow

`func (o *MetaBusinessAgentBudgetInput) GetTimeWindow() string`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *MetaBusinessAgentBudgetInput) GetTimeWindowOk() (*string, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *MetaBusinessAgentBudgetInput) SetTimeWindow(v string)`

SetTimeWindow sets TimeWindow field to given value.


### GetMaxBudget

`func (o *MetaBusinessAgentBudgetInput) GetMaxBudget() int64`

GetMaxBudget returns the MaxBudget field if non-nil, zero value otherwise.

### GetMaxBudgetOk

`func (o *MetaBusinessAgentBudgetInput) GetMaxBudgetOk() (*int64, bool)`

GetMaxBudgetOk returns a tuple with the MaxBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBudget

`func (o *MetaBusinessAgentBudgetInput) SetMaxBudget(v int64)`

SetMaxBudget sets MaxBudget field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
