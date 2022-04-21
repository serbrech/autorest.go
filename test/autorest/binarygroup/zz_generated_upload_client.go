//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package binarygroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"io"
	"net/http"
)

// UploadClient contains the methods for the Upload group.
// Don't use this type directly, use NewUploadClient() instead.
type UploadClient struct {
	pl runtime.Pipeline
}

// NewUploadClient creates a new instance of UploadClient with the specified values.
// options - pass nil to accept the default values.
func NewUploadClient(options *azcore.ClientOptions) *UploadClient {
	if options == nil {
		options = &azcore.ClientOptions{}
	}
	client := &UploadClient{
		pl: runtime.NewPipeline(moduleName, moduleVersion, runtime.PipelineOptions{}, options),
	}
	return client
}

// Binary - Uploading binary file
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// fileParam - Non-empty binary file
// options - UploadClientBinaryOptions contains the optional parameters for the UploadClient.Binary method.
func (client *UploadClient) Binary(ctx context.Context, fileParam io.ReadSeekCloser, options *UploadClientBinaryOptions) (UploadClientBinaryResponse, error) {
	req, err := client.binaryCreateRequest(ctx, fileParam, options)
	if err != nil {
		return UploadClientBinaryResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return UploadClientBinaryResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return UploadClientBinaryResponse{}, runtime.NewResponseError(resp)
	}
	return UploadClientBinaryResponse{}, nil
}

// binaryCreateRequest creates the Binary request.
func (client *UploadClient) binaryCreateRequest(ctx context.Context, fileParam io.ReadSeekCloser, options *UploadClientBinaryOptions) (*policy.Request, error) {
	urlPath := "/binary/octet"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, req.SetBody(fileParam, "application/octet-stream")
}

// File - Uploading json file
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// fileParam - JSON file with payload { "more": "cowbell" }
// options - UploadClientFileOptions contains the optional parameters for the UploadClient.File method.
func (client *UploadClient) File(ctx context.Context, fileParam io.ReadSeekCloser, options *UploadClientFileOptions) (UploadClientFileResponse, error) {
	req, err := client.fileCreateRequest(ctx, fileParam, options)
	if err != nil {
		return UploadClientFileResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return UploadClientFileResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return UploadClientFileResponse{}, runtime.NewResponseError(resp)
	}
	return UploadClientFileResponse{}, nil
}

// fileCreateRequest creates the File request.
func (client *UploadClient) fileCreateRequest(ctx context.Context, fileParam io.ReadSeekCloser, options *UploadClientFileOptions) (*policy.Request, error) {
	urlPath := "/binary/file"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, req.SetBody(fileParam, "application/json")
}
