//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package filegroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// FilesClient contains the methods for the Files group.
// Don't use this type directly, use NewFilesClient() instead.
type FilesClient struct {
	pl runtime.Pipeline
}

// NewFilesClient creates a new instance of FilesClient with the specified values.
// options - pass nil to accept the default values.
func NewFilesClient(options *azcore.ClientOptions) *FilesClient {
	if options == nil {
		options = &azcore.ClientOptions{}
	}
	client := &FilesClient{
		pl: runtime.NewPipeline(moduleName, moduleVersion, runtime.PipelineOptions{}, options),
	}
	return client
}

// GetEmptyFile - Get empty file
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - FilesClientGetEmptyFileOptions contains the optional parameters for the FilesClient.GetEmptyFile method.
func (client *FilesClient) GetEmptyFile(ctx context.Context, options *FilesClientGetEmptyFileOptions) (FilesClientGetEmptyFileResponse, error) {
	req, err := client.getEmptyFileCreateRequest(ctx, options)
	if err != nil {
		return FilesClientGetEmptyFileResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FilesClientGetEmptyFileResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FilesClientGetEmptyFileResponse{}, runtime.NewResponseError(resp)
	}
	return FilesClientGetEmptyFileResponse{Body: resp.Body}, nil
}

// getEmptyFileCreateRequest creates the GetEmptyFile request.
func (client *FilesClient) getEmptyFileCreateRequest(ctx context.Context, options *FilesClientGetEmptyFileOptions) (*policy.Request, error) {
	urlPath := "/files/stream/empty"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	runtime.SkipBodyDownload(req)
	req.Raw().Header.Set("Accept", "image/png, application/json")
	return req, nil
}

// GetFile - Get file
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - FilesClientGetFileOptions contains the optional parameters for the FilesClient.GetFile method.
func (client *FilesClient) GetFile(ctx context.Context, options *FilesClientGetFileOptions) (FilesClientGetFileResponse, error) {
	req, err := client.getFileCreateRequest(ctx, options)
	if err != nil {
		return FilesClientGetFileResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FilesClientGetFileResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FilesClientGetFileResponse{}, runtime.NewResponseError(resp)
	}
	return FilesClientGetFileResponse{Body: resp.Body}, nil
}

// getFileCreateRequest creates the GetFile request.
func (client *FilesClient) getFileCreateRequest(ctx context.Context, options *FilesClientGetFileOptions) (*policy.Request, error) {
	urlPath := "/files/stream/nonempty"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	runtime.SkipBodyDownload(req)
	req.Raw().Header.Set("Accept", "image/png, application/json")
	return req, nil
}

// GetFileLarge - Get a large file
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - FilesClientGetFileLargeOptions contains the optional parameters for the FilesClient.GetFileLarge method.
func (client *FilesClient) GetFileLarge(ctx context.Context, options *FilesClientGetFileLargeOptions) (FilesClientGetFileLargeResponse, error) {
	req, err := client.getFileLargeCreateRequest(ctx, options)
	if err != nil {
		return FilesClientGetFileLargeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FilesClientGetFileLargeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FilesClientGetFileLargeResponse{}, runtime.NewResponseError(resp)
	}
	return FilesClientGetFileLargeResponse{Body: resp.Body}, nil
}

// getFileLargeCreateRequest creates the GetFileLarge request.
func (client *FilesClient) getFileLargeCreateRequest(ctx context.Context, options *FilesClientGetFileLargeOptions) (*policy.Request, error) {
	urlPath := "/files/stream/verylarge"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	runtime.SkipBodyDownload(req)
	req.Raw().Header.Set("Accept", "image/png, application/json")
	return req, nil
}
