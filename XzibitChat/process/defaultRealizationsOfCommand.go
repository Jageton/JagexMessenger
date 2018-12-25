package process

import (
	"XzibitChat/chat"
	"encoding/json"
	"errors"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

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

func (DefaultAnswerParser) ParseAnswer(args []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(args, &m)
	if err != nil {
		return err
	}
	if exc, ok := m["Exception"]; ok {
		exception := exc.(string)
		if exception == "" {
			return nil
		} else {
			return errors.New(exception)
		}
	}
	return errors.New("servers problem")
}

