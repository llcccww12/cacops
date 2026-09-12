package fetcher

import "code.gitea.io/gitea/entity"

type Fetcher interface {
	Fetch(path string) (*entity.DynConfig, error)
}
