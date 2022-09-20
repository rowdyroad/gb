package main

import (
	"fmt"
	"log"
	"math"
	"reflect"
)

func print[T any](v T) {
	log.Println(v)
}

func printType(a interface{}) {
	switch v := a.(type) {
	case int:
		fmt.Println("this is int: ", v)
	case float64:
		fmt.Println("this is float: ", v)
	default:
		fmt.Println("this is something:", v)
	}
}

func libMethod(v interface{}) {
	s := reflect.ValueOf(v)
	t := reflect.TypeOf(v)
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name, s.Field(i))

		//if s.Field(i).Type().Kind() == reflect.String && s.Field(i).CanSet() {
		//	s.Field(i).SetString("Hello")
		//}
	}
	fmt.Println(reflect.TypeOf(v).Kind(), reflect.TypeOf(v))
	fmt.Println(v.(A).S)

	var x float64 = 3.4
	p := reflect.ValueOf(&x)
	p.Elem().SetFloat(100.3)

	fmt.Println(p.Elem().CanSet())
}

type A struct {
	C int
	D float64
	S string
}

func min[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float64](a, b T) T {
	return T(math.Max(float64(a), float64(b)))
}

func main() {

	var i int = 9
	var c float64 = 100.101

	printType(i)
	printType(c)

	str := "hello"

	a := A{C: 100}
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(str).Kind())
	fmt.Println(reflect.TypeOf(a).Kind())

	fmt.Println(reflect.ValueOf(a))

	fmt.Println("libmethod")
	libMethod(a)

	print(i)
	print(c)
	print(a)

	b := 10

	min(i, b)
}
