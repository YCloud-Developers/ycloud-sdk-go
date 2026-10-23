# MetaBusinessAgentApi

All URIs are relative to *https://api.ycloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BudgetReplace**](MetaBusinessAgentApi.md#BudgetReplace) | **Put** /metaBusinessAgents/{entityId}/budgets | Replace Meta Business Agent budgets
[**BudgetRetrieve**](MetaBusinessAgentApi.md#BudgetRetrieve) | **Get** /metaBusinessAgents/{entityId}/budgets | Retrieve Meta Business Agent budgets
[**BusinessAgentCreateConnector**](MetaBusinessAgentApi.md#BusinessAgentCreateConnector) | **Post** /metaBusinessAgents/{entityId}/connectors | Create Connector
[**BusinessAgentCreateConnectorTool**](MetaBusinessAgentApi.md#BusinessAgentCreateConnectorTool) | **Post** /metaBusinessAgents/{entityId}/connectors/{connectorId}/tools | Create Connector Tool
[**BusinessAgentDeleteConnector**](MetaBusinessAgentApi.md#BusinessAgentDeleteConnector) | **Delete** /metaBusinessAgents/{entityId}/connectors/{connectorId} | Delete Connector
[**BusinessAgentDeleteConnectorTool**](MetaBusinessAgentApi.md#BusinessAgentDeleteConnectorTool) | **Delete** /metaBusinessAgents/{entityId}/connectors/{connectorId}/tools/{toolId} | Delete Connector Tool
[**BusinessAgentGetConnector**](MetaBusinessAgentApi.md#BusinessAgentGetConnector) | **Get** /metaBusinessAgents/{entityId}/connectors/{connectorId} | Get Connector
[**BusinessAgentGetConnectorTool**](MetaBusinessAgentApi.md#BusinessAgentGetConnectorTool) | **Get** /metaBusinessAgents/{entityId}/connectors/{connectorId}/tools/{toolId} | Get Connector Tool
[**BusinessAgentListConnectorLogs**](MetaBusinessAgentApi.md#BusinessAgentListConnectorLogs) | **Get** /metaBusinessAgents/{entityId}/connectors/{connectorId}/logs | List Connector Logs
[**BusinessAgentListConnectorTools**](MetaBusinessAgentApi.md#BusinessAgentListConnectorTools) | **Get** /metaBusinessAgents/{entityId}/connectors/{connectorId}/tools | List Connector Tools
[**BusinessAgentListConnectors**](MetaBusinessAgentApi.md#BusinessAgentListConnectors) | **Get** /metaBusinessAgents/{entityId}/connectors | List Connectors
[**BusinessAgentRefreshConnectorMcpTools**](MetaBusinessAgentApi.md#BusinessAgentRefreshConnectorMcpTools) | **Post** /metaBusinessAgents/{entityId}/connectors/{connectorId}/refreshMCPTools | Refresh MCP Connector Tools
[**BusinessAgentRunConnectorTool**](MetaBusinessAgentApi.md#BusinessAgentRunConnectorTool) | **Post** /metaBusinessAgents/{entityId}/connectors/{connectorId}/tools/{toolId}/runs | Run Connector Tool
[**BusinessAgentUpdateConnector**](MetaBusinessAgentApi.md#BusinessAgentUpdateConnector) | **Patch** /metaBusinessAgents/{entityId}/connectors/{connectorId} | Update Connector
[**BusinessAgentUpdateConnectorTool**](MetaBusinessAgentApi.md#BusinessAgentUpdateConnectorTool) | **Patch** /metaBusinessAgents/{entityId}/connectors/{connectorId}/tools/{toolId} | Update Connector Tool
[**BusinessAgentUpsertConnectorApiKey**](MetaBusinessAgentApi.md#BusinessAgentUpsertConnectorApiKey) | **Put** /metaBusinessAgents/{entityId}/connectors/{connectorId}/credentials/apiKey | Upsert Connector Api Key
[**BusinessAgentUpsertConnectorCertificate**](MetaBusinessAgentApi.md#BusinessAgentUpsertConnectorCertificate) | **Put** /metaBusinessAgents/{entityId}/connectors/{connectorId}/credentials/certificate | Upsert Connector Certificate
[**BusinessAgentUpsertConnectorOAuth**](MetaBusinessAgentApi.md#BusinessAgentUpsertConnectorOAuth) | **Put** /metaBusinessAgents/{entityId}/connectors/{connectorId}/credentials/oauth | Upsert Connector O Auth
[**ConversationTurnsRetrieve**](MetaBusinessAgentApi.md#ConversationTurnsRetrieve) | **Get** /metaBusinessAgents/{entityId}/conversationTurns | Retrieve Meta Business Agent conversation turns
[**SettingsReplace**](MetaBusinessAgentApi.md#SettingsReplace) | **Put** /metaBusinessAgents/{entityId}/settings | Replace Meta Business Agent settings
[**SettingsRetrieve**](MetaBusinessAgentApi.md#SettingsRetrieve) | **Get** /metaBusinessAgents/{entityId}/settings | Retrieve Meta Business Agent settings



## BudgetReplace

> MetaBusinessAgentBudgetResponse BudgetReplace(ctx, entityId).MetaBusinessAgentBudgetReplaceRequest(metaBusinessAgentBudgetReplaceRequest).Execute()

Replace Meta Business Agent budgets



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    entityId := "1249005978294267" // string | Endpoint-specific Meta Business Agent entity ID. Budget and Conversation Turns operations use the tenant-owned WhatsApp Business Phone Number ID.
    metaBusinessAgentBudgetReplaceRequest := *ycloud.NewMetaBusinessAgentBudgetReplaceRequest([]ycloud.MetaBusinessAgentBudgetInput{*ycloud.NewMetaBusinessAgentBudgetInput("UnitType_example", "TimeWindow_example", int64(123))}) // MetaBusinessAgentBudgetReplaceRequest |

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaBusinessAgentApi.BudgetReplace(context.Background(), entityId).MetaBusinessAgentBudgetReplaceRequest(metaBusinessAgentBudgetReplaceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaBusinessAgentApi.BudgetReplace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BudgetReplace`: MetaBusinessAgentBudgetResponse
    fmt.Fprintf(os.Stdout, "Response from `MetaBusinessAgentApi.BudgetReplace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | Endpoint-specific Meta Business Agent entity ID. Budget and Conversation Turns operations use the tenant-owned WhatsApp Business Phone Number ID. |

### Other Parameters

Other parameters are passed through a pointer to a apiBudgetReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metaBusinessAgentBudgetReplaceRequest** | [**MetaBusinessAgentBudgetReplaceRequest**](MetaBusinessAgentBudgetReplaceRequest.md) |  |

### Return type

[**MetaBusinessAgentBudgetResponse**](MetaBusinessAgentBudgetResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BudgetRetrieve

> MetaBusinessAgentBudgetResponse BudgetRetrieve(ctx, entityId).Execute()

Retrieve Meta Business Agent budgets



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    entityId := "1249005978294267" // string | Endpoint-specific Meta Business Agent entity ID. Budget and Conversation Turns operations use the tenant-owned WhatsApp Business Phone Number ID.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaBusinessAgentApi.BudgetRetrieve(context.Background(), entityId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaBusinessAgentApi.BudgetRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BudgetRetrieve`: MetaBusinessAgentBudgetResponse
    fmt.Fprintf(os.Stdout, "Response from `MetaBusinessAgentApi.BudgetRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | Endpoint-specific Meta Business Agent entity ID. Budget and Conversation Turns operations use the tenant-owned WhatsApp Business Phone Number ID. |

### Other Parameters

Other parameters are passed through a pointer to a apiBudgetRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetaBusinessAgentBudgetResponse**](MetaBusinessAgentBudgetResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BusinessAgentCreateConnector

> ConnectorDTO BusinessAgentCreateConnector(ctx, entityId).PublicCreateConnectorRequest(publicCreateConnectorRequest).Execute()

Create Connector



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    entityId := "entityId_example" // string | WhatsApp Business Phone Number ID that owns the Meta Business Agent.
    publicCreateConnectorRequest := *ycloud.NewPublicCreateConnectorRequest("Name_example", "Description_example", "BaseUrl_example", "AuthType_example") // PublicCreateConnectorRequest |

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaBusinessAgentApi.BusinessAgentCreateConnector(context.Background(), entityId).PublicCreateConnectorRequest(publicCreateConnectorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaBusinessAgentApi.BusinessAgentCreateConnector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BusinessAgentCreateConnector`: ConnectorDTO
    fmt.Fprintf(os.Stdout, "Response from `MetaBusinessAgentApi.BusinessAgentCreateConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | WhatsApp Business Phone Number ID that owns the Meta Business Agent. |

### Other Parameters

Other parameters are passed through a pointer to a apiBusinessAgentCreateConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicCreateConnectorRequest** | [**PublicCreateConnectorRequest**](PublicCreateConnectorRequest.md) |  |

### Return type

[**ConnectorDTO**](ConnectorDTO.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BusinessAgentCreateConnectorTool

> ConnectorToolDTO BusinessAgentCreateConnectorTool(ctx, entityId, connectorId).PublicCreateConnectorToolRequest(publicCreateConnectorToolRequest).Execute()

Create Connector Tool



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    entityId := "entityId_example" // string | WhatsApp Business Phone Number ID that owns the Meta Business Agent.
    connectorId := "connectorId_example" // string | Connector resource ID.
    publicCreateConnectorToolRequest := *ycloud.NewPublicCreateConnectorToolRequest("Name_example", "Description_example", *ycloud.NewConnectorToolRequestDefinitionDTO("Method_example", "Path_example"), false) // PublicCreateConnectorToolRequest |

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaBusinessAgentApi.BusinessAgentCreateConnectorTool(context.Background(), entityId, connectorId).PublicCreateConnectorToolRequest(publicCreateConnectorToolRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaBusinessAgentApi.BusinessAgentCreateConnectorTool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BusinessAgentCreateConnectorTool`: ConnectorToolDTO
    fmt.Fprintf(os.Stdout, "Response from `MetaBusinessAgentApi.BusinessAgentCreateConnectorTool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | WhatsApp Business Phone Number ID that owns the Meta Business Agent. |
**connectorId** | **string** | Connector resource ID. |

### Other Parameters

Other parameters are passed through a pointer to a apiBusinessAgentCreateConnectorToolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **publicCreateConnectorToolRequest** | [**PublicCreateConnectorToolRequest**](PublicCreateConnectorToolRequest.md) |  |

### Return type

[**ConnectorToolDTO**](ConnectorToolDTO.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BusinessAgentDeleteConnector

> BusinessAgentDeleteConnector(ctx, entityId, connectorId).Execute()

Delete Connector



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    entityId := "entityId_example" // string | WhatsApp Business Phone Number ID that owns the Meta Business Agent.
    connectorId := "connectorId_example" // string | Connector resource ID.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaBusinessAgentApi.BusinessAgentDeleteConnector(context.Background(), entityId, connectorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaBusinessAgentApi.BusinessAgentDeleteConnector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | WhatsApp Business Phone Number ID that owns the Meta Business Agent. |
**connectorId** | **string** | Connector resource ID. |

### Other Parameters

Other parameters are passed through a pointer to a apiBusinessAgentDeleteConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BusinessAgentDeleteConnectorTool

> BusinessAgentDeleteConnectorTool(ctx, entityId, connectorId, toolId).Execute()

Delete Connector Tool



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    entityId := "entityId_example" // string | WhatsApp Business Phone Number ID that owns the Meta Business Agent.
    connectorId := "connectorId_example" // string | Connector resource ID.
    toolId := "toolId_example" // string | Connector tool resource ID.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaBusinessAgentApi.BusinessAgentDeleteConnectorTool(context.Background(), entityId, connectorId, toolId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaBusinessAgentApi.BusinessAgentDeleteConnectorTool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | WhatsApp Business Phone Number ID that owns the Meta Business Agent. |
**connectorId** | **string** | Connector resource ID. |
**toolId** | **string** | Connector tool resource ID. |

### Other Parameters

Other parameters are passed through a pointer to a apiBusinessAgentDeleteConnectorToolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BusinessAgentGetConnector

> ConnectorDTO BusinessAgentGetConnector(ctx, entityId, connectorId).Execute()

Get Connector



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    entityId := "entityId_example" // string | WhatsApp Business Phone Number ID that owns the Meta Business Agent.
    connectorId := "connectorId_example" // string | Connector resource ID.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaBusinessAgentApi.BusinessAgentGetConnector(context.Background(), entityId, connectorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaBusinessAgentApi.BusinessAgentGetConnector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BusinessAgentGetConnector`: ConnectorDTO
    fmt.Fprintf(os.Stdout, "Response from `MetaBusinessAgentApi.BusinessAgentGetConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | WhatsApp Business Phone Number ID that owns the Meta Business Agent. |
**connectorId** | **string** | Connector resource ID. |

### Other Parameters

Other parameters are passed through a pointer to a apiBusinessAgentGetConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConnectorDTO**](ConnectorDTO.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BusinessAgentGetConnectorTool

> ConnectorToolDTO BusinessAgentGetConnectorTool(ctx, entityId, connectorId, toolId).Execute()

Get Connector Tool



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    entityId := "entityId_example" // string | WhatsApp Business Phone Number ID that owns the Meta Business Agent.
    connectorId := "connectorId_example" // string | Connector resource ID.
    toolId := "toolId_example" // string | Connector tool resource ID.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaBusinessAgentApi.BusinessAgentGetConnectorTool(context.Background(), entityId, connectorId, toolId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaBusinessAgentApi.BusinessAgentGetConnectorTool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BusinessAgentGetConnectorTool`: ConnectorToolDTO
    fmt.Fprintf(os.Stdout, "Response from `MetaBusinessAgentApi.BusinessAgentGetConnectorTool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | WhatsApp Business Phone Number ID that owns the Meta Business Agent. |
**connectorId** | **string** | Connector resource ID. |
**toolId** | **string** | Connector tool resource ID. |

### Other Parameters

Other parameters are passed through a pointer to a apiBusinessAgentGetConnectorToolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ConnectorToolDTO**](ConnectorToolDTO.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BusinessAgentListConnectorLogs

> ConnectorLogStatsDTO BusinessAgentListConnectorLogs(ctx, entityId, connectorId).StartTime(startTime).EndTime(endTime).Limit(limit).ToolId(toolId).IncludeStats(includeStats).SummaryOnly(summaryOnly).TopN(topN).Execute()

List Connector Logs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    entityId := "entityId_example" // string | WhatsApp Business Phone Number ID that owns the Meta Business Agent.
    connectorId := "connectorId_example" // string | Connector resource ID.
    startTime := int64(789) // int64 |  (optional)
    endTime := int64(789) // int64 |  (optional)
    limit := int32(56) // int32 |  (optional)
    toolId := "toolId_example" // string |  (optional)
    includeStats := true // bool |  (optional)
    summaryOnly := true // bool |  (optional)
    topN := int32(56) // int32 |  (optional)

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaBusinessAgentApi.BusinessAgentListConnectorLogs(context.Background(), entityId, connectorId).StartTime(startTime).EndTime(endTime).Limit(limit).ToolId(toolId).IncludeStats(includeStats).SummaryOnly(summaryOnly).TopN(topN).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaBusinessAgentApi.BusinessAgentListConnectorLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BusinessAgentListConnectorLogs`: ConnectorLogStatsDTO
    fmt.Fprintf(os.Stdout, "Response from `MetaBusinessAgentApi.BusinessAgentListConnectorLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | WhatsApp Business Phone Number ID that owns the Meta Business Agent. |
**connectorId** | **string** | Connector resource ID. |

### Other Parameters

Other parameters are passed through a pointer to a apiBusinessAgentListConnectorLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startTime** | **int64** |  |
 **endTime** | **int64** |  |
 **limit** | **int32** |  |
 **toolId** | **string** |  |
 **includeStats** | **bool** |  |
 **summaryOnly** | **bool** |  |
 **topN** | **int32** |  |

### Return type

[**ConnectorLogStatsDTO**](ConnectorLogStatsDTO.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BusinessAgentListConnectorTools

> []ConnectorToolDTO BusinessAgentListConnectorTools(ctx, entityId, connectorId).Execute()

List Connector Tools



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    entityId := "entityId_example" // string | WhatsApp Business Phone Number ID that owns the Meta Business Agent.
    connectorId := "connectorId_example" // string | Connector resource ID.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaBusinessAgentApi.BusinessAgentListConnectorTools(context.Background(), entityId, connectorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaBusinessAgentApi.BusinessAgentListConnectorTools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BusinessAgentListConnectorTools`: []ConnectorToolDTO
    fmt.Fprintf(os.Stdout, "Response from `MetaBusinessAgentApi.BusinessAgentListConnectorTools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | WhatsApp Business Phone Number ID that owns the Meta Business Agent. |
**connectorId** | **string** | Connector resource ID. |

### Other Parameters

Other parameters are passed through a pointer to a apiBusinessAgentListConnectorToolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ConnectorToolDTO**](ConnectorToolDTO.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BusinessAgentListConnectors

> []ConnectorDTO BusinessAgentListConnectors(ctx, entityId).Execute()

List Connectors



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    entityId := "entityId_example" // string | WhatsApp Business Phone Number ID that owns the Meta Business Agent.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaBusinessAgentApi.BusinessAgentListConnectors(context.Background(), entityId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaBusinessAgentApi.BusinessAgentListConnectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BusinessAgentListConnectors`: []ConnectorDTO
    fmt.Fprintf(os.Stdout, "Response from `MetaBusinessAgentApi.BusinessAgentListConnectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | WhatsApp Business Phone Number ID that owns the Meta Business Agent. |

### Other Parameters

Other parameters are passed through a pointer to a apiBusinessAgentListConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ConnectorDTO**](ConnectorDTO.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BusinessAgentRefreshConnectorMcpTools

> ConnectorDTO BusinessAgentRefreshConnectorMcpTools(ctx, entityId, connectorId).Execute()

Refresh MCP Connector Tools



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    entityId := "entityId_example" // string | WhatsApp Business Phone Number ID that owns the Meta Business Agent.
    connectorId := "connectorId_example" // string | Meta Connector ID returned by the Connector API. This is not a YCloud local Connector UUID.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaBusinessAgentApi.BusinessAgentRefreshConnectorMcpTools(context.Background(), entityId, connectorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaBusinessAgentApi.BusinessAgentRefreshConnectorMcpTools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BusinessAgentRefreshConnectorMcpTools`: ConnectorDTO
    fmt.Fprintf(os.Stdout, "Response from `MetaBusinessAgentApi.BusinessAgentRefreshConnectorMcpTools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | WhatsApp Business Phone Number ID that owns the Meta Business Agent. |
**connectorId** | **string** | Meta Connector ID returned by the Connector API. This is not a YCloud local Connector UUID. |

### Other Parameters

Other parameters are passed through a pointer to a apiBusinessAgentRefreshConnectorMcpToolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConnectorDTO**](ConnectorDTO.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BusinessAgentRunConnectorTool

> ConnectorToolRunResultDTO BusinessAgentRunConnectorTool(ctx, entityId, connectorId, toolId).PublicRunConnectorToolRequest(publicRunConnectorToolRequest).Execute()

Run Connector Tool



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    entityId := "entityId_example" // string | WhatsApp Business Phone Number ID that owns the Meta Business Agent.
    connectorId := "connectorId_example" // string | Connector resource ID.
    toolId := "toolId_example" // string | Connector tool resource ID.
    publicRunConnectorToolRequest := *ycloud.NewPublicRunConnectorToolRequest() // PublicRunConnectorToolRequest |

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaBusinessAgentApi.BusinessAgentRunConnectorTool(context.Background(), entityId, connectorId, toolId).PublicRunConnectorToolRequest(publicRunConnectorToolRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaBusinessAgentApi.BusinessAgentRunConnectorTool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BusinessAgentRunConnectorTool`: ConnectorToolRunResultDTO
    fmt.Fprintf(os.Stdout, "Response from `MetaBusinessAgentApi.BusinessAgentRunConnectorTool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | WhatsApp Business Phone Number ID that owns the Meta Business Agent. |
**connectorId** | **string** | Connector resource ID. |
**toolId** | **string** | Connector tool resource ID. |

### Other Parameters

Other parameters are passed through a pointer to a apiBusinessAgentRunConnectorToolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **publicRunConnectorToolRequest** | [**PublicRunConnectorToolRequest**](PublicRunConnectorToolRequest.md) |  |

### Return type

[**ConnectorToolRunResultDTO**](ConnectorToolRunResultDTO.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BusinessAgentUpdateConnector

> ConnectorDTO BusinessAgentUpdateConnector(ctx, entityId, connectorId).PublicUpdateConnectorRequest(publicUpdateConnectorRequest).Execute()

Update Connector



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    entityId := "entityId_example" // string | WhatsApp Business Phone Number ID that owns the Meta Business Agent.
    connectorId := "connectorId_example" // string | Connector resource ID.
    publicUpdateConnectorRequest := *ycloud.NewPublicUpdateConnectorRequest("Name_example", "Description_example", "BaseUrl_example", "AuthType_example") // PublicUpdateConnectorRequest |

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaBusinessAgentApi.BusinessAgentUpdateConnector(context.Background(), entityId, connectorId).PublicUpdateConnectorRequest(publicUpdateConnectorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaBusinessAgentApi.BusinessAgentUpdateConnector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BusinessAgentUpdateConnector`: ConnectorDTO
    fmt.Fprintf(os.Stdout, "Response from `MetaBusinessAgentApi.BusinessAgentUpdateConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | WhatsApp Business Phone Number ID that owns the Meta Business Agent. |
**connectorId** | **string** | Connector resource ID. |

### Other Parameters

Other parameters are passed through a pointer to a apiBusinessAgentUpdateConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **publicUpdateConnectorRequest** | [**PublicUpdateConnectorRequest**](PublicUpdateConnectorRequest.md) |  |

### Return type

[**ConnectorDTO**](ConnectorDTO.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BusinessAgentUpdateConnectorTool

> ConnectorToolDTO BusinessAgentUpdateConnectorTool(ctx, entityId, connectorId, toolId).PublicUpdateConnectorToolRequest(publicUpdateConnectorToolRequest).Execute()

Update Connector Tool



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    entityId := "entityId_example" // string | WhatsApp Business Phone Number ID that owns the Meta Business Agent.
    connectorId := "connectorId_example" // string | Connector resource ID.
    toolId := "toolId_example" // string | Connector tool resource ID.
    publicUpdateConnectorToolRequest := *ycloud.NewPublicUpdateConnectorToolRequest("Name_example", "Description_example", *ycloud.NewConnectorToolRequestDefinitionDTO("Method_example", "Path_example"), false) // PublicUpdateConnectorToolRequest |

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaBusinessAgentApi.BusinessAgentUpdateConnectorTool(context.Background(), entityId, connectorId, toolId).PublicUpdateConnectorToolRequest(publicUpdateConnectorToolRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaBusinessAgentApi.BusinessAgentUpdateConnectorTool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BusinessAgentUpdateConnectorTool`: ConnectorToolDTO
    fmt.Fprintf(os.Stdout, "Response from `MetaBusinessAgentApi.BusinessAgentUpdateConnectorTool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | WhatsApp Business Phone Number ID that owns the Meta Business Agent. |
**connectorId** | **string** | Connector resource ID. |
**toolId** | **string** | Connector tool resource ID. |

### Other Parameters

Other parameters are passed through a pointer to a apiBusinessAgentUpdateConnectorToolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **publicUpdateConnectorToolRequest** | [**PublicUpdateConnectorToolRequest**](PublicUpdateConnectorToolRequest.md) |  |

### Return type

[**ConnectorToolDTO**](ConnectorToolDTO.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BusinessAgentUpsertConnectorApiKey

> ConnectorDTO BusinessAgentUpsertConnectorApiKey(ctx, entityId, connectorId).PublicUpsertConnectorApiKeyRequest(publicUpsertConnectorApiKeyRequest).Execute()

Upsert Connector Api Key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    entityId := "entityId_example" // string | WhatsApp Business Phone Number ID that owns the Meta Business Agent.
    connectorId := "connectorId_example" // string | Connector resource ID.
    publicUpsertConnectorApiKeyRequest := *ycloud.NewPublicUpsertConnectorApiKeyRequest(*ycloud.NewApiKeyDTO(interface{}(123), interface{}(123), interface{}(123))) // PublicUpsertConnectorApiKeyRequest |

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaBusinessAgentApi.BusinessAgentUpsertConnectorApiKey(context.Background(), entityId, connectorId).PublicUpsertConnectorApiKeyRequest(publicUpsertConnectorApiKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaBusinessAgentApi.BusinessAgentUpsertConnectorApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BusinessAgentUpsertConnectorApiKey`: ConnectorDTO
    fmt.Fprintf(os.Stdout, "Response from `MetaBusinessAgentApi.BusinessAgentUpsertConnectorApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | WhatsApp Business Phone Number ID that owns the Meta Business Agent. |
**connectorId** | **string** | Connector resource ID. |

### Other Parameters

Other parameters are passed through a pointer to a apiBusinessAgentUpsertConnectorApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **publicUpsertConnectorApiKeyRequest** | [**PublicUpsertConnectorApiKeyRequest**](PublicUpsertConnectorApiKeyRequest.md) |  |

### Return type

[**ConnectorDTO**](ConnectorDTO.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BusinessAgentUpsertConnectorCertificate

> ConnectorDTO BusinessAgentUpsertConnectorCertificate(ctx, entityId, connectorId).PublicUpsertConnectorCertificateRequest(publicUpsertConnectorCertificateRequest).Execute()

Upsert Connector Certificate



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    entityId := "entityId_example" // string | WhatsApp Business Phone Number ID that owns the Meta Business Agent.
    connectorId := "connectorId_example" // string | Connector resource ID.
    publicUpsertConnectorCertificateRequest := *ycloud.NewPublicUpsertConnectorCertificateRequest("ClientCertificate_example", "ClientKey_example") // PublicUpsertConnectorCertificateRequest |

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaBusinessAgentApi.BusinessAgentUpsertConnectorCertificate(context.Background(), entityId, connectorId).PublicUpsertConnectorCertificateRequest(publicUpsertConnectorCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaBusinessAgentApi.BusinessAgentUpsertConnectorCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BusinessAgentUpsertConnectorCertificate`: ConnectorDTO
    fmt.Fprintf(os.Stdout, "Response from `MetaBusinessAgentApi.BusinessAgentUpsertConnectorCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | WhatsApp Business Phone Number ID that owns the Meta Business Agent. |
**connectorId** | **string** | Connector resource ID. |

### Other Parameters

Other parameters are passed through a pointer to a apiBusinessAgentUpsertConnectorCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **publicUpsertConnectorCertificateRequest** | [**PublicUpsertConnectorCertificateRequest**](PublicUpsertConnectorCertificateRequest.md) |  |

### Return type

[**ConnectorDTO**](ConnectorDTO.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BusinessAgentUpsertConnectorOAuth

> ConnectorDTO BusinessAgentUpsertConnectorOAuth(ctx, entityId, connectorId).PublicUpsertConnectorOAuthRequest(publicUpsertConnectorOAuthRequest).Execute()

Upsert Connector O Auth



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    entityId := "entityId_example" // string | WhatsApp Business Phone Number ID that owns the Meta Business Agent.
    connectorId := "connectorId_example" // string | Connector resource ID.
    publicUpsertConnectorOAuthRequest := *ycloud.NewPublicUpsertConnectorOAuthRequest(*ycloud.NewOAuth2ClientCredentialsDTO("TokenUrl_example", []string{"ScopesToRequest_example"}, "ClientId_example", "ClientSecret_example")) // PublicUpsertConnectorOAuthRequest |

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaBusinessAgentApi.BusinessAgentUpsertConnectorOAuth(context.Background(), entityId, connectorId).PublicUpsertConnectorOAuthRequest(publicUpsertConnectorOAuthRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaBusinessAgentApi.BusinessAgentUpsertConnectorOAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BusinessAgentUpsertConnectorOAuth`: ConnectorDTO
    fmt.Fprintf(os.Stdout, "Response from `MetaBusinessAgentApi.BusinessAgentUpsertConnectorOAuth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | WhatsApp Business Phone Number ID that owns the Meta Business Agent. |
**connectorId** | **string** | Connector resource ID. |

### Other Parameters

Other parameters are passed through a pointer to a apiBusinessAgentUpsertConnectorOAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **publicUpsertConnectorOAuthRequest** | [**PublicUpsertConnectorOAuthRequest**](PublicUpsertConnectorOAuthRequest.md) |  |

### Return type

[**ConnectorDTO**](ConnectorDTO.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConversationTurnsRetrieve

> MetaBusinessAgentConversationTurnPage ConversationTurnsRetrieve(ctx, entityId).UserPhoneNumber(userPhoneNumber).StartTimestampMs(startTimestampMs).EndTimestampMs(endTimestampMs).Before(before).After(after).Limit(limit).Execute()

Retrieve Meta Business Agent conversation turns



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    entityId := "1249005978294267" // string | Endpoint-specific Meta Business Agent entity ID. Budget and Conversation Turns operations use the tenant-owned WhatsApp Business Phone Number ID.
    userPhoneNumber := "14155550123" // string | WhatsApp user phone number including the country code, with digits only and no leading `+`, spaces, or separators.
    startTimestampMs := int64(789) // int64 | Inclusive start time as Unix epoch milliseconds. (optional)
    endTimestampMs := int64(789) // int64 | Inclusive end time as Unix epoch milliseconds. Must be greater than or equal to `start_timestamp_ms` when both are supplied. (optional)
    before := "before_example" // string | Opaque backward-pagination cursor. Do not combine with `after`. (optional)
    after := "after_example" // string | Opaque forward-pagination cursor. Do not combine with `before`. (optional)
    limit := int32(56) // int32 | Maximum number of conversation turns to return. (optional)

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaBusinessAgentApi.ConversationTurnsRetrieve(context.Background(), entityId).UserPhoneNumber(userPhoneNumber).StartTimestampMs(startTimestampMs).EndTimestampMs(endTimestampMs).Before(before).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaBusinessAgentApi.ConversationTurnsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConversationTurnsRetrieve`: MetaBusinessAgentConversationTurnPage
    fmt.Fprintf(os.Stdout, "Response from `MetaBusinessAgentApi.ConversationTurnsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | Endpoint-specific Meta Business Agent entity ID. Budget and Conversation Turns operations use the tenant-owned WhatsApp Business Phone Number ID. |

### Other Parameters

Other parameters are passed through a pointer to a apiConversationTurnsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userPhoneNumber** | **string** | WhatsApp user phone number including the country code, with digits only and no leading &#x60;+&#x60;, spaces, or separators. |
 **startTimestampMs** | **int64** | Inclusive start time as Unix epoch milliseconds. |
 **endTimestampMs** | **int64** | Inclusive end time as Unix epoch milliseconds. Must be greater than or equal to &#x60;start_timestamp_ms&#x60; when both are supplied. |
 **before** | **string** | Opaque backward-pagination cursor. Do not combine with &#x60;after&#x60;. |
 **after** | **string** | Opaque forward-pagination cursor. Do not combine with &#x60;before&#x60;. |
 **limit** | **int32** | Maximum number of conversation turns to return. |

### Return type

[**MetaBusinessAgentConversationTurnPage**](MetaBusinessAgentConversationTurnPage.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SettingsReplace

> MetaBusinessAgentSettings SettingsReplace(ctx, entityId).MetaBusinessAgentSettingsReplaceRequest(metaBusinessAgentSettingsReplaceRequest).Execute()

Replace Meta Business Agent settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    entityId := "1249005978294267" // string | Endpoint-specific Meta Business Agent entity ID. Budget and Conversation Turns operations use the tenant-owned WhatsApp Business Phone Number ID.
    metaBusinessAgentSettingsReplaceRequest := *ycloud.NewMetaBusinessAgentSettingsReplaceRequest() // MetaBusinessAgentSettingsReplaceRequest |

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaBusinessAgentApi.SettingsReplace(context.Background(), entityId).MetaBusinessAgentSettingsReplaceRequest(metaBusinessAgentSettingsReplaceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaBusinessAgentApi.SettingsReplace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SettingsReplace`: MetaBusinessAgentSettings
    fmt.Fprintf(os.Stdout, "Response from `MetaBusinessAgentApi.SettingsReplace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | Endpoint-specific Meta Business Agent entity ID. Budget and Conversation Turns operations use the tenant-owned WhatsApp Business Phone Number ID. |

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metaBusinessAgentSettingsReplaceRequest** | [**MetaBusinessAgentSettingsReplaceRequest**](MetaBusinessAgentSettingsReplaceRequest.md) |  |

### Return type

[**MetaBusinessAgentSettings**](MetaBusinessAgentSettings.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SettingsRetrieve

> []MetaBusinessAgentSettings SettingsRetrieve(ctx, entityId).AgentId(agentId).Execute()

Retrieve Meta Business Agent settings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    ycloud "github.com/ycloud-developers/ycloud-sdk-go"
)

func main() {
    entityId := "1249005978294267" // string | Endpoint-specific Meta Business Agent entity ID. Budget and Conversation Turns operations use the tenant-owned WhatsApp Business Phone Number ID.
    agentId := "agentId_example" // string | Optional Meta Agent ID used to select one settings record. (optional)

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaBusinessAgentApi.SettingsRetrieve(context.Background(), entityId).AgentId(agentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaBusinessAgentApi.SettingsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SettingsRetrieve`: []MetaBusinessAgentSettings
    fmt.Fprintf(os.Stdout, "Response from `MetaBusinessAgentApi.SettingsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | Endpoint-specific Meta Business Agent entity ID. Budget and Conversation Turns operations use the tenant-owned WhatsApp Business Phone Number ID. |

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agentId** | **string** | Optional Meta Agent ID used to select one settings record. |

### Return type

[**[]MetaBusinessAgentSettings**](MetaBusinessAgentSettings.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
