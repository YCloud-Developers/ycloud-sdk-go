# WebsiteDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Website resource ID. |
**Url** | **string** | Source website URL. |
**CrawlStatus** | Pointer to **string** | Current Meta crawl state. | [optional]
**PagesCrawled** | Pointer to **int32** | Number of pages Meta has processed. | [optional]
**LastCrawledAt** | Pointer to **int64** | Most recent crawl timestamp. | [optional]
**CreatedAt** | Pointer to **int64** | Resource creation timestamp. | [optional]
**IncludedSubDomains** | Pointer to **[]string** | Subdomains included in the website crawl, when returned by Meta. | [optional]
**IncludedUrlPatterns** | Pointer to **[]string** | URL patterns included in the website crawl, when returned by Meta. | [optional]
**ExcludedSubDomains** | Pointer to **[]string** | Subdomains excluded from the website crawl, when returned by Meta. | [optional]
**ExcludedUrlPatterns** | Pointer to **[]string** | URL patterns excluded from the website crawl, when returned by Meta. | [optional]
**SingleUrls** | Pointer to **[]string** | Individual URLs included in the website source, when returned by Meta. | [optional]
**CrawlError** | Pointer to **string** | Optional crawl error reported by Meta; an empty string is preserved. | [optional]

## Methods

### NewWebsiteDTO

`func NewWebsiteDTO(id string, url string, ) *WebsiteDTO`

NewWebsiteDTO instantiates a new WebsiteDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsiteDTOWithDefaults

`func NewWebsiteDTOWithDefaults() *WebsiteDTO`

NewWebsiteDTOWithDefaults instantiates a new WebsiteDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebsiteDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebsiteDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebsiteDTO) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *WebsiteDTO) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebsiteDTO) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebsiteDTO) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCrawlStatus

`func (o *WebsiteDTO) GetCrawlStatus() string`

GetCrawlStatus returns the CrawlStatus field if non-nil, zero value otherwise.

### GetCrawlStatusOk

`func (o *WebsiteDTO) GetCrawlStatusOk() (*string, bool)`

GetCrawlStatusOk returns a tuple with the CrawlStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawlStatus

`func (o *WebsiteDTO) SetCrawlStatus(v string)`

SetCrawlStatus sets CrawlStatus field to given value.

### HasCrawlStatus

`func (o *WebsiteDTO) HasCrawlStatus() bool`

HasCrawlStatus returns a boolean if a field has been set.

### GetPagesCrawled

`func (o *WebsiteDTO) GetPagesCrawled() int32`

GetPagesCrawled returns the PagesCrawled field if non-nil, zero value otherwise.

### GetPagesCrawledOk

`func (o *WebsiteDTO) GetPagesCrawledOk() (*int32, bool)`

GetPagesCrawledOk returns a tuple with the PagesCrawled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagesCrawled

`func (o *WebsiteDTO) SetPagesCrawled(v int32)`

SetPagesCrawled sets PagesCrawled field to given value.

### HasPagesCrawled

`func (o *WebsiteDTO) HasPagesCrawled() bool`

HasPagesCrawled returns a boolean if a field has been set.

### GetLastCrawledAt

`func (o *WebsiteDTO) GetLastCrawledAt() int64`

GetLastCrawledAt returns the LastCrawledAt field if non-nil, zero value otherwise.

### GetLastCrawledAtOk

`func (o *WebsiteDTO) GetLastCrawledAtOk() (*int64, bool)`

GetLastCrawledAtOk returns a tuple with the LastCrawledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCrawledAt

`func (o *WebsiteDTO) SetLastCrawledAt(v int64)`

SetLastCrawledAt sets LastCrawledAt field to given value.

### HasLastCrawledAt

`func (o *WebsiteDTO) HasLastCrawledAt() bool`

HasLastCrawledAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WebsiteDTO) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebsiteDTO) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebsiteDTO) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WebsiteDTO) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetIncludedSubDomains

`func (o *WebsiteDTO) GetIncludedSubDomains() []string`

GetIncludedSubDomains returns the IncludedSubDomains field if non-nil, zero value otherwise.

### GetIncludedSubDomainsOk

