package kafka

import (
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)



var (
	kafkaConsumer *kafka.Writer
)

const (
	kafkaURL = "localhost:19092"
	kafktaTopic = "user_topic_vip"
)
// for producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}
// for consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		CommitInterval: time.Second,
		StartOffset: kafka.FirstOffset,
	})
}


type StockInfo struct{
	Message string `json:"message"`
	Type string `json:"type"`
}
