/*
 * DC/OS
 *
 * DC/OS API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dcos

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type SecretsApiService service

/*
SecretsApiService Create a secret in the store at the path.
Create a secret in the store at the path.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param store The backend to store the secret in.
 * @param pathToSecret The path to store the secret in.
 * @param secretsV1Secret Secret value.
*/
func (a *SecretsApiService) CreateSecret(ctx context.Context, store string, pathToSecret string, secretsV1Secret SecretsV1Secret) (*http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/secrets/v1/secret/{store}/{path-to-secret}"
	localVarPath = strings.Replace(localVarPath, "{"+"store"+"}", fmt.Sprintf("%v", store), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"path-to-secret"+"}", fmt.Sprintf("%v", pathToSecret), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &secretsV1Secret
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
SecretsApiService Delete a secret.
Delete a secret.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param store The backend to delete the secret from.
 * @param pathToSecret The path to the secret to delete.
*/
func (a *SecretsApiService) DeleteSecret(ctx context.Context, store string, pathToSecret string) (*http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/secrets/v1/secret/{store}/{path-to-secret}"
	localVarPath = strings.Replace(localVarPath, "{"+"store"+"}", fmt.Sprintf("%v", store), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"path-to-secret"+"}", fmt.Sprintf("%v", pathToSecret), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
SecretsApiService Read or list a secret from the store by its path.
Read or list a secret from the store by its path.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param store The backend store from which to get the secret.
 * @param pathToSecret The path to store the secret in.
 * @param optional nil or *GetSecretOpts - Optional Parameters:
 * @param "List" (optional.String) -  If set to true, list only secret keys.
@return SecretsV1Secret
*/

type GetSecretOpts struct {
	List optional.String
}

func (a *SecretsApiService) GetSecret(ctx context.Context, store string, pathToSecret string, localVarOptionals *GetSecretOpts) (SecretsV1Secret, *http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SecretsV1Secret
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/secrets/v1/secret/{store}/{path-to-secret}"
	localVarPath = strings.Replace(localVarPath, "{"+"store"+"}", fmt.Sprintf("%v", store), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"path-to-secret"+"}", fmt.Sprintf("%v", pathToSecret), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.List.IsSet() {
		localVarQueryParams.Add("list", parameterToString(localVarOptionals.List.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v SecretsV1Secret
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
SecretsApiService Update secret.
Update existing secret in the specified store.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param store The backend to store the secret in.
 * @param pathToSecret The path for the secret to update.
 * @param secretsV1Secret Secret value.
*/
func (a *SecretsApiService) UpdateSecret(ctx context.Context, store string, pathToSecret string, secretsV1Secret SecretsV1Secret) (*http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/secrets/v1/secret/{store}/{path-to-secret}"
	localVarPath = strings.Replace(localVarPath, "{"+"store"+"}", fmt.Sprintf("%v", store), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"path-to-secret"+"}", fmt.Sprintf("%v", pathToSecret), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &secretsV1Secret
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
