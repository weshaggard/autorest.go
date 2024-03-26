// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azalias

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Client contains the methods for the Alias group.
// Don't use this type directly, use a constructor function instead.
type Client struct {
	internal            *azcore.Client
	endpoint            string
	clientGroup         ClientGroup
	clientOptionalGroup *ClientOptionalGroup
	optionalString      *string
}

// Create - Applies to: see pricing tiers [https://aka.ms/AzureMapsPricingTier].
// Creator makes it possible to develop applications based on your private indoor map data using Azure Maps API and SDK. This
// [https://docs.microsoft.com/azure/azure-maps/creator-indoor-maps] article
// introduces concepts and tools that apply to Azure Maps Creator.
// This API allows the caller to create an alias. You can also assign the alias during the create request. An alias can reference
// an ID generated by a creator service, but cannot reference another alias
// ID.
// SUBMIT CREATE REQUEST To create your alias, you will use a POST request. If you would like to assign the alias during the
// creation, you will pass the resourceId query parameter.
// CREATE ALIAS RESPONSE The Create API returns a HTTP 201 Created response with the alias resource in the body.
// A sample response from creating an alias:
// { "createdTimestamp": "2020-02-13T21:19:11.123Z", "aliasId": "a8a4b8bb-ecf4-fb27-a618-f41721552766", "creatorDataItemId":
// "e89aebb9-70a3-8fe1-32bb-1fbd0c725f14", "lastUpdatedTimestamp":
// "2020-02-13T21:19:22.123Z" }
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0
//   - stringQuery - The unique id that references the assigned data item to be aliased.
//   - boolHeaderEnum - Some enums that are boolean values.
//   - unixTimeQuery - Required unix time passed via query param.
//   - options - CreateOptions contains the optional parameters for the Client.Create method.
func (client *Client) Create(ctx context.Context, headerBools []bool, stringQuery string, boolHeaderEnum BooleanEnum, unixTimeQuery time.Time, headerEnum SomeEnum, queryEnum SomeEnum, options *CreateOptions) (CreateResponse, error) {
	var err error
	const operationName = "Client.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, headerBools, stringQuery, boolHeaderEnum, unixTimeQuery, headerEnum, queryEnum, options)
	if err != nil {
		return CreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return CreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *Client) createCreateRequest(ctx context.Context, headerBools []bool, stringQuery string, boolHeaderEnum BooleanEnum, unixTimeQuery time.Time, headerEnum SomeEnum, queryEnum SomeEnum, options *CreateOptions) (*policy.Request, error) {
	urlPath := "/aliases"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.BoolHeaderEnum1 != nil {
		reqQP.Set("boolHeaderEnum", fmt.Sprintf("%v", *options.BoolHeaderEnum1))
	}
	reqQP.Set("client-version", client.clientGroup.ClientVersion)
	creatorIDDefault := int32(123)
	if options != nil && options.CreatorID != nil {
		creatorIDDefault = *options.CreatorID
	}
	reqQP.Set("creator-id", strconv.FormatInt(int64(creatorIDDefault), 10))
	if options != nil && options.GroupBy != nil {
		for _, qv := range options.GroupBy {
			reqQP.Add("groupBy", fmt.Sprintf("%d", qv))
		}
	}
	reqQP.Set("queryEnum", string(queryEnum))
	reqQP.Set("stringQuery", stringQuery)
	reqQP.Set("unixTimeQuery", timeUnix(unixTimeQuery).String())
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	assignedIDDefault := float32(8989)
	if options != nil && options.AssignedID != nil {
		assignedIDDefault = *options.AssignedID
	}
	req.Raw().Header["assigned-id"] = []string{strconv.FormatFloat(float64(assignedIDDefault), 'f', -1, 32)}
	req.Raw().Header["boolHeaderEnum"] = []string{fmt.Sprintf("%v", boolHeaderEnum)}
	req.Raw().Header["client-index"] = []string{strconv.FormatInt(int64(client.clientGroup.ClientIndex), 10)}
	req.Raw().Header["headerBools"] = []string{strings.Join(strings.Fields(strings.Trim(fmt.Sprint(headerBools), "[]")), ",")}
	req.Raw().Header["headerEnum"] = []string{string(headerEnum)}
	if options != nil && options.OptionalUnixTime != nil {
		req.Raw().Header["optionalUnixTime"] = []string{timeUnix(*options.OptionalUnixTime).String()}
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *Client) createHandleResponse(resp *http.Response) (CreateResponse, error) {
	result := CreateResponse{}
	if val := resp.Header.Get("Access-Control-Expose-Headers"); val != "" {
		result.AccessControlExposeHeaders = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AliasesCreateResponse); err != nil {
		return CreateResponse{}, err
	}
	return result, nil
}

// GetScript - Retrieve the configuration script identified by configuration name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0
//   - SomeGroup - SomeGroup contains a group of parameters for the Client.GetScript method.
//   - ExplodedGroup - ExplodedGroup contains a group of parameters for the Client.GetScript method.
//   - options - GetScriptOptions contains the optional parameters for the Client.GetScript method.
func (client *Client) GetScript(ctx context.Context, headerCounts []int32, queryCounts []int64, explodedStringStuff []string, numericHeader int32, headerTime time.Time, props GeoJSONObjectNamedCollection, someGroup SomeGroup, explodedGroup ExplodedGroup, options *GetScriptOptions) (GetScriptResponse, error) {
	var err error
	const operationName = "Client.GetScript"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getScriptCreateRequest(ctx, headerCounts, queryCounts, explodedStringStuff, numericHeader, headerTime, props, someGroup, explodedGroup, options)
	if err != nil {
		return GetScriptResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GetScriptResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GetScriptResponse{}, err
	}
	resp, err := client.getScriptHandleResponse(httpResp)
	return resp, err
}

// getScriptCreateRequest creates the GetScript request.
func (client *Client) getScriptCreateRequest(ctx context.Context, headerCounts []int32, queryCounts []int64, explodedStringStuff []string, numericHeader int32, headerTime time.Time, props GeoJSONObjectNamedCollection, someGroup SomeGroup, explodedGroup ExplodedGroup, options *GetScriptOptions) (*policy.Request, error) {
	urlPath := "/scripts"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	for _, qv := range explodedStringStuff {
		reqQP.Add("explodedStringStuff", qv)
	}
	for _, qv := range explodedGroup.ExplodedStuff {
		reqQP.Add("explodedStuff", fmt.Sprintf("%v", qv))
	}
	if options != nil && options.OptionalExplodedStuff != nil {
		for _, qv := range options.OptionalExplodedStuff {
			reqQP.Add("optionalExplodedStuff", qv)
		}
	}
	reqQP.Set("queryCounts", strings.Join(strings.Fields(strings.Trim(fmt.Sprint(queryCounts), "[]")), ","))
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"text/powershell"}
	req.Raw().Header["headerCounts"] = []string{strings.Join(strings.Fields(strings.Trim(fmt.Sprint(headerCounts), "[]")), ",")}
	req.Raw().Header["headerStrings"] = []string{strings.Join(someGroup.HeaderStrings, ",")}
	req.Raw().Header["headerTime"] = []string{timeRFC3339(headerTime).String()}
	req.Raw().Header["numericHeader"] = []string{strconv.FormatInt(int64(numericHeader), 10)}
	if err := runtime.MarshalAsJSON(req, props); err != nil {
		return nil, err
	}
	return req, nil
}

// getScriptHandleResponse handles the GetScript response.
func (client *Client) getScriptHandleResponse(resp *http.Response) (GetScriptResponse, error) {
	result := GetScriptResponse{}
	body, err := runtime.Payload(resp)
	if err != nil {
		return GetScriptResponse{}, err
	}
	txt := string(body)
	result.Value = &txt
	return result, nil
}

// NewListPager - Applies to: see pricing tiers [https://aka.ms/AzureMapsPricingTier].
// Creator makes it possible to develop applications based on your private indoor map data using Azure Maps API and SDK. This
// [https://docs.microsoft.com/azure/azure-maps/creator-indoor-maps] article
// introduces concepts and tools that apply to Azure Maps Creator.
// This API allows the caller to fetch a list of all previously successfully created aliases.
// SUBMIT LIST REQUEST To list all your aliases, you will issue a GET request with no additional parameters.
// LIST DATA RESPONSE The List API returns the complete list of all aliases in json format. The response contains the following
// details for each alias resource:
// > createdTimestamp - The timestamp that the alias was created. Format yyyy-MM-ddTHH:mm:ss.sssZ aliasId - The id for the
// alias. creatorDataItemId - The id for the creator data item that this alias
// references (could be null if the alias has not been assigned). lastUpdatedTimestamp - The last time the alias was assigned
// to a resource. Format yyyy-MM-ddTHH:mm:ss.sssZ
// A sample response returning 2 alias resources:
// { "aliases": [ { "createdTimestamp": "2020-02-13T21:19:11.123Z", "aliasId": "a8a4b8bb-ecf4-fb27-a618-f41721552766", "creatorDataItemId":
// "e89aebb9-70a3-8fe1-32bb-1fbd0c725f14", "lastUpdatedTimestamp":
// "2020-02-13T21:19:22.123Z" }, { "createdTimestamp": "2020-02-18T19:53:33.123Z", "aliasId": "1856dbfc-7a66-ee5a-bf8d-51dbfe1906f6",
// "creatorDataItemId": null, "lastUpdatedTimestamp":
// "2020-02-18T19:53:33.123Z" } ] }
//
// Generated from API version 2.0
//   - options - ListOptions contains the optional parameters for the Client.NewListPager method.
func (client *Client) NewListPager(headerEnums []IntEnum, queryEnum IntEnum, options *ListOptions) *runtime.Pager[ListResponseEnvelope] {
	return runtime.NewPager(runtime.PagingHandler[ListResponseEnvelope]{
		More: func(page ListResponseEnvelope) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ListResponseEnvelope) (ListResponseEnvelope, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "Client.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, headerEnums, queryEnum, options)
			}, nil)
			if err != nil {
				return ListResponseEnvelope{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *Client) listCreateRequest(ctx context.Context, headerEnums []IntEnum, queryEnum IntEnum, options *ListOptions) (*policy.Request, error) {
	urlPath := "/aliases"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("client-version", client.clientGroup.ClientVersion)
	if options != nil && options.GroupBy != nil {
		for _, qv := range options.GroupBy {
			reqQP.Add("groupBy", string(qv))
		}
	}
	reqQP.Set("queryEnum", fmt.Sprintf("%v", queryEnum))
	if options != nil && options.QueryEnums != nil {
		for _, qv := range options.QueryEnums {
			reqQP.Add("queryEnums", fmt.Sprintf("%d", qv))
		}
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.HeaderEnum != nil {
		req.Raw().Header["headerEnum"] = []string{fmt.Sprintf("%v", *options.HeaderEnum)}
	}
	req.Raw().Header["headerEnums"] = []string{strings.Join(strings.Fields(strings.Trim(fmt.Sprint(headerEnums), "[]")), ",")}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *Client) listHandleResponse(resp *http.Response) (ListResponseEnvelope, error) {
	result := ListResponseEnvelope{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResponse); err != nil {
		return ListResponseEnvelope{}, err
	}
	return result, nil
}

// BeginListLRO - A long-running paged operation that uses a next link operation
//
// Generated from API version 2.0
//   - options - BeginListLROOptions contains the optional parameters for the Client.BeginListLRO method.
func (client *Client) BeginListLRO(ctx context.Context, options *BeginListLROOptions) (*runtime.Poller[*runtime.Pager[ListLROResponse]], error) {
	pager := runtime.NewPager(runtime.PagingHandler[ListLROResponse]{
		More: func(page ListLROResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ListLROResponse) (ListLROResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "Client.BeginListLRO")
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), *page.NextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listLROCreateRequest(ctx, options)
			}, &runtime.FetcherForNextLinkOptions{
				NextReq: func(ctx context.Context, encodedNextLink string) (*policy.Request, error) {
					return client.listLRONextCreateRequest(ctx, encodedNextLink)
				},
			})
			if err != nil {
				return ListLROResponse{}, err
			}
			return client.listLROHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
	if options == nil || options.ResumeToken == "" {
		resp, err := client.listLRO(ctx, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[*runtime.Pager[ListLROResponse]]{
			Response: &pager,
			Tracer:   client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[*runtime.Pager[ListLROResponse]]{
			Response: &pager,
			Tracer:   client.internal.Tracer(),
		})
	}
}

// ListLRO - A long-running paged operation that uses a next link operation
//
// Generated from API version 2.0
func (client *Client) listLRO(ctx context.Context, options *BeginListLROOptions) (*http.Response, error) {
	var err error
	const operationName = "Client.BeginListLRO"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listLROCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// listLROCreateRequest creates the ListLRO request.
func (client *Client) listLROCreateRequest(ctx context.Context, options *BeginListLROOptions) (*policy.Request, error) {
	urlPath := "/paged"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listLROHandleResponse handles the ListLRO response.
func (client *Client) listLROHandleResponse(resp *http.Response) (ListLROResponse, error) {
	result := ListLROResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PagesOfThings); err != nil {
		return ListLROResponse{}, err
	}
	return result, nil
}

// - options - ListWithSharedNextOneOptions contains the optional parameters for the Client.NewListWithSharedNextOnePager method.
func (client *Client) NewListWithSharedNextOnePager(options *ListWithSharedNextOneOptions) *runtime.Pager[ListWithSharedNextOneResponse] {
	return runtime.NewPager(runtime.PagingHandler[ListWithSharedNextOneResponse]{
		More: func(page ListWithSharedNextOneResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ListWithSharedNextOneResponse) (ListWithSharedNextOneResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "Client.NewListWithSharedNextOnePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listWithSharedNextOneCreateRequest(ctx, options)
			}, &runtime.FetcherForNextLinkOptions{
				NextReq: func(ctx context.Context, encodedNextLink string) (*policy.Request, error) {
					return client.listWithSharedNextCreateRequest(ctx, encodedNextLink)
				},
			})
			if err != nil {
				return ListWithSharedNextOneResponse{}, err
			}
			return client.listWithSharedNextOneHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listWithSharedNextOneCreateRequest creates the ListWithSharedNextOne request.
func (client *Client) listWithSharedNextOneCreateRequest(ctx context.Context, options *ListWithSharedNextOneOptions) (*policy.Request, error) {
	urlPath := "/listWithSharedNextOne"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listWithSharedNextOneHandleResponse handles the ListWithSharedNextOne response.
func (client *Client) listWithSharedNextOneHandleResponse(resp *http.Response) (ListWithSharedNextOneResponse, error) {
	result := ListWithSharedNextOneResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResponse); err != nil {
		return ListWithSharedNextOneResponse{}, err
	}
	return result, nil
}

// - options - ListWithSharedNextTwoOptions contains the optional parameters for the Client.NewListWithSharedNextTwoPager method.
func (client *Client) NewListWithSharedNextTwoPager(options *ListWithSharedNextTwoOptions) *runtime.Pager[ListWithSharedNextTwoResponse] {
	return runtime.NewPager(runtime.PagingHandler[ListWithSharedNextTwoResponse]{
		More: func(page ListWithSharedNextTwoResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ListWithSharedNextTwoResponse) (ListWithSharedNextTwoResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "Client.NewListWithSharedNextTwoPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listWithSharedNextTwoCreateRequest(ctx, options)
			}, &runtime.FetcherForNextLinkOptions{
				NextReq: func(ctx context.Context, encodedNextLink string) (*policy.Request, error) {
					return client.listWithSharedNextCreateRequest(ctx, encodedNextLink)
				},
			})
			if err != nil {
				return ListWithSharedNextTwoResponse{}, err
			}
			return client.listWithSharedNextTwoHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listWithSharedNextTwoCreateRequest creates the ListWithSharedNextTwo request.
func (client *Client) listWithSharedNextTwoCreateRequest(ctx context.Context, options *ListWithSharedNextTwoOptions) (*policy.Request, error) {
	urlPath := "/listWithSharedNextTwo"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listWithSharedNextTwoHandleResponse handles the ListWithSharedNextTwo response.
func (client *Client) listWithSharedNextTwoHandleResponse(resp *http.Response) (ListWithSharedNextTwoResponse, error) {
	result := ListWithSharedNextTwoResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResponse); err != nil {
		return ListWithSharedNextTwoResponse{}, err
	}
	return result, nil
}

// PolicyAssignment -
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0
//   - options - PolicyAssignmentOptions contains the optional parameters for the Client.PolicyAssignment method.
func (client *Client) PolicyAssignment(ctx context.Context, things []Things, polymorphicParam GeoJSONObjectClassification, options *PolicyAssignmentOptions) (PolicyAssignmentResponse, error) {
	var err error
	const operationName = "Client.PolicyAssignment"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.policyAssignmentCreateRequest(ctx, things, polymorphicParam, options)
	if err != nil {
		return PolicyAssignmentResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PolicyAssignmentResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PolicyAssignmentResponse{}, err
	}
	resp, err := client.policyAssignmentHandleResponse(httpResp)
	return resp, err
}

// policyAssignmentCreateRequest creates the PolicyAssignment request.
func (client *Client) policyAssignmentCreateRequest(ctx context.Context, things []Things, polymorphicParam GeoJSONObjectClassification, options *PolicyAssignmentOptions) (*policy.Request, error) {
	urlPath := "/policy"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Interval != nil {
		reqQP.Set("interval", *options.Interval)
	}
	if client.clientOptionalGroup != nil && client.clientOptionalGroup.OptionalVersion != nil {
		reqQP.Set("optional-version", *client.clientOptionalGroup.OptionalVersion)
	}
	reqQP.Set("things", strings.Join(strings.Fields(strings.Trim(fmt.Sprint(things), "[]")), ","))
	if options != nil && options.Unique != nil {
		reqQP.Set("unique", *options.Unique)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if client.clientOptionalGroup != nil && client.clientOptionalGroup.OptionalIndex != nil {
		req.Raw().Header["optional-index"] = []string{strconv.FormatInt(int64(*client.clientOptionalGroup.OptionalIndex), 10)}
	}
	if client.optionalString != nil {
		req.Raw().Header["optional-string"] = []string{*client.optionalString}
	}
	if err := runtime.MarshalAsJSON(req, polymorphicParam); err != nil {
		return nil, err
	}
	return req, nil
}

// policyAssignmentHandleResponse handles the PolicyAssignment response.
func (client *Client) policyAssignmentHandleResponse(resp *http.Response) (PolicyAssignmentResponse, error) {
	result := PolicyAssignmentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyAssignmentProperties); err != nil {
		return PolicyAssignmentResponse{}, err
	}
	return result, nil
}

// UploadForm -
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0
//   - options - UploadFormOptions contains the optional parameters for the Client.UploadForm method.
func (client *Client) UploadForm(ctx context.Context, requiredString string, requiredEnum DataSetting, requiredInt int32, options *UploadFormOptions) (UploadFormResponse, error) {
	var err error
	const operationName = "Client.UploadForm"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.uploadFormCreateRequest(ctx, requiredString, requiredEnum, requiredInt, options)
	if err != nil {
		return UploadFormResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UploadFormResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return UploadFormResponse{}, err
	}
	return UploadFormResponse{}, nil
}

// uploadFormCreateRequest creates the UploadForm request.
func (client *Client) uploadFormCreateRequest(ctx context.Context, requiredString string, requiredEnum DataSetting, requiredInt int32, options *UploadFormOptions) (*policy.Request, error) {
	urlPath := "/formdata"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	formData := map[string]any{}
	formData["requiredString"] = requiredString
	if options != nil && options.OptionalString != nil {
		formData["OptionalString"] = *options.OptionalString
	}
	formData["requiredEnum"] = requiredEnum
	formData["requiredInt"] = requiredInt
	if options != nil && options.OptionalBool != nil {
		formData["OptionalBool"] = *options.OptionalBool
	}
	if options != nil && options.OptionalIntEnum != nil {
		formData["OptionalIntEnum"] = *options.OptionalIntEnum
	}
	if err := runtime.SetMultipartFormData(req, formData); err != nil {
		return nil, err
	}
	return req, nil
}

// listLRONextCreateRequest creates the listLRONextCreateRequest request.
func (client *Client) listLRONextCreateRequest(ctx context.Context, nextLink string) (*policy.Request, error) {
	urlPath := "/paged"
	urlPath = strings.ReplaceAll(urlPath, "{nextLink}", nextLink)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listWithSharedNextCreateRequest creates the listWithSharedNextCreateRequest request.
func (client *Client) listWithSharedNextCreateRequest(ctx context.Context, nextLink string) (*policy.Request, error) {
	urlPath := "/listWithSharedNext"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["nextLink"] = []string{nextLink}
	return req, nil
}
