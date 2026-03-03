package models

type Snippet struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `json:"name"`
	Language    string `json:"language"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Code string `json:"code"`
}

