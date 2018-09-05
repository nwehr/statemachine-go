package main

import _ "fmt"
import "StateMachine"

func main() {
	sm := StateMachine.Load()

	{
		unique, _ := sm.State().(StateMachine.Unique)
		println(unique.Name())
	}
	
	sm.Next(sm.State().Transitions()[0])
	
	{
		unique, _ := sm.State().(StateMachine.Unique)
		println(unique.Name())
	}
}
