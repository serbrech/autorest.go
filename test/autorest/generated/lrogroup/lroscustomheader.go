// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package lrogroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"time"
)

// LrOSCustomHeaderOperations contains the methods for the LrOSCustomHeader group.
type LrOSCustomHeaderOperations interface {
	// BeginPost202Retry200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running post request, service returns a 202 to the initial request, with 'Location' and 'Retry-After' headers, Polls return a 200 with a response body after success
	BeginPost202Retry200(ctx context.Context, lrOSCustomHeaderPost202Retry200Options *LrOSCustomHeaderPost202Retry200Options) (*HTTPPollerResponse, error)
	// ResumePost202Retry200 - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumePost202Retry200(token string) (HTTPPoller, error)
	// BeginPostAsyncRetrySucceeded - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running post request, service returns a 202 to the initial request, with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
	BeginPostAsyncRetrySucceeded(ctx context.Context, lrOSCustomHeaderPostAsyncRetrySucceededOptions *LrOSCustomHeaderPostAsyncRetrySucceededOptions) (*HTTPPollerResponse, error)
	// ResumePostAsyncRetrySucceeded - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumePostAsyncRetrySucceeded(token string) (HTTPPoller, error)
	// BeginPut201CreatingSucceeded200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running put request, service returns a 201 to the initial request, with an entity that contains ProvisioningState=’Creating’.  Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
	BeginPut201CreatingSucceeded200(ctx context.Context, lrOSCustomHeaderPut201CreatingSucceeded200Options *LrOSCustomHeaderPut201CreatingSucceeded200Options) (*ProductPollerResponse, error)
	// ResumePut201CreatingSucceeded200 - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumePut201CreatingSucceeded200(token string) (ProductPoller, error)
	// BeginPutAsyncRetrySucceeded - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running put request, service returns a 200 to the initial request, with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
	BeginPutAsyncRetrySucceeded(ctx context.Context, lrOSCustomHeaderPutAsyncRetrySucceededOptions *LrOSCustomHeaderPutAsyncRetrySucceededOptions) (*ProductPollerResponse, error)
	// ResumePutAsyncRetrySucceeded - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumePutAsyncRetrySucceeded(token string) (ProductPoller, error)
}

// lrOSCustomHeaderOperations implements the LrOSCustomHeaderOperations interface.
type lrOSCustomHeaderOperations struct {
	*Client
}

