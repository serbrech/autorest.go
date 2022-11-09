//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azartifacts

// BigDataPoolsClientGetResponse contains the response from method bigDataPoolsClient.Get.
type BigDataPoolsClientGetResponse struct {
	BigDataPoolResourceInfo
}

// BigDataPoolsClientListResponse contains the response from method bigDataPoolsClient.List.
type BigDataPoolsClientListResponse struct {
	BigDataPoolResourceInfoListResult
}

// DataFlowClientCreateOrUpdateDataFlowResponse contains the response from method dataFlowClient.BeginCreateOrUpdateDataFlow.
type DataFlowClientCreateOrUpdateDataFlowResponse struct {
	DataFlowResource
}

// DataFlowClientDeleteDataFlowResponse contains the response from method dataFlowClient.BeginDeleteDataFlow.
type DataFlowClientDeleteDataFlowResponse struct {
	// placeholder for future response values
}

// DataFlowClientGetDataFlowResponse contains the response from method dataFlowClient.GetDataFlow.
type DataFlowClientGetDataFlowResponse struct {
	DataFlowResource
}

// DataFlowClientGetDataFlowsByWorkspaceResponse contains the response from method dataFlowClient.NewGetDataFlowsByWorkspacePager.
type DataFlowClientGetDataFlowsByWorkspaceResponse struct {
	DataFlowListResponse
}

// DataFlowClientRenameDataFlowResponse contains the response from method dataFlowClient.BeginRenameDataFlow.
type DataFlowClientRenameDataFlowResponse struct {
	// placeholder for future response values
}

// DataFlowDebugSessionClientAddDataFlowResponse contains the response from method dataFlowDebugSessionClient.AddDataFlow.
type DataFlowDebugSessionClientAddDataFlowResponse struct {
	AddDataFlowToDebugSessionResponse
}

// DataFlowDebugSessionClientCreateDataFlowDebugSessionResponse contains the response from method dataFlowDebugSessionClient.BeginCreateDataFlowDebugSession.
type DataFlowDebugSessionClientCreateDataFlowDebugSessionResponse struct {
	CreateDataFlowDebugSessionResponse
}

// DataFlowDebugSessionClientDeleteDataFlowDebugSessionResponse contains the response from method dataFlowDebugSessionClient.DeleteDataFlowDebugSession.
type DataFlowDebugSessionClientDeleteDataFlowDebugSessionResponse struct {
	// placeholder for future response values
}

// DataFlowDebugSessionClientExecuteCommandResponse contains the response from method dataFlowDebugSessionClient.BeginExecuteCommand.
type DataFlowDebugSessionClientExecuteCommandResponse struct {
	DataFlowDebugCommandResponse
}

// DataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspaceResponse contains the response from method dataFlowDebugSessionClient.NewQueryDataFlowDebugSessionsByWorkspacePager.
type DataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspaceResponse struct {
	QueryDataFlowDebugSessionsResponse
}

// DatasetClientCreateOrUpdateDatasetResponse contains the response from method datasetClient.BeginCreateOrUpdateDataset.
type DatasetClientCreateOrUpdateDatasetResponse struct {
	DatasetResource
}

// DatasetClientDeleteDatasetResponse contains the response from method datasetClient.BeginDeleteDataset.
type DatasetClientDeleteDatasetResponse struct {
	// placeholder for future response values
}

// DatasetClientGetDatasetResponse contains the response from method datasetClient.GetDataset.
type DatasetClientGetDatasetResponse struct {
	DatasetResource
}

// DatasetClientGetDatasetsByWorkspaceResponse contains the response from method datasetClient.NewGetDatasetsByWorkspacePager.
type DatasetClientGetDatasetsByWorkspaceResponse struct {
	DatasetListResponse
}

// DatasetClientRenameDatasetResponse contains the response from method datasetClient.BeginRenameDataset.
type DatasetClientRenameDatasetResponse struct {
	// placeholder for future response values
}

// IntegrationRuntimesClientGetResponse contains the response from method integrationRuntimesClient.Get.
type IntegrationRuntimesClientGetResponse struct {
	IntegrationRuntimeResource
}

