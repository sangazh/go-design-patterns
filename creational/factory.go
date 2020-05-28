package creational

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	Cash      = 1
	DebitCard = 2
)

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m{
	case Cash:
		return new(CardPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	}
	return nil, errors.New("not implemented yet")
}

type CardPM struct{}
type DebitCardPM struct{}
type CreditCardPM struct{}

func (c *CardPM) Pay(amount float32) string{
	return fmt.Sprintf("%.02f paid using cash", amount)
}

func (c *DebitCardPM) Pay(amount float32) string{
	return fmt.Sprintf("%.02f paid using debit card", amount)
}

func (c *CreditCardPM) Pay(amount float32) string{
	return fmt.Sprintf("%.02f paid using credit card", amount)
}
