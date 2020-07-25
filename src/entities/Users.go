
package entities

import "encoding/json"

func UnmarshalUsers(data []byte) (Users, error) {
	var r Users
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Users) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Users struct {
	ID        *uuid `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
	Nasc      *string `json:"nasc,omitempty"`
	Aboutme   *string `json:"aboutme,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}
