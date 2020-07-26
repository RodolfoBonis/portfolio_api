package entities

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func UnmarshalUser(data []byte) (User, error) {
	var r User
	err := json.Unmarshal(data, &r)
	return r, err
}

func (base *User) Marshal() ([]byte, error) {
	return json.Marshal(base)
}

type User struct {
	ID          *uuid.UUID     `json:"id" gorm:"type:uuid;PRIMARY_KEY"`
	Name        *string        `json:"name" gorm:"type:text"`
	Password    *string        `json:"password,omitempty" gorm:"type:text"`
	Email       *string        `json:"email" gorm:"type:text;UNIQUE"`
	Nasc        *time.Time     `json:"nasc,omitempty" gorm:"type:timestamp"`
	Aboutme     *string        `json:"aboutme,omitempty" gorm:"type:text"`
	SocialMedia []*SocialMedia `json:"socialmedia,omitempty"`
	CreatedAt   *time.Time     `json:"created_at,omitempty" gorm:"type:timestamp"`
	UpdatedAt   *time.Time     `json:"updated_at,omitempty" gorm:"type:timestamp"`
	DeletedAt   *time.Time     `json:"deleted_at,omitempty" sql:"index"`
}

func hashPassword(password *string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), 14)
	return string(bytes), err
}

func (base *User) BeforeCreate(scope *gorm.Scope) {
	uuid := uuid.NewV4()
	hash, _ := hashPassword(base.Password)
	scope.SetColumn("Password", hash)
	scope.SetColumn("ID", uuid)
}
