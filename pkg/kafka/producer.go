package kafka

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/Shopify/sarama"
)

// ProducerConfig configures a Kafka producer
type ProducerConfig struct {
	Brokers   string
	TopicName string
}

// Producer produces Kafka messages
type Producer struct {
	syncProducer sarama.SyncProducer
	topicName    string
	mutex        sync.Mutex
}

// NewProducer creates a new Producer
func NewProducer(config *ProducerConfig) (*Producer, error) {
	if len(config.Brokers) < 1 {
		return nil, errors.New("must specify at least 1 broker")
	}
	if config.TopicName == "" {
		return nil, errors.New("must provide a topic name")
	}

	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	saramaConfig := sarama.NewConfig()
	saramaConfig.ClientID = hostname
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Return.Successes = true
	saramaConfig.Producer.Return.Errors = true

	kafkaBrokers := strings.Split(config.Brokers, ",")
	producer, err := sarama.NewSyncProducer(kafkaBrokers, saramaConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to initialise Kafka producer: %v", err)
	}
	return &Producer{
		syncProducer: producer,
		topicName:    config.TopicName,
		mutex:        sync.Mutex{},
	}, nil
}

// Message represents a Kafka message
type Message struct {
	Key   string
	Value []byte
}

// Produce produces a Kafka message or returns an error
func (p *Producer) Produce(msg *Message) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	produceMsg := &sarama.ProducerMessage{
		Topic: p.topicName,
		Key:   sarama.StringEncoder(msg.Key),
		Value: sarama.ByteEncoder(msg.Value),
	}

	// TODO: add retries
	if _, _, err := p.syncProducer.SendMessage(produceMsg); err != nil {
		return fmt.Errorf("failed to produce Kafka message: %v", err)
	}
	return nil
}

// Close closes the producer
func (p *Producer) Close() {
	p.Close()
}
