package main

import "fmt"
import "statemachine"

func main() {
	sm := statemachine.Load()
	
	fmt.Println(sm.State().Id())
	
	sm.Next(sm.State().Transitions()[0])
	
	fmt.Println(sm.State().Id())
}
