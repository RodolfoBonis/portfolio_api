package entities

import "encoding/json"

func UnmarshalPosts(data []byte) (Posts, error) {
	var r Posts
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Posts) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Posts struct {
	ID          *string `json:"id,omitempty"`
	UserID      *string `json:"userId,omitempty"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Like        *int64  `json:"like,omitempty"`
	Photos      *string `json:"photos,omitempty"`
	CreatedAt   *string `json:"created_at,omitempty"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
	DeletedAt   *string `json:"deleted_at,omitempty"`
}
