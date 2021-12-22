package viacep

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAddressViaCEP(t *testing.T) {
	tests := []struct {
		cep     string
		want    *AddressViaCEP
		wantErr assert.ErrorAssertionFunc
	}{
		{
			cep: "01001000",
			want: &AddressViaCEP{
				ZipCode:  "01001-000",
				Street:   "Praça da Sé",
				Number:   "lado ímpar",
				District: "Sé",
				City:     "São Paulo",
				State:    "SP",
				DDD:      "11",
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.cep, func(t *testing.T) {
			got, err := NewAddressViaCEP(tt.cep)
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
