package invoice

import (
	"github.com/jhoguer/MySql-Con-Go/pkg/invoiceheader"
	"github.com/jhoguer/MySql-Con-Go/pkg/invoiceitem"
)

// Model of invoice
type Model struct {
	Header *invoiceheader.Model
	Items  invoiceitem.Models
}

// Storage interface that must implement a db storage
type Storage interface {
	Create(*Model) error
}

// Service of invoice
type Service struct {
	storage Storage
}

// NewService return a pinter of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Create a new invoice
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
