package data

import "testing"

func TestValidation(t *testing.T) {
	p := &Product{
		Name:  "Cofeee",
		Price: 1,
		SKU:   "aa-avb-aa",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
