package context

import "code.gitea.io/gitea/models"

type AccessContext struct {
	models.Permission
	*models.SubjectAccessContext
}
