package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsersScan(t *testing.T) {
	var u Users
	err := u.Scan("1,2,3")
	assert.NoError(t, err)
	assert.Equal(t, Users{"1", "2", "3"}, u)

	err = u.Scan(123)
	assert.Error(t, err)
}

func TestUsersValue(t *testing.T) {
	u := Users{"1", "2", "3"}
	s, err := u.Value()
	assert.NoError(t, err)
	assert.Equal(t, "1,2,3", s)

	u = Users{}
	s, err = u.Value()
	assert.NoError(t, err)
	assert.Nil(t, s)
}
