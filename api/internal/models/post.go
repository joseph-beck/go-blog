package models

type Post struct {
	Model
	BlogRefer   uint   `json:"blogRefer"`
	AuthorRefer uint   `json:"authorRefer"`
	Author      Author `gorm:"foreignKey:AuthorRefer" json:"author"`

	Title string `gorm:"type:varchar(100)" json:"title"`
	Body  string `gorm:"type:text" json:"body"`
	Likes uint   `gorm:"default:0" json:"likes"`
}
