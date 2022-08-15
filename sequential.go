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
var mtx sync.Mutex

func main() {
	//queue, eventTriggered = make([]func(), 0), false
	//var wg sync.WaitGroup

	// Register 7 functions --> call Event --> Add 2 functions in register
	wg.Add(10)
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)
	ch4 := make(chan bool)
	ch5 := make(chan bool)
	ch6 := make(chan bool)
	ch7 := make(chan bool)
	che := make(chan bool)
	ch8 := make(chan bool)
	ch9 := make(chan bool)

	go RegisterCallback(CallbackOne, &wg, &mtx, ch1)
	go RegisterCallback(CallbackTwo, &wg, &mtx, ch2)
	go RegisterCallback(CallbackThree, &wg, &mtx, ch3)
	go RegisterCallback(CallbackFour, &wg, &mtx, ch4)
	go RegisterCallback(CallbackFive, &wg, &mtx, ch5)
	go RegisterCallback(CallbackSix, &wg, &mtx, ch6)
	go RegisterCallback(CallbackSeven, &wg, &mtx, ch7)
	go Events(&wg, &mtx, che)
	go RegisterCallback(CallbackEight, &wg, &mtx, ch8)
	go RegisterCallback(CallbackNine, &wg, &mtx, ch9)

	<-ch1
	<-ch2
	<-ch3
	<-ch4
	<-ch5
	<-ch6
	<-ch7
	<-che
	<-ch8
	<-ch9
	wg.Wait()

}

func Events(wg *sync.WaitGroup, mtx *sync.Mutex, ch chan bool) {
	defer wg.Done()
	//mtx.Lock()
	eventTriggered = true
	//mtx.Unlock()
	for len(queue) > 0 {
		queue[0]()
		queue = queue[1:]
	}
	ch <- true
}

func RegisterCallback(callbackfunc func(), wg *sync.WaitGroup, mtx *sync.Mutex, ch chan bool) {
	defer wg.Done()
	if eventTriggered == false {
		mtx.Lock()
		queue = append(queue, callbackfunc)
		mtx.Unlock()
	} else {
		callbackfunc()
	}
	ch <- true
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
