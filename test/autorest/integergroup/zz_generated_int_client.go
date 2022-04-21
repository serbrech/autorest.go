//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package integergroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"time"
)

// IntClient contains the methods for the Int group.
// Don't use this type directly, use NewIntClient() instead.
type IntClient struct {
	pl runtime.Pipeline
}

// NewIntClient creates a new instance of IntClient with the specified values.
// options - pass nil to accept the default values.
func NewIntClient(options *azcore.ClientOptions) *IntClient {
	if options == nil {
		options = &azcore.ClientOptions{}
	}
	client := &IntClient{
		pl: runtime.NewPipeline(moduleName, moduleVersion, runtime.PipelineOptions{}, options),
	}
	return client
}

// GetInvalid - Get invalid Int value
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - IntClientGetInvalidOptions contains the optional parameters for the IntClient.GetInvalid method.
func (client *IntClient) GetInvalid(ctx context.Context, options *IntClientGetInvalidOptions) (IntClientGetInvalidResponse, error) {
	req, err := client.getInvalidCreateRequest(ctx, options)
	if err != nil {
		return IntClientGetInvalidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntClientGetInvalidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntClientGetInvalidResponse{}, runtime.NewResponseError(resp)
	}
	return client.getInvalidHandleResponse(resp)
}

// getInvalidCreateRequest creates the GetInvalid request.
func (client *IntClient) getInvalidCreateRequest(ctx context.Context, options *IntClientGetInvalidOptions) (*policy.Request, error) {
	urlPath := "/int/invalid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getInvalidHandleResponse handles the GetInvalid response.
func (client *IntClient) getInvalidHandleResponse(resp *http.Response) (IntClientGetInvalidResponse, error) {
	result := IntClientGetInvalidResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return IntClientGetInvalidResponse{}, err
	}
	return result, nil
}

// GetInvalidUnixTime - Get invalid Unix time value
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - IntClientGetInvalidUnixTimeOptions contains the optional parameters for the IntClient.GetInvalidUnixTime method.
func (client *IntClient) GetInvalidUnixTime(ctx context.Context, options *IntClientGetInvalidUnixTimeOptions) (IntClientGetInvalidUnixTimeResponse, error) {
	req, err := client.getInvalidUnixTimeCreateRequest(ctx, options)
	if err != nil {
		return IntClientGetInvalidUnixTimeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntClientGetInvalidUnixTimeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntClientGetInvalidUnixTimeResponse{}, runtime.NewResponseError(resp)
	}
	return client.getInvalidUnixTimeHandleResponse(resp)
}

