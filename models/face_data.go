package models

type FaceData struct {
	UserId        string `json:"user_id" fake:"{uuid}"`
	FaceId        string `json:"face_id" fake:"{uuid}"`
	FaceEmbedding []byte `json:"face_embedding" fake:"{bytearray}"`
}

type Meta struct {
	UserId string `json:"user_id"`
	FaceId string `json:"face_id"`
}
