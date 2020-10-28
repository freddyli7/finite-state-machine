package main

import "fmt"

func onOper() bool {
	fmt.Println("gonna turn off")
	return true
}

func offOper() bool {
	fmt.Println("gonna turn on")
	return true
}

func runSimpleStateMachine() {
	machine := &Machine{
		id:           "M-simple-light",
		InitialState: "on",
		States: StateMap{
			"on": MachineState{
				On: TransitionMap{
					"TOGGLE": MachineTransition{
						To:        "off",
						Operation: onOper,
					},
					"DOUBLE_TOGGLE": MachineTransition{
						To:        "on",
						Operation: offOper,
					},
				},
			},
			"off": MachineState{
				On: TransitionMap{
					"TOGGLE": MachineTransition{
						To:        "on",
						Operation: offOper,
					},
					"DOUBLE_TOGGLE": MachineTransition{
						To:        "off",
						Operation: onOper,
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
