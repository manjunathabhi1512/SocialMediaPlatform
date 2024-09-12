package model

import (
	"time"

	"github.com/google/uuid"
)

type PostData struct {
	PostId   uuid.UUID
	Context  string
	Likes    int
	Dislikes int
	Comments []CommentsData
	PostedAt time.Time
}

type CommentsData struct {
	CommentId   uuid.UUID
	Comment     string
	CommentedAt time.Time
}

type PostPayload struct {
	ContextData string `json:"context"`
}

//db tables

type PostDataToDB struct {
	PostID   uuid.UUID `gorm:"column:postid;type:uuid;primary_key"`
	Context  string    `gorm:"column:context;type:text;not null"`
	Likes    int       `gorm:"column:likes;default:0"`
	Dislikes int       `gorm:"column:dislikes;default:0"`
	PostedAt time.Time `gorm:"column:posted_at;default:current_timestamp"`
}

type CommentDataToDB struct {
	CommentID   uuid.UUID `gorm:"column:commentid;type:uuid;primary_key"`
	Comment     string    `gorm:"column:comment;type:text;not null"`
	PostID      uuid.UUID `gorm:"column:post_id;type:uuid;not null"`
	CommentedAt time.Time `gorm:"column:commented_at;default:current_timestamp"`
}
