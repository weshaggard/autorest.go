//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azacr

import "time"

type AccessToken struct {
	// The access token for performing authenticated requests
	AccessToken *string
}

// Annotations - Additional information provided through arbitrary metadata.
type Annotations struct {
	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]any

	// Contact details of the people or organization responsible for the image.
	Authors *string

	// Date and time on which the image was built (string, date-time as defined by https://tools.ietf.org/html/rfc3339#section-5.6)
	Created *time.Time

	// Human-readable description of the software packaged in the image
	Description *string

	// URL to get documentation on the image.
	Documentation *string

	// License(s) under which contained software is distributed as an SPDX License Expression.
	Licenses *string

	// Name of the reference for a target.
	Name *string

	// Source control revision identifier for the packaged software.
	Revision *string

	// URL to get source code for building the image.
	Source *string

	// Human-readable title of the image
	Title *string

	// URL to find more information on the image.
	URL *string

	// Name of the distributing entity, organization or individual.
	Vendor *string

	// Version of the packaged software. The version MAY match a label or tag in the source code repository, may also be Semantic
	// versioning-compatible
	Version *string
}

// ArtifactManifestPlatform - The artifact's platform, consisting of operating system and architecture.
type ArtifactManifestPlatform struct {
	// READ-ONLY; Manifest digest
	Digest *string

	// READ-ONLY; CPU architecture
	Architecture *ArtifactArchitecture

	// READ-ONLY; Operating system
	OperatingSystem *ArtifactOperatingSystem
}

// ArtifactManifestProperties - Manifest attributes details
type ArtifactManifestProperties struct {
	// READ-ONLY; Manifest attributes
	Manifest *ManifestAttributesBase

	// READ-ONLY; Registry login server name. This is likely to be similar to {registry-name}.azurecr.io.
	RegistryLoginServer *string

	// READ-ONLY; Repository name
	RepositoryName *string
}

// ArtifactTagProperties - Tag attributes
type ArtifactTagProperties struct {
	// READ-ONLY; Registry login server name. This is likely to be similar to {registry-name}.azurecr.io.
	RegistryLoginServer *string

	// READ-ONLY; Image name
	RepositoryName *string

	// READ-ONLY; List of tag attribute details
	Tag *TagAttributesBase
}

// AuthenticationClientExchangeAADAccessTokenForAcrRefreshTokenOptions contains the optional parameters for the AuthenticationClient.ExchangeAADAccessTokenForAcrRefreshToken
// method.
type AuthenticationClientExchangeAADAccessTokenForAcrRefreshTokenOptions struct {
	// AAD access token, mandatory when granttype is accesstokenrefreshtoken or access_token.
	AccessToken *string

	// AAD refresh token, mandatory when granttype is accesstokenrefreshtoken or refresh_token
	RefreshToken *string

	// AAD tenant associated to the AAD credentials.
	Tenant *string
}

// AuthenticationClientExchangeAcrRefreshTokenForAcrAccessTokenOptions contains the optional parameters for the AuthenticationClient.ExchangeAcrRefreshTokenForAcrAccessToken
// method.
type AuthenticationClientExchangeAcrRefreshTokenForAcrAccessTokenOptions struct {
	// Grant type is expected to be refresh_token
	GrantType *TokenGrantType
}

