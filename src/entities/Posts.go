package entities

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"time"
)

func UnmarshalPost(data []byte) (Post, error) {
	var r Post
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Post) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Post struct {
	ID          *uuid.UUID `json:"id,omitempty" gorm:"type:uuid;PRIMARY_KEY"`
	UserID      *uuid.UUID `json:"userId,omitempty" gorm:"type:uuid"`
	Title       *string    `json:"title,omitempty" gorm:"type:text"`
	Description *string    `json:"description,omitempty" gorm:"type:text"`
	User        *User      `json:"user,omitempty"`
	Tags        *[]Tag     `json:"tags,omitempty"`
	Comments    *[]Comment `json:"comments,omitempty"`
	Like        *int       `json:"like,omitempty" gorm:"type:integer"`
	Photos      *string    `json:"photos,omitempty" gorm:"type:text"`
	CreatedAt   *time.Time `json:"created_at,omitempty" gorm:"type:timestamp"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" gorm:"type:timestamp"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" sql:"index"`
}
