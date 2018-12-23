package rabbitmq

type Request struct {
	Id            int64
	CommandName   string
	Value       []byte
}
