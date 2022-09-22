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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// DataFlowDebugSessionClient contains the methods for the DataFlowDebugSession group.
// Don't use this type directly, use NewDataFlowDebugSessionClient() instead.
type DataFlowDebugSessionClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// NewDataFlowDebugSessionClient creates a new instance of DataFlowDebugSessionClient with the specified values.
// endpoint - The workspace development endpoint, for example https://myworkspace.dev.azuresynapse.net.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDataFlowDebugSessionClient(endpoint string, credential azcore.TokenCredential, options *DataFlowDebugSessionClientOptions) *DataFlowDebugSessionClient {
	if options == nil {
		options = &DataFlowDebugSessionClientOptions{}
	}
	authPolicy := runtime.NewBearerTokenPolicy(credential, []string{"https://dev.azuresynapse.net/.default"}, nil)
	pl := runtime.NewPipeline(moduleName, moduleVersion, runtime.PipelineOptions{PerRetry: []policy.Policy{authPolicy}}, &options.ClientOptions)
	client := &DataFlowDebugSessionClient{
		endpoint: endpoint,
		pl:       pl,
	}
	return client
}

// AddDataFlow - Add a data flow into debug session.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// request - Data flow debug session definition with debug content.
// options - DataFlowDebugSessionClientAddDataFlowOptions contains the optional parameters for the DataFlowDebugSessionClient.AddDataFlow
// method.
func (client *DataFlowDebugSessionClient) AddDataFlow(ctx context.Context, request DataFlowDebugPackage, options *DataFlowDebugSessionClientAddDataFlowOptions) (DataFlowDebugSessionClientAddDataFlowResponse, error) {
	req, err := client.addDataFlowCreateRequest(ctx, request, options)
	if err != nil {
		return DataFlowDebugSessionClientAddDataFlowResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataFlowDebugSessionClientAddDataFlowResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataFlowDebugSessionClientAddDataFlowResponse{}, runtime.NewResponseError(resp)
	}
	return client.addDataFlowHandleResponse(resp)
}

// addDataFlowCreateRequest creates the AddDataFlow request.
func (client *DataFlowDebugSessionClient) addDataFlowCreateRequest(ctx context.Context, request DataFlowDebugPackage, options *DataFlowDebugSessionClientAddDataFlowOptions) (*policy.Request, error) {
	urlPath := "/addDataFlowToDebugSession"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, request)
}

// addDataFlowHandleResponse handles the AddDataFlow response.
func (client *DataFlowDebugSessionClient) addDataFlowHandleResponse(resp *http.Response) (DataFlowDebugSessionClientAddDataFlowResponse, error) {
	result := DataFlowDebugSessionClientAddDataFlowResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AddDataFlowToDebugSessionResponse); err != nil {
		return DataFlowDebugSessionClientAddDataFlowResponse{}, err
	}
	return result, nil
}

