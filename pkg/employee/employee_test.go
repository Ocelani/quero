package employee

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseSalary(t *testing.T) {
	tests := []struct {
		sallary string
		want    float64
	}{
		{sallary: "R$1.000,99", want: 1000.99},
		{sallary: "R$0,99", want: 0.99},
		{sallary: "R$10", want: 10},
		{sallary: "", want: 0},
		{sallary: "R$", want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.sallary, func(t *testing.T) {
			got := parseSallary(tt.sallary)
			assert.Equal(t, tt.want, got)
		})
	}
}
