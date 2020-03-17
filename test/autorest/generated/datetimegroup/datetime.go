// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package datetimegroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
	"time"
)

// DatetimeOperations contains the methods for the Datetime group.
type DatetimeOperations interface {
	// GetInvalid - Get invalid datetime value
	GetInvalid(ctx context.Context) (*DatetimeGetInvalidResponse, error)
	// GetLocalNegativeOffsetLowercaseMaxDateTime - Get max datetime value with positive num offset 9999-12-31t23:59:59.999-14:00
	GetLocalNegativeOffsetLowercaseMaxDateTime(ctx context.Context) (*DatetimeGetLocalNegativeOffsetLowercaseMaxDateTimeResponse, error)
	// GetLocalNegativeOffsetMinDateTime - Get min datetime value 0001-01-01T00:00:00-14:00
	GetLocalNegativeOffsetMinDateTime(ctx context.Context) (*DatetimeGetLocalNegativeOffsetMinDateTimeResponse, error)
	// GetLocalNegativeOffsetUppercaseMaxDateTime - Get max datetime value with positive num offset 9999-12-31T23:59:59.999-14:00
	GetLocalNegativeOffsetUppercaseMaxDateTime(ctx context.Context) (*DatetimeGetLocalNegativeOffsetUppercaseMaxDateTimeResponse, error)
	// GetLocalPositiveOffsetLowercaseMaxDateTime - Get max datetime value with positive num offset 9999-12-31t23:59:59.999+14:00
	GetLocalPositiveOffsetLowercaseMaxDateTime(ctx context.Context) (*DatetimeGetLocalPositiveOffsetLowercaseMaxDateTimeResponse, error)
	// GetLocalPositiveOffsetMinDateTime - Get min datetime value 0001-01-01T00:00:00+14:00
	GetLocalPositiveOffsetMinDateTime(ctx context.Context) (*DatetimeGetLocalPositiveOffsetMinDateTimeResponse, error)
	// GetLocalPositiveOffsetUppercaseMaxDateTime - Get max datetime value with positive num offset 9999-12-31T23:59:59.999+14:00
	GetLocalPositiveOffsetUppercaseMaxDateTime(ctx context.Context) (*DatetimeGetLocalPositiveOffsetUppercaseMaxDateTimeResponse, error)
	// GetNull - Get null datetime value
	GetNull(ctx context.Context) (*DatetimeGetNullResponse, error)
	// GetOverflow - Get overflow datetime value
	GetOverflow(ctx context.Context) (*DatetimeGetOverflowResponse, error)
	// GetUTCLowercaseMaxDateTime - Get max datetime value 9999-12-31t23:59:59.999z
	GetUTCLowercaseMaxDateTime(ctx context.Context) (*DatetimeGetUTCLowercaseMaxDateTimeResponse, error)
	// GetUTCMinDateTime - Get min datetime value 0001-01-01T00:00:00Z
	GetUTCMinDateTime(ctx context.Context) (*DatetimeGetUTCMinDateTimeResponse, error)
	// GetUTCUppercaseMaxDateTime - Get max datetime value 9999-12-31T23:59:59.999Z
	GetUTCUppercaseMaxDateTime(ctx context.Context) (*DatetimeGetUTCUppercaseMaxDateTimeResponse, error)
	// GetUTCUppercaseMaxDateTime7Digits - Get max datetime value 9999-12-31T23:59:59.9999999Z
	GetUTCUppercaseMaxDateTime7Digits(ctx context.Context) (*DatetimeGetUTCUppercaseMaxDateTime7DigitsResponse, error)
	// GetUnderflow - Get underflow datetime value
	GetUnderflow(ctx context.Context) (*DatetimeGetUnderflowResponse, error)
	// PutLocalNegativeOffsetMaxDateTime - Put max datetime value with positive numoffset 9999-12-31t23:59:59.999-14:00
	PutLocalNegativeOffsetMaxDateTime(ctx context.Context, datetimeBody time.Time) (*DatetimePutLocalNegativeOffsetMaxDateTimeResponse, error)
	// PutLocalNegativeOffsetMinDateTime - Put min datetime value 0001-01-01T00:00:00-14:00
	PutLocalNegativeOffsetMinDateTime(ctx context.Context, datetimeBody time.Time) (*DatetimePutLocalNegativeOffsetMinDateTimeResponse, error)
	// PutLocalPositiveOffsetMaxDateTime - Put max datetime value with positive numoffset 9999-12-31t23:59:59.999+14:00
	PutLocalPositiveOffsetMaxDateTime(ctx context.Context, datetimeBody time.Time) (*DatetimePutLocalPositiveOffsetMaxDateTimeResponse, error)
	// PutLocalPositiveOffsetMinDateTime - Put min datetime value 0001-01-01T00:00:00+14:00
	PutLocalPositiveOffsetMinDateTime(ctx context.Context, datetimeBody time.Time) (*DatetimePutLocalPositiveOffsetMinDateTimeResponse, error)
	// PutUTCMaxDateTime - Put max datetime value 9999-12-31T23:59:59.999Z
	PutUTCMaxDateTime(ctx context.Context, datetimeBody time.Time) (*DatetimePutUTCMaxDateTimeResponse, error)
	// PutUTCMaxDateTime7Digits - Put max datetime value 9999-12-31T23:59:59.9999999Z
	PutUTCMaxDateTime7Digits(ctx context.Context, datetimeBody time.Time) (*DatetimePutUTCMaxDateTime7DigitsResponse, error)
	// PutUTCMinDateTime - Put min datetime value 0001-01-01T00:00:00Z
	PutUTCMinDateTime(ctx context.Context, datetimeBody time.Time) (*DatetimePutUTCMinDateTimeResponse, error)
}

