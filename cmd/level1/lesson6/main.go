package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func f(a *int) {
	fmt.Println("a: %p", a)
	*a = 1
}

type T struct {
	a int
}

func (t *T) Function(b int) {
	t.a = b
}

func (t T) Function2() int {
	return t.a *t.a
}

func main() {

	v := 999
	fmt.Println(v)
	fmt.Println("&v: %p", &v)
	f(&v)
	fmt.Println(v)


	//f, err := os.Open("text.txt")
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()
	//
	//for i := 0; i < 100; i++ {
	//	if i == 19 {
	//		return
	//	}
	//}

	var i int64 = 9423
	fmt.Println(unsafe.Sizeof(i))

	b := &i

	fmt.Printf("%d %p %d %p", i, &i, *(&i), b, *b)
	//
	//for {
	//	time.Sleep(time.Second)
	//	f := make([]int, 100000)
	//	PrintMemUsage()
	//	_ = f
	//}

}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}