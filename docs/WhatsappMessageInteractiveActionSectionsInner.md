# WhatsappMessageInteractiveActionSectionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title of the section. | [optional] 
**Rows** | Pointer to [**[]WhatsappMessageInteractiveActionSectionsInnerRowsInner**](WhatsappMessageInteractiveActionSectionsInnerRowsInner.md) | Contains a list of rows. You can have a total of 10 rows across your sections. Each row must have a title (Maximum length: 24 characters) and an ID (Maximum length: 200 characters). You can add a description (Maximum length: 72 characters), but it is optional. | [optional] 

## Methods

### NewWhatsappMessageInteractiveActionSectionsInner

`func NewWhatsappMessageInteractiveActionSectionsInner() *WhatsappMessageInteractiveActionSectionsInner`

NewWhatsappMessageInteractiveActionSectionsInner instantiates a new WhatsappMessageInteractiveActionSectionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappMessageInteractiveActionSectionsInnerWithDefaults

`func NewWhatsappMessageInteractiveActionSectionsInnerWithDefaults() *WhatsappMessageInteractiveActionSectionsInner`

NewWhatsappMessageInteractiveActionSectionsInnerWithDefaults instantiates a new WhatsappMessageInteractiveActionSectionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *WhatsappMessageInteractiveActionSectionsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WhatsappMessageInteractiveActionSectionsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WhatsappMessageInteractiveActionSectionsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WhatsappMessageInteractiveActionSectionsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetRows

`func (o *WhatsappMessageInteractiveActionSectionsInner) GetRows() []WhatsappMessageInteractiveActionSectionsInnerRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *WhatsappMessageInteractiveActionSectionsInner) GetRowsOk() (*[]WhatsappMessageInteractiveActionSectionsInnerRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *WhatsappMessageInteractiveActionSectionsInner) SetRows(v []WhatsappMessageInteractiveActionSectionsInnerRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *WhatsappMessageInteractiveActionSectionsInner) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