// datetimeOperations implements the DatetimeOperations interface.
type datetimeOperations struct {
	*Client
}

// GetInvalid - Get invalid datetime value
func (client *datetimeOperations) GetInvalid(ctx context.Context) (*DatetimeGetInvalidResponse, error) {
	req, err := client.getInvalidCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getInvalidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getInvalidCreateRequest creates the GetInvalid request.
func (client *datetimeOperations) getInvalidCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/datetime/invalid"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// getInvalidHandleResponse handles the GetInvalid response.
func (client *datetimeOperations) getInvalidHandleResponse(resp *azcore.Response) (*DatetimeGetInvalidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	result := DatetimeGetInvalidResponse{RawResponse: resp.Response}
	result.Value = (*time.Time)(aux)
	return &result, err
}

// GetLocalNegativeOffsetLowercaseMaxDateTime - Get max datetime value with positive num offset 9999-12-31t23:59:59.999-14:00
func (client *datetimeOperations) GetLocalNegativeOffsetLowercaseMaxDateTime(ctx context.Context) (*DatetimeGetLocalNegativeOffsetLowercaseMaxDateTimeResponse, error) {
	req, err := client.getLocalNegativeOffsetLowercaseMaxDateTimeCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getLocalNegativeOffsetLowercaseMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getLocalNegativeOffsetLowercaseMaxDateTimeCreateRequest creates the GetLocalNegativeOffsetLowercaseMaxDateTime request.
func (client *datetimeOperations) getLocalNegativeOffsetLowercaseMaxDateTimeCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/datetime/max/localnegativeoffset/lowercase"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// getLocalNegativeOffsetLowercaseMaxDateTimeHandleResponse handles the GetLocalNegativeOffsetLowercaseMaxDateTime response.
func (client *datetimeOperations) getLocalNegativeOffsetLowercaseMaxDateTimeHandleResponse(resp *azcore.Response) (*DatetimeGetLocalNegativeOffsetLowercaseMaxDateTimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	result := DatetimeGetLocalNegativeOffsetLowercaseMaxDateTimeResponse{RawResponse: resp.Response}
	result.Value = (*time.Time)(aux)
	return &result, err
}

// GetLocalNegativeOffsetMinDateTime - Get min datetime value 0001-01-01T00:00:00-14:00
func (client *datetimeOperations) GetLocalNegativeOffsetMinDateTime(ctx context.Context) (*DatetimeGetLocalNegativeOffsetMinDateTimeResponse, error) {
	req, err := client.getLocalNegativeOffsetMinDateTimeCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getLocalNegativeOffsetMinDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getLocalNegativeOffsetMinDateTimeCreateRequest creates the GetLocalNegativeOffsetMinDateTime request.
func (client *datetimeOperations) getLocalNegativeOffsetMinDateTimeCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/datetime/min/localnegativeoffset"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// getLocalNegativeOffsetMinDateTimeHandleResponse handles the GetLocalNegativeOffsetMinDateTime response.
func (client *datetimeOperations) getLocalNegativeOffsetMinDateTimeHandleResponse(resp *azcore.Response) (*DatetimeGetLocalNegativeOffsetMinDateTimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	result := DatetimeGetLocalNegativeOffsetMinDateTimeResponse{RawResponse: resp.Response}
	result.Value = (*time.Time)(aux)
	return &result, err
}

// GetLocalNegativeOffsetUppercaseMaxDateTime - Get max datetime value with positive num offset 9999-12-31T23:59:59.999-14:00
func (client *datetimeOperations) GetLocalNegativeOffsetUppercaseMaxDateTime(ctx context.Context) (*DatetimeGetLocalNegativeOffsetUppercaseMaxDateTimeResponse, error) {
	req, err := client.getLocalNegativeOffsetUppercaseMaxDateTimeCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getLocalNegativeOffsetUppercaseMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getLocalNegativeOffsetUppercaseMaxDateTimeCreateRequest creates the GetLocalNegativeOffsetUppercaseMaxDateTime request.
func (client *datetimeOperations) getLocalNegativeOffsetUppercaseMaxDateTimeCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/datetime/max/localnegativeoffset/uppercase"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// getLocalNegativeOffsetUppercaseMaxDateTimeHandleResponse handles the GetLocalNegativeOffsetUppercaseMaxDateTime response.
func (client *datetimeOperations) getLocalNegativeOffsetUppercaseMaxDateTimeHandleResponse(resp *azcore.Response) (*DatetimeGetLocalNegativeOffsetUppercaseMaxDateTimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	result := DatetimeGetLocalNegativeOffsetUppercaseMaxDateTimeResponse{RawResponse: resp.Response}
	result.Value = (*time.Time)(aux)
	return &result, err
}

// GetLocalPositiveOffsetLowercaseMaxDateTime - Get max datetime value with positive num offset 9999-12-31t23:59:59.999+14:00
func (client *datetimeOperations) GetLocalPositiveOffsetLowercaseMaxDateTime(ctx context.Context) (*DatetimeGetLocalPositiveOffsetLowercaseMaxDateTimeResponse, error) {
	req, err := client.getLocalPositiveOffsetLowercaseMaxDateTimeCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getLocalPositiveOffsetLowercaseMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getLocalPositiveOffsetLowercaseMaxDateTimeCreateRequest creates the GetLocalPositiveOffsetLowercaseMaxDateTime request.
func (client *datetimeOperations) getLocalPositiveOffsetLowercaseMaxDateTimeCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/datetime/max/localpositiveoffset/lowercase"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// getLocalPositiveOffsetLowercaseMaxDateTimeHandleResponse handles the GetLocalPositiveOffsetLowercaseMaxDateTime response.
func (client *datetimeOperations) getLocalPositiveOffsetLowercaseMaxDateTimeHandleResponse(resp *azcore.Response) (*DatetimeGetLocalPositiveOffsetLowercaseMaxDateTimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	result := DatetimeGetLocalPositiveOffsetLowercaseMaxDateTimeResponse{RawResponse: resp.Response}
	result.Value = (*time.Time)(aux)
	return &result, err
}

// GetLocalPositiveOffsetMinDateTime - Get min datetime value 0001-01-01T00:00:00+14:00
func (client *datetimeOperations) GetLocalPositiveOffsetMinDateTime(ctx context.Context) (*DatetimeGetLocalPositiveOffsetMinDateTimeResponse, error) {
	req, err := client.getLocalPositiveOffsetMinDateTimeCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getLocalPositiveOffsetMinDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getLocalPositiveOffsetMinDateTimeCreateRequest creates the GetLocalPositiveOffsetMinDateTime request.
func (client *datetimeOperations) getLocalPositiveOffsetMinDateTimeCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/datetime/min/localpositiveoffset"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// getLocalPositiveOffsetMinDateTimeHandleResponse handles the GetLocalPositiveOffsetMinDateTime response.
func (client *datetimeOperations) getLocalPositiveOffsetMinDateTimeHandleResponse(resp *azcore.Response) (*DatetimeGetLocalPositiveOffsetMinDateTimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	result := DatetimeGetLocalPositiveOffsetMinDateTimeResponse{RawResponse: resp.Response}
	result.Value = (*time.Time)(aux)
	return &result, err
}

// GetLocalPositiveOffsetUppercaseMaxDateTime - Get max datetime value with positive num offset 9999-12-31T23:59:59.999+14:00
func (client *datetimeOperations) GetLocalPositiveOffsetUppercaseMaxDateTime(ctx context.Context) (*DatetimeGetLocalPositiveOffsetUppercaseMaxDateTimeResponse, error) {
	req, err := client.getLocalPositiveOffsetUppercaseMaxDateTimeCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getLocalPositiveOffsetUppercaseMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getLocalPositiveOffsetUppercaseMaxDateTimeCreateRequest creates the GetLocalPositiveOffsetUppercaseMaxDateTime request.
func (client *datetimeOperations) getLocalPositiveOffsetUppercaseMaxDateTimeCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/datetime/max/localpositiveoffset/uppercase"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// getLocalPositiveOffsetUppercaseMaxDateTimeHandleResponse handles the GetLocalPositiveOffsetUppercaseMaxDateTime response.
func (client *datetimeOperations) getLocalPositiveOffsetUppercaseMaxDateTimeHandleResponse(resp *azcore.Response) (*DatetimeGetLocalPositiveOffsetUppercaseMaxDateTimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	result := DatetimeGetLocalPositiveOffsetUppercaseMaxDateTimeResponse{RawResponse: resp.Response}
	result.Value = (*time.Time)(aux)
	return &result, err
}

// GetNull - Get null datetime value
func (client *datetimeOperations) GetNull(ctx context.Context) (*DatetimeGetNullResponse, error) {
	req, err := client.getNullCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getNullCreateRequest creates the GetNull request.
func (client *datetimeOperations) getNullCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/datetime/null"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *datetimeOperations) getNullHandleResponse(resp *azcore.Response) (*DatetimeGetNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	result := DatetimeGetNullResponse{RawResponse: resp.Response}
	result.Value = (*time.Time)(aux)
	return &result, err
}

// GetOverflow - Get overflow datetime value
func (client *datetimeOperations) GetOverflow(ctx context.Context) (*DatetimeGetOverflowResponse, error) {
	req, err := client.getOverflowCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getOverflowHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getOverflowCreateRequest creates the GetOverflow request.
func (client *datetimeOperations) getOverflowCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/datetime/overflow"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// getOverflowHandleResponse handles the GetOverflow response.
func (client *datetimeOperations) getOverflowHandleResponse(resp *azcore.Response) (*DatetimeGetOverflowResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	result := DatetimeGetOverflowResponse{RawResponse: resp.Response}
	result.Value = (*time.Time)(aux)
	return &result, err
}

// GetUTCLowercaseMaxDateTime - Get max datetime value 9999-12-31t23:59:59.999z
func (client *datetimeOperations) GetUTCLowercaseMaxDateTime(ctx context.Context) (*DatetimeGetUTCLowercaseMaxDateTimeResponse, error) {
	req, err := client.getUtcLowercaseMaxDateTimeCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getUtcLowercaseMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getUtcLowercaseMaxDateTimeCreateRequest creates the GetUTCLowercaseMaxDateTime request.
func (client *datetimeOperations) getUtcLowercaseMaxDateTimeCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/datetime/max/utc/lowercase"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// getUtcLowercaseMaxDateTimeHandleResponse handles the GetUTCLowercaseMaxDateTime response.
func (client *datetimeOperations) getUtcLowercaseMaxDateTimeHandleResponse(resp *azcore.Response) (*DatetimeGetUTCLowercaseMaxDateTimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	result := DatetimeGetUTCLowercaseMaxDateTimeResponse{RawResponse: resp.Response}
	result.Value = (*time.Time)(aux)
	return &result, err
}

// GetUTCMinDateTime - Get min datetime value 0001-01-01T00:00:00Z
func (client *datetimeOperations) GetUTCMinDateTime(ctx context.Context) (*DatetimeGetUTCMinDateTimeResponse, error) {
	req, err := client.getUtcMinDateTimeCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getUtcMinDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getUtcMinDateTimeCreateRequest creates the GetUTCMinDateTime request.
func (client *datetimeOperations) getUtcMinDateTimeCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/datetime/min/utc"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// getUtcMinDateTimeHandleResponse handles the GetUTCMinDateTime response.
func (client *datetimeOperations) getUtcMinDateTimeHandleResponse(resp *azcore.Response) (*DatetimeGetUTCMinDateTimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	result := DatetimeGetUTCMinDateTimeResponse{RawResponse: resp.Response}
	result.Value = (*time.Time)(aux)
	return &result, err
}

// GetUTCUppercaseMaxDateTime - Get max datetime value 9999-12-31T23:59:59.999Z
func (client *datetimeOperations) GetUTCUppercaseMaxDateTime(ctx context.Context) (*DatetimeGetUTCUppercaseMaxDateTimeResponse, error) {
	req, err := client.getUtcUppercaseMaxDateTimeCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getUtcUppercaseMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getUtcUppercaseMaxDateTimeCreateRequest creates the GetUTCUppercaseMaxDateTime request.
func (client *datetimeOperations) getUtcUppercaseMaxDateTimeCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/datetime/max/utc/uppercase"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// getUtcUppercaseMaxDateTimeHandleResponse handles the GetUTCUppercaseMaxDateTime response.
func (client *datetimeOperations) getUtcUppercaseMaxDateTimeHandleResponse(resp *azcore.Response) (*DatetimeGetUTCUppercaseMaxDateTimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	result := DatetimeGetUTCUppercaseMaxDateTimeResponse{RawResponse: resp.Response}
	result.Value = (*time.Time)(aux)
	return &result, err
}

// GetUTCUppercaseMaxDateTime7Digits - Get max datetime value 9999-12-31T23:59:59.9999999Z
func (client *datetimeOperations) GetUTCUppercaseMaxDateTime7Digits(ctx context.Context) (*DatetimeGetUTCUppercaseMaxDateTime7DigitsResponse, error) {
	req, err := client.getUtcUppercaseMaxDateTime7DigitsCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getUtcUppercaseMaxDateTime7DigitsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getUtcUppercaseMaxDateTime7DigitsCreateRequest creates the GetUTCUppercaseMaxDateTime7Digits request.
func (client *datetimeOperations) getUtcUppercaseMaxDateTime7DigitsCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/datetime/max/utc7ms/uppercase"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// getUtcUppercaseMaxDateTime7DigitsHandleResponse handles the GetUTCUppercaseMaxDateTime7Digits response.
func (client *datetimeOperations) getUtcUppercaseMaxDateTime7DigitsHandleResponse(resp *azcore.Response) (*DatetimeGetUTCUppercaseMaxDateTime7DigitsResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	result := DatetimeGetUTCUppercaseMaxDateTime7DigitsResponse{RawResponse: resp.Response}
	result.Value = (*time.Time)(aux)
	return &result, err
}

// GetUnderflow - Get underflow datetime value
func (client *datetimeOperations) GetUnderflow(ctx context.Context) (*DatetimeGetUnderflowResponse, error) {
	req, err := client.getUnderflowCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getUnderflowHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getUnderflowCreateRequest creates the GetUnderflow request.
func (client *datetimeOperations) getUnderflowCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/datetime/underflow"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// getUnderflowHandleResponse handles the GetUnderflow response.
func (client *datetimeOperations) getUnderflowHandleResponse(resp *azcore.Response) (*DatetimeGetUnderflowResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	var aux *timeRFC3339
	err := resp.UnmarshalAsJSON(&aux)
	result := DatetimeGetUnderflowResponse{RawResponse: resp.Response}
	result.Value = (*time.Time)(aux)
	return &result, err
}

// PutLocalNegativeOffsetMaxDateTime - Put max datetime value with positive numoffset 9999-12-31t23:59:59.999-14:00
func (client *datetimeOperations) PutLocalNegativeOffsetMaxDateTime(ctx context.Context, datetimeBody time.Time) (*DatetimePutLocalNegativeOffsetMaxDateTimeResponse, error) {
	req, err := client.putLocalNegativeOffsetMaxDateTimeCreateRequest(*client.u, datetimeBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putLocalNegativeOffsetMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putLocalNegativeOffsetMaxDateTimeCreateRequest creates the PutLocalNegativeOffsetMaxDateTime request.
func (client *datetimeOperations) putLocalNegativeOffsetMaxDateTimeCreateRequest(u url.URL, datetimeBody time.Time) (*azcore.Request, error) {
	urlPath := "/datetime/max/localnegativeoffset"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(datetimeBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// putLocalNegativeOffsetMaxDateTimeHandleResponse handles the PutLocalNegativeOffsetMaxDateTime response.
func (client *datetimeOperations) putLocalNegativeOffsetMaxDateTimeHandleResponse(resp *azcore.Response) (*DatetimePutLocalNegativeOffsetMaxDateTimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &DatetimePutLocalNegativeOffsetMaxDateTimeResponse{RawResponse: resp.Response}, nil
}

// PutLocalNegativeOffsetMinDateTime - Put min datetime value 0001-01-01T00:00:00-14:00
func (client *datetimeOperations) PutLocalNegativeOffsetMinDateTime(ctx context.Context, datetimeBody time.Time) (*DatetimePutLocalNegativeOffsetMinDateTimeResponse, error) {
	req, err := client.putLocalNegativeOffsetMinDateTimeCreateRequest(*client.u, datetimeBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putLocalNegativeOffsetMinDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putLocalNegativeOffsetMinDateTimeCreateRequest creates the PutLocalNegativeOffsetMinDateTime request.
func (client *datetimeOperations) putLocalNegativeOffsetMinDateTimeCreateRequest(u url.URL, datetimeBody time.Time) (*azcore.Request, error) {
	urlPath := "/datetime/min/localnegativeoffset"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(datetimeBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// putLocalNegativeOffsetMinDateTimeHandleResponse handles the PutLocalNegativeOffsetMinDateTime response.
func (client *datetimeOperations) putLocalNegativeOffsetMinDateTimeHandleResponse(resp *azcore.Response) (*DatetimePutLocalNegativeOffsetMinDateTimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &DatetimePutLocalNegativeOffsetMinDateTimeResponse{RawResponse: resp.Response}, nil
}

// PutLocalPositiveOffsetMaxDateTime - Put max datetime value with positive numoffset 9999-12-31t23:59:59.999+14:00
func (client *datetimeOperations) PutLocalPositiveOffsetMaxDateTime(ctx context.Context, datetimeBody time.Time) (*DatetimePutLocalPositiveOffsetMaxDateTimeResponse, error) {
	req, err := client.putLocalPositiveOffsetMaxDateTimeCreateRequest(*client.u, datetimeBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putLocalPositiveOffsetMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putLocalPositiveOffsetMaxDateTimeCreateRequest creates the PutLocalPositiveOffsetMaxDateTime request.
func (client *datetimeOperations) putLocalPositiveOffsetMaxDateTimeCreateRequest(u url.URL, datetimeBody time.Time) (*azcore.Request, error) {
	urlPath := "/datetime/max/localpositiveoffset"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(datetimeBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// putLocalPositiveOffsetMaxDateTimeHandleResponse handles the PutLocalPositiveOffsetMaxDateTime response.
func (client *datetimeOperations) putLocalPositiveOffsetMaxDateTimeHandleResponse(resp *azcore.Response) (*DatetimePutLocalPositiveOffsetMaxDateTimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &DatetimePutLocalPositiveOffsetMaxDateTimeResponse{RawResponse: resp.Response}, nil
}

// PutLocalPositiveOffsetMinDateTime - Put min datetime value 0001-01-01T00:00:00+14:00
func (client *datetimeOperations) PutLocalPositiveOffsetMinDateTime(ctx context.Context, datetimeBody time.Time) (*DatetimePutLocalPositiveOffsetMinDateTimeResponse, error) {
	req, err := client.putLocalPositiveOffsetMinDateTimeCreateRequest(*client.u, datetimeBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putLocalPositiveOffsetMinDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putLocalPositiveOffsetMinDateTimeCreateRequest creates the PutLocalPositiveOffsetMinDateTime request.
func (client *datetimeOperations) putLocalPositiveOffsetMinDateTimeCreateRequest(u url.URL, datetimeBody time.Time) (*azcore.Request, error) {
	urlPath := "/datetime/min/localpositiveoffset"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(datetimeBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// putLocalPositiveOffsetMinDateTimeHandleResponse handles the PutLocalPositiveOffsetMinDateTime response.
func (client *datetimeOperations) putLocalPositiveOffsetMinDateTimeHandleResponse(resp *azcore.Response) (*DatetimePutLocalPositiveOffsetMinDateTimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &DatetimePutLocalPositiveOffsetMinDateTimeResponse{RawResponse: resp.Response}, nil
}

// PutUTCMaxDateTime - Put max datetime value 9999-12-31T23:59:59.999Z
func (client *datetimeOperations) PutUTCMaxDateTime(ctx context.Context, datetimeBody time.Time) (*DatetimePutUTCMaxDateTimeResponse, error) {
	req, err := client.putUtcMaxDateTimeCreateRequest(*client.u, datetimeBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putUtcMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putUtcMaxDateTimeCreateRequest creates the PutUTCMaxDateTime request.
func (client *datetimeOperations) putUtcMaxDateTimeCreateRequest(u url.URL, datetimeBody time.Time) (*azcore.Request, error) {
	urlPath := "/datetime/max/utc"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(datetimeBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// putUtcMaxDateTimeHandleResponse handles the PutUTCMaxDateTime response.
func (client *datetimeOperations) putUtcMaxDateTimeHandleResponse(resp *azcore.Response) (*DatetimePutUTCMaxDateTimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &DatetimePutUTCMaxDateTimeResponse{RawResponse: resp.Response}, nil
}

// PutUTCMaxDateTime7Digits - Put max datetime value 9999-12-31T23:59:59.9999999Z
func (client *datetimeOperations) PutUTCMaxDateTime7Digits(ctx context.Context, datetimeBody time.Time) (*DatetimePutUTCMaxDateTime7DigitsResponse, error) {
	req, err := client.putUtcMaxDateTime7DigitsCreateRequest(*client.u, datetimeBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putUtcMaxDateTime7DigitsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putUtcMaxDateTime7DigitsCreateRequest creates the PutUTCMaxDateTime7Digits request.
func (client *datetimeOperations) putUtcMaxDateTime7DigitsCreateRequest(u url.URL, datetimeBody time.Time) (*azcore.Request, error) {
	urlPath := "/datetime/max/utc7ms"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(datetimeBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// putUtcMaxDateTime7DigitsHandleResponse handles the PutUTCMaxDateTime7Digits response.
func (client *datetimeOperations) putUtcMaxDateTime7DigitsHandleResponse(resp *azcore.Response) (*DatetimePutUTCMaxDateTime7DigitsResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &DatetimePutUTCMaxDateTime7DigitsResponse{RawResponse: resp.Response}, nil
}

// PutUTCMinDateTime - Put min datetime value 0001-01-01T00:00:00Z
func (client *datetimeOperations) PutUTCMinDateTime(ctx context.Context, datetimeBody time.Time) (*DatetimePutUTCMinDateTimeResponse, error) {
	req, err := client.putUtcMinDateTimeCreateRequest(*client.u, datetimeBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.putUtcMinDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// putUtcMinDateTimeCreateRequest creates the PutUTCMinDateTime request.
func (client *datetimeOperations) putUtcMinDateTimeCreateRequest(u url.URL, datetimeBody time.Time) (*azcore.Request, error) {
	urlPath := "/datetime/min/utc"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(datetimeBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// putUtcMinDateTimeHandleResponse handles the PutUTCMinDateTime response.
func (client *datetimeOperations) putUtcMinDateTimeHandleResponse(resp *azcore.Response) (*DatetimePutUTCMinDateTimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &DatetimePutUTCMinDateTimeResponse{RawResponse: resp.Response}, nil
}