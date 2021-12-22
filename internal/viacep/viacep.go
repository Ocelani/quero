package viacep

import (
	"encoding/json"
	"net/http"
	"quero2pay/pkg/address"
)

// AddressViaCEP represents a complete address
// information data received from ViaCEP API.
type AddressViaCEP struct {
	ZipCode  string `json:"cep"`
	City     string `json:"localidade"`
	State    string `json:"uf"`
	District string `json:"bairro"`
	Street   string `json:"logradouro"`
	Number   string `json:"complemento"`
	DDD      string `json:"ddd"`
}

// NewAddressViaCEP instantiates a new adress entity.
// It receives a CEP and returns the complete address
// based on the information avaiable on ViaCEP API.
func NewAddressViaCEP(cep string) (*AddressViaCEP, error) {
	url := "https://viacep.com.br/ws/" + cep + "/json/"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var addr AddressViaCEP
	err = json.NewDecoder(resp.Body).Decode(&addr)

	return &addr, err
}

// ToAddress converts the instantiated AddressViaCEP to an address.Address entity.
func (a *AddressViaCEP) ConvertToAddress() *address.Address {
	return &address.Address{
		ZipCode:  a.ZipCode,
		City:     a.City,
		State:    a.State,
		District: a.District,
		Street:   a.Street,
		DDD:      a.DDD,
	}
}
