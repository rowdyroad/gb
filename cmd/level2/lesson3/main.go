package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)
type Struct struct {
	Str string
}


func waitForSignal() <-chan error {
	return make(chan error)
}

func main() {
	var sg sync.WaitGroup
	sg.Add(10)
	for i := 0; i < 10;i++ {
		go func() {
			defer sg.Done()
			defer fmt.Println("done")
		}()
	}
	sg.Wait()

	ch := make(chan Struct)
	ch1 := make(chan Struct)
	doneChan := make(chan struct{},1)

	go func() {
		defer func() {
			close(ch)
		}()
		for {
			select {
				case <-doneChan:
					return
				default:
					ch <- Struct{"1"}
			}
		}
	}()
	go func() {
		defer func() {
			close(ch1)
		}()
		for {
			select {
			case <-doneChan:
				return
			default:
				ch1 <- Struct{"2"}
			}
		}
	}()


	go func() {
		defer func() {
			fmt.Println("closed")
		}()
		arr := []Struct{}
		q := 2
		for q > 0 {
			select {
			case s, opened := <-ch:
				if opened {
					fmt.Println("from ch")
					arr = append(arr, s)
					if len(arr) > 10 {
						close(doneChan)
					}
				} else {
					q--
				}
			case s1, opened := <-ch1:
				if opened {
					fmt.Println("from ch1")
					arr = append(arr, s1)
					if len(arr) > 10 {
						close(doneChan)
						return
					}
				} else {
					q--
				}
			}
		}
	}()

	time.Sleep(time.Second)

	ctx, cancel := context.WithCancel(context.Background())

	ctx2, _ := context.WithTimeout(ctx, time.Second)

	context.WithDeadline(ctx2, time.Now().Add(time.Hour))

	ctx3 := context.WithValue(ctx, "key", "value")
	ctx3.Value()
	go func() {
		select {
			case <-ctx.Done():
			case <-ctx2.Done():
		}
		fmt.Println("closed")
	}()
	time.Sleep(time.Second)
	cancel()
	time.Sleep(time.Second)

}