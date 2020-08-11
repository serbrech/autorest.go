// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package lrogroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"time"
)

// HTTPPoller provides polling facilities until the operation completes
type HTTPPoller interface {
	Done() bool
	Poll(ctx context.Context) (*http.Response, error)
	FinalResponse(ctx context.Context) (*http.Response, error)
	ResumeToken() (string, error)
}

type httpPoller struct {
	// the client for making the request
	pipeline azcore.Pipeline
	pt       armcore.Poller
}

// Done returns true if there was an error or polling has reached a terminal state
func (p *httpPoller) Done() bool {
	return p.pt.Done()
}

// Poll will send poll the service endpoint and return an http.Response or error received from the service
func (p *httpPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx, p.pipeline)
}

func (p *httpPoller) FinalResponse(ctx context.Context) (*http.Response, error) {
	return p.pt.FinalResponse(ctx, p.pipeline, nil)
}

// ResumeToken generates the string token that can be used with the ResumeHTTPPoller method
// on the client to create a new poller from the data held in the current poller type
func (p *httpPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *httpPoller) pollUntilDone(ctx context.Context, frequency time.Duration) (*http.Response, error) {
	return p.pt.PollUntilDone(ctx, frequency, p.pipeline, nil)
}

// ProductArrayPoller provides polling facilities until the operation completes
type ProductArrayPoller interface {
	Done() bool
	Poll(ctx context.Context) (*http.Response, error)
	FinalResponse(ctx context.Context) (*ProductArrayResponse, error)
	ResumeToken() (string, error)
}

type productArrayPoller struct {
	// the client for making the request
	pipeline azcore.Pipeline
	pt       armcore.Poller
}

// Done returns true if there was an error or polling has reached a terminal state
func (p *productArrayPoller) Done() bool {
	return p.pt.Done()
}

// Poll will send poll the service endpoint and return an http.Response or error received from the service
func (p *productArrayPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx, p.pipeline)
}

func (p *productArrayPoller) FinalResponse(ctx context.Context) (*ProductArrayResponse, error) {
	respType := &ProductArrayResponse{ProductArray: &[]Product{}}
	resp, err := p.pt.FinalResponse(ctx, p.pipeline, respType.ProductArray)
	if err != nil {
		return nil, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken generates the string token that can be used with the ResumeProductArrayPoller method
// on the client to create a new poller from the data held in the current poller type
func (p *productArrayPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *productArrayPoller) pollUntilDone(ctx context.Context, frequency time.Duration) (*ProductArrayResponse, error) {
	respType := &ProductArrayResponse{ProductArray: &[]Product{}}
	resp, err := p.pt.PollUntilDone(ctx, frequency, p.pipeline, respType.ProductArray)
	if err != nil {
		return nil, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ProductPoller provides polling facilities until the operation completes
type ProductPoller interface {
	Done() bool
	Poll(ctx context.Context) (*http.Response, error)
	FinalResponse(ctx context.Context) (*ProductResponse, error)
	ResumeToken() (string, error)
}

type productPoller struct {
	// the client for making the request
	pipeline azcore.Pipeline
	pt       armcore.Poller
}

// Done returns true if there was an error or polling has reached a terminal state
func (p *productPoller) Done() bool {
	return p.pt.Done()
}

// Poll will send poll the service endpoint and return an http.Response or error received from the service
func (p *productPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx, p.pipeline)
}

func (p *productPoller) FinalResponse(ctx context.Context) (*ProductResponse, error) {
	respType := &ProductResponse{Product: &Product{}}
	resp, err := p.pt.FinalResponse(ctx, p.pipeline, respType.Product)
	if err != nil {
		return nil, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken generates the string token that can be used with the ResumeProductPoller method
// on the client to create a new poller from the data held in the current poller type
func (p *productPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *productPoller) pollUntilDone(ctx context.Context, frequency time.Duration) (*ProductResponse, error) {
	respType := &ProductResponse{Product: &Product{}}
	resp, err := p.pt.PollUntilDone(ctx, frequency, p.pipeline, respType.Product)
	if err != nil {
		return nil, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// SkuPoller provides polling facilities until the operation completes
type SkuPoller interface {
	Done() bool
	Poll(ctx context.Context) (*http.Response, error)
	FinalResponse(ctx context.Context) (*SkuResponse, error)
	ResumeToken() (string, error)
}

type skuPoller struct {
	// the client for making the request
	pipeline azcore.Pipeline
	pt       armcore.Poller
}

// Done returns true if there was an error or polling has reached a terminal state
func (p *skuPoller) Done() bool {
	return p.pt.Done()
}

// Poll will send poll the service endpoint and return an http.Response or error received from the service
func (p *skuPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx, p.pipeline)
}

func (p *skuPoller) FinalResponse(ctx context.Context) (*SkuResponse, error) {
	respType := &SkuResponse{Sku: &Sku{}}
	resp, err := p.pt.FinalResponse(ctx, p.pipeline, respType.Sku)
	if err != nil {
		return nil, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken generates the string token that can be used with the ResumeSkuPoller method
// on the client to create a new poller from the data held in the current poller type
func (p *skuPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *skuPoller) pollUntilDone(ctx context.Context, frequency time.Duration) (*SkuResponse, error) {
	respType := &SkuResponse{Sku: &Sku{}}
	resp, err := p.pt.PollUntilDone(ctx, frequency, p.pipeline, respType.Sku)
	if err != nil {
		return nil, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// SubProductPoller provides polling facilities until the operation completes
type SubProductPoller interface {
	Done() bool
	Poll(ctx context.Context) (*http.Response, error)
	FinalResponse(ctx context.Context) (*SubProductResponse, error)
	ResumeToken() (string, error)
}

type subProductPoller struct {
	// the client for making the request
	pipeline azcore.Pipeline
	pt       armcore.Poller
}

// Done returns true if there was an error or polling has reached a terminal state
func (p *subProductPoller) Done() bool {
	return p.pt.Done()
}

// Poll will send poll the service endpoint and return an http.Response or error received from the service
func (p *subProductPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx, p.pipeline)
}

func (p *subProductPoller) FinalResponse(ctx context.Context) (*SubProductResponse, error) {
	respType := &SubProductResponse{SubProduct: &SubProduct{}}
	resp, err := p.pt.FinalResponse(ctx, p.pipeline, respType.SubProduct)
	if err != nil {
		return nil, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken generates the string token that can be used with the ResumeSubProductPoller method
// on the client to create a new poller from the data held in the current poller type
func (p *subProductPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *subProductPoller) pollUntilDone(ctx context.Context, frequency time.Duration) (*SubProductResponse, error) {
	respType := &SubProductResponse{SubProduct: &SubProduct{}}
	resp, err := p.pt.PollUntilDone(ctx, frequency, p.pipeline, respType.SubProduct)
	if err != nil {
		return nil, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func delay(ctx context.Context, delay time.Duration) error {
	select {
	case <-time.After(delay):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
