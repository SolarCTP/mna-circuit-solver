package main

type Voltage float64
type Current float64

// The component is logically the edge of the graph representing the entire
// circuit
type Component interface {
	GetVoltage() Voltage
	GetCurrent() Current
}
