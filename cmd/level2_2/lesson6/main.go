package main

func main() {
	b := 333
	go func() {
		b++
	}()
	b--

	//for k, v := range m {
	//	fmt.Println(k, v)
	//}
}
