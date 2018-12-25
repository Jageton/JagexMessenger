package process

import (
	"XzibitChat/chat"
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

type CreateDialogCommand struct {
	DefaultCommandNameGetterSetter
	DefaultUserHandler
	DefaultAnswerParser
	DialogId              int64
	DialogName     string
	UserLogins   []string
	UsersIds     []int64
}

func (e CreateDialogCommand) CommandName() string {
    return "CreateDialog"
}

func (CreateDialogCommand) IsGlobal() bool {
    return true
}

func (e *CreateDialogCommand) PreExecuteCheck() bool {
    msg := ""
    isOk := true
    switch {
    case e.user.InHub():
    	msg = "Please leave dialogue!"
    	isOk = false
    case !e.user.IsLogon():
        msg = "You are not logged in!"
        isOk = false
    }
    if !isOk {
    	SendMessage(e.user.ID(), msg)
    }
    return isOk
}

func (e CreateDialogCommand) Execute() {
	hub := chat.NewHub()
	hub.EnterHub(e.user)
	Processor.Hubs.Add(e.DialogId, hub)
	SendMessage(e.user.ID(), "Dialogue was created!")
	for _, id := range e.UsersIds {
		SendMessage(id, "You was added to dialog: " + e.DialogName)
	}
}

func (e *CreateDialogCommand) GetParsedJSON() []byte {
	m := map[string]interface{}{
		"DialogueName" : e.DialogName,
		"Users" : e.UserLogins,
		"FromUserId" : e.user.ID(),
	}
	bytes, _ := json.Marshal(m)
	return bytes
}

func (e *CreateDialogCommand) ParseData(msg *tgbotapi.Message) bool {
	str := strings.Split(msg.CommandArguments(), " ")
	if str[0] == "" {
		return false
	}
	if len(str) < 1 {
		return false
	}
	e.DialogName = str[0]
	e.UserLogins = str[1:]
	return true
}

func ( CreateDialogCommand) Help() string {
    return "/CreateHub Name Users..."
}

