package main

import (
	"fmt"
	"time"
)

func main() {
	semaphore := make(chan int, 3)
	done := make(chan struct{})
	i := 0

	go func() {
		for ; ; i++ {
			semaphore <- i
			time.Sleep(time.Millisecond * 100)
		}
	}()
	go func() {
		time.Sleep(time.Millisecond * 1000)
		done <- struct{}{}
	}()

	msg := 0
L:
	for {
		select {
		case msg = <-semaphore:
			fmt.Println(msg)
		case <-done:
			fmt.Println("done")
			break L
		default:
			if msg >= 20 {
				break L
			}

			fmt.Println("waiting")
			time.Sleep(time.Millisecond * 200)
		}
	}

	fmt.Println("Success")
}
