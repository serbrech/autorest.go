//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package optionalgroup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ImplicitClient contains the methods for the Implicit group.
// Don't use this type directly, use NewImplicitClient() instead.
type ImplicitClient struct {
	requiredGlobalPath  string
	requiredGlobalQuery string
	optionalGlobalQuery *int32
	pl                  runtime.Pipeline
}

// NewImplicitClient creates a new instance of ImplicitClient with the specified values.
// requiredGlobalPath - number of items to skip
// requiredGlobalQuery - number of items to skip
// optionalGlobalQuery - number of items to skip
// options - pass nil to accept the default values.
func NewImplicitClient(requiredGlobalPath string, requiredGlobalQuery string, optionalGlobalQuery *int32, options *azcore.ClientOptions) *ImplicitClient {
	if options == nil {
		options = &azcore.ClientOptions{}
	}
	client := &ImplicitClient{
		requiredGlobalPath:  requiredGlobalPath,
		requiredGlobalQuery: requiredGlobalQuery,
		optionalGlobalQuery: optionalGlobalQuery,
		pl:                  runtime.NewPipeline(moduleName, moduleVersion, runtime.PipelineOptions{}, options),
	}
	return client
}

// GetOptionalGlobalQuery - Test implicitly optional query parameter
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - ImplicitClientGetOptionalGlobalQueryOptions contains the optional parameters for the ImplicitClient.GetOptionalGlobalQuery
// method.
func (client *ImplicitClient) GetOptionalGlobalQuery(ctx context.Context, options *ImplicitClientGetOptionalGlobalQueryOptions) (ImplicitClientGetOptionalGlobalQueryResponse, error) {
	req, err := client.getOptionalGlobalQueryCreateRequest(ctx, options)
	if err != nil {
		return ImplicitClientGetOptionalGlobalQueryResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ImplicitClientGetOptionalGlobalQueryResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ImplicitClientGetOptionalGlobalQueryResponse{}, runtime.NewResponseError(resp)
	}
	return ImplicitClientGetOptionalGlobalQueryResponse{}, nil
}

