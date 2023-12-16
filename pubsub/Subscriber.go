package pubsub

import (
	"fmt"
)

type Subscriber struct {
	id string
}

func CreateNewSubscriber() (string, *Subscriber) {
	fmt.Println(" func : CreateNewSubscriber")

}

func (s *Subscriber) AddTopic(topic string) {
	fmt.Println(" func : AddTopic")

}

func (s *Subscriber) RemoveTopic(topic string) {
	fmt.Println(" func : RemoveTopic")

}

func (s *Subscriber) GetTopics() []string {
	fmt.Println(" func : GetTopics")

	return []string{}

}
