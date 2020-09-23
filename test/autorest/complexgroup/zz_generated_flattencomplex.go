// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
)

// FlattencomplexOperations contains the methods for the Flattencomplex group.
type FlattencomplexOperations interface {
	GetValid(ctx context.Context) (*MyBaseTypeResponse, error)
}

// FlattencomplexClient implements the FlattencomplexOperations interface.
// Don't use this type directly, use NewFlattencomplexClient() instead.
type FlattencomplexClient struct {
	*Client
}

// NewFlattencomplexClient creates a new instance of FlattencomplexClient with the specified values.
func NewFlattencomplexClient(c *Client) FlattencomplexOperations {
	return &FlattencomplexClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *FlattencomplexClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

func (client *FlattencomplexClient) GetValid(ctx context.Context) (*MyBaseTypeResponse, error) {
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
func (client *FlattencomplexClient) GetValidCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/complex/flatten/valid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetValidHandleResponse handles the GetValid response.
func (client *FlattencomplexClient) GetValidHandleResponse(resp *azcore.Response) (*MyBaseTypeResponse, error) {
	result := MyBaseTypeResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result)
}

// GetValidHandleError handles the GetValid error response.
func (client *FlattencomplexClient) GetValidHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
