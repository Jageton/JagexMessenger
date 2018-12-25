package process

import (
	"XzibitChat/tmessage"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type SendMessageCommand struct {
	DefaultUserHandler
	DefaultCommandNameGetterSetter
	DefaultAnswerParser
	msg  *tmessage.TMessage
}

func (s *SendMessageCommand) CommandName() string {
    return ""
}

func (s *SendMessageCommand) Help() string {
    return "[/send] ''You must be in Dialogue"
}

func (s *SendMessageCommand) GetParsedJSON() (bytes []byte) {
	return nil
}

func (s *SendMessageCommand) ParseData(msg *tgbotapi.Message) bool {
	s.msg = tmessage.NewTMessage(msg.From.UserName, msg.Text, s.user.ID())
    return true
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
