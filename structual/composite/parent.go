package composite

import (
	"fmt"
)

type Parent struct{
	SomeField int
}

type Son struct{
	p Parent
}

func GetParentField(p *Parent) int {
	fmt.Println(p.SomeField)
	return p.SomeField
}
