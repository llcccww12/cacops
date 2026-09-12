package memory

import (
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"

	"code.gitea.io/gitea/modules/setting"
)

func PprofMemory(ctx *context.Context) {
	t := time.Now()
	file := setting.LogRootPath + "/" + "heap" + t.Format("20060102150405")
	f, err := os.Create(file)
	if err != nil {
		log.Error("", err)
		ctx.JSON(http.StatusOK, models.BaseErrorMessageApi("fail to create pprof file."))
		return
	}
	defer f.Close()
	runtime.GC() // get up-to-date statistics
	pprof.WriteHeapProfile(f)
	ctx.JSON(http.StatusOK, models.BaseOKMessage)

}
