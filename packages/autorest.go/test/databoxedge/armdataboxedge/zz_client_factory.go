//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataboxedge

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential     azcore.TokenCredential
	options        *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

// NewAddonsClient creates a new instance of AddonsClient.
func (c *ClientFactory) NewAddonsClient() *AddonsClient {
	subClient, _ := NewAddonsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewAlertsClient creates a new instance of AlertsClient.
func (c *ClientFactory) NewAlertsClient() *AlertsClient {
	subClient, _ := NewAlertsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewAvailableSKUsClient creates a new instance of AvailableSKUsClient.
func (c *ClientFactory) NewAvailableSKUsClient() *AvailableSKUsClient {
	subClient, _ := NewAvailableSKUsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewBandwidthSchedulesClient creates a new instance of BandwidthSchedulesClient.
func (c *ClientFactory) NewBandwidthSchedulesClient() *BandwidthSchedulesClient {
	subClient, _ := NewBandwidthSchedulesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewContainersClient creates a new instance of ContainersClient.
func (c *ClientFactory) NewContainersClient() *ContainersClient {
	subClient, _ := NewContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewDevicesClient creates a new instance of DevicesClient.
func (c *ClientFactory) NewDevicesClient() *DevicesClient {
	subClient, _ := NewDevicesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewDiagnosticSettingsClient creates a new instance of DiagnosticSettingsClient.
func (c *ClientFactory) NewDiagnosticSettingsClient() *DiagnosticSettingsClient {
	subClient, _ := NewDiagnosticSettingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewJobsClient creates a new instance of JobsClient.
func (c *ClientFactory) NewJobsClient() *JobsClient {
	subClient, _ := NewJobsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewMonitoringConfigClient creates a new instance of MonitoringConfigClient.
func (c *ClientFactory) NewMonitoringConfigClient() *MonitoringConfigClient {
	subClient, _ := NewMonitoringConfigClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewNodesClient creates a new instance of NodesClient.
func (c *ClientFactory) NewNodesClient() *NodesClient {
	subClient, _ := NewNodesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

// NewOperationsStatusClient creates a new instance of OperationsStatusClient.
func (c *ClientFactory) NewOperationsStatusClient() *OperationsStatusClient {
	subClient, _ := NewOperationsStatusClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOrdersClient creates a new instance of OrdersClient.
func (c *ClientFactory) NewOrdersClient() *OrdersClient {
	subClient, _ := NewOrdersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewRolesClient creates a new instance of RolesClient.
func (c *ClientFactory) NewRolesClient() *RolesClient {
	subClient, _ := NewRolesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewSharesClient creates a new instance of SharesClient.
func (c *ClientFactory) NewSharesClient() *SharesClient {
	subClient, _ := NewSharesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewStorageAccountCredentialsClient creates a new instance of StorageAccountCredentialsClient.
func (c *ClientFactory) NewStorageAccountCredentialsClient() *StorageAccountCredentialsClient {
	subClient, _ := NewStorageAccountCredentialsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewStorageAccountsClient creates a new instance of StorageAccountsClient.
func (c *ClientFactory) NewStorageAccountsClient() *StorageAccountsClient {
	subClient, _ := NewStorageAccountsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewSupportPackagesClient creates a new instance of SupportPackagesClient.
func (c *ClientFactory) NewSupportPackagesClient() *SupportPackagesClient {
	subClient, _ := NewSupportPackagesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewTriggersClient creates a new instance of TriggersClient.
func (c *ClientFactory) NewTriggersClient() *TriggersClient {
	subClient, _ := NewTriggersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewUsersClient creates a new instance of UsersClient.
func (c *ClientFactory) NewUsersClient() *UsersClient {
	subClient, _ := NewUsersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
