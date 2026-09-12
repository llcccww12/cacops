package urchin

import (
	"code.gitea.io/gitea/modules/urfs_client/config"
	urfs "code.gitea.io/gitea/modules/urfs_client/dfstore"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
)

type Urchinfs interface {

	// schedule source dataset to target peer
	ScheduleDataToPeer(sourceUrl, destPeerHost string) (*PeerResult, error)

	// check schedule data to peer task status
	CheckScheduleTaskStatus(sourceUrl, destPeerHost string) (*PeerResult, error)

	ScheduleDataToPeerByKey(endpoint, bucketName, objectKey, destPeerHost string) (*PeerResult, error)

	CheckScheduleTaskStatusByKey(endpoint, bucketName, objectKey, destPeerHost string) (*PeerResult, error)

	ScheduleDirToPeerByKey(endpoint, bucketName, objectKey, destPeerHost string) (*PeerResult, error)

	CheckScheduleDirTaskStatusByKey(endpoint, bucketName, objectKey, destPeerHost string) (*PeerResult, error)
}

type urchinfs struct {
	// Initialize default urfs config.
	cfg *config.DfstoreConfig
}

// New urchinfs instance.
func New() Urchinfs {

	urfs := &urchinfs{
		cfg: config.NewDfstore(),
	}
	return urfs
}

const (
	// UrfsScheme if the scheme of object storage.
	UrfsScheme = "urfs"
)

func (urfs *urchinfs) ScheduleDataToPeer(sourceUrl, destPeerHost string) (*PeerResult, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := urfs.cfg.Validate(); err != nil {
		return nil, err
	}

	if err := validateSchedulelArgs(sourceUrl, destPeerHost); err != nil {
		return nil, err
	}

	// Copy object storage to local file.
	endpoint, bucketName, objectKey, err := parseUrfsURL(sourceUrl)
	if err != nil {
		return nil, err
	}
	peerResult, err := processScheduleDataToPeer(ctx, urfs.cfg, endpoint, bucketName, objectKey, destPeerHost)
	if err != nil {
		return nil, err
	}

	return peerResult, err
}

func (urfs *urchinfs) ScheduleDataToPeerByKey(endpoint, bucketName, objectKey, destPeerHost string) (*PeerResult, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	peerResult, err := processScheduleDataToPeer(ctx, urfs.cfg, endpoint, bucketName, objectKey, destPeerHost)
	if err != nil {
		return nil, err
	}

	return peerResult, err
}

func (urfs *urchinfs) ScheduleDirToPeerByKey(endpoint, bucketName, objectKey, destPeerHost string) (*PeerResult, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	peerResult, err := processScheduleDirToPeer(ctx, urfs.cfg, endpoint, bucketName, objectKey, destPeerHost)
	if err != nil {
		return nil, err
	}

	return peerResult, err
}

func (urfs *urchinfs) CheckScheduleTaskStatus(sourceUrl, destPeerHost string) (*PeerResult, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := urfs.cfg.Validate(); err != nil {
		return nil, err
	}

	if err := validateSchedulelArgs(sourceUrl, destPeerHost); err != nil {
		return nil, err
	}

	// Copy object storage to local file.
	endpoint, bucketName, objectKey, err := parseUrfsURL(sourceUrl)
	if err != nil {
		return nil, err
	}
	peerResult, err := processCheckScheduleTaskStatus(ctx, urfs.cfg, endpoint, bucketName, objectKey, destPeerHost)
	if err != nil {
		return nil, err
	}

	return peerResult, err
}

func (urfs *urchinfs) CheckScheduleTaskStatusByKey(endpoint, bucketName, objectKey, destPeerHost string) (*PeerResult, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	peerResult, err := processCheckScheduleTaskStatus(ctx, urfs.cfg, endpoint, bucketName, objectKey, destPeerHost)
	if err != nil {
		return nil, err
	}

	return peerResult, err
}

func (urfs *urchinfs) CheckScheduleDirTaskStatusByKey(endpoint, bucketName, objectKey, destPeerHost string) (*PeerResult, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	peerResult, err := processCheckScheduleDirTaskStatus(ctx, urfs.cfg, endpoint, bucketName, objectKey, destPeerHost)
	if err != nil {
		return nil, err
	}

	return peerResult, err
}

// isUrfsURL determines whether the raw url is urfs url.
func isUrfsURL(rawURL string) bool {
	u, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return false
	}

	if u.Scheme != UrfsScheme || u.Host == "" || u.Path == "" {
		return false
	}

	return true
}

// Validate copy arguments.
func validateSchedulelArgs(sourceUrl, destPeer string) error {
	if !isUrfsURL(sourceUrl) {
		return errors.New("source url should be urfs:// protocol")
	}

	return nil
}

