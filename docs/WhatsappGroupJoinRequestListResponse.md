# WhatsappGroupJoinRequestListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]WhatsappGroupJoinRequest**](WhatsappGroupJoinRequest.md) |  | [optional] 
**Paging** | Pointer to [**WhatsappGroupPaging**](WhatsappGroupPaging.md) |  | [optional] 

## Methods

### NewWhatsappGroupJoinRequestListResponse

`func NewWhatsappGroupJoinRequestListResponse() *WhatsappGroupJoinRequestListResponse`

NewWhatsappGroupJoinRequestListResponse instantiates a new WhatsappGroupJoinRequestListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsappGroupJoinRequestListResponseWithDefaults

`func NewWhatsappGroupJoinRequestListResponseWithDefaults() *WhatsappGroupJoinRequestListResponse`

NewWhatsappGroupJoinRequestListResponseWithDefaults instantiates a new WhatsappGroupJoinRequestListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WhatsappGroupJoinRequestListResponse) GetData() []WhatsappGroupJoinRequest`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WhatsappGroupJoinRequestListResponse) GetDataOk() (*[]WhatsappGroupJoinRequest, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WhatsappGroupJoinRequestListResponse) SetData(v []WhatsappGroupJoinRequest)`

SetData sets Data field to given value.

### HasData

`func (o *WhatsappGroupJoinRequestListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPaging

`func (o *WhatsappGroupJoinRequestListResponse) GetPaging() WhatsappGroupPaging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *WhatsappGroupJoinRequestListResponse) GetPagingOk() (*WhatsappGroupPaging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *WhatsappGroupJoinRequestListResponse) SetPaging(v WhatsappGroupPaging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *WhatsappGroupJoinRequestListResponse) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


