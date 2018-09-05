package StateMachine

func asyncLoadStates(st chan []UniqueState){
	s1 := UniqueState{ImplUnique{1, "Start"}, ImplState{[]Transition{}}}
	s2 := UniqueState{ImplUnique{2, "End"}, ImplState{[]Transition{}}}

	st <- []UniqueState{s1, s2}
}

func asyncLoadTransitions(tr chan []UniqueTransition) {
	t1 := UniqueTransition{ImplUnique{1, "Complete"}, ImplTransition{nil, func(args map[string]interface{}) TransitionResult {
		return TransitionResult{Ok: true}
	}}}

	tr <- []UniqueTransition{t1}
}

func Load() UniqueStateMachine {
	stChan := make(chan []UniqueState)
	trChan := make(chan []UniqueTransition)

	go asyncLoadStates(stChan)
	go asyncLoadTransitions(trChan)

	states, transitions := <-stChan, <-trChan

	// set all destinations
	(&transitions[0]).destination = &states[1]

	// append all transitions
	(&states[0]).transitions = append((&states[0]).transitions, transitions[0])
	
	return UniqueStateMachine{ImplUnique{1, "Main"}, ImplStateMachine{states[0], map[string]interface{} {}}}
}