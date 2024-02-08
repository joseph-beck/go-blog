package models

type Author struct {
	Model
	Username string `gorm:"type:varchar(25);unique;not null" json:"username"`              // username of the author, must be unique
	Name     string `gorm:"type:varchar(25);not null" json:"name"`                         // name of the author, can be a duplicate name
	State    State  `gorm:"type:integer;check:(State >= 0) and (State <= 4)" json:"state"` // state of the user, are they online etc?
	Status   string `gorm:"type:varchar(50)" json:"status"`                                // the current status of the user, a small message
	About    string `gorm:"type:varchar(255)" json:"about"`                                // about the author, a larger body of text
	Picture  string `gorm:"type:text" json:"picture"`                                      // link to the image of the author
}
