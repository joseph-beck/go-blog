package database

import (
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

// Create a new data Store
func New(c Conn) *Store {
	return &Store{
		DB: c(),
	}
}

// Close the data Store
func (s *Store) Close() error {
	db, err := s.DB.DB()
	if err != nil {
		return err
	}

	err = db.Close()
	if err != nil {
		return err
	}

	return nil
}

// Ping the data store
func (s *Store) Ping() error {
	db, err := s.DB.DB()
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

// Get a list of given models from the given table, use a reference for i.
// Queries can be customized with the p, for example "id = ?", 1
func (s *Store) List(i interface{}, t string, p ...interface{}) error {
	res := s.DB.Table(t).Find(i, p...)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// Get a single instance of a given model from the given table, use a reference for i.
// Queries can be customized with p, for example "id = ?", 1
func (s *Store) Get(i interface{}, t string, p ...interface{}) error {
	res := s.DB.Table(t).Find(i, p...).First(i)
	if res.Error != nil {
		return res.Error
	}

	return nil

}

// Create a given model in the given table, use a reference for i
func (s *Store) Create(i interface{}, t string) error {
	res := s.DB.Table(t).Create(i)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// Update a given model in the given table, use a reference for i
func (s *Store) Update(i interface{}, t string) error {
	res := s.DB.Table(t).Save(i)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// Delete a given model in the given table, use a reference for i.
// Queries can be customized with p, for example "id = ?", 1
func (s *Store) Delete(i interface{}, t string, p ...interface{}) error {
	res := s.DB.Table(t).Delete(i, p...)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// Check if the database contains the given model in the given table, use a reference for i.
// Queries can be customized with p, for example "id = ?", 1
func (s *Store) Contains(i interface{}, t string, p ...interface{}) (bool, error) {
	cnt := int64(0)
	res := s.DB.Table(t).Find(&i, p...).Count(&cnt)
	if res.Error != nil {
		return false, res.Error
	}

	return cnt != 0, nil
}

// Automatically migrate all of the given models
func (s *Store) AutoMigrate(i ...interface{}) error {
	if len(i) <= 0 {
		return nil
	}

	for _, t := range i {
		err := s.DB.Migrator().AutoMigrate(t)
		if err != nil {
			return err
		}
	}

	return nil
}

// Create a table of the given model
func (s *Store) CreateTable(m interface{}) error {
	err := s.DB.Migrator().CreateTable(m)
	return err
}

// Drop the tabe of a given name
func (s *Store) DropTable(t string) error {
	err := s.DB.Migrator().DropTable(t)
	return err
}

// Check if a table of the given name exists
func (s *Store) HasTable(t string) bool {
	e := s.DB.Migrator().HasTable(t)
	return e
}
