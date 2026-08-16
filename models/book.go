package models

type Book struct {
	ID          uint     `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string   `json:"title" gorm:"not null"`
	Author      string   `json:"author,omitempty" gorm:"default:Unknown"`
	ReleaseYear int      `json:"release_year"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `json:"category,omitempty"`
}
