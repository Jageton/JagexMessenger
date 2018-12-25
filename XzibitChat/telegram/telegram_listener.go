package telegram

import (
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type BotListener struct {
	bot *tgbotapi.BotAPI
	Token string
}

type Listener struct {
	listener chan *tgbotapi.Message
	done     chan bool
}

func (l *Listener) GetMessage() *tgbotapi.Message{
	return <-l.listener
}

func (l *Listener) Close() {
	l.done <- true
}

func (t *BotListener) ParseConfig(config []byte) error {
	return json.Unmarshal(config, t)
}

func (t *BotListener) Connect() error {
	bot, err := tgbotapi.NewBotAPI(t.Token)
	if err != nil {
		return err
	}
	t.bot = bot
	return nil
}

func NewListener() *BotListener {
	return &BotListener{}
}

func (t *BotListener) Listener(queueSize int) (*Listener, error) {
	waiter := make(chan *tgbotapi.Message, queueSize)
	done := make(chan bool)
	up := tgbotapi.NewUpdate(0)
	updateChannel, err := t.bot.GetUpdatesChan(up)
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			select {
			case update, ok := <-updateChannel:
				if !ok {
					log.Println("Cannot get message")
				}
				waiter <- update.Message
				break
			case <-done:
				close(done)
				close(waiter)
				return
			}
		}
	}()
	return &Listener{
		listener: waiter,
		done: done,
	}, nil
}

type Sender struct {
	sender chan tgbotapi.MessageConfig
}

func (s Sender) Send(msg tgbotapi.MessageConfig) {
	s.sender <- msg
}

func (t *BotListener) Sender() *Sender {
	listener := make(chan tgbotapi.MessageConfig)
	go func() {
		for {
			select {
			case m, ok := <-listener:
				if !ok {
					log.Println("Cannot send message")
					continue
				}
				_, err := t.bot.Send(m)
				if err != nil {
					log.Println("Cannot send message ", err)
				}
			}
		}
	}()
	return &Sender{
		sender: listener,
	}
}
