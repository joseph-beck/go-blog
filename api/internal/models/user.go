package models

type Email string

type State int

const (
	Unknown State = iota // 0
	Online               // 1
	Away                 // 2
	Busy                 // 3
	Offline              // 4
)

type Role int

const (
	Guest   Role = iota // 0
	Blogger             // 1
	Admin               // 2
	Owner               // 3
)

type User struct {
	Model
	Username string `gorm:"type:varchar(25);unique;not null" json:"username"`                       // username of the user, must be unique
	Name     string `gorm:"type:varchar(25);not null" json:"name"`                                  // name of the user, can be a duplicate name
	Email    Email  `gorm:"type:varchar(100);unique;not null;unique_index" json:"email"`            // email of user, may require more validation
	Password string `gorm:"type:varchar(64);not null" json:"password"`                              // stored with a SHA256
	State    State  `gorm:"type:integer;not null;check:(State >= 0) and (State <= 4)" json:"state"` // state of the user, are they online etc?
	Role     Role   `gorm:"type:integer;not null;check:(Role >= 0) and (Role <= 3)" json:"role"`    // role of the user, used for user access levels
	Status   string `gorm:"type:varchar(50)" json:"status"`                                         // the current status of the user, a small message
	About    string `gorm:"type:varchar(255)" json:"about"`                                         // about the user, a larger body of text
}
