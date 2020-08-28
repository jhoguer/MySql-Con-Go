package main

import (
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

	m := &product.Model{
		Name:  "Curso Go desde cero",
		Price: 60,
	}

	err = serviceProduct.Create(m)
	if err != nil {
		log.Fatalf("product.Create: %v", err)
	}

}
