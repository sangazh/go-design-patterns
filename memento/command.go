package memento

type Command interface {
	GetValue() interface{}
}

type Volume byte

func (v Volume) GetValue() interface{} {
	return v
}

type Mute bool

func (m Mute) GetValue() interface{} {
	return m
}

type Memento struct {
	memento Command
}

type Originator struct {
	Command Command
}

func (o *Originator) NewMemento() Memento {
	return Memento{o.Command}
}

func (o *Originator) ExtractAndStoreCommand(m Memento) {
	o.Command = m.memento
}

type CareTaker struct {
	mementoStack []Memento
}

func (c *CareTaker) Add(m Memento) {
	c.mementoStack = append(c.mementoStack, m)
}

func (c *CareTaker) Pop() Memento {
	if len(c.mementoStack) > 0 {
		tempMemento := c.mementoStack[len(c.mementoStack)-1]
		c.mementoStack = c.mementoStack[0 : len(c.mementoStack)-1]
		return tempMemento
	}
	return Memento{}
}

type MementoFacade struct {
	originator Originator
	careTaker  CareTaker
}

func (m *MementoFacade) SaveSettings(s Command) {
	m.originator.Command = s
	m.careTaker.Add(m.originator.NewMemento())
}

func (m *MementoFacade) RestoreSettings(i int) Command {
	m.originator.ExtractAndStoreCommand(m.careTaker.Pop())
	return m.originator.Command
}
