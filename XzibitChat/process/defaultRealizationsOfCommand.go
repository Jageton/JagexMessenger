package process

import (
	"XzibitChat/chat"
	"XzibitChat/commands"
	"errors"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mitchellh/mapstructure"
)

func Parse(args map[string]interface{}, cmd *commands.Command) error {
	return mapstructure.Decode(args, cmd)
}

type DefaultUserHandler struct {
	user *chat.User
}

func (d *DefaultUserHandler) SetUser(user *chat.User){
	d.user = user
}

func (d *DefaultUserHandler) Fail(msg string){
	SendMessage(d.user.ID(), msg)
}

type DefaultFail struct {
}

func (DefaultFail) Fail(id int64, msg string){
	message := tgbotapi.NewMessage(id, msg)
	Processor.Sender.Send(message)
}

type DefaultCommandNameGetterSetter struct {
	CommandName string
}

func (c *DefaultCommandNameGetterSetter) SetCommandName(cmdName string) {
	c.CommandName = cmdName
}

func (c *DefaultCommandNameGetterSetter) GetCommandName() string {
	return c.CommandName
}

type DefaultAnswerParser struct {

}

func (DefaultAnswerParser) ParseAnswer(args map[string]interface{}) error {
	if exc, ok := args["exception"]; ok {
		exception := exc.(string)
		if exception == "" {
			return nil
		} else {
			return errors.New(exception)
		}
	}
	return errors.New("servers problem")
}

