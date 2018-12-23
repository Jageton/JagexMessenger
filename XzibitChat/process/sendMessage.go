package process

import (
	"XzibitChat/tmessage"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type SendMessageCommand struct {
	DefaultUserHandler
	DefaultCommandNameGetterSetter
	msg  *tmessage.TMessage
}

func (s *SendMessageCommand) CommandName() string {
    return ""
}

func (s *SendMessageCommand) Help() string {
    return "You must be in Dialogue"
}

func (s *SendMessageCommand) GetParsedJSON() (bytes []byte) {
	return nil
}

func (s *SendMessageCommand) ParseData(*tgbotapi.Message) bool {
    return true
}

func (s *SendMessageCommand) ParseAnswer(map[string]interface{}) error {
    return nil
}

func (s *SendMessageCommand) Execute() {
	s.user.SendMessageToHub(Processor.Sender, s.msg)
}

func (s *SendMessageCommand) IsGlobal() bool {
    return false
}

func (s *SendMessageCommand) PreExecuteCheck() bool {
	isOk := true
	msg := ""
	switch {
	case !s.user.IsLogon():
		isOk = false
		msg = "You are not logged in!"
	case !s.user.InHub():
		isOk = false
		msg = "You are not in any hub!"
	}
	if !isOk {
		msg := tgbotapi.NewMessage(s.user.ID(), msg)
		Processor.Sender.Send(msg)
	}
	return isOk
}
