package main

import "fmt"

func main() {
	var a = "initial" // Declears 1 variable. String infered by use of speech
	var b, c = 1, 2   // multiple variables declared at once
	var d = "true"
	var e int //Variables declared without a corresponding initialization are zero-valued
	var f = "apple"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	//Discovered
	var g, h int = 1, 2
	fmt.Println(g, h)
}
