package main

import (
	"fmt"
	"sync"
)

type Config struct {
	Password string
}


func main() {

	var wg sync.WaitGroup

	bytePool := sync.Pool{
		New: func() interface{} {
			return make([]byte, 100)
		},
	}

	wg.Add(2)
	go func() {
		defer wg.Done()
		a := bytePool.Get().([]byte)
		fmt.Println("1->", a)
		for i := range a {
			a[i] = 1
		}
		bytePool.Put(a)
	}()

	go func() {
		defer wg.Done()
		a := bytePool.Get().([]byte)
		fmt.Println("2->", a)
		for i := range a {
			a[i] = 2
		}
		bytePool.Put(a)
	}()

	wg.Wait()

	a := bytePool.Get().([]byte)
	fmt.Println("3->", a)
	bytePool.Put(a)

	//var sg sync.WaitGroup
	//// sg.Done() == sg.Add(-1)
	//for i := 0; i < 10; i++ {
	//	sg.Add(1)
	//	var a int64
	//	go func() {
	//		defer sg.Done()
	//		defer fmt.Println("done")
	//		for j := 0; j < 100; j++ {
	//			atomic.AddInt64(&a, 1)
	//		}
	//	}()
	//}
	//sg.Wait()
	//fmt.Println("done")
	//for j := 0; j < 100;j++ {
	//	a := []int{}
	//	//var mux sync.Mutex
	//	var rmmux sync.RWMutex
	//	for k := 0; k < 10; k++ {
	//		go func() {
	//			for i := 0; i < 3; i++ {
	//				rmmux.Lock()
	//				a = append(a, len(a))
	//				for m := 0; m < len(a)/2;m++ {
	//					a[m],a[len(a)-m-1] = a[len(a)-m-1],a[m]
	//				}
	//				rmmux.Unlock()
	//			}
	//		}()
	//
	//		//go func() {
	//		//	for {
	//		//		rmmux.RLock()
	//		//		fmt.Println(len(a))
	//		//		rmmux.RUnlock()
	//		//	}
	//		//}()
	//
	//	}
	//	time.Sleep(50 * time.Millisecond)
	//	fmt.Println(a)
	//}

	//var a int64
	//var aMutex sync.Mute
	//
	//cfg := atomic.Value{}
	//
	//cfg.Store(Config{
	//	Password:"12345",
	//})
	//
	//var c *Config
	////
	////var g atomic.Value
	//for  j := 0; j < 10; j++ {
	//	var a int64
	//
	//	go func() { // http
	//			if c != nil && c.Password == "12345" {
	//
	//			}
	//
	//	}()
	//
	//	go func() {
	//		c = nil
	//	}()
	//
	//	time.Sleep(time.Second)
	//
	//	fmt.Println(a)
	//}


}
