package routers

import (
	"net/http"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/services/socketwrap"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var SocketManager = socketwrap.NewClientsManager()

func ActionNotification(ctx *context.Context) {

	conn, err := upgrader.Upgrade(ctx.Resp, ctx.Req.Request, nil)
	if err != nil {
		log.Warn("can not create connection.", err)
		return
	}
	client := &socketwrap.Client{Manager: SocketManager, Conn: conn, Send: make(chan *models.Action, 256)}

	WriteLastActionsIfHave(conn)

	client.Manager.Register <- client

	go client.WritePump()

}

func WriteLastActionsIfHave(conn *websocket.Conn) {
	socketwrap.LastActionsQueue.Mutex.RLock()
	{
		size := socketwrap.LastActionsQueue.Queue.Len()
		if size > 0 {
			tempE := socketwrap.LastActionsQueue.Queue.Front()
			conn.WriteJSON(tempE.Value)
			for i := 1; i < size; i++ {
				tempE = tempE.Next()
				conn.WriteJSON(tempE.Value)

			}
		}

	}
	socketwrap.LastActionsQueue.Mutex.RUnlock()
}

func LatestActions(ctx *context.Context) {
	if !setting.EnableWebSocket {
		actions := socketwrap.GetLastActions()
		ctx.JSON(200, map[string]interface{}{"socket": false, "actions": actions})
	} else {
		ctx.JSON(200, map[string]interface{}{"socket": true})

	}

}
