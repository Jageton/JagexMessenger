package process

import (
	"XzibitChat/chat"
	"XzibitChat/rabbitmq"
	"XzibitChat/telegram"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Process struct {
	Hubs     *chat.DialogManager
	Users    *chat.UserManager
	Sender   *telegram.Sender
	RabbitMQ *rabbitmq.MsgBroker
}

var Processor Process

func StartProcess(bot *telegram.BotListener) error {
	listener, err := bot.Listener(128)
	if err != nil {
		return err
	}
	for {
		message := listener.GetMessage()
		go processMessage(message)
	}
}

func processMessage(message *tgbotapi.Message) {
	chatId := message.Chat.ID
	user, ok := Processor.Users.Get(chatId)
	if !ok {
		user = chat.NewUser(chatId)
		Processor.Users.Add(chatId, user)
	}
	if message.IsCommand() {
		processCommand(user, message.Command(), message)
	} else {
		processCommand(user, "send", message)
	}
}

func processCommand(user *chat.User, cmdName string, message *tgbotapi.Message) {
	cmd, ok := GetCommand(cmdName)
	if !ok {
		SendMessage(user.ID(), "Such command doesn't exist")
		return
	}
	cmd.SetUser(user)
	cmd.SetCommandName(cmdName)
	if !cmd.ParseData(message) {
		SendMessage(user.ID(), cmd.Help())
		return
	}
	if !cmd.PreExecuteCheck() {
		return
	}
	if cmd.IsGlobal() {
		Processor.RabbitMQ.Send(cmd)
	} else {
		cmd.Execute()
	}
}

func SendMessage(userId int64, msg string) {
	message := tgbotapi.NewMessage(userId, msg)
	Processor.Sender.Send(message)
}
