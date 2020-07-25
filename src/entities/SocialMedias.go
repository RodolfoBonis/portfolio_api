package entities

import "encoding/json"

func UnmarshalSocialMedias(data []byte) (SocialMedias, error) {
	var r SocialMedias
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SocialMedias) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SocialMedias struct {
	ID        *string `json:"id,omitempty"`
	UserID    *string `json:"userId,omitempty"`
	Name      *string `json:"name,omitempty"`
	URL       *string `json:"url,omitempty"`
	Icon      *string `json:"icon,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}
