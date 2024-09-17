// Copyright 2024 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package cloudprofiler provides access to the Cloud Profiler API.
//
// For product documentation, see: https://cloud.google.com/profiler/
//
// # Library status
//
// These client libraries are officially supported by Google. However, this
// library is considered complete and is in maintenance mode. This means
// that we will address critical bugs and security issues but will not add
// any new features.
//
// When possible, we recommend using our newer
// [Cloud Client Libraries for Go](https://pkg.go.dev/cloud.google.com/go)
// that are still actively being worked and iterated on.
//
// # Creating a client
//
// Usage example:
//
//	import "google.golang.org/api/cloudprofiler/v2"
//	...
//	ctx := context.Background()
//	cloudprofilerService, err := cloudprofiler.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for
// authentication. For information on how to create and obtain Application
// Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// # Other authentication options
//
// By default, all available scopes (see "Constants") are used to authenticate.
// To restrict scopes, use [google.golang.org/api/option.WithScopes]:
//
//	cloudprofilerService, err := cloudprofiler.NewService(ctx, option.WithScopes(cloudprofiler.MonitoringWriteScope))
//
// To use an API key for authentication (note: some APIs do not support API
// keys), use [google.golang.org/api/option.WithAPIKey]:
//
//	cloudprofilerService, err := cloudprofiler.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth
// flow, use [google.golang.org/api/option.WithTokenSource]:
//
//	config := &oauth2.Config{...}
//	// ...
//	token, err := config.Exchange(ctx, ...)
//	cloudprofilerService, err := cloudprofiler.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See [google.golang.org/api/option.ClientOption] for details on options.
package cloudprofiler // import "google.golang.org/api/cloudprofiler/v2"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	googleapi "google.golang.org/api/googleapi"
	internal "google.golang.org/api/internal"
	gensupport "google.golang.org/api/internal/gensupport"
	option "google.golang.org/api/option"
	internaloption "google.golang.org/api/option/internaloption"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = internaloption.WithDefaultEndpoint
var _ = internal.Version

const apiId = "cloudprofiler:v2"
const apiName = "cloudprofiler"
const apiVersion = "v2"
const basePath = "https://cloudprofiler.googleapis.com/"
const basePathTemplate = "https://cloudprofiler.UNIVERSE_DOMAIN/"
const mtlsBasePath = "https://cloudprofiler.mtls.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// See, edit, configure, and delete your Google Cloud data and see the email
	// address for your Google Account.
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View and write monitoring data for all of your Google and third-party Cloud
	// and API projects
	MonitoringScope = "https://www.googleapis.com/auth/monitoring"

	// Publish metric data to your Google Cloud projects
	MonitoringWriteScope = "https://www.googleapis.com/auth/monitoring.write"
)

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	scopesOption := internaloption.WithDefaultScopes(
		"https://www.googleapis.com/auth/cloud-platform",
		"https://www.googleapis.com/auth/monitoring",
		"https://www.googleapis.com/auth/monitoring.write",
	)
	// NOTE: prepend, so we don't override user-specified scopes.
	opts = append([]option.ClientOption{scopesOption}, opts...)
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
	opts = append(opts, internaloption.WithDefaultEndpointTemplate(basePathTemplate))
	opts = append(opts, internaloption.WithDefaultMTLSEndpoint(mtlsBasePath))
	opts = append(opts, internaloption.EnableNewAuthLibrary())
	client, endpoint, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	s, err := New(client)
	if err != nil {
		return nil, err
	}
	if endpoint != "" {
		s.BasePath = endpoint
	}
	return s, nil
}

// New creates a new Service. It uses the provided http.Client for requests.
//
// Deprecated: please use NewService instead.
// To provide a custom HTTP client, use option.WithHTTPClient.
// If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Profiles = NewProjectsProfilesService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Profiles *ProjectsProfilesService
}

func NewProjectsProfilesService(s *Service) *ProjectsProfilesService {
	rs := &ProjectsProfilesService{s: s}
	return rs
}

type ProjectsProfilesService struct {
	s *Service
}

