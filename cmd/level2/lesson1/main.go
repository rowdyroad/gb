package main

import (
	"errors"
	"fmt"
	"github.com/Chekunin/wraperr"
	"log"
	"runtime/debug"
)

var (
	ErrDivideByZero = errors.New("divide by zero")
	ErrNumberTooBig = errors.New("number is too big")
	ErrNumberTooSmall = errors.New("number is too small")

	ErrCalculation = errors.New("calculation")
	ErrHTTP = errors.New("http")
)


type MyError struct {
	A float64
	B float64
	Message string
	StackTrace string
}

func NewMyError( a,b float64, msg string) MyError {
	return MyError{a,b, msg, string(debug.Stack())}
}

func (r MyError) Error() string {
	return fmt.Sprintf("Error: %s (%f / %f) %s", r.Message, r.A, r.B, r.StackTrace)
}

func div(a,b float64) (float64, error) {
	if b == 0 {
		return 0,wraperr.Wrap(ErrCalculation,ErrDivideByZero)
	}
	return a / b, nil
}


func calc(a,b float64) (float64, error) {
	return div(a,b)
}

func deferExample() {
	x := 0
	defer fmt.Println(x+1)
	defer fmt.Println(x+2)
	defer fmt.Println(x+3)

	defer func() {
		fmt.Println(x+1)
	}()
	defer func() {
		fmt.Println(x+2)
	}()
	defer func() {
		fmt.Println(x+3)
	}()

	x = 100
}

func panicFunc() {
	panic("oops")
}


func test() (x int, err error) {
	defer func() {
		log.Println(x)
	}()
	x = 5
	return 99, nil
}



func main() {
	var a,b float64
	fmt.Scan(&a,&b)
	res, err := calc(a,b)


	if errors.Is(err, ErrCalculation) {
		log.Println("error calculation", err)
	}

	if errors.Is(err, ErrDivideByZero) {
		log.Println("divide by zero!", err)
	}

	switch err  {
	case ErrNumberTooBig:
		log.Println("Correct your the second number")
	case ErrDivideByZero:
		log.Println("Correct your the second number")
	case nil:
		log.Println("Result:",res)
	}
	defer fmt.Println("hello")
	for i := 0; i < 10; i++ {
		defer fmt.Printf("i:%d\n",i)
	}
	defer fmt.Println("bye")

	defer func() {
		if err := recover(); err != nil {
			log.Println("Oops:", err)
		}
	}()

	defer func() {
		if err := recover(); err != nil {
			log.Println("oops2:",err)
			panic(err)
		}
	}()
	panicFunc()

	test()
}
