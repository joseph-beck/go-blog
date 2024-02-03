package database

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/joseph-beck/chit-chat/api/internal/models"
	"github.com/stretchr/testify/assert"
)

type MockModel struct {
	models.Model
	Key   string `gorm:"type:text" json:"key"`
	Value string `gorm:"type:text" json:"value"`
}

func before() error {
	s := New(SQLiteConn())
	err := s.AutoMigrate(&MockModel{})
	return err
}

func after() error {
	s := New(SQLiteConn())
	err := s.DropTable("mock_models")
	return err
}

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	err = before()
	if err != nil {
		panic(err)
	}

	v := m.Run()

	err = after()
	if err != nil {
		panic(err)
	}

	os.Exit(v)
}

func TestNew(t *testing.T) {
	s := New(SQLiteConn())
	assert.NotNil(t, s)
}

func TestClose(t *testing.T) {
	s := New(SQLiteConn())
	err := s.Close()
	assert.NoError(t, err)
}

func TestPing(t *testing.T) {
	s := New(SQLiteConn())
	err := s.Ping()
	assert.NoError(t, err)
}

func TestList(t *testing.T) {
	s := New(SQLiteConn())

	var m []MockModel
	err := s.List(&m, "mock_models")
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	s := New(SQLiteConn())

	m := MockModel{Model: models.Model{}, Key: "david2001", Value: "david@gmail.com"}
	err := s.Create(&m, "mock_models")
	assert.NoError(t, err)

	m = MockModel{Model: models.Model{ID: 1}}
	err = s.Get(&m, "mock_models", "id = ?", 1)
	assert.NoError(t, err)
}

func TestCreate(t *testing.T) {
	s := New(SQLiteConn())

	m := MockModel{Model: models.Model{}, Key: "david2001", Value: "david@gmail.com"}
	err := s.Create(&m, "mock_models")
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	s := New(SQLiteConn())

	m := MockModel{Model: models.Model{}, Key: "david2001", Value: "david@gmail.com"}
	err := s.Create(&m, "mock_models")
	assert.NoError(t, err)

	o := m
	m.Key = "david2002"
	err = s.Update(&m, "mock_models")
	assert.NoError(t, err)
	assert.NotEqual(t, o.Key, m.Key)
}

func TestDelete(t *testing.T) {
	s := New(SQLiteConn())

	m := MockModel{Model: models.Model{}, Key: "david2001", Value: "david@gmail.com"}
	err := s.Create(&m, "mock_models")
	assert.NoError(t, err)

	err = s.Delete(&MockModel{}, "mock_models", "id = ?", m.ID)
	assert.NoError(t, err)
}

func TestContains(t *testing.T) {
	s := New(SQLiteConn())

	e, err := s.Contains(&MockModel{}, "mock_models", "id = ?", uint(500))
	assert.NoError(t, err)
	assert.False(t, e)
}

type TestTable struct {
	models.Model
	Key   string `gorm:"type:text" json:"key"`
	Value string `gorm:"type:text" json:"value"`
}

func TestAutoMigrate(t *testing.T) {
	s := New(SQLiteConn())

	err := s.AutoMigrate()
	assert.NoError(t, err)

	err = s.AutoMigrate(MockModel{})
	assert.NoError(t, err)
}

func TestCreateTable(t *testing.T) {
	s := New(SQLiteConn())

	err := s.CreateTable(&TestTable{})
	assert.NoError(t, err)

	e := s.HasTable("test_tables")
	assert.True(t, e)
}

func TestDropTable(t *testing.T) {
	s := New(SQLiteConn())

	err := s.DropTable("test_tables")
	assert.NoError(t, err)

	e := s.HasTable("test_tables")
	assert.False(t, e)
}

func TestHasTable(t *testing.T) {
	s := New(SQLiteConn())

	e := s.HasTable("test_tables")
	assert.False(t, e)
}