// CreateProfileRequest: CreateProfileRequest describes a profile resource
// online creation request. The deployment field must be populated. The
// profile_type specifies the list of profile types supported by the agent. The
// creation call will hang until a profile of one of these types needs to be
// collected.
type CreateProfileRequest struct {
	// Deployment: Deployment details.
	Deployment *Deployment `json:"deployment,omitempty"`
	// ProfileType: One or more profile types that the agent is capable of
	// providing.
	//
	// Possible values:
	//   "PROFILE_TYPE_UNSPECIFIED" - Unspecified profile type.
	//   "CPU" - Thread CPU time sampling.
	//   "WALL" - Wallclock time sampling. More expensive as stops all threads.
	//   "HEAP" - In-use heap profile. Represents a snapshot of the allocations
	// that are live at the time of the profiling.
	//   "THREADS" - Single-shot collection of all thread stacks.
	//   "CONTENTION" - Synchronization contention profile.
	//   "PEAK_HEAP" - Peak heap profile.
	//   "HEAP_ALLOC" - Heap allocation profile. It represents the aggregation of
	// all allocations made over the duration of the profile. All allocations are
	// included, including those that might have been freed by the end of the
	// profiling interval. The profile is in particular useful for garbage
	// collecting languages to understand which parts of the code create most of
	// the garbage collection pressure to see if those can be optimized.
	ProfileType []string `json:"profileType,omitempty"`
	// ForceSendFields is a list of field names (e.g. "Deployment") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "Deployment") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s *CreateProfileRequest) MarshalJSON() ([]byte, error) {
	type NoMethod CreateProfileRequest
	return gensupport.MarshalJSON(NoMethod(*s), s.ForceSendFields, s.NullFields)
}

// Deployment: Deployment contains the deployment identification information.
type Deployment struct {
	// Labels: Labels identify the deployment within the user universe and same
	// target. Validation regex for label names: `^a-z0-9
	// ([a-z0-9-]{0,61}[a-z0-9])?$`. Value for an individual label must be <= 512
	// bytes, the total size of all label names and values must be <= 1024 bytes.
	// Label named "language" can be used to record the programming language of the
	// profiled deployment. The standard choices for the value include "java",
	// "go", "python", "ruby", "nodejs", "php", "dotnet". For deployments running
	// on Google Cloud Platform, "zone" or "region" label should be present
	// describing the deployment location. An example of a zone is "us-central1-a",
	// an example of a region is "us-central1" or "us-central".
	Labels map[string]string `json:"labels,omitempty"`
	// ProjectId: Project ID is the ID of a cloud project. Validation regex:
	// `^a-z{4,61}[a-z0-9]$`.
	ProjectId string `json:"projectId,omitempty"`
	// Target: Target is the service name used to group related deployments: *
	// Service name for App Engine Flex / Standard. * Cluster and container name
	// for GKE. * User-specified string for direct Compute Engine profiling (e.g.
	// Java). * Job name for Dataflow. Validation regex: `^a-z0-9
	// ([-a-z0-9_.]{0,253}[a-z0-9])?$`.
	Target string `json:"target,omitempty"`
	// ForceSendFields is a list of field names (e.g. "Labels") to unconditionally
	// include in API requests. By default, fields with empty or default values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "Labels") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s *Deployment) MarshalJSON() ([]byte, error) {
	type NoMethod Deployment
	return gensupport.MarshalJSON(NoMethod(*s), s.ForceSendFields, s.NullFields)
}

