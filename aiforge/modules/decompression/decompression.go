package decompression

import "code.gitea.io/gitea/modules/worker"

func NewContext() {
	worker.NewTaskCenter()
}
