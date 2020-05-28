package composite

import (
	"testing"
)

func TestGetParentField(t *testing.T) {
	p := new(Parent)
	GetParentField(p)

	s := new(Son)
	GetParentField(&s.p)
}
