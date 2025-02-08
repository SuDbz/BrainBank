package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/IBM/sarama"
)

// if topic created return true, otherwise return false
func createTopicIfNoteExist(config *sarama.Config, topic string, brokers []string) bool {

	client, err := sarama.NewClient(brokers, config)
	if err != nil {
		log.Fatalf("Failed to start the producer : %v", err)
	}
	defer client.Close()

	//Topic create is part of admin client hence create an admin client
	adminClient, err := sarama.NewClusterAdminFromClient(client)
	if err != nil {
		log.Fatalf("Failed to start admin client : %v", err)
	}

	defer adminClient.Close()

	//List all topics
	topicListMap, err := adminClient.ListTopics()
	if err != nil {
		log.Fatalf("Failed to fetch topics from admin client : %v", err)
	}

	if topicDetails, exist := topicListMap[topic]; !exist {
		//create the topic
		newTopic := sarama.TopicDetail{
			NumPartitions:     1,
			ReplicationFactor: 1,
			ConfigEntries:     nil,
		}

		//create topic
		err = adminClient.CreateTopic(topic, &newTopic, false)
		if err != nil {
			log.Fatalf("Failed to create topic %v with details :%v using admin client :%v", topic, newTopic, err)
		}

		fmt.Printf("Successfully created topic :%v", topic)
		return true
	} else {
		fmt.Printf("topic already exist  :%v ", topicDetails)
	}

	return false

}

func startProducer(config sarama.Config, topic string, brokers []string) {
	fmt.Println("\nstarting producer...")
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, &config)
	if err != nil {
		log.Fatalf("Failed to start the producer : %v", err)
	}
	defer producer.Close()

	//start a ticker
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {

		select {

		case <-ticker.C: //its time to create a message
			idMin := 1
			idMax := 10
			itemUUID := rand.Intn(idMax-idMin) + idMin
			min := 1
			max := 5
			randomNum := rand.Intn(max-min) + min

			order := map[string]interface{}{
				"id":      fmt.Sprint(itemUUID),
				"type":    fmt.Sprintf("P%d", randomNum),
				"item-id": fmt.Sprint(randomNum),
			}

			fmt.Printf("Message to publish : %v \n", order)

			message := &sarama.ProducerMessage{
				Topic: topic,
				Value: sarama.StringEncoder(fmt.Sprint(order)),
				Key:   sarama.StringEncoder(fmt.Sprint(itemUUID)),
			}

			partition, offset, err := producer.SendMessage(message)
			if err != nil {
				log.Fatalf("Failed to send message :%v", err)
			} else {
				fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
			}
		}
	}
}

func main() {
	topic := "my-topic"
	config := sarama.NewConfig()
	brokers := []string{"localhost:29092", // External port for kafka1
		"localhost:29093", // External port for kafka2
		"localhost:29094", // External port for kafka3}
	}
	createTopicIfNoteExist(config, topic, brokers)
	startProducer(*config, topic, brokers)
}