// AuthenticationClientGetAcrAccessTokenFromLoginOptions contains the optional parameters for the AuthenticationClient.GetAcrAccessTokenFromLogin
// method.
type AuthenticationClientGetAcrAccessTokenFromLoginOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryBlobClientCancelUploadOptions contains the optional parameters for the ContainerRegistryBlobClient.CancelUpload
// method.
type ContainerRegistryBlobClientCancelUploadOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryBlobClientCheckBlobExistsOptions contains the optional parameters for the ContainerRegistryBlobClient.CheckBlobExists
// method.
type ContainerRegistryBlobClientCheckBlobExistsOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryBlobClientCheckChunkExistsOptions contains the optional parameters for the ContainerRegistryBlobClient.CheckChunkExists
// method.
type ContainerRegistryBlobClientCheckChunkExistsOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryBlobClientCompleteUploadOptions contains the optional parameters for the ContainerRegistryBlobClient.CompleteUpload
// method.
type ContainerRegistryBlobClientCompleteUploadOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryBlobClientDeleteBlobOptions contains the optional parameters for the ContainerRegistryBlobClient.DeleteBlob
// method.
type ContainerRegistryBlobClientDeleteBlobOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryBlobClientGetBlobOptions contains the optional parameters for the ContainerRegistryBlobClient.GetBlob
// method.
type ContainerRegistryBlobClientGetBlobOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryBlobClientGetChunkOptions contains the optional parameters for the ContainerRegistryBlobClient.GetChunk
// method.
type ContainerRegistryBlobClientGetChunkOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryBlobClientGetUploadStatusOptions contains the optional parameters for the ContainerRegistryBlobClient.GetUploadStatus
// method.
type ContainerRegistryBlobClientGetUploadStatusOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryBlobClientMountBlobOptions contains the optional parameters for the ContainerRegistryBlobClient.MountBlob
// method.
type ContainerRegistryBlobClientMountBlobOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryBlobClientStartUploadOptions contains the optional parameters for the ContainerRegistryBlobClient.StartUpload
// method.
type ContainerRegistryBlobClientStartUploadOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryBlobClientUploadChunkOptions contains the optional parameters for the ContainerRegistryBlobClient.UploadChunk
// method.
type ContainerRegistryBlobClientUploadChunkOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryClientCheckDockerV2SupportOptions contains the optional parameters for the ContainerRegistryClient.CheckDockerV2Support
// method.
type ContainerRegistryClientCheckDockerV2SupportOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryClientCreateManifestOptions contains the optional parameters for the ContainerRegistryClient.CreateManifest
// method.
type ContainerRegistryClientCreateManifestOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryClientDeleteManifestOptions contains the optional parameters for the ContainerRegistryClient.DeleteManifest
// method.
type ContainerRegistryClientDeleteManifestOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryClientDeleteRepositoryOptions contains the optional parameters for the ContainerRegistryClient.DeleteRepository
// method.
type ContainerRegistryClientDeleteRepositoryOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryClientDeleteTagOptions contains the optional parameters for the ContainerRegistryClient.DeleteTag method.
type ContainerRegistryClientDeleteTagOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryClientGetManifestOptions contains the optional parameters for the ContainerRegistryClient.GetManifest
// method.
type ContainerRegistryClientGetManifestOptions struct {
	// Accept header string delimited by comma. For example, application/vnd.docker.distribution.manifest.v2+json
	Accept *string
}

// ContainerRegistryClientGetManifestPropertiesOptions contains the optional parameters for the ContainerRegistryClient.GetManifestProperties
// method.
type ContainerRegistryClientGetManifestPropertiesOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryClientGetManifestsOptions contains the optional parameters for the ContainerRegistryClient.NewGetManifestsPager
// method.
type ContainerRegistryClientGetManifestsOptions struct {
	// Query parameter for the last item in previous query. Result set will include values lexically after last.
	Last *string

	// query parameter for max number of items
	N *int32

	// orderby query parameter
	Orderby *string
}

// ContainerRegistryClientGetPropertiesOptions contains the optional parameters for the ContainerRegistryClient.GetProperties
// method.
type ContainerRegistryClientGetPropertiesOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryClientGetRepositoriesOptions contains the optional parameters for the ContainerRegistryClient.NewGetRepositoriesPager
// method.
type ContainerRegistryClientGetRepositoriesOptions struct {
	// Query parameter for the last item in previous query. Result set will include values lexically after last.
	Last *string

	// query parameter for max number of items
	N *int32
}

// ContainerRegistryClientGetTagPropertiesOptions contains the optional parameters for the ContainerRegistryClient.GetTagProperties
// method.
type ContainerRegistryClientGetTagPropertiesOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryClientGetTagsOptions contains the optional parameters for the ContainerRegistryClient.NewGetTagsPager
// method.
type ContainerRegistryClientGetTagsOptions struct {
	// filter by digest
	Digest *string

	// Query parameter for the last item in previous query. Result set will include values lexically after last.
	Last *string

	// query parameter for max number of items
	N *int32

	// orderby query parameter
	Orderby *string
}

