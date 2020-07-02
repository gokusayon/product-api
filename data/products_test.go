package dataimport

import "testing"

func TestCheckValidation(t *testing.T){
	p := &Product{
		Name: "hello",
		Price: 10,
		SKU: "abs-sadf-a",
	}

	err := p.Validate()
	if err != nil{
		t.Fatal(err)
	}
}


