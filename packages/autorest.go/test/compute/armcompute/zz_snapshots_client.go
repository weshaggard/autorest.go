//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SnapshotsClient contains the methods for the Snapshots group.
// Don't use this type directly, use NewSnapshotsClient() instead.
type SnapshotsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSnapshotsClient creates a new instance of SnapshotsClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSnapshotsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SnapshotsClient, error) {
	cl, err := arm.NewClient(moduleName+".SnapshotsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SnapshotsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01
//   - resourceGroupName - The name of the resource group.
//   - snapshotName - The name of the snapshot that is being created. The name can't be changed after the snapshot is created.
//     Supported characters for the name are a-z, A-Z, 0-9, _ and -. The max name length is 80
//     characters.
//   - snapshot - Snapshot object supplied in the body of the Put disk operation.
//   - options - SnapshotsClientBeginCreateOrUpdateOptions contains the optional parameters for the SnapshotsClient.BeginCreateOrUpdate
//     method.
func (client *SnapshotsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, snapshotName string, snapshot Snapshot, options *SnapshotsClientBeginCreateOrUpdateOptions) (*runtime.Poller[SnapshotsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, snapshotName, snapshot, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[SnapshotsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SnapshotsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01
func (client *SnapshotsClient) createOrUpdate(ctx context.Context, resourceGroupName string, snapshotName string, snapshot Snapshot, options *SnapshotsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "SnapshotsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, snapshotName, snapshot, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SnapshotsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, snapshotName string, snapshot Snapshot, options *SnapshotsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if snapshotName == "" {
		return nil, errors.New("parameter snapshotName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotName}", url.PathEscape(snapshotName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, snapshot); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01
//   - resourceGroupName - The name of the resource group.
//   - snapshotName - The name of the snapshot that is being created. The name can't be changed after the snapshot is created.
//     Supported characters for the name are a-z, A-Z, 0-9, _ and -. The max name length is 80
//     characters.
//   - options - SnapshotsClientBeginDeleteOptions contains the optional parameters for the SnapshotsClient.BeginDelete method.
func (client *SnapshotsClient) BeginDelete(ctx context.Context, resourceGroupName string, snapshotName string, options *SnapshotsClientBeginDeleteOptions) (*runtime.Poller[SnapshotsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, snapshotName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[SnapshotsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SnapshotsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01
func (client *SnapshotsClient) deleteOperation(ctx context.Context, resourceGroupName string, snapshotName string, options *SnapshotsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "SnapshotsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, snapshotName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SnapshotsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, snapshotName string, options *SnapshotsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if snapshotName == "" {
		return nil, errors.New("parameter snapshotName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotName}", url.PathEscape(snapshotName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets information about a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01
//   - resourceGroupName - The name of the resource group.
//   - snapshotName - The name of the snapshot that is being created. The name can't be changed after the snapshot is created.
//     Supported characters for the name are a-z, A-Z, 0-9, _ and -. The max name length is 80
//     characters.
//   - options - SnapshotsClientGetOptions contains the optional parameters for the SnapshotsClient.Get method.
func (client *SnapshotsClient) Get(ctx context.Context, resourceGroupName string, snapshotName string, options *SnapshotsClientGetOptions) (SnapshotsClientGetResponse, error) {
	var err error
	const operationName = "SnapshotsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, snapshotName, options)
	if err != nil {
		return SnapshotsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SnapshotsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SnapshotsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SnapshotsClient) getCreateRequest(ctx context.Context, resourceGroupName string, snapshotName string, options *SnapshotsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if snapshotName == "" {
		return nil, errors.New("parameter snapshotName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotName}", url.PathEscape(snapshotName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SnapshotsClient) getHandleResponse(resp *http.Response) (SnapshotsClientGetResponse, error) {
	result := SnapshotsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Snapshot); err != nil {
		return SnapshotsClientGetResponse{}, err
	}
	return result, nil
}

// BeginGrantAccess - Grants access to a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01
//   - resourceGroupName - The name of the resource group.
//   - snapshotName - The name of the snapshot that is being created. The name can't be changed after the snapshot is created.
//     Supported characters for the name are a-z, A-Z, 0-9, _ and -. The max name length is 80
//     characters.
//   - grantAccessData - Access data object supplied in the body of the get snapshot access operation.
//   - options - SnapshotsClientBeginGrantAccessOptions contains the optional parameters for the SnapshotsClient.BeginGrantAccess
//     method.
func (client *SnapshotsClient) BeginGrantAccess(ctx context.Context, resourceGroupName string, snapshotName string, grantAccessData GrantAccessData, options *SnapshotsClientBeginGrantAccessOptions) (*runtime.Poller[SnapshotsClientGrantAccessResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.grantAccess(ctx, resourceGroupName, snapshotName, grantAccessData, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SnapshotsClientGrantAccessResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SnapshotsClientGrantAccessResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// GrantAccess - Grants access to a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01
func (client *SnapshotsClient) grantAccess(ctx context.Context, resourceGroupName string, snapshotName string, grantAccessData GrantAccessData, options *SnapshotsClientBeginGrantAccessOptions) (*http.Response, error) {
	var err error
	const operationName = "SnapshotsClient.BeginGrantAccess"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.grantAccessCreateRequest(ctx, resourceGroupName, snapshotName, grantAccessData, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// grantAccessCreateRequest creates the GrantAccess request.
func (client *SnapshotsClient) grantAccessCreateRequest(ctx context.Context, resourceGroupName string, snapshotName string, grantAccessData GrantAccessData, options *SnapshotsClientBeginGrantAccessOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}/beginGetAccess"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if snapshotName == "" {
		return nil, errors.New("parameter snapshotName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotName}", url.PathEscape(snapshotName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, grantAccessData); err != nil {
		return nil, err
	}
	return req, nil
}

// NewListPager - Lists snapshots under a subscription.
//
// Generated from API version 2021-12-01
//   - options - SnapshotsClientListOptions contains the optional parameters for the SnapshotsClient.NewListPager method.
func (client *SnapshotsClient) NewListPager(options *SnapshotsClientListOptions) *runtime.Pager[SnapshotsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SnapshotsClientListResponse]{
		More: func(page SnapshotsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SnapshotsClientListResponse) (SnapshotsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SnapshotsClient.NewListPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SnapshotsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SnapshotsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SnapshotsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *SnapshotsClient) listCreateRequest(ctx context.Context, options *SnapshotsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/snapshots"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SnapshotsClient) listHandleResponse(resp *http.Response) (SnapshotsClientListResponse, error) {
	result := SnapshotsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SnapshotList); err != nil {
		return SnapshotsClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists snapshots under a resource group.
//
// Generated from API version 2021-12-01
//   - resourceGroupName - The name of the resource group.
//   - options - SnapshotsClientListByResourceGroupOptions contains the optional parameters for the SnapshotsClient.NewListByResourceGroupPager
//     method.
func (client *SnapshotsClient) NewListByResourceGroupPager(resourceGroupName string, options *SnapshotsClientListByResourceGroupOptions) *runtime.Pager[SnapshotsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[SnapshotsClientListByResourceGroupResponse]{
		More: func(page SnapshotsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SnapshotsClientListByResourceGroupResponse) (SnapshotsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SnapshotsClient.NewListByResourceGroupPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SnapshotsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SnapshotsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SnapshotsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SnapshotsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SnapshotsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SnapshotsClient) listByResourceGroupHandleResponse(resp *http.Response) (SnapshotsClientListByResourceGroupResponse, error) {
	result := SnapshotsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SnapshotList); err != nil {
		return SnapshotsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginRevokeAccess - Revokes access to a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01
//   - resourceGroupName - The name of the resource group.
//   - snapshotName - The name of the snapshot that is being created. The name can't be changed after the snapshot is created.
//     Supported characters for the name are a-z, A-Z, 0-9, _ and -. The max name length is 80
//     characters.
//   - options - SnapshotsClientBeginRevokeAccessOptions contains the optional parameters for the SnapshotsClient.BeginRevokeAccess
//     method.
func (client *SnapshotsClient) BeginRevokeAccess(ctx context.Context, resourceGroupName string, snapshotName string, options *SnapshotsClientBeginRevokeAccessOptions) (*runtime.Poller[SnapshotsClientRevokeAccessResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.revokeAccess(ctx, resourceGroupName, snapshotName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SnapshotsClientRevokeAccessResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SnapshotsClientRevokeAccessResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// RevokeAccess - Revokes access to a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01
func (client *SnapshotsClient) revokeAccess(ctx context.Context, resourceGroupName string, snapshotName string, options *SnapshotsClientBeginRevokeAccessOptions) (*http.Response, error) {
	var err error
	const operationName = "SnapshotsClient.BeginRevokeAccess"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.revokeAccessCreateRequest(ctx, resourceGroupName, snapshotName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// revokeAccessCreateRequest creates the RevokeAccess request.
func (client *SnapshotsClient) revokeAccessCreateRequest(ctx context.Context, resourceGroupName string, snapshotName string, options *SnapshotsClientBeginRevokeAccessOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}/endGetAccess"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if snapshotName == "" {
		return nil, errors.New("parameter snapshotName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotName}", url.PathEscape(snapshotName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// BeginUpdate - Updates (patches) a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01
//   - resourceGroupName - The name of the resource group.
//   - snapshotName - The name of the snapshot that is being created. The name can't be changed after the snapshot is created.
//     Supported characters for the name are a-z, A-Z, 0-9, _ and -. The max name length is 80
//     characters.
//   - snapshot - Snapshot object supplied in the body of the Patch snapshot operation.
//   - options - SnapshotsClientBeginUpdateOptions contains the optional parameters for the SnapshotsClient.BeginUpdate method.
func (client *SnapshotsClient) BeginUpdate(ctx context.Context, resourceGroupName string, snapshotName string, snapshot SnapshotUpdate, options *SnapshotsClientBeginUpdateOptions) (*runtime.Poller[SnapshotsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, snapshotName, snapshot, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[SnapshotsClientUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[SnapshotsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Updates (patches) a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01
func (client *SnapshotsClient) update(ctx context.Context, resourceGroupName string, snapshotName string, snapshot SnapshotUpdate, options *SnapshotsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "SnapshotsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, snapshotName, snapshot, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *SnapshotsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, snapshotName string, snapshot SnapshotUpdate, options *SnapshotsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if snapshotName == "" {
		return nil, errors.New("parameter snapshotName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotName}", url.PathEscape(snapshotName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, snapshot); err != nil {
		return nil, err
	}
	return req, nil
}
