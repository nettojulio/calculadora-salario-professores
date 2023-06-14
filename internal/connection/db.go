package connection

import (
	"database/sql"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetDBInstance() *sql.DB {
	return connectDB()
}

func connectDB() *sql.DB {
	once.Do(func() {
		connection := "user=postgres dbname=postgres password=postgres host=localhost sslmode=disable"

		var err error
		db, err = sql.Open("postgres", connection)
		if err != nil {
			panic(err.Error())
		}
	})

	return db
}
