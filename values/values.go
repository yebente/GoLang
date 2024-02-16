package main

import "fmt"

func main() {
	//strings which can be added together with +
	fmt.Println("go" + "lang")

	//integers and floats
	fmt.Println("1+1 =", 1+1)
	fmt.Println("0.7/3.0 =", 0.7/0.3)

	//Booleans, with boolean operators
	fmt.Println(true && false)
	fmt.Println(true && true)
	fmt.Println(false && true)
	fmt.Println(false && false)

	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)

	fmt.Println(!true)
	fmt.Println(!false)

}
