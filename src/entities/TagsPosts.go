package entities

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"time"
)

func UnmarshalTagsPosts(data []byte) (TagsPosts, error) {
	var r TagsPosts
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TagsPosts) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TagsPosts struct {
	ID        *uuid.UUID `json:"id,omitempty" gorm:"type:uuid;PRIMARY_KEY"`
	PostID    *uuid.UUID `json:"post_id,omitempty" gorm:"type:uuid"`
	TagID     *uuid.UUID `json:"tagId,omitempty" gorm:"type:uuid"`
	Tags      *[]Tag     `json:"tags,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"type:timestamp"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"type:timestamp"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index"`
}
