package entities

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

func UnmarshalSocialMedia(data []byte) (SocialMedia, error) {
	var r SocialMedia
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SocialMedia) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SocialMedia struct {
	ID        *uuid.UUID `json:"id,omitempty" gorm:"type:uuid;PRIMARY_KEY"`
	UserID    *uuid.UUID `json:"userId,omitempty" gorm:"type:uuid"`
	Name      *string    `json:"name,omitempty" gorm:"type:text"`
	URL       *string    `json:"url,omitempty" gorm:"type:text"`
	Icon      *string    `json:"icon,omitempty" gorm:"type:text"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"type:timestamp"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"type:timestamp"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index"`
}

func (base *SocialMedia) BeforeCreate(scope *gorm.Scope) {
	uuid := uuid.NewV4()
	scope.SetColumn("ID", uuid)
}
