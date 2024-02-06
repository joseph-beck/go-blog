package models

type State int

const (
	Unknown State = iota // 0
	Online               // 1
	Away                 // 2
	Busy                 // 3
	Offline              // 4
)

type User struct {
	Model
	Username string `gorm:"type:varchar(25);unique;not null" json:"username"`
	Email    string `gorm:"type:varchar(100);unique;not null;unique_index" json:"email"`
	Password string `gorm:"type:varchar(64);not null" json:"password"`
	State    State  `gorm:"type:integer;not null;check:(state >= 0) and (state <= 4)" json:"state"`
	Status   string `gorm:"type:varchar(50)" json:"status"`
	About    string `gorm:"type:varchar(255)" json:"about"`
}
