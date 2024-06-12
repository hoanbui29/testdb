package database

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

func GetKafkaConsumer(partition int) *kafka.Reader {
	consumer := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{fmt.Sprintf("%s:%d", viper.GetString("base.host"), viper.GetInt("kafka.port"))},
		Topic:   viper.GetString("kafka.topic"),
		// GroupID:   viper.GetString("kafka.consumer.group_id"),
		Partition: partition,
		MaxBytes:  10e8,
	})

	return consumer
}

func GetLatestOffset(partition int) int64 {
	conn, err := kafka.DialLeader(context.Background(), "tcp", fmt.Sprintf("%s:%d", viper.GetString("base.host"), viper.GetInt("kafka.port")), viper.GetString("kafka.topic"), partition)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	endOffset, err := conn.ReadLastOffset()
	if err != nil {
		log.Fatal(err)
	}
	return endOffset
}

func GetKafkaWriter() *kafka.Writer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{fmt.Sprintf("%s:%d", viper.GetString("base.host"), viper.GetInt("kafka.port"))},
		Topic:    "test",
		Balancer: &kafka.LeastBytes{},
	})

	return writer
}