// ContainerRegistryClientUpdateManifestPropertiesOptions contains the optional parameters for the ContainerRegistryClient.UpdateManifestProperties
// method.
type ContainerRegistryClientUpdateManifestPropertiesOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryClientUpdatePropertiesOptions contains the optional parameters for the ContainerRegistryClient.UpdateProperties
// method.
type ContainerRegistryClientUpdatePropertiesOptions struct {
	// placeholder for future optional parameters
}

// ContainerRegistryClientUpdateTagAttributesOptions contains the optional parameters for the ContainerRegistryClient.UpdateTagAttributes
// method.
type ContainerRegistryClientUpdateTagAttributesOptions struct {
	// placeholder for future optional parameters
}

// ContainerRepositoryProperties - Properties of this repository.
type ContainerRepositoryProperties struct {
	// REQUIRED; Writeable properties of the resource
	ChangeableAttributes *RepositoryWriteableProperties

	// READ-ONLY; Image created time
	CreatedOn *time.Time

	// READ-ONLY; Image last update time
	LastUpdatedOn *time.Time

	// READ-ONLY; Number of the manifests
	ManifestCount *int32

	// READ-ONLY; Image name
	Name *string

	// READ-ONLY; Registry login server name. This is likely to be similar to {registry-name}.azurecr.io.
	RegistryLoginServer *string

	// READ-ONLY; Number of the tags
	TagCount *int32
}

// DeleteRepositoryResult - Deleted repository
type DeleteRepositoryResult struct {
	// READ-ONLY; SHA of the deleted image
	DeletedManifests []*string

	// READ-ONLY; Tag of the deleted image
	DeletedTags []*string
}

// Descriptor - Docker V2 image layer descriptor including config and layers
type Descriptor struct {
	// Additional information provided through arbitrary metadata.
	Annotations *Annotations

	// Layer digest
	Digest *string

	// Layer media type
	MediaType *string

	// Layer size
	Size *int64

	// Specifies a list of URIs from which this object may be downloaded.
	Urls []*string
}

// ErrorInfo - Error information
type ErrorInfo struct {
	// Error code
	Code *string

	// Error details
	Detail []byte

	// Error message
	Message *string
}

// Errors - Acr error response describing why the operation failed
type Errors struct {
	// Array of detailed error
	Errors []*ErrorInfo
}

// FsLayer - Image layer information
type FsLayer struct {
	// SHA of an image layer
	BlobSum *string
}

// History - A list of unstructured historical data for v1 compatibility
type History struct {
	// The raw v1 compatibility information
	V1Compatibility *string
}

// ImageSignature - Signature of a signed manifest
type ImageSignature struct {
	// A JSON web signature
	Header *JWK

	// The signed protected header
	Protected *string

	// A signature for the image manifest, signed by a libtrust private key
	Signature *string
}

// JWK - A JSON web signature
type JWK struct {
	// The algorithm used to sign or encrypt the JWT
	Alg *string

	// JSON web key parameter
	Jwk *JWKHeader
}

// JWKHeader - JSON web key parameter
type JWKHeader struct {
	// crv value
	Crv *string

	// kid value
	Kid *string

	// kty value
	Kty *string

	// x value
	X *string

	// y value
	Y *string
}

// Manifest - Returns the requested manifest file
type Manifest struct {
	// Schema version
	SchemaVersion *int32
}

// ManifestAttributesBase - Manifest details
type ManifestAttributesBase struct {
	// READ-ONLY; Created time
	CreatedOn *time.Time

	// READ-ONLY; Manifest
	Digest *string

	// READ-ONLY; Last update time
	LastUpdatedOn *time.Time

	// Writeable properties of the resource
	ChangeableAttributes *ManifestWriteableProperties

	// Config blob media type
	ConfigMediaType *string

	// READ-ONLY; CPU architecture
	Architecture *ArtifactArchitecture

	// READ-ONLY; Operating system
	OperatingSystem *ArtifactOperatingSystem

	// READ-ONLY; List of artifacts that are referenced by this manifest list, with information about the platform each supports.
	// This list will be empty if this is a leaf manifest and not a manifest list.
	RelatedArtifacts []*ArtifactManifestPlatform

	// READ-ONLY; Image size
	Size *int64

	// READ-ONLY; List of tags
	Tags []*string
}

