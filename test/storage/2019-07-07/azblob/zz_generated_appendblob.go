// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azblob

import (
	"context"
	"encoding/base64"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type appendBlobClient struct {
	con *connection
}

// Pipeline returns the pipeline associated with this client.
func (client appendBlobClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// AppendBlock - The Append Block operation commits a new block of data to the end of an existing append blob. The Append Block operation is permitted only
// if the blob was created with x-ms-blob-type set to
// AppendBlob. Append Block is supported only on version 2015-02-21 version or later.
func (client appendBlobClient) AppendBlock(ctx context.Context, contentLength int64, body azcore.ReadSeekCloser, appendBlobAppendBlockOptions *AppendBlobAppendBlockOptions, leaseAccessConditions *LeaseAccessConditions, appendPositionAccessConditions *AppendPositionAccessConditions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, modifiedAccessConditions *ModifiedAccessConditions) (AppendBlobAppendBlockResponse, error) {
	req, err := client.appendBlockCreateRequest(ctx, contentLength, body, appendBlobAppendBlockOptions, leaseAccessConditions, appendPositionAccessConditions, cpkInfo, cpkScopeInfo, modifiedAccessConditions)
	if err != nil {
		return AppendBlobAppendBlockResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return AppendBlobAppendBlockResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusCreated) {
		return AppendBlobAppendBlockResponse{}, client.appendBlockHandleError(resp)
	}
	return client.appendBlockHandleResponse(resp)
}

// appendBlockCreateRequest creates the AppendBlock request.
func (client appendBlobClient) appendBlockCreateRequest(ctx context.Context, contentLength int64, body azcore.ReadSeekCloser, appendBlobAppendBlockOptions *AppendBlobAppendBlockOptions, leaseAccessConditions *LeaseAccessConditions, appendPositionAccessConditions *AppendPositionAccessConditions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, modifiedAccessConditions *ModifiedAccessConditions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPut, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("comp", "appendblock")
	if appendBlobAppendBlockOptions != nil && appendBlobAppendBlockOptions.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*appendBlobAppendBlockOptions.Timeout), 10))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	if appendBlobAppendBlockOptions != nil && appendBlobAppendBlockOptions.TransactionalContentMd5 != nil {
		req.Header.Set("Content-MD5", base64.StdEncoding.EncodeToString(*appendBlobAppendBlockOptions.TransactionalContentMd5))
	}
	if appendBlobAppendBlockOptions != nil && appendBlobAppendBlockOptions.TransactionalContentCrc64 != nil {
		req.Header.Set("x-ms-content-crc64", base64.StdEncoding.EncodeToString(*appendBlobAppendBlockOptions.TransactionalContentCrc64))
	}
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseId != nil {
		req.Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseId)
	}
	if appendPositionAccessConditions != nil && appendPositionAccessConditions.MaxSize != nil {
		req.Header.Set("x-ms-blob-condition-maxsize", strconv.FormatInt(*appendPositionAccessConditions.MaxSize, 10))
	}
	if appendPositionAccessConditions != nil && appendPositionAccessConditions.AppendPosition != nil {
		req.Header.Set("x-ms-blob-condition-appendpos", strconv.FormatInt(*appendPositionAccessConditions.AppendPosition, 10))
	}
	if cpkInfo != nil && cpkInfo.EncryptionKey != nil {
		req.Header.Set("x-ms-encryption-key", *cpkInfo.EncryptionKey)
	}
	if cpkInfo != nil && cpkInfo.EncryptionKeySha256 != nil {
		req.Header.Set("x-ms-encryption-key-sha256", *cpkInfo.EncryptionKeySha256)
	}
	if cpkInfo != nil && cpkInfo.EncryptionAlgorithm != nil {
		req.Header.Set("x-ms-encryption-algorithm", "AES256")
	}
	if cpkScopeInfo != nil && cpkScopeInfo.EncryptionScope != nil {
		req.Header.Set("x-ms-encryption-scope", *cpkScopeInfo.EncryptionScope)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Header.Set("If-Modified-Since", modifiedAccessConditions.IfModifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", modifiedAccessConditions.IfUnmodifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Header.Set("If-Match", *modifiedAccessConditions.IfMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *modifiedAccessConditions.IfNoneMatch)
	}
	req.Header.Set("x-ms-version", "2019-07-07")
	if appendBlobAppendBlockOptions != nil && appendBlobAppendBlockOptions.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *appendBlobAppendBlockOptions.RequestId)
	}
	req.Header.Set("Accept", "application/xml")
	return req, req.SetBody(body, "application/octet-stream")
}

