package lock

import (
	"code.gitea.io/gitea/models"
)

type LockContext struct {
	Repo *models.Repository
	Task *models.Cloudbrain
	User *models.User
}

type Lock interface {
	IsMatch(ctx *LockContext) bool
	Lock(ctx *LockContext) string
	Unlock(ctx *LockContext) error
}
