package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var Count int

func main() {
	startTime := time.Now()
	buffer := make(chan int, 10)
	var wg sync.WaitGroup
	var mtx sync.Mutex

	//go FillBuffer(buffer)

	/*
		for v := range buffer {
			fmt.Println(v)
		}
	*/

	for i := 1; i <= 30; i++ {
		wg.Add(1)
		go GetUid(buffer, &wg, &mtx)
	}
	wg.Wait()
	endTime := time.Now()
	timeTaken := endTime.Sub(startTime)
	fmt.Println(timeTaken)
}

func GetUid(buf chan int, wg *sync.WaitGroup, mtx *sync.Mutex) {
	//fmt.Println("COUNT IS: ", Count)
	mtx.Lock()
	if Count == 0 {
		go FillBuffer(buf)
		Count = cap(buf)
	}

	v := <-buf

	fmt.Printf("%v  \t", v)
	Count -= 1
	mtx.Unlock()
	wg.Done()
}

func FillBuffer(buff chan int) {
	/*
		if len(buff) == 0 {
			rand.Seed(time.Now().UnixNano())
			for i := 0; i < cap(buff); i++ {
				buff <- rand.Intn(100)
			}
		}
	*/
	//fmt.Println("LENGTH OF BUFFER: ", len(buff))
	//fmt.Println("Inside FILL")
	time.Sleep(1000 * time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < cap(buff); i++ {
		buff <- rand.Intn(1000)
	}
	//fmt.Println("OTSIDE")
	//fmt.Println("COUNT IS: ", Count)
	//time.Sleep(1000 * time.Millisecond)
	//Count = cap(buff)

	//time.Sleep(100 * time.Millisecond)

	//close(buff)
}
