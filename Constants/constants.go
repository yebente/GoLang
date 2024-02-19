package main

import (
	"fmt"
	"math"
)

func main() {

	const a string = "constant"
	const n int = 500000000
	const d = 3e20 / n

	fmt.Println(a)
	fmt.Println(n)
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