// BeginCreateDataFlowDebugSession - Creates a data flow debug session.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// request - Data flow debug session definition
// options - DataFlowDebugSessionClientBeginCreateDataFlowDebugSessionOptions contains the optional parameters for the DataFlowDebugSessionClient.BeginCreateDataFlowDebugSession
// method.
func (client *DataFlowDebugSessionClient) BeginCreateDataFlowDebugSession(ctx context.Context, request CreateDataFlowDebugSessionRequest, options *DataFlowDebugSessionClientBeginCreateDataFlowDebugSessionOptions) (*runtime.Poller[DataFlowDebugSessionClientCreateDataFlowDebugSessionResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createDataFlowDebugSession(ctx, request, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DataFlowDebugSessionClientCreateDataFlowDebugSessionResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DataFlowDebugSessionClientCreateDataFlowDebugSessionResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateDataFlowDebugSession - Creates a data flow debug session.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
func (client *DataFlowDebugSessionClient) createDataFlowDebugSession(ctx context.Context, request CreateDataFlowDebugSessionRequest, options *DataFlowDebugSessionClientBeginCreateDataFlowDebugSessionOptions) (*http.Response, error) {
	req, err := client.createDataFlowDebugSessionCreateRequest(ctx, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createDataFlowDebugSessionCreateRequest creates the CreateDataFlowDebugSession request.
func (client *DataFlowDebugSessionClient) createDataFlowDebugSessionCreateRequest(ctx context.Context, request CreateDataFlowDebugSessionRequest, options *DataFlowDebugSessionClientBeginCreateDataFlowDebugSessionOptions) (*policy.Request, error) {
	urlPath := "/createDataFlowDebugSession"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, request)
}

// DeleteDataFlowDebugSession - Deletes a data flow debug session.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// request - Data flow debug session definition for deletion
// options - DataFlowDebugSessionClientDeleteDataFlowDebugSessionOptions contains the optional parameters for the DataFlowDebugSessionClient.DeleteDataFlowDebugSession
// method.
func (client *DataFlowDebugSessionClient) DeleteDataFlowDebugSession(ctx context.Context, request DeleteDataFlowDebugSessionRequest, options *DataFlowDebugSessionClientDeleteDataFlowDebugSessionOptions) (DataFlowDebugSessionClientDeleteDataFlowDebugSessionResponse, error) {
	req, err := client.deleteDataFlowDebugSessionCreateRequest(ctx, request, options)
	if err != nil {
		return DataFlowDebugSessionClientDeleteDataFlowDebugSessionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataFlowDebugSessionClientDeleteDataFlowDebugSessionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataFlowDebugSessionClientDeleteDataFlowDebugSessionResponse{}, runtime.NewResponseError(resp)
	}
	return DataFlowDebugSessionClientDeleteDataFlowDebugSessionResponse{}, nil
}

// deleteDataFlowDebugSessionCreateRequest creates the DeleteDataFlowDebugSession request.
func (client *DataFlowDebugSessionClient) deleteDataFlowDebugSessionCreateRequest(ctx context.Context, request DeleteDataFlowDebugSessionRequest, options *DataFlowDebugSessionClientDeleteDataFlowDebugSessionOptions) (*policy.Request, error) {
	urlPath := "/deleteDataFlowDebugSession"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, request)
}

// BeginExecuteCommand - Execute a data flow debug command.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// request - Data flow debug command definition.
// options - DataFlowDebugSessionClientBeginExecuteCommandOptions contains the optional parameters for the DataFlowDebugSessionClient.BeginExecuteCommand
// method.
func (client *DataFlowDebugSessionClient) BeginExecuteCommand(ctx context.Context, request DataFlowDebugCommandRequest, options *DataFlowDebugSessionClientBeginExecuteCommandOptions) (*runtime.Poller[DataFlowDebugSessionClientExecuteCommandResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.executeCommand(ctx, request, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DataFlowDebugSessionClientExecuteCommandResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DataFlowDebugSessionClientExecuteCommandResponse](options.ResumeToken, client.pl, nil)
	}
}

// ExecuteCommand - Execute a data flow debug command.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
func (client *DataFlowDebugSessionClient) executeCommand(ctx context.Context, request DataFlowDebugCommandRequest, options *DataFlowDebugSessionClientBeginExecuteCommandOptions) (*http.Response, error) {
	req, err := client.executeCommandCreateRequest(ctx, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// executeCommandCreateRequest creates the ExecuteCommand request.
func (client *DataFlowDebugSessionClient) executeCommandCreateRequest(ctx context.Context, request DataFlowDebugCommandRequest, options *DataFlowDebugSessionClientBeginExecuteCommandOptions) (*policy.Request, error) {
	urlPath := "/executeDataFlowDebugCommand"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, request)
}

// NewQueryDataFlowDebugSessionsByWorkspacePager - Query all active data flow debug sessions.
// Generated from API version 2019-06-01-preview
// options - DataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspaceOptions contains the optional parameters for the
// DataFlowDebugSessionClient.QueryDataFlowDebugSessionsByWorkspace method.
func (client *DataFlowDebugSessionClient) NewQueryDataFlowDebugSessionsByWorkspacePager(options *DataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspaceOptions) *runtime.Pager[DataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspaceResponse]{
		More: func(page DataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspaceResponse) (DataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.queryDataFlowDebugSessionsByWorkspaceCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspaceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.queryDataFlowDebugSessionsByWorkspaceHandleResponse(resp)
		},
	})
}

// queryDataFlowDebugSessionsByWorkspaceCreateRequest creates the QueryDataFlowDebugSessionsByWorkspace request.
func (client *DataFlowDebugSessionClient) queryDataFlowDebugSessionsByWorkspaceCreateRequest(ctx context.Context, options *DataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/queryDataFlowDebugSessions"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// queryDataFlowDebugSessionsByWorkspaceHandleResponse handles the QueryDataFlowDebugSessionsByWorkspace response.
func (client *DataFlowDebugSessionClient) queryDataFlowDebugSessionsByWorkspaceHandleResponse(resp *http.Response) (DataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspaceResponse, error) {
	result := DataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.QueryDataFlowDebugSessionsResponse); err != nil {
		return DataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspaceResponse{}, err
	}
	return result, nil
}
