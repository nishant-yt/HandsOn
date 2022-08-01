/*
Implementation of Register Callback and Events using thread with no LOCK/MUTEX

Goal :  Register functions in buffer but all the calls to registercallback will run concurrently in different threads

		In case any function is registered in the buffer after the event, it should trigger the function without storing it in the
		buffer , note we are calling these function concurrently using multiple threads


		NOTE: With no mutex / lock here , the buffer will have different number of elements everytime because buffer will get overwritten
			  and we will have data corruption / loss and therefore not all callbacks will get executed when event will be triggered
*/
package main

import (
	"fmt"
	"sync"
)

var queue = make([]func(), 0)
var eventTriggered = false
var wg sync.WaitGroup

func main() {

	wg.Add(10)
	go RegisterCallback(CallbackOne, &wg)
	go RegisterCallback(CallbackTwo, &wg)
	go RegisterCallback(CallbackThree, &wg)
	go RegisterCallback(CallbackFour, &wg)
	go RegisterCallback(CallbackFive, &wg)
	go RegisterCallback(CallbackSix, &wg)
	go RegisterCallback(CallbackSeven, &wg)
	// Triggering The event
	go Events(&wg)
	// Registering another 2 events after the event
	go RegisterCallback(CallbackEight, &wg)
	go RegisterCallback(CallbackNine, &wg)
	wg.Wait()

}

func Events(wg *sync.WaitGroup) {
	defer wg.Done()

	eventTriggered = true
	for len(queue) > 0 {
		queue[0]()
		queue = queue[1:]
	}
}

func RegisterCallback(callbackfunc func(), wg *sync.WaitGroup) {
	defer wg.Done()
	if eventTriggered == false {
		queue = append(queue, callbackfunc)
	} else {
		callbackfunc()
	}
}

func CallbackOne() {
	fmt.Println("TEST ONE")
}

func CallbackTwo() {
	fmt.Println("TEST TWO")
}

func CallbackThree() {
	fmt.Println("TEST THREE")
}

func CallbackFour() {
	fmt.Println("TEST FOUR")
}

func CallbackFive() {
	fmt.Println("TEST FIVE")
}

func CallbackSix() {
	fmt.Println("TEST SIX")
}

func CallbackSeven() {
	fmt.Println("TEST SEVEN")
}

func CallbackEight() {
	fmt.Println("TEST EIGHT")
}

func CallbackNine() {
	fmt.Println("TEST NINE")
}
