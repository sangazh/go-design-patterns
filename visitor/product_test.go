package visitor

import (
	"fmt"
	"testing"
)

func TestProduct_Accept(t *testing.T) {
	products := make([]visitable, 3)
	products[0] = &Rice{Product{
		Price: 32.0,
		Name:  "Some rice",
	}}
	products[1] = &Rice{Product{
		Price: 40,
		Name:  "Some pasta",
	}}
	products[2] = &Fridge{Product{
		Price: 50,
		Name:  "A fridge",
	}}

	PriceVisitor := new(PriceVisitor)

	for _, p := range products {
		p.Accept(PriceVisitor)
	}

	fmt.Printf("Total: %f\n", PriceVisitor.Sum)

	nameVisitor := new(NamePrinter)
	for _, p := range products {
		p.Accept(nameVisitor)
	}

	fmt.Printf("\nProduct list:\n -------\n%s", nameVisitor.ProductList)
}
