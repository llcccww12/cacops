package dync_parser

import (
	"fmt"

	"code.gitea.io/gitea/entity"
)

type DynConfigParser interface {
	ParseObject(entity.DynConfig) (interface{}, error)
	ParseList(entity.DynConfig) (interface{}, error)
}

func RegisterParser(configType string, parser DynConfigParser) {
	parsers[configType] = parser
}

func GetParser(configType string) (DynConfigParser, bool) {
	parser, exists := parsers[configType]
	return parser, exists
}

func ParseConfig(config entity.DynConfig) (interface{}, error) {
	parser, exists := GetParser(config.ConfigType)
	if !exists {
		return nil, fmt.Errorf("no parser registered for config type: %s", config.ConfigType)
	}
	if config.ValueType == entity.ValueTypeObject {
		return parser.ParseObject(config)
	}
	return parser.ParseList(config)
}

var parsers = make(map[string]DynConfigParser)