// ManifestAttributesManifest - List of manifest attributes
type ManifestAttributesManifest struct {
	// List of manifest attributes details
	References []*ArtifactManifestPlatform
}

// ManifestList - Returns the requested Docker multi-arch-manifest file
type ManifestList struct {
	// List of V2 image layer information
	Manifests []*ManifestListAttributes

	// Media type for this Manifest
	MediaType *string

	// Schema version
	SchemaVersion *int32
}

type ManifestListAttributes struct {
	// The digest of the content, as defined by the Registry V2 HTTP API Specification
	Digest *string

	// The MIME type of the referenced object. This will generally be application/vnd.docker.image.manifest.v2+json, but it could
	// also be application/vnd.docker.image.manifest.v1+json
	MediaType *string

	// The platform object describes the platform which the image in the manifest runs on. A full list of valid operating system
	// and architecture values are listed in the Go language documentation for $GOOS
	// and $GOARCH
	Platform *Platform

	// The size in bytes of the object
	Size *int64
}

// ManifestWrapper - Returns the requested manifest file
type ManifestWrapper struct {
	// (OCI, OCIIndex) Additional metadata
	Annotations *Annotations

	// (V1) CPU architecture
	Architecture *string

	// (V2, OCI) Image config descriptor
	Config *Descriptor

	// (V1) List of layer information
	FsLayers []*FsLayer

	// (V1) Image history
	History []*History

	// (V2, OCI) List of V2 image layer information
	Layers []*Descriptor

	// (ManifestList, OCIIndex) List of V2 image layer information
	Manifests []*ManifestListAttributes

	// Media type for this Manifest
	MediaType *string

	// (V1) Image name
	Name *string

	// Schema version
	SchemaVersion *int32

	// (V1) Image signature
	Signatures []*ImageSignature

	// (V1) Image tag
	Tag *string
}

// ManifestWriteableProperties - Changeable attributes
type ManifestWriteableProperties struct {
	// Delete enabled
	CanDelete *bool

	// List enabled
	CanList *bool

	// Read enabled
	CanRead *bool

	// Write enabled
	CanWrite *bool
}

// Manifests - Manifest attributes
type Manifests struct {
	Link *string

	// List of manifests
	Manifests []*ManifestAttributesBase

	// Registry login server name. This is likely to be similar to {registry-name}.azurecr.io.
	RegistryLoginServer *string

	// Image name
	Repository *string
}

// OCIIndex - Returns the requested OCI index file
type OCIIndex struct {
	// Additional information provided through arbitrary metadata.
	Annotations *Annotations

	// List of OCI image layer information
	Manifests []*ManifestListAttributes

	// Schema version
	SchemaVersion *int32
}

// OCIManifest - Returns the requested OCI Manifest file
type OCIManifest struct {
	// Additional information provided through arbitrary metadata.
	Annotations *Annotations

	// V2 image config descriptor
	Config *Descriptor

	// List of V2 image layer information
	Layers []*Descriptor

	// Schema version
	SchemaVersion *int32
}

type Paths108HwamOauth2ExchangePostRequestbodyContentApplicationXWwwFormUrlencodedSchema struct {
	// REQUIRED; Can take a value of accesstokenrefreshtoken, or accesstoken, or refresh_token
	GrantType *PostContentSchemaGrantType

	// REQUIRED; Indicates the name of your Azure container registry.
	Service *string

	// AAD access token, mandatory when granttype is accesstokenrefreshtoken or access_token.
	AADAccessToken *string

	// AAD refresh token, mandatory when granttype is accesstokenrefreshtoken or refresh_token
	RefreshToken *string

	// AAD tenant associated to the AAD credentials.
	Tenant *string
}

