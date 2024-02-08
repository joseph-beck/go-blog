package models

type Blog struct {
	Model
	Owner uint `json:"owner"`

	Name        string `gorm:"type:varchar(100)" json:"name"`
	Description string `gorm:"type:varchar(255)" json:"description"`
}
