package main

import (
	"fmt"
	"math/rand"
	"time"
)

var buffer = make([]int, 0, 4)

func main() {
	startTime := time.Now()
	GetId()
	GetId()
	GetId()
	GetId()
	GetId()
	GetId()
	GetId()
	GetId()
	GetId()
	GetId()
	GetId()
	GetId()
	endTime := time.Now()
	timeTaken := endTime.Sub(startTime)
	fmt.Println(timeTaken)
}

func GetId() {
	if len(buffer) > 0 {
		// Removing an elemetn from front will reduce the slice capacity so therefore it's better to remove from end
		top := buffer[len(buffer)-1]
		buffer = buffer[:len(buffer)-1]
		fmt.Println(top)
		//return
	} else {
		GetListOfIds()
		top := buffer[len(buffer)-1]
		buffer = buffer[:len(buffer)-1]
		fmt.Println(top)
		//return
	}
}

func GetListOfIds() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < cap(buffer); i++ {
		buffer = append(buffer, rand.Intn(100))
	}
	//fmt.Println(buffer)
	time.Sleep(1000 * time.Millisecond)
}
