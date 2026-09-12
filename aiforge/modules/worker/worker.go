package worker

import (
	"code.gitea.io/gitea/modules/setting"
	"github.com/RichardKnop/machinery/v1"
	mchConf "github.com/RichardKnop/machinery/v1/config"
)

var (
	AsyncTaskCenter *machinery.Server
)

func NewTaskCenter() {
	cnf := &mchConf.Config{
		Broker:        setting.Broker,
		DefaultQueue:  setting.DefaultQueue,
		ResultBackend: setting.ResultBackend,
	}
	tc, err := machinery.NewServer(cnf)
	if err != nil {
		panic(err)
	}
	AsyncTaskCenter = tc
}
