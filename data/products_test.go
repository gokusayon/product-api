package dataimport

import (
	"testing"
)

func TestCheckValidation(t *testing.T){
	p := &Product{
		Name: "hello",
		Price: 10,
		SKU: "abs-sadf-a",
	}

	validate := NewValidation()
	err := validate.Validate(p)
	if err != nil{
		t.Fatal(err)
	}
}


