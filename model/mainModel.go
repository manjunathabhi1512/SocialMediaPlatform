package model

import (
	"time"

	"github.com/google/uuid"
)

type PostData struct {
	PostId    uuid.UUID
	Context   string
	Likes     int
	Dislikes  int
	Comments  []CommentsData
	PostedAt time.Time
}

type CommentsData struct {
	CommentId   uuid.UUID
	Comment     string
	CommentedAt time.Time
}

type PostPayload struct{
	ContextData string `json:"context"`
}