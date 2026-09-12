package pipeline

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

var pathProxys = map[string]*httputil.ReverseProxy{}
var patternProxys = map[string]*httputil.ReverseProxy{}

func GetOrCreatePathProxy(scheme, host, path string) *httputil.ReverseProxy {
	key := fmt.Sprintf("%s-%s-%s", scheme, host, path)
	proxy, ok := pathProxys[key]
	if ok {
		return proxy
	}

	director := func(req *http.Request) {
		req.URL.Scheme = scheme
		req.URL.Host = host
		req.URL.Path = path
	}
	proxy = &httputil.ReverseProxy{Director: director}
	pathProxys[key] = proxy
	return proxy
}