// ListProfilesResponse: ListProfileResponse contains the list of collected
// profiles for deployments in projects which the user has permissions to view.
type ListProfilesResponse struct {
	// NextPageToken: Token to receive the next page of results. This field maybe
	// empty if there are no more profiles to fetch.
	NextPageToken string `json:"nextPageToken,omitempty"`
	// Profiles: List of profiles fetched.
	Profiles []*Profile `json:"profiles,omitempty"`
	// SkippedProfiles: Number of profiles that were skipped in the current page
	// since they were not able to be fetched successfully. This should typically
	// be zero. A non-zero value may indicate a transient failure, in which case if
	// the number is too high for your use case, the call may be retried.
	SkippedProfiles int64 `json:"skippedProfiles,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the server.
	googleapi.ServerResponse `json:"-"`
	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "NextPageToken") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s *ListProfilesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListProfilesResponse
	return gensupport.MarshalJSON(NoMethod(*s), s.ForceSendFields, s.NullFields)
}

// Profile: Profile resource.
type Profile struct {
	// Deployment: Deployment this profile corresponds to.
	Deployment *Deployment `json:"deployment,omitempty"`
	// Duration: Duration of the profiling session. Input (for the offline mode) or
	// output (for the online mode). The field represents requested profiling
	// duration. It may slightly differ from the effective profiling duration,
	// which is recorded in the profile data, in case the profiling can't be
	// stopped immediately (e.g. in case stopping the profiling is handled
	// asynchronously).
	Duration string `json:"duration,omitempty"`
	// Labels: Input only. Labels associated to this specific profile. These labels
	// will get merged with the deployment labels for the final data set. See
	// documentation on deployment labels for validation rules and limits.
	Labels map[string]string `json:"labels,omitempty"`
	// Name: Output only. Opaque, server-assigned, unique ID for this profile.
	Name string `json:"name,omitempty"`
	// ProfileBytes: Input only. Profile bytes, as a gzip compressed serialized
	// proto, the format is
	// https://github.com/google/pprof/blob/master/proto/profile.proto.
	ProfileBytes string `json:"profileBytes,omitempty"`
	// ProfileType: Type of profile. For offline mode, this must be specified when
	// creating the profile. For online mode it is assigned and returned by the
	// server.
	//
	// Possible values:
	//   "PROFILE_TYPE_UNSPECIFIED" - Unspecified profile type.
	//   "CPU" - Thread CPU time sampling.
	//   "WALL" - Wallclock time sampling. More expensive as stops all threads.
	//   "HEAP" - In-use heap profile. Represents a snapshot of the allocations
	// that are live at the time of the profiling.
	//   "THREADS" - Single-shot collection of all thread stacks.
	//   "CONTENTION" - Synchronization contention profile.
	//   "PEAK_HEAP" - Peak heap profile.
	//   "HEAP_ALLOC" - Heap allocation profile. It represents the aggregation of
	// all allocations made over the duration of the profile. All allocations are
	// included, including those that might have been freed by the end of the
	// profiling interval. The profile is in particular useful for garbage
	// collecting languages to understand which parts of the code create most of
	// the garbage collection pressure to see if those can be optimized.
	ProfileType string `json:"profileType,omitempty"`
	// StartTime: Output only. Start time for the profile. This output is only
	// present in response from the ListProfiles method.
	StartTime string `json:"startTime,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the server.
	googleapi.ServerResponse `json:"-"`
	// ForceSendFields is a list of field names (e.g. "Deployment") to
	// unconditionally include in API requests. By default, fields with empty or
	// default values are omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-ForceSendFields for more
	// details.
	ForceSendFields []string `json:"-"`
	// NullFields is a list of field names (e.g. "Deployment") to include in API
	// requests with the JSON null value. By default, fields with empty values are
	// omitted from API requests. See
	// https://pkg.go.dev/google.golang.org/api#hdr-NullFields for more details.
	NullFields []string `json:"-"`
}

func (s *Profile) MarshalJSON() ([]byte, error) {
	type NoMethod Profile
	return gensupport.MarshalJSON(NoMethod(*s), s.ForceSendFields, s.NullFields)
}

type ProjectsProfilesCreateCall struct {
	s                    *Service
	parent               string
	createprofilerequest *CreateProfileRequest
	urlParams_           gensupport.URLParams
	ctx_                 context.Context
	header_              http.Header
}

