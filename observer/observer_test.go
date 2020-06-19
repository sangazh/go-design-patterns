package observer

import (
	"testing"
)

func TestTestObserver_Notify(t *testing.T) {
	o1 := &TestObserver{ID: 1}
	o2 := &TestObserver{ID: 2}
	o3 := &TestObserver{ID: 3}
	o4 := &TestObserver{ID: 4}
	p := new(Publisher)

	t.Run("AddObserver", func(t *testing.T) {
		p.AddObserver(o1)
		p.AddObserver(o2)
		p.AddObserver(o3)
		p.AddObserver(o4)
		if len(p.ObserverList) != 4 {
			t.Fail()
		}
	})

	t.Run("RemoveObserver", func(t *testing.T) {
		p.RemoveObserver(o2)

		if len(p.ObserverList) != 3 {
			t.Errorf("should be three left, got: %d", len(p.ObserverList))
		}

		for _, o := range p.ObserverList {
			to, ok := o.(*TestObserver)
			if !ok {
				t.Fail()
			}
			if to.ID == 2 {
				t.Fail()
			}
		}
	})

	t.Run("Notify", func(t *testing.T) {
		if len(p.ObserverList) == 0 {
			t.Error("nothing to test. empty list")
		}
		msg := "Hello World"
		p.NotifyObserver(msg)
		for _, o := range p.ObserverList {
			t.Log(o.Message())
		}
	})
}
