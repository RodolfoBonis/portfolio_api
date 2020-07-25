package entities

import "encoding/json"

func UnmarshalComments(data []byte) (Comments, error) {
	var r Comments
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Comments) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Comments struct {
	ID        *string `json:"id,omitempty"`
	PostID    *string `json:"postId,omitempty"`
	Author    *string `json:"author,omitempty"`
	Content   *string `json:"content,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}
