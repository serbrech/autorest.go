//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package formdatagroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"io"
	"net/http"
)

// FormdataClient contains the methods for the Formdata group.
// Don't use this type directly, use NewFormdataClient() instead.
type FormdataClient struct {
	pl runtime.Pipeline
}

// NewFormdataClient creates a new instance of FormdataClient with the specified values.
// options - pass nil to accept the default values.
func NewFormdataClient(options *azcore.ClientOptions) *FormdataClient {
	if options == nil {
		options = &azcore.ClientOptions{}
	}
	client := &FormdataClient{
		pl: runtime.NewPipeline(moduleName, moduleVersion, runtime.PipelineOptions{}, options),
	}
	return client
}

// UploadFile - Upload file
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// fileContent - File to upload.
// fileName - File name to upload. Name has to be spelled exactly as written here.
// options - FormdataClientUploadFileOptions contains the optional parameters for the FormdataClient.UploadFile method.
func (client *FormdataClient) UploadFile(ctx context.Context, fileContent io.ReadSeekCloser, fileName string, options *FormdataClientUploadFileOptions) (FormdataClientUploadFileResponse, error) {
	req, err := client.uploadFileCreateRequest(ctx, fileContent, fileName, options)
	if err != nil {
		return FormdataClientUploadFileResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FormdataClientUploadFileResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FormdataClientUploadFileResponse{}, runtime.NewResponseError(resp)
	}
	return FormdataClientUploadFileResponse{Body: resp.Body}, nil
}

// uploadFileCreateRequest creates the UploadFile request.
func (client *FormdataClient) uploadFileCreateRequest(ctx context.Context, fileContent io.ReadSeekCloser, fileName string, options *FormdataClientUploadFileOptions) (*policy.Request, error) {
	urlPath := "/formdata/stream/uploadfile"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	runtime.SkipBodyDownload(req)
	req.Raw().Header.Set("Accept", "application/octet-stream, application/json")
	if err := runtime.SetMultipartFormData(req, map[string]interface{}{
		"fileContent": fileContent,
		"fileName":    fileName,
	}); err != nil {
		return nil, err
	}
	return req, nil
}

// UploadFileViaBody - Upload file
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// fileContent - File to upload.
// options - FormdataClientUploadFileViaBodyOptions contains the optional parameters for the FormdataClient.UploadFileViaBody
// method.
func (client *FormdataClient) UploadFileViaBody(ctx context.Context, fileContent io.ReadSeekCloser, options *FormdataClientUploadFileViaBodyOptions) (FormdataClientUploadFileViaBodyResponse, error) {
	req, err := client.uploadFileViaBodyCreateRequest(ctx, fileContent, options)
	if err != nil {
		return FormdataClientUploadFileViaBodyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FormdataClientUploadFileViaBodyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FormdataClientUploadFileViaBodyResponse{}, runtime.NewResponseError(resp)
	}
	return FormdataClientUploadFileViaBodyResponse{Body: resp.Body}, nil
}

// uploadFileViaBodyCreateRequest creates the UploadFileViaBody request.
func (client *FormdataClient) uploadFileViaBodyCreateRequest(ctx context.Context, fileContent io.ReadSeekCloser, options *FormdataClientUploadFileViaBodyOptions) (*policy.Request, error) {
	urlPath := "/formdata/stream/uploadfile"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	runtime.SkipBodyDownload(req)
	req.Raw().Header.Set("Accept", "application/octet-stream, application/json")
	return req, req.SetBody(fileContent, "application/octet-stream")
}

// UploadFiles - Upload multiple files
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// fileContent - Files to upload.
// options - FormdataClientUploadFilesOptions contains the optional parameters for the FormdataClient.UploadFiles method.
func (client *FormdataClient) UploadFiles(ctx context.Context, fileContent []io.ReadSeekCloser, options *FormdataClientUploadFilesOptions) (FormdataClientUploadFilesResponse, error) {
	req, err := client.uploadFilesCreateRequest(ctx, fileContent, options)
	if err != nil {
		return FormdataClientUploadFilesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FormdataClientUploadFilesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FormdataClientUploadFilesResponse{}, runtime.NewResponseError(resp)
	}
	return FormdataClientUploadFilesResponse{Body: resp.Body}, nil
}

// uploadFilesCreateRequest creates the UploadFiles request.
func (client *FormdataClient) uploadFilesCreateRequest(ctx context.Context, fileContent []io.ReadSeekCloser, options *FormdataClientUploadFilesOptions) (*policy.Request, error) {
	urlPath := "/formdata/stream/uploadfiles"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	runtime.SkipBodyDownload(req)
	req.Raw().Header.Set("Accept", "application/octet-stream, application/json")
	if err := runtime.SetMultipartFormData(req, map[string]interface{}{
		"fileContent": fileContent,
	}); err != nil {
		return nil, err
	}
	return req, nil
}
