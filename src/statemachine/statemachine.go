package StateMachine

////////////////////////////////////////////
// StateMachine
////////////////////////////////////////////
type StateMachine interface {
	State() State
	Next(Transition)
}

////////////////////////////////////////////
// ImplStateMachine
////////////////////////////////////////////
type ImplStateMachine struct {
	state State
	args map[string]interface{}
}

func (this ImplStateMachine) State() State {
	return this.state
}

func (this *ImplStateMachine) Next(t Transition) {
	this.state = t.Destination()
}