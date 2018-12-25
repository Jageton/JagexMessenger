package tmessage



type TMessage struct {
	From string
	Text string
	Id   int64
}

func NewTMessage(from, text string, id int64) *TMessage{
	return &TMessage{
		From: from,
		Text: text,
		Id: id,
	}
}