package entities

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"time"
)

func UnmarshalComment(data []byte) (Comment, error) {
	var r Comment
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Comment) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Comment struct {
	ID        *uuid.UUID `json:"id,omitempty" gorm:"type:uuid;PRIMARY_KEY"`
	PostID    *uuid.UUID    `json:"postId,omitempty" gorm:"type:uuid"`
	Author    *string    `json:"author,omitempty"`
	Content   *string    `json:"content,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index"`
}