// getOptionalGlobalQueryCreateRequest creates the GetOptionalGlobalQuery request.
func (client *ImplicitClient) getOptionalGlobalQueryCreateRequest(ctx context.Context, options *ImplicitClientGetOptionalGlobalQueryOptions) (*policy.Request, error) {
	urlPath := "/reqopt/global/optional/query"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if client.optionalGlobalQuery != nil {
		reqQP.Set("optional-global-query", strconv.FormatInt(int64(*client.optionalGlobalQuery), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// GetRequiredGlobalPath - Test implicitly required path parameter
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - ImplicitClientGetRequiredGlobalPathOptions contains the optional parameters for the ImplicitClient.GetRequiredGlobalPath
// method.
func (client *ImplicitClient) GetRequiredGlobalPath(ctx context.Context, options *ImplicitClientGetRequiredGlobalPathOptions) (ImplicitClientGetRequiredGlobalPathResponse, error) {
	req, err := client.getRequiredGlobalPathCreateRequest(ctx, options)
	if err != nil {
		return ImplicitClientGetRequiredGlobalPathResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ImplicitClientGetRequiredGlobalPathResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ImplicitClientGetRequiredGlobalPathResponse{}, runtime.NewResponseError(resp)
	}
	return ImplicitClientGetRequiredGlobalPathResponse{}, nil
}

// getRequiredGlobalPathCreateRequest creates the GetRequiredGlobalPath request.
func (client *ImplicitClient) getRequiredGlobalPathCreateRequest(ctx context.Context, options *ImplicitClientGetRequiredGlobalPathOptions) (*policy.Request, error) {
	urlPath := "/reqopt/global/required/path/{required-global-path}"
	if client.requiredGlobalPath == "" {
		return nil, errors.New("parameter client.requiredGlobalPath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{required-global-path}", url.PathEscape(client.requiredGlobalPath))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// GetRequiredGlobalQuery - Test implicitly required query parameter
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - ImplicitClientGetRequiredGlobalQueryOptions contains the optional parameters for the ImplicitClient.GetRequiredGlobalQuery
// method.
func (client *ImplicitClient) GetRequiredGlobalQuery(ctx context.Context, options *ImplicitClientGetRequiredGlobalQueryOptions) (ImplicitClientGetRequiredGlobalQueryResponse, error) {
	req, err := client.getRequiredGlobalQueryCreateRequest(ctx, options)
	if err != nil {
		return ImplicitClientGetRequiredGlobalQueryResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ImplicitClientGetRequiredGlobalQueryResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ImplicitClientGetRequiredGlobalQueryResponse{}, runtime.NewResponseError(resp)
	}
	return ImplicitClientGetRequiredGlobalQueryResponse{}, nil
}

// getRequiredGlobalQueryCreateRequest creates the GetRequiredGlobalQuery request.
func (client *ImplicitClient) getRequiredGlobalQueryCreateRequest(ctx context.Context, options *ImplicitClientGetRequiredGlobalQueryOptions) (*policy.Request, error) {
	urlPath := "/reqopt/global/required/query"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("required-global-query", client.requiredGlobalQuery)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// GetRequiredPath - Test implicitly required path parameter
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - ImplicitClientGetRequiredPathOptions contains the optional parameters for the ImplicitClient.GetRequiredPath
// method.
func (client *ImplicitClient) GetRequiredPath(ctx context.Context, pathParameter string, options *ImplicitClientGetRequiredPathOptions) (ImplicitClientGetRequiredPathResponse, error) {
	req, err := client.getRequiredPathCreateRequest(ctx, pathParameter, options)
	if err != nil {
		return ImplicitClientGetRequiredPathResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ImplicitClientGetRequiredPathResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ImplicitClientGetRequiredPathResponse{}, runtime.NewResponseError(resp)
	}
	return ImplicitClientGetRequiredPathResponse{}, nil
}

// getRequiredPathCreateRequest creates the GetRequiredPath request.
func (client *ImplicitClient) getRequiredPathCreateRequest(ctx context.Context, pathParameter string, options *ImplicitClientGetRequiredPathOptions) (*policy.Request, error) {
	urlPath := "/reqopt/implicit/required/path/{pathParameter}"
	if pathParameter == "" {
		return nil, errors.New("parameter pathParameter cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pathParameter}", url.PathEscape(pathParameter))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// PutOptionalBinaryBody - Test implicitly optional body parameter
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - ImplicitClientPutOptionalBinaryBodyOptions contains the optional parameters for the ImplicitClient.PutOptionalBinaryBody
// method.
func (client *ImplicitClient) PutOptionalBinaryBody(ctx context.Context, options *ImplicitClientPutOptionalBinaryBodyOptions) (ImplicitClientPutOptionalBinaryBodyResponse, error) {
	req, err := client.putOptionalBinaryBodyCreateRequest(ctx, options)
	if err != nil {
		return ImplicitClientPutOptionalBinaryBodyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ImplicitClientPutOptionalBinaryBodyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ImplicitClientPutOptionalBinaryBodyResponse{}, runtime.NewResponseError(resp)
	}
	return ImplicitClientPutOptionalBinaryBodyResponse{}, nil
}

// putOptionalBinaryBodyCreateRequest creates the PutOptionalBinaryBody request.
func (client *ImplicitClient) putOptionalBinaryBodyCreateRequest(ctx context.Context, options *ImplicitClientPutOptionalBinaryBodyOptions) (*policy.Request, error) {
	urlPath := "/reqopt/implicit/optional/binary-body"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.BodyParameter != nil {
		return req, req.SetBody(options.BodyParameter, "application/octet-stream")
	}
	return req, nil
}

// PutOptionalBody - Test implicitly optional body parameter
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - ImplicitClientPutOptionalBodyOptions contains the optional parameters for the ImplicitClient.PutOptionalBody
// method.
func (client *ImplicitClient) PutOptionalBody(ctx context.Context, options *ImplicitClientPutOptionalBodyOptions) (ImplicitClientPutOptionalBodyResponse, error) {
	req, err := client.putOptionalBodyCreateRequest(ctx, options)
	if err != nil {
		return ImplicitClientPutOptionalBodyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ImplicitClientPutOptionalBodyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ImplicitClientPutOptionalBodyResponse{}, runtime.NewResponseError(resp)
	}
	return ImplicitClientPutOptionalBodyResponse{}, nil
}

// putOptionalBodyCreateRequest creates the PutOptionalBody request.
func (client *ImplicitClient) putOptionalBodyCreateRequest(ctx context.Context, options *ImplicitClientPutOptionalBodyOptions) (*policy.Request, error) {
	urlPath := "/reqopt/implicit/optional/body"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.BodyParameter != nil {
		body := streaming.NopCloser(strings.NewReader(*options.BodyParameter))
		return req, req.SetBody(body, "application/json")
	}
	return req, nil
}

// PutOptionalHeader - Test implicitly optional header parameter
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - ImplicitClientPutOptionalHeaderOptions contains the optional parameters for the ImplicitClient.PutOptionalHeader
// method.
func (client *ImplicitClient) PutOptionalHeader(ctx context.Context, options *ImplicitClientPutOptionalHeaderOptions) (ImplicitClientPutOptionalHeaderResponse, error) {
	req, err := client.putOptionalHeaderCreateRequest(ctx, options)
	if err != nil {
		return ImplicitClientPutOptionalHeaderResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ImplicitClientPutOptionalHeaderResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ImplicitClientPutOptionalHeaderResponse{}, runtime.NewResponseError(resp)
	}
	return ImplicitClientPutOptionalHeaderResponse{}, nil
}

// putOptionalHeaderCreateRequest creates the PutOptionalHeader request.
func (client *ImplicitClient) putOptionalHeaderCreateRequest(ctx context.Context, options *ImplicitClientPutOptionalHeaderOptions) (*policy.Request, error) {
	urlPath := "/reqopt/implicit/optional/header"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	if options != nil && options.QueryParameter != nil {
		req.Raw().Header.Set("queryParameter", *options.QueryParameter)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// PutOptionalQuery - Test implicitly optional query parameter
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - ImplicitClientPutOptionalQueryOptions contains the optional parameters for the ImplicitClient.PutOptionalQuery
// method.
func (client *ImplicitClient) PutOptionalQuery(ctx context.Context, options *ImplicitClientPutOptionalQueryOptions) (ImplicitClientPutOptionalQueryResponse, error) {
	req, err := client.putOptionalQueryCreateRequest(ctx, options)
	if err != nil {
		return ImplicitClientPutOptionalQueryResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ImplicitClientPutOptionalQueryResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ImplicitClientPutOptionalQueryResponse{}, runtime.NewResponseError(resp)
	}
	return ImplicitClientPutOptionalQueryResponse{}, nil
}

// putOptionalQueryCreateRequest creates the PutOptionalQuery request.
func (client *ImplicitClient) putOptionalQueryCreateRequest(ctx context.Context, options *ImplicitClientPutOptionalQueryOptions) (*policy.Request, error) {
	urlPath := "/reqopt/implicit/optional/query"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.QueryParameter != nil {
		reqQP.Set("queryParameter", *options.QueryParameter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}
