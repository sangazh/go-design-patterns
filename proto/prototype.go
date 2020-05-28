package proto

import (
	"errors"
	"fmt"
)

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

const (
	White = iota + 1
	Black
	Blue
)

func GetShirtsCloner() ShirtCloner {
	return new(ShirtsCache)
}

type ShirtsCache struct{}

func (s *ShirtsCache) GetClone(c int) (ItemInfoGetter, error) {
	switch c {
	case White:
		item := *whitePrototype
		return &item, nil
	}

	return nil, errors.New("Not implemented yet")
}

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtColor byte

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Color:%v, Price:%.2f, SKU: %s", s.Color, s.GetPrice(), s.SKU)
}

var whitePrototype = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

func (s *Shirt) GetPrice() float32 {
	return s.Price
}
