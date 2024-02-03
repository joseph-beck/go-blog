package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPgConn(t *testing.T) {
	conn := PgConn()
	assert.NotNil(t, conn)
}

func TestSQLiteConn(t *testing.T) {
	conn := SQLiteConn()
	assert.NotNil(t, conn)
}
