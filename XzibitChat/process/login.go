package process

import (
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

type LoginCommand struct {
	DefaultUserHandler
	DefaultAnswerParser
	DefaultCommandNameGetterSetter
	password string
}

func (l *LoginCommand) CommandName() string {
    return "Auth"
}

func (l *LoginCommand) GetParsedJSON() []byte {
    m := map[string]interface{}{
    	"Password": l.password,
    	"UserId": l.user.ID(),
    }
    bytes, _ := json.Marshal(m)
    return bytes

}

func (l *LoginCommand) ParseData(msg *tgbotapi.Message) bool {
	passw := strings.Split(msg.CommandArguments(), " ")
	if passw[0] == "" {
		return false
	}
	if len(passw) != 1 {
		return false
	}
	l.password = msg.CommandArguments()
	return true
}

func (l *LoginCommand) Help() string {
    return "/login password"
}

func (l *LoginCommand) Execute() {
	l.user.SetLogin(true)
}

func (l *LoginCommand) IsGlobal() bool {
    return true
}

func (l *LoginCommand) PreExecuteCheck() bool {
	if l.user.IsLogon() {
		msg := tgbotapi.NewMessage(l.user.ID(), "You are already logged in!")
		Processor.Sender.Send(msg)
		return false
	}
	return true
}
