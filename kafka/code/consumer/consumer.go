package main

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func main() {

	topic := "my-topic"
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	brokers := []string{"localhost:29092", // External port for kafka1
		"localhost:29093", // External port for kafka2
		"localhost:29094", // External port for kafka3}
	}
	consumerGroup := "my-topic-consumer-group"

	startConsumer(config, topic,consumerGroup, brokers)
}

func startConsumer(config *sarama.Config, topic string,consumerGroupName string, brokers []string) {
	fmt.Println("Consumer started")
	consumerGroup,err := sarama.NewConsumerGroup(brokers,consumerGroupName,config)
	if err != nil {
		log.Fatalf("Failed to create consumer group :%v",err)
	}	

	defer consumerGroup.Close()

	//get messages 
	for {
		err := consumerGroup.Consume(context.Background(),[]string{topic},&consumer{})
		if err != nil {
			log.Fatalf("Failed to consume messages: %v", err)
		}
	}

}

type consumer struct {}
func (consumer) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (consumer) Cleanup(sarama.ConsumerGroupSession) error { return nil }
func (consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	log.Printf("Consumer is focused on partition : %v",claim.Partition())
    for msg := range claim.Messages() {
        log.Printf("Message claimed: value = %s, Partition = %v, topic = %s \n", string(msg.Value), msg.Partition, msg.Topic)
		//mark the message as consumed 
        session.MarkMessage(msg, "")
    }
    return nil
}

