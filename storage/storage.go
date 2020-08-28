package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

// NewMySQLDB create new connection to mySQL db
func NewMySQLDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", "root:toor@tcp(localhost:3306)/godb")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Println("conectado a mySQL")
	})
}
