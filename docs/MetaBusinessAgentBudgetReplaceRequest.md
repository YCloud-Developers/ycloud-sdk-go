# MetaBusinessAgentBudgetReplaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budgets** | [**[]MetaBusinessAgentBudgetInput**](MetaBusinessAgentBudgetInput.md) | The complete replacement collection. Each &#x60;(unit_type, time_window)&#x60; pair must be unique. An empty array removes all limits. |

## Methods

### NewMetaBusinessAgentBudgetReplaceRequest

`func NewMetaBusinessAgentBudgetReplaceRequest(budgets []MetaBusinessAgentBudgetInput, ) *MetaBusinessAgentBudgetReplaceRequest`

NewMetaBusinessAgentBudgetReplaceRequest instantiates a new MetaBusinessAgentBudgetReplaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaBusinessAgentBudgetReplaceRequestWithDefaults

`func NewMetaBusinessAgentBudgetReplaceRequestWithDefaults() *MetaBusinessAgentBudgetReplaceRequest`

NewMetaBusinessAgentBudgetReplaceRequestWithDefaults instantiates a new MetaBusinessAgentBudgetReplaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgets

`func (o *MetaBusinessAgentBudgetReplaceRequest) GetBudgets() []MetaBusinessAgentBudgetInput`

GetBudgets returns the Budgets field if non-nil, zero value otherwise.

### GetBudgetsOk

`func (o *MetaBusinessAgentBudgetReplaceRequest) GetBudgetsOk() (*[]MetaBusinessAgentBudgetInput, bool)`

GetBudgetsOk returns a tuple with the Budgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgets

`func (o *MetaBusinessAgentBudgetReplaceRequest) SetBudgets(v []MetaBusinessAgentBudgetInput)`

SetBudgets sets Budgets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
