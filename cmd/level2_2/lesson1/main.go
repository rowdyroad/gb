package main

import (
	"errors"
	"fmt"
	//"io"
)

type Math interface {
	Div(a,b float64) (float64, error)
}

type HTTPError struct {
	code int
	message string
	err error
}

func (e HTTPError) Unwrap() error {
	return e.err
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("%s (%d)", e.message, e.code)
}

func NewHTTPError(code int, message string, categoryErr error) HTTPError {
	return HTTPError{code, message, categoryErr}
}

var (
	ErrHTTP = errors.New("http error")

	ErrHTTPNotFound = fmt.Errorf("%w:not found", ErrHTTP)
	ErrHTTPForbidden = fmt.Errorf("%w:forbidden", ErrHTTP)

	ErrNotFound = NewHTTPError(404, "not found", ErrHTTP)
	ErrForbidden  = NewHTTPError(403, "forbidden",ErrHTTP)
)

func getPage(path string) error {
	switch path {
	case "/passwd":
		return ErrForbidden
	case "/user/1":
		return nil
	case "/user/2":
		return ErrHTTPForbidden
	case "/user/3":
		return ErrHTTPNotFound
	default:
		return ErrNotFound
	}
}

func div(a,b float64) (ret float64, err error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("err in dev2:", e)
			err = e.(error)
		}
	}()
	defer func() {
		if e  := recover(); e != nil {
			fmt.Println("err in dev:", e)
			panic(e)
		}
	}()
	if b == 0 {
		panic(errors.New("divide by zero"))
	}
	return a/b, nil
}



func main() {
	//fmt.Println(getPage("/user/1"))
	//fmt.Println(getPage("/passwd"))
	//fmt.Println(getPage("/users"))
	//if err := getPage("/users"); err != nil {
	//	fmt.Println(err)
	//}
	//err := getPage("/passwd")
	//switch {
	//case errors.Is(err, ErrHTTP):
	//	fmt.Println("Err HTTP", err)
	//case errors.Is(err, io.EOF):
	//	fmt.Println("err eof", err)
	//default:
	//	panic(err)
	//}
	//
	//defer fmt.Println("1")
	//defer fmt.Println("2")
	//defer fmt.Println("3")
	//
	//for i := 0; i < 10;i++ {
	//	defer fmt.Printf("%d\n",i)
	//}
	//
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("We've got an error:", err)
	//	}
	//}()
	res, err := div(1,0)
	fmt.Println("div:", res, err)

}