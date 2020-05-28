package decorator

import (
	"errors"
	"fmt"
)

type IngredientAdd interface {
	AddIngredient() (string, error)
}

type PizzaDecorator struct {
	Ingredient IngredientAdd
}

func (p *PizzaDecorator) AddIngredient() (string, error) {
	return "Pizza with the following ingredients:", nil
}

type Meat struct {
	Ingredient IngredientAdd
}

func (m *Meat) AddIngredient() (string, error) {
	if m.Ingredient == nil {
		return "", errors.New("decorator is not init")
	}
	s, err := m.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s meat", s), nil
}

type Onion struct {
	Ingredient IngredientAdd
}

func (m *Onion) AddIngredient() (string, error) {
	if m.Ingredient == nil {
		return "", errors.New("decorator is not init")
	}
	s, err := m.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s onion", s), nil
}
