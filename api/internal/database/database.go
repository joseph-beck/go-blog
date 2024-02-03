package database

import (
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

func New(c Conn) *Store {
	return &Store{
		DB: c(),
	}
}

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

func (s *Store) List(i interface{}, t string) error {
	res := s.DB.Table(t).Find(i)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (s *Store) Get(i interface{}, t string, p ...interface{}) error {
	res := s.DB.Table(t).Find(i, p...).First(i)
	if res.Error != nil {
		return res.Error
	}

	return nil

}

func (s *Store) Create(i interface{}, t string) error {
	res := s.DB.Table(t).Create(i)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (s *Store) Update(i interface{}, t string) error {
	res := s.DB.Table(t).Save(i)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (s *Store) Delete(i interface{}, t string, p ...interface{}) error {
	res := s.DB.Table(t).Delete(i, p...)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (s *Store) Contains(i interface{}, t string, p ...interface{}) (bool, error) {
	cnt := int64(0)
	res := s.DB.Table(t).Find(&i, p...).Count(&cnt)
	if res.Error != nil {
		return false, res.Error
	}

	return cnt != 0, nil
}

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

func (s *Store) CreateTable(m interface{}) error {
	err := s.DB.Migrator().CreateTable(m)
	return err
}

func (s *Store) DropTable(t string) error {
	err := s.DB.Migrator().DropTable(t)
	return err
}

func (s *Store) HasTable(t string) bool {
	e := s.DB.Migrator().HasTable(t)
	return e
}
