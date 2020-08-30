package product

import (
	"fmt"
	"strings"
	"time"
)

// Model of product
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Para mostrar mas ordenada la info de la BD
func (m *Model) String() string {
	return fmt.Sprintf(
		"%02d | %-50s | %-20s | %5d | %10s | %10s",
		m.ID, m.Name, m.Observations, m.Price,
		m.CreatedAt.Format("2006-01-02"),
		m.UpdatedAt.Format("2006-01-02"))
}

// Models slice of Model
type Models []*Model

func (m Models) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%02s | %-50s | %-20s | %5s | %10s | %10s\n",
		"id", "name", "observations", "price", "created_at", "updated_at"))
	for _, model := range m {
		builder.WriteString(model.String() + "\n")
	}
	return builder.String()
}

// Storage of Product
type Storage interface {
	Migrate() error
	Create(*Model) error
	GetAll() (Models, error)
	GetByID(uint) (*Model, error)
}

// Service of Product
type Service struct {
	storage Storage
}

// NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used for migrate product
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}

// Create is used for create a product
func (s *Service) Create(m *Model) error {
	m.CreatedAt = time.Now()
	return s.storage.Create(m)
}

// GetAll is used for get all the products
func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}

// GetById is used for get a product
func (s *Service) GetById(id uint) (*Model, error) {
	return s.storage.GetByID(id)
}