// IntegrationRuntimesClientListResponse contains the response from method integrationRuntimesClient.List.
type IntegrationRuntimesClientListResponse struct {
	IntegrationRuntimeListResponse
}

// LibraryClientAppendResponse contains the response from method libraryClient.Append.
type LibraryClientAppendResponse struct {
	// placeholder for future response values
}

// LibraryClientCreateResponse contains the response from method libraryClient.BeginCreate.
type LibraryClientCreateResponse struct {
	LibraryResourceInfo
}

// LibraryClientDeleteResponse contains the response from method libraryClient.BeginDelete.
type LibraryClientDeleteResponse struct {
	LibraryResourceInfo
}

// LibraryClientFlushResponse contains the response from method libraryClient.BeginFlush.
type LibraryClientFlushResponse struct {
	LibraryResourceInfo
}

// LibraryClientGetOperationResultResponse contains the response from method libraryClient.GetOperationResult.
type LibraryClientGetOperationResultResponse struct {
	// Possible types are LibraryResource, OperationResult
	Value any
}

// LibraryClientGetResponse contains the response from method libraryClient.Get.
type LibraryClientGetResponse struct {
	LibraryResource
}

// LibraryClientListResponse contains the response from method libraryClient.NewListPager.
type LibraryClientListResponse struct {
	LibraryListResponse
}

// LinkedServiceClientCreateOrUpdateLinkedServiceResponse contains the response from method linkedServiceClient.BeginCreateOrUpdateLinkedService.
type LinkedServiceClientCreateOrUpdateLinkedServiceResponse struct {
	LinkedServiceResource
}

// LinkedServiceClientDeleteLinkedServiceResponse contains the response from method linkedServiceClient.BeginDeleteLinkedService.
type LinkedServiceClientDeleteLinkedServiceResponse struct {
	// placeholder for future response values
}

// LinkedServiceClientGetLinkedServiceResponse contains the response from method linkedServiceClient.GetLinkedService.
type LinkedServiceClientGetLinkedServiceResponse struct {
	LinkedServiceResource
}

// LinkedServiceClientGetLinkedServicesByWorkspaceResponse contains the response from method linkedServiceClient.NewGetLinkedServicesByWorkspacePager.
type LinkedServiceClientGetLinkedServicesByWorkspaceResponse struct {
	LinkedServiceListResponse
}

// LinkedServiceClientRenameLinkedServiceResponse contains the response from method linkedServiceClient.BeginRenameLinkedService.
type LinkedServiceClientRenameLinkedServiceResponse struct {
	// placeholder for future response values
}

// NotebookClientCreateOrUpdateNotebookResponse contains the response from method notebookClient.BeginCreateOrUpdateNotebook.
type NotebookClientCreateOrUpdateNotebookResponse struct {
	NotebookResource
}

// NotebookClientDeleteNotebookResponse contains the response from method notebookClient.BeginDeleteNotebook.
type NotebookClientDeleteNotebookResponse struct {
	// placeholder for future response values
}

// NotebookClientGetNotebookResponse contains the response from method notebookClient.GetNotebook.
type NotebookClientGetNotebookResponse struct {
	NotebookResource
}

// NotebookClientGetNotebookSummaryByWorkSpaceResponse contains the response from method notebookClient.NewGetNotebookSummaryByWorkSpacePager.
type NotebookClientGetNotebookSummaryByWorkSpaceResponse struct {
	NotebookListResponse
}

// NotebookClientGetNotebooksByWorkspaceResponse contains the response from method notebookClient.NewGetNotebooksByWorkspacePager.
type NotebookClientGetNotebooksByWorkspaceResponse struct {
	NotebookListResponse
}

// NotebookClientRenameNotebookResponse contains the response from method notebookClient.BeginRenameNotebook.
type NotebookClientRenameNotebookResponse struct {
	// placeholder for future response values
}

// PipelineClientCreateOrUpdatePipelineResponse contains the response from method pipelineClient.BeginCreateOrUpdatePipeline.
type PipelineClientCreateOrUpdatePipelineResponse struct {
	PipelineResource
}

// PipelineClientCreatePipelineRunResponse contains the response from method pipelineClient.CreatePipelineRun.
type PipelineClientCreatePipelineRunResponse struct {
	CreateRunResponse
}

