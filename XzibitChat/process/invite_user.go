package process

import (
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

type InviteUserCommand struct {
	DefaultUserHandler
	DefaultAnswerParser
	DefaultCommandNameGetterSetter
	userName  string
	userID    int64
}

func (e InviteUserCommand) CommandName() string {
    return "InviteUserToDialog"
}

func (e InviteUserCommand) GetParsedJSON() []byte {
    m := map[string]interface{}{
    	"ToUserLogin": e.userName,
    	"DialogId": e.user.Hub().ID(),
    	"FromUserId": e.user.ID(),
    }
    bytes, _ := json.Marshal(m)
    return bytes
}

func (e InviteUserCommand) Help() string {
    return "/invite username"
}

func (e InviteUserCommand) IsGlobal() bool {
    return true
}

func (e InviteUserCommand) PreExecuteCheck() bool {
    msg := ""
    isOk := true
	switch {
	case !e.user.IsLogon():
		msg = "You are not logged in!"
		isOk = false
	case !e.user.InHub():
		msg = "You are not in hub!"
		isOk = false
	}
    if !isOk {
        e.Fail(msg)
    }
    return isOk
}

func (e *InviteUserCommand) Execute() {
	str := fmt.Sprintf("You is added to %s dialogue!ID: %d",
		e.user.Hub().Name(), e.user.Hub().ID())
    SendMessage(e.userID, str)
}

func (e *InviteUserCommand) ParseData(msg *tgbotapi.Message) bool {
	str := strings.Split(msg.CommandArguments(), " ")
	if str[0] == "" {
		return false
	}
	if len(str) != 1 {
		return false
	}
    e.userName = str[0]
    return true
}
