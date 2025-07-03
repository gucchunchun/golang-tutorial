// Namespace
package main

import "fmt"

// Entry point
func main() {
	fmt.Println("## 課題１")
	fmt.Println("Hello World")
	fmt.Println()

	fmt.Println("## 課題2")
	// ○: initialize the zero-value of the type
	var str_variable1 string
	// △: not best practice?
	var str_variable2 string = "Hello World1"
	// ○: initialize the actual value
	str_variable3 := "Hello World2"

	str_variable1 = "Hello World3"
	fmt.Println(str_variable1)
	fmt.Println(str_variable2)
	fmt.Println(str_variable3)
	fmt.Println()
}
