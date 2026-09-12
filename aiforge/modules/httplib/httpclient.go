package httplib

import (
	"crypto/tls"
	"crypto/x509"
	"net"
	"net/http"
	"os"
	"time"
)

// Config HTTP客户端配置
type Config struct {
	Timeout               time.Duration
	DialTimeout           time.Duration
	TLSHandshakeTimeout   time.Duration
	ResponseHeaderTimeout time.Duration
	IdleConnTimeout       time.Duration
	MaxIdleConns          int
	MaxIdleConnsPerHost   int
	MaxConnsPerHost       int
	KeepAlive             time.Duration
	InsecureSkipVerify    bool   // 谨慎使用！
	CustomCAFile          string // 自定义CA证书文件
	MinTLSVersion         uint16
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		Timeout:               120 * time.Second,
		DialTimeout:           10 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: 15 * time.Second,
		IdleConnTimeout:       90 * time.Second,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   20,
		MaxConnsPerHost:       50,
		KeepAlive:             30 * time.Second,
		InsecureSkipVerify:    false,
		MinTLSVersion:         tls.VersionTLS12,
	}
}

// NewClientTimeOut 根据超时创建HTTP客户端
func NewClientTimeOut(cfg *Config) *http.Client {
	if cfg == nil || cfg.Timeout == 0 {
		cfg = DefaultConfig()
	}

	return &http.Client{
		Timeout: cfg.Timeout,
	}
}

// NewClientWithConfig 根据配置创建HTTP客户端
func NewClientWithConfig(cfg *Config) (*http.Client, error) {
	if cfg == nil {
		cfg = DefaultConfig()
	}

	// 创建TLS配置
	tlsConfig := &tls.Config{
		MinVersion: cfg.MinTLSVersion,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
		InsecureSkipVerify: cfg.InsecureSkipVerify,
	}

	// 加载自定义CA证书
	if cfg.CustomCAFile != "" {
		caCert, err := os.ReadFile(cfg.CustomCAFile)
		if err != nil {
			return nil, err
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		tlsConfig.RootCAs = caCertPool
	}

	transport := &http.Transport{
		MaxIdleConns:        cfg.MaxIdleConns,
		MaxIdleConnsPerHost: cfg.MaxIdleConnsPerHost,
		MaxConnsPerHost:     cfg.MaxConnsPerHost,
		IdleConnTimeout:     cfg.IdleConnTimeout,

		DialContext: (&net.Dialer{
			Timeout:   cfg.DialTimeout,
			KeepAlive: cfg.KeepAlive,
			DualStack: true,
		}).DialContext,

		TLSHandshakeTimeout:   cfg.TLSHandshakeTimeout,
		ResponseHeaderTimeout: cfg.ResponseHeaderTimeout,
		ExpectContinueTimeout: 5 * time.Second,

		TLSClientConfig:    tlsConfig,
		ForceAttemptHTTP2:  true,
		DisableCompression: false,
	}

	return &http.Client{
		Timeout:   cfg.Timeout,
		Transport: transport,
	}, nil
}
