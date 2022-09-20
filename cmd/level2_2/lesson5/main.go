package main

import (
	"fmt"
	"sync"
)

func main() {
	//var c uint64
	//
	//for i := 0; i < 1000; i++ {
	//	go func() {
	//		atomic.AddUint64(&c, 1)
	//	}()
	//}
	//
	//time.Sleep(time.Second)
	//
	//fmt.Println(c)
	//
	//var s uint64 = 1
	//
	//if atomic.CompareAndSwapUint64(&s, 1, 0) {
	//	//critical section
	//
	//	//
	//	atomic.SwapUint64(&s, 1)
	//}
	//var m sync.Mutex
	//
	//for i := 0; i < 1000; i++ {
	//	go func(j int) {
	//		m.Lock() //<-999
	//		defer m.Unlock()
	//		if c > 0 {
	//			c = c / c
	//		}
	//
	//		c++
	//	}(i)
	//}
	//
	//fmt.Println("before lock")
	//
	//var v uint64
	//
	//var rw sync.RWMutex
	//
	//for i := 0; i < 1000; i++ {
	//	go func() {
	//		rw.RLock() //<-
	//		defer rw.RUnlock()
	//		if v%10 == 0 {
	//			fmt.Println("ok")
	//		}
	//	}()
	//}
	//
	//for i := 0; i < 5; i++ {
	//	go func() {
	//		for {
	//			rw.Lock()
	//			v++
	//			rw.Unlock()
	//		}
	//	}()
	//}
	//
	//af := func(m2 sync.Mutex) {
	//	m2.Lock()
	//	defer m2.Unlock()
	//}

	//mp := map[string]string{}
	//var m sync.Mutex
	//var sm sync.Map
	//go func() {
	//	for {
	//		// v:= mp["hello"];
	//		v, has := sm.Load("hello")
	//		if !has {
	//			v.(string)
	//		}
	//		//
	//		sm.Store("hello", "hello")
	//		sm.Delete("hello")
	//		//delete(mp, "hello")
	//		//for key, value := range mp {
	//		//
	//		//}
	//		sm.Range(func(key, value any) bool {
	//
	//			return true;
	//		})
	//
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		m.Lock()
	//		mp["bye"] = "bye"
	//		m.Unlock()
	//	}
	//}()
	//
	//time.Sleep(time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()

			fmt.Println("hello", j)

		}(i)
	}

	wg.Wait() // <- wait

	f := func() {
		fmt.Println("done")
	}
	var on sync.Once
	var k int
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			on.Do(f)
			k++
		}()
	}

	wg.Wait()

	sp := sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024)
		},
	}

	bts := sp.Get().([]byte)

	sp.Put(bts)

}
