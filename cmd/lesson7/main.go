package lesson7

import (
	"errors"
	"fmt"
)

var ErrHttp = errors.New("Http error" )
var ErrHttpNotFound = errors.New("http not found")

type Error struct {
	err error
	prev *Error
}

func (e Error) Error() string {
	return fmt.Sprint(e.err, e.prev.Error())
}

func (e Error) Unwrap() error {
	return e.prev
}

func main() {






}