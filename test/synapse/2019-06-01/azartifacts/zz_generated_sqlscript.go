// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

type sqlScriptClient struct {
	*client
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *sqlScriptClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// CreateOrUpdateSQLScript - Creates or updates a Sql Script.
func (client *sqlScriptClient) CreateOrUpdateSQLScript(ctx context.Context, sqlScriptName string, sqlScript SQLScriptResource, sqlScriptCreateOrUpdateSqlscriptOptions *SQLScriptCreateOrUpdateSQLScriptOptions) (*SQLScriptResourceResponse, error) {
	req, err := client.CreateOrUpdateSQLScriptCreateRequest(ctx, sqlScriptName, sqlScript, sqlScriptCreateOrUpdateSqlscriptOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.CreateOrUpdateSQLScriptHandleError(resp)
	}
	result, err := client.CreateOrUpdateSQLScriptHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateOrUpdateSQLScriptCreateRequest creates the CreateOrUpdateSQLScript request.
func (client *sqlScriptClient) CreateOrUpdateSQLScriptCreateRequest(ctx context.Context, sqlScriptName string, sqlScript SQLScriptResource, sqlScriptCreateOrUpdateSqlscriptOptions *SQLScriptCreateOrUpdateSQLScriptOptions) (*azcore.Request, error) {
	urlPath := "/sqlScripts/{sqlScriptName}"
	urlPath = strings.ReplaceAll(urlPath, "{sqlScriptName}", url.PathEscape(sqlScriptName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if sqlScriptCreateOrUpdateSqlscriptOptions != nil && sqlScriptCreateOrUpdateSqlscriptOptions.IfMatch != nil {
		req.Header.Set("If-Match", *sqlScriptCreateOrUpdateSqlscriptOptions.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(sqlScript)
}

// CreateOrUpdateSQLScriptHandleResponse handles the CreateOrUpdateSQLScript response.
func (client *sqlScriptClient) CreateOrUpdateSQLScriptHandleResponse(resp *azcore.Response) (*SQLScriptResourceResponse, error) {
	result := SQLScriptResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SQLScriptResource)
}

// CreateOrUpdateSQLScriptHandleError handles the CreateOrUpdateSQLScript error response.
func (client *sqlScriptClient) CreateOrUpdateSQLScriptHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DeleteSQLScript - Deletes a Sql Script.
func (client *sqlScriptClient) DeleteSQLScript(ctx context.Context, sqlScriptName string) (*http.Response, error) {
	req, err := client.DeleteSQLScriptCreateRequest(ctx, sqlScriptName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.DeleteSQLScriptHandleError(resp)
	}
	return resp.Response, nil
}

// DeleteSQLScriptCreateRequest creates the DeleteSQLScript request.
func (client *sqlScriptClient) DeleteSQLScriptCreateRequest(ctx context.Context, sqlScriptName string) (*azcore.Request, error) {
	urlPath := "/sqlScripts/{sqlScriptName}"
	urlPath = strings.ReplaceAll(urlPath, "{sqlScriptName}", url.PathEscape(sqlScriptName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteSQLScriptHandleError handles the DeleteSQLScript error response.
func (client *sqlScriptClient) DeleteSQLScriptHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetSQLScript - Gets a sql script.
func (client *sqlScriptClient) GetSQLScript(ctx context.Context, sqlScriptName string, sqlScriptGetSqlscriptOptions *SQLScriptGetSQLScriptOptions) (*SQLScriptResourceResponse, error) {
	req, err := client.GetSQLScriptCreateRequest(ctx, sqlScriptName, sqlScriptGetSqlscriptOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNotModified) {
		return nil, client.GetSQLScriptHandleError(resp)
	}
	result, err := client.GetSQLScriptHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetSQLScriptCreateRequest creates the GetSQLScript request.
func (client *sqlScriptClient) GetSQLScriptCreateRequest(ctx context.Context, sqlScriptName string, sqlScriptGetSqlscriptOptions *SQLScriptGetSQLScriptOptions) (*azcore.Request, error) {
	urlPath := "/sqlScripts/{sqlScriptName}"
	urlPath = strings.ReplaceAll(urlPath, "{sqlScriptName}", url.PathEscape(sqlScriptName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if sqlScriptGetSqlscriptOptions != nil && sqlScriptGetSqlscriptOptions.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *sqlScriptGetSqlscriptOptions.IfNoneMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetSQLScriptHandleResponse handles the GetSQLScript response.
func (client *sqlScriptClient) GetSQLScriptHandleResponse(resp *azcore.Response) (*SQLScriptResourceResponse, error) {
	result := SQLScriptResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SQLScriptResource)
}

// GetSQLScriptHandleError handles the GetSQLScript error response.
func (client *sqlScriptClient) GetSQLScriptHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetSQLScriptsByWorkspace - Lists sql scripts.
func (client *sqlScriptClient) GetSQLScriptsByWorkspace() SQLScriptsListResponsePager {
	return &sqlScriptsListResponsePager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.GetSQLScriptsByWorkspaceCreateRequest(ctx)
		},
		responder: client.GetSQLScriptsByWorkspaceHandleResponse,
		errorer:   client.GetSQLScriptsByWorkspaceHandleError,
		advancer: func(ctx context.Context, resp *SQLScriptsListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.SQLScriptsListResponse.NextLink)
		},
	}
}

// GetSQLScriptsByWorkspaceCreateRequest creates the GetSQLScriptsByWorkspace request.
func (client *sqlScriptClient) GetSQLScriptsByWorkspaceCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/sqlScripts"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetSQLScriptsByWorkspaceHandleResponse handles the GetSQLScriptsByWorkspace response.
func (client *sqlScriptClient) GetSQLScriptsByWorkspaceHandleResponse(resp *azcore.Response) (*SQLScriptsListResponseResponse, error) {
	result := SQLScriptsListResponseResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SQLScriptsListResponse)
}

// GetSQLScriptsByWorkspaceHandleError handles the GetSQLScriptsByWorkspace error response.
func (client *sqlScriptClient) GetSQLScriptsByWorkspaceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
