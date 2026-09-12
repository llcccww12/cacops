package socketwrap

import (
	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"github.com/gorilla/websocket"
)

type Client struct {
	Manager *ClientsManager

	Conn *websocket.Conn

	Send chan *models.Action
}

func (c *Client) Close() {
	close(c.Send)
	c.Conn.Close()
}

func (c *Client) WritePump() {

	defer func() {
		c.Manager.Unregister <- c
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:

			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				log.Warn("send socket is closed")
				return
			}
			log.Warn("socket:", message)
			err := c.Conn.WriteJSON(message)
			if err != nil {
				log.Warn("can not send message", err)
				return
			}

			n := len(c.Send)
			for i := 0; i < n; i++ {
				err = c.Conn.WriteJSON(<-c.Send)
				if err != nil {
					log.Warn("can not send message", err)
					return
				}
			}

		}
	}
}
