package main

import "fmt"

//go:generate stringer -type=Pill
type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)



func main() {
	////runtime.GOMAXPROCS(1)
	//
	////go fmt.Println("hello")
	////
	////runtime.GC()
	////
	////for i := 0;;i++ {
	////	runtime.Gosched()
	////}
	////fmt.Println("CPUs:", runtime.NumCPU())
	////
	////a := map[string]string{}
	////for i := 0; i < 100;i++ {
	////	go func() {
	////		a["test"]  = fmt.Sprintf("%d",i)
	////	}()
	////}
	////time.Sleep(time.Second)
	//
	//trace.Start(os.Stderr)
	//defer trace.Stop()
	//wg := sync.WaitGroup{}
	//for i := 0; i < 1<<4; i += 1 {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		for i := 0; i < 1e8; i += 1 {
	//		}
	//	}()
	//}
	//wg.Wait()

	fmt.Println(Placebo)
}
