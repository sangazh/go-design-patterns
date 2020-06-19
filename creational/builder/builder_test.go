package builder

import (
	"strings"
	"testing"

	"design/creational"
)

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("car wheel must be 4, got %d", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("car structure must be 'Car', got %s", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("car seats must be 5, got %d", car.Seats)
	}

	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	bike := bikeBuilder.GetVehicle()

	if bike.Wheels != 2 {
		t.Errorf("bike wheel must be 4, got %d", bike.Wheels)
	}

	if bike.Structure != "Motorbike" {
		t.Errorf("bike structure must be 'Car', got %s", bike.Structure)
	}

	if bike.Seats != 1 {
		t.Errorf("bike seats must be 5, got %d", bike.Seats)
	}
}

func TestGetPaymentMethodCash(t *testing.T) {
	payment, err := creational.GetPaymentMethod(creational.Cash)
	if err != nil {
		t.Fatal("A payment method of type 'Cash' must exist")
	}

	msg := creational.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("The cash payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodDebitCard(t *testing.T) {
	payment, err := creational.GetPaymentMethod(creational.DebitCard)
	if err != nil {
		t.Fatal("A payment method of type 'DebitCard' must exist")
	}

	msg := creational.Pay(22.30)
	if !strings.Contains(msg, "paid using debit card") {
		t.Error("The debit card payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethod(t *testing.T) {
	_, err := creational.GetPaymentMethod(20)
	if err == nil {
		t.Error("A payment method with ID 20 must return an error")
	}

	t.Log("LOG:", err)
}
