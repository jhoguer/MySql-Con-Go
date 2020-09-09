package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jhoguer/MySql-Con-Go/pkg/product"
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

// Driver of storage
type Driver string

// Drivers
const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

// New created the connection with DB
func New(d Driver) {
	switch d {
	case MySQL:
		newMySQLDB()
	case Postgres:
		newPostgresDB()
	}
}

// newMySQLDB create new connection to mySQL db
func newMySQLDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", "root:toor@tcp(localhost:3306)/godb?parseTime=true")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Println("conectado a mySQL")
	})
}

// newPostgresDB create new connection to mySQL db
func newPostgresDB() {
	once.Do(func() {

		// Lo que este aqui solo se ejecutara 1 vez
		var err error
		db, err = sql.Open("postgres", "postgres://jhoguer:jhon198615@localhost:5432/godb?sslmode=disable")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Println("Conectado a Postgres")
	})
}

// Pool return a unique instace of db
func Pool() *sql.DB {
	return db
}

// Controlando los nulos string
func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

// Controlando los nulos Time
func timeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{Time: t}
	if !null.Time.IsZero() {
		null.Valid = true
	}
	return null
}

// DAOProduct factory of product.Storage
func DAOProduct(driver Driver) (product.Storage, error) {
	switch driver {
	case Postgres:
		return newPsqlProduct(db), nil
	case MySQL:
		return newMySQLProduct(db), nil
	default:
		return nil, fmt.Errorf("Driver not implement")
	}
}