// Parse object storage url. eg: urfs://源数据$endpoint/源数据$bucket/源数据filepath
func parseUrfsURL(rawURL string) (string, string, string, error) {
	u, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return "", "", "", err
	}

	if u.Scheme != UrfsScheme {
		return "", "", "", fmt.Errorf("invalid scheme, e.g. %s://endpoint/bucket_name/object_key", UrfsScheme)
	}

	if u.Host == "" {
		return "", "", "", errors.New("empty endpoint name")
	}

	if u.Path == "" {
		return "", "", "", errors.New("empty object path")
	}

	s := strings.Trim(u.Path, "/")
	t := strings.IndexByte(s, '/')
	if t == -1 {
		return "", "", "", errors.New("invalid bucket and object key " + u.Path)
	}
	bucket, key := s[:t], s[t+1:]

	return u.Host, bucket, key, nil
}

// Schedule object storage to peer.
func processScheduleDataToPeer(ctx context.Context, cfg *config.DfstoreConfig, endpoint, bucketName, objectKey, dstPeer string) (*PeerResult, error) {
	dfs := urfs.New(cfg.Endpoint)
	meta, err := dfs.GetUrfsMetadataWithContext(ctx, &urfs.GetUrfsMetadataInput{
		Endpoint:   endpoint,
		BucketName: bucketName,
		ObjectKey:  objectKey,
		DstPeer:    dstPeer,
	}, false)
	if err != nil {
		return nil, err
	}

	reader, err := dfs.GetUrfsWithContext(ctx, &urfs.GetUrfsInput{
		Endpoint:   endpoint,
		BucketName: bucketName,
		ObjectKey:  objectKey,
		DstPeer:    dstPeer,
	}, false)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	body, err := ioutil.ReadAll(reader)

	var peerResult PeerResult
	if err == nil {
		err = json.Unmarshal((body), &peerResult)
	}
	peerResult.SignedUrl = strings.ReplaceAll(peerResult.SignedUrl, "\\u0026", "&")

	fileContentLength, err := strconv.ParseInt(peerResult.ContentLength, 10, 64)
	if err != nil {
		return nil, err
	}
	if fileContentLength != meta.ContentLength {
		return nil, errors.New("content length inconsistent with meta")
	}

	return &peerResult, err
}

// Schedule object storage dir to peer.
func processScheduleDirToPeer(ctx context.Context, cfg *config.DfstoreConfig, endpoint, bucketName, objectKey, dstPeer string) (*PeerResult, error) {
	dfs := urfs.New(cfg.Endpoint)

	reader, err := dfs.GetUrfsWithContext(ctx, &urfs.GetUrfsInput{
		Endpoint:   endpoint,
		BucketName: bucketName,
		ObjectKey:  objectKey,
		DstPeer:    dstPeer,
	}, true)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	body, err := ioutil.ReadAll(reader)

	var peerResult PeerResult
	if err == nil {
		err = json.Unmarshal((body), &peerResult)
	}
	peerResult.SignedUrl = strings.ReplaceAll(peerResult.SignedUrl, "\\u0026", "&")

	return &peerResult, err
}

// check schedule task status.
func processCheckScheduleTaskStatus(ctx context.Context, cfg *config.DfstoreConfig, endpoint, bucketName, objectKey, dstPeer string) (*PeerResult, error) {
	dfs := urfs.New(cfg.Endpoint)
	meta, err := dfs.GetUrfsMetadataWithContext(ctx, &urfs.GetUrfsMetadataInput{
		Endpoint:   endpoint,
		BucketName: bucketName,
		ObjectKey:  objectKey,
		DstPeer:    dstPeer,
	}, false)
	if err != nil {
		return nil, err
	}

	reader, err := dfs.GetUrfsStatusWithContext(ctx, &urfs.GetUrfsInput{
		Endpoint:   endpoint,
		BucketName: bucketName,
		ObjectKey:  objectKey,
		DstPeer:    dstPeer,
	}, false)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	body, err := ioutil.ReadAll(reader)

	var peerResult PeerResult
	if err == nil {
		err = json.Unmarshal((body), &peerResult)
	}
	peerResult.SignedUrl = strings.ReplaceAll(peerResult.SignedUrl, "\\u0026", "&")

	fileContentLength, err := strconv.ParseInt(peerResult.ContentLength, 10, 64)
	if err != nil {
		return nil, err
	}
	if fileContentLength != meta.ContentLength {
		return nil, err
	}
	return &peerResult, err
}

// check schedule task status.
func processCheckScheduleDirTaskStatus(ctx context.Context, cfg *config.DfstoreConfig, endpoint, bucketName, objectKey, dstPeer string) (*PeerResult, error) {
	dfs := urfs.New(cfg.Endpoint)

	reader, err := dfs.GetUrfsStatusWithContext(ctx, &urfs.GetUrfsInput{
		Endpoint:   endpoint,
		BucketName: bucketName,
		ObjectKey:  objectKey,
		DstPeer:    dstPeer,
	}, true)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	body, err := ioutil.ReadAll(reader)

	var peerResult PeerResult
	if err == nil {
		err = json.Unmarshal((body), &peerResult)
	}
	peerResult.SignedUrl = strings.ReplaceAll(peerResult.SignedUrl, "\\u0026", "&")
	return &peerResult, err
}

type PeerResult struct {
	ContentType   string `json:"Content-Type"`
	ContentLength string `json:"Content-Length"`
	SignedUrl     string
	DataRoot      string
	DataPath      string
	DataEndpoint  string
	StatusCode    int
	StatusMsg     string
	TaskID        string
}