// appendBlockHandleResponse handles the AppendBlock response.
func (client appendBlobClient) appendBlockHandleResponse(resp *azcore.Response) (AppendBlobAppendBlockResponse, error) {
	result := AppendBlobAppendBlockResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return AppendBlobAppendBlockResponse{}, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("Content-MD5"); val != "" {
		contentMd5, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return AppendBlobAppendBlockResponse{}, err
		}
		result.ContentMD5 = &contentMd5
	}
	if val := resp.Header.Get("x-ms-content-crc64"); val != "" {
		contentCrc64, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return AppendBlobAppendBlockResponse{}, err
		}
		result.ContentCRC64 = &contentCrc64
	}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return AppendBlobAppendBlockResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-blob-append-offset"); val != "" {
		result.BlobAppendOffset = &val
	}
	if val := resp.Header.Get("x-ms-blob-committed-block-count"); val != "" {
		blobCommittedBlockCount32, err := strconv.ParseInt(val, 10, 32)
		blobCommittedBlockCount := int32(blobCommittedBlockCount32)
		if err != nil {
			return AppendBlobAppendBlockResponse{}, err
		}
		result.BlobCommittedBlockCount = &blobCommittedBlockCount
	}
	if val := resp.Header.Get("x-ms-request-server-encrypted"); val != "" {
		isServerEncrypted, err := strconv.ParseBool(val)
		if err != nil {
			return AppendBlobAppendBlockResponse{}, err
		}
		result.IsServerEncrypted = &isServerEncrypted
	}
	if val := resp.Header.Get("x-ms-encryption-key-sha256"); val != "" {
		result.EncryptionKeySHA256 = &val
	}
	if val := resp.Header.Get("x-ms-encryption-scope"); val != "" {
		result.EncryptionScope = &val
	}
	return result, nil
}

// appendBlockHandleError handles the AppendBlock error response.
func (client appendBlobClient) appendBlockHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// AppendBlockFromURL - The Append Block operation commits a new block of data to the end of an existing append blob where the contents are read from a
// source url. The Append Block operation is permitted only if the blob was
// created with x-ms-blob-type set to AppendBlob. Append Block is supported only on version 2015-02-21 version or later.
func (client appendBlobClient) AppendBlockFromURL(ctx context.Context, sourceUrl url.URL, contentLength int64, appendBlobAppendBlockFromUrlOptions *AppendBlobAppendBlockFromURLOptions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, leaseAccessConditions *LeaseAccessConditions, appendPositionAccessConditions *AppendPositionAccessConditions, modifiedAccessConditions *ModifiedAccessConditions, sourceModifiedAccessConditions *SourceModifiedAccessConditions) (AppendBlobAppendBlockFromURLResponse, error) {
	req, err := client.appendBlockFromUrlCreateRequest(ctx, sourceUrl, contentLength, appendBlobAppendBlockFromUrlOptions, cpkInfo, cpkScopeInfo, leaseAccessConditions, appendPositionAccessConditions, modifiedAccessConditions, sourceModifiedAccessConditions)
	if err != nil {
		return AppendBlobAppendBlockFromURLResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return AppendBlobAppendBlockFromURLResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusCreated) {
		return AppendBlobAppendBlockFromURLResponse{}, client.appendBlockFromUrlHandleError(resp)
	}
	return client.appendBlockFromUrlHandleResponse(resp)
}

