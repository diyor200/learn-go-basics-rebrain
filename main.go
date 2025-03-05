package main

import (
	"math/rand"
	"time"
)

func main() {

}

func MySlowFunction() int {
	a := rand.Intn(100)
	b := rand.Intn(100)
	time.Sleep(time.Millisecond * 100)
	return a + b
}

func MyOptimizedFunction() int {
	a := rand.Intn(100)
	b := rand.Intn(100)
	time.Sleep(time.Millisecond * 10)
	return a + b
}
