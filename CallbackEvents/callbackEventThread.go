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

var queue []func()
var eventTriggered bool

func main() {
	queue, eventTriggered = make([]func(), 0), false
	var wg sync.WaitGroup

	// Registring 7 functions concurrently in the buffer
	wg.Add(7)
	go RegisterCallback(CallbackOne, &queue, &wg)
	go RegisterCallback(CallbackTwo, &queue, &wg)
	go RegisterCallback(CallbackThree, &queue, &wg)
	go RegisterCallback(CallbackFour, &queue, &wg)
	go RegisterCallback(CallbackFive, &queue, &wg)
	go RegisterCallback(CallbackSix, &queue, &wg)
	go RegisterCallback(CallbackSeven, &queue, &wg)

	wg.Wait()

	fmt.Println(eventTriggered)
	ch := make(chan bool)

	// Event has been triggred
	go Events(&queue, ch, &eventTriggered)

	//<-ch

	// Registering another 2 events
	wg.Add(2)

	go RegisterCallback(CallbackEight, &queue, &wg)
	go RegisterCallback(CallbackNine, &queue, &wg)
	wg.Wait()
	<-ch
	fmt.Println(eventTriggered)

}

func Events(rqueue *[]func(), ch chan bool, isEventDone *bool) {
	fmt.Println("Queue: ", *rqueue, "Len of Queue: ", len(*rqueue))
	*isEventDone = true
	for len(*rqueue) > 0 {
		(*rqueue)[0]()
		(*rqueue) = (*rqueue)[1:]
	}
	ch <- true
}

func RegisterCallback(callbackfunc func(), rqueue *[]func(), wg *sync.WaitGroup) {
	if eventTriggered == false {
		(*rqueue) = append((*rqueue), callbackfunc)
	} else {
		callbackfunc()
	}
	wg.Done()
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
