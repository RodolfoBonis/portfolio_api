package entities

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
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
	ID        *uuid.UUID `json:"id,omitempty" gorm:"type:uuid;PRIMARY_KEY"`
	User      *User      `json:"user" gorm:"association_autoupdate:false;association_autocreate:false"`
	UserID    uuid.UUID  `gorm:"type:uuid;" json:"user_id,omitempty"`
	Title     *string    `json:"title,omitempty" gorm:"type:text"`
	Content   *string    `json:"content,omitempty" gorm:"type:text"`
	Tags      *[]Tag     `json:"tags,omitempty" gorm:"association_autoupdate:false;association_autocreate:false"`
	Like      *int       `json:"like,omitempty" gorm:"type:integer"`
	Photos    *string    `json:"photos,omitempty" gorm:"type:text"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"type:timestamp"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"type:timestamp"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index"`
}

func (base *Post) BeforeCreate(scope *gorm.Scope) {
	uuid := uuid.NewV4()
	scope.SetColumn("ID", uuid)
}
