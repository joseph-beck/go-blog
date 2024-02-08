package models

type Author struct {
	Model
	Username string `json:"username"`
	Name     string `json:"name"`
}
