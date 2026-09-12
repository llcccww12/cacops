package dfstore

import (
	"code.gitea.io/gitea/modules/urfs_client/config"
	pkgobjectstorage "code.gitea.io/gitea/modules/urfs_client/objectstorage"
	"context"
	"errors"
	"fmt"
	"github.com/go-http-utils/headers"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

// Dfstore is the interface used for object storage.
type Dfstore interface {

	// GetUrfsMetadataRequestWithContext returns *http.Request of getting Urfs metadata.
	GetUrfsMetadataRequestWithContext(ctx context.Context, input *GetUrfsMetadataInput, isDir bool) (*http.Request, error)

	// GetUrfsMetadataWithContext returns matedata of Urfs.
	GetUrfsMetadataWithContext(ctx context.Context, input *GetUrfsMetadataInput, isDir bool) (*pkgobjectstorage.ObjectMetadata, error)

	// GetUrfsRequestWithContext returns *http.Request of getting Urfs.
	GetUrfsRequestWithContext(ctx context.Context, input *GetUrfsInput, isDir bool) (*http.Request, error)

	// GetUrfsWithContext returns data of Urfs.
	GetUrfsWithContext(ctx context.Context, input *GetUrfsInput, isDir bool) (io.ReadCloser, error)

	// GetUrfsStatusRequestWithContext returns *http.Request of getting Urfs status.
	GetUrfsStatusRequestWithContext(ctx context.Context, input *GetUrfsInput, isDir bool) (*http.Request, error)

	// GetUrfsStatusWithContext returns schedule status of Urfs.
	GetUrfsStatusWithContext(ctx context.Context, input *GetUrfsInput, isDir bool) (io.ReadCloser, error)
}

// dfstore provides object storage function.
type dfstore struct {
	endpoint   string
	httpClient *http.Client
}

// Option is a functional option for configuring the dfstore.
type Option func(dfs *dfstore)

// New dfstore instance.
func New(endpoint string, options ...Option) Dfstore {
	dfs := &dfstore{
		endpoint:   endpoint,
		httpClient: http.DefaultClient,
	}

	for _, opt := range options {
		opt(dfs)
	}

	return dfs
}

// GetUrfsMetadataInput is used to construct request of getting object metadata.
type GetUrfsMetadataInput struct {

	// Endpoint is endpoint name.
	Endpoint string

	// BucketName is bucket name.
	BucketName string

	// ObjectKey is object key.
	ObjectKey string

	// DstPeer is target peerHost.
	DstPeer string
}

// Validate validates GetUrfsMetadataInput fields.
func (i *GetUrfsMetadataInput) Validate() error {

	if i.Endpoint == "" {
		return errors.New("invalid Endpoint")

	}

	if i.BucketName == "" {
		return errors.New("invalid BucketName")

	}

	if i.ObjectKey == "" {
		return errors.New("invalid ObjectKey")
	}

	return nil
}

// GetObjectMetadataRequestWithContext returns *http.Request of getting object metadata.
func (dfs *dfstore) GetUrfsMetadataRequestWithContext(ctx context.Context, input *GetUrfsMetadataInput, isDir bool) (*http.Request, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	dstUrl := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s", input.DstPeer),
	}

	u, err := url.Parse(dstUrl.String())
	if err != nil {
		return nil, err
	}

	u.Path = path.Join("buckets", input.BucketName+"."+input.Endpoint, "objects", input.ObjectKey)
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, u.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// GetObjectMetadataWithContext returns metadata of object.
func (dfs *dfstore) GetUrfsMetadataWithContext(ctx context.Context, input *GetUrfsMetadataInput, isDir bool) (*pkgobjectstorage.ObjectMetadata, error) {
	req, err := dfs.GetUrfsMetadataRequestWithContext(ctx, input, isDir)
	if err != nil {
		return nil, err
	}

	resp, err := dfs.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("bad response status %s", resp.Status)
	}

	contentLength, err := strconv.ParseInt(resp.Header.Get(headers.ContentLength), 10, 64)
	if err != nil {
		return nil, err
	}

	return &pkgobjectstorage.ObjectMetadata{
		ContentDisposition: resp.Header.Get(headers.ContentDisposition),
		ContentEncoding:    resp.Header.Get(headers.ContentEncoding),
		ContentLanguage:    resp.Header.Get(headers.ContentLanguage),
		ContentLength:      int64(contentLength),
		ContentType:        resp.Header.Get(headers.ContentType),
		ETag:               resp.Header.Get(headers.ContentType),
		Digest:             resp.Header.Get(config.HeaderDragonflyObjectMetaDigest),
	}, nil
}