// PipelineClientDeletePipelineResponse contains the response from method pipelineClient.BeginDeletePipeline.
type PipelineClientDeletePipelineResponse struct {
	// placeholder for future response values
}

// PipelineClientGetPipelineResponse contains the response from method pipelineClient.GetPipeline.
type PipelineClientGetPipelineResponse struct {
	PipelineResource
}

// PipelineClientGetPipelinesByWorkspaceResponse contains the response from method pipelineClient.NewGetPipelinesByWorkspacePager.
type PipelineClientGetPipelinesByWorkspaceResponse struct {
	PipelineListResponse
}

// PipelineClientRenamePipelineResponse contains the response from method pipelineClient.BeginRenamePipeline.
type PipelineClientRenamePipelineResponse struct {
	// placeholder for future response values
}

// PipelineRunClientCancelPipelineRunResponse contains the response from method pipelineRunClient.CancelPipelineRun.
type PipelineRunClientCancelPipelineRunResponse struct {
	// placeholder for future response values
}

// PipelineRunClientGetPipelineRunResponse contains the response from method pipelineRunClient.GetPipelineRun.
type PipelineRunClientGetPipelineRunResponse struct {
	PipelineRun
}

// PipelineRunClientQueryActivityRunsResponse contains the response from method pipelineRunClient.QueryActivityRuns.
type PipelineRunClientQueryActivityRunsResponse struct {
	ActivityRunsQueryResponse
}

// PipelineRunClientQueryPipelineRunsByWorkspaceResponse contains the response from method pipelineRunClient.QueryPipelineRunsByWorkspace.
type PipelineRunClientQueryPipelineRunsByWorkspaceResponse struct {
	PipelineRunsQueryResponse
}

// SparkJobDefinitionClientCreateOrUpdateSparkJobDefinitionResponse contains the response from method sparkJobDefinitionClient.BeginCreateOrUpdateSparkJobDefinition.
type SparkJobDefinitionClientCreateOrUpdateSparkJobDefinitionResponse struct {
	SparkJobDefinitionResource
}

// SparkJobDefinitionClientDebugSparkJobDefinitionResponse contains the response from method sparkJobDefinitionClient.BeginDebugSparkJobDefinition.
type SparkJobDefinitionClientDebugSparkJobDefinitionResponse struct {
	SparkBatchJob
}

// SparkJobDefinitionClientDeleteSparkJobDefinitionResponse contains the response from method sparkJobDefinitionClient.BeginDeleteSparkJobDefinition.
type SparkJobDefinitionClientDeleteSparkJobDefinitionResponse struct {
	// placeholder for future response values
}

// SparkJobDefinitionClientExecuteSparkJobDefinitionResponse contains the response from method sparkJobDefinitionClient.BeginExecuteSparkJobDefinition.
type SparkJobDefinitionClientExecuteSparkJobDefinitionResponse struct {
	SparkBatchJob
}

// SparkJobDefinitionClientGetSparkJobDefinitionResponse contains the response from method sparkJobDefinitionClient.GetSparkJobDefinition.
type SparkJobDefinitionClientGetSparkJobDefinitionResponse struct {
	SparkJobDefinitionResource
}

// SparkJobDefinitionClientGetSparkJobDefinitionsByWorkspaceResponse contains the response from method sparkJobDefinitionClient.NewGetSparkJobDefinitionsByWorkspacePager.
type SparkJobDefinitionClientGetSparkJobDefinitionsByWorkspaceResponse struct {
	SparkJobDefinitionsListResponse
}

// SparkJobDefinitionClientRenameSparkJobDefinitionResponse contains the response from method sparkJobDefinitionClient.BeginRenameSparkJobDefinition.
type SparkJobDefinitionClientRenameSparkJobDefinitionResponse struct {
	// placeholder for future response values
}

// SqlPoolsClientGetResponse contains the response from method sqlPoolsClient.Get.
type SqlPoolsClientGetResponse struct {
	SQLPool
}

// SqlPoolsClientListResponse contains the response from method sqlPoolsClient.List.
type SqlPoolsClientListResponse struct {
	SQLPoolInfoListResult
}

