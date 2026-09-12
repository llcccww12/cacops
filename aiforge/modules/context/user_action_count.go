package context

import (
	"fmt"
	"sync"
	"time"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
)

var userActionMap sync.Map

var userChannel = make(chan int64, 10000)

func UserActionChannelInit() {
	go consumerUserAction(userChannel)
}

func UserActionMapClear() {
	userActionMap.Range(func(key, value any) bool {
		userActionMap.Delete(key)
		return true
	})
}

func toCache(ctx *Context) {
	if ctx.User != nil {
		_, ok := userActionMap.Load(ctx.User.ID)
		if !ok {
			log.Info("add user uid to login action db. uid=" + fmt.Sprint(ctx.User.ID))
			userActionMap.Store(ctx.User.ID, time.Now())
			userChannel <- ctx.User.ID
		}
	}
}

func consumerUserAction(in <-chan int64) {
	for uid := range in {
		models.SaveLoginActionToDb(uid)
	}
}