// Post202Retry200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running post request, service returns a 202 to the initial request, with 'Location' and 'Retry-After' headers, Polls return a 200 with a response body after success
func (client *lrOSCustomHeaderOperations) BeginPost202Retry200(ctx context.Context, lrOSCustomHeaderPost202Retry200Options *LrOSCustomHeaderPost202Retry200Options) (*HTTPPollerResponse, error) {
	req, err := client.post202Retry200CreateRequest(lrOSCustomHeaderPost202Retry200Options)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.post202Retry200HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("lrOSCustomHeaderOperations.Post202Retry200", "", resp, client.post202Retry200HandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *lrOSCustomHeaderOperations) ResumePost202Retry200(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("lrOSCustomHeaderOperations.Post202Retry200", token, client.post202Retry200HandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// post202Retry200CreateRequest creates the Post202Retry200 request.
func (client *lrOSCustomHeaderOperations) post202Retry200CreateRequest(lrOSCustomHeaderPost202Retry200Options *LrOSCustomHeaderPost202Retry200Options) (*azcore.Request, error) {
	urlPath := "/lro/customheader/post/202/retry/200"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPost, *u)
	if lrOSCustomHeaderPost202Retry200Options != nil {
		return req, req.MarshalAsJSON(lrOSCustomHeaderPost202Retry200Options.Product)
	}
	return req, nil
}

// post202Retry200HandleResponse handles the Post202Retry200 response.
func (client *lrOSCustomHeaderOperations) post202Retry200HandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusAccepted, http.StatusNoContent) {
		return nil, client.post202Retry200HandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// post202Retry200HandleError handles the Post202Retry200 error response.
func (client *lrOSCustomHeaderOperations) post202Retry200HandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PostAsyncRetrySucceeded - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running post request, service returns a 202 to the initial request, with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
func (client *lrOSCustomHeaderOperations) BeginPostAsyncRetrySucceeded(ctx context.Context, lrOSCustomHeaderPostAsyncRetrySucceededOptions *LrOSCustomHeaderPostAsyncRetrySucceededOptions) (*HTTPPollerResponse, error) {
	req, err := client.postAsyncRetrySucceededCreateRequest(lrOSCustomHeaderPostAsyncRetrySucceededOptions)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.postAsyncRetrySucceededHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("lrOSCustomHeaderOperations.PostAsyncRetrySucceeded", "", resp, client.postAsyncRetrySucceededHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *lrOSCustomHeaderOperations) ResumePostAsyncRetrySucceeded(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("lrOSCustomHeaderOperations.PostAsyncRetrySucceeded", token, client.postAsyncRetrySucceededHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// postAsyncRetrySucceededCreateRequest creates the PostAsyncRetrySucceeded request.
func (client *lrOSCustomHeaderOperations) postAsyncRetrySucceededCreateRequest(lrOSCustomHeaderPostAsyncRetrySucceededOptions *LrOSCustomHeaderPostAsyncRetrySucceededOptions) (*azcore.Request, error) {
	urlPath := "/lro/customheader/postasync/retry/succeeded"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPost, *u)
	if lrOSCustomHeaderPostAsyncRetrySucceededOptions != nil {
		return req, req.MarshalAsJSON(lrOSCustomHeaderPostAsyncRetrySucceededOptions.Product)
	}
	return req, nil
}

// postAsyncRetrySucceededHandleResponse handles the PostAsyncRetrySucceeded response.
func (client *lrOSCustomHeaderOperations) postAsyncRetrySucceededHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusAccepted, http.StatusNoContent) {
		return nil, client.postAsyncRetrySucceededHandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// postAsyncRetrySucceededHandleError handles the PostAsyncRetrySucceeded error response.
func (client *lrOSCustomHeaderOperations) postAsyncRetrySucceededHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Put201CreatingSucceeded200 - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running put request, service returns a 201 to the initial request, with an entity that contains ProvisioningState=’Creating’.  Polls return this value until the last poll returns a ‘200’ with ProvisioningState=’Succeeded’
func (client *lrOSCustomHeaderOperations) BeginPut201CreatingSucceeded200(ctx context.Context, lrOSCustomHeaderPut201CreatingSucceeded200Options *LrOSCustomHeaderPut201CreatingSucceeded200Options) (*ProductPollerResponse, error) {
	req, err := client.put201CreatingSucceeded200CreateRequest(lrOSCustomHeaderPut201CreatingSucceeded200Options)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.put201CreatingSucceeded200HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("lrOSCustomHeaderOperations.Put201CreatingSucceeded200", "", resp, client.put201CreatingSucceeded200HandleError)
	if err != nil {
		return nil, err
	}
	poller := &productPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ProductResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *lrOSCustomHeaderOperations) ResumePut201CreatingSucceeded200(token string) (ProductPoller, error) {
	pt, err := resumePollingTracker("lrOSCustomHeaderOperations.Put201CreatingSucceeded200", token, client.put201CreatingSucceeded200HandleError)
	if err != nil {
		return nil, err
	}
	return &productPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// put201CreatingSucceeded200CreateRequest creates the Put201CreatingSucceeded200 request.
func (client *lrOSCustomHeaderOperations) put201CreatingSucceeded200CreateRequest(lrOSCustomHeaderPut201CreatingSucceeded200Options *LrOSCustomHeaderPut201CreatingSucceeded200Options) (*azcore.Request, error) {
	urlPath := "/lro/customheader/put/201/creating/succeeded/200"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	if lrOSCustomHeaderPut201CreatingSucceeded200Options != nil {
		return req, req.MarshalAsJSON(lrOSCustomHeaderPut201CreatingSucceeded200Options.Product)
	}
	return req, nil
}

// put201CreatingSucceeded200HandleResponse handles the Put201CreatingSucceeded200 response.
func (client *lrOSCustomHeaderOperations) put201CreatingSucceeded200HandleResponse(resp *azcore.Response) (*ProductPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.put201CreatingSucceeded200HandleError(resp)
	}
	return &ProductPollerResponse{RawResponse: resp.Response}, nil
}

// put201CreatingSucceeded200HandleError handles the Put201CreatingSucceeded200 error response.
func (client *lrOSCustomHeaderOperations) put201CreatingSucceeded200HandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutAsyncRetrySucceeded - x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 is required message header for all requests. Long running put request, service returns a 200 to the initial request, with an entity that contains ProvisioningState=’Creating’. Poll the endpoint indicated in the Azure-AsyncOperation header for operation status
func (client *lrOSCustomHeaderOperations) BeginPutAsyncRetrySucceeded(ctx context.Context, lrOSCustomHeaderPutAsyncRetrySucceededOptions *LrOSCustomHeaderPutAsyncRetrySucceededOptions) (*ProductPollerResponse, error) {
	req, err := client.putAsyncRetrySucceededCreateRequest(lrOSCustomHeaderPutAsyncRetrySucceededOptions)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putAsyncRetrySucceededHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("lrOSCustomHeaderOperations.PutAsyncRetrySucceeded", "", resp, client.putAsyncRetrySucceededHandleError)
	if err != nil {
		return nil, err
	}
	poller := &productPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ProductResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *lrOSCustomHeaderOperations) ResumePutAsyncRetrySucceeded(token string) (ProductPoller, error) {
	pt, err := resumePollingTracker("lrOSCustomHeaderOperations.PutAsyncRetrySucceeded", token, client.putAsyncRetrySucceededHandleError)
	if err != nil {
		return nil, err
	}
	return &productPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// putAsyncRetrySucceededCreateRequest creates the PutAsyncRetrySucceeded request.
func (client *lrOSCustomHeaderOperations) putAsyncRetrySucceededCreateRequest(lrOSCustomHeaderPutAsyncRetrySucceededOptions *LrOSCustomHeaderPutAsyncRetrySucceededOptions) (*azcore.Request, error) {
	urlPath := "/lro/customheader/putasync/retry/succeeded"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPut, *u)
	if lrOSCustomHeaderPutAsyncRetrySucceededOptions != nil {
		return req, req.MarshalAsJSON(lrOSCustomHeaderPutAsyncRetrySucceededOptions.Product)
	}
	return req, nil
}

// putAsyncRetrySucceededHandleResponse handles the PutAsyncRetrySucceeded response.
func (client *lrOSCustomHeaderOperations) putAsyncRetrySucceededHandleResponse(resp *azcore.Response) (*ProductPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.putAsyncRetrySucceededHandleError(resp)
	}
	return &ProductPollerResponse{RawResponse: resp.Response}, nil
}

// putAsyncRetrySucceededHandleError handles the PutAsyncRetrySucceeded error response.
func (client *lrOSCustomHeaderOperations) putAsyncRetrySucceededHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}