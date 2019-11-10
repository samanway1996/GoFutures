package main

import (
	"fmt"
	"time"
	"math/rand"
)

type isDone struct {
	isCompleteChannel chan int
	isCompleteFlag bool
}

var factorial [100000]int

func myfun(n int) *isDone {
	
	// Creation and initialization of isDone structure
	var isComplete isDone
	isComplete.isCompleteChannel = make(chan int)
	isComplete.isCompleteFlag = false
	
	// This is our actual function to be called in async way
	go func() {
		factorial[0] = 1
		for i:=1;i<n;i++ {
			factorial[i] = factorial[i-1] * i
			factorial[i] = factorial[i] % 1000007	
		}
		// Writing into chan before exiting the function
		isComplete.isCompleteChannel <- 1
		fmt.Printf("Computation done\n")
	}()
	
	go pollRoutineStatus(&isComplete)
	return &isComplete
}

// Handle func to check if async routine is complete or not
func pollRoutineStatus(isComplete *isDone) {
	select {
		case <- isComplete.isCompleteChannel:
			isComplete.isCompleteFlag = true
	}
}

func main() {
	n := 1000
	isFinished := myfun(n)

	if isFinished.isCompleteFlag {
		rnd := rand.Intn(n)
		fmt.Printf("Value of factorial %d: %d\n", rnd, factorial[rnd])
	} else {
		fmt.Println("Routine still not finished")
	}
	
	// Sleep
	time.Sleep(time.Second*2)
	
	if isFinished.isCompleteFlag {
		rnd := rand.Intn(n)
		fmt.Printf("Value of factorial %d: %d\n", rnd, factorial[rnd])
	} else {
		fmt.Println("Routine still not finished")
	}
	
	//Sleep
	time.Sleep(time.Second*1)
	
	if isFinished.isCompleteFlag {
		rnd := rand.Intn(n)
		fmt.Printf("Value of factorial %d: %d\n", rnd, factorial[rnd])
	} else {
		fmt.Println("Routine still not finished")
	}
}
