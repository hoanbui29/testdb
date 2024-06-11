package tests

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testdb/models"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestProto(t *testing.T) {
	meta := models.Meta{
		PersonId: "123",
		FaceId:   "456",
	}
	metaJson, err := json.Marshal(meta)
	if err != nil {
		t.Fatal(err)
	}
	data := models.SingleImageEnroll{
		Meta:          string(metaJson),
		FaceEmbedding: []byte(base64.StdEncoding.EncodeToString([]byte("face_embedding"))),
	}

	mar, err := proto.Marshal(&data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("test")
	fmt.Printf("%s\n", string(mar))
}
