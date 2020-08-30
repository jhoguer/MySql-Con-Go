package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jhoguer/MySql-Con-Go/pkg/invoiceheader"
	"github.com/jhoguer/MySql-Con-Go/pkg/invoiceitem"
	"github.com/jhoguer/MySql-Con-Go/pkg/product"
	"github.com/jhoguer/MySql-Con-Go/storage"
)

func main() {
	storage.NewMySQLDB()

	storageProduct := storage.NewMySQLProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	err := serviceProduct.Migrate()
	if err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}

	storageInvoiceHeader := storage.NewMySQLInvoiceHeader(storage.Pool())
	serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)

	err = serviceInvoiceHeader.Migrate()
	if err != nil {
		log.Fatalf("invoiceheader.Migrate: %v", err)
	}

	storageInvoiceItem := storage.NewMySQLInvoiceItem(storage.Pool())
	serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)

	err = serviceInvoiceItem.Migrate()
	if err != nil {
		log.Fatalf("invoiceitem.Migrate: %v", err)
	}

	//
	// m := &product.Model{
	// 	Name:         "Curso de JavaScript",
	// 	Price:        70,
	// 	Observations: "Este curso es basico",
	// }

	// err = serviceProduct.Create(m)
	// if err != nil {
	// 	log.Fatalf("product.Create: %v", err)
	// }

	res, err := serviceProduct.GetAll()
	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}

	fmt.Println(res)

	// GetById return a pruduct by id
	m, err := serviceProduct.GetById(2)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("No hay un producto con ese id")
	case err != nil:
		log.Fatalf("product.GetByID: %v", err)
	default:
		fmt.Println(m)
	}

}
