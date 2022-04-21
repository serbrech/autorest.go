//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type pipelineRunClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// newPipelineRunClient creates a new instance of pipelineRunClient with the specified values.
// endpoint - The workspace development endpoint, for example https://myworkspace.dev.azuresynapse.net.
// pl - the pipeline used for sending requests and handling responses.
func newPipelineRunClient(endpoint string, pl runtime.Pipeline) *pipelineRunClient {
	client := &pipelineRunClient{
		endpoint: endpoint,
		pl:       pl,
	}
	return client
}

// CancelPipelineRun - Cancel a pipeline run by its run ID.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// runID - The pipeline run identifier.
// options - pipelineRunClientCancelPipelineRunOptions contains the optional parameters for the pipelineRunClient.CancelPipelineRun
// method.
func (client *pipelineRunClient) CancelPipelineRun(ctx context.Context, runID string, options *pipelineRunClientCancelPipelineRunOptions) (pipelineRunClientCancelPipelineRunResponse, error) {
	req, err := client.cancelPipelineRunCreateRequest(ctx, runID, options)
	if err != nil {
		return pipelineRunClientCancelPipelineRunResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return pipelineRunClientCancelPipelineRunResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return pipelineRunClientCancelPipelineRunResponse{}, runtime.NewResponseError(resp)
	}
	return pipelineRunClientCancelPipelineRunResponse{}, nil
}

// cancelPipelineRunCreateRequest creates the CancelPipelineRun request.
func (client *pipelineRunClient) cancelPipelineRunCreateRequest(ctx context.Context, runID string, options *pipelineRunClientCancelPipelineRunOptions) (*policy.Request, error) {
	urlPath := "/pipelineruns/{runId}/cancel"
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.IsRecursive != nil {
		reqQP.Set("isRecursive", strconv.FormatBool(*options.IsRecursive))
	}
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// GetPipelineRun - Get a pipeline run by its run ID.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// runID - The pipeline run identifier.
// options - pipelineRunClientGetPipelineRunOptions contains the optional parameters for the pipelineRunClient.GetPipelineRun
// method.
func (client *pipelineRunClient) GetPipelineRun(ctx context.Context, runID string, options *pipelineRunClientGetPipelineRunOptions) (pipelineRunClientGetPipelineRunResponse, error) {
	req, err := client.getPipelineRunCreateRequest(ctx, runID, options)
	if err != nil {
		return pipelineRunClientGetPipelineRunResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return pipelineRunClientGetPipelineRunResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return pipelineRunClientGetPipelineRunResponse{}, runtime.NewResponseError(resp)
	}
	return client.getPipelineRunHandleResponse(resp)
}

// getPipelineRunCreateRequest creates the GetPipelineRun request.
func (client *pipelineRunClient) getPipelineRunCreateRequest(ctx context.Context, runID string, options *pipelineRunClientGetPipelineRunOptions) (*policy.Request, error) {
	urlPath := "/pipelineruns/{runId}"
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getPipelineRunHandleResponse handles the GetPipelineRun response.
func (client *pipelineRunClient) getPipelineRunHandleResponse(resp *http.Response) (pipelineRunClientGetPipelineRunResponse, error) {
	result := pipelineRunClientGetPipelineRunResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineRun); err != nil {
		return pipelineRunClientGetPipelineRunResponse{}, err
	}
	return result, nil
}

// QueryActivityRuns - Query activity runs based on input filter conditions.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// pipelineName - The pipeline name.
// runID - The pipeline run identifier.
// filterParameters - Parameters to filter the activity runs.
// options - pipelineRunClientQueryActivityRunsOptions contains the optional parameters for the pipelineRunClient.QueryActivityRuns
// method.
func (client *pipelineRunClient) QueryActivityRuns(ctx context.Context, pipelineName string, runID string, filterParameters RunFilterParameters, options *pipelineRunClientQueryActivityRunsOptions) (pipelineRunClientQueryActivityRunsResponse, error) {
	req, err := client.queryActivityRunsCreateRequest(ctx, pipelineName, runID, filterParameters, options)
	if err != nil {
		return pipelineRunClientQueryActivityRunsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return pipelineRunClientQueryActivityRunsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return pipelineRunClientQueryActivityRunsResponse{}, runtime.NewResponseError(resp)
	}
	return client.queryActivityRunsHandleResponse(resp)
}

// queryActivityRunsCreateRequest creates the QueryActivityRuns request.
func (client *pipelineRunClient) queryActivityRunsCreateRequest(ctx context.Context, pipelineName string, runID string, filterParameters RunFilterParameters, options *pipelineRunClientQueryActivityRunsOptions) (*policy.Request, error) {
	urlPath := "/pipelines/{pipelineName}/pipelineruns/{runId}/queryActivityruns"
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, filterParameters)
}

// queryActivityRunsHandleResponse handles the QueryActivityRuns response.
func (client *pipelineRunClient) queryActivityRunsHandleResponse(resp *http.Response) (pipelineRunClientQueryActivityRunsResponse, error) {
	result := pipelineRunClientQueryActivityRunsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ActivityRunsQueryResponse); err != nil {
		return pipelineRunClientQueryActivityRunsResponse{}, err
	}
	return result, nil
}

// QueryPipelineRunsByWorkspace - Query pipeline runs in the workspace based on input filter conditions.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// filterParameters - Parameters to filter the pipeline run.
// options - pipelineRunClientQueryPipelineRunsByWorkspaceOptions contains the optional parameters for the pipelineRunClient.QueryPipelineRunsByWorkspace
// method.
func (client *pipelineRunClient) QueryPipelineRunsByWorkspace(ctx context.Context, filterParameters RunFilterParameters, options *pipelineRunClientQueryPipelineRunsByWorkspaceOptions) (pipelineRunClientQueryPipelineRunsByWorkspaceResponse, error) {
	req, err := client.queryPipelineRunsByWorkspaceCreateRequest(ctx, filterParameters, options)
	if err != nil {
		return pipelineRunClientQueryPipelineRunsByWorkspaceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return pipelineRunClientQueryPipelineRunsByWorkspaceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return pipelineRunClientQueryPipelineRunsByWorkspaceResponse{}, runtime.NewResponseError(resp)
	}
	return client.queryPipelineRunsByWorkspaceHandleResponse(resp)
}

// queryPipelineRunsByWorkspaceCreateRequest creates the QueryPipelineRunsByWorkspace request.
func (client *pipelineRunClient) queryPipelineRunsByWorkspaceCreateRequest(ctx context.Context, filterParameters RunFilterParameters, options *pipelineRunClientQueryPipelineRunsByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/queryPipelineRuns"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, filterParameters)
}

// queryPipelineRunsByWorkspaceHandleResponse handles the QueryPipelineRunsByWorkspace response.
func (client *pipelineRunClient) queryPipelineRunsByWorkspaceHandleResponse(resp *http.Response) (pipelineRunClientQueryPipelineRunsByWorkspaceResponse, error) {
	result := pipelineRunClientQueryPipelineRunsByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineRunsQueryResponse); err != nil {
		return pipelineRunClientQueryPipelineRunsByWorkspaceResponse{}, err
	}
	return result, nil
}
