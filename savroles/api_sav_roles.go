/*
Saviynt SAV Roles API

Saviynt SAV Roles API

API version: 1.0
Contact: https://github.com/saviynt
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package savroles

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// SAVRolesAPIService SAVRolesAPI service
type SAVRolesAPIService service

type ApiGetAllSAVRolesRequest struct {
	ctx context.Context
	ApiService *SAVRolesAPIService
}

func (r ApiGetAllSAVRolesRequest) Execute() (*GetAllSAVRolesResponse, *http.Response, error) {
	return r.ApiService.GetAllSAVRolesExecute(r)
}

/*
GetAllSAVRoles Get All SAV Roles

This API returns all out-of-the-box (OOTB) and custom SAV roles along with their properties.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAllSAVRolesRequest
*/
func (a *SAVRolesAPIService) GetAllSAVRoles(ctx context.Context) ApiGetAllSAVRolesRequest {
	return ApiGetAllSAVRolesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetAllSAVRolesResponse
func (a *SAVRolesAPIService) GetAllSAVRolesExecute(r ApiGetAllSAVRolesRequest) (*GetAllSAVRolesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetAllSAVRolesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SAVRolesAPIService.GetAllSAVRoles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ECMv6/api/userms/savroles"

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetSAVRoleUsersRequest struct {
	ctx context.Context
	ApiService *SAVRolesAPIService
	limit *string
	offset *string
	savRoleName string
}

// Specifies the number of retrieved results
func (r ApiGetSAVRoleUsersRequest) Limit(limit string) ApiGetSAVRoleUsersRequest {
	r.limit = &limit
	return r
}

// Specifies the number of rows of the result to skip before any rows are retrieved, and must be used with the &#x60;limit&#x60; parameter
func (r ApiGetSAVRoleUsersRequest) Offset(offset string) ApiGetSAVRoleUsersRequest {
	r.offset = &offset
	return r
}

func (r ApiGetSAVRoleUsersRequest) Execute() (*GetSAVRoleUsersResponse, *http.Response, error) {
	return r.ApiService.GetSAVRoleUsersExecute(r)
}

/*
GetSAVRoleUsers Get Users Associated with a Particular SAV Role

This API returns a list of users associated with a particular SAV role.



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param savRoleName The `ROLENAME` field in geAllSAVRoles API
 @return ApiGetSAVRoleUsersRequest
*/
func (a *SAVRolesAPIService) GetSAVRoleUsers(ctx context.Context, savRoleName string) ApiGetSAVRoleUsersRequest {
	return ApiGetSAVRoleUsersRequest{
		ApiService: a,
		ctx: ctx,
		savRoleName: savRoleName,
	}
}

// Execute executes the request
//  @return GetSAVRoleUsersResponse
func (a *SAVRolesAPIService) GetSAVRoleUsersExecute(r ApiGetSAVRoleUsersRequest) (*GetSAVRoleUsersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetSAVRoleUsersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SAVRolesAPIService.GetSAVRoleUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ECMv6/api/userms/savroles/{savRoleName}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"savRoleName"+"}", url.PathEscape(parameterValueToString(r.savRoleName, "savRoleName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.limit == nil {
		return localVarReturnValue, nil, reportError("limit is required and must be specified")
	}
	if r.offset == nil {
		return localVarReturnValue, nil, reportError("offset is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
