package main

import "fmt"

// MachineTransition transition map
type MachineTransition struct {
	To string // to next state
}

// TransitionMap map with transitions: key=>event, value=>MachineTransition
type TransitionMap map[string]MachineTransition

// MachineState is State of machine
type MachineState struct {
	On TransitionMap // means on the event, do MachineTransition
}

// StateMap maps state, key=>state name, MachineState=>MachineState
type StateMap map[string]MachineState

// Machine datatype
type Machine struct {
	id           string
	InitialState string
	currentState string
	States       StateMap
}

// IMachine machine interface
type IMachine interface {
	Transition() string
	Current() string
	GetID() string
}

// Current returns current state
func (m *Machine) Current() string {
	if m.currentState == "" {
		return m.InitialState
	}
	return m.currentState
}

// Transition transitions to next state
func (m *Machine) Transition(event string) string {
	current := m.Current()
	transitions := m.States[current].On
	next := transitions[event].To
	if next != "" {
		m.currentState = next
		return next
	}
	return current
}

func (m *Machine) GetID() string {
	return m.id
}

func run() {
	machine := &Machine{
		id:           "M-1",
		InitialState: "on",
		States: StateMap{
			"on": MachineState{
				On: TransitionMap{
					"TOGGLE": MachineTransition{
						To: "off",
					},
					"DOUBLE_TOGGLE": MachineTransition{
						To: "on",
					},
				},
			},
			"off": MachineState{
				On: TransitionMap{
					"TOGGLE": MachineTransition{
						To: "on",
					},
					"DOUBLE_TOGGLE": MachineTransition{
						To: "off",
					},
				},
			},
		},
	}
	fmt.Println("ID: ", machine.GetID())
	output := machine.Transition("TOGGLE")
	fmt.Println(output)
	output = machine.Transition("TOGGLE")
	fmt.Println(output)
	output = machine.Transition("TOGGLE")
	fmt.Println(output)
	output = machine.Transition("DOUBLE_TOGGLE")
	fmt.Println(output)
	output = machine.Transition("TOGGLE")
	fmt.Println(output)
	output = machine.Transition("DOUBLE_TOGGLE")
	fmt.Println(output)
}

