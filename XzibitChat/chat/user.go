package chat

import (
	"XzibitChat/telegram"
	"XzibitChat/tmessage"
)

type User struct {
	isLogon bool
	id      int64
	hub     *Dialog
	inHub   bool
}

func NewUser(id int64) *User{
	return &User{
		id: id,
	}
}

func (u *User) SetHub(hub *Dialog) {
	u.hub = hub
	u.inHub = true
}

func (u *User) RemoveHub() {
	u.hub = nil
	u.inHub = false
}

func (u *User) Hub() *Dialog {
	return u.hub
}

func (u *User) SendMessageToHub(sender *telegram.Sender, msg *tmessage.TMessage){
	u.hub.SendToAllFrom(sender, u.id, msg)
}

func (u *User) InHub() bool {
	return u.inHub
}

func (u *User) IsLogon() bool {
	return u.isLogon
}

func (u *User) ID() int64 {
	return u.id
}

func (u *User) SetLogin(v bool) {
	u.isLogon = v
}
