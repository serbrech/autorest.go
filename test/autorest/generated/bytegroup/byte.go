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

// byteOperations implements the ByteOperations interface.
type byteOperations struct {
	*Client
}

// GetEmpty - Get empty byte value ''
func (client *byteOperations) GetEmpty(ctx context.Context) (*ByteArrayResponse, error) {
	req, err := client.getEmptyCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getEmptyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getEmptyCreateRequest creates the GetEmpty request.
func (client *byteOperations) getEmptyCreateRequest() (*azcore.Request, error) {
	urlPath := "/byte/empty"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getEmptyHandleResponse handles the GetEmpty response.
func (client *byteOperations) getEmptyHandleResponse(resp *azcore.Response) (*ByteArrayResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := ByteArrayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetInvalid - Get invalid byte value ':::SWAGGER::::'
func (client *byteOperations) GetInvalid(ctx context.Context) (*ByteArrayResponse, error) {
	req, err := client.getInvalidCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getInvalidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getInvalidCreateRequest creates the GetInvalid request.
func (client *byteOperations) getInvalidCreateRequest() (*azcore.Request, error) {
	urlPath := "/byte/invalid"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getInvalidHandleResponse handles the GetInvalid response.
func (client *byteOperations) getInvalidHandleResponse(resp *azcore.Response) (*ByteArrayResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := ByteArrayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetNonASCII - Get non-ascii byte string hex(FF FE FD FC FB FA F9 F8 F7 F6)
func (client *byteOperations) GetNonASCII(ctx context.Context) (*ByteArrayResponse, error) {
	req, err := client.getNonAsciiCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getNonAsciiHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getNonAsciiCreateRequest creates the GetNonASCII request.
func (client *byteOperations) getNonAsciiCreateRequest() (*azcore.Request, error) {
	urlPath := "/byte/nonAscii"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getNonAsciiHandleResponse handles the GetNonASCII response.
func (client *byteOperations) getNonAsciiHandleResponse(resp *azcore.Response) (*ByteArrayResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := ByteArrayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetNull - Get null byte value
func (client *byteOperations) GetNull(ctx context.Context) (*ByteArrayResponse, error) {
	req, err := client.getNullCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getNullCreateRequest creates the GetNull request.
func (client *byteOperations) getNullCreateRequest() (*azcore.Request, error) {
	urlPath := "/byte/null"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *byteOperations) getNullHandleResponse(resp *azcore.Response) (*ByteArrayResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := ByteArrayResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// PutNonASCII - Put non-ascii byte string hex(FF FE FD FC FB FA F9 F8 F7 F6)
func (client *byteOperations) PutNonASCII(ctx context.Context, byteBody []byte) (*http.Response, error) {
	req, err := client.putNonAsciiCreateRequest(byteBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putNonAsciiHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putNonAsciiCreateRequest creates the PutNonASCII request.
func (client *byteOperations) putNonAsciiCreateRequest(byteBody []byte) (*azcore.Request, error) {
	urlPath := "/byte/nonAscii"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	if err := req.MarshalAsJSON(byteBody); err != nil {
		if err != nil {
			return nil, err
		}
	}
	return req, nil
}

// putNonAsciiHandleResponse handles the PutNonASCII response.
func (client *byteOperations) putNonAsciiHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return resp.Response, nil
}
