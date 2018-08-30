package StateMachine

////////////////////////////////////////////
// State
////////////////////////////////////////////
type State struct {
	id uint
	name string

	Transitions []Transition
}

func (s State) Id() uint {
	return s.id
}

func (s State) Name() string {
	return s.name
}
