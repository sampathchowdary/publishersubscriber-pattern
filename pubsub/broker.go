package pubsub

import "fmt"

type Subscribers map[string]*Subscriber

type Broker struct {
	subscribers Subscribers
	topics      map[string]Subscribers
}

// create Broker
func NewBroker() *Broker {
	return &Broker{
		subscribers: Subscribers{},
		topics:      map[string]Subscribers{},
	}
}

// Add Subscriber
func (b *Broker) AddSubscriber() *Subscriber {
	id, s := CreateNewSubscriber()
	b.subscribers[id] = s
	return s
}

// Remove Subscriber
func (b *Broker) RemoveSubscriber(s *Subscriber) {
	for topic := range s.topics {
		b.Unsubscribe(s, topic)
	}
	delete(b.subscribers, s.id)
	s.Destruct()
}

// Send Messages - msg tp topic (signal to update message)
func (b *Broker) SendMessage(message string, topics []string) {
	for _, topic := range topics {
		for _, s := range b.topics[topic] {
			mes := NewMessage(message, topic)
			go (func(s *Subscriber) {
				s.Signal(mes) // signal function
			})(s)
		}
	}
}

// get subscribers specific to topic
func (b *Broker) GetSubscribers(topic string) int {
	return len(b.topics[topic])
}

// subscribe to specific topic
func (b *Broker) Subscribe(s *Subscriber, topic string) {
	if b.topics[topic] == nil {
		b.topics[topic] = Subscribers{}
	}
	s.AddTopic(topic)
	b.topics[topic][s.id] = s
	fmt.Printf(" %s Subscribed to topic: %s\n", s.id, topic)
}

// unsubscribe from specific topic
func (b *Broker) Unsubscribe(s *Subscriber, topic string) {
	delete(b.topics[topic], s.id)
	s.RemoveTopic(topic)

	fmt.Printf(" %s Removed from Subscribed topic: %s\n", s.id, topic)
}

// publish messages
func (b *Broker) PublishMessage(topic string, message string) {
	for _, s := range b.topics[topic] {
		m := NewMessage(message, topic)
		if !s.active {
			return
		}
		go (func(s *Subscriber) {
			s.Signal(m) // signal function
		})(s)
	}
}
