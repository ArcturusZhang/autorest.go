// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package reportgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// AutoRestReportServiceClient contains the methods for the AutoRestReportService group.
// Don't use this type directly, use NewAutoRestReportServiceClient() instead.
type AutoRestReportServiceClient struct {
	con *Connection
}

// NewAutoRestReportServiceClient creates a new instance of AutoRestReportServiceClient with the specified values.
func NewAutoRestReportServiceClient(con *Connection) *AutoRestReportServiceClient {
	return &AutoRestReportServiceClient{con: con}
}

// GetOptionalReport - Get optional test coverage report
func (client *AutoRestReportServiceClient) GetOptionalReport(ctx context.Context, options *AutoRestReportServiceGetOptionalReportOptions) (MapOfInt32Response, error) {
	req, err := client.getOptionalReportCreateRequest(ctx, options)
	if err != nil {
		return MapOfInt32Response{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return MapOfInt32Response{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return MapOfInt32Response{}, client.getOptionalReportHandleError(resp)
	}
	return client.getOptionalReportHandleResponse(resp)
}

// getOptionalReportCreateRequest creates the GetOptionalReport request.
func (client *AutoRestReportServiceClient) getOptionalReportCreateRequest(ctx context.Context, options *AutoRestReportServiceGetOptionalReportOptions) (*azcore.Request, error) {
	urlPath := "/report/optional"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	if options != nil && options.Qualifier != nil {
		query.Set("qualifier", *options.Qualifier)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getOptionalReportHandleResponse handles the GetOptionalReport response.
func (client *AutoRestReportServiceClient) getOptionalReportHandleResponse(resp *azcore.Response) (MapOfInt32Response, error) {
	var val map[string]int32
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return MapOfInt32Response{}, err
	}
	return MapOfInt32Response{RawResponse: resp.Response, Value: val}, nil
}

// getOptionalReportHandleError handles the GetOptionalReport error response.
func (client *AutoRestReportServiceClient) getOptionalReportHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetReport - Get test coverage report
func (client *AutoRestReportServiceClient) GetReport(ctx context.Context, options *AutoRestReportServiceGetReportOptions) (MapOfInt32Response, error) {
	req, err := client.getReportCreateRequest(ctx, options)
	if err != nil {
		return MapOfInt32Response{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return MapOfInt32Response{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return MapOfInt32Response{}, client.getReportHandleError(resp)
	}
	return client.getReportHandleResponse(resp)
}

// getReportCreateRequest creates the GetReport request.
func (client *AutoRestReportServiceClient) getReportCreateRequest(ctx context.Context, options *AutoRestReportServiceGetReportOptions) (*azcore.Request, error) {
	urlPath := "/report"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	if options != nil && options.Qualifier != nil {
		query.Set("qualifier", *options.Qualifier)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getReportHandleResponse handles the GetReport response.
func (client *AutoRestReportServiceClient) getReportHandleResponse(resp *azcore.Response) (MapOfInt32Response, error) {
	var val map[string]int32
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return MapOfInt32Response{}, err
	}
	return MapOfInt32Response{RawResponse: resp.Response, Value: val}, nil
}

// getReportHandleError handles the GetReport error response.
func (client *AutoRestReportServiceClient) getReportHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
