package main

import (
	"fmt"
	"math/rand"
	"publisherSubscriber/pubsub"
	"time"
)

var availableTopics = map[string]string{
	"BTC": "BITCOIN",
	"ETH": "ETHEREUM",
	"SAM": "SAMPATH",
}

func pricePublisher(broker *pubsub.Broker) {
	topicKeys := make([]string, 0, len(availableTopics))
	topicValues := make([]string, 0, len(availableTopics))

	for k, v := range availableTopics {
		topicKeys = append(topicKeys, k)
		topicValues = append(topicValues, v)
	}
	for {
		randValue := topicValues[rand.Intn(len(topicValues))]
		message := fmt.Sprintf("%f", rand.Float64())
		go broker.PublishMessage(randValue, message)

		r := rand.Intn(5)
		time.Sleep(time.Duration(r) * time.Second)

	}
}

func main() {
	broker := pubsub.NewBroker()
	s1 := broker.AddSubscriber()
	s2 := broker.AddSubscriber()
	s3 := broker.AddSubscriber()

	// s1 & s3 to only those specific topics
	broker.Subscribe(s1, availableTopics["BTC"])
	broker.Subscribe(s3, availableTopics["ETC"])
	// s2 to both topics
	broker.Subscribe(s2, availableTopics["BTC"])
	broker.Subscribe(s2, availableTopics["ETC"])

	defer broker.Subscribe(s3, availableTopics["SAM"])

	go pricePublisher(broker)

	go s3.Listen()
	go s1.Listen()

	go s2.Listen()
	fmt.Scanln()
	fmt.Println("Done !")
}
