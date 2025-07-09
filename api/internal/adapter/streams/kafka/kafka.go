package kafka

import (
	"os"
	"sync"

	sarama "gopkg.in/Shopify/sarama.v1"
)

type KafkaStreamSender struct {
	producer sarama.SyncProducer
	topic    string
}

var (
	producerInstance sarama.SyncProducer
	producerOnce     sync.Once
	producerErr      error
)

func NewKafkaStreamSender(producerInstance sarama.SyncProducer, topic string) *KafkaStreamSender {
	return &KafkaStreamSender{
		producer: producerInstance,
		topic:    topic,
	}
}

func (k *KafkaStreamSender) Send(payload []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: k.topic,
		Value: sarama.ByteEncoder(payload),
	}
	_, _, err := k.producer.SendMessage(msg)
	return err
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func (k *KafkaStreamSender) Get(key string) (string, error) {
	return "", nil
}
