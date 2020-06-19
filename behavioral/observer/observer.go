package observer

import (
	"fmt"
)

type Observer interface {
	Notify(string)
	Message() string
}

type Publisher struct {
	ObserverList []Observer
}

func (p *Publisher) AddObserver(o Observer) {
	p.ObserverList = append(p.ObserverList, o)
}

func (p *Publisher) RemoveObserver(o Observer) {
	for i, observer := range p.ObserverList {
		if o == observer {
			p.ObserverList = append(p.ObserverList[:i], p.ObserverList[i+1:]...)
			break
		}
	}

}

func (p *Publisher) NotifyObserver(m string) {
	fmt.Printf("Publisher received message '%s' to notify observers\n", m)
	for _, o := range p.ObserverList {
		o.Notify(m)
	}
}

type TestObserver struct {
	ID      int
	message string
}

func (p *TestObserver) Notify(m string) {
	fmt.Printf("Ovserver %d: message '%s' received \n", p.ID, m)
	p.message = m
}

func (p *TestObserver) Message() string {
	return p.message
}
