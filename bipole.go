package main

type Bipole interface {
	Component
	SetNodes(nFrom *Node, nTo *Node)
	GetNodes() (A *Node, B *Node)
}

type Resistance float64

type Resistor struct {
	nodeFrom   *Node
	nodeTo     *Node
	voltage    Voltage
	current    Current
	Resistance Resistance
}

func NewResistor(r Resistance) *Resistor {
	return &Resistor{
		nodeFrom:   nil,
		nodeTo:     nil,
		Resistance: r,
	}
}

func (r *Resistor) GetVoltage() Voltage {
	return r.voltage
}

func (r *Resistor) GetCurrent() Current {
	return r.current
}

func (r *Resistor) GetNodes() (A *Node, B *Node) {
	return r.nodeFrom, r.nodeTo
}

func (r *Resistor) SetNodes(nFrom *Node, nTo *Node) {
	r.nodeFrom = nFrom
	r.nodeTo = nTo
}
