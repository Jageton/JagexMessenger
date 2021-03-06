package chat

import (
	"XzibitChat/telegram"
	"XzibitChat/tmessage"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Dialog struct {
	users *UserManager
	name  string
	id    int64
}


func NewHub(id int64) *Dialog {
	return &Dialog{
		users: NewUserManager(),
		id : id,
	}
}

func (h Dialog) ID() int64 {
	return h.id
}

func (h Dialog) Name() string {
	return h.name
}

func (h *Dialog) SendToAllFrom(sender *telegram.Sender, idFrom int64, msg *tmessage.TMessage) {
	for _, id := range h.users.Ides() {
		if id == idFrom {
			continue
		}
		nmessage := tgbotapi.NewMessage(id, msg.From + ": " + msg.Text)
		sender.Send(nmessage)
	}
}

func (h *Dialog) LeaveHub(id int64) bool {
	if user, ok := h.users.Get(id); ok {
		user.RemoveHub()
		h.users.Remove(id)
		return true
	}
	return false
}

func (h *Dialog) EnterHub(user *User) {
	h.users.Add(user.ID(), user)
	user.SetHub(h)
}
