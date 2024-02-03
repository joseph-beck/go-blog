package models

type User struct {
	Model
	Username string `gorm:"type:varchar(25);unique;not null" json:"username"`
	Email    string `gorm:"type:varchar(100);unique;not null;unique_index" json:"email"`
}