type PathsV3R3RxOauth2TokenPostRequestbodyContentApplicationXWwwFormUrlencodedSchema struct {
	// REQUIRED; Must be a valid ACR refresh token
	AcrRefreshToken *string

	// REQUIRED; Grant type is expected to be refresh_token
	GrantType *TokenGrantType

	// REQUIRED; Which is expected to be a valid scope, and can be specified more than once for multiple scope requests. You obtained
	// this from the Www-Authenticate response header from the challenge.
	Scope *string

	// REQUIRED; Indicates the name of your Azure container registry.
	Service *string
}

// Platform - The platform object describes the platform which the image in the manifest runs on. A full list of valid operating
// system and architecture values are listed in the Go language documentation for $GOOS
// and $GOARCH
type Platform struct {
	// Specifies the CPU architecture, for example amd64 or ppc64le.
	Architecture *string

	// The optional features field specifies an array of strings, each listing a required CPU feature (for example sse4 or aes
	Features []*string

	// The os field specifies the operating system, for example linux or windows.
	OS *string

	// The optional os.features field specifies an array of strings, each listing a required OS feature (for example on Windows
	// win32k
	OSFeatures []*string

	// The optional os.version field specifies the operating system version, for example 10.0.10586.
	OSVersion *string

	// The optional variant field specifies a variant of the CPU, for example armv6l to specify a particular CPU variant of the
	// ARM CPU.
	Variant *string
}

type RefreshToken struct {
	// The refresh token to be used for generating access tokens
	RefreshToken *string
}

// Repositories - List of repositories
type Repositories struct {
	Link *string

	// Repository names
	Repositories []*string
}

// RepositoryTags - Result of the request to list tags of the image
type RepositoryTags struct {
	// Name of the image
	Name *string

	// List of tags
	Tags []*string
}

// RepositoryWriteableProperties - Changeable attributes for Repository
type RepositoryWriteableProperties struct {
	// Delete enabled
	CanDelete *bool

	// List enabled
	CanList *bool

	// Read enabled
	CanRead *bool

	// Write enabled
	CanWrite *bool
}

// TagAttributesBase - Tag attribute details
type TagAttributesBase struct {
	// REQUIRED; Writeable properties of the resource
	ChangeableAttributes *TagWriteableProperties

	// Is signed
	Signed *bool

	// READ-ONLY; Tag created time
	CreatedOn *time.Time

	// READ-ONLY; Tag digest
	Digest *string

	// READ-ONLY; Tag last update time
	LastUpdatedOn *time.Time

	// READ-ONLY; Tag name
	Name *string
}

// TagAttributesTag - Tag
type TagAttributesTag struct {
	// SignatureRecord value
	SignatureRecord *string
}

// TagList - List of tag details
type TagList struct {
	// REQUIRED; Registry login server name. This is likely to be similar to {registry-name}.azurecr.io.
	RegistryLoginServer *string

	// REQUIRED; Image name
	Repository *string

	// REQUIRED; List of tag attribute details
	TagAttributeBases []*TagAttributesBase
	Link              *string
}

// TagWriteableProperties - Changeable attributes
type TagWriteableProperties struct {
	// Delete enabled
	CanDelete *bool

	// List enabled
	CanList *bool

	// Read enabled
	CanRead *bool

	// Write enabled
	CanWrite *bool
}

// V1Manifest - Returns the requested V1 manifest file
type V1Manifest struct {
	// CPU architecture
	Architecture *string

	// List of layer information
	FsLayers []*FsLayer

	// Image history
	History []*History

	// Image name
	Name *string

	// Schema version
	SchemaVersion *int32

	// Image signature
	Signatures []*ImageSignature

	// Image tag
	Tag *string
}

// V2Manifest - Returns the requested Docker V2 Manifest file
type V2Manifest struct {
	// V2 image config descriptor
	Config *Descriptor

	// List of V2 image layer information
	Layers []*Descriptor

	// Media type for this Manifest
	MediaType *string

	// Schema version
	SchemaVersion *int32
}
