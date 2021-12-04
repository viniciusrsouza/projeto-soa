package events

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/viniciusrsouza/projeto-soa/order/config"
)

type KafkaPublisher struct {
	producer *kafka.Producer
	config   config.Kafka
}

func NewKafkaPublisher(kafkaCfg config.Kafka) KafkaPublisher {
	return KafkaPublisher{
		config: kafkaCfg,
	}
}

type Message []byte

func (k *KafkaPublisher) PublishMessage(topic Topic, msg Message) error {
	deliveryCh := make(chan kafka.Event)

	t := topic.String()
	message := kafka.Message{
		Value: msg,
		TopicPartition: kafka.TopicPartition{
			Topic:     &t,
			Partition: kafka.PartitionAny,
		},
	}

	err := k.producer.Produce(&message, deliveryCh)
	if err != nil {
		return err
	}

	e := <-deliveryCh
	if e.(*kafka.Message).TopicPartition.Error != nil {
		return fmt.Errorf("delivery failed: %w", e.(*kafka.Message).TopicPartition.Error)
	}

	return nil
}

func (k *KafkaPublisher) Start() error {
	var err error
	k.producer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": k.config.Servers,
		"acks":              -1,
	})
	if err != nil {
		return fmt.Errorf("could not start kafka producer: %w", err)
	}

	return nil
}
