//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azartifacts

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

type workspaceGitRepoManagementClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// newWorkspaceGitRepoManagementClient creates a new instance of workspaceGitRepoManagementClient with the specified values.
//   - endpoint - The workspace development endpoint, for example https://myworkspace.dev.azuresynapse.net.
//   - pl - the pipeline used for sending requests and handling responses.
func newWorkspaceGitRepoManagementClient(endpoint string, pl runtime.Pipeline) *workspaceGitRepoManagementClient {
	client := &workspaceGitRepoManagementClient{
		endpoint: endpoint,
		pl:       pl,
	}
	return client
}

// GetGitHubAccessToken - Get the GitHub access token.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - options - WorkspaceGitRepoManagementClientGetGitHubAccessTokenOptions contains the optional parameters for the workspaceGitRepoManagementClient.GetGitHubAccessToken
//     method.
func (client *workspaceGitRepoManagementClient) GetGitHubAccessToken(ctx context.Context, gitHubAccessTokenRequest GitHubAccessTokenRequest, options *WorkspaceGitRepoManagementClientGetGitHubAccessTokenOptions) (WorkspaceGitRepoManagementClientGetGitHubAccessTokenResponse, error) {
	req, err := client.getGitHubAccessTokenCreateRequest(ctx, gitHubAccessTokenRequest, options)
	if err != nil {
		return WorkspaceGitRepoManagementClientGetGitHubAccessTokenResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceGitRepoManagementClientGetGitHubAccessTokenResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceGitRepoManagementClientGetGitHubAccessTokenResponse{}, runtime.NewResponseError(resp)
	}
	return client.getGitHubAccessTokenHandleResponse(resp)
}

// getGitHubAccessTokenCreateRequest creates the GetGitHubAccessToken request.
func (client *workspaceGitRepoManagementClient) getGitHubAccessTokenCreateRequest(ctx context.Context, gitHubAccessTokenRequest GitHubAccessTokenRequest, options *WorkspaceGitRepoManagementClientGetGitHubAccessTokenOptions) (*policy.Request, error) {
	urlPath := "/getGitHubAccessToken"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.ClientRequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*options.ClientRequestID}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, gitHubAccessTokenRequest)
}

// getGitHubAccessTokenHandleResponse handles the GetGitHubAccessToken response.
func (client *workspaceGitRepoManagementClient) getGitHubAccessTokenHandleResponse(resp *http.Response) (WorkspaceGitRepoManagementClientGetGitHubAccessTokenResponse, error) {
	result := WorkspaceGitRepoManagementClientGetGitHubAccessTokenResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GitHubAccessTokenResponse); err != nil {
		return WorkspaceGitRepoManagementClientGetGitHubAccessTokenResponse{}, err
	}
	return result, nil
}
