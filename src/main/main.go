package main

import "fmt"

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

func load() StateMachine {
	defaultHandler := func(args map[string]interface{}) TransitionResult {
		return TransitionResult{Ok: true}
	}

	args := map[string]interface{} {"charge_id": 1}
	
	// create all states
	s1 := &ImplState{1, "Start", []Transition{}}
	s2 := &ImplState{2, "End", []Transition{}}
	
	// create all transitions
	t1 := &ImplTransition{1, "Complete", s2, defaultHandler}
	
	// add transitions to states
	s1.transitions = append(s1.transitions, t1)
	
	// create state machine with initial and current state
	return &ImplStateMachine{1, "Main", s1, args}
}

func main() {
	sm := load()
	
	fmt.Println(sm.State().Id())
	
	sm.Next(sm.State().Transitions()[0])
	
	fmt.Println(sm.State().Id())
}
