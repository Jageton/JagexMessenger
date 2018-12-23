package process

import (
	"XzibitChat/chat"
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"strings"
)

type EnterDialogCommand struct {
	DefaultUserHandler
	DefaultCommandNameGetterSetter
	DefaultAnswerParser
	DialogId  int64
	isGlobal  bool
}

func (e *EnterDialogCommand) CommandName() string {
    return "EnterDialog"
}

func (e *EnterDialogCommand) GetParsedJSON() []byte {
    m := map[string]int64{
    	"DialogId": e.DialogId,
    	"UserId" : e.user.ID(),
    }
    bytes, _ := json.Marshal(m)
    return bytes
}

func (e *EnterDialogCommand) ParseData(msg *tgbotapi.Message) bool {
	str := strings.Split(msg.CommandArguments(), " ")
	if str[0] == "" {
		return false
	}
	if len(str) != 1 {
		return false
	}
	id, err := strconv.ParseInt(str[0], 10, 64)
	if err != nil {
		return false
	}
	e.DialogId = id
	return true
}

func (e *EnterDialogCommand) Help() string {
    return "/enterHub id"
}

func (e *EnterDialogCommand) IsGlobal() bool {
    return true
}

func (e *EnterDialogCommand) PreExecuteCheck() bool {
	msg := ""
	isOk := true
	switch{
	case !e.user.IsLogon():
		msg = "You are not logged in"
		isOk = false
	case e.user.InHub():
		msg = "You are in dialog. Please leave dialog!"
		isOk = false
	}
    if !isOk {
    	e.Fail(msg)
    }
    return isOk
}

func (e *EnterDialogCommand) Execute() {
	if hub, ok := Processor.Hubs.Get(e.DialogId); ok {
		hub.EnterHub(e.user)
	}else{
		h := chat.NewHub()
		h.EnterHub(e.user)
		Processor.Hubs.Add(e.DialogId, h)
	}
}
