package StateMachine

////////////////////////////////////////////
// StateMachine
////////////////////////////////////////////
type StateMachine struct {
	id uint
	name string
	
	State State
	Args map[string]interface{}
}

func (sm StateMachine) Id() uint {
	return sm.id
}

func (sm StateMachine) Name() string {
	return sm.name
}

func (sm *StateMachine) Next(t Transition) {
	sm.State = t.Destination
}
