package common

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"code.gitea.io/gitea/modules/log"
	"github.com/hashicorp/go-retryablehttp"
)

func Post(
	ctx context.Context,
	url string,
	reqBody []byte,
	client *retryablehttp.Client) (err error, respBody []byte) {

	respBuf, err := client.Post(
		url,
		"application/json; charset=UTF-8",
		strings.NewReader(string(reqBody)))
	if nil != err {
		return err, respBody
	}
	defer func(body io.ReadCloser) {
		err := body.Close()
		if nil != err {
		}
	}(respBuf.Body)

	respBody, err = io.ReadAll(respBuf.Body)
	if nil != err {
		return err, respBody
	}

	return nil, respBody
}

func Get(
	ctx context.Context,
	url string,
	client *retryablehttp.Client) (err error, respBody []byte) {

	respBuf, err := client.Get(url)
	if nil != err {
		return err, respBody
	}
	defer func(body io.ReadCloser) {
		err := body.Close()
		if nil != err {
		}
	}(respBuf.Body)

	respBody, err = io.ReadAll(respBuf.Body)
	if nil != err {
		return err, respBody
	}

	return nil, respBody
}

func Do(
	ctx context.Context,
	url,
	method string,
	header http.Header,
	reqBody interface{},
	client *retryablehttp.Client) (err error, respBody []byte) {

	request, err := retryablehttp.NewRequest(
		method,
		url,
		reqBody)
	if nil != err {
		log.Error("urchin_v2 request(%s %s) NewRequest error. err=%v", method, url, err)
		return err, respBody
	}

	request.Header = header
	log.Info("urchin_v2 send request(%s %s) request body=%s", method, url, jsonString(reqBody))

	resp, err := client.Do(request)
	if nil != err {
		log.Error("urchin_v2 request(%s %s) response error. err=%v", method, url, err)
		return err, respBody
	}
	defer func(body io.ReadCloser) {
		_err := body.Close()
		if nil != _err {
		}
	}(resp.Body)

	respBody, err = io.ReadAll(resp.Body)
	if nil != err {
		log.Error("urchin_v2 request(%s %s) io ReadAll error. err=%v", method, url, err)
		return err, respBody
	}
	log.Info("urchin_v2 request(%s %s) response body=%s", method, url, jsonString(respBody))

	return nil, respBody
}

func jsonString(v interface{}) string {
	var bodyStr string
	if b, ok := v.([]byte); ok {
		bodyStr = string(b)
	} else if s, ok := v.(string); ok {
		bodyStr = s
	} else if v != nil {
		tmp, _ := json.Marshal(v)
		bodyStr = string(tmp)
	}
	return bodyStr
}
