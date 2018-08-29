package statemachine

////////////////////////////////////////////
// Unique
////////////////////////////////////////////
type Unique interface {
	Id() uint
	Name() string
}

////////////////////////////////////////////
// TransitionResult
////////////////////////////////////////////
type TransitionResult struct {
	Ok bool
}

////////////////////////////////////////////
// TransitionHandler
////////////////////////////////////////////
type TransitionHandler func(args map[string]interface{}) (TransitionResult)

////////////////////////////////////////////
// State
////////////////////////////////////////////
type State interface {
	Unique
	Transitions() []Transition
}

////////////////////////////////////////////
// Transition
////////////////////////////////////////////
type Transition interface {
	Unique

	Destination() State
	Handle(map[string]interface{}) TransitionResult
}

////////////////////////////////////////////
// StateMachine
////////////////////////////////////////////
type StateMachine interface {
	Unique

	State() State
	Next(Transition)
}

////////////////////////////////////////////
// ImplState
////////////////////////////////////////////
type ImplState struct {
	id uint
	name string
	transitions []Transition
}

func (s ImplState) Id() uint {
	return s.id
}

func (s ImplState) Name() string {
	return s.name
}

func (s ImplState) Transitions() []Transition {
	return s.transitions
}

////////////////////////////////////////////
// ImplTransition
////////////////////////////////////////////
type ImplTransition struct {
	id uint
	name string
	
	destination State
	
	handler TransitionHandler
}

func (t ImplTransition) Id() uint {
	return t.id
}

func (t ImplTransition) Name() string {
	return t.name
}

func (t ImplTransition) Destination() State {
	return t.destination
}

func (t ImplTransition) Handle(args map[string]interface{}) TransitionResult {
	return t.handler(args)
}

////////////////////////////////////////////
// ImplStateMachine
////////////////////////////////////////////
type ImplStateMachine struct {
	id uint
	name string
	
	state State
	
	Args map[string]interface{}
}

func (sm *ImplStateMachine) Id() uint {
	return sm.id
}

func (sm *ImplStateMachine) Name() string {
	return sm.name
}

func (sm *ImplStateMachine) State() State {
	return sm.state
}

func (sm *ImplStateMachine) Next(t Transition) {
	sm.state = t.Destination()
}

func asyncLoadStates(st chan []ImplState){
	s1 := ImplState{1, "Start", []Transition{}}
	s2 := ImplState{2, "End", []Transition{}}

	st <- []ImplState{s1, s2}
}

func asyncLoadTransitions(tr chan []ImplTransition) {
	t1 := ImplTransition{1, "Complete", ImplState{0, "", []Transition{}}, func(args map[string]interface{}) TransitionResult {
		return TransitionResult{Ok: true}
	}}

	tr <- []ImplTransition{t1}
}

func Load() ImplStateMachine {
	stChan := make(chan []ImplState)
	trChan := make(chan []ImplTransition)

	go asyncLoadStates(stChan)
	go asyncLoadTransitions(trChan)

	states, transitions := <-stChan, <-trChan

	// set all destinations
	transitions[0].destination = states[1]

	// append all transitions
	states[0].transitions = append(states[0].transitions, transitions[0])
	
	return ImplStateMachine{1, "Main", states[0], map[string]interface{} {"charge_id": 1}}
}
