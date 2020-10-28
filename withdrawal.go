package main

import "fmt"

func write2UwDB() bool {
	result := true
	fmt.Println("operation: write 2 Uw DB: ", result)
	return result
}

func sendThirdPartyRequest() bool {
	result := true
	fmt.Println("operation: send payout request to third party service: ", result)
	return result
}

func updateUwDB() bool {
	result := true
	fmt.Println("operation: update Uw DB: ", result)
	return result
}

func queryTxResult() bool {
	result := true
	fmt.Println("operation: query third party payout tx: ", result)
	return result
}

func write2DepositRecordDB() bool {
	result := false
	fmt.Println("operation: write withdrawal record into DB: ", result)
	return result
}

func mintBack() bool {
	result := true
	fmt.Println("operation: mint spwn back: ", result)
	return result
}

func doNothing() bool {
	return true
}

func runWithdrawal() {
	machine := &Machine{
		id:           "M-withdrawal",
		InitialState: "on_hold",
		States: StateMap{
			"on_hold": MachineState{
				On: TransitionMap{
					"write_uw": MachineTransition{
						To:        "init_unconfirmed_withdrawal",
						Operation: write2UwDB,
					},
					"failed": MachineTransition{
						To:        "done",
						Operation: mintBack,
					},
				},
			},
			"init_unconfirmed_withdrawal": MachineState{
				On: TransitionMap{
					"send": MachineTransition{
						To:        "request_sent",
						Operation: sendThirdPartyRequest,
					},
					"failed": MachineTransition{
						To:        "done",
						Operation: mintBack,
					},
				},
			},
			"request_sent": MachineState{
				On: TransitionMap{
					"update_uw": MachineTransition{
						To:        "unconfirmed",
						Operation: updateUwDB,
					},
					"failed": MachineTransition{
						To:        "unconfirmed",
						Operation: doNothing,
					},
				},
			},
			"unconfirmed": MachineState{
				On: TransitionMap{
					"query": MachineTransition{
						To:        "confirmed",
						Operation: queryTxResult,
					},
				},
			},
			"confirmed": MachineState{
				On: TransitionMap{
					"finalize": MachineTransition{
						To:        "done",
						Operation: write2DepositRecordDB,
					},
					"failed": MachineTransition{
						To:        "done",
						Operation: doNothing,
					},
				},
			},
			"done": MachineState{
				On: TransitionMap{
				},
			},
		},
	}
	fmt.Println("ID: ", machine.GetID())
	output := machine.Transition("write_uw")
	fmt.Println("state: " + output)
	output = machine.Transition("send")
	fmt.Println("state: " + output)
	output = machine.Transition("update_uw")
	fmt.Println("state: " + output)
	output = machine.Transition("query")
	fmt.Println("state: " + output)
	output = machine.Transition("finalize")
	fmt.Println("state: " + output)
}
