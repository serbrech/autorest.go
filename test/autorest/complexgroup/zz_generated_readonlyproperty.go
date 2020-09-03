// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
)

// ReadonlypropertyOperations contains the methods for the Readonlyproperty group.
type ReadonlypropertyOperations interface {
	// GetValid - Get complex types that have readonly properties
	GetValid(ctx context.Context) (*ReadonlyObjResponse, error)
	// PutValid - Put complex types that have readonly properties
	PutValid(ctx context.Context, complexBody ReadonlyObj) (*http.Response, error)
}

// ReadonlypropertyClient implements the ReadonlypropertyOperations interface.
// Don't use this type directly, use NewReadonlypropertyClient() instead.
type ReadonlypropertyClient struct {
	*Client
}

// NewReadonlypropertyClient creates a new instance of ReadonlypropertyClient with the specified values.
func NewReadonlypropertyClient(c *Client) ReadonlypropertyOperations {
	return &ReadonlypropertyClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *ReadonlypropertyClient) Do(ctx context.Context, req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(ctx, req)
}

// GetValid - Get complex types that have readonly properties
func (client *ReadonlypropertyClient) GetValid(ctx context.Context) (*ReadonlyObjResponse, error) {
	req, err := client.GetValidCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetValidCreateRequest creates the GetValid request.
func (client *ReadonlypropertyClient) GetValidCreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/complex/readonlyproperty/valid"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// GetValidHandleResponse handles the GetValid response.
func (client *ReadonlypropertyClient) GetValidHandleResponse(resp *azcore.Response) (*ReadonlyObjResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetValidHandleError(resp)
	}
	result := ReadonlyObjResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ReadonlyObj)
}

// GetValidHandleError handles the GetValid error response.
func (client *ReadonlypropertyClient) GetValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutValid - Put complex types that have readonly properties
func (client *ReadonlypropertyClient) PutValid(ctx context.Context, complexBody ReadonlyObj) (*http.Response, error) {
	req, err := client.PutValidCreateRequest(complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.PutValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PutValidCreateRequest creates the PutValid request.
func (client *ReadonlypropertyClient) PutValidCreateRequest(complexBody ReadonlyObj) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/complex/readonlyproperty/valid"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(complexBody)
}

// PutValidHandleResponse handles the PutValid response.
func (client *ReadonlypropertyClient) PutValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutValidHandleError(resp)
	}
	return resp.Response, nil
}

// PutValidHandleError handles the PutValid error response.
func (client *ReadonlypropertyClient) PutValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}