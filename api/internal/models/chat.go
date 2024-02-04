package models

import (
	"database/sql/driver"
	"errors"
	"strings"
)

type Users []string

func (u *Users) Scan(src any) error {
	str, ok := src.(string)
	if !ok {
		return errors.New("src value cannot cast to string")
	}

	*u = strings.Split(string(str), ",")

	return nil

}

func (u Users) Value() (driver.Value, error) {
	if len(u) == 0 {
		return nil, nil
	}

	return strings.Join(u, ","), nil
}

type Chat struct {
	Model
	Users Users
}
