/*
Saviynt mTLS Authentication API

mTLS Authentication

API version: 1.0
Contact: https://github.com/saviynt
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mtlsauthentication

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"os"
)


// MTLSAuthenticationAPIService MTLSAuthenticationAPI service
type MTLSAuthenticationAPIService service

type ApiDeleteKeyStoreRequest struct {
	ctx context.Context
	ApiService *MTLSAuthenticationAPIService
	alias string
}

func (r ApiDeleteKeyStoreRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteKeyStoreExecute(r)
}

/*
DeleteKeyStore Delete KeyStore

Introduced in Release v24.7, this API lets administrators delete an mTLS certificate that is uploaded in the EIC keystore.

Note: Ensure that you have specified the name of the alias in the request URL.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param alias
 @return ApiDeleteKeyStoreRequest
*/
func (a *MTLSAuthenticationAPIService) DeleteKeyStore(ctx context.Context, alias string) ApiDeleteKeyStoreRequest {
	return ApiDeleteKeyStoreRequest{
		ApiService: a,
		ctx: ctx,
		alias: alias,
	}
}

// Execute executes the request
func (a *MTLSAuthenticationAPIService) DeleteKeyStoreExecute(r ApiDeleteKeyStoreRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MTLSAuthenticationAPIService.DeleteKeyStore")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ECM/api/v5/deleteKeyStoreAlias/{alias}"
	localVarPath = strings.Replace(localVarPath, "{"+"alias"+"}", url.PathEscape(parameterValueToString(r.alias, "alias")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetKeyStoreCertificateDetailsRequest struct {
	ctx context.Context
	ApiService *MTLSAuthenticationAPIService
}

func (r ApiGetKeyStoreCertificateDetailsRequest) Execute() (*GetKeyStoreCertificateDetailsResponse, *http.Response, error) {
	return r.ApiService.GetKeyStoreCertificateDetailsExecute(r)
}

/*
GetKeyStoreCertificateDetails Get KeyStore Details

Introduced in Release v24.7, this API enables administrators to view the details of mTLS certificates that are uploaded in the EIC keystore.

A thumbprint is an unique identification of the certificate. The thumbprints attribute of this API provides a unique identification of the certificate. It displays expiry and thumbprint details in the following order:

* Leaf certificate
* Intermediate certificate
* Root certificate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetKeyStoreCertificateDetailsRequest
*/
func (a *MTLSAuthenticationAPIService) GetKeyStoreCertificateDetails(ctx context.Context) ApiGetKeyStoreCertificateDetailsRequest {
	return ApiGetKeyStoreCertificateDetailsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetKeyStoreCertificateDetailsResponse
func (a *MTLSAuthenticationAPIService) GetKeyStoreCertificateDetailsExecute(r ApiGetKeyStoreCertificateDetailsRequest) (*GetKeyStoreCertificateDetailsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetKeyStoreCertificateDetailsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MTLSAuthenticationAPIService.GetKeyStoreCertificateDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ECM/api/v5/getKeyStoreCertificateDetails"

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

type ApiUploadKeyStoreRequest struct {
	ctx context.Context
	ApiService *MTLSAuthenticationAPIService
	keyStoreFile *os.File
	keyStorePassword *string
}

func (r ApiUploadKeyStoreRequest) KeyStoreFile(keyStoreFile *os.File) ApiUploadKeyStoreRequest {
	r.keyStoreFile = keyStoreFile
	return r
}

func (r ApiUploadKeyStoreRequest) KeyStorePassword(keyStorePassword string) ApiUploadKeyStoreRequest {
	r.keyStorePassword = &keyStorePassword
	return r
}

func (r ApiUploadKeyStoreRequest) Execute() (*UploadKeyStoreResponse, *http.Response, error) {
	return r.ApiService.UploadKeyStoreExecute(r)
}

/*
UploadKeyStore Upload KeyStore

Introduced in Release v24.7, this API lets administrators upload Mutual Transport Layer Security (mTLS) certificates in the EIC keystore.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUploadKeyStoreRequest
*/
func (a *MTLSAuthenticationAPIService) UploadKeyStore(ctx context.Context) ApiUploadKeyStoreRequest {
	return ApiUploadKeyStoreRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UploadKeyStoreResponse
func (a *MTLSAuthenticationAPIService) UploadKeyStoreExecute(r ApiUploadKeyStoreRequest) (*UploadKeyStoreResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UploadKeyStoreResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MTLSAuthenticationAPIService.UploadKeyStore")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ECM/api/v5/uploadKeyStore"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	var keyStoreFileLocalVarFormFileName string
	var keyStoreFileLocalVarFileName     string
	var keyStoreFileLocalVarFileBytes    []byte

	keyStoreFileLocalVarFormFileName = "keyStoreFile"
	keyStoreFileLocalVarFile := r.keyStoreFile

	if keyStoreFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(keyStoreFileLocalVarFile)

		keyStoreFileLocalVarFileBytes = fbs
		keyStoreFileLocalVarFileName = keyStoreFileLocalVarFile.Name()
		keyStoreFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: keyStoreFileLocalVarFileBytes, fileName: keyStoreFileLocalVarFileName, formFileName: keyStoreFileLocalVarFormFileName})
	}
	if r.keyStorePassword != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "keyStorePassword", r.keyStorePassword, "", "")
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
