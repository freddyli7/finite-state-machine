package main

import "fmt"

func changeColor() bool {
	fmt.Println("color changing")
	return true
}

func startFlashing() bool {
	fmt.Println("light is flashing")
	return true
}

func goOff() bool {
	fmt.Println("gonna turn off")
	return true
}

func goOn() bool {
	fmt.Println("gonna turn on")
	return true
}

func runColorfulLight() {
	machine := &Machine{
		id:           "M-colorful-light",
		InitialState: "off",
		States: StateMap{
			"off": MachineState{
				On: TransitionMap{
					"HOLD": MachineTransition{
						To:        "red",
						Operation: goOn,
					},
				},
			},
			"red": MachineState{
				On: TransitionMap{
					"TOGGLE_ONCE": MachineTransition{
						To:        "green",
						Operation: changeColor,
					},
					"TOGGLE_TWICE": MachineTransition{
						To:        "red",
						Operation: startFlashing,
					},
					"HOLD": MachineTransition{
						To:        "off",
						Operation: goOff,
					},
				},
			},
			"green": MachineState{
				On: TransitionMap{
					"TOGGLE_ONCE": MachineTransition{
						To:        "blue",
						Operation: changeColor,
					},
					"TOGGLE_TWICE": MachineTransition{
						To:        "green",
						Operation: startFlashing,
					},
					"HOLD": MachineTransition{
						To:        "off",
						Operation: goOff,
					},
				},
			},
			"blue": MachineState{
				On: TransitionMap{
					"TOGGLE_ONCE": MachineTransition{
						To:        "yellow",
						Operation: changeColor,
					},
					"TOGGLE_TWICE": MachineTransition{
						To:        "blue",
						Operation: startFlashing,
					},
					"HOLD": MachineTransition{
						To:        "off",
						Operation: goOff,
					},
				},
			},
			"yellow": MachineState{
				On: TransitionMap{
					"TOGGLE_ONCE": MachineTransition{
						To:        "red",
						Operation: changeColor,
					},
					"TOGGLE_TWICE": MachineTransition{
						To:        "yellow",
						Operation: startFlashing,
					},
					"HOLD": MachineTransition{
						To:        "off",
						Operation: goOff,
					},
				},
			},
		},
	}
	fmt.Println("ID: ", machine.GetID())
	output := machine.Transition("HOLD")
	fmt.Println(output)
	output = machine.Transition("TOGGLE_ONCE")
	fmt.Println(output)
	output = machine.Transition("TOGGLE_ONCE")
	fmt.Println(output)
	output = machine.Transition("TOGGLE_ONCE")
	fmt.Println(output)
	output = machine.Transition("TOGGLE_TWICE")
	fmt.Println(output)
	output = machine.Transition("HOLD")
	fmt.Println(output)
	output = machine.Transition("HOLD")
	fmt.Println(output)
}
