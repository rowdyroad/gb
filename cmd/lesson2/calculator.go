package main

import (
	"fmt"
	"math"
)

const (
	OperationAdd = "+"
	OperationSub = "-"
	OperationMul = "*"
	OperationDev = "/"
	OperationPow = "^"
)

func main() {
	var a, b float64
	var operator string
	fmt.Println("Введите операнд операцию и второй операнд:")
	fmt.Scan(&a, &operator, &b)
	var res float64
	switch operator {
	case OperationAdd:
		res = a + b
	case OperationSub:
		res = a - b
	case OperationMul:
		res = a * b
	case OperationDev:
		if b == 0 {
			fmt.Println("Ошибка! Деление на ноль")
			return
		}
		res = a / b
	case OperationPow:
		res = math.Pow(a, b)
	default:
		fmt.Println("Ошибка! Неизвестная операция")
		return
	}

	fmt.Println("Результат: ", res)
}
