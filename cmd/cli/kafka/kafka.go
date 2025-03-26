package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL    = "localhost:9092"
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
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		StartOffset:    kafka.FirstOffset,
	})
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func newStock(msg, typeMsg string) *StockInfo {
	return &StockInfo{
		Message: msg,
		Type:    typeMsg,
	}
}

func actionStock(c *gin.Context) {
	s := newStock(c.Query("message"), c.Query("type"))
	body := make(map[string]interface{})
	body["action"] = "action"
	body["info"] = s

	jsonBody, _ := json.Marshal(body)
	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}

	err := kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"error": "",
		"msg":   "success",
	})
}

func RegisterConsumer(id int) {
	// group
	kafkaGroupId := "consumer-group-"
	reader := getKafkaReader(kafkaURL, kafktaTopic, kafkaGroupId)
	defer reader.Close()

	fmt.Printf("Consumer %d is running\n", id)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Error reading message", err)
		}
		fmt.Printf("Consumer %d: Message on topic: %v, offset: %v , time: %d %s = %s \n", id, m.Topic, m.Partition, m.Offset, m.Time.Unix(),
			string(m.Key), string(m.Value))

	}
}

func main() {
	r := gin.Default()
	kafkaProducer = getKafkaWriter(kafkaURL, kafktaTopic)
	defer kafkaProducer.Close()

	r.POST("action/stock", actionStock)

	// register 2 user buy stock id 1 va id 2
	go RegisterConsumer(1)
	go RegisterConsumer(2)

	r.Run(":8999")
}
