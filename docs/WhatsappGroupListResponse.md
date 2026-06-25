# WhatsappGroupListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]WhatsappGroupListItem**](WhatsappGroupListItem.md) |  | [optional] 
**Paging** | Pointer to [**WhatsappGroupPaging**](WhatsappGroupPaging.md) |  | [optional] 

## Methods

### NewWhatsappGroupListResponse

`func NewWhatsappGroupListResponse() *WhatsappGroupListResponse`

NewWhatsappGroupListResponse instantiates a new WhatsappGroupListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappGroupListResponseWithDefaults

`func NewWhatsappGroupListResponseWithDefaults() *WhatsappGroupListResponse`

NewWhatsappGroupListResponseWithDefaults instantiates a new WhatsappGroupListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WhatsappGroupListResponse) GetData() []WhatsappGroupListItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WhatsappGroupListResponse) GetDataOk() (*[]WhatsappGroupListItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WhatsappGroupListResponse) SetData(v []WhatsappGroupListItem)`

SetData sets Data field to given value.

### HasData

`func (o *WhatsappGroupListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPaging

`func (o *WhatsappGroupListResponse) GetPaging() WhatsappGroupPaging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *WhatsappGroupListResponse) GetPagingOk() (*WhatsappGroupPaging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *WhatsappGroupListResponse) SetPaging(v WhatsappGroupPaging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *WhatsappGroupListResponse) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


