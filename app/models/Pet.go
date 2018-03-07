package models

// Pet model
type Pet struct {
	ID         uint `gorm:"primary_key"`
	Name       string
	Status     string
	CategoryID uint
	TagID      uint
}
