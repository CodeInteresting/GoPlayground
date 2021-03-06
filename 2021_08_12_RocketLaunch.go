package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// create abort channel
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <- tick:
			// do nothing
		case <- abort:
			fmt.Println("Launch Aborted!")
			return
		}
	}

	launch()
}

func launch() {
	fmt.Println("Rocket is into sky!!!")
}