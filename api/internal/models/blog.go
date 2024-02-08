package models

type Blog struct {
	Model
	AuthorRefer uint   `json:"authorRefer"`
	Author      Author `gorm:"foreignKey:AuthorRefer" json:"author"`

	Name        string `gorm:"type:varchar(100)" json:"name"`
	Description string `gorm:"type:varchar(255)" json:"description"`
}
