package main

import (
	"book/ch2/tempconv"
	"fmt"
)

func main() {
	tempconv.CToF(1)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
	fmt.Println(a)
}

var a = b + c
var b = f()
var c = 1

func f() int {
	return c + 1
}
