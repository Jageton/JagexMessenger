package rabbitmq

type Answer struct {
	Request int64
	Value map[string]interface{}
}