// Create: CreateProfile creates a new profile resource in the online mode.
// _Direct use of this API is discouraged, please use a supported profiler
// agent
// (https://cloud.google.com/profiler/docs/about-profiler#profiling_agent)
// instead for profile collection._ The server ensures that the new profiles
// are created at a constant rate per deployment, so the creation request may
// hang for some time until the next profile session is available. The request
// may fail with ABORTED error if the creation is not available within ~1m, the
// response will indicate the duration of the backoff the client should take
// before attempting creating a profile again. The backoff duration is returned
// in google.rpc.RetryInfo extension on the response status. To a gRPC client,
// the extension will be return as a binary-serialized proto in the trailing
// metadata item named "google.rpc.retryinfo-bin".
//
// - parent: Parent project to create the profile in.
func (r *ProjectsProfilesService) Create(parent string, createprofilerequest *CreateProfileRequest) *ProjectsProfilesCreateCall {
	c := &ProjectsProfilesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.createprofilerequest = createprofilerequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *ProjectsProfilesCreateCall) Fields(s ...googleapi.Field) *ProjectsProfilesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *ProjectsProfilesCreateCall) Context(ctx context.Context) *ProjectsProfilesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *ProjectsProfilesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsProfilesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "application/json", c.header_)
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.createprofilerequest)
	if err != nil {
		return nil, err
	}
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+parent}/profiles")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudprofiler.projects.profiles.create" call.
// Any non-2xx status code is an error. Response headers are in either
// *Profile.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was returned.
func (c *ProjectsProfilesCreateCall) Do(opts ...googleapi.CallOption) (*Profile, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &Profile{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

type ProjectsProfilesCreateOfflineCall struct {
	s          *Service
	parent     string
	profile    *Profile
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// CreateOffline: CreateOfflineProfile creates a new profile resource in the
// offline mode. The client provides the profile to create along with the
// profile bytes, the server records it. _Direct use of this API is
// discouraged, please use a supported profiler agent
// (https://cloud.google.com/profiler/docs/about-profiler#profiling_agent)
// instead for profile collection._
//
// - parent: Parent project to create the profile in.
func (r *ProjectsProfilesService) CreateOffline(parent string, profile *Profile) *ProjectsProfilesCreateOfflineCall {
	c := &ProjectsProfilesCreateOfflineCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.profile = profile
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *ProjectsProfilesCreateOfflineCall) Fields(s ...googleapi.Field) *ProjectsProfilesCreateOfflineCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *ProjectsProfilesCreateOfflineCall) Context(ctx context.Context) *ProjectsProfilesCreateOfflineCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *ProjectsProfilesCreateOfflineCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsProfilesCreateOfflineCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "application/json", c.header_)
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.profile)
	if err != nil {
		return nil, err
	}
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+parent}/profiles:createOffline")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudprofiler.projects.profiles.createOffline" call.
// Any non-2xx status code is an error. Response headers are in either
// *Profile.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was returned.
func (c *ProjectsProfilesCreateOfflineCall) Do(opts ...googleapi.CallOption) (*Profile, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &Profile{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

type ProjectsProfilesListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists profiles which have been collected so far and for which the
// caller has permission to view.
//
//   - parent: The parent, which owns this collection of profiles. Format:
//     projects/{user_project_id}.
func (r *ProjectsProfilesService) List(parent string) *ProjectsProfilesListCall {
	c := &ProjectsProfilesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number of items
// to return. Default page_size is 1000. Max limit is 1000.
func (c *ProjectsProfilesListCall) PageSize(pageSize int64) *ProjectsProfilesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The token to continue
// pagination and get profiles from a particular page. When paginating, all
// other parameters provided to `ListProfiles` must match the call that
// provided the page token.
func (c *ProjectsProfilesListCall) PageToken(pageToken string) *ProjectsProfilesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *ProjectsProfilesListCall) Fields(s ...googleapi.Field) *ProjectsProfilesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets an optional parameter which makes the operation fail if the
// object's ETag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *ProjectsProfilesListCall) IfNoneMatch(entityTag string) *ProjectsProfilesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *ProjectsProfilesListCall) Context(ctx context.Context) *ProjectsProfilesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *ProjectsProfilesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsProfilesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "", c.header_)
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+parent}/profiles")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudprofiler.projects.profiles.list" call.
// Any non-2xx status code is an error. Response headers are in either
// *ListProfilesResponse.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsProfilesListCall) Do(opts ...googleapi.CallOption) (*ListProfilesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &ListProfilesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsProfilesListCall) Pages(ctx context.Context, f func(*ListProfilesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken"))
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

type ProjectsProfilesPatchCall struct {
	s          *Service
	name       string
	profile    *Profile
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Patch: UpdateProfile updates the profile bytes and labels on the profile
// resource created in the online mode. Updating the bytes for profiles created
// in the offline mode is currently not supported: the profile content must be
// provided at the time of the profile creation. _Direct use of this API is
// discouraged, please use a supported profiler agent
// (https://cloud.google.com/profiler/docs/about-profiler#profiling_agent)
// instead for profile collection._
//
// - name: Output only. Opaque, server-assigned, unique ID for this profile.
func (r *ProjectsProfilesService) Patch(name string, profile *Profile) *ProjectsProfilesPatchCall {
	c := &ProjectsProfilesPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.profile = profile
	return c
}

// UpdateMask sets the optional parameter "updateMask": Field mask used to
// specify the fields to be overwritten. Currently only profile_bytes and
// labels fields are supported by UpdateProfile, so only those fields can be
// specified in the mask. When no mask is provided, all fields are overwritten.
func (c *ProjectsProfilesPatchCall) UpdateMask(updateMask string) *ProjectsProfilesPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more
// details.
func (c *ProjectsProfilesPatchCall) Fields(s ...googleapi.Field) *ProjectsProfilesPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method.
func (c *ProjectsProfilesPatchCall) Context(ctx context.Context) *ProjectsProfilesPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns a http.Header that can be modified by the caller to add
// headers to the request.
func (c *ProjectsProfilesPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsProfilesPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := gensupport.SetHeaders(c.s.userAgent(), "application/json", c.header_)
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.profile)
	if err != nil {
		return nil, err
	}
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("PATCH", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudprofiler.projects.profiles.patch" call.
// Any non-2xx status code is an error. Response headers are in either
// *Profile.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was returned.
func (c *ProjectsProfilesPatchCall) Do(opts ...googleapi.CallOption) (*Profile, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, gensupport.WrapError(&googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		})
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, gensupport.WrapError(err)
	}
	ret := &Profile{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}
