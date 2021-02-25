// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrogroup

import (
	"context"
	"errors"
	"net/http"
	"net/http/cookiejar"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/to"
	"github.com/google/go-cmp/cmp"
)

func newLROSClient() *LrOSClient {
	options := ConnectionOptions{}
	options.Retry.RetryDelay = 10 * time.Millisecond
	options.HTTPClient = httpClientWithCookieJar()
	return NewLrOSClient(NewDefaultConnection(&options))
}

func httpClientWithCookieJar() azcore.Transport {
	j, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	http.DefaultClient.Jar = j
	return azcore.TransportFunc(func(req *http.Request) (*http.Response, error) {
		return http.DefaultClient.Do(req)
	})
}

func TestLROResumeWrongPoller(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDelete202NoRetry204(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	diffPoller, err := op.ResumePost200WithPayload(rt)
	if err == nil {
		t.Fatal("expected an error but did not find receive one")
	}
	if diffPoller != nil {
		t.Fatal("expected a nil poller from the resume operation")
	}
}

func TestLROBeginDelete202NoRetry204(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDelete202NoRetry204(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDelete202NoRetry204(rt)
	if err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusNoContent {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestLROBeginDelete202Retry200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDelete202Retry200(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDelete202Retry200(rt)
	if err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestLROBeginDelete204Succeeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDelete204Succeeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	_, err = poller.ResumeToken()
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := res.StatusCode; s != http.StatusNoContent {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestLROBeginDeleteAsyncNoHeaderInRetry(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteAsyncNoHeaderInRetry(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteAsyncNoHeaderInRetry(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := res.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestLROBeginDeleteAsyncNoRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteAsyncNoRetrySucceeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteAsyncNoRetrySucceeded(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := res.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestLROBeginDeleteAsyncRetryFailed(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteAsyncRetryFailed(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteAsyncRetryFailed(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response from the polling operation")
	}
	var cloudErr *CloudError
	if !errors.As(err, &cloudErr) {
		t.Fatal("expected a CloudError but did not receive one")
	}
	var httpResp azcore.HTTPResponse
	if !errors.As(err, &httpResp) {
		t.Fatal("expected azcore.HTTPResponse error")
	} else if sc := httpResp.RawResponse().StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginDeleteAsyncRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteAsyncRetrySucceeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteAsyncRetrySucceeded(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := res.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestLROBeginDeleteAsyncRetrycanceled(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteAsyncRetrycanceled(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteAsyncRetrycanceled(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response from the polling operation")
	}
	var cloudErr *CloudError
	if !errors.As(err, &cloudErr) {
		t.Fatal("expected a CloudError but did not receive one")
	}
	var httpResp azcore.HTTPResponse
	if !errors.As(err, &httpResp) {
		t.Fatal("expected azcore.HTTPResponse error")
	} else if sc := httpResp.RawResponse().StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginDeleteNoHeaderInRetry(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteNoHeaderInRetry(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteNoHeaderInRetry(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := res.StatusCode; s != http.StatusNoContent {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestLROBeginDeleteProvisioning202Accepted200Succeeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteProvisioning202Accepted200Succeeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteProvisioning202Accepted200Succeeded(rt)
	if err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestLROBeginDeleteProvisioning202DeletingFailed200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteProvisioning202DeletingFailed200(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteProvisioning202DeletingFailed200(rt)
	if err != nil {
		t.Fatal(err)
	}
	_, err = resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROBeginDeleteProvisioning202Deletingcanceled200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteProvisioning202Deletingcanceled200(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteProvisioning202Deletingcanceled200(rt)
	if err != nil {
		t.Fatal(err)
	}
	_, err = resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROBeginPost200WithPayload(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPost200WithPayload(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePost200WithPayload(rt)
	if err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.SKU, &SKU{
		ID:   to.StringPtr("1"),
		Name: to.StringPtr("product"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPost202List(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPost202List(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePost202List(rt)
	if err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.ProductArray, []Product{
		{
			Resource: Resource{
				ID:   to.StringPtr("100"),
				Name: to.StringPtr("foo"),
			},
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPost202NoRetry204(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPost202NoRetry204(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePost202NoRetry204(rt)
	if err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusNoContent {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestLROBeginPost202Retry200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPost202Retry200(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePost202Retry200(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := res.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestLROBeginPostAsyncNoRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostAsyncNoRetrySucceeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePostAsyncNoRetrySucceeded(rt)
	if err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, &Product{
		Resource: Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPostAsyncRetryFailed(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostAsyncRetryFailed(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePostAsyncRetryFailed(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response from the polling operation")
	}
	var cloudErr *CloudError
	if !errors.As(err, &cloudErr) {
		t.Fatal("expected a CloudError but did not receive one")
	}
	var httpResp azcore.HTTPResponse
	if !errors.As(err, &httpResp) {
		t.Fatal("expected azcore.HTTPResponse error")
	} else if sc := httpResp.RawResponse().StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPostAsyncRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostAsyncRetrySucceeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePostAsyncRetrySucceeded(rt)
	if err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, &Product{
		Resource: Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPostAsyncRetrycanceled(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostAsyncRetrycanceled(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePostAsyncRetrycanceled(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response from the polling operation")
	}
	var cloudErr *CloudError
	if !errors.As(err, &cloudErr) {
		t.Fatal("expected a CloudError but did not receive one")
	}
	var httpResp azcore.HTTPResponse
	if !errors.As(err, &httpResp) {
		t.Fatal("expected azcore.HTTPResponse error")
	} else if sc := httpResp.RawResponse().StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPostDoubleHeadersFinalAzureHeaderGet(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostDoubleHeadersFinalAzureHeaderGet(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePostDoubleHeadersFinalAzureHeaderGet(rt)
	if err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, &Product{
		Resource: Resource{
			ID: to.StringPtr("100"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPostDoubleHeadersFinalAzureHeaderGetDefault(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostDoubleHeadersFinalAzureHeaderGetDefault(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePostDoubleHeadersFinalAzureHeaderGetDefault(rt)
	if err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, &Product{
		Resource: Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPostDoubleHeadersFinalLocationGet(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostDoubleHeadersFinalLocationGet(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePostDoubleHeadersFinalLocationGet(rt)
	if err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, &Product{}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPut200Acceptedcanceled200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut200Acceptedcanceled200(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePut200Acceptedcanceled200(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if r := cmp.Diff(res, ProductResponse{}); r != "" {
		t.Fatal(r)
	}
	var cloudErr *CloudError
	if !errors.As(err, &cloudErr) {
		t.Fatal("expected a CloudError but did not receive one")
	}
	var httpResp azcore.HTTPResponse
	if !errors.As(err, &httpResp) {
		t.Fatal("expected azcore.HTTPResponse error")
	} else if sc := httpResp.RawResponse().StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPut200Succeeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut200Succeeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	_, err = poller.ResumeToken()
	if err == nil {
		t.Fatal("Expected an error but did not receive one")
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, &Product{
		Resource: Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPut200SucceededNoState(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut200SucceededNoState(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	_, err = poller.ResumeToken()
	if err == nil {
		t.Fatal("Expected an error but did not receive one")
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, &Product{
		Resource: Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

// TODO check if this test should actually be returning a 200 or a 204
func TestLROBeginPut200UpdatingSucceeded204(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut200UpdatingSucceeded204(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePut200UpdatingSucceeded204(rt)
	if err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, &Product{
		Resource: Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPut201CreatingFailed200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut201CreatingFailed200(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePut201CreatingFailed200(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if r := cmp.Diff(res, ProductResponse{}); r != "" {
		t.Fatal(r)
	}
	var cloudErr *CloudError
	if !errors.As(err, &cloudErr) {
		t.Fatal("expected a CloudError but did not receive one")
	}
	var httpResp azcore.HTTPResponse
	if !errors.As(err, &httpResp) {
		t.Fatal("expected azcore.HTTPResponse error")
	} else if sc := httpResp.RawResponse().StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPut201CreatingSucceeded200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut201CreatingSucceeded200(context.Background(), &LrOSBeginPut201CreatingSucceeded200Options{Product: &Product{}})
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePut201CreatingSucceeded200(rt)
	if err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, &Product{
		Resource: Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPut202Retry200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut202Retry200(context.Background(), &LrOSBeginPut202Retry200Options{Product: &Product{}})
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePut202Retry200(rt)
	if err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, &Product{
		Resource: Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPutAsyncNoHeaderInRetry(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncNoHeaderInRetry(context.Background(), &LrOSBeginPutAsyncNoHeaderInRetryOptions{Product: &Product{}})
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePutAsyncNoHeaderInRetry(rt)
	if err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, &Product{
		Resource: Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPutAsyncNoRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncNoRetrySucceeded(context.Background(), &LrOSBeginPutAsyncNoRetrySucceededOptions{Product: &Product{}})
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePutAsyncNoRetrySucceeded(rt)
	if err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, &Product{
		Resource: Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPutAsyncNoRetrycanceled(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncNoRetrycanceled(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePutAsyncNoRetrycanceled(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if r := cmp.Diff(res, ProductResponse{}); r != "" {
		t.Fatal(r)
	}
	var cloudErr *CloudError
	if !errors.As(err, &cloudErr) {
		t.Fatal("expected a CloudError but did not receive one")
	}
	var httpResp azcore.HTTPResponse
	if !errors.As(err, &httpResp) {
		t.Fatal("expected azcore.HTTPResponse error")
	} else if sc := httpResp.RawResponse().StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPutAsyncNonResource(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncNonResource(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePutAsyncNonResource(rt)
	if err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.SKU, &SKU{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("sku"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPutAsyncRetryFailed(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncRetryFailed(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePutAsyncRetryFailed(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if r := cmp.Diff(res, ProductResponse{}); r != "" {
		t.Fatal(r)
	}
	var cloudErr *CloudError
	if !errors.As(err, &cloudErr) {
		t.Fatal("expected a CloudError but did not receive one")
	}
	var httpResp azcore.HTTPResponse
	if !errors.As(err, &httpResp) {
		t.Fatal("expected azcore.HTTPResponse error")
	} else if sc := httpResp.RawResponse().StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPutAsyncRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncRetrySucceeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePutAsyncRetrySucceeded(rt)
	if err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, &Product{
		Resource: Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPutAsyncSubResource(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncSubResource(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePutAsyncSubResource(rt)
	if err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.SubProduct, &SubProduct{
		SubResource: SubResource{
			ID: to.StringPtr("100"),
		},
		Properties: &SubProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPutNoHeaderInRetry(t *testing.T) {
	t.Skip("problem with put flow")
	op := newLROSClient()
	resp, err := op.BeginPutNoHeaderInRetry(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePutNoHeaderInRetry(rt)
	if err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.Product, &Product{
		Resource: Resource{
			ID: to.StringPtr("100"),
		},
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPutNonResource(t *testing.T) {
	t.Skip("problem with put flow")
	op := newLROSClient()
	resp, err := op.BeginPutNonResource(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePutNonResource(rt)
	if err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.SKU, &SKU{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("sku"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestLROBeginPutSubResource(t *testing.T) {
	t.Skip("problem with put flow")
	op := newLROSClient()
	resp, err := op.BeginPutSubResource(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePutSubResource(rt)
	if err != nil {
		t.Fatal(err)
	}
	pollResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if s := pollResp.RawResponse.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
	if r := cmp.Diff(pollResp.SubProduct, &SubProduct{
		SubResource: SubResource{
			ID: to.StringPtr("100"),
		},
		Properties: &SubProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	}); r != "" {
		t.Fatal(r)
	}
}
