package main

import (
	"fmt"
	"gb/internal/player"
	"os"
)

type A struct {
	String string
	Number int64
	Real float64
}

type B struct {
	A
	IsOK bool
}

func (this B) Hello() {
	fmt.Println("from hello method:", this.String)
}


func main() {
	x := &B{A: A{"ok", 10,11}, IsOK: true}

	fmt.Println(x)
	y := A{"ok", 10, 11}
	fmt.Println(y)
	x.Hello()


	p, err := player.NewPlayer("Boris", "X" , 9999)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(p.GetAccountAsUSD())
	p.Income(1000)
	fmt.Println(p.GetAccountAsUSD())

	var a,b,c int
	var s string
	if _, err = fmt.Scan(&a,&b,&c);  err != nil {
		fmt.Println("Некорректно введены данные")
	}

	if _, err = fmt.Scan(&a,&s);  err != nil {
		fmt.Println("Некорректно введены данные")
	}
	fmt.Println(err)
}
