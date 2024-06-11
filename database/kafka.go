package database

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
)

func GetKafkaConsumer() *kafka.Reader {
	consumer := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{fmt.Sprintf("%s:%d", viper.GetString("base.host"), viper.GetInt("kafka.port"))},
		Topic:    "test",
		GroupID:  viper.GetString("kafka.consumer.group_id"),
		MaxBytes: 10e8,
	})

	return consumer
}

func GetKafkaWriter() *kafka.Writer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{fmt.Sprintf("%s:%d", viper.GetString("base.host"), viper.GetInt("kafka.port"))},
		Topic:    "test",
		Balancer: &kafka.LeastBytes{},
	})

	return writer
}
