package dataimport

import (
	"fmt"
	"time"
)

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// the ID of the product
	//
	// required: true
	// min: 1
	ID          int     `json:"id"`

	// the Name of the Product
	//
	// required: true
	// max length: 25
	Name        string  `json:"name" validate:"required"`


	// the Description of the Product
	//
	// required: false
	// max length: 10000
	Description string  `json:"description"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price       float32 `json:"price" validate:"gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU         string  `json:"sku" validate:"sku"`

	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

//func (p *Product) Validate() error{
//	validate := validator.New()
//	validate.RegisterValidation("sku", validateSKU)
//	return validate.Struct(p)
//}


// Products is a collection of Product
type Products []*Product

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product){
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error{
	_, pos, err := findNextProduct(id)

	if err != nil {
		return err
	}

	productList[pos] = p
	p.ID = id
	
	return nil
}

var ErrorProductNotFound error = fmt.Errorf("Product Not Found!")

func findNextProduct(id int) (*Product , int, error){
	for i, p := range productList {
		if p.ID ==  id{
			return p, i ,nil
		}
	}

	return nil, -1,  ErrorProductNotFound
}

func DeleteProduct(id int) (error){
	_, i, err := findNextProduct(id)

	if err != nil {
		return err
	}

	productList = append(productList[:i], productList[i+1:]...)
	return nil
}

func getNextID() int{
	lp := productList[len(productList) -1]
	return lp.ID + 1
}

func GetProductByID(id int) (*Product, error){
	pd, _, err := findNextProduct(id)

	if err != nil {
		return nil, err
	}
	return pd, nil
}

// productList is a hard coded list of products for this
// example data source
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
