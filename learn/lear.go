package main

import "fmt"

//func main() {
//	gen := func(ctx context.Context) <-chan int {
//		intChan := make(chan int)
//		n := 1
//
//		go func() {
//			defer close(intChan)
//			for {
//				select {
//				case <-ctx.Done():
//					fmt.Println(ctx.Err())
//					return
//				case intChan <- n:
//					n++
//				}
//			}
//		}()
//
//		return intChan
//	}
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	for i := range gen(ctx) {
//		fmt.Println(i)
//		if i == 13 {
//			break
//		}
//	}
//}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
