package main

func main() {
	a := "hello"

	helloChan := make(chan *string)
	byeChan := make(chan *string)

	go func() {
		for hello := range helloChan {
			if *hello != "hello" {
				panic("not hello")
			}
			*hello = "bye"
			byeChan <- hello
		}
	}()

	go func() {
		for bye := range byeChan {
			if *bye != "bye" {
				panic("not bye")
			}
				*bye = "hello"
				helloChan <- bye
		}
	}()


	helloChan <- &a

	select {

	}
}
