package address

import (
	"time"
)

// Address represents a complete address
// information data based on its postal code (CEP).
type Address struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	ZipCode  string `json:"zip_code,omitempty"`
	City     string `json:"city,omitempty"`
	State    string `json:"state,omitempty"`
	District string `json:"district,omitempty"`
	Street   string `json:"street,omitempty"`
	Number   int    `json:"number,omitempty"`
	DDD      string `json:"ddd,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	BusinessID int `json:"business_id,omitempty"`
}

// NewAddress instantiates a new adress entity.
func NewAddress(zipCode string) *Address {
	return &Address{
		ZipCode: zipCode,
	}
}
