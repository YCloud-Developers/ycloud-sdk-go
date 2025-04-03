/*
YCloud API

The [YCloud](https://ycloud.com) API is organized around [REST](https://en.wikipedia.org/wiki/Representational_state_transfer). Our API is designed to have predictable, resource-oriented URLs, return [JSON](https://www.json.org) responses, and use standard HTTP response codes and verbs.

API version: v2
Contact: service@ycloud.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ycloud

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// WhatsappBusinessAccountsApiService WhatsappBusinessAccountsApi service
type WhatsappBusinessAccountsApiService service

type WhatsappBusinessAccountsApiListRequest struct {
	ctx context.Context
	ApiService *WhatsappBusinessAccountsApiService
	page *int32
	limit *int32
	includeTotal *bool
	filterAccountReviewStatus *string
}

// Page number of the results to be returned, 1-based.
func (r WhatsappBusinessAccountsApiListRequest) Page(page int32) WhatsappBusinessAccountsApiListRequest {
	r.page = &page
	return r
}

// A limit on the number of results to be returned, or number of results per page, between 1 and 100, defaults to 10.
func (r WhatsappBusinessAccountsApiListRequest) Limit(limit int32) WhatsappBusinessAccountsApiListRequest {
	r.limit = &limit
	return r
}

// Return results inside an object that contains the total result count or not.
func (r WhatsappBusinessAccountsApiListRequest) IncludeTotal(includeTotal bool) WhatsappBusinessAccountsApiListRequest {
	r.includeTotal = &includeTotal
	return r
}

// WhatsApp Business Account review status.
func (r WhatsappBusinessAccountsApiListRequest) FilterAccountReviewStatus(filterAccountReviewStatus string) WhatsappBusinessAccountsApiListRequest {
	r.filterAccountReviewStatus = &filterAccountReviewStatus
	return r
}

func (r WhatsappBusinessAccountsApiListRequest) Execute() (*WhatsappBusinessAccountPage, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List List WABAs

Returns a paginated list of WhatsApp business accounts you've registered.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return WhatsappBusinessAccountsApiListRequest
*/
func (a *WhatsappBusinessAccountsApiService) List(ctx context.Context) WhatsappBusinessAccountsApiListRequest {
	return WhatsappBusinessAccountsApiListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WhatsappBusinessAccountPage
func (a *WhatsappBusinessAccountsApiService) ListExecute(r WhatsappBusinessAccountsApiListRequest) (*WhatsappBusinessAccountPage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WhatsappBusinessAccountPage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WhatsappBusinessAccountsApiService.List")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/whatsapp/businessAccounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.includeTotal != nil {
		localVarQueryParams.Add("includeTotal", parameterToString(*r.includeTotal, ""))
	}
	if r.filterAccountReviewStatus != nil {
		localVarQueryParams.Add("filter.accountReviewStatus", parameterToString(*r.filterAccountReviewStatus, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type WhatsappBusinessAccountsApiRetrieveRequest struct {
	ctx context.Context
	ApiService *WhatsappBusinessAccountsApiService
	id string
}

func (r WhatsappBusinessAccountsApiRetrieveRequest) Execute() (*WhatsappBusinessAccount, *http.Response, error) {
	return r.ApiService.RetrieveExecute(r)
}

/*
Retrieve Retrieve a WABA

Retrieves a WABA you've registered.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id WhatsApp Business Account ID.
 @return WhatsappBusinessAccountsApiRetrieveRequest
*/
func (a *WhatsappBusinessAccountsApiService) Retrieve(ctx context.Context, id string) WhatsappBusinessAccountsApiRetrieveRequest {
	return WhatsappBusinessAccountsApiRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return WhatsappBusinessAccount
func (a *WhatsappBusinessAccountsApiService) RetrieveExecute(r WhatsappBusinessAccountsApiRetrieveRequest) (*WhatsappBusinessAccount, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WhatsappBusinessAccount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WhatsappBusinessAccountsApiService.Retrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/whatsapp/businessAccounts/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
