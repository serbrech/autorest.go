// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// PolymorphismOperations contains the methods for the Polymorphism group.
type PolymorphismOperations interface {
	// GetComplicated - Get complex types that are polymorphic, but not at the root of the hierarchy; also have additional properties
	GetComplicated(ctx context.Context) (*SalmonResponse, error)
	// GetComposedWithDiscriminator - Get complex object composing a polymorphic scalar property and array property with polymorphic element type, with discriminator specified. Deserialization must NOT fail and use the discriminator type specified on the wire.
	GetComposedWithDiscriminator(ctx context.Context) (*DotFishMarketResponse, error)
	// GetComposedWithoutDiscriminator - Get complex object composing a polymorphic scalar property and array property with polymorphic element type, without discriminator specified on wire. Deserialization must NOT fail and use the explicit type of the property.
	GetComposedWithoutDiscriminator(ctx context.Context) (*DotFishMarketResponse, error)
	// GetDotSyntax - Get complex types that are polymorphic, JSON key contains a dot
	GetDotSyntax(ctx context.Context) (*DotFishResponse, error)
	// GetValid - Get complex types that are polymorphic
	GetValid(ctx context.Context) (*FishResponse, error)
	// PutComplicated - Put complex types that are polymorphic, but not at the root of the hierarchy; also have additional properties
	PutComplicated(ctx context.Context, complexBody SalmonClassification) (*http.Response, error)
	// PutMissingDiscriminator - Put complex types that are polymorphic, omitting the discriminator
	PutMissingDiscriminator(ctx context.Context, complexBody SalmonClassification) (*SalmonResponse, error)
	// PutValid - Put complex types that are polymorphic
	PutValid(ctx context.Context, complexBody FishClassification) (*http.Response, error)
	// PutValidMissingRequired - Put complex types that are polymorphic, attempting to omit required 'birthday' field - the request should not be allowed from the client
	PutValidMissingRequired(ctx context.Context, complexBody FishClassification) (*http.Response, error)
}

// PolymorphismClient implements the PolymorphismOperations interface.
// Don't use this type directly, use NewPolymorphismClient() instead.
type PolymorphismClient struct {
	*Client
}

// NewPolymorphismClient creates a new instance of PolymorphismClient with the specified values.
func NewPolymorphismClient(c *Client) PolymorphismOperations {
	return &PolymorphismClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *PolymorphismClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// GetComplicated - Get complex types that are polymorphic, but not at the root of the hierarchy; also have additional properties
func (client *PolymorphismClient) GetComplicated(ctx context.Context) (*SalmonResponse, error) {
	req, err := client.GetComplicatedCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetComplicatedHandleError(resp)
	}
	result, err := client.GetComplicatedHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetComplicatedCreateRequest creates the GetComplicated request.
func (client *PolymorphismClient) GetComplicatedCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/complex/polymorphism/complicated"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetComplicatedHandleResponse handles the GetComplicated response.
func (client *PolymorphismClient) GetComplicatedHandleResponse(resp *azcore.Response) (*SalmonResponse, error) {
	result := SalmonResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result)
}

// GetComplicatedHandleError handles the GetComplicated error response.
func (client *PolymorphismClient) GetComplicatedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetComposedWithDiscriminator - Get complex object composing a polymorphic scalar property and array property with polymorphic element type, with discriminator specified. Deserialization must NOT fail and use the discriminator type specified on the wire.
func (client *PolymorphismClient) GetComposedWithDiscriminator(ctx context.Context) (*DotFishMarketResponse, error) {
	req, err := client.GetComposedWithDiscriminatorCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetComposedWithDiscriminatorHandleError(resp)
	}
	result, err := client.GetComposedWithDiscriminatorHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetComposedWithDiscriminatorCreateRequest creates the GetComposedWithDiscriminator request.
func (client *PolymorphismClient) GetComposedWithDiscriminatorCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/complex/polymorphism/composedWithDiscriminator"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetComposedWithDiscriminatorHandleResponse handles the GetComposedWithDiscriminator response.
func (client *PolymorphismClient) GetComposedWithDiscriminatorHandleResponse(resp *azcore.Response) (*DotFishMarketResponse, error) {
	result := DotFishMarketResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DotFishMarket)
}

// GetComposedWithDiscriminatorHandleError handles the GetComposedWithDiscriminator error response.
func (client *PolymorphismClient) GetComposedWithDiscriminatorHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetComposedWithoutDiscriminator - Get complex object composing a polymorphic scalar property and array property with polymorphic element type, without discriminator specified on wire. Deserialization must NOT fail and use the explicit type of the property.
func (client *PolymorphismClient) GetComposedWithoutDiscriminator(ctx context.Context) (*DotFishMarketResponse, error) {
	req, err := client.GetComposedWithoutDiscriminatorCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetComposedWithoutDiscriminatorHandleError(resp)
	}
	result, err := client.GetComposedWithoutDiscriminatorHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetComposedWithoutDiscriminatorCreateRequest creates the GetComposedWithoutDiscriminator request.
func (client *PolymorphismClient) GetComposedWithoutDiscriminatorCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/complex/polymorphism/composedWithoutDiscriminator"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetComposedWithoutDiscriminatorHandleResponse handles the GetComposedWithoutDiscriminator response.
func (client *PolymorphismClient) GetComposedWithoutDiscriminatorHandleResponse(resp *azcore.Response) (*DotFishMarketResponse, error) {
	result := DotFishMarketResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DotFishMarket)
}

