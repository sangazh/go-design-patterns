package memento

import (
	"fmt"
	"testing"
)

func TestMementoFacade_SaveSettings(t *testing.T) {
	m := MementoFacade{}
	m.SaveSettings(Volume(4))
	m.SaveSettings(Mute(false))

	assertAndPrint(m.RestoreSettings(0))
	assertAndPrint(m.RestoreSettings(1))
}

func assertAndPrint(c Command) {
	switch cast := c.(type) {
	case Volume:
		fmt.Printf("Volume: \t%d\n", cast)
	case Mute:
		fmt.Printf("Mute: \t%t\n", cast)
	}
}
