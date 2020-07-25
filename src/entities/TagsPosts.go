package entities

import "encoding/json"

func UnmarshalTagsPosts(data []byte) (TagsPosts, error) {
	var r TagsPosts
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TagsPosts) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TagsPosts struct {
	ID        *string `json:"id,omitempty"`
	PostID    *string `json:"postId,omitempty"`
	TagID     *string `json:"tagId,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}
