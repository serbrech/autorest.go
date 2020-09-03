// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package additionalpropsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
)

// PetsOperations contains the methods for the Pets group.
type PetsOperations interface {
	// CreateApInProperties - Create a Pet which contains more properties than what is defined.
	CreateApInProperties(ctx context.Context, createParameters PetApInProperties) (*PetApInPropertiesResponse, error)
	// CreateApInPropertiesWithApstring - Create a Pet which contains more properties than what is defined.
	CreateApInPropertiesWithApstring(ctx context.Context, createParameters PetApInPropertiesWithApstring) (*PetApInPropertiesWithApstringResponse, error)
	// CreateApObject - Create a Pet which contains more properties than what is defined.
	CreateApObject(ctx context.Context, createParameters PetApObject) (*PetApObjectResponse, error)
	// CreateApString - Create a Pet which contains more properties than what is defined.
	CreateApString(ctx context.Context, createParameters PetApString) (*PetApStringResponse, error)
	// CreateApTrue - Create a Pet which contains more properties than what is defined.
	CreateApTrue(ctx context.Context, createParameters PetApTrue) (*PetApTrueResponse, error)
	// CreateCatApTrue - Create a CatAPTrue which contains more properties than what is defined.
	CreateCatApTrue(ctx context.Context, createParameters CatApTrue) (*CatApTrueResponse, error)
}

// PetsClient implements the PetsOperations interface.
// Don't use this type directly, use NewPetsClient() instead.
type PetsClient struct {
	*Client
}

// NewPetsClient creates a new instance of PetsClient with the specified values.
func NewPetsClient(c *Client) PetsOperations {
	return &PetsClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *PetsClient) Do(ctx context.Context, req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(ctx, req)
}

// CreateApInProperties - Create a Pet which contains more properties than what is defined.
func (client *PetsClient) CreateApInProperties(ctx context.Context, createParameters PetApInProperties) (*PetApInPropertiesResponse, error) {
	req, err := client.CreateApInPropertiesCreateRequest(createParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.CreateApInPropertiesHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateApInPropertiesCreateRequest creates the CreateApInProperties request.
func (client *PetsClient) CreateApInPropertiesCreateRequest(createParameters PetApInProperties) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/additionalProperties/in/properties"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(createParameters)
}

// CreateApInPropertiesHandleResponse handles the CreateApInProperties response.
func (client *PetsClient) CreateApInPropertiesHandleResponse(resp *azcore.Response) (*PetApInPropertiesResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.CreateApInPropertiesHandleError(resp)
	}
	result := PetApInPropertiesResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PetApInProperties)
}

// CreateApInPropertiesHandleError handles the CreateApInProperties error response.
func (client *PetsClient) CreateApInPropertiesHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// CreateApInPropertiesWithApstring - Create a Pet which contains more properties than what is defined.
func (client *PetsClient) CreateApInPropertiesWithApstring(ctx context.Context, createParameters PetApInPropertiesWithApstring) (*PetApInPropertiesWithApstringResponse, error) {
	req, err := client.CreateApInPropertiesWithApstringCreateRequest(createParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.CreateApInPropertiesWithApstringHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateApInPropertiesWithApstringCreateRequest creates the CreateApInPropertiesWithApstring request.
func (client *PetsClient) CreateApInPropertiesWithApstringCreateRequest(createParameters PetApInPropertiesWithApstring) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/additionalProperties/in/properties/with/additionalProperties/string"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(createParameters)
}

// CreateApInPropertiesWithApstringHandleResponse handles the CreateApInPropertiesWithApstring response.
func (client *PetsClient) CreateApInPropertiesWithApstringHandleResponse(resp *azcore.Response) (*PetApInPropertiesWithApstringResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.CreateApInPropertiesWithApstringHandleError(resp)
	}
	result := PetApInPropertiesWithApstringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PetApInPropertiesWithApstring)
}

// CreateApInPropertiesWithApstringHandleError handles the CreateApInPropertiesWithApstring error response.
func (client *PetsClient) CreateApInPropertiesWithApstringHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// CreateApObject - Create a Pet which contains more properties than what is defined.
func (client *PetsClient) CreateApObject(ctx context.Context, createParameters PetApObject) (*PetApObjectResponse, error) {
	req, err := client.CreateApObjectCreateRequest(createParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.CreateApObjectHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateApObjectCreateRequest creates the CreateApObject request.
func (client *PetsClient) CreateApObjectCreateRequest(createParameters PetApObject) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/additionalProperties/type/object"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(createParameters)
}

// CreateApObjectHandleResponse handles the CreateApObject response.
func (client *PetsClient) CreateApObjectHandleResponse(resp *azcore.Response) (*PetApObjectResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.CreateApObjectHandleError(resp)
	}
	result := PetApObjectResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PetApObject)
}

// CreateApObjectHandleError handles the CreateApObject error response.
func (client *PetsClient) CreateApObjectHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// CreateApString - Create a Pet which contains more properties than what is defined.
func (client *PetsClient) CreateApString(ctx context.Context, createParameters PetApString) (*PetApStringResponse, error) {
	req, err := client.CreateApStringCreateRequest(createParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.CreateApStringHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateApStringCreateRequest creates the CreateApString request.
func (client *PetsClient) CreateApStringCreateRequest(createParameters PetApString) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/additionalProperties/type/string"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(createParameters)
}

// CreateApStringHandleResponse handles the CreateApString response.
func (client *PetsClient) CreateApStringHandleResponse(resp *azcore.Response) (*PetApStringResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.CreateApStringHandleError(resp)
	}
	result := PetApStringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PetApString)
}

// CreateApStringHandleError handles the CreateApString error response.
func (client *PetsClient) CreateApStringHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// CreateApTrue - Create a Pet which contains more properties than what is defined.
func (client *PetsClient) CreateApTrue(ctx context.Context, createParameters PetApTrue) (*PetApTrueResponse, error) {
	req, err := client.CreateApTrueCreateRequest(createParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.CreateApTrueHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateApTrueCreateRequest creates the CreateApTrue request.
func (client *PetsClient) CreateApTrueCreateRequest(createParameters PetApTrue) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/additionalProperties/true"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(createParameters)
}

// CreateApTrueHandleResponse handles the CreateApTrue response.
func (client *PetsClient) CreateApTrueHandleResponse(resp *azcore.Response) (*PetApTrueResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.CreateApTrueHandleError(resp)
	}
	result := PetApTrueResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PetApTrue)
}

// CreateApTrueHandleError handles the CreateApTrue error response.
func (client *PetsClient) CreateApTrueHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// CreateCatApTrue - Create a CatAPTrue which contains more properties than what is defined.
func (client *PetsClient) CreateCatApTrue(ctx context.Context, createParameters CatApTrue) (*CatApTrueResponse, error) {
	req, err := client.CreateCatApTrueCreateRequest(createParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.CreateCatApTrueHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateCatApTrueCreateRequest creates the CreateCatApTrue request.
func (client *PetsClient) CreateCatApTrueCreateRequest(createParameters CatApTrue) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/additionalProperties/true-subclass"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(createParameters)
}

// CreateCatApTrueHandleResponse handles the CreateCatApTrue response.
func (client *PetsClient) CreateCatApTrueHandleResponse(resp *azcore.Response) (*CatApTrueResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.CreateCatApTrueHandleError(resp)
	}
	result := CatApTrueResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.CatApTrue)
}

// CreateCatApTrueHandleError handles the CreateCatApTrue error response.
func (client *PetsClient) CreateCatApTrueHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}