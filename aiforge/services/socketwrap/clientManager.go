package socketwrap

import (
	"os"
	"os/signal"
	"syscall"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"github.com/elliotchance/orderedmap"
)

var opTypes = []int{1, 2, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15, 17, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 35, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 53, 54, 55, 56, 57, 60, 62, 63, 65, 66, 67, 68, 70, 71, 72}

type ClientsManager struct {
	Clients    *orderedmap.OrderedMap
	Register   chan *Client
	Unregister chan *Client
}

func NewClientsManager() *ClientsManager {
	return &ClientsManager{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    orderedmap.NewOrderedMap(),
	}
}

const MaxClients = 100

var LastActionsQueue = NewSyncQueue(15)

func (h *ClientsManager) Run() {
	initActionQueue()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	var signalsReceived uint
	for {
		select {
		case client := <-h.Register:
			h.Clients.Set(client, true)
			if h.Clients.Len() > MaxClients {
				h.Clients.Delete(h.Clients.Front().Key)
			}

		case client := <-h.Unregister:
			if _, ok := h.Clients.Get(client); ok {
				h.Clients.Delete(client)
				close(client.Send)
			}
		case message := <-models.ActionChan:
			if isInOpTypes(opTypes, message.OpType) {
				filterUserPrivateInfo(message)
				message.FilterCloudbrainInfo()
				message.FilterDatasetInfo()
				message.FilterAimodelInfo()
				LastActionsQueue.Push(message)
				for _, client := range h.Clients.Keys() {
					select {
					case client.(*Client).Send <- message:
					default:
						close(client.(*Client).Send)
						h.Clients.Delete(client)
					}
				}
			}
		case s := <-sig:
			log.Info("received signal", s)
			signalsReceived++
			if signalsReceived < 2 {
				for _, client := range h.Clients.Keys() {
					h.Clients.Delete(client)
					client.(*Client).Close()
				}
				break

			}
		}
	}
}

func isInOpTypes(types []int, opType models.ActionType) bool {
	isFound := false
	for _, value := range types {
		if value == int(opType) {
			isFound = true
			break
		}
	}
	return isFound
}

func initActionQueue() {
	actions, err := models.GetLast20PublicFeeds(opTypes)
	if err == nil {
		for i := len(actions) - 1; i >= 0; i-- {

			user, err := models.GetUserByID(actions[i].UserID)
			if err == nil {
				if !user.IsOrganization() {
					filterUserPrivateInfo(actions[i])
					actions[i].FilterCloudbrainInfo()
					actions[i].FilterDatasetInfo()
					actions[i].FilterAimodelInfo()
					LastActionsQueue.Push(actions[i])
				}

			}

		}
	}
}

func filterUserPrivateInfo(action *models.Action) {
	action.Comment = nil
	if action.ActUser != nil {
		action.ActUser.Email = ""
		action.ActUser.Passwd = ""
		action.ActUser.PasswdHashAlgo = ""
		action.ActUser.PrivateKey = ""
		action.ActUser.PublicKey = ""
		action.ActUser.Salt = ""
		action.ActUser.FullName = ""
		action.ActUser.AvatarEmail = ""
		action.ActUser.IsAdmin = false
		action.ActUser.EmailNotificationsPreference = ""
		action.ActUser.IsOperator = false
		action.ActUser.PhoneNumber = ""
		action.ActUser.LastLoginUnix = 0
		action.ActUser.Location = ""
		action.ActUser.LoginSource = 0
		action.ActUser.Rands = ""
		action.ActUser.WechatOpenId = ""
		action.ActUser.WechatBindUnix = 0

	}
	if action.Repo != nil {
		action.Repo.Owner = nil
	}

}

func GetLastActions() []*models.Action {
	list := LastActionsQueue.List()
	var actions []*models.Action
	for _, value := range list {
		actions = append(actions, value.(*models.Action))
	}
	return actions

}
