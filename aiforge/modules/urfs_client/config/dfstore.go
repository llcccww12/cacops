package config

import (
	"errors"
	"fmt"
	"net/url"
)

type DfstoreConfig struct {
	// Address of the object storage service.
	Endpoint string `yaml:"endpoint,omitempty" mapstructure:"endpoint,omitempty"`

	// Filter is used to generate a unique Task ID by
	// filtering unnecessary query params in the URL,
	// it is separated by & character.
	Filter string `yaml:"filter,omitempty" mapstructure:"filter,omitempty"`

	// Mode is the mode in which the backend is written,
	// including WriteBack and AsyncWriteBack.
	Mode int `yaml:"mode,omitempty" mapstructure:"mode,omitempty"`

	// MaxReplicas is the maximum number of
	// replicas of an object cache in seed peers.
	MaxReplicas int `yaml:"maxReplicas,omitempty" mapstructure:"mode,maxReplicas"`
}

// New dfstore configuration.
func NewDfstore() *DfstoreConfig {
	url := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%d", "127.0.0.1", DefaultObjectStorageStartPort),
	}

	return &DfstoreConfig{
		Endpoint:    url.String(),
		MaxReplicas: DefaultObjectMaxReplicas,
	}
}

func (cfg *DfstoreConfig) Validate() error {
	if cfg.Endpoint == "" {
		return errors.New("dfstore requires parameter endpoint")
	}

	if _, err := url.ParseRequestURI(cfg.Endpoint); err != nil {
		return fmt.Errorf("invalid endpoint: %w", err)
	}

	return nil
}
