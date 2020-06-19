package visitor

import (
	"fmt"
)

type ProductInfoRetriever interface {
	GetPrice() float64
	GetName() string
}

type visitor interface {
	Visit(ProductInfoRetriever)
}

type visitable interface {
	Accept(visitor)
}

type Product struct {
	Price float64
	Name  string
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) Accept(v visitor) {
	v.Visit(p)
}

type Rice struct {
	Product
}

type Pasta struct {
	Product
}

type PriceVisitor struct {
	Sum float64
}

func (v *PriceVisitor) Visit(p ProductInfoRetriever) {
	v.Sum += p.GetPrice()
}

type NamePrinter struct {
	ProductList string
}

func (n *NamePrinter) Visit(p ProductInfoRetriever) {
	n.ProductList = fmt.Sprintf("%s\n%s", p.GetName(), n.ProductList)
}

type Fridge struct{
	Product
}

func (f *Fridge) GetPrice() float64{
	return f.Product.Price + 20
}

func (f *Fridge) Accept(v visitor) {
	v.Visit(f)
}
