// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package paginggroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// OdataProductResultPager provides iteration over OdataProductResult pages.
type OdataProductResultPager interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current OdataProductResultResponse.
	PageResponse() *OdataProductResultResponse

	// Err returns the last error encountered while paging.
	Err() error
}

type odataProductResultHandleResponse func(*azcore.Response) (*OdataProductResultResponse, error)

type odataProductResultAdvancePage func(*OdataProductResultResponse) (*azcore.Request, error)

type odataProductResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// contains the pending request
	request *azcore.Request
	// callback for handling the HTTP response
	responder odataProductResultHandleResponse
	// callback for advancing to the next page
	advancer odataProductResultAdvancePage
	// contains the current response
	current *OdataProductResultResponse
	// any error encountered
	err error
}

func (p *odataProductResultPager) Err() error {
	return p.err
}

func (p *odataProductResultPager) NextPage(ctx context.Context) bool {
	if p.current != nil {
		if p.current.OdataProductResult.OdataNextLink == nil || len(*p.current.OdataProductResult.OdataNextLink) == 0 {
			return false
		}
		req, err := p.advancer(p.current)
		if err != nil {
			p.err = err
			return false
		}
		p.request = req
	}
	resp, err := p.pipeline.Do(ctx, p.request)
	if err != nil {
		p.err = err
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *odataProductResultPager) PageResponse() *OdataProductResultResponse {
	return p.current
}

// ProductResultPager provides iteration over ProductResult pages.
type ProductResultPager interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current ProductResultResponse.
	PageResponse() *ProductResultResponse

	// Err returns the last error encountered while paging.
	Err() error
}

type productResultHandleResponse func(*azcore.Response) (*ProductResultResponse, error)

type productResultAdvancePage func(*ProductResultResponse) (*azcore.Request, error)

type productResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// contains the pending request
	request *azcore.Request
	// callback for handling the HTTP response
	responder productResultHandleResponse
	// callback for advancing to the next page
	advancer productResultAdvancePage
	// contains the current response
	current *ProductResultResponse
	// any error encountered
	err error
	// previous response from the endpoint
	resp *azcore.Response
}

func (p *productResultPager) Err() error {
	return p.err
}

func (p *productResultPager) NextPage(ctx context.Context) bool {
	if p.current != nil {
		if p.current.ProductResult.NextLink == nil || len(*p.current.ProductResult.NextLink) == 0 {
			return false
		}
		req, err := p.advancer(p.current)
		if err != nil {
			p.err = err
			return false
		}
		p.request = req
	}
	resp := p.resp
	var err error
	if resp == nil {
		resp, err = p.pipeline.Do(ctx, p.request)
	} else {
		p.resp = nil
	}
	if err != nil {
		p.err = err
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *productResultPager) PageResponse() *ProductResultResponse {
	return p.current
}

// ProductResultValuePager provides iteration over ProductResultValue pages.
type ProductResultValuePager interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current ProductResultValueResponse.
	PageResponse() *ProductResultValueResponse

	// Err returns the last error encountered while paging.
	Err() error
}

type productResultValueHandleResponse func(*azcore.Response) (*ProductResultValueResponse, error)

type productResultValueAdvancePage func(*ProductResultValueResponse) (*azcore.Request, error)

type productResultValuePager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// contains the pending request
	request *azcore.Request
	// callback for handling the HTTP response
	responder productResultValueHandleResponse
	// callback for advancing to the next page
	advancer productResultValueAdvancePage
	// contains the current response
	current *ProductResultValueResponse
	// any error encountered
	err error
}

func (p *productResultValuePager) Err() error {
	return p.err
}

func (p *productResultValuePager) NextPage(ctx context.Context) bool {
	if p.current != nil {
		if p.current.ProductResultValue.NextLink == nil || len(*p.current.ProductResultValue.NextLink) == 0 {
			return false
		}
		req, err := p.advancer(p.current)
		if err != nil {
			p.err = err
			return false
		}
		p.request = req
	}
	resp, err := p.pipeline.Do(ctx, p.request)
	if err != nil {
		p.err = err
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *productResultValuePager) PageResponse() *ProductResultValueResponse {
	return p.current
}

// ProductResultValueWithXmsClientNamePager provides iteration over ProductResultValueWithXmsClientName pages.
type ProductResultValueWithXmsClientNamePager interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current ProductResultValueWithXmsClientNameResponse.
	PageResponse() *ProductResultValueWithXmsClientNameResponse

	// Err returns the last error encountered while paging.
	Err() error
}

type productResultValueWithXmsClientNameHandleResponse func(*azcore.Response) (*ProductResultValueWithXmsClientNameResponse, error)

type productResultValueWithXmsClientNameAdvancePage func(*ProductResultValueWithXmsClientNameResponse) (*azcore.Request, error)

type productResultValueWithXmsClientNamePager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// contains the pending request
	request *azcore.Request
	// callback for handling the HTTP response
	responder productResultValueWithXmsClientNameHandleResponse
	// callback for advancing to the next page
	advancer productResultValueWithXmsClientNameAdvancePage
	// contains the current response
	current *ProductResultValueWithXmsClientNameResponse
	// any error encountered
	err error
}

func (p *productResultValueWithXmsClientNamePager) Err() error {
	return p.err
}

func (p *productResultValueWithXmsClientNamePager) NextPage(ctx context.Context) bool {
	if p.current != nil {
		if p.current.ProductResultValueWithXmsClientName.NextLink == nil || len(*p.current.ProductResultValueWithXmsClientName.NextLink) == 0 {
			return false
		}
		req, err := p.advancer(p.current)
		if err != nil {
			p.err = err
			return false
		}
		p.request = req
	}
	resp, err := p.pipeline.Do(ctx, p.request)
	if err != nil {
		p.err = err
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *productResultValueWithXmsClientNamePager) PageResponse() *ProductResultValueWithXmsClientNameResponse {
	return p.current
}