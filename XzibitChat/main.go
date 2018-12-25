package main

import (
	"XzibitChat/chat"
	"XzibitChat/process"
	"XzibitChat/rabbitmq"
	"XzibitChat/telegram"
	"encoding/json"
	"log"
	"os"
)

func main() {
	bot := ReadTelegramBotConfig()
	rabbit := ReadRabbitConfig()
	err := rabbit.Connect()
	if err != nil {
		log.Fatal(err)
	}
	err = bot.Connect()
	if err != nil {
		log.Fatal(err)
	}
	process.Processor = process.Process{
		Users:    chat.NewUserManager(),
		Hubs:     chat.NewHubManager(),
		Sender:   bot.Sender(),
		RabbitMQ: rabbit,
	}
	err = process.StartProcess(bot)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan bool)
	<-done
}


func ReadTelegramBotConfig() *telegram.BotListener{
	tbconfig, err := os.Open("TelegramBotConfig.json")
	if err != nil {
		log.Fatal("Error open TelegramBotConfig.json")
	}
	tbdecoder := json.NewDecoder(tbconfig)
	botListener := telegram.NewListener()
	err = tbdecoder.Decode(botListener)
	if err != nil {
		log.Fatal(err)
	}
	return botListener
}

func ReadRabbitConfig() *rabbitmq.MsgBroker {
	rabbitconfigfile, err := os.Open("RabbitMQConfig.json")
	if err != nil {
		log.Fatal("Error open RabbitMQConfig.json")
	}
	rabbitdecoder := json.NewDecoder(rabbitconfigfile)
	rabbit := rabbitmq.NewMsgBroker()
	err = rabbitdecoder.Decode(rabbit)
	if err != nil {
		log.Fatal(err)
	}
	return rabbit
}
