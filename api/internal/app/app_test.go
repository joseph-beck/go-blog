package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	a := New()
	assert.NotNil(t, a)
}
