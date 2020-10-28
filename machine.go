package main

import "fmt"

// MachineTransition transition map
type MachineTransition struct {
	To        string // to next state
	Operation func() bool
}

func (m MachineTransition) isOperationExist() bool {
	return m.Operation != nil
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

	if !transitions[event].isOperationExist() {
		fmt.Println("Error: " + event + " is not allow in the current state")
		return current
	}
	operExecOk := transitions[event].Operation()
	if !operExecOk {
		if !transitions["failed"].isOperationExist() {
			return current
		}
		_ = transitions["failed"].Operation()
		t := transitions["failed"].To
		m.currentState = t
		return t
	}
	next := transitions[event].To
	m.currentState = next

	return next
}

func (m *Machine) GetID() string {
	return m.id
}