// GetUrfsInput is used to construct request of getting object.
type GetUrfsInput struct {

	// Endpoint is endpoint name.
	Endpoint string

	// BucketName is bucket name.
	BucketName string

	// ObjectKey is object key.
	ObjectKey string

	// Filter is used to generate a unique Task ID by
	// filtering unnecessary query params in the URL,
	// it is separated by & character.
	Filter string

	// Range is the HTTP range header.
	Range string

	// DstPeer is target peerHost.
	DstPeer string
}

// GetObjectWithContext returns data of object.
func (dfs *dfstore) GetUrfsWithContext(ctx context.Context, input *GetUrfsInput, isDir bool) (io.ReadCloser, error) {
	req, err := dfs.GetUrfsRequestWithContext(ctx, input, isDir)
	if err != nil {
		return nil, err
	}

	resp, err := dfs.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("bad response status %s", resp.Status)
	}

	return resp.Body, nil
}

// GetObjectRequestWithContext returns *http.Request of getting object.
func (dfs *dfstore) GetUrfsRequestWithContext(ctx context.Context, input *GetUrfsInput, isDir bool) (*http.Request, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	dstUrl := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s", input.DstPeer),
	}

	u, err := url.Parse(dstUrl.String())
	if err != nil {
		return nil, err
	}

	if isDir {
		u.Path = path.Join("buckets", input.BucketName+"."+input.Endpoint, "cache_folder", input.ObjectKey)
	} else {
		u.Path = path.Join("buckets", input.BucketName+"."+input.Endpoint, "cache_object", input.ObjectKey)
	}

	query := u.Query()
	if input.Filter != "" {
		query.Set("filter", input.Filter)
	}
	u.RawQuery = query.Encode()
	println("u.string", u.String())
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), nil)
	if err != nil {
		return nil, err
	}

	if input.Range != "" {
		req.Header.Set(headers.Range, input.Range)
	}

	return req, nil
}

// Validate validates GetUrfsInput fields.
func (i *GetUrfsInput) Validate() error {

	if i.Endpoint == "" {
		return errors.New("invalid Endpoint")

	}

	if i.BucketName == "" {
		return errors.New("invalid BucketName")

	}

	if i.ObjectKey == "" {
		return errors.New("invalid ObjectKey")
	}

	return nil
}

// GetUrfsStatusWithContext returns schedule task status.
func (dfs *dfstore) GetUrfsStatusWithContext(ctx context.Context, input *GetUrfsInput, isDir bool) (io.ReadCloser, error) {
	req, err := dfs.GetUrfsStatusRequestWithContext(ctx, input, isDir)
	if err != nil {
		return nil, err
	}

	resp, err := dfs.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("bad response status %s", resp.Status)
	}

	return resp.Body, nil
}

// GetObjectStatusRequestWithContext returns *http.Request of check schedule task status.
func (dfs *dfstore) GetUrfsStatusRequestWithContext(ctx context.Context, input *GetUrfsInput, isDir bool) (*http.Request, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	dstUrl := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s", input.DstPeer),
	}

	u, err := url.Parse(dstUrl.String())
	if err != nil {
		return nil, err
	}

	if isDir {
		u.Path = path.Join("buckets", input.BucketName+"."+input.Endpoint, "check_folder", input.ObjectKey)
	} else {
		u.Path = path.Join("buckets", input.BucketName+"."+input.Endpoint, "check_object", input.ObjectKey)
	}

	query := u.Query()
	if input.Filter != "" {
		query.Set("filter", input.Filter)
	}
	u.RawQuery = query.Encode()
	//println("u.string ", u.String())
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	if input.Range != "" {
		req.Header.Set(headers.Range, input.Range)
	}

	return req, nil
}
