package utils

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"testdb/models"

	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
)

func ProcessMessage(db *sql.DB, redis *redis.Client, message []byte) {
	var data models.SingleImageEnroll
	err := proto.Unmarshal(message, &data)
	if err != nil {
		log.Fatalf("Failed to unmarshal message: %v", err)
	}
	var meta models.Meta
	err = json.Unmarshal([]byte(data.Meta), &meta)
	if err != nil {
		log.Fatal(err)
	}
	faceData := models.FaceData{
		UserId:        meta.UserId,
		FaceId:        meta.FaceId,
		FaceEmbedding: data.FaceEmbedding,
	}

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO public.face_data (user_id, face_id, face_embedding) VALUES ($1, $2, $3)", faceData.UserId, faceData.FaceId, faceData.FaceEmbedding)
	if err != nil {
		log.Fatal(err)
	}

	err = redis.Set(context.Background(), faceData.FaceId, faceData.UserId, 0).Err()
	if err != nil {
		log.Fatal(err)
	}
}
