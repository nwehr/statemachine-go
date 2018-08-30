package StateMachine

func asyncLoadStates(st chan []State){
	s1 := State{1, "Start", []Transition{}}
	s2 := State{2, "End", []Transition{}}

	st <- []State{s1, s2}
}

func asyncLoadTransitions(tr chan []Transition) {
	t1 := Transition{1, "Complete", State{0, "", []Transition{}}, func(args map[string]interface{}) TransitionResult {
		return TransitionResult{Ok: true}
	}}

	tr <- []Transition{t1}
}

func Load() StateMachine {
	stChan := make(chan []State)
	trChan := make(chan []Transition)

	go asyncLoadStates(stChan)
	go asyncLoadTransitions(trChan)

	states, transitions := <-stChan, <-trChan

	// set all destinations
	transitions[0].Destination = states[1]

	// append all transitions
	states[0].Transitions = append(states[0].Transitions, transitions[0])
	
	return StateMachine{1, "Main", states[0], map[string]interface{} {}}
}