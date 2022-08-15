package main

import (
	"reflect"
	"testing"
)


func BenchmarkTypeSwitch(t *testing.B) {
	var a interface{} = "string" // Struct{1,2,"str"}
	var c int
	for i := 0; i < t.N;i++ {
		switch a.(type) {
		case int,int8,int16,uint32:
			c++
		case string:
			c++
		case Struct:
			c++
		}
	}
}

func BenchmarkReflect(t *testing.B) {
	st := reflect.TypeOf(Struct{})
	var a interface{} = "string"
	var c int
	for i := 0; i < t.N; i++ {
		t :=  reflect.TypeOf(a)
		if t == st {
			c++
		}
	}
}
