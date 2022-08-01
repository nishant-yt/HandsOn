/*
Basic Implementation of Register Callback and Events

Goal :  When the register event is triggered , the callback function should store those events in an buffer/ queue and
		when event is fired , all those functions from queue should be executed one by one

		In case any function is registered in the buffer after the event, it should trigger the function withou storing it in the
		buffer
*/

package main

import "fmt"

// declaring buffer and tracker to keep track of event triggered
var registerQueue []func()
var eventTriggered bool

func main() {

	registerQueue, eventTriggered = make([]func(), 0), false

	// Registring two functions in register callback
	RegisterCallback(&registerQueue, CallbackOne)
	RegisterCallback(&registerQueue, CallbackTwo)

	// Event is triggerd
	Events(&registerQueue)

	// After the event , another function is registered in callback
	RegisterCallback(&registerQueue, CallbackThree)
}

func Events(registerQueue *[]func()) {
	// If event is trigerred , we set the variable
	eventTriggered = true
	for len(*registerQueue) > 0 {
		(*registerQueue)[0]()
		(*registerQueue) = (*registerQueue)[1:]
	}
}

func RegisterCallback(registerQueue *[]func(), callbackfunc func()) {
	// If event is not triggered the store the function in the buffer for later execution
	// otherwise directly execute the function
	if eventTriggered == false {
		(*registerQueue) = append((*registerQueue), callbackfunc)
	} else {
		callbackfunc()
	}
}

func CallbackOne() {
	fmt.Println("lorem ipsum")
}

func CallbackTwo() {
	fmt.Println("Hello")
}

func CallbackThree() {
	fmt.Println("Test")
}
