package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := New()
	assert.NotNil(t, s)
}
