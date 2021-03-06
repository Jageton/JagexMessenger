package process

import (
	"XzibitChat/chat"
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mitchellh/mapstructure"
	"strings"
)

type CreateDialogCommand struct {
	DefaultCommandNameGetterSetter
	DefaultUserHandler
	DialogId              int64
	DialogName     string
	UserLogins   []string
	UserIds     []int64
}

func (e CreateDialogCommand) CommandName() string {
    return "create_dialog"
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
	hub := chat.NewHub(e.DialogId)
	hub.EnterHub(e.user)
	Processor.Hubs.Add(e.DialogId, hub)
	SendMessage(e.user.ID(), "Dialogue was created!")
	for _, id := range e.UserIds {
		if e.user.ID() == id {
			continue
		}
		SendMessage(id, "You was added to dialog: " + e.DialogName)
	}
}

func (e *CreateDialogCommand) GetParsedJSON() []byte {
	m := map[string]interface{}{
		"DialogName" : e.DialogName,
		"UserLogins" : e.UserLogins,
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
    return "/create Name Users..."
}

func (e * CreateDialogCommand) ParseAnswer(args []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(args, &m)
	if err != nil {
		return err
	}
	return mapstructure.Decode(m, &e)
}
