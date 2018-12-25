package process

import "github.com/go-telegram-bot-api/telegram-bot-api"

type ExitDialogCommand struct {
	DefaultUserHandler
	DefaultCommandNameGetterSetter
	DefaultAnswerParser
}

func (e ExitDialogCommand) CommandName() string {
    return ""
}

func (e ExitDialogCommand) GetParsedJSON() []byte {
	panic("implement me")
}

func (e ExitDialogCommand) ParseData(*tgbotapi.Message) bool {
	return true
}

func (e ExitDialogCommand) Help() string {
	return "/exit You will exit dialog for a time. You can come back in dialog."
}

func ( ExitDialogCommand) IsGlobal() bool {
    return false
}

func (e *ExitDialogCommand) PreExecuteCheck() bool {
	msg := ""
	isOK := true
	switch  {
	case !e.user.IsLogon():
		msg = "You are not logged in"
		isOK = false
	case !e.user.InHub():
		msg = "You are not in Hub"
		isOK = false
	}
	if !isOK {
		e.Fail(msg)
	}
	return isOK
}

func (e *ExitDialogCommand) Execute() {
	e.user.Hub().LeaveHub(e.user.ID())
	SendMessage(e.user.ID(), "You left dialog!")
}

