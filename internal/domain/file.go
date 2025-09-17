package domain

import "time"

type File struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	IconPath   *string   `json:"icon_path"`
	FilePath   *string   `json:"file_path"`
	CategoryID *string   `json:"category_id,omitempty"` // nullable FK
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
