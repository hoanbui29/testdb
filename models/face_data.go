package models

type FaceData struct {
	PersonId      string `json:"person_id" fake:"{uuid}"`
	FaceId        string `json:"face_id" fake:"{uuid}"`
	FaceEmbedding []byte `json:"face_embedding" fake:"{bytearray}"`
}

type Meta struct {
	PersonId string `json:"person_id"`
	FaceId   string `json:"face_id"`
}
