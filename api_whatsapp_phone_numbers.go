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

// WhatsappPhoneNumbersApiService WhatsappPhoneNumbersApi service
type WhatsappPhoneNumbersApiService service

type WhatsappPhoneNumbersApiListRequest struct {
	ctx          context.Context
	ApiService   *WhatsappPhoneNumbersApiService
	page         *int32
	limit        *int32
	includeTotal *bool
	filterWabaId *string
}

// Page number of the results to be returned, 1-based.
func (r WhatsappPhoneNumbersApiListRequest) Page(page int32) WhatsappPhoneNumbersApiListRequest {
	r.page = &page
	return r
}

// A limit on the number of results to be returned, or number of results per page, between 1 and 100, defaults to 10.
func (r WhatsappPhoneNumbersApiListRequest) Limit(limit int32) WhatsappPhoneNumbersApiListRequest {
	r.limit = &limit
	return r
}

// Return results inside an object that contains the total result count or not.
func (r WhatsappPhoneNumbersApiListRequest) IncludeTotal(includeTotal bool) WhatsappPhoneNumbersApiListRequest {
	r.includeTotal = &includeTotal
	return r
}

// WhatsApp Business Account ID.
func (r WhatsappPhoneNumbersApiListRequest) FilterWabaId(filterWabaId string) WhatsappPhoneNumbersApiListRequest {
	r.filterWabaId = &filterWabaId
	return r
}

