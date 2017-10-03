package custombaseurlgroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// PathsClient is the test Infrastructure for AutoRest
type PathsClient struct {
	ManagementClient
}

// NewPathsClient creates an instance of the PathsClient client.
func NewPathsClient() PathsClient {
	return PathsClient{New()}
}

// GetEmpty get a 200 to test a valid base uri
//
// accountName is account Name
func (client PathsClient) GetEmpty(accountName string) (result autorest.Response, err error) {
	req, err := client.GetEmptyPreparer(accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "custombaseurlgroup.PathsClient", "GetEmpty", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetEmptySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "custombaseurlgroup.PathsClient", "GetEmpty", resp, "Failure sending request")
		return
	}

	result, err = client.GetEmptyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "custombaseurlgroup.PathsClient", "GetEmpty", resp, "Failure responding to request")
	}

	return
}

// GetEmptyPreparer prepares the GetEmpty request.
func (client PathsClient) GetEmptyPreparer(accountName string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"accountName": accountName,
		"host":        client.Host,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("http://{accountName}{host}", urlParameters),
		autorest.WithPath("/customuri"))
	return preparer.Prepare(&http.Request{})
}

// GetEmptySender sends the GetEmpty request. The method will close the
// http.Response Body if it receives an error.
func (client PathsClient) GetEmptySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetEmptyResponder handles the response to the GetEmpty request. The method always
// closes the http.Response Body.
func (client PathsClient) GetEmptyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}