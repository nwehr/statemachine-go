package main

import "fmt"
import "StateMachine"

func main() {
	sm := StateMachine.Load()
	
	fmt.Println(sm.State.Name())
	
	sm.Next(sm.State.Transitions[0])
	
	fmt.Println(sm.State.Name())
}
