package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Conn func() *gorm.DB

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
func SQLiteConn() Conn {
	return func() *gorm.DB {
		db, err := gorm.Open(sqlite.Open(os.Getenv("SQLITE_DB")), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
		if err != nil {
			panic(err)
		}

		return db
	}
}
