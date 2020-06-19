package composite

import (
	"testing"
)

func TestAthlete_Train(t *testing.T) {
	localSwim := Swim

	swimmer := CompositeSwimmerA{
		MySwim: localSwim,
	}
	swimmer.MyAthlete.Train()
	swimmer.MySwim()
}

func TestAnimal_Eat(t *testing.T) {
	fish := Shark{Swim: Swim}
	fish.Eat()
	fish.Swim()
}

func TestCompositeSwimmerB(t *testing.T) {
	swimmer := CompositeSwimmerB{
		Trainer: &Athlete{},
		Swimmer: &SwimmerImpl{},
	}
	Train()
	Swim()
}
