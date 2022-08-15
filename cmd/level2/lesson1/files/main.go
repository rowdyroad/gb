package main


import "os"

import "fmt"

func main() {

    defer func() {
	if err := recover(); err != nil {
	    fmt.Println(err)
	}
    }()

    for {
	_, err := os.Open("1.txt")
	if err != nil {
	    panic(err)
	}
    }
}