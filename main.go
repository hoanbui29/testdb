package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"testdb/database"
	"testdb/models"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"
)

func processMessage(db *sql.DB, redis *redis.Client, message []byte) {
	var data models.SingleImageEnroll
	err := proto.Unmarshal(message, &data)
	if err != nil {
		log.Fatal(err)
	}
	var meta models.Meta
	err = json.Unmarshal([]byte(data.Meta), &meta)
	if err != nil {
		log.Fatal(err)
	}
	faceData := models.FaceData{
		PersonId:      meta.PersonId,
		FaceId:        meta.FaceId,
		FaceEmbedding: data.FaceEmbedding,
	}

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO public.face_data (person_id, face_id, face_embedding) VALUES ($1, $2, $3)", faceData.PersonId, faceData.FaceId, faceData.FaceEmbedding)
	if err != nil {
		log.Fatal(err)
	}
}

func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	return viper.ReadInConfig()
}

func main() {
	err := Init()
	if err != nil {
		log.Fatal(err)
	}
	pgDb := database.OpenPg()
	redis := database.OpenRedis()
	consumer := database.GetKafkaConsumer()
	defer consumer.Close()
	defer pgDb.Close()
	defer redis.Close()
	for {
		m, err := consumer.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
			break
		}
		processMessage(pgDb, redis, m.Value)
	}
}
