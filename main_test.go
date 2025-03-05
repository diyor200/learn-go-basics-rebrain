package main

import "testing"

func BenchmarkMySlowFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MySlowFunction()

	}
}

func BenchmarkMyOptimizedFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MyOptimizedFunction()
	}
}
