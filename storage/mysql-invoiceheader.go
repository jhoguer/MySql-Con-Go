package storage

import (
	"database/sql"
	"fmt"

	"github.com/jhoguer/MySql-Con-Go/pkg/invoiceheader"
)

const (
	mySQLMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
		id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
		client VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP
	)`
	mySQLCreateInvoiceHeader = `INSERT INTO invoice_headers(client)
														VALUES (?)`
)

// MySQLInvoiceHeader used for work with postgres - invoiceHeader
type MySQLInvoiceHeader struct {
	db *sql.DB
}

// NewMySQLInvoiceHeader return a new pointer of PsqlInvoiceHeader
func NewMySQLInvoiceHeader(db *sql.DB) *MySQLInvoiceHeader {
	return &MySQLInvoiceHeader{db}
}

// Migrate implement the interface invoideHeader.Storage
func (p *MySQLInvoiceHeader) Migrate() error {
	stmt, err := p.db.Prepare(mySQLMigrateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migración de invoiceHeader ejecutada correctamente")
	return nil
}

// CreateTx interface that must implement a db storage
func (p *MySQLInvoiceHeader) CreateTx(tx *sql.Tx, m *invoiceheader.Model) error {
	stmt, err := tx.Prepare(mySQLCreateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(m.Client)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	m.ID = uint(id)

	return nil
}
