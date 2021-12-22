package employee

import (
	"quero2pay/pkg/business"
	"strconv"
	"strings"
	"time"
)

type Role string

const (
	Developer Role = "Developer"
	Designer  Role = "Designer"
	Admin     Role = "Admin"
)

// Employee represents an employee entity.
type Employee struct {
	ID      int     `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`
	Name    string  `json:"name,omitempty"` // Nome (até 200 caracteres)
	Sallary float64 `json:"sallary,omitempty"`
	Role    Role    `json:"role,omitempty"`

	BusinessID int
	Business   *business.Business `json:"business_id,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// NewEmployee instantiates a new employee entity.
func NewEmployee(name, sallary string, role Role, bsns *business.Business) *Employee {
	return &Employee{
		Name:       name,
		Sallary:    parseSallary(sallary),
		Role:       role,
		BusinessID: bsns.ID,
		Business:   bsns,
	}
}

func parseSallary(sallary string) float64 {
	s := strings.ReplaceAll(sallary, "R$", "")
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, ",", ".")
	slry, _ := strconv.ParseFloat(s, 64)
	return slry
}
