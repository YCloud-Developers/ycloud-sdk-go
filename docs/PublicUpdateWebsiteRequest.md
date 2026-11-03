# PublicUpdateWebsiteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  |
**IncludedSubDomains** | Pointer to **[]string** | Subdomains to include in the website crawl. Omit this field to leave it unspecified. Send &#x60;[]&#x60; to clear it. | [optional]
**IncludedUrlPatterns** | Pointer to **[]string** | URL patterns to include in the website crawl. Omit this field to leave it unspecified. Send &#x60;[]&#x60; to clear it. | [optional]
**ExcludedSubDomains** | Pointer to **[]string** | Subdomains to exclude from the website crawl. Omit this field to leave it unspecified. Send &#x60;[]&#x60; to clear it. | [optional]
**ExcludedUrlPatterns** | Pointer to **[]string** | URL patterns to exclude from the website crawl. Omit this field to leave it unspecified. Send &#x60;[]&#x60; to clear it. | [optional]
**SingleUrls** | Pointer to **[]string** | Individual URLs to include in the website source. Omit this field to leave it unspecified. Send &#x60;[]&#x60; to clear it. | [optional]

## Methods

### NewPublicUpdateWebsiteRequest

`func NewPublicUpdateWebsiteRequest(url string, ) *PublicUpdateWebsiteRequest`

NewPublicUpdateWebsiteRequest instantiates a new PublicUpdateWebsiteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicUpdateWebsiteRequestWithDefaults

`func NewPublicUpdateWebsiteRequestWithDefaults() *PublicUpdateWebsiteRequest`

NewPublicUpdateWebsiteRequestWithDefaults instantiates a new PublicUpdateWebsiteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *PublicUpdateWebsiteRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PublicUpdateWebsiteRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PublicUpdateWebsiteRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetIncludedSubDomains

`func (o *PublicUpdateWebsiteRequest) GetIncludedSubDomains() []string`

GetIncludedSubDomains returns the IncludedSubDomains field if non-nil, zero value otherwise.

### GetIncludedSubDomainsOk

`func (o *PublicUpdateWebsiteRequest) GetIncludedSubDomainsOk() (*[]string, bool)`

GetIncludedSubDomainsOk returns a tuple with the IncludedSubDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSubDomains

`func (o *PublicUpdateWebsiteRequest) SetIncludedSubDomains(v []string)`

SetIncludedSubDomains sets IncludedSubDomains field to given value.

### HasIncludedSubDomains

`func (o *PublicUpdateWebsiteRequest) HasIncludedSubDomains() bool`

HasIncludedSubDomains returns a boolean if a field has been set.

### GetIncludedUrlPatterns

`func (o *PublicUpdateWebsiteRequest) GetIncludedUrlPatterns() []string`

GetIncludedUrlPatterns returns the IncludedUrlPatterns field if non-nil, zero value otherwise.

### GetIncludedUrlPatternsOk

`func (o *PublicUpdateWebsiteRequest) GetIncludedUrlPatternsOk() (*[]string, bool)`

GetIncludedUrlPatternsOk returns a tuple with the IncludedUrlPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUrlPatterns

`func (o *PublicUpdateWebsiteRequest) SetIncludedUrlPatterns(v []string)`

SetIncludedUrlPatterns sets IncludedUrlPatterns field to given value.

### HasIncludedUrlPatterns

`func (o *PublicUpdateWebsiteRequest) HasIncludedUrlPatterns() bool`

HasIncludedUrlPatterns returns a boolean if a field has been set.

### GetExcludedSubDomains

`func (o *PublicUpdateWebsiteRequest) GetExcludedSubDomains() []string`

GetExcludedSubDomains returns the ExcludedSubDomains field if non-nil, zero value otherwise.

### GetExcludedSubDomainsOk

`func (o *PublicUpdateWebsiteRequest) GetExcludedSubDomainsOk() (*[]string, bool)`

GetExcludedSubDomainsOk returns a tuple with the ExcludedSubDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSubDomains

`func (o *PublicUpdateWebsiteRequest) SetExcludedSubDomains(v []string)`

SetExcludedSubDomains sets ExcludedSubDomains field to given value.

### HasExcludedSubDomains

`func (o *PublicUpdateWebsiteRequest) HasExcludedSubDomains() bool`

HasExcludedSubDomains returns a boolean if a field has been set.

### GetExcludedUrlPatterns

`func (o *PublicUpdateWebsiteRequest) GetExcludedUrlPatterns() []string`

GetExcludedUrlPatterns returns the ExcludedUrlPatterns field if non-nil, zero value otherwise.

### GetExcludedUrlPatternsOk

`func (o *PublicUpdateWebsiteRequest) GetExcludedUrlPatternsOk() (*[]string, bool)`

GetExcludedUrlPatternsOk returns a tuple with the ExcludedUrlPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUrlPatterns

`func (o *PublicUpdateWebsiteRequest) SetExcludedUrlPatterns(v []string)`

SetExcludedUrlPatterns sets ExcludedUrlPatterns field to given value.

### HasExcludedUrlPatterns

`func (o *PublicUpdateWebsiteRequest) HasExcludedUrlPatterns() bool`

HasExcludedUrlPatterns returns a boolean if a field has been set.

### GetSingleUrls

`func (o *PublicUpdateWebsiteRequest) GetSingleUrls() []string`

GetSingleUrls returns the SingleUrls field if non-nil, zero value otherwise.

### GetSingleUrlsOk

`func (o *PublicUpdateWebsiteRequest) GetSingleUrlsOk() (*[]string, bool)`

GetSingleUrlsOk returns a tuple with the SingleUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleUrls

`func (o *PublicUpdateWebsiteRequest) SetSingleUrls(v []string)`

SetSingleUrls sets SingleUrls field to given value.

### HasSingleUrls

`func (o *PublicUpdateWebsiteRequest) HasSingleUrls() bool`

HasSingleUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
