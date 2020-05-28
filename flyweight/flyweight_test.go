package flyweight

import (
	"fmt"
	"testing"
)

func TestTeamFlyweightFactory_GetTeam(t *testing.T) {
	factory := NewTeamFactory()
	teamA1 := factory.GetTeam(TeamA)
	if teamA1 == nil {
		t.Error("teamA is nil")
	}

	teamA2 := factory.GetTeam(TeamA)
	if teamA2 == nil {
		t.Error("teamA is nil")
	}

	if teamA1 != teamA2 {
		t.Error("teamA pointers should be the same")
	}

	if factory.GetNumberOfObjects() != 1 {
		t.Errorf("number of objectes should be 1, got: %d", factory.GetNumberOfObjects())
	}

}

func Test_HighVolume(t *testing.T) {
	factory := NewTeamFactory()
	teams := make([]*Team, 500000*2)
	for i := 0; i < 500000; i++ {
		teams[i] = factory.GetTeam(TeamA)
	}
	for i := 500000; i < 2*500000; i++ {
		teams[i] = factory.GetTeam(TeamB)
	}

	if factory.GetNumberOfObjects() != 2 {
		t.Errorf("should be 2, got: %d", factory.GetNumberOfObjects())
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("pointer %d points to %p and is located in %p\n", i, teams[i], &teams[i])
	}
}
