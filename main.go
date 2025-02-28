package main

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	//fmt.Println(runTasks())
	fmt.Println(runTasksWithErrorGroup())
}

func runTasksWithErrorGroup() error {
	var (
		eg    = new(errgroup.Group)
		tasks = []func() error{
			func() error { time.Sleep(time.Second * 2); return errors.New("task 1 fail") },
			func() error { time.Sleep(time.Second * 1); return errors.New("task 2 fail") },
			func() error { time.Sleep(time.Second * 3); return errors.New("task 3 fail") },
		}
	)

	for _, task := range tasks {
		t := task
		eg.Go(func() error {
			return t()
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil

}

// without sync.errorGroup
//func runTasks() error {
//	var (
//		wg      sync.WaitGroup
//		errChan = make(chan error, 1)
//		tasks   = []func() error{
//			func() error { time.Sleep(time.Second * 2); return errors.New("task 1 fail") },
//			func() error { time.Sleep(time.Second * 1); return errors.New("task 2 fail") },
//			func() error { time.Sleep(time.Second * 3); return errors.New("task 3 fail") },
//		}
//	)
//
//	for _, task := range tasks {
//		wg.Add(1)
//		go func(t func() error) {
//			defer wg.Done()
//			if err := t(); err != nil {
//				select {
//				case errChan <- err:
//				default:
//					fmt.Println("can't write error to errChan, because it is full! err:", err)
//				}
//			}
//		}(task)
//	}
//
//	wg.Wait()
//	close(errChan)
//
//	if err, ok := <-errChan; ok {
//		return err
//	}
//
//	return nil
//}
