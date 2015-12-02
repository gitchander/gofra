package main

import (
	"fmt"
	"log"
	"math/big"
)

func main() {
	if err := fastTest(); err != nil {
		log.Fatal(err.Error())
	}
}

func bigSample() {

	a := new(big.Float)

	fmt.Println("1: ", a.Prec())

	a.SetPrec(200)
	a.SetFloat64(0.01)

	a.Mul(a, a)
	a.Mul(a, a)
	a.Mul(a, a)
	a.Mul(a, a)
	a.Mul(a, a)
	a.Mul(a, a)
	a.Mul(a, a)
	a.Mul(a, a)
	a.Mul(a, a)
	a.Mul(a, a)
	a.Mul(a, a)
	a.Mul(a, a)
	a.Mul(a, a)
	a.Mul(a, a)
	a.Mul(a, a)
	a.Mul(a, a)

	fmt.Println("2: ", a.Prec())

	fmt.Println(a.Text('g', 200))
}
