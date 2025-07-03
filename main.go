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

	fmt.Println(str_variable1)
	str_variable1 = "Hello World3"
	fmt.Println(str_variable1)
	fmt.Println(str_variable2)
	fmt.Println(str_variable3)
	fmt.Println()

	fmt.Println("## 課題3")
	var (
		b bool
		i int
		f float64
	)
	b = true
	// int: -9223372036854775808 to 9223372036854775807
	i = 9223372036854775807
	f = 3.14
	fmt.Println(b)
	fmt.Printf("%T", b)
	fmt.Println()
	fmt.Println(i)
	fmt.Printf("%T", i)
	fmt.Println()
	fmt.Println(f)
	fmt.Printf("%T", f)
	fmt.Println()
	fmt.Println()

	// int8: -128 to 127
	var i8 int8 = 127
	fmt.Println(i8)
	// int16: -32768 to 32767
	var i16 int16 = 225
	fmt.Println(i16)
	// uint8: 0 to 255
	var ui8 uint8 = 225
	fmt.Println(ui8)

	fmt.Println(i + i)
	// fmt.Println(i + i8) invalid operation: i + i8 (mismatched types int and int8)
	fmt.Println()

	fmt.Println("## 課題4")

	fmt.Printf("%v: ", b)
	fmt.Printf("%T, ", b)
	fmt.Printf("%p, ", &b)
	fmt.Printf("%t, ", b)
	fmt.Println()

	fmt.Printf("%v: ", i)
	fmt.Printf("%T, ", i)
	fmt.Printf("%p, ", &i)
	fmt.Printf("%b, ", i)
	fmt.Printf("%d, ", i)
	fmt.Printf("%o, ", i)
	fmt.Println()

	fmt.Printf("%v: ", f)
	fmt.Printf("%T, ", f)
	fmt.Printf("%p, ", &f)
	fmt.Printf("%b, ", f)
	fmt.Printf("%g, ", f)
	fmt.Println()

	fmt.Printf("%v: ", str_variable1)
	fmt.Printf("%T, ", str_variable1)
	fmt.Printf("%p, ", &str_variable1)
	fmt.Printf("%s, ", str_variable1)
	fmt.Println()

}
