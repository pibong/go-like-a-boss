package main

import (
	"fmt"
	"math"
	"gitlab.com/foo/go-like-a-boss/mylib"
)

func add(x int, y int) int {
	return x + y
}


func main() {
	var n, m int
	n = 3
	m = 5
	fmt.Printf("** Welcome in this stupid code **\n")
	fmt.Printf("%d + %d = %d\n", n, m, add(n, m))

	fmt.Printf("my lib says: %d\n", mylib.Multiply(n, m))

	var i, j int = 1, 2
	k := 3	//short assignment statement can be used in place of a var declaration with implicit type
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)

	var o int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", o, f, b, s)

	f = math.Sqrt(float64(n*n + m*m))
	var z = uint(f)
	fmt.Println(f, z)
}
