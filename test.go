package main

import (
	"fmt"
	"math"
)

const a = "A"

// Test function prints hello world to con
func Test() {
	var x string
	x = "y"
	fmt.Println(x)
	A()
	B()
	C()
}

func T() {
	var r float64
	r = 6.1
	fmt.Println(math.Exp(r))
}

func A() {
	fmt.Println("Hello World")
}

func B() {
	var d = "d"
	fmt.Println(d)
}

func C() {
	fmt.Println("Hello World")
}
