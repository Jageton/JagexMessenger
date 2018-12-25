package process

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type HelpCommand struct {
	DefaultUserHandler
	DefaultCommandNameGetterSetter
	DefaultAnswerParser
}

func (HelpCommand) IsGlobal() bool {
	return false
}

func (HelpCommand) PreExecuteCheck() bool {
	return true
}

func (h *HelpCommand) Execute() {
	str := ""
	for _, c := range Commands {
		str += c.Help() + "\n"
	}
	SendMessage(h.user.ID(), str)
}

func (HelpCommand) GetParsedJSON() []byte {
	return nil
}

func (HelpCommand) ParseData(*tgbotapi.Message) bool {
	return true
}

func (HelpCommand) Help() string {
	return "/help"
}

func (HelpCommand) CommandName() string {
	return ""
}

