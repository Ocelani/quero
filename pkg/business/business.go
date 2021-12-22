package business

import (
	"quero2pay/pkg/address"
	"time"
)

// Business represents the business enterprise entity type.
type Business struct {
	ID    int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name,omitempty"`  // Nome (Limite de 200 caracteres)
	Phone string `json:"phone,omitempty"` // Telefone, com máscara e validação que pode aceitar a quantidade de dígitos de celulares e telefones fixos, com DDD.

	Address *address.Address `json:"address,omitempty" gorm:"foreignkey:business_id;constraint:OnUpdate:CASCADE"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// NewBusiness instantiates a new business entity.
func NewBusiness(name, phone string, addr *address.Address) *Business {
	return &Business{
		Name:    name,
		Phone:   phone,
		Address: addr,
	}
}
