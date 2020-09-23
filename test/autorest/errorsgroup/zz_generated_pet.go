// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package errorsgroup

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// PetOperations contains the methods for the Pet group.
type PetOperations interface {
	// DoSomething - Asks pet to do something
	DoSomething(ctx context.Context, whatAction string) (*PetActionResponse, error)
	// GetPetByID - Gets pets by id.
	GetPetByID(ctx context.Context, petId string) (*PetResponse, error)
}

// PetClient implements the PetOperations interface.
// Don't use this type directly, use NewPetClient() instead.
type PetClient struct {
	*Client
}

// NewPetClient creates a new instance of PetClient with the specified values.
func NewPetClient(c *Client) PetOperations {
	return &PetClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *PetClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// DoSomething - Asks pet to do something
func (client *PetClient) DoSomething(ctx context.Context, whatAction string) (*PetActionResponse, error) {
	req, err := client.DoSomethingCreateRequest(ctx, whatAction)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.DoSomethingHandleError(resp)
	}
	result, err := client.DoSomethingHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DoSomethingCreateRequest creates the DoSomething request.
func (client *PetClient) DoSomethingCreateRequest(ctx context.Context, whatAction string) (*azcore.Request, error) {
	urlPath := "/errorStatusCodes/Pets/doSomething/{whatAction}"
	urlPath = strings.ReplaceAll(urlPath, "{whatAction}", url.PathEscape(whatAction))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DoSomethingHandleResponse handles the DoSomething response.
func (client *PetClient) DoSomethingHandleResponse(resp *azcore.Response) (*PetActionResponse, error) {
	result := PetActionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PetAction)
}

// DoSomethingHandleError handles the DoSomething error response.
func (client *PetClient) DoSomethingHandleError(resp *azcore.Response) error {
	switch resp.StatusCode {
	case http.StatusInternalServerError:
		var err petActionError
		if err := resp.UnmarshalAsJSON(&err); err != nil {
			return err
		}
		return azcore.NewResponseError(err.wrapped, resp.Response)
	default:
		var err petActionError
		if err := resp.UnmarshalAsJSON(&err); err != nil {
			return err
		}
		return azcore.NewResponseError(err.wrapped, resp.Response)
	}
}

// GetPetByID - Gets pets by id.
func (client *PetClient) GetPetByID(ctx context.Context, petId string) (*PetResponse, error) {
	req, err := client.GetPetByIDCreateRequest(ctx, petId)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.GetPetByIDHandleError(resp)
	}
	result, err := client.GetPetByIDHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPetByIDCreateRequest creates the GetPetByID request.
func (client *PetClient) GetPetByIDCreateRequest(ctx context.Context, petId string) (*azcore.Request, error) {
	urlPath := "/errorStatusCodes/Pets/{petId}/GetPet"
	urlPath = strings.ReplaceAll(urlPath, "{petId}", url.PathEscape(petId))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetPetByIDHandleResponse handles the GetPetByID response.
func (client *PetClient) GetPetByIDHandleResponse(resp *azcore.Response) (*PetResponse, error) {
	result := PetResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Pet)
}

// GetPetByIDHandleError handles the GetPetByID error response.
func (client *PetClient) GetPetByIDHandleError(resp *azcore.Response) error {
	switch resp.StatusCode {
	case http.StatusBadRequest:
		var err string
		if err := resp.UnmarshalAsJSON(&err); err != nil {
			return err
		}
		return azcore.NewResponseError(fmt.Errorf("%v", err), resp.Response)
	case http.StatusNotFound:
		var err notFoundErrorBase
		if err := resp.UnmarshalAsJSON(&err); err != nil {
			return err
		}
		return azcore.NewResponseError(err.wrapped, resp.Response)
	case http.StatusNotImplemented:
		var err int32
		if err := resp.UnmarshalAsJSON(&err); err != nil {
			return err
		}
		return azcore.NewResponseError(fmt.Errorf("%v", err), resp.Response)
	default:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
		}
		if len(body) == 0 {
			return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
		}
		return azcore.NewResponseError(errors.New(string(body)), resp.Response)
	}
}
