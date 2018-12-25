package process

import (
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

type RegistrationCommand struct {
	DefaultUserHandler
	DefaultAnswerParser
	DefaultCommandNameGetterSetter
	password string
	login    string
}

func (r RegistrationCommand) CommandName() string {
    return "Registration"
}

func ( RegistrationCommand) IsGlobal() bool {
    return true
}

func (r *RegistrationCommand) PreExecuteCheck() bool {
    msg := ""
    isOk := true
	switch  {
	case r.user.IsLogon():
		msg = "You are already logged in"
		isOk = false
	}
    if !isOk {
    	r.Fail(msg)
    }
    return isOk
}

func (r *RegistrationCommand) Execute() {
    SendMessage(r.user.ID(), "You are in system now!")
    r.user.SetLogin(true)
}

func (r *RegistrationCommand) GetParsedJSON() []byte {
    m := map[string]interface{}{
    	"Login"    : r.login,
    	"Password" : r.password,
    	"UserId"   : r.user.ID(),
    }
    bytes, _ := json.Marshal(m)
    return bytes
}

func (r *RegistrationCommand) ParseData(msg *tgbotapi.Message) bool {
    str := strings.Split(msg.CommandArguments(), " ")
    if str[0] == "" {
    	return false
    }
    if len(str) != 1 {
    	return false
    }
    r.password = str[0]
    r.login = msg.From.UserName
    return true
}

func ( RegistrationCommand) Help() string {
    return "/regr password"
}