func (r WhatsappPhoneNumbersApiListRequest) Execute() (*WhatsappPhoneNumberPage, *http.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List List phone numbers

Returns a paginated list of WhatsApp business phone numbers you've registered on YCloud.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return WhatsappPhoneNumbersApiListRequest
*/
func (a *WhatsappPhoneNumbersApiService) List(ctx context.Context) WhatsappPhoneNumbersApiListRequest {
	return WhatsappPhoneNumbersApiListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return WhatsappPhoneNumberPage
func (a *WhatsappPhoneNumbersApiService) ListExecute(r WhatsappPhoneNumbersApiListRequest) (*WhatsappPhoneNumberPage, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *WhatsappPhoneNumberPage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WhatsappPhoneNumbersApiService.List")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/whatsapp/phoneNumbers"

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
	if r.filterWabaId != nil {
		localVarQueryParams.Add("filter.wabaId", parameterToString(*r.filterWabaId, ""))
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

type WhatsappPhoneNumbersApiRegisterRequest struct {
	ctx         context.Context
	ApiService  *WhatsappPhoneNumbersApiService
	wabaId      string
	phoneNumber string
}

func (r WhatsappPhoneNumbersApiRegisterRequest) Execute() (*WhatsappPhoneNumber, *http.Response, error) {
	return r.ApiService.RegisterExecute(r)
}

/*
Register Register a phone number

Registers a WhatsApp business phone number.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param wabaId WhatsApp Business Account ID.
	@param phoneNumber Phone number ID.
	@return WhatsappPhoneNumbersApiRegisterRequest
*/
func (a *WhatsappPhoneNumbersApiService) Register(ctx context.Context, wabaId string, phoneNumber string) WhatsappPhoneNumbersApiRegisterRequest {
	return WhatsappPhoneNumbersApiRegisterRequest{
		ApiService:  a,
		ctx:         ctx,
		wabaId:      wabaId,
		phoneNumber: phoneNumber,
	}
}

// Execute executes the request
//
//	@return WhatsappPhoneNumber
func (a *WhatsappPhoneNumbersApiService) RegisterExecute(r WhatsappPhoneNumbersApiRegisterRequest) (*WhatsappPhoneNumber, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *WhatsappPhoneNumber
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WhatsappPhoneNumbersApiService.Register")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/whatsapp/phoneNumbers/{wabaId}/{phoneNumber}/register"
	localVarPath = strings.Replace(localVarPath, "{"+"wabaId"+"}", url.PathEscape(parameterToString(r.wabaId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"phoneNumber"+"}", url.PathEscape(parameterToString(r.phoneNumber, "")), -1)

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

type WhatsappPhoneNumbersApiRetrieveRequest struct {
	ctx         context.Context
	ApiService  *WhatsappPhoneNumbersApiService
	wabaId      string
	phoneNumber string
}

func (r WhatsappPhoneNumbersApiRetrieveRequest) Execute() (*WhatsappPhoneNumber, *http.Response, error) {
	return r.ApiService.RetrieveExecute(r)
}

/*
Retrieve Retrieve a phone number

Retrieves a WhatsApp business phone number you've registered on YCloud.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param wabaId WhatsApp Business Account ID.
	@param phoneNumber Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
	@return WhatsappPhoneNumbersApiRetrieveRequest
*/
func (a *WhatsappPhoneNumbersApiService) Retrieve(ctx context.Context, wabaId string, phoneNumber string) WhatsappPhoneNumbersApiRetrieveRequest {
	return WhatsappPhoneNumbersApiRetrieveRequest{
		ApiService:  a,
		ctx:         ctx,
		wabaId:      wabaId,
		phoneNumber: phoneNumber,
	}
}

// Execute executes the request
//
//	@return WhatsappPhoneNumber
func (a *WhatsappPhoneNumbersApiService) RetrieveExecute(r WhatsappPhoneNumbersApiRetrieveRequest) (*WhatsappPhoneNumber, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *WhatsappPhoneNumber
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WhatsappPhoneNumbersApiService.Retrieve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/whatsapp/phoneNumbers/{wabaId}/{phoneNumber}"
	localVarPath = strings.Replace(localVarPath, "{"+"wabaId"+"}", url.PathEscape(parameterToString(r.wabaId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"phoneNumber"+"}", url.PathEscape(parameterToString(r.phoneNumber, "")), -1)

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

type WhatsappPhoneNumbersApiRetrieveCommerceSettingsRequest struct {
	ctx         context.Context
	ApiService  *WhatsappPhoneNumbersApiService
	wabaId      string
	phoneNumber string
}

func (r WhatsappPhoneNumbersApiRetrieveCommerceSettingsRequest) Execute() (*WhatsappCommerceSettings, *http.Response, error) {
	return r.ApiService.RetrieveCommerceSettingsExecute(r)
}

/*
RetrieveCommerceSettings Retrieve commerce settings

Retrieves a WhatsApp business phone number's commerce settings.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param wabaId WhatsApp Business Account ID.
	@param phoneNumber Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
	@return WhatsappPhoneNumbersApiRetrieveCommerceSettingsRequest
*/
func (a *WhatsappPhoneNumbersApiService) RetrieveCommerceSettings(ctx context.Context, wabaId string, phoneNumber string) WhatsappPhoneNumbersApiRetrieveCommerceSettingsRequest {
	return WhatsappPhoneNumbersApiRetrieveCommerceSettingsRequest{
		ApiService:  a,
		ctx:         ctx,
		wabaId:      wabaId,
		phoneNumber: phoneNumber,
	}
}

// Execute executes the request
//
//	@return WhatsappCommerceSettings
func (a *WhatsappPhoneNumbersApiService) RetrieveCommerceSettingsExecute(r WhatsappPhoneNumbersApiRetrieveCommerceSettingsRequest) (*WhatsappCommerceSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *WhatsappCommerceSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WhatsappPhoneNumbersApiService.RetrieveCommerceSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/whatsapp/phoneNumbers/{wabaId}/{phoneNumber}/whatsappCommerceSettings"
	localVarPath = strings.Replace(localVarPath, "{"+"wabaId"+"}", url.PathEscape(parameterToString(r.wabaId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"phoneNumber"+"}", url.PathEscape(parameterToString(r.phoneNumber, "")), -1)

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

type WhatsappPhoneNumbersApiRetrieveProfileRequest struct {
	ctx         context.Context
	ApiService  *WhatsappPhoneNumbersApiService
	wabaId      string
	phoneNumber string
}

func (r WhatsappPhoneNumbersApiRetrieveProfileRequest) Execute() (*WhatsappPhoneNumberProfile, *http.Response, error) {
	return r.ApiService.RetrieveProfileExecute(r)
}

/*
RetrieveProfile Retrieve a phone number profile

Retrieves a WhatsApp business phone number's profile. Customers can view your business profile by clicking your business's name or number in a conversation thread.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param wabaId WhatsApp Business Account ID.
	@param phoneNumber Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
	@return WhatsappPhoneNumbersApiRetrieveProfileRequest
*/
func (a *WhatsappPhoneNumbersApiService) RetrieveProfile(ctx context.Context, wabaId string, phoneNumber string) WhatsappPhoneNumbersApiRetrieveProfileRequest {
	return WhatsappPhoneNumbersApiRetrieveProfileRequest{
		ApiService:  a,
		ctx:         ctx,
		wabaId:      wabaId,
		phoneNumber: phoneNumber,
	}
}

// Execute executes the request
//
//	@return WhatsappPhoneNumberProfile
func (a *WhatsappPhoneNumbersApiService) RetrieveProfileExecute(r WhatsappPhoneNumbersApiRetrieveProfileRequest) (*WhatsappPhoneNumberProfile, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *WhatsappPhoneNumberProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WhatsappPhoneNumbersApiService.RetrieveProfile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/whatsapp/phoneNumbers/{wabaId}/{phoneNumber}/profile"
	localVarPath = strings.Replace(localVarPath, "{"+"wabaId"+"}", url.PathEscape(parameterToString(r.wabaId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"phoneNumber"+"}", url.PathEscape(parameterToString(r.phoneNumber, "")), -1)

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

type WhatsappPhoneNumbersApiUpdateCommerceSettingsRequest struct {
	ctx                                   context.Context
	ApiService                            *WhatsappPhoneNumbersApiService
	wabaId                                string
	phoneNumber                           string
	whatsappCommerceSettingsUpdateRequest *WhatsappCommerceSettingsUpdateRequest
}

func (r WhatsappPhoneNumbersApiUpdateCommerceSettingsRequest) WhatsappCommerceSettingsUpdateRequest(whatsappCommerceSettingsUpdateRequest WhatsappCommerceSettingsUpdateRequest) WhatsappPhoneNumbersApiUpdateCommerceSettingsRequest {
	r.whatsappCommerceSettingsUpdateRequest = &whatsappCommerceSettingsUpdateRequest
	return r
}

func (r WhatsappPhoneNumbersApiUpdateCommerceSettingsRequest) Execute() (*WhatsappCommerceSettings, *http.Response, error) {
	return r.ApiService.UpdateCommerceSettingsExecute(r)
}

/*
UpdateCommerceSettings Update commerce settings

Updates a WhatsApp business phone number's commerce settings.
Use this endpoint to enable or disable the shopping cart or the product catalog for a specific business phone number.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param wabaId WhatsApp Business Account ID.
	@param phoneNumber Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
	@return WhatsappPhoneNumbersApiUpdateCommerceSettingsRequest
*/
func (a *WhatsappPhoneNumbersApiService) UpdateCommerceSettings(ctx context.Context, wabaId string, phoneNumber string) WhatsappPhoneNumbersApiUpdateCommerceSettingsRequest {
	return WhatsappPhoneNumbersApiUpdateCommerceSettingsRequest{
		ApiService:  a,
		ctx:         ctx,
		wabaId:      wabaId,
		phoneNumber: phoneNumber,
	}
}

// Execute executes the request
//
//	@return WhatsappCommerceSettings
func (a *WhatsappPhoneNumbersApiService) UpdateCommerceSettingsExecute(r WhatsappPhoneNumbersApiUpdateCommerceSettingsRequest) (*WhatsappCommerceSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *WhatsappCommerceSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WhatsappPhoneNumbersApiService.UpdateCommerceSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/whatsapp/phoneNumbers/{wabaId}/{phoneNumber}/whatsappCommerceSettings"
	localVarPath = strings.Replace(localVarPath, "{"+"wabaId"+"}", url.PathEscape(parameterToString(r.wabaId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"phoneNumber"+"}", url.PathEscape(parameterToString(r.phoneNumber, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.whatsappCommerceSettingsUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("whatsappCommerceSettingsUpdateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.whatsappCommerceSettingsUpdateRequest
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

type WhatsappPhoneNumbersApiUpdateProfileRequest struct {
	ctx                                     context.Context
	ApiService                              *WhatsappPhoneNumbersApiService
	wabaId                                  string
	phoneNumber                             string
	whatsappPhoneNumberProfileUpdateRequest *WhatsappPhoneNumberProfileUpdateRequest
}

func (r WhatsappPhoneNumbersApiUpdateProfileRequest) WhatsappPhoneNumberProfileUpdateRequest(whatsappPhoneNumberProfileUpdateRequest WhatsappPhoneNumberProfileUpdateRequest) WhatsappPhoneNumbersApiUpdateProfileRequest {
	r.whatsappPhoneNumberProfileUpdateRequest = &whatsappPhoneNumberProfileUpdateRequest
	return r
}

func (r WhatsappPhoneNumbersApiUpdateProfileRequest) Execute() (*WhatsappPhoneNumberProfile, *http.Response, error) {
	return r.ApiService.UpdateProfileExecute(r)
}

/*
UpdateProfile Update a phone number profile

Updates a WhatsApp business phone number profile.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param wabaId WhatsApp Business Account ID.
	@param phoneNumber Phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format.
	@return WhatsappPhoneNumbersApiUpdateProfileRequest
*/
func (a *WhatsappPhoneNumbersApiService) UpdateProfile(ctx context.Context, wabaId string, phoneNumber string) WhatsappPhoneNumbersApiUpdateProfileRequest {
	return WhatsappPhoneNumbersApiUpdateProfileRequest{
		ApiService:  a,
		ctx:         ctx,
		wabaId:      wabaId,
		phoneNumber: phoneNumber,
	}
}

// Execute executes the request
//
//	@return WhatsappPhoneNumberProfile
func (a *WhatsappPhoneNumbersApiService) UpdateProfileExecute(r WhatsappPhoneNumbersApiUpdateProfileRequest) (*WhatsappPhoneNumberProfile, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *WhatsappPhoneNumberProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WhatsappPhoneNumbersApiService.UpdateProfile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/whatsapp/phoneNumbers/{wabaId}/{phoneNumber}/profile"
	localVarPath = strings.Replace(localVarPath, "{"+"wabaId"+"}", url.PathEscape(parameterToString(r.wabaId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"phoneNumber"+"}", url.PathEscape(parameterToString(r.phoneNumber, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.whatsappPhoneNumberProfileUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("whatsappPhoneNumberProfileUpdateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.whatsappPhoneNumberProfileUpdateRequest
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
