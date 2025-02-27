package main

import (
	"fmt"
	"time"
)

func main() {
	semaphore := make(chan int, 3)

	for i := 0; i < 10; i++ {
		semaphore <- i
		go func() {
			defer func() {
				msg := <-semaphore
				fmt.Println(msg)
			}()
			time.Sleep(time.Millisecond * 1000) // some operations
		}()
	}
	for len(semaphore) > 0 {
		time.Sleep(time.Millisecond * 10)
	}
}
