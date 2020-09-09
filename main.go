package main

import (
	"fmt"
	"log"

	"github.com/jhoguer/MySql-Con-Go/pkg/product"
	"github.com/jhoguer/MySql-Con-Go/storage"
)

func main() {
	// storage.NewMySQLDB()

	// storageProduct := storage.NewMySQLProduct(storage.Pool())
	// serviceProduct := product.NewService(storageProduct)

	// err := serviceProduct.Migrate()
	// if err != nil {
	// 	log.Fatalf("product.Migrate: %v", err)
	// }

	// storageInvoiceHeader := storage.NewMySQLInvoiceHeader(storage.Pool())
	// serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)

	// err = serviceInvoiceHeader.Migrate()
	// if err != nil {
	// 	log.Fatalf("invoiceheader.Migrate: %v", err)
	// }

	// storageInvoiceItem := storage.NewMySQLInvoiceItem(storage.Pool())
	// serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)

	// err = serviceInvoiceItem.Migrate()
	// if err != nil {
	// 	log.Fatalf("invoiceitem.Migrate: %v", err)
	// }

	//
	// m := &product.Model{
	// 	Name:         "Curso Html",
	// 	Price:        25,
	// 	Observations: "Edicion 2020",
	// }

	// err = serviceProduct.Create(m)
	// if err != nil {
	// 	log.Fatalf("product.Create: %v", err)
	// }

	// res, err := serviceProduct.GetAll()
	// if err != nil {
	// 	log.Fatalf("product.GetAll: %v", err)
	// }

	// fmt.Println(res)

	// GetById return a pruduct by id
	// m, err := serviceProduct.GetById(2)
	// switch {
	// case errors.Is(err, sql.ErrNoRows):
	// 	fmt.Println("No hay un producto con ese id")
	// case err != nil:
	// 	log.Fatalf("product.GetByID: %v", err)
	// default:
	// 	fmt.Println(m)
	// }

	// k := &product.Model{
	// 	ID:    1,
	// 	Name:  "Curso de CSS",
	// 	Price: 40,
	// }
	// err = serviceProduct.Update(k)
	// if err != nil {
	// 	log.Fatalf("product.Update: %v", err)
	// }

	// err = serviceProduct.Delete(3)
	// if err != nil {
	// 	log.Fatalf("product.Delete: %v", err)
	// }

	// Transactions
	// storageHeader := storage.NewMySQLInvoiceHeader(storage.Pool())
	// storageItems := storage.NewMySQLInvoiceItem(storage.Pool())

	// storageInvoice := storage.NewMySQLInvoice(storage.Pool(), storageInvoiceHeader, storageInvoiceItem)

	// j := &invoice.Model{
	// 	Header: &invoiceheader.Model{
	// 		Client: "Bety",
	// 	},
	// 	Items: invoiceitem.Models{
	// 		&invoiceitem.Model{ProductID: 2},
	// 		&invoiceitem.Model{ProductID: 5},
	// 		&invoiceitem.Model{ProductID: 6},
	// 	},
	// }
	// serviceInvoice := invoice.NewService(storageInvoice)
	// if err := serviceInvoice.Create(j); err != nil {
	// 	log.Fatalf("invoice.Create: %v", err)
	// }

	driver := storage.MySQL

	storage.New(driver)

	myStorage, err := storage.DAOProduct(driver)
	if err != nil {
		log.Fatalf("DAOProduct: %v", err)
	}

	serviceProduct := product.NewService(myStorage)

	ms, err := serviceProduct.GetById(5)
	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}

	fmt.Println(ms)

}
