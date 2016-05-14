package main

import (
	"flag"
	"fmt"
	"os"
	// "strconv"
	"strings"
)

var n = flag.Bool("n", false, "new line")
var sep = flag.String("s", " ", "separator")

func main() {
	// var s, sep string
	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	// fmt.Println(s)

	// s, sep := "", ""
	// for _, arg := range os.Args[1:] {
	// 	s += sep + arg
	// 	sep = " "
	// }
	// fmt.Println(s)

	// for i, arg := range os.Args {
	// 	fmt.Println(strconv.Itoa(i) + " " + arg)
	// }
	// fmt.Println(os.Args)
	fmt.Print(strings.Join(os.Args, *sep))
	if !*n {
		fmt.Println()
	}
}

var global *int

func f() {
	var x int
	x = 1
	global = &x
}

func g() {
	y := new(int)
	*y = 1
}
