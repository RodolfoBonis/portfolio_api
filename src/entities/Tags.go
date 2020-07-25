package entities

import "encoding/json"

func UnmarshalTags(data []byte) (Tags, error) {
	var r Tags
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Tags) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Tags struct {
	ID        *string `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}
