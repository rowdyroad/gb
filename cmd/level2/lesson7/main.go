package main

import (
	"fmt"
	"reflect"
)


//go:generate stringer -type=Pill
type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)



type Struct struct {
	B int
	A int
	C string
}

func (s Struct) Hello(a int) string {
	return fmt.Sprintf("hello: %d", a)
}


type Hello interface {
	Hello(a int) string
}

func a(a float64, b string, c int) (float64,string,int) {
	return a+1.0,b+"2",c+3
}


func f(fn interface{}) {
	if reflect.TypeOf(fn).Kind() != reflect.Func {
		panic("not function")
	}
	fmt.Println("in:",reflect.TypeOf(fn).NumIn())
	for i := 0; i < reflect.TypeOf(fn).NumIn(); i++ {
		fmt.Println("\t", reflect.TypeOf(fn).In(i))
	}
	fmt.Println("out:",reflect.TypeOf(fn).NumOut())
	for i := 0; i < reflect.TypeOf(fn).NumOut(); i++ {
		fmt.Println("\t", reflect.TypeOf(fn).Out(i))
	}

	ret := reflect.ValueOf(fn).Call([]reflect.Value{reflect.ValueOf(1.0),reflect.ValueOf("hell"),reflect.ValueOf(1)})
	for i := 0; i < len(ret);i++ {
		fmt.Println("i", ret[i])
	}


}

func main() {
	//var i interface{} // (tabl, *ptr)

	s := Struct{
		A: 1,
		B: 4,
		C:"hello",
	}
	t := reflect.TypeOf(s)

	var x  interface{} = s
	if _, ok := x.(Hello); ok {

	}

	var ff float64 = 3.10

	//reflect.ValueOf(&ff).SetFloat(100)
	fmt.Println(reflect.ValueOf(&ff).Elem().CanSet())

	reflect.ValueOf(&ff).Elem().SetFloat(100)
	fmt.Println(ff)
	if t.Implements(reflect.TypeOf((*Hello)(nil)).Elem()) {

	}

	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name, t.Field(i).Type, reflect.ValueOf(s).Field(i).Interface())
	}

	f(a)
}