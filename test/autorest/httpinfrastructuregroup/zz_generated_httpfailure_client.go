//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package httpinfrastructuregroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// HTTPFailureClient contains the methods for the HTTPFailure group.
// Don't use this type directly, use NewHTTPFailureClient() instead.
type HTTPFailureClient struct {
	pl runtime.Pipeline
}

// NewHTTPFailureClient creates a new instance of HTTPFailureClient with the specified values.
// options - pass nil to accept the default values.
func NewHTTPFailureClient(options *azcore.ClientOptions) *HTTPFailureClient {
	if options == nil {
		options = &azcore.ClientOptions{}
	}
	client := &HTTPFailureClient{
		pl: runtime.NewPipeline(moduleName, moduleVersion, runtime.PipelineOptions{}, options),
	}
	return client
}

// GetEmptyError - Get empty error form server
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - HTTPFailureClientGetEmptyErrorOptions contains the optional parameters for the HTTPFailureClient.GetEmptyError
// method.
func (client *HTTPFailureClient) GetEmptyError(ctx context.Context, options *HTTPFailureClientGetEmptyErrorOptions) (HTTPFailureClientGetEmptyErrorResponse, error) {
	req, err := client.getEmptyErrorCreateRequest(ctx, options)
	if err != nil {
		return HTTPFailureClientGetEmptyErrorResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPFailureClientGetEmptyErrorResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPFailureClientGetEmptyErrorResponse{}, runtime.NewResponseError(resp)
	}
	return client.getEmptyErrorHandleResponse(resp)
}

// getEmptyErrorCreateRequest creates the GetEmptyError request.
func (client *HTTPFailureClient) getEmptyErrorCreateRequest(ctx context.Context, options *HTTPFailureClientGetEmptyErrorOptions) (*policy.Request, error) {
	urlPath := "/http/failure/emptybody/error"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getEmptyErrorHandleResponse handles the GetEmptyError response.
func (client *HTTPFailureClient) getEmptyErrorHandleResponse(resp *http.Response) (HTTPFailureClientGetEmptyErrorResponse, error) {
	result := HTTPFailureClientGetEmptyErrorResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return HTTPFailureClientGetEmptyErrorResponse{}, err
	}
	return result, nil
}

// GetNoModelEmpty - Get empty response from server
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - HTTPFailureClientGetNoModelEmptyOptions contains the optional parameters for the HTTPFailureClient.GetNoModelEmpty
// method.
func (client *HTTPFailureClient) GetNoModelEmpty(ctx context.Context, options *HTTPFailureClientGetNoModelEmptyOptions) (HTTPFailureClientGetNoModelEmptyResponse, error) {
	req, err := client.getNoModelEmptyCreateRequest(ctx, options)
	if err != nil {
		return HTTPFailureClientGetNoModelEmptyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPFailureClientGetNoModelEmptyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPFailureClientGetNoModelEmptyResponse{}, runtime.NewResponseError(resp)
	}
	return client.getNoModelEmptyHandleResponse(resp)
}

// getNoModelEmptyCreateRequest creates the GetNoModelEmpty request.
func (client *HTTPFailureClient) getNoModelEmptyCreateRequest(ctx context.Context, options *HTTPFailureClientGetNoModelEmptyOptions) (*policy.Request, error) {
	urlPath := "/http/failure/nomodel/empty"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getNoModelEmptyHandleResponse handles the GetNoModelEmpty response.
func (client *HTTPFailureClient) getNoModelEmptyHandleResponse(resp *http.Response) (HTTPFailureClientGetNoModelEmptyResponse, error) {
	result := HTTPFailureClientGetNoModelEmptyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return HTTPFailureClientGetNoModelEmptyResponse{}, err
	}
	return result, nil
}

// GetNoModelError - Get empty error form server
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - HTTPFailureClientGetNoModelErrorOptions contains the optional parameters for the HTTPFailureClient.GetNoModelError
// method.
func (client *HTTPFailureClient) GetNoModelError(ctx context.Context, options *HTTPFailureClientGetNoModelErrorOptions) (HTTPFailureClientGetNoModelErrorResponse, error) {
	req, err := client.getNoModelErrorCreateRequest(ctx, options)
	if err != nil {
		return HTTPFailureClientGetNoModelErrorResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPFailureClientGetNoModelErrorResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HTTPFailureClientGetNoModelErrorResponse{}, runtime.NewResponseError(resp)
	}
	return client.getNoModelErrorHandleResponse(resp)
}

// getNoModelErrorCreateRequest creates the GetNoModelError request.
func (client *HTTPFailureClient) getNoModelErrorCreateRequest(ctx context.Context, options *HTTPFailureClientGetNoModelErrorOptions) (*policy.Request, error) {
	urlPath := "/http/failure/nomodel/error"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getNoModelErrorHandleResponse handles the GetNoModelError response.
func (client *HTTPFailureClient) getNoModelErrorHandleResponse(resp *http.Response) (HTTPFailureClientGetNoModelErrorResponse, error) {
	result := HTTPFailureClientGetNoModelErrorResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return HTTPFailureClientGetNoModelErrorResponse{}, err
	}
	return result, nil
}
