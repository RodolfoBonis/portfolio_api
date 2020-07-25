package entities

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"time"
)

func UnmarshalUser(data []byte) (User, error) {
	var r User
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *User) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type User struct {
	ID          *uuid.UUID     `json:"id,omitempty" gorm:"type:uuid;PRIMARY_KEY"`
	Name        *string        `json:"name,omitempty" gorm:"type:text"`
	Nasc        *time.Time     `json:"nasc,omitempty" gorm:"type:timestamp"`
	Aboutme     *string        `json:"aboutme,omitempty" gorm:"type:text"`
	SocialMedia *[]SocialMedia `json:"socialmedia,omitempty"`
	CreatedAt   *time.Time     `json:"created_at,omitempty" gorm:"type:timestamp"`
	UpdatedAt   *time.Time     `json:"updated_at,omitempty" gorm:"type:timestamp"`
	DeletedAt   *time.Time     `json:"deleted_at,omitempty" sql:"index"`
}
