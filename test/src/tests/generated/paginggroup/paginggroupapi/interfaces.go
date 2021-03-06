package paginggroupapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"tests/generated/paginggroup"
)

// PagingClientAPI contains the set of methods on the PagingClient type.
type PagingClientAPI interface {
	GetMultiplePages(ctx context.Context, clientRequestID string, maxresults *int32, timeout *int32) (result paginggroup.ProductResultPage, err error)
	GetMultiplePagesFailure(ctx context.Context) (result paginggroup.ProductResultPage, err error)
	GetMultiplePagesFailureURI(ctx context.Context) (result paginggroup.ProductResultPage, err error)
	GetMultiplePagesFragmentNextLink(ctx context.Context, APIVersion string, tenant string) (result paginggroup.OdataProductResultPage, err error)
	GetMultiplePagesFragmentWithGroupingNextLink(ctx context.Context, APIVersion string, tenant string) (result paginggroup.OdataProductResultPage, err error)
	GetMultiplePagesLRO(ctx context.Context, clientRequestID string, maxresults *int32, timeout *int32) (result paginggroup.PagingGetMultiplePagesLROFuture, err error)
	GetMultiplePagesRetryFirst(ctx context.Context) (result paginggroup.ProductResultPage, err error)
	GetMultiplePagesRetrySecond(ctx context.Context) (result paginggroup.ProductResultPage, err error)
	GetMultiplePagesWithOffset(ctx context.Context, offset int32, clientRequestID string, maxresults *int32, timeout *int32) (result paginggroup.ProductResultPage, err error)
	GetOdataMultiplePages(ctx context.Context, clientRequestID string, maxresults *int32, timeout *int32) (result paginggroup.OdataProductResultPage, err error)
	GetSinglePages(ctx context.Context) (result paginggroup.ProductResultPage, err error)
	GetSinglePagesFailure(ctx context.Context) (result paginggroup.ProductResultPage, err error)
	NextFragment(ctx context.Context, APIVersion string, tenant string, nextLink string) (result paginggroup.OdataProductResult, err error)
	NextFragmentWithGrouping(ctx context.Context, APIVersion string, tenant string, nextLink string) (result paginggroup.OdataProductResult, err error)
}

var _ PagingClientAPI = (*paginggroup.PagingClient)(nil)
