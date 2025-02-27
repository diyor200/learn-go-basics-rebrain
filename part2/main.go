package main

import (
	"context"
	"fmt"
	"time"
)

//func main() {
//	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
//	tick := time.NewTicker(time.Millisecond * 50)
//
//	for {
//		select {
//		case t := <-tick.C:
//			fmt.Println(t)
//			cancel()
//		case <-ctx.Done():
//			fmt.Println("context deadline exceeded")
//			return
//		default:
//			fmt.Println("waiting")
//			time.Sleep(time.Millisecond * 20)
//		}
//	}
//}

func main() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		ch := make(chan int)
		i := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case ch <- i:
					i++
					time.Sleep(time.Millisecond * 300)
				}
			}
		}()

		return ch
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	startTime := time.Now()
	ticker := time.NewTicker(time.Second * 5)
	for n := range gen(ctx) {
		fmt.Println(n)
		select {
		case <-ticker.C:
			fmt.Printf("program stopped. Started:%v  Ended: %v\n", startTime, time.Now())
			return
		default:
		}
	}
}
