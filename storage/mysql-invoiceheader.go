package storage

import (
	"database/sql"
	"fmt"
)

const (
	mySQLMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
		id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
		client VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP
	)`
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
