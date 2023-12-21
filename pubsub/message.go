package pubsub

import "fmt"

type Message struct {
	topic string
	body  string
}

// New Message object with topic and message
func NewMessage(message string, topic string) *Message {
	fmt.Println(" func : NewMessage")
	return &Message{
		topic: topic,
		body:  message,
	}

}

// get message topic from message object
func (m *Message) GetTopic() string {
	fmt.Println(" func : GetTopic")
	return m.topic
}

// get message body from message object
func (m *Message) GetMessageBody() string {
	fmt.Println(" func : GetMessageBody")
	return m.body
}