// SqlScriptClientCreateOrUpdateSQLScriptResponse contains the response from method sqlScriptClient.BeginCreateOrUpdateSQLScript.
type SqlScriptClientCreateOrUpdateSQLScriptResponse struct {
	SQLScriptResource
}

// SqlScriptClientDeleteSQLScriptResponse contains the response from method sqlScriptClient.BeginDeleteSQLScript.
type SqlScriptClientDeleteSQLScriptResponse struct {
	// placeholder for future response values
}

// SqlScriptClientGetSQLScriptResponse contains the response from method sqlScriptClient.GetSQLScript.
type SqlScriptClientGetSQLScriptResponse struct {
	SQLScriptResource
}

// SqlScriptClientGetSQLScriptsByWorkspaceResponse contains the response from method sqlScriptClient.NewGetSQLScriptsByWorkspacePager.
type SqlScriptClientGetSQLScriptsByWorkspaceResponse struct {
	SQLScriptsListResponse
}

// SqlScriptClientRenameSQLScriptResponse contains the response from method sqlScriptClient.BeginRenameSQLScript.
type SqlScriptClientRenameSQLScriptResponse struct {
	// placeholder for future response values
}

// TriggerClientCreateOrUpdateTriggerResponse contains the response from method triggerClient.BeginCreateOrUpdateTrigger.
type TriggerClientCreateOrUpdateTriggerResponse struct {
	TriggerResource
}

// TriggerClientDeleteTriggerResponse contains the response from method triggerClient.BeginDeleteTrigger.
type TriggerClientDeleteTriggerResponse struct {
	// placeholder for future response values
}

// TriggerClientGetEventSubscriptionStatusResponse contains the response from method triggerClient.GetEventSubscriptionStatus.
type TriggerClientGetEventSubscriptionStatusResponse struct {
	TriggerSubscriptionOperationStatus
}

// TriggerClientGetTriggerResponse contains the response from method triggerClient.GetTrigger.
type TriggerClientGetTriggerResponse struct {
	TriggerResource
}

// TriggerClientGetTriggersByWorkspaceResponse contains the response from method triggerClient.NewGetTriggersByWorkspacePager.
type TriggerClientGetTriggersByWorkspaceResponse struct {
	TriggerListResponse
}

// TriggerClientStartTriggerResponse contains the response from method triggerClient.BeginStartTrigger.
type TriggerClientStartTriggerResponse struct {
	// placeholder for future response values
}

// TriggerClientStopTriggerResponse contains the response from method triggerClient.BeginStopTrigger.
type TriggerClientStopTriggerResponse struct {
	// placeholder for future response values
}

// TriggerClientSubscribeTriggerToEventsResponse contains the response from method triggerClient.BeginSubscribeTriggerToEvents.
type TriggerClientSubscribeTriggerToEventsResponse struct {
	TriggerSubscriptionOperationStatus
}

// TriggerClientUnsubscribeTriggerFromEventsResponse contains the response from method triggerClient.BeginUnsubscribeTriggerFromEvents.
type TriggerClientUnsubscribeTriggerFromEventsResponse struct {
	TriggerSubscriptionOperationStatus
}

// TriggerRunClientCancelTriggerInstanceResponse contains the response from method triggerRunClient.CancelTriggerInstance.
type TriggerRunClientCancelTriggerInstanceResponse struct {
	// placeholder for future response values
}

// TriggerRunClientQueryTriggerRunsByWorkspaceResponse contains the response from method triggerRunClient.QueryTriggerRunsByWorkspace.
type TriggerRunClientQueryTriggerRunsByWorkspaceResponse struct {
	TriggerRunsQueryResponse
}

// TriggerRunClientRerunTriggerInstanceResponse contains the response from method triggerRunClient.RerunTriggerInstance.
type TriggerRunClientRerunTriggerInstanceResponse struct {
	// placeholder for future response values
}

// WorkspaceClientGetResponse contains the response from method workspaceClient.Get.
type WorkspaceClientGetResponse struct {
	Workspace
}

// WorkspaceGitRepoManagementClientGetGitHubAccessTokenResponse contains the response from method workspaceGitRepoManagementClient.GetGitHubAccessToken.
type WorkspaceGitRepoManagementClientGetGitHubAccessTokenResponse struct {
	GitHubAccessTokenResponse
}