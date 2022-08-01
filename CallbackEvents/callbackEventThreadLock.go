/*
Implementation of Register Callback and Events using thread with MUTEX/LOCK

Goal :  Register functions in buffer but all the calls to registercallback will run concurrently in different threads

		In case any function is registered in the buffer after the event, it should trigger the function without storing it in the
		buffer , note we are calling these function concurrently using multiple threads


		NOTE: Here we have mutex /lock , so at any time only one thread/ goroutine will access the buffer so that all the calls
			  get resgitered without corrupting/losing the data
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
	// creating a mutex
	var mtx sync.Mutex

	wg.Add(7)

	go RegisterCallback(CallbackOne, &queue, &wg, &mtx)
	go RegisterCallback(CallbackTwo, &queue, &wg, &mtx)
	go RegisterCallback(CallbackThree, &queue, &wg, &mtx)
	go RegisterCallback(CallbackFour, &queue, &wg, &mtx)
	go RegisterCallback(CallbackFive, &queue, &wg, &mtx)
	go RegisterCallback(CallbackSix, &queue, &wg, &mtx)
	go RegisterCallback(CallbackSeven, &queue, &wg, &mtx)

	wg.Wait()

	fmt.Println(eventTriggered)
	ch := make(chan bool)
	go Events(&queue, ch, &eventTriggered)
	//<-ch
	wg.Add(2)

	go RegisterCallback(CallbackEight, &queue, &wg, &mtx)
	go RegisterCallback(CallbackNine, &queue, &wg, &mtx)
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

func RegisterCallback(callbackfunc func(), rqueue *[]func(), wg *sync.WaitGroup, mtx *sync.Mutex) {
	if eventTriggered == false {
		mtx.Lock()
		(*rqueue) = append((*rqueue), callbackfunc)
		mtx.Unlock()
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
