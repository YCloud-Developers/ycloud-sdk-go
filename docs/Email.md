# Email

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID for the object. | 
**From** | Pointer to [**Mailbox**](Mailbox.md) |  | [optional] 
**To** | Pointer to [**[]Mailbox**](Mailbox.md) | The intended recipients&#39; email addresses. | [optional] 
**Cc** | Pointer to [**[]Mailbox**](Mailbox.md) | Recipients who will receive a copy of the email. | [optional] 
**Bcc** | Pointer to [**[]Mailbox**](Mailbox.md) | Recipients who will receive a blind carbon copy of the email. | [optional] 
**ReplyTo** | Pointer to [**[]Mailbox**](Mailbox.md) | If this field exists, then the reply should go to the addresses indicated in that field and not to the address(es) indicated in the &#x60;from&#x60; field. | [optional] 
**Subject** | Pointer to **string** | The email subject, which contains a short string identifying the topic of the message. | [optional] 
**Summary** | Pointer to **string** | This is a summary of your email. Max length: 70. | [optional] 
**ContentType** | Pointer to [**EmailContentType**](EmailContentType.md) |  | [optional] 
**ExternalId** | Pointer to **string** | A unique string to reference the object. This can be an order number or similar, and can be used to reconcile the object with your internal systems. | [optional] 
**CallbackUrl** | Pointer to **string** | Delivery report URL. You can provide a URL, and we will push the updated status report to your server in time. e.g., https://httpbin.org/anything?tag&#x3D;api. Note: We recommend configuring Webhook Endpoints instead. | [optional] 
**CreateTime** | Pointer to **time.Time** | The time at which this message was created, formatted in [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339). e.g., &#x60;2022-06-01T12:00:00.000Z&#x60;. | [optional] 
**TotalRecipients** | Pointer to **int32** | Total recipients of this message, including &#x60;to&#x60;, &#x60;cc&#x60; and &#x60;bcc&#x60;. | [optional] 
**TotalPrice** | Pointer to **float64** | Total price of this message. | [optional] 
**Currency** | Pointer to **string** | Price currency. [ISO 4217 currency code](https://en.wikipedia.org/wiki/ISO_4217). | [optional] 

## Methods

### NewEmail

`func NewEmail(id string, ) *Email`

NewEmail instantiates a new Email object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailWithDefaults

`func NewEmailWithDefaults() *Email`

NewEmailWithDefaults instantiates a new Email object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Email) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Email) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Email) SetId(v string)`

SetId sets Id field to given value.


### GetFrom

`func (o *Email) GetFrom() Mailbox`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Email) GetFromOk() (*Mailbox, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Email) SetFrom(v Mailbox)`

SetFrom sets From field to given value.

### HasFrom

`func (o *Email) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *Email) GetTo() []Mailbox`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Email) GetToOk() (*[]Mailbox, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Email) SetTo(v []Mailbox)`

SetTo sets To field to given value.

### HasTo

`func (o *Email) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetCc

`func (o *Email) GetCc() []Mailbox`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *Email) GetCcOk() (*[]Mailbox, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *Email) SetCc(v []Mailbox)`

SetCc sets Cc field to given value.

### HasCc

`func (o *Email) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetBcc

`func (o *Email) GetBcc() []Mailbox`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *Email) GetBccOk() (*[]Mailbox, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *Email) SetBcc(v []Mailbox)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *Email) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetReplyTo

`func (o *Email) GetReplyTo() []Mailbox`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *Email) GetReplyToOk() (*[]Mailbox, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *Email) SetReplyTo(v []Mailbox)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *Email) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSubject

`func (o *Email) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Email) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Email) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Email) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSummary

`func (o *Email) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Email) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Email) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *Email) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetContentType

`func (o *Email) GetContentType() EmailContentType`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *Email) GetContentTypeOk() (*EmailContentType, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *Email) SetContentType(v EmailContentType)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *Email) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetExternalId

`func (o *Email) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Email) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Email) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Email) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *Email) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *Email) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *Email) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *Email) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetCreateTime

`func (o *Email) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *Email) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *Email) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *Email) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetTotalRecipients

`func (o *Email) GetTotalRecipients() int32`

GetTotalRecipients returns the TotalRecipients field if non-nil, zero value otherwise.

### GetTotalRecipientsOk

`func (o *Email) GetTotalRecipientsOk() (*int32, bool)`

GetTotalRecipientsOk returns a tuple with the TotalRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecipients

`func (o *Email) SetTotalRecipients(v int32)`

SetTotalRecipients sets TotalRecipients field to given value.

### HasTotalRecipients

`func (o *Email) HasTotalRecipients() bool`

HasTotalRecipients returns a boolean if a field has been set.

### GetTotalPrice

`func (o *Email) GetTotalPrice() float64`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *Email) GetTotalPriceOk() (*float64, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *Email) SetTotalPrice(v float64)`

SetTotalPrice sets TotalPrice field to given value.

### HasTotalPrice

`func (o *Email) HasTotalPrice() bool`

HasTotalPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *Email) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Email) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Email) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Email) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


