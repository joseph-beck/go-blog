package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	f := Login()
	assert.NotNil(t, f)
}
