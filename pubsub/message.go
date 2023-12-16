package pubsub

import "fmt"

type Message struct {
	topic string
	body  string
}

func NewMessage(msg string, topic string) *Message {
	fmt.Println(" func : NewMessage")
	return &Message{
		topic: topic,
		body:  msg,
	}

}

func (m *Message) GetTopic() string {
	fmt.Println(" func : GetTopic")
	return m.topic
}

func (m *Message) GetMessageBody() string {
	fmt.Println(" func : GetMessageBody")
	return m.body
}
