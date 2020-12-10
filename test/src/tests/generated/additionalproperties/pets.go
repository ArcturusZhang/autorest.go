package additionalproperties

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// PetsClient is the test Infrastructure for AutoRest
type PetsClient struct {
	BaseClient
}

// NewPetsClient creates an instance of the PetsClient client.
func NewPetsClient() PetsClient {
	return NewPetsClientWithBaseURI(DefaultBaseURI)
}

// NewPetsClientWithBaseURI creates an instance of the PetsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewPetsClientWithBaseURI(baseURI string) PetsClient {
	return PetsClient{NewWithBaseURI(baseURI)}
}

// CreateAPInProperties create a Pet which contains more properties than what is defined.
func (client PetsClient) CreateAPInProperties(ctx context.Context, createParameters PetAPInProperties) (result PetAPInProperties, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PetsClient.CreateAPInProperties")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: createParameters,
			Constraints: []validation.Constraint{{Target: "createParameters.ID", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("additionalproperties.PetsClient", "CreateAPInProperties", err.Error())
	}

	req, err := client.CreateAPInPropertiesPreparer(ctx, createParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "additionalproperties.PetsClient", "CreateAPInProperties", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateAPInPropertiesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "additionalproperties.PetsClient", "CreateAPInProperties", resp, "Failure sending request")
		return
	}

	result, err = client.CreateAPInPropertiesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "additionalproperties.PetsClient", "CreateAPInProperties", resp, "Failure responding to request")
		return
	}

	return
}

