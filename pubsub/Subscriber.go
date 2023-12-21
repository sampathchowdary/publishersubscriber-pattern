package pubsub

import (
	"fmt"
	"log"
	"math/rand"
)

// Subscriber object with id, messages, topics, active or not
type Subscriber struct {
	id       string
	messages chan *Message // message channel
	topics   map[string]bool
	active   bool
}

// create new Subscriber
func CreateNewSubscriber() (string, *Subscriber) {
	fmt.Println(" func : CreateNewSubscriber")
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	id := fmt.Sprintf("%X-%X", b[0:4], b[4:8]) // create id
	return id, &Subscriber{
		id:       id,
		messages: make(chan *Message),
		topics:   map[string]bool{},
		active:   true,
	}
}

// add topic to subscriber
func (s *Subscriber) AddTopic(topic string) {
	// fmt.Println(" func : AddTopic")
	s.topics[topic] = true

}

// remove topics func to remove topics further
func (s *Subscriber) RemoveTopic(topic string) {
	// fmt.Println(" func : RemoveTopic")
	delete(s.topics, topic)
}

// get topics - all available topics
func (s *Subscriber) GetTopics() []string {
	// fmt.Println(" func : GetTopics")
	Totaltopics := []string{}
	for topic, _ := range s.topics {
		Totaltopics = append(Totaltopics, topic)
	}

	return Totaltopics
}

// get messages of subscriber is active
func (s *Subscriber) Signal(msg *Message) {
	// fmt.Println(" func : Signal")
	if s.active {
		s.messages <- msg
	}
}

// after removing subscriber - need to stop messages
func (s *Subscriber) Destruct() {
	// fmt.Println(" func : Destruct")
	s.active = false
	close(s.messages)
}

// Listen messages
func (s *Subscriber) Listen() {
	// fmt.Println(" func : Listen")
	for {
		if message, ok := <-s.messages; ok {
			fmt.Printf(" subscriber: %s, receiver: %s topic: %s", s.id, message.GetMessageBody(), message.GetTopic())
		}
	}
}
