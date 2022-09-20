package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Config struct {
	Address string
	Port    int
	Timeout time.Duration
}

//MyStruct struct of object
type MyStruct struct{}

func (f MyStruct) MarshalJSON() (string, error) {
	return `{"a": 0}`, nil
}
func main() {
	j, _ := json.Marshal(MyStruct{})
	fmt.Println(string(j))
}