// getInvalidUnixTimeCreateRequest creates the GetInvalidUnixTime request.
func (client *IntClient) getInvalidUnixTimeCreateRequest(ctx context.Context, options *IntClientGetInvalidUnixTimeOptions) (*policy.Request, error) {
	urlPath := "/int/invalidunixtime"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getInvalidUnixTimeHandleResponse handles the GetInvalidUnixTime response.
func (client *IntClient) getInvalidUnixTimeHandleResponse(resp *http.Response) (IntClientGetInvalidUnixTimeResponse, error) {
	result := IntClientGetInvalidUnixTimeResponse{}
	var aux *timeUnix
	if err := runtime.UnmarshalAsJSON(resp, &aux); err != nil {
		return IntClientGetInvalidUnixTimeResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// GetNull - Get null Int value
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - IntClientGetNullOptions contains the optional parameters for the IntClient.GetNull method.
func (client *IntClient) GetNull(ctx context.Context, options *IntClientGetNullOptions) (IntClientGetNullResponse, error) {
	req, err := client.getNullCreateRequest(ctx, options)
	if err != nil {
		return IntClientGetNullResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntClientGetNullResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntClientGetNullResponse{}, runtime.NewResponseError(resp)
	}
	return client.getNullHandleResponse(resp)
}

// getNullCreateRequest creates the GetNull request.
func (client *IntClient) getNullCreateRequest(ctx context.Context, options *IntClientGetNullOptions) (*policy.Request, error) {
	urlPath := "/int/null"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *IntClient) getNullHandleResponse(resp *http.Response) (IntClientGetNullResponse, error) {
	result := IntClientGetNullResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return IntClientGetNullResponse{}, err
	}
	return result, nil
}

// GetNullUnixTime - Get null Unix time value
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - IntClientGetNullUnixTimeOptions contains the optional parameters for the IntClient.GetNullUnixTime method.
func (client *IntClient) GetNullUnixTime(ctx context.Context, options *IntClientGetNullUnixTimeOptions) (IntClientGetNullUnixTimeResponse, error) {
	req, err := client.getNullUnixTimeCreateRequest(ctx, options)
	if err != nil {
		return IntClientGetNullUnixTimeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntClientGetNullUnixTimeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntClientGetNullUnixTimeResponse{}, runtime.NewResponseError(resp)
	}
	return client.getNullUnixTimeHandleResponse(resp)
}

// getNullUnixTimeCreateRequest creates the GetNullUnixTime request.
func (client *IntClient) getNullUnixTimeCreateRequest(ctx context.Context, options *IntClientGetNullUnixTimeOptions) (*policy.Request, error) {
	urlPath := "/int/nullunixtime"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getNullUnixTimeHandleResponse handles the GetNullUnixTime response.
func (client *IntClient) getNullUnixTimeHandleResponse(resp *http.Response) (IntClientGetNullUnixTimeResponse, error) {
	result := IntClientGetNullUnixTimeResponse{}
	var aux *timeUnix
	if err := runtime.UnmarshalAsJSON(resp, &aux); err != nil {
		return IntClientGetNullUnixTimeResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// GetOverflowInt32 - Get overflow Int32 value
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - IntClientGetOverflowInt32Options contains the optional parameters for the IntClient.GetOverflowInt32 method.
func (client *IntClient) GetOverflowInt32(ctx context.Context, options *IntClientGetOverflowInt32Options) (IntClientGetOverflowInt32Response, error) {
	req, err := client.getOverflowInt32CreateRequest(ctx, options)
	if err != nil {
		return IntClientGetOverflowInt32Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntClientGetOverflowInt32Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntClientGetOverflowInt32Response{}, runtime.NewResponseError(resp)
	}
	return client.getOverflowInt32HandleResponse(resp)
}

// getOverflowInt32CreateRequest creates the GetOverflowInt32 request.
func (client *IntClient) getOverflowInt32CreateRequest(ctx context.Context, options *IntClientGetOverflowInt32Options) (*policy.Request, error) {
	urlPath := "/int/overflowint32"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getOverflowInt32HandleResponse handles the GetOverflowInt32 response.
func (client *IntClient) getOverflowInt32HandleResponse(resp *http.Response) (IntClientGetOverflowInt32Response, error) {
	result := IntClientGetOverflowInt32Response{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return IntClientGetOverflowInt32Response{}, err
	}
	return result, nil
}

// GetOverflowInt64 - Get overflow Int64 value
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - IntClientGetOverflowInt64Options contains the optional parameters for the IntClient.GetOverflowInt64 method.
func (client *IntClient) GetOverflowInt64(ctx context.Context, options *IntClientGetOverflowInt64Options) (IntClientGetOverflowInt64Response, error) {
	req, err := client.getOverflowInt64CreateRequest(ctx, options)
	if err != nil {
		return IntClientGetOverflowInt64Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntClientGetOverflowInt64Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntClientGetOverflowInt64Response{}, runtime.NewResponseError(resp)
	}
	return client.getOverflowInt64HandleResponse(resp)
}

// getOverflowInt64CreateRequest creates the GetOverflowInt64 request.
func (client *IntClient) getOverflowInt64CreateRequest(ctx context.Context, options *IntClientGetOverflowInt64Options) (*policy.Request, error) {
	urlPath := "/int/overflowint64"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getOverflowInt64HandleResponse handles the GetOverflowInt64 response.
func (client *IntClient) getOverflowInt64HandleResponse(resp *http.Response) (IntClientGetOverflowInt64Response, error) {
	result := IntClientGetOverflowInt64Response{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return IntClientGetOverflowInt64Response{}, err
	}
	return result, nil
}

// GetUnderflowInt32 - Get underflow Int32 value
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - IntClientGetUnderflowInt32Options contains the optional parameters for the IntClient.GetUnderflowInt32 method.
func (client *IntClient) GetUnderflowInt32(ctx context.Context, options *IntClientGetUnderflowInt32Options) (IntClientGetUnderflowInt32Response, error) {
	req, err := client.getUnderflowInt32CreateRequest(ctx, options)
	if err != nil {
		return IntClientGetUnderflowInt32Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntClientGetUnderflowInt32Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntClientGetUnderflowInt32Response{}, runtime.NewResponseError(resp)
	}
	return client.getUnderflowInt32HandleResponse(resp)
}

// getUnderflowInt32CreateRequest creates the GetUnderflowInt32 request.
func (client *IntClient) getUnderflowInt32CreateRequest(ctx context.Context, options *IntClientGetUnderflowInt32Options) (*policy.Request, error) {
	urlPath := "/int/underflowint32"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getUnderflowInt32HandleResponse handles the GetUnderflowInt32 response.
func (client *IntClient) getUnderflowInt32HandleResponse(resp *http.Response) (IntClientGetUnderflowInt32Response, error) {
	result := IntClientGetUnderflowInt32Response{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return IntClientGetUnderflowInt32Response{}, err
	}
	return result, nil
}

// GetUnderflowInt64 - Get underflow Int64 value
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - IntClientGetUnderflowInt64Options contains the optional parameters for the IntClient.GetUnderflowInt64 method.
func (client *IntClient) GetUnderflowInt64(ctx context.Context, options *IntClientGetUnderflowInt64Options) (IntClientGetUnderflowInt64Response, error) {
	req, err := client.getUnderflowInt64CreateRequest(ctx, options)
	if err != nil {
		return IntClientGetUnderflowInt64Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntClientGetUnderflowInt64Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntClientGetUnderflowInt64Response{}, runtime.NewResponseError(resp)
	}
	return client.getUnderflowInt64HandleResponse(resp)
}

// getUnderflowInt64CreateRequest creates the GetUnderflowInt64 request.
func (client *IntClient) getUnderflowInt64CreateRequest(ctx context.Context, options *IntClientGetUnderflowInt64Options) (*policy.Request, error) {
	urlPath := "/int/underflowint64"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getUnderflowInt64HandleResponse handles the GetUnderflowInt64 response.
func (client *IntClient) getUnderflowInt64HandleResponse(resp *http.Response) (IntClientGetUnderflowInt64Response, error) {
	result := IntClientGetUnderflowInt64Response{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return IntClientGetUnderflowInt64Response{}, err
	}
	return result, nil
}

// GetUnixTime - Get datetime encoded as Unix time value
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// options - IntClientGetUnixTimeOptions contains the optional parameters for the IntClient.GetUnixTime method.
func (client *IntClient) GetUnixTime(ctx context.Context, options *IntClientGetUnixTimeOptions) (IntClientGetUnixTimeResponse, error) {
	req, err := client.getUnixTimeCreateRequest(ctx, options)
	if err != nil {
		return IntClientGetUnixTimeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntClientGetUnixTimeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntClientGetUnixTimeResponse{}, runtime.NewResponseError(resp)
	}
	return client.getUnixTimeHandleResponse(resp)
}

// getUnixTimeCreateRequest creates the GetUnixTime request.
func (client *IntClient) getUnixTimeCreateRequest(ctx context.Context, options *IntClientGetUnixTimeOptions) (*policy.Request, error) {
	urlPath := "/int/unixtime"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getUnixTimeHandleResponse handles the GetUnixTime response.
func (client *IntClient) getUnixTimeHandleResponse(resp *http.Response) (IntClientGetUnixTimeResponse, error) {
	result := IntClientGetUnixTimeResponse{}
	var aux *timeUnix
	if err := runtime.UnmarshalAsJSON(resp, &aux); err != nil {
		return IntClientGetUnixTimeResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// PutMax32 - Put max int32 value
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// intBody - int body
// options - IntClientPutMax32Options contains the optional parameters for the IntClient.PutMax32 method.
func (client *IntClient) PutMax32(ctx context.Context, intBody int32, options *IntClientPutMax32Options) (IntClientPutMax32Response, error) {
	req, err := client.putMax32CreateRequest(ctx, intBody, options)
	if err != nil {
		return IntClientPutMax32Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntClientPutMax32Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntClientPutMax32Response{}, runtime.NewResponseError(resp)
	}
	return IntClientPutMax32Response{}, nil
}

// putMax32CreateRequest creates the PutMax32 request.
func (client *IntClient) putMax32CreateRequest(ctx context.Context, intBody int32, options *IntClientPutMax32Options) (*policy.Request, error) {
	urlPath := "/int/max/32"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, intBody)
}

// PutMax64 - Put max int64 value
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// intBody - int body
// options - IntClientPutMax64Options contains the optional parameters for the IntClient.PutMax64 method.
func (client *IntClient) PutMax64(ctx context.Context, intBody int64, options *IntClientPutMax64Options) (IntClientPutMax64Response, error) {
	req, err := client.putMax64CreateRequest(ctx, intBody, options)
	if err != nil {
		return IntClientPutMax64Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntClientPutMax64Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntClientPutMax64Response{}, runtime.NewResponseError(resp)
	}
	return IntClientPutMax64Response{}, nil
}

// putMax64CreateRequest creates the PutMax64 request.
func (client *IntClient) putMax64CreateRequest(ctx context.Context, intBody int64, options *IntClientPutMax64Options) (*policy.Request, error) {
	urlPath := "/int/max/64"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, intBody)
}

// PutMin32 - Put min int32 value
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// intBody - int body
// options - IntClientPutMin32Options contains the optional parameters for the IntClient.PutMin32 method.
func (client *IntClient) PutMin32(ctx context.Context, intBody int32, options *IntClientPutMin32Options) (IntClientPutMin32Response, error) {
	req, err := client.putMin32CreateRequest(ctx, intBody, options)
	if err != nil {
		return IntClientPutMin32Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntClientPutMin32Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntClientPutMin32Response{}, runtime.NewResponseError(resp)
	}
	return IntClientPutMin32Response{}, nil
}

// putMin32CreateRequest creates the PutMin32 request.
func (client *IntClient) putMin32CreateRequest(ctx context.Context, intBody int32, options *IntClientPutMin32Options) (*policy.Request, error) {
	urlPath := "/int/min/32"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, intBody)
}

// PutMin64 - Put min int64 value
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// intBody - int body
// options - IntClientPutMin64Options contains the optional parameters for the IntClient.PutMin64 method.
func (client *IntClient) PutMin64(ctx context.Context, intBody int64, options *IntClientPutMin64Options) (IntClientPutMin64Response, error) {
	req, err := client.putMin64CreateRequest(ctx, intBody, options)
	if err != nil {
		return IntClientPutMin64Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntClientPutMin64Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntClientPutMin64Response{}, runtime.NewResponseError(resp)
	}
	return IntClientPutMin64Response{}, nil
}

// putMin64CreateRequest creates the PutMin64 request.
func (client *IntClient) putMin64CreateRequest(ctx context.Context, intBody int64, options *IntClientPutMin64Options) (*policy.Request, error) {
	urlPath := "/int/min/64"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, intBody)
}

// PutUnixTimeDate - Put datetime encoded as Unix time
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 1.0.0
// intBody - int body
// options - IntClientPutUnixTimeDateOptions contains the optional parameters for the IntClient.PutUnixTimeDate method.
func (client *IntClient) PutUnixTimeDate(ctx context.Context, intBody time.Time, options *IntClientPutUnixTimeDateOptions) (IntClientPutUnixTimeDateResponse, error) {
	req, err := client.putUnixTimeDateCreateRequest(ctx, intBody, options)
	if err != nil {
		return IntClientPutUnixTimeDateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntClientPutUnixTimeDateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntClientPutUnixTimeDateResponse{}, runtime.NewResponseError(resp)
	}
	return IntClientPutUnixTimeDateResponse{}, nil
}

// putUnixTimeDateCreateRequest creates the PutUnixTimeDate request.
func (client *IntClient) putUnixTimeDateCreateRequest(ctx context.Context, intBody time.Time, options *IntClientPutUnixTimeDateOptions) (*policy.Request, error) {
	urlPath := "/int/unixtime"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	aux := timeUnix(intBody)
	return req, runtime.MarshalAsJSON(req, aux)
}
