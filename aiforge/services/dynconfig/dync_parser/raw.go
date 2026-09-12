package dync_parser

import (
	"code.gitea.io/gitea/entity"
)

func init() {
	// Register the parser for "promote" config type
	RegisterParser(entity.DynConfigTypeRaw, &RawParser{})
}

type RawParser struct {
}

func (p *RawParser) ParseObject(config entity.DynConfig) (interface{}, error) {
	return config.Value, nil
}

func (p *RawParser) ParseList(config entity.DynConfig) (interface{}, error) {
	return config.Value, nil
}
