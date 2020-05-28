package decorator

import (
	"strings"
	"testing"
)

func TestPizzaDecorator_AddIngredient(t *testing.T) {
	pizza := new(PizzaDecorator)
	pizzaResult, err := pizza.AddIngredient()
	expectedText := "Pizza with the following ingredients:"
	if err != nil {
		t.Error(err.Error())
		return
	}
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("got: %s, expect: %s", pizzaResult, expectedText)
	}

}


func TestOnion_AddIngredient(t *testing.T) {
	onion := new(Onion)
	onionResult, err := onion.AddIngredient()
	if err == nil {
		t.Errorf("expect error, result: %s", onionResult)
	}

	onion = &Onion{&PizzaDecorator{}}
	onionResult, err = onion.AddIngredient()
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(onionResult, "onion") {
		t.Errorf("should contain 'onion', but got: %s", onionResult)
	}
}

func TestMeat_AddIngredient(t *testing.T) {
	meat := new(Meat)
	meatResult, err := meat.AddIngredient()
	if err == nil {
		t.Errorf("expect error, result: %s", meatResult)
	}

	meat = &Meat{&PizzaDecorator{}}
	meatResult, err = meat.AddIngredient()
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(meatResult, "meat") {
		t.Errorf("should contain 'meat', but got: %s", meatResult)
	}
}

func TestPizzaDecorator_FullStack(t *testing.T) {
	pizza := &Onion{&Meat{&PizzaDecorator{}}}

	pizzaResult, err := pizza.AddIngredient()
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(pizzaResult, "Pizza") {
		t.Errorf("should contain 'pizza', but got: %s", pizzaResult)
	}
}
