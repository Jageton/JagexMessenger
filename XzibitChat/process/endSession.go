package process

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type EndSession struct {
	DefaultUserHandler
	DefaultAnswerParser
	DefaultCommandNameGetterSetter
}

func ( EndSession) IsGlobal() bool {
    return false
}

func (e *EndSession) PreExecuteCheck() bool {
	return true
}

func (e *EndSession) Execute() {
    if e.user.InHub() {
    	e.user.Hub().LeaveHub(e.user.ID())
    }
    Processor.Users.Remove(e.user.ID())
}

func ( EndSession) GetParsedJSON() []byte {
    return nil
}

func ( EndSession) ParseData(*tgbotapi.Message) bool {
    return true
}

func ( EndSession) Help() string {
    return "/end"
}

func ( EndSession) CommandName() string {
    return ""
}