`func (o *WebsiteDTO) GetIncludedSubDomainsOk() (*[]string, bool)`

GetIncludedSubDomainsOk returns a tuple with the IncludedSubDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSubDomains

`func (o *WebsiteDTO) SetIncludedSubDomains(v []string)`

SetIncludedSubDomains sets IncludedSubDomains field to given value.

### HasIncludedSubDomains

`func (o *WebsiteDTO) HasIncludedSubDomains() bool`

HasIncludedSubDomains returns a boolean if a field has been set.

### GetIncludedUrlPatterns

`func (o *WebsiteDTO) GetIncludedUrlPatterns() []string`

GetIncludedUrlPatterns returns the IncludedUrlPatterns field if non-nil, zero value otherwise.

### GetIncludedUrlPatternsOk

`func (o *WebsiteDTO) GetIncludedUrlPatternsOk() (*[]string, bool)`

GetIncludedUrlPatternsOk returns a tuple with the IncludedUrlPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUrlPatterns

`func (o *WebsiteDTO) SetIncludedUrlPatterns(v []string)`

SetIncludedUrlPatterns sets IncludedUrlPatterns field to given value.

### HasIncludedUrlPatterns

`func (o *WebsiteDTO) HasIncludedUrlPatterns() bool`

HasIncludedUrlPatterns returns a boolean if a field has been set.

### GetExcludedSubDomains

`func (o *WebsiteDTO) GetExcludedSubDomains() []string`

GetExcludedSubDomains returns the ExcludedSubDomains field if non-nil, zero value otherwise.

### GetExcludedSubDomainsOk

`func (o *WebsiteDTO) GetExcludedSubDomainsOk() (*[]string, bool)`

GetExcludedSubDomainsOk returns a tuple with the ExcludedSubDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSubDomains

`func (o *WebsiteDTO) SetExcludedSubDomains(v []string)`

SetExcludedSubDomains sets ExcludedSubDomains field to given value.

### HasExcludedSubDomains

`func (o *WebsiteDTO) HasExcludedSubDomains() bool`

HasExcludedSubDomains returns a boolean if a field has been set.

### GetExcludedUrlPatterns

`func (o *WebsiteDTO) GetExcludedUrlPatterns() []string`

GetExcludedUrlPatterns returns the ExcludedUrlPatterns field if non-nil, zero value otherwise.

### GetExcludedUrlPatternsOk

`func (o *WebsiteDTO) GetExcludedUrlPatternsOk() (*[]string, bool)`

GetExcludedUrlPatternsOk returns a tuple with the ExcludedUrlPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUrlPatterns

`func (o *WebsiteDTO) SetExcludedUrlPatterns(v []string)`

SetExcludedUrlPatterns sets ExcludedUrlPatterns field to given value.

### HasExcludedUrlPatterns

`func (o *WebsiteDTO) HasExcludedUrlPatterns() bool`

HasExcludedUrlPatterns returns a boolean if a field has been set.

### GetSingleUrls

`func (o *WebsiteDTO) GetSingleUrls() []string`

GetSingleUrls returns the SingleUrls field if non-nil, zero value otherwise.

### GetSingleUrlsOk

`func (o *WebsiteDTO) GetSingleUrlsOk() (*[]string, bool)`

GetSingleUrlsOk returns a tuple with the SingleUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleUrls

`func (o *WebsiteDTO) SetSingleUrls(v []string)`

SetSingleUrls sets SingleUrls field to given value.

### HasSingleUrls

`func (o *WebsiteDTO) HasSingleUrls() bool`

HasSingleUrls returns a boolean if a field has been set.

### GetCrawlError

`func (o *WebsiteDTO) GetCrawlError() string`

GetCrawlError returns the CrawlError field if non-nil, zero value otherwise.

### GetCrawlErrorOk

`func (o *WebsiteDTO) GetCrawlErrorOk() (*string, bool)`

GetCrawlErrorOk returns a tuple with the CrawlError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawlError

`func (o *WebsiteDTO) SetCrawlError(v string)`

SetCrawlError sets CrawlError field to given value.

### HasCrawlError

`func (o *WebsiteDTO) HasCrawlError() bool`

HasCrawlError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
