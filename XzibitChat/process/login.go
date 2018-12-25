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
	login string
	password string
}

func (l *LoginCommand) CommandName() string {
    return "auth"
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
	str := strings.Split(msg.CommandArguments(), " ")
	if str[0] == "" {
		return false
	}
	if len(str) != 2 {
		return false
	}
	l.login = str[0]
	l.password = str[1]
	return true
}

func (l *LoginCommand) Help() string {
    return "/login login password"
}

func (l *LoginCommand) Execute() {
	l.user.SetLogin(true)
	SendMessage(l.user.ID(), "Success!")
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