// CreateAPInPropertiesPreparer prepares the CreateAPInProperties request.
func (client PetsClient) CreateAPInPropertiesPreparer(ctx context.Context, createParameters PetAPInProperties) (*http.Request, error) {
	createParameters.Status = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/additionalProperties/in/properties"),
		autorest.WithJSON(createParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateAPInPropertiesSender sends the CreateAPInProperties request. The method will close the
// http.Response Body if it receives an error.
func (client PetsClient) CreateAPInPropertiesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateAPInPropertiesResponder handles the response to the CreateAPInProperties request. The method always
// closes the http.Response Body.
func (client PetsClient) CreateAPInPropertiesResponder(resp *http.Response) (result PetAPInProperties, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateAPInPropertiesWithAPString create a Pet which contains more properties than what is defined.
func (client PetsClient) CreateAPInPropertiesWithAPString(ctx context.Context, createParameters PetAPInPropertiesWithAPString) (result PetAPInPropertiesWithAPString, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PetsClient.CreateAPInPropertiesWithAPString")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: createParameters,
			Constraints: []validation.Constraint{{Target: "createParameters.ID", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "createParameters.OdataLocation", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("additionalproperties.PetsClient", "CreateAPInPropertiesWithAPString", err.Error())
	}

	req, err := client.CreateAPInPropertiesWithAPStringPreparer(ctx, createParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "additionalproperties.PetsClient", "CreateAPInPropertiesWithAPString", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateAPInPropertiesWithAPStringSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "additionalproperties.PetsClient", "CreateAPInPropertiesWithAPString", resp, "Failure sending request")
		return
	}

	result, err = client.CreateAPInPropertiesWithAPStringResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "additionalproperties.PetsClient", "CreateAPInPropertiesWithAPString", resp, "Failure responding to request")
		return
	}

	return
}

// CreateAPInPropertiesWithAPStringPreparer prepares the CreateAPInPropertiesWithAPString request.
func (client PetsClient) CreateAPInPropertiesWithAPStringPreparer(ctx context.Context, createParameters PetAPInPropertiesWithAPString) (*http.Request, error) {
	createParameters.Status = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/additionalProperties/in/properties/with/additionalProperties/string"),
		autorest.WithJSON(createParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateAPInPropertiesWithAPStringSender sends the CreateAPInPropertiesWithAPString request. The method will close the
// http.Response Body if it receives an error.
func (client PetsClient) CreateAPInPropertiesWithAPStringSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateAPInPropertiesWithAPStringResponder handles the response to the CreateAPInPropertiesWithAPString request. The method always
// closes the http.Response Body.
func (client PetsClient) CreateAPInPropertiesWithAPStringResponder(resp *http.Response) (result PetAPInPropertiesWithAPString, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateAPObject create a Pet which contains more properties than what is defined.
func (client PetsClient) CreateAPObject(ctx context.Context, createParameters PetAPObject) (result PetAPObject, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PetsClient.CreateAPObject")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: createParameters,
			Constraints: []validation.Constraint{{Target: "createParameters.ID", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("additionalproperties.PetsClient", "CreateAPObject", err.Error())
	}

	req, err := client.CreateAPObjectPreparer(ctx, createParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "additionalproperties.PetsClient", "CreateAPObject", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateAPObjectSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "additionalproperties.PetsClient", "CreateAPObject", resp, "Failure sending request")
		return
	}

	result, err = client.CreateAPObjectResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "additionalproperties.PetsClient", "CreateAPObject", resp, "Failure responding to request")
		return
	}

	return
}

// CreateAPObjectPreparer prepares the CreateAPObject request.
func (client PetsClient) CreateAPObjectPreparer(ctx context.Context, createParameters PetAPObject) (*http.Request, error) {
	createParameters.Status = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/additionalProperties/type/object"),
		autorest.WithJSON(createParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateAPObjectSender sends the CreateAPObject request. The method will close the
// http.Response Body if it receives an error.
func (client PetsClient) CreateAPObjectSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateAPObjectResponder handles the response to the CreateAPObject request. The method always
// closes the http.Response Body.
func (client PetsClient) CreateAPObjectResponder(resp *http.Response) (result PetAPObject, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateAPString create a Pet which contains more properties than what is defined.
func (client PetsClient) CreateAPString(ctx context.Context, createParameters PetAPString) (result PetAPString, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PetsClient.CreateAPString")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: createParameters,
			Constraints: []validation.Constraint{{Target: "createParameters.ID", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("additionalproperties.PetsClient", "CreateAPString", err.Error())
	}

	req, err := client.CreateAPStringPreparer(ctx, createParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "additionalproperties.PetsClient", "CreateAPString", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateAPStringSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "additionalproperties.PetsClient", "CreateAPString", resp, "Failure sending request")
		return
	}

	result, err = client.CreateAPStringResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "additionalproperties.PetsClient", "CreateAPString", resp, "Failure responding to request")
		return
	}

	return
}

// CreateAPStringPreparer prepares the CreateAPString request.
func (client PetsClient) CreateAPStringPreparer(ctx context.Context, createParameters PetAPString) (*http.Request, error) {
	createParameters.Status = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/additionalProperties/type/string"),
		autorest.WithJSON(createParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateAPStringSender sends the CreateAPString request. The method will close the
// http.Response Body if it receives an error.
func (client PetsClient) CreateAPStringSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateAPStringResponder handles the response to the CreateAPString request. The method always
// closes the http.Response Body.
func (client PetsClient) CreateAPStringResponder(resp *http.Response) (result PetAPString, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateAPTrue create a Pet which contains more properties than what is defined.
func (client PetsClient) CreateAPTrue(ctx context.Context, createParameters PetAPTrue) (result PetAPTrue, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PetsClient.CreateAPTrue")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: createParameters,
			Constraints: []validation.Constraint{{Target: "createParameters.ID", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("additionalproperties.PetsClient", "CreateAPTrue", err.Error())
	}

	req, err := client.CreateAPTruePreparer(ctx, createParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "additionalproperties.PetsClient", "CreateAPTrue", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateAPTrueSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "additionalproperties.PetsClient", "CreateAPTrue", resp, "Failure sending request")
		return
	}

	result, err = client.CreateAPTrueResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "additionalproperties.PetsClient", "CreateAPTrue", resp, "Failure responding to request")
		return
	}

	return
}

// CreateAPTruePreparer prepares the CreateAPTrue request.
func (client PetsClient) CreateAPTruePreparer(ctx context.Context, createParameters PetAPTrue) (*http.Request, error) {
	createParameters.Status = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/additionalProperties/true"),
		autorest.WithJSON(createParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateAPTrueSender sends the CreateAPTrue request. The method will close the
// http.Response Body if it receives an error.
func (client PetsClient) CreateAPTrueSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateAPTrueResponder handles the response to the CreateAPTrue request. The method always
// closes the http.Response Body.
func (client PetsClient) CreateAPTrueResponder(resp *http.Response) (result PetAPTrue, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateCatAPTrue create a CatAPTrue which contains more properties than what is defined.
func (client PetsClient) CreateCatAPTrue(ctx context.Context, createParameters CatAPTrue) (result CatAPTrue, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PetsClient.CreateCatAPTrue")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateCatAPTruePreparer(ctx, createParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "additionalproperties.PetsClient", "CreateCatAPTrue", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateCatAPTrueSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "additionalproperties.PetsClient", "CreateCatAPTrue", resp, "Failure sending request")
		return
	}

	result, err = client.CreateCatAPTrueResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "additionalproperties.PetsClient", "CreateCatAPTrue", resp, "Failure responding to request")
		return
	}

	return
}

// CreateCatAPTruePreparer prepares the CreateCatAPTrue request.
func (client PetsClient) CreateCatAPTruePreparer(ctx context.Context, createParameters CatAPTrue) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/additionalProperties/true-subclass"),
		autorest.WithJSON(createParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateCatAPTrueSender sends the CreateCatAPTrue request. The method will close the
// http.Response Body if it receives an error.
func (client PetsClient) CreateCatAPTrueSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateCatAPTrueResponder handles the response to the CreateCatAPTrue request. The method always
// closes the http.Response Body.
func (client PetsClient) CreateCatAPTrueResponder(resp *http.Response) (result CatAPTrue, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
