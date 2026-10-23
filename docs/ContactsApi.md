# ContactsApi

All URIs are relative to *https://api.ycloud.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttributesList**](ContactsApi.md#AttributesList) | **Get** /contact/contacts/attributes | List contact attributes
[**Create**](ContactsApi.md#Create) | **Post** /contact/contacts | Create a contact
[**Delete**](ContactsApi.md#Delete) | **Delete** /contact/contacts/{id} | Delete a contact
[**List**](ContactsApi.md#List) | **Get** /contact/contacts | List contacts
[**NoteCreate**](ContactsApi.md#NoteCreate) | **Post** /contact/contacts/{contactIdentifier}/notes | Create a contact note
[**NoteDelete**](ContactsApi.md#NoteDelete) | **Delete** /contact/notes/{noteId} | Delete a contact note
[**NoteUpdate**](ContactsApi.md#NoteUpdate) | **Patch** /contact/notes/{noteId} | Update a contact note
[**NotesList**](ContactsApi.md#NotesList) | **Get** /contact/contacts/{contactIdentifier}/notes | List contact notes
[**Retrieve**](ContactsApi.md#Retrieve) | **Get** /contact/contacts/{id} | Retrieve a contact
[**Update**](ContactsApi.md#Update) | **Patch** /contact/contacts/{id} | Update a contact



## AttributesList

> []ContactAttribute AttributesList(ctx).Execute()

List contact attributes



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

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactsApi.AttributesList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsApi.AttributesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttributesList`: []ContactAttribute
    fmt.Fprintf(os.Stdout, "Response from `ContactsApi.AttributesList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAttributesListRequest struct via the builder pattern


### Return type

[**[]ContactAttribute**](ContactAttribute.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> Contact Create(ctx).ContactCreateRequest(contactCreateRequest).Execute()

Create a contact



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
    contactCreateRequest := *ycloud.NewContactCreateRequest("+16315551111") // ContactCreateRequest |

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactsApi.Create(context.Background()).ContactCreateRequest(contactCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Contact
    fmt.Fprintf(os.Stdout, "Response from `ContactsApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactCreateRequest** | [**ContactCreateRequest**](ContactCreateRequest.md) |  |

### Return type

[**Contact**](Contact.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Contact Delete(ctx, id).Execute()

Delete a contact



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
    id := "alice_01" // string | Identifier of the contact. Supports a contact ID, a phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format starting with `+` (for example, `+16315551111`), or a Meta username without the leading `@`. A Meta username must contain 3 to 35 letters, digits, periods, or underscores. If a numeric value can be interpreted as both a Meta username and a contact ID, it is resolved as a Meta username first. If no contact has that Meta username, the value is resolved as a contact ID.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactsApi.Delete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: Contact
    fmt.Fprintf(os.Stdout, "Response from `ContactsApi.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the contact. Supports a contact ID, a phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format starting with &#x60;+&#x60; (for example, &#x60;+16315551111&#x60;), or a Meta username without the leading &#x60;@&#x60;. A Meta username must contain 3 to 35 letters, digits, periods, or underscores. If a numeric value can be interpreted as both a Meta username and a contact ID, it is resolved as a Meta username first. If no contact has that Meta username, the value is resolved as a contact ID. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Contact**](Contact.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ContactPage List(ctx).Page(page).Limit(limit).IncludeTotal(includeTotal).PageAfter(pageAfter).FilterTags(filterTags).FilterCountryCode(filterCountryCode).FilterPhoneNumber(filterPhoneNumber).FilterEmail(filterEmail).Execute()

List contacts



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
    page := int32(56) // int32 | Page number of the results to be returned, 1-based. (optional) (default to 1)
    limit := int32(56) // int32 | A limit on the number of results to be returned, or number of results per page, between 1 and 100, defaults to 10. (optional) (default to 10)
    includeTotal := true // bool | Return results inside an object that contains the total result count or not. (optional) (default to false)
    pageAfter := "0" // string | Contact ID cursor for forward pagination. Use `0` to start a new cursor traversal. For each subsequent page, pass the exact `cursor.after` value from the previous response. The value must be a non-negative decimal integer in the signed 64-bit range (`0` through `9223372036854775807`). It cannot be combined with `page`, `pageBefore`, `offset`, or `sort`. (optional)
    filterTags := "tag1,tag2" // string | Comma-separated list of tag names. If any tag does not exist, the request fails with a parameter error. (optional)
    filterCountryCode := "US" // string | Two-letter country abbreviation. See [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). (optional)
    filterPhoneNumber := "+16315551111" // string | Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. (optional)
    filterEmail := "support@example.com" // string | The contact's email address. (optional)

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactsApi.List(context.Background()).Page(page).Limit(limit).IncludeTotal(includeTotal).PageAfter(pageAfter).FilterTags(filterTags).FilterCountryCode(filterCountryCode).FilterPhoneNumber(filterPhoneNumber).FilterEmail(filterEmail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsApi.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ContactPage
    fmt.Fprintf(os.Stdout, "Response from `ContactsApi.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number of the results to be returned, 1-based. | [default to 1]
 **limit** | **int32** | A limit on the number of results to be returned, or number of results per page, between 1 and 100, defaults to 10. | [default to 10]
 **includeTotal** | **bool** | Return results inside an object that contains the total result count or not. | [default to false]
 **pageAfter** | **string** | Contact ID cursor for forward pagination. Use &#x60;0&#x60; to start a new cursor traversal. For each subsequent page, pass the exact &#x60;cursor.after&#x60; value from the previous response. The value must be a non-negative decimal integer in the signed 64-bit range (&#x60;0&#x60; through &#x60;9223372036854775807&#x60;). It cannot be combined with &#x60;page&#x60;, &#x60;pageBefore&#x60;, &#x60;offset&#x60;, or &#x60;sort&#x60;. |
 **filterTags** | **string** | Comma-separated list of tag names. If any tag does not exist, the request fails with a parameter error. |
 **filterCountryCode** | **string** | Two-letter country abbreviation. See [ISO 3166-1 alpha-2 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). |
 **filterPhoneNumber** | **string** | Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. |
 **filterEmail** | **string** | The contact&#39;s email address. |

### Return type

[**ContactPage**](ContactPage.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NoteCreate

> ContactNote NoteCreate(ctx, contactIdentifier).ContactNoteWriteRequest(contactNoteWriteRequest).Execute()

Create a contact note



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
    contactIdentifier := "alice_01" // string | Identifier of the contact that owns the note. Supports a contact ID, a username without the leading `@`, or a phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format starting with `+`. Usernames contain 3 to 35 letters, digits, periods, or underscores. A numeric value is resolved as a username first and then as a contact ID when no matching username exists.
    contactNoteWriteRequest := *ycloud.NewContactNoteWriteRequest("Customer now prefers afternoon follow-up.") // ContactNoteWriteRequest |

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactsApi.NoteCreate(context.Background(), contactIdentifier).ContactNoteWriteRequest(contactNoteWriteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsApi.NoteCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NoteCreate`: ContactNote
    fmt.Fprintf(os.Stdout, "Response from `ContactsApi.NoteCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactIdentifier** | **string** | Identifier of the contact that owns the note. Supports a contact ID, a username without the leading &#x60;@&#x60;, or a phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format starting with &#x60;+&#x60;. Usernames contain 3 to 35 letters, digits, periods, or underscores. A numeric value is resolved as a username first and then as a contact ID when no matching username exists. |

### Other Parameters

Other parameters are passed through a pointer to a apiNoteCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contactNoteWriteRequest** | [**ContactNoteWriteRequest**](ContactNoteWriteRequest.md) |  |

### Return type

[**ContactNote**](ContactNote.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NoteDelete

> ContactNote NoteDelete(ctx, noteId).Execute()

Delete a contact note



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
    noteId := "6a3de646e18f344f743aaa4d" // string | The 24-character ObjectId of the contact note.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactsApi.NoteDelete(context.Background(), noteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsApi.NoteDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NoteDelete`: ContactNote
    fmt.Fprintf(os.Stdout, "Response from `ContactsApi.NoteDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteId** | **string** | The 24-character ObjectId of the contact note. |

### Other Parameters

Other parameters are passed through a pointer to a apiNoteDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContactNote**](ContactNote.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NoteUpdate

> ContactNote NoteUpdate(ctx, noteId).ContactNoteWriteRequest(contactNoteWriteRequest).Execute()

Update a contact note



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
    noteId := "6a3de646e18f344f743aaa4d" // string | The 24-character ObjectId of the contact note.
    contactNoteWriteRequest := *ycloud.NewContactNoteWriteRequest("Customer now prefers afternoon follow-up.") // ContactNoteWriteRequest |

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactsApi.NoteUpdate(context.Background(), noteId).ContactNoteWriteRequest(contactNoteWriteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsApi.NoteUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NoteUpdate`: ContactNote
    fmt.Fprintf(os.Stdout, "Response from `ContactsApi.NoteUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteId** | **string** | The 24-character ObjectId of the contact note. |

### Other Parameters

Other parameters are passed through a pointer to a apiNoteUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contactNoteWriteRequest** | [**ContactNoteWriteRequest**](ContactNoteWriteRequest.md) |  |

### Return type

[**ContactNote**](ContactNote.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotesList

> []ContactNote NotesList(ctx, contactIdentifier).Execute()

List contact notes



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
    contactIdentifier := "alice_01" // string | Identifier of the contact that owns the note. Supports a contact ID, a username without the leading `@`, or a phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format starting with `+`. Usernames contain 3 to 35 letters, digits, periods, or underscores. A numeric value is resolved as a username first and then as a contact ID when no matching username exists.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactsApi.NotesList(context.Background(), contactIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsApi.NotesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NotesList`: []ContactNote
    fmt.Fprintf(os.Stdout, "Response from `ContactsApi.NotesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactIdentifier** | **string** | Identifier of the contact that owns the note. Supports a contact ID, a username without the leading &#x60;@&#x60;, or a phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format starting with &#x60;+&#x60;. Usernames contain 3 to 35 letters, digits, periods, or underscores. A numeric value is resolved as a username first and then as a contact ID when no matching username exists. |

### Other Parameters

Other parameters are passed through a pointer to a apiNotesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ContactNote**](ContactNote.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Retrieve

> Contact Retrieve(ctx, id).Execute()

Retrieve a contact



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
    id := "alice_01" // string | Identifier of the contact. Supports a contact ID, a phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format starting with `+` (for example, `+16315551111`), or a Meta username without the leading `@`. A Meta username must contain 3 to 35 letters, digits, periods, or underscores. If a numeric value can be interpreted as both a Meta username and a contact ID, it is resolved as a Meta username first. If no contact has that Meta username, the value is resolved as a contact ID.

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactsApi.Retrieve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsApi.Retrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Retrieve`: Contact
    fmt.Fprintf(os.Stdout, "Response from `ContactsApi.Retrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the contact. Supports a contact ID, a phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format starting with &#x60;+&#x60; (for example, &#x60;+16315551111&#x60;), or a Meta username without the leading &#x60;@&#x60;. A Meta username must contain 3 to 35 letters, digits, periods, or underscores. If a numeric value can be interpreted as both a Meta username and a contact ID, it is resolved as a Meta username first. If no contact has that Meta username, the value is resolved as a contact ID. |

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Contact**](Contact.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Contact Update(ctx, id).ContactUpdateRequest(contactUpdateRequest).Execute()

Update a contact



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
    id := "alice_01" // string | Identifier of the contact. Supports a contact ID, a phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format starting with `+` (for example, `+16315551111`), or a Meta username without the leading `@`. A Meta username must contain 3 to 35 letters, digits, periods, or underscores. If a numeric value can be interpreted as both a Meta username and a contact ID, it is resolved as a Meta username first. If no contact has that Meta username, the value is resolved as a contact ID.
    contactUpdateRequest := *ycloud.NewContactUpdateRequest() // ContactUpdateRequest |  (optional)

    configuration := ycloud.NewConfiguration()
    apiClient := ycloud.NewAPIClient(configuration)
    resp, r, err := apiClient.ContactsApi.Update(context.Background(), id).ContactUpdateRequest(contactUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContactsApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Contact
    fmt.Fprintf(os.Stdout, "Response from `ContactsApi.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the contact. Supports a contact ID, a phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format starting with &#x60;+&#x60; (for example, &#x60;+16315551111&#x60;), or a Meta username without the leading &#x60;@&#x60;. A Meta username must contain 3 to 35 letters, digits, periods, or underscores. If a numeric value can be interpreted as both a Meta username and a contact ID, it is resolved as a Meta username first. If no contact has that Meta username, the value is resolved as a contact ID. |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contactUpdateRequest** | [**ContactUpdateRequest**](ContactUpdateRequest.md) |  |

### Return type

[**Contact**](Contact.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
