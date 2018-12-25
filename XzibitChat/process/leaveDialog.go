package process

import (
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"strings"
)

type LeaveDialogCommand struct {
	DefaultCommandNameGetterSetter
	DefaultUserHandler
	DefaultAnswerParser
	DialogId int64
}

func (e LeaveDialogCommand) CommandName() string {
    return "LeaveDialog"
}

func ( LeaveDialogCommand) IsGlobal() bool {
    return true
}

func (e LeaveDialogCommand) PreExecuteCheck() bool {
    msg := ""
    isOk := true
	switch  {
	case !e.user.IsLogon():
		msg = "You are not logged in!"
		isOk = false
	}
    if !isOk {
        e.Fail(msg)
    }
    return isOk
}

func (e LeaveDialogCommand) Execute() {
	SendMessage(e.user.ID(), "You left from the dialog!")
}

func (e LeaveDialogCommand) GetParsedJSON() []byte {
    m := map[string]int64{
    	"DialogId" : e.DialogId,
    	"UserId" : e.user.ID(),
    }
    bytes, _ := json.Marshal(m)
    return bytes
}

func (e *LeaveDialogCommand) ParseData(msg *tgbotapi.Message) bool {
    str := strings.Split(msg.CommandArguments(), " ")
    if str[0] == "" {
    	return false
    }
    if len(str) != 1 {
    	return false
    }
    n, err := strconv.ParseInt(str[0], 10, 64)
    if err != nil {
    	return false
    }
    e.DialogId = n
    return true

}

func ( LeaveDialogCommand) Help() string {
    return "/leave id"
}
