package complexgroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// InheritanceClient is the test Infrastructure for AutoRest
type InheritanceClient struct {
	BaseClient
}

// NewInheritanceClient creates an instance of the InheritanceClient client.
func NewInheritanceClient() InheritanceClient {
	return NewInheritanceClientWithBaseURI(DefaultBaseURI)
}

// NewInheritanceClientWithBaseURI creates an instance of the InheritanceClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewInheritanceClientWithBaseURI(baseURI string) InheritanceClient {
	return InheritanceClient{NewWithBaseURI(baseURI)}
}

// GetValid get complex types that extend others
func (client InheritanceClient) GetValid(ctx context.Context) (result Siamese, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InheritanceClient.GetValid")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetValidPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.InheritanceClient", "GetValid", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetValidSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "complexgroup.InheritanceClient", "GetValid", resp, "Failure sending request")
		return
	}

	result, err = client.GetValidResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.InheritanceClient", "GetValid", resp, "Failure responding to request")
		return
	}

	return
}

// GetValidPreparer prepares the GetValid request.
func (client InheritanceClient) GetValidPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/inheritance/valid"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetValidSender sends the GetValid request. The method will close the
// http.Response Body if it receives an error.
func (client InheritanceClient) GetValidSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetValidResponder handles the response to the GetValid request. The method always
// closes the http.Response Body.
func (client InheritanceClient) GetValidResponder(resp *http.Response) (result Siamese, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutValid put complex types that extend others
// Parameters:
// complexBody - please put a siamese with id=2, name="Siameee", color=green, breed=persion, which hates 2
// dogs, the 1st one named "Potato" with id=1 and food="tomato", and the 2nd one named "Tomato" with id=-1 and
// food="french fries".
func (client InheritanceClient) PutValid(ctx context.Context, complexBody Siamese) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InheritanceClient.PutValid")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PutValidPreparer(ctx, complexBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.InheritanceClient", "PutValid", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutValidSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "complexgroup.InheritanceClient", "PutValid", resp, "Failure sending request")
		return
	}

	result, err = client.PutValidResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.InheritanceClient", "PutValid", resp, "Failure responding to request")
		return
	}

	return
}

// PutValidPreparer prepares the PutValid request.
func (client InheritanceClient) PutValidPreparer(ctx context.Context, complexBody Siamese) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/inheritance/valid"),
		autorest.WithJSON(complexBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutValidSender sends the PutValid request. The method will close the
// http.Response Body if it receives an error.
func (client InheritanceClient) PutValidSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutValidResponder handles the response to the PutValid request. The method always
// closes the http.Response Body.
func (client InheritanceClient) PutValidResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
