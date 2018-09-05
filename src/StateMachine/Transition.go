package StateMachine

////////////////////////////////////////////
// TransitionResult
////////////////////////////////////////////
type TransitionResult struct {
	Ok bool
	Msg string
}

////////////////////////////////////////////
// TransitionHandler
////////////////////////////////////////////
type TransitionHandler func(args map[string]interface{}) TransitionResult

////////////////////////////////////////////
// Transition
////////////////////////////////////////////
type Transition interface {
	Destination() State
	Handle(map[string]interface{}) TransitionResult
}

////////////////////////////////////////////
// ImplTransition
////////////////////////////////////////////
type ImplTransition struct {
	destination State
	handler TransitionHandler
}

func (this ImplTransition) Destination() State {
	return this.destination
}

func (this ImplTransition) Handle(args map[string]interface{}) TransitionResult {
	return this.handler(args)
}
