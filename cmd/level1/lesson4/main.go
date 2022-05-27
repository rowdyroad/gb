package main

import (
	"fmt"
)

//slice: [len, cap, *arrayptr]

func f(a []int) {
	a = append(a, 1)
	a[1] = 666
}

func f2(m map[string]int) {
	m["f2"] = 999
	delete(m, "hello")

}

func main() {
	var ns []int
	_ = ns

	a := []int{1, 2, 34, 5, 6, 7, 8}
	b:= make([]int, len(a))
	copy(b, a)
	fmt.Println(a,b)
	a[1] = 99999
	fmt.Println(a,b)

	f(a)
	fmt.Println(a,b)

	fmt.Println("==========MAP==========")
	m := map[string]struct{
		A int
	}{}

	m["hello"] = struct{ A int }{A: 10}
	if val, ok := m["hello"]; ok {
		fmt.Println("its here:", val)
	}

	if _, ok := m["bye"]; !ok {
		fmt.Println("its not here")
	}

	var nm map[string]int
	fmt.Println(nm)

	nm2 := map[string]int{
		"hello": 2,
		"bye":3,
	}
	fmt.Println(nm2)

	nm3 := make(map[string]int, 100)
	fmt.Println(nm3)

	for key, value := range nm2 {
		fmt.Println(key, "=", value)
	}

	fmt.Println(len(nm2))

	delete(nm2, "bye")
	fmt.Println(nm2)

	x := map[string]map[string]int{}

	if _, ok := x["ggg"]; !ok {
		x["ggg"] = map[string]int{}
	}
	x["ggg"]["sfd"] = 1

	fmt.Println(x)

	f2(nm2)
	fmt.Println(nm2)

	f3 := func(a []int) {
		a = append(a, 1)
		a[1] = 666
	}
	newar := make([]int, 3, 10)
	newar[0], newar[1], newar[2] = 1,2,3
	fmt.Println(newar)
	f3(newar)
	fmt.Println(newar)
}