// appendBlockFromUrlCreateRequest creates the AppendBlockFromURL request.
func (client appendBlobClient) appendBlockFromUrlCreateRequest(ctx context.Context, sourceUrl url.URL, contentLength int64, appendBlobAppendBlockFromUrlOptions *AppendBlobAppendBlockFromURLOptions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, leaseAccessConditions *LeaseAccessConditions, appendPositionAccessConditions *AppendPositionAccessConditions, modifiedAccessConditions *ModifiedAccessConditions, sourceModifiedAccessConditions *SourceModifiedAccessConditions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPut, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("comp", "appendblock")
	if appendBlobAppendBlockFromUrlOptions != nil && appendBlobAppendBlockFromUrlOptions.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*appendBlobAppendBlockFromUrlOptions.Timeout), 10))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("x-ms-copy-source", sourceUrl.String())
	if appendBlobAppendBlockFromUrlOptions != nil && appendBlobAppendBlockFromUrlOptions.SourceRange != nil {
		req.Header.Set("x-ms-source-range", *appendBlobAppendBlockFromUrlOptions.SourceRange)
	}
	if appendBlobAppendBlockFromUrlOptions != nil && appendBlobAppendBlockFromUrlOptions.SourceContentMd5 != nil {
		req.Header.Set("x-ms-source-content-md5", base64.StdEncoding.EncodeToString(*appendBlobAppendBlockFromUrlOptions.SourceContentMd5))
	}
	if appendBlobAppendBlockFromUrlOptions != nil && appendBlobAppendBlockFromUrlOptions.SourceContentcrc64 != nil {
		req.Header.Set("x-ms-source-content-crc64", base64.StdEncoding.EncodeToString(*appendBlobAppendBlockFromUrlOptions.SourceContentcrc64))
	}
	req.Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	if appendBlobAppendBlockFromUrlOptions != nil && appendBlobAppendBlockFromUrlOptions.TransactionalContentMd5 != nil {
		req.Header.Set("Content-MD5", base64.StdEncoding.EncodeToString(*appendBlobAppendBlockFromUrlOptions.TransactionalContentMd5))
	}
	if cpkInfo != nil && cpkInfo.EncryptionKey != nil {
		req.Header.Set("x-ms-encryption-key", *cpkInfo.EncryptionKey)
	}
	if cpkInfo != nil && cpkInfo.EncryptionKeySha256 != nil {
		req.Header.Set("x-ms-encryption-key-sha256", *cpkInfo.EncryptionKeySha256)
	}
	if cpkInfo != nil && cpkInfo.EncryptionAlgorithm != nil {
		req.Header.Set("x-ms-encryption-algorithm", "AES256")
	}
	if cpkScopeInfo != nil && cpkScopeInfo.EncryptionScope != nil {
		req.Header.Set("x-ms-encryption-scope", *cpkScopeInfo.EncryptionScope)
	}
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseId != nil {
		req.Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseId)
	}
	if appendPositionAccessConditions != nil && appendPositionAccessConditions.MaxSize != nil {
		req.Header.Set("x-ms-blob-condition-maxsize", strconv.FormatInt(*appendPositionAccessConditions.MaxSize, 10))
	}
	if appendPositionAccessConditions != nil && appendPositionAccessConditions.AppendPosition != nil {
		req.Header.Set("x-ms-blob-condition-appendpos", strconv.FormatInt(*appendPositionAccessConditions.AppendPosition, 10))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Header.Set("If-Modified-Since", modifiedAccessConditions.IfModifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", modifiedAccessConditions.IfUnmodifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Header.Set("If-Match", *modifiedAccessConditions.IfMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *modifiedAccessConditions.IfNoneMatch)
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfModifiedSince != nil {
		req.Header.Set("x-ms-source-if-modified-since", sourceModifiedAccessConditions.SourceIfModifiedSince.Format(time.RFC1123))
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfUnmodifiedSince != nil {
		req.Header.Set("x-ms-source-if-unmodified-since", sourceModifiedAccessConditions.SourceIfUnmodifiedSince.Format(time.RFC1123))
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfMatch != nil {
		req.Header.Set("x-ms-source-if-match", *sourceModifiedAccessConditions.SourceIfMatch)
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfNoneMatch != nil {
		req.Header.Set("x-ms-source-if-none-match", *sourceModifiedAccessConditions.SourceIfNoneMatch)
	}
	req.Header.Set("x-ms-version", "2019-07-07")
	if appendBlobAppendBlockFromUrlOptions != nil && appendBlobAppendBlockFromUrlOptions.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *appendBlobAppendBlockFromUrlOptions.RequestId)
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// appendBlockFromUrlHandleResponse handles the AppendBlockFromURL response.
func (client appendBlobClient) appendBlockFromUrlHandleResponse(resp *azcore.Response) (AppendBlobAppendBlockFromURLResponse, error) {
	result := AppendBlobAppendBlockFromURLResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return AppendBlobAppendBlockFromURLResponse{}, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("Content-MD5"); val != "" {
		contentMd5, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return AppendBlobAppendBlockFromURLResponse{}, err
		}
		result.ContentMD5 = &contentMd5
	}
	if val := resp.Header.Get("x-ms-content-crc64"); val != "" {
		contentCrc64, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return AppendBlobAppendBlockFromURLResponse{}, err
		}
		result.ContentCRC64 = &contentCrc64
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return AppendBlobAppendBlockFromURLResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-blob-append-offset"); val != "" {
		result.BlobAppendOffset = &val
	}
	if val := resp.Header.Get("x-ms-blob-committed-block-count"); val != "" {
		blobCommittedBlockCount32, err := strconv.ParseInt(val, 10, 32)
		blobCommittedBlockCount := int32(blobCommittedBlockCount32)
		if err != nil {
			return AppendBlobAppendBlockFromURLResponse{}, err
		}
		result.BlobCommittedBlockCount = &blobCommittedBlockCount
	}
	if val := resp.Header.Get("x-ms-encryption-key-sha256"); val != "" {
		result.EncryptionKeySHA256 = &val
	}
	if val := resp.Header.Get("x-ms-encryption-scope"); val != "" {
		result.EncryptionScope = &val
	}
	if val := resp.Header.Get("x-ms-request-server-encrypted"); val != "" {
		isServerEncrypted, err := strconv.ParseBool(val)
		if err != nil {
			return AppendBlobAppendBlockFromURLResponse{}, err
		}
		result.IsServerEncrypted = &isServerEncrypted
	}
	return result, nil
}

// appendBlockFromUrlHandleError handles the AppendBlockFromURL error response.
func (client appendBlobClient) appendBlockFromUrlHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Create - The Create Append Blob operation creates a new append blob.
func (client appendBlobClient) Create(ctx context.Context, contentLength int64, appendBlobCreateOptions *AppendBlobCreateOptions, blobHttpHeaders *BlobHttpHeaders, leaseAccessConditions *LeaseAccessConditions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, modifiedAccessConditions *ModifiedAccessConditions) (AppendBlobCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, contentLength, appendBlobCreateOptions, blobHttpHeaders, leaseAccessConditions, cpkInfo, cpkScopeInfo, modifiedAccessConditions)
	if err != nil {
		return AppendBlobCreateResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return AppendBlobCreateResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusCreated) {
		return AppendBlobCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client appendBlobClient) createCreateRequest(ctx context.Context, contentLength int64, appendBlobCreateOptions *AppendBlobCreateOptions, blobHttpHeaders *BlobHttpHeaders, leaseAccessConditions *LeaseAccessConditions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, modifiedAccessConditions *ModifiedAccessConditions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPut, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	if appendBlobCreateOptions != nil && appendBlobCreateOptions.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*appendBlobCreateOptions.Timeout), 10))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("x-ms-blob-type", "AppendBlob")
	req.Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	if blobHttpHeaders != nil && blobHttpHeaders.BlobContentType != nil {
		req.Header.Set("x-ms-blob-content-type", *blobHttpHeaders.BlobContentType)
	}
	if blobHttpHeaders != nil && blobHttpHeaders.BlobContentEncoding != nil {
		req.Header.Set("x-ms-blob-content-encoding", *blobHttpHeaders.BlobContentEncoding)
	}
	if blobHttpHeaders != nil && blobHttpHeaders.BlobContentLanguage != nil {
		req.Header.Set("x-ms-blob-content-language", *blobHttpHeaders.BlobContentLanguage)
	}
	if blobHttpHeaders != nil && blobHttpHeaders.BlobContentMd5 != nil {
		req.Header.Set("x-ms-blob-content-md5", base64.StdEncoding.EncodeToString(*blobHttpHeaders.BlobContentMd5))
	}
	if blobHttpHeaders != nil && blobHttpHeaders.BlobCacheControl != nil {
		req.Header.Set("x-ms-blob-cache-control", *blobHttpHeaders.BlobCacheControl)
	}
	if appendBlobCreateOptions != nil && appendBlobCreateOptions.Metadata != nil {
		for k, v := range *appendBlobCreateOptions.Metadata {
			req.Header.Set("x-ms-meta-"+k, v)
		}
	}
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseId != nil {
		req.Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseId)
	}
	if blobHttpHeaders != nil && blobHttpHeaders.BlobContentDisposition != nil {
		req.Header.Set("x-ms-blob-content-disposition", *blobHttpHeaders.BlobContentDisposition)
	}
	if cpkInfo != nil && cpkInfo.EncryptionKey != nil {
		req.Header.Set("x-ms-encryption-key", *cpkInfo.EncryptionKey)
	}
	if cpkInfo != nil && cpkInfo.EncryptionKeySha256 != nil {
		req.Header.Set("x-ms-encryption-key-sha256", *cpkInfo.EncryptionKeySha256)
	}
	if cpkInfo != nil && cpkInfo.EncryptionAlgorithm != nil {
		req.Header.Set("x-ms-encryption-algorithm", "AES256")
	}
	if cpkScopeInfo != nil && cpkScopeInfo.EncryptionScope != nil {
		req.Header.Set("x-ms-encryption-scope", *cpkScopeInfo.EncryptionScope)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Header.Set("If-Modified-Since", modifiedAccessConditions.IfModifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", modifiedAccessConditions.IfUnmodifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Header.Set("If-Match", *modifiedAccessConditions.IfMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *modifiedAccessConditions.IfNoneMatch)
	}
	req.Header.Set("x-ms-version", "2019-07-07")
	if appendBlobCreateOptions != nil && appendBlobCreateOptions.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *appendBlobCreateOptions.RequestId)
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// createHandleResponse handles the Create response.
func (client appendBlobClient) createHandleResponse(resp *azcore.Response) (AppendBlobCreateResponse, error) {
	result := AppendBlobCreateResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return AppendBlobCreateResponse{}, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("Content-MD5"); val != "" {
		contentMd5, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return AppendBlobCreateResponse{}, err
		}
		result.ContentMD5 = &contentMd5
	}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return AppendBlobCreateResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-request-server-encrypted"); val != "" {
		isServerEncrypted, err := strconv.ParseBool(val)
		if err != nil {
			return AppendBlobCreateResponse{}, err
		}
		result.IsServerEncrypted = &isServerEncrypted
	}
	if val := resp.Header.Get("x-ms-encryption-key-sha256"); val != "" {
		result.EncryptionKeySHA256 = &val
	}
	if val := resp.Header.Get("x-ms-encryption-scope"); val != "" {
		result.EncryptionScope = &val
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client appendBlobClient) createHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
