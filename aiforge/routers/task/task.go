package task

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/services/task"
)

func RunTask() {
	for {
		select {
		case action := <-models.ActionChan4Task:
			task.Accomplish(action)
		}
	}
}
