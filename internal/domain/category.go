package domain

import "time"

type Category struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	IconPath  *string   `json:"icon_path"`
	ParentId  *uint     `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
