package main

import (
	"fmt"
	"gb/internal/mymath"
	"hash/crc32"
	"strconv"
	"unsafe"
)

var GlobalA = 1

type mymap struct {
	buckets [][]int
}

func NewMap() mymap {
	return mymap{
		buckets: make([][]int, 4),
	}
}

func (m *mymap) Insert(key string, val int) {
	hash := crc32.ChecksumIEEE([]byte(key))
	bucketId := int(hash) % len(m.buckets)
	m.buckets[bucketId] = append(m.buckets[bucketId], int(hash), val)
}

func (m mymap) Get(key string) (int,bool) {
	hash := crc32.ChecksumIEEE([]byte(key))
	bucketId := int(hash) % len(m.buckets)
	for i := 0; i < len(m.buckets[bucketId]);i+=2 {
		if m.buckets[bucketId][i] == int(hash) {
			return m.buckets[bucketId][i+1],true
		}
	}
	return 0, false
}

type Collection []int

func (c Collection) Sum() int {
	ret := 0
	for _, v := range c {
		ret += v
	}
	return ret
}

type MyString string

func (t MyString) ToInt(a ...int) int {
	r, _ := strconv.Atoi(string(t))
	return r
}

func (t MyString) ToIntWithArgs(a int) int {
	r, _ := strconv.Atoi(string(t))
	return r
}

func main() {

	a := []int{10:1, 15:5}
	fmt.Println(a)

	//key=>value // t->min o(n)
	//бинарный поиск

	//key=>value
	//1=>5
	//8=>9
	//9=>9
	//4=>3
	//
	//1=>5
	//4=>3
	//8=>9  len(ar) == 4
	//9=>9
	//
	//map[4] ?
	//
	//len(ar) == 4 / 2 => 2 => 8 => o(log2n)
	//

	//hashmap
	//hash(x) => y, !hash(y) != x
	//
	//hash(str)=>int, int64,int128
	//
	//sha256 => hash(str)=>256bit
	//crc16("hello") != crc("hell") != crc16("hl")
	//коллизия
	//x1 != x2 => hash(x1) == hash(x2)
	//
	//fmt.Printf("%x %d\n",crc32.ChecksumIEEE([]byte("hello")), crc32.ChecksumIEEE([]byte("hello")) % 4)
	//fmt.Printf("%x %d\n",crc32.ChecksumIEEE([]byte("helhellohellohellohellohellohellol1")),crc32.ChecksumIEEE([]byte("helhellohellohellohellohellohellol1")) % 4)
	//fmt.Printf("%x %d\n",crc32.ChecksumIEEE([]byte("helhellohellohellohellohellohellohellohellohellohellohellohellohellohellohellol2")),crc32.ChecksumIEEE([]byte("helhellohellohellohellohellohellohellohellohellohellohellohellohellohellohellol2")) % 4)
	//

	//
	m := NewMap()
	m.Insert("hello1", 2134)
	m.Insert("hello2", 1134)
	m.Insert("hello3", 35134)
	fmt.Println(m)
	fmt.Println(m.Get("hello1"))



	type SomeStruct struct {
		A int32 //0
		B int32 //0
		s [8]uint8 // [0,0,0,0,0,0,0,0]
		Da []uint8 // []
	}



	var x SomeStruct
	x.Da = append(x.Da, 1,2,3,4,5,6,7,8,9)

	type EmptyStruct struct{}
	var ex EmptyStruct
	fmt.Println(unsafe.Sizeof(ex))

	xmap := map[string]struct{
		A int
		B int
	}{}


	xmap["ok"] = struct{
		A int
		B int
	}{A:1,B:1}


	var f func()

	f = func() {
		fmt.Println("1")
	}

	f()

	type MyFunc func(a,b int) bool

	var mf MyFunc

	mf = func(a,b int) bool {
		fmt.Println(a+b)
		return a+b > 0
	}

	fmt.Println(mf(10,20))


	var mf3 func(a,b,c int) (string,string,string)


	mf3 = func(a,b,c int) (string,string,string) {
		return "hello1", "hello2", "hello3"
	}

	fmt.Println(mf3(1,2,3))


	mf4 := func(a,b,c int) (ret1 string,ret2 string, ret3 string) {
		ret1 = "a"
		ret2 = "b"
		return
	}

	fmt.Println(mf4(1,2,3))

	mf5 := func() MyFunc {
		return func(a,b int) bool {
			fmt.Println("callback", a,b)
			return true
		}
	}

	res := mf5()
	if res(11,21) {
		fmt.Println("ok")
	}

	type Struct struct {
		F MyFunc
		F1 func(a,b int) bool
	}

	var st Struct

	st.F = func(a,b int) bool {
		return true
	}

	st.F(10,10)


	vaFunc := func (prefix string, temp int, b ...string)  {
		fmt.Println(prefix,temp)
		for _, v := range a {
			fmt.Println("v:",v)
		}
	}

	vaFunc("hello", 1234)

	var val int = 1
	{
		var val int = 3
		fmt.Println(val)
		func() {
			fmt.Println(val)
		}()

	}

	fmt.Println(val)

	for _, x := range []int{1,2,33,4,5,6,7,8,9} {
		go func() {
			fmt.Println(x)
		}()
	}


	deferredFunc("hello1")


	fmt.Println("recursion:", recursion(5))
}

func recursion(n int64) int64 {
	if n == 1 {
		return 1
	}
	return n * recursion(n-1)
}


func deferredFunc(str string) {
	fmt.Println(str)
	defer func() {
		fmt.Println("bye1")
	}()
	defer func() {
		fmt.Println("bye2")
	}()
	defer func() {
		fmt.Println("bye3")
	}()
	defer func() {
		fmt.Println("bye4")
	}()

	if str == "a" {
		return
	}
	defer func() {
		fmt.Println("bye5")
	}()
	if str == "hello" {
		return
	}
	fmt.Println(str)

	for i := 0; i < 10;i++ {
		j := i
		defer func() {
			fmt.Println(j)
		}()

	}

	fmt.Println("test")



	cl := Collection{1,2,3,4,5,6}

	fmt.Println("Sum:", cl.Sum())

	var float mymath.Float = 10.2


	fmt.Println(float.Square())


}





