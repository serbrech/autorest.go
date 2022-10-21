//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/compute/armcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/LogAnalyticsRequestRateByInterval.json
func ExampleLogAnalyticsClient_BeginExportRequestRateByInterval() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewLogAnalyticsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginExportRequestRateByInterval(ctx, "westus", armcompute.RequestRateByIntervalInput{
		BlobContainerSasURI: to.Ptr("https://somesasuri"),
		FromTime:            to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-21T01:54:06.862601Z"); return t }()),
		GroupByResourceName: to.Ptr(true),
		ToTime:              to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-23T01:54:06.862601Z"); return t }()),
		IntervalLength:      to.Ptr(armcompute.IntervalInMinsFiveMins),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LogAnalyticsOperationResult = armcompute.LogAnalyticsOperationResult{
	// 	Properties: &armcompute.LogAnalyticsOutput{
	// 		Output: to.Ptr("https://crptestar4227.blob.core.windows.net:443/sascontainer/RequestRateByInterval_20180121-0154_20180123-0154.csv"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/LogAnalyticsThrottledRequests.json
func ExampleLogAnalyticsClient_BeginExportThrottledRequests() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewLogAnalyticsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginExportThrottledRequests(ctx, "westus", armcompute.ThrottledRequestsInput{
		BlobContainerSasURI:        to.Ptr("https://somesasuri"),
		FromTime:                   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-21T01:54:06.862601Z"); return t }()),
		GroupByClientApplicationID: to.Ptr(false),
		GroupByOperationName:       to.Ptr(true),
		GroupByResourceName:        to.Ptr(false),
		GroupByUserAgent:           to.Ptr(false),
		ToTime:                     to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-23T01:54:06.862601Z"); return t }()),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LogAnalyticsOperationResult = armcompute.LogAnalyticsOperationResult{
	// 	Properties: &armcompute.LogAnalyticsOutput{
	// 		Output: to.Ptr("https://crptestar4227.blob.core.windows.net:443/sascontainer/ThrottledRequests_20180121-0154_20180123-0154.csv"),
	// 	},
	// }
}
