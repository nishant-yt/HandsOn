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
var wg sync.WaitGroup
var mtx sync.Mutex

func main() {

	wg.Add(10)
	go RegisterCallback(CallbackOne, &wg, &mtx)
	go RegisterCallback(CallbackTwo, &wg, &mtx)
	go RegisterCallback(CallbackThree, &wg, &mtx)
	go RegisterCallback(CallbackFour, &wg, &mtx)
	go RegisterCallback(CallbackFive, &wg, &mtx)
	go RegisterCallback(CallbackSix, &wg, &mtx)
	go RegisterCallback(CallbackSeven, &wg, &mtx)
	go Events(&wg, &mtx)
	go RegisterCallback(CallbackEight, &wg, &mtx)
	go RegisterCallback(CallbackNine, &wg, &mtx)
	wg.Wait()
}

func Events(wg *sync.WaitGroup, mtx *sync.Mutex) {
	defer wg.Done()
	mtx.Lock()
	eventTriggered = true
	mtx.Unlock()
	for len(queue) > 0 {
		queue[0]()
		queue = queue[1:]
	}
}

func RegisterCallback(callbackfunc func(), wg *sync.WaitGroup, mtx *sync.Mutex) {
	defer wg.Done()
	if eventTriggered == false {
		mtx.Lock()
		queue = append(queue, callbackfunc)
		mtx.Unlock()
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
