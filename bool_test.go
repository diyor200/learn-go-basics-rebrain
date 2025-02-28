package main

import (
	"sync"
	"testing"
)

type Counter struct {
	A int
	B int
}

func (c *Counter) Inc() {
	c.A++
	c.B++
}

//func BenchmarkCounter_Inc(b *testing.B) {
//	counter := &Counter{}
//	for i := 0; i < b.N; i++ {
//		for j := 0; j < 10000; j++ {
//			b.StopTimer()
//			counter.Inc()
//			b.StartTimer()
//		}
//	}
//}

func BenchmarkCounter_WithPool(b *testing.B) {
	var counterPool = sync.Pool{
		New: func() interface{} { return Counter{} },
	}
	counter := counterPool.Get().(Counter)
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			b.StopTimer()
			counter.Inc()
			b.StartTimer()
		}
	}
}
