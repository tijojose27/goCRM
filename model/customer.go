package model

import (
	"github.com/google/uuid"
)

// Customer struct definition
type Customer struct {
	ID        string
	Name      string
	Role      string
	Email     string
	Phone     string
	Contacted bool
}

// CreateSampleCustomers use this functino to initalize list
func CreateSampleCustomers() []Customer {
	customers := []Customer{
		{ID: uuid.New().String(), Name: "John Doe", Role: "Manager", Email: "john.doe@example.com", Phone: "555-1234", Contacted: true},
		{ID: uuid.New().String(), Name: "Jane Smith", Role: "Salesperson", Email: "jane.smith@example.com", Phone: "555-5678", Contacted: false},
		{ID: uuid.New().String(), Name: "Alice Johnson", Role: "Developer", Email: "alice.johnson@example.com", Phone: "555-9876", Contacted: true},
		{ID: uuid.New().String(), Name: "Bob Williams", Role: "CEO", Email: "bob.williams@example.com", Phone: "555-4321", Contacted: false},
		{ID: uuid.New().String(), Name: "Charlie Brown", Role: "Support", Email: "charlie.brown@example.com", Phone: "555-8765", Contacted: true},
		{ID: uuid.New().String(), Name: "Diana Prince", Role: "HR", Email: "diana.prince@example.com", Phone: "555-3456", Contacted: false},
		{ID: uuid.New().String(), Name: "Evan Thompson", Role: "Marketing", Email: "evan.thompson@example.com", Phone: "555-6543", Contacted: true},
		{ID: uuid.New().String(), Name: "Fiona White", Role: "Designer", Email: "fiona.white@example.com", Phone: "555-9087", Contacted: false},
		{ID: uuid.New().String(), Name: "George Green", Role: "Consultant", Email: "george.green@example.com", Phone: "555-2345", Contacted: true},
		{ID: uuid.New().String(), Name: "Hannah Black", Role: "Project Manager", Email: "hannah.black@example.com", Phone: "555-6789", Contacted: false},
	}
	return customers
}

func NewCustomer(name, role, email, phone string, contacted bool) Customer {
	return Customer{
		ID:        uuid.New().String(),
		Name:      name,
		Role:      role,
		Email:     email,
		Phone:     phone,
		Contacted: contacted,
	}
}
