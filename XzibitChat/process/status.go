package process

import "github.com/go-telegram-bot-api/telegram-bot-api"

type StatusCommand struct {
	DefaultUserHandler
	DefaultCommandNameGetterSetter
	DefaultAnswerParser
}

func ( StatusCommand) IsGlobal() bool {
	return false
}

func ( StatusCommand) PreExecuteCheck() bool {
	return true
}

func (s *StatusCommand) Execute() {
	msg := ""
	if !s.user.IsLogon() {
		msg = "You are not logged in!"
	}else {
		if !s.user.InHub() {
			msg = "You are not in any dialog!"
		} else {
			msg = "You are in " + s.user.Hub().Name() + " dialog!"
		}
	}
	SendMessage(s.user.ID(), msg)
}

func ( StatusCommand) GetParsedJSON() []byte {
	return nil
}

func ( StatusCommand) ParseData(*tgbotapi.Message) bool {
	return true
}

func ( StatusCommand) Help() string {
	return "/status"
}

func ( StatusCommand) CommandName() string {
	return ""
}

