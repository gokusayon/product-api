package main

import (
	"fmt"
	"gokusyon/github.com/products-api/sdk/client"
	"gokusyon/github.com/products-api/sdk/client/products"
	"testing"
)

func TestClientListProducts(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:8080")
	c := client.NewHTTPClientWithConfig(nil, cfg)
	params := products.NewListProductsParams()
	prod, err := c.Products.ListProducts(params)
	if err != nil{
		t.Fatal(err)
	}

	fmt.Println(prod.GetPayload()[0])
}

func TestClientSingleProduct(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:8080")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	param := products.NewListSingleProductParams().WithID(1)
	prod, err := c.Products.ListSingleProduct(param)
	if err != nil{
		t.Fatal(err)
	}

	fmt.Printf("%#v", prod.GetPayload())

}
