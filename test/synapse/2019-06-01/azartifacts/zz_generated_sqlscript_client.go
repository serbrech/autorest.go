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
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

type sqlScriptClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// newSQLScriptClient creates a new instance of sqlScriptClient with the specified values.
//   - endpoint - The workspace development endpoint, for example https://myworkspace.dev.azuresynapse.net.
//   - pl - the pipeline used for sending requests and handling responses.
func newSQLScriptClient(endpoint string, pl runtime.Pipeline) *sqlScriptClient {
	client := &sqlScriptClient{
		endpoint: endpoint,
		pl:       pl,
	}
	return client
}

// BeginCreateOrUpdateSQLScript - Creates or updates a Sql Script.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - sqlScriptName - The sql script name.
//   - sqlScript - Sql Script resource definition.
//   - options - SqlScriptClientBeginCreateOrUpdateSQLScriptOptions contains the optional parameters for the sqlScriptClient.BeginCreateOrUpdateSQLScript
//     method.
func (client *sqlScriptClient) BeginCreateOrUpdateSQLScript(ctx context.Context, sqlScriptName string, sqlScript SQLScriptResource, options *SqlScriptClientBeginCreateOrUpdateSQLScriptOptions) (*runtime.Poller[SqlScriptClientCreateOrUpdateSQLScriptResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdateSQLScript(ctx, sqlScriptName, sqlScript, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SqlScriptClientCreateOrUpdateSQLScriptResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[SqlScriptClientCreateOrUpdateSQLScriptResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdateSQLScript - Creates or updates a Sql Script.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *sqlScriptClient) createOrUpdateSQLScript(ctx context.Context, sqlScriptName string, sqlScript SQLScriptResource, options *SqlScriptClientBeginCreateOrUpdateSQLScriptOptions) (*http.Response, error) {
	req, err := client.createOrUpdateSQLScriptCreateRequest(ctx, sqlScriptName, sqlScript, options)
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

// createOrUpdateSQLScriptCreateRequest creates the CreateOrUpdateSQLScript request.
func (client *sqlScriptClient) createOrUpdateSQLScriptCreateRequest(ctx context.Context, sqlScriptName string, sqlScript SQLScriptResource, options *SqlScriptClientBeginCreateOrUpdateSQLScriptOptions) (*policy.Request, error) {
	urlPath := "/sqlScripts/{sqlScriptName}"
	if sqlScriptName == "" {
		return nil, errors.New("parameter sqlScriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlScriptName}", url.PathEscape(sqlScriptName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, sqlScript)
}

// BeginDeleteSQLScript - Deletes a Sql Script.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - sqlScriptName - The sql script name.
//   - options - SqlScriptClientBeginDeleteSQLScriptOptions contains the optional parameters for the sqlScriptClient.BeginDeleteSQLScript
//     method.
func (client *sqlScriptClient) BeginDeleteSQLScript(ctx context.Context, sqlScriptName string, options *SqlScriptClientBeginDeleteSQLScriptOptions) (*runtime.Poller[SqlScriptClientDeleteSQLScriptResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteSQLScript(ctx, sqlScriptName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SqlScriptClientDeleteSQLScriptResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[SqlScriptClientDeleteSQLScriptResponse](options.ResumeToken, client.pl, nil)
	}
}

// DeleteSQLScript - Deletes a Sql Script.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *sqlScriptClient) deleteSQLScript(ctx context.Context, sqlScriptName string, options *SqlScriptClientBeginDeleteSQLScriptOptions) (*http.Response, error) {
	req, err := client.deleteSQLScriptCreateRequest(ctx, sqlScriptName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteSQLScriptCreateRequest creates the DeleteSQLScript request.
func (client *sqlScriptClient) deleteSQLScriptCreateRequest(ctx context.Context, sqlScriptName string, options *SqlScriptClientBeginDeleteSQLScriptOptions) (*policy.Request, error) {
	urlPath := "/sqlScripts/{sqlScriptName}"
	if sqlScriptName == "" {
		return nil, errors.New("parameter sqlScriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlScriptName}", url.PathEscape(sqlScriptName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetSQLScript - Gets a sql script.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - sqlScriptName - The sql script name.
//   - options - SqlScriptClientGetSQLScriptOptions contains the optional parameters for the sqlScriptClient.GetSQLScript method.
func (client *sqlScriptClient) GetSQLScript(ctx context.Context, sqlScriptName string, options *SqlScriptClientGetSQLScriptOptions) (SqlScriptClientGetSQLScriptResponse, error) {
	req, err := client.getSQLScriptCreateRequest(ctx, sqlScriptName, options)
	if err != nil {
		return SqlScriptClientGetSQLScriptResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SqlScriptClientGetSQLScriptResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotModified) {
		return SqlScriptClientGetSQLScriptResponse{}, runtime.NewResponseError(resp)
	}
	return client.getSQLScriptHandleResponse(resp)
}

// getSQLScriptCreateRequest creates the GetSQLScript request.
func (client *sqlScriptClient) getSQLScriptCreateRequest(ctx context.Context, sqlScriptName string, options *SqlScriptClientGetSQLScriptOptions) (*policy.Request, error) {
	urlPath := "/sqlScripts/{sqlScriptName}"
	if sqlScriptName == "" {
		return nil, errors.New("parameter sqlScriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlScriptName}", url.PathEscape(sqlScriptName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSQLScriptHandleResponse handles the GetSQLScript response.
func (client *sqlScriptClient) getSQLScriptHandleResponse(resp *http.Response) (SqlScriptClientGetSQLScriptResponse, error) {
	result := SqlScriptClientGetSQLScriptResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLScriptResource); err != nil {
		return SqlScriptClientGetSQLScriptResponse{}, err
	}
	return result, nil
}

// NewGetSQLScriptsByWorkspacePager - Lists sql scripts.
//
// Generated from API version 2019-06-01-preview
//   - options - SqlScriptClientGetSQLScriptsByWorkspaceOptions contains the optional parameters for the sqlScriptClient.GetSQLScriptsByWorkspace
//     method.
func (client *sqlScriptClient) NewGetSQLScriptsByWorkspacePager(options *SqlScriptClientGetSQLScriptsByWorkspaceOptions) *runtime.Pager[SqlScriptClientGetSQLScriptsByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[SqlScriptClientGetSQLScriptsByWorkspaceResponse]{
		More: func(page SqlScriptClientGetSQLScriptsByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SqlScriptClientGetSQLScriptsByWorkspaceResponse) (SqlScriptClientGetSQLScriptsByWorkspaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.getSQLScriptsByWorkspaceCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SqlScriptClientGetSQLScriptsByWorkspaceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SqlScriptClientGetSQLScriptsByWorkspaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SqlScriptClientGetSQLScriptsByWorkspaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.getSQLScriptsByWorkspaceHandleResponse(resp)
		},
	})
}

// getSQLScriptsByWorkspaceCreateRequest creates the GetSQLScriptsByWorkspace request.
func (client *sqlScriptClient) getSQLScriptsByWorkspaceCreateRequest(ctx context.Context, options *SqlScriptClientGetSQLScriptsByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/sqlScripts"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSQLScriptsByWorkspaceHandleResponse handles the GetSQLScriptsByWorkspace response.
func (client *sqlScriptClient) getSQLScriptsByWorkspaceHandleResponse(resp *http.Response) (SqlScriptClientGetSQLScriptsByWorkspaceResponse, error) {
	result := SqlScriptClientGetSQLScriptsByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLScriptsListResponse); err != nil {
		return SqlScriptClientGetSQLScriptsByWorkspaceResponse{}, err
	}
	return result, nil
}

// BeginRenameSQLScript - Renames a sqlScript.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - sqlScriptName - The sql script name.
//   - request - proposed new name.
//   - options - SqlScriptClientBeginRenameSQLScriptOptions contains the optional parameters for the sqlScriptClient.BeginRenameSQLScript
//     method.
func (client *sqlScriptClient) BeginRenameSQLScript(ctx context.Context, sqlScriptName string, request ArtifactRenameRequest, options *SqlScriptClientBeginRenameSQLScriptOptions) (*runtime.Poller[SqlScriptClientRenameSQLScriptResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.renameSQLScript(ctx, sqlScriptName, request, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SqlScriptClientRenameSQLScriptResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[SqlScriptClientRenameSQLScriptResponse](options.ResumeToken, client.pl, nil)
	}
}

// RenameSQLScript - Renames a sqlScript.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *sqlScriptClient) renameSQLScript(ctx context.Context, sqlScriptName string, request ArtifactRenameRequest, options *SqlScriptClientBeginRenameSQLScriptOptions) (*http.Response, error) {
	req, err := client.renameSQLScriptCreateRequest(ctx, sqlScriptName, request, options)
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

// renameSQLScriptCreateRequest creates the RenameSQLScript request.
func (client *sqlScriptClient) renameSQLScriptCreateRequest(ctx context.Context, sqlScriptName string, request ArtifactRenameRequest, options *SqlScriptClientBeginRenameSQLScriptOptions) (*policy.Request, error) {
	urlPath := "/sqlScripts/{sqlScriptName}/rename"
	if sqlScriptName == "" {
		return nil, errors.New("parameter sqlScriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlScriptName}", url.PathEscape(sqlScriptName))
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
