package main

import (
	"encoding/json"
	"fmt"
)

func funcA() {
	for i := 0; i < 100;i++ {
		for j := 0; j < 100; i++ {

		}
	}
}

type MyStruct struct {}
func (f MyStruct) MarshalJSON() ([]byte, error) {
	return []byte(`{"a": 0}`), nil
}

func funcB() {
		j, _ := json.Marshal(MyStruct{})
		fmt.Println(string(j))
}
func main() {
	fmt.Println("hello")

	funcB()
}
