// файл gosched/main.go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("Start")
	go func() {
		fmt.Println("Start goroutine")
		for {
			someWorkload()
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("End")
}

func someWorkload() {
}
