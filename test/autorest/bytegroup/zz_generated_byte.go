// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package bytegroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// ByteOperations contains the methods for the Byte group.
type ByteOperations interface {
	// GetEmpty - Get empty byte value ''
	GetEmpty(ctx context.Context) (*ByteArrayResponse, error)
	// GetInvalid - Get invalid byte value ':::SWAGGER::::'
	GetInvalid(ctx context.Context) (*ByteArrayResponse, error)
	// GetNonASCII - Get non-ascii byte string hex(FF FE FD FC FB FA F9 F8 F7 F6)
	GetNonASCII(ctx context.Context) (*ByteArrayResponse, error)
	// GetNull - Get null byte value
	GetNull(ctx context.Context) (*ByteArrayResponse, error)
	// PutNonASCII - Put non-ascii byte string hex(FF FE FD FC FB FA F9 F8 F7 F6)
	PutNonASCII(ctx context.Context, byteBody []byte) (*http.Response, error)
}

// ByteClient implements the ByteOperations interface.
// Don't use this type directly, use NewByteClient() instead.
type ByteClient struct {
	*Client
}

// NewByteClient creates a new instance of ByteClient with the specified values.
func NewByteClient(c *Client) ByteOperations {
	return &ByteClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *ByteClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// GetEmpty - Get empty byte value ''
func (client *ByteClient) GetEmpty(ctx context.Context) (*ByteArrayResponse, error) {
	req, err := client.GetEmptyCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetEmptyHandleError(resp)
	}
	result, err := client.GetEmptyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetEmptyCreateRequest creates the GetEmpty request.
func (client *ByteClient) GetEmptyCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/byte/empty"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetEmptyHandleResponse handles the GetEmpty response.
func (client *ByteClient) GetEmptyHandleResponse(resp *azcore.Response) (*ByteArrayResponse, error) {
	result := ByteArrayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsByteArray(&result.Value, azcore.Base64StdFormat)
}

// GetEmptyHandleError handles the GetEmpty error response.
func (client *ByteClient) GetEmptyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetInvalid - Get invalid byte value ':::SWAGGER::::'
func (client *ByteClient) GetInvalid(ctx context.Context) (*ByteArrayResponse, error) {
	req, err := client.GetInvalidCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetInvalidHandleError(resp)
	}
	result, err := client.GetInvalidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetInvalidCreateRequest creates the GetInvalid request.
func (client *ByteClient) GetInvalidCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/byte/invalid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetInvalidHandleResponse handles the GetInvalid response.
func (client *ByteClient) GetInvalidHandleResponse(resp *azcore.Response) (*ByteArrayResponse, error) {
	result := ByteArrayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsByteArray(&result.Value, azcore.Base64StdFormat)
}

// GetInvalidHandleError handles the GetInvalid error response.
func (client *ByteClient) GetInvalidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNonASCII - Get non-ascii byte string hex(FF FE FD FC FB FA F9 F8 F7 F6)
func (client *ByteClient) GetNonASCII(ctx context.Context) (*ByteArrayResponse, error) {
	req, err := client.GetNonASCIICreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetNonASCIIHandleError(resp)
	}
	result, err := client.GetNonASCIIHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetNonASCIICreateRequest creates the GetNonASCII request.
func (client *ByteClient) GetNonASCIICreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/byte/nonAscii"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetNonASCIIHandleResponse handles the GetNonASCII response.
func (client *ByteClient) GetNonASCIIHandleResponse(resp *azcore.Response) (*ByteArrayResponse, error) {
	result := ByteArrayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsByteArray(&result.Value, azcore.Base64StdFormat)
}

// GetNonASCIIHandleError handles the GetNonASCII error response.
func (client *ByteClient) GetNonASCIIHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNull - Get null byte value
func (client *ByteClient) GetNull(ctx context.Context) (*ByteArrayResponse, error) {
	req, err := client.GetNullCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetNullHandleError(resp)
	}
	result, err := client.GetNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetNullCreateRequest creates the GetNull request.
func (client *ByteClient) GetNullCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/byte/null"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetNullHandleResponse handles the GetNull response.
func (client *ByteClient) GetNullHandleResponse(resp *azcore.Response) (*ByteArrayResponse, error) {
	result := ByteArrayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsByteArray(&result.Value, azcore.Base64StdFormat)
}

// GetNullHandleError handles the GetNull error response.
func (client *ByteClient) GetNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutNonASCII - Put non-ascii byte string hex(FF FE FD FC FB FA F9 F8 F7 F6)
func (client *ByteClient) PutNonASCII(ctx context.Context, byteBody []byte) (*http.Response, error) {
	req, err := client.PutNonASCIICreateRequest(ctx, byteBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutNonASCIIHandleError(resp)
	}
	return resp.Response, nil
}

// PutNonASCIICreateRequest creates the PutNonASCII request.
func (client *ByteClient) PutNonASCIICreateRequest(ctx context.Context, byteBody []byte) (*azcore.Request, error) {
	urlPath := "/byte/nonAscii"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsByteArray(byteBody, azcore.Base64StdFormat)
}

// PutNonASCIIHandleError handles the PutNonASCII error response.
func (client *ByteClient) PutNonASCIIHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
