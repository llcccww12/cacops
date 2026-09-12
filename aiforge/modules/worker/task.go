package worker

import (
	"context"

	"code.gitea.io/gitea/modules/log"

	"github.com/RichardKnop/machinery/v1/tasks"
)

// 方法名
const (
	DecompressTaskName = "Decompress"
)

func SendDecompressTask(ctx context.Context, uuid string, name string) error {
	args := []tasks.Arg{{Name: "uuid", Type: "string", Value: uuid}, {Name: "name", Type: "string", Value: name}}
	task, err := tasks.NewSignature(DecompressTaskName, args)
	if err != nil {
		log.Error("NewSignature failed:", err.Error())
		return err
	}

	task.RetryCount = 3
	_, err = AsyncTaskCenter.SendTaskWithContext(ctx, task)
	if err != nil {
		log.Error("SendTaskWithContext failed:", err.Error())
		return err
	}

	log.Info("SendDecompressTask(%s) success", uuid)

	return nil
}