// GetComposedWithoutDiscriminatorHandleError handles the GetComposedWithoutDiscriminator error response.
func (client *PolymorphismClient) GetComposedWithoutDiscriminatorHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetDotSyntax - Get complex types that are polymorphic, JSON key contains a dot
func (client *PolymorphismClient) GetDotSyntax(ctx context.Context) (*DotFishResponse, error) {
	req, err := client.GetDotSyntaxCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetDotSyntaxHandleError(resp)
	}
	result, err := client.GetDotSyntaxHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDotSyntaxCreateRequest creates the GetDotSyntax request.
func (client *PolymorphismClient) GetDotSyntaxCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/complex/polymorphism/dotsyntax"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetDotSyntaxHandleResponse handles the GetDotSyntax response.
func (client *PolymorphismClient) GetDotSyntaxHandleResponse(resp *azcore.Response) (*DotFishResponse, error) {
	result := DotFishResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result)
}

// GetDotSyntaxHandleError handles the GetDotSyntax error response.
func (client *PolymorphismClient) GetDotSyntaxHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetValid - Get complex types that are polymorphic
func (client *PolymorphismClient) GetValid(ctx context.Context) (*FishResponse, error) {
	req, err := client.GetValidCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetValidHandleError(resp)
	}
	result, err := client.GetValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetValidCreateRequest creates the GetValid request.
func (client *PolymorphismClient) GetValidCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/complex/polymorphism/valid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetValidHandleResponse handles the GetValid response.
func (client *PolymorphismClient) GetValidHandleResponse(resp *azcore.Response) (*FishResponse, error) {
	result := FishResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result)
}

// GetValidHandleError handles the GetValid error response.
func (client *PolymorphismClient) GetValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutComplicated - Put complex types that are polymorphic, but not at the root of the hierarchy; also have additional properties
func (client *PolymorphismClient) PutComplicated(ctx context.Context, complexBody SalmonClassification) (*http.Response, error) {
	req, err := client.PutComplicatedCreateRequest(ctx, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutComplicatedHandleError(resp)
	}
	return resp.Response, nil
}

// PutComplicatedCreateRequest creates the PutComplicated request.
func (client *PolymorphismClient) PutComplicatedCreateRequest(ctx context.Context, complexBody SalmonClassification) (*azcore.Request, error) {
	urlPath := "/complex/polymorphism/complicated"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// PutComplicatedHandleError handles the PutComplicated error response.
func (client *PolymorphismClient) PutComplicatedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutMissingDiscriminator - Put complex types that are polymorphic, omitting the discriminator
func (client *PolymorphismClient) PutMissingDiscriminator(ctx context.Context, complexBody SalmonClassification) (*SalmonResponse, error) {
	req, err := client.PutMissingDiscriminatorCreateRequest(ctx, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutMissingDiscriminatorHandleError(resp)
	}
	result, err := client.PutMissingDiscriminatorHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PutMissingDiscriminatorCreateRequest creates the PutMissingDiscriminator request.
func (client *PolymorphismClient) PutMissingDiscriminatorCreateRequest(ctx context.Context, complexBody SalmonClassification) (*azcore.Request, error) {
	urlPath := "/complex/polymorphism/missingdiscriminator"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// PutMissingDiscriminatorHandleResponse handles the PutMissingDiscriminator response.
func (client *PolymorphismClient) PutMissingDiscriminatorHandleResponse(resp *azcore.Response) (*SalmonResponse, error) {
	result := SalmonResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result)
}

// PutMissingDiscriminatorHandleError handles the PutMissingDiscriminator error response.
func (client *PolymorphismClient) PutMissingDiscriminatorHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutValid - Put complex types that are polymorphic
func (client *PolymorphismClient) PutValid(ctx context.Context, complexBody FishClassification) (*http.Response, error) {
	req, err := client.PutValidCreateRequest(ctx, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutValidHandleError(resp)
	}
	return resp.Response, nil
}

// PutValidCreateRequest creates the PutValid request.
func (client *PolymorphismClient) PutValidCreateRequest(ctx context.Context, complexBody FishClassification) (*azcore.Request, error) {
	urlPath := "/complex/polymorphism/valid"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// PutValidHandleError handles the PutValid error response.
func (client *PolymorphismClient) PutValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutValidMissingRequired - Put complex types that are polymorphic, attempting to omit required 'birthday' field - the request should not be allowed from the client
func (client *PolymorphismClient) PutValidMissingRequired(ctx context.Context, complexBody FishClassification) (*http.Response, error) {
	req, err := client.PutValidMissingRequiredCreateRequest(ctx, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutValidMissingRequiredHandleError(resp)
	}
	return resp.Response, nil
}

// PutValidMissingRequiredCreateRequest creates the PutValidMissingRequired request.
func (client *PolymorphismClient) PutValidMissingRequiredCreateRequest(ctx context.Context, complexBody FishClassification) (*azcore.Request, error) {
	urlPath := "/complex/polymorphism/missingrequired/invalid"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// PutValidMissingRequiredHandleError handles the PutValidMissingRequired error response.
func (client *PolymorphismClient) PutValidMissingRequiredHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
