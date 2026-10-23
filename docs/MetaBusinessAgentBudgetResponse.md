# MetaBusinessAgentBudgetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budgets** | [**[]MetaBusinessAgentBudget**](MetaBusinessAgentBudget.md) | The complete Budget collection. Each &#x60;(unit_type, time_window)&#x60; pair is unique. An empty array means unlimited usage. |

## Methods

### NewMetaBusinessAgentBudgetResponse

`func NewMetaBusinessAgentBudgetResponse(budgets []MetaBusinessAgentBudget, ) *MetaBusinessAgentBudgetResponse`

NewMetaBusinessAgentBudgetResponse instantiates a new MetaBusinessAgentBudgetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaBusinessAgentBudgetResponseWithDefaults

`func NewMetaBusinessAgentBudgetResponseWithDefaults() *MetaBusinessAgentBudgetResponse`

NewMetaBusinessAgentBudgetResponseWithDefaults instantiates a new MetaBusinessAgentBudgetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgets

`func (o *MetaBusinessAgentBudgetResponse) GetBudgets() []MetaBusinessAgentBudget`

GetBudgets returns the Budgets field if non-nil, zero value otherwise.

### GetBudgetsOk

`func (o *MetaBusinessAgentBudgetResponse) GetBudgetsOk() (*[]MetaBusinessAgentBudget, bool)`

GetBudgetsOk returns a tuple with the Budgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgets

`func (o *MetaBusinessAgentBudgetResponse) SetBudgets(v []MetaBusinessAgentBudget)`

SetBudgets sets Budgets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
