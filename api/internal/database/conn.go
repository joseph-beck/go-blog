package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Database connection function, returns a pointer to a database
type Conn func() *gorm.DB

// Return a connection to a postgres database
func PgConn() Conn {
	return func() *gorm.DB {
		dsn := fmt.Sprintf(`
		host=%s 
		user=%s 
		password=%s 
		dbname=%s 
		port=%s 
		sslmode=disable 
		TimeZone=Europe/London`,
			os.Getenv("PG_ADDR"),
			os.Getenv("PG_USER"),
			os.Getenv("PG_PASS"),
			os.Getenv("PG_NAME"),
			os.Getenv("PG_PORT"),
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		return db
	}
}

// Return a connection using to an sqlite database
func SQLiteConn() Conn {
	return func() *gorm.DB {
		db, err := gorm.Open(sqlite.Open(os.Getenv("SQLITE_DB")), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		return db
	}
}
