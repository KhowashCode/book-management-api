package models

type Category struct {
	ID    uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"category_name" gorm:"not null"`
	Books []Book `json:"books,omitempty"`
}
