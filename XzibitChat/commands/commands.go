package commands

import (
	"XzibitChat/chat"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type UserHandler interface {
	SetUser(*chat.User)
	Fail(string)
}

type CommandNameGetterSetter interface {
	SetCommandName(string)
	GetCommandName() string
}

type AnswerParser interface {
	ParseAnswer([]byte) error
}

type Command interface {
	UserHandler
	CommandNameGetterSetter
	AnswerParser
	IsGlobal() bool
	PreExecuteCheck() bool
	Execute()
	GetParsedJSON() []byte
	ParseData(*tgbotapi.Message) bool
	Help() string
	CommandName() string
}
