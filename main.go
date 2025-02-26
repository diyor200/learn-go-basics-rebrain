package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(6)
	for i := 0; i < 10; i++ {
		go func(j int) {
			fmt.Println("Starting goroutine: ", j)
			for {
			}
		}(i)
	}

	select {}
}
