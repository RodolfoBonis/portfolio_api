package entities

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"time"
)

func UnmarshalTag(data []byte) (Tag, error) {
	var r Tag
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Tag) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Tag struct {
	ID        *uuid.UUID `json:"id,omitempty" gorm:"type:uuid;PRIMARY_KEY"`
	Name      *string `json:"name,omitempty" gorm:"type:text"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"type:timestamp"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"type:timestamp"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index"`
}
