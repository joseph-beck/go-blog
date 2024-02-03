package cache

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	v := m.Run()
	os.Exit(v)
}

func TestNew(t *testing.T) {
	c := New()
	assert.NotNil(t, c)
}

func TestPing(t *testing.T) {
	c := New()

	_, err := c.Ping()
	assert.NoError(t, err)
}

func TestClose(t *testing.T) {
	c := New()

	err := c.Close()
	assert.NoError(t, err)
}

func TestSet(t *testing.T) {

}

func TestGet(t *testing.T) {

}

func TestDelete(t *testing.T) {

}

func TestContains(t *testing.T) {

}
