package rabbitmq

type Request struct {
	Id            int64
	Command   string
	Data       []byte
}
