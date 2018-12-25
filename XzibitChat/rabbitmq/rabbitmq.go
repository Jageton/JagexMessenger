package rabbitmq

import (
	"XzibitChat/commands"
	"encoding/json"
	"github.com/streadway/amqp"
	"sync"
	"time"
)

type Queue struct {
	Name      string
	Consumer  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
}

type PublishConfig struct {
	Name       string
	Key        string
	Mandatory  bool
	Immediate  bool
}

type MsgBroker struct {
	connection         *amqp.Connection
	requests           *Requests
	channel            *amqp.Channel
	mutex              *sync.Mutex
	ConnectionString   string
	ConsumeQueue       *Queue
	PublishConfig       *PublishConfig
}

func NewMsgBroker() *MsgBroker {
	return &MsgBroker{
		mutex : &sync.Mutex{},
	}
}

func (m *MsgBroker) ParseConfig(config []byte) error {
	return json.Unmarshal(config, m)
}

func (m *MsgBroker) Connect() error {
	conn, err := amqp.Dial(m.ConnectionString)
	if err != nil {
		return err
	}

	m.requests = NewReqests()
	m.connection = conn
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	m.channel = ch
	answers, err := ch.Consume(m.ConsumeQueue.Name,
		                       m.ConsumeQueue.Consumer,
		                       m.ConsumeQueue.AutoAck,
		                       m.ConsumeQueue.Exclusive,
		                       m.ConsumeQueue.NoLocal,
		                       m.ConsumeQueue.NoWait,
		                       m.ConsumeQueue.Args)
	if err != nil {
		return err
	}
	go m.Listen(answers)
	return nil
}

func (m *MsgBroker) Send(command commands.Command){
	id := m.requests.Add(command)
	request := &Request{
		Id: id,
		Command: command.CommandName(),
		Data: command.GetParsedJSON(),
	}
	bytes, _ := json.Marshal(request)
	msg := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType: "text/plain",
		Body:         bytes,
	}
	m.mutex.Lock()
	err := m.channel.Publish(m.PublishConfig.Name,
		                     m.PublishConfig.Key,
		                     m.PublishConfig.Mandatory,
		                     m.PublishConfig.Immediate, msg)
	m.mutex.Unlock()
	if err != nil {
		command.Fail("Server is not responding!")
	}
}

func (m *MsgBroker) Listen(ch <-chan amqp.Delivery){
	for msg := range ch {
		answer, _ := parseAnswer(msg.Body)
		m.requests.Execute(answer)
		msg.Ack(false)
	}
}

func parseAnswer(body []byte) (*Answer, error) {
	a := &Answer{}
	err := json.Unmarshal(body, a)
	return a, err
}



