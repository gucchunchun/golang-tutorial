// Namespace
package main

import (
	"fmt"
	"math/rand"
)

// Entry point
func main() {
	task1()
	fmt.Println()

	task2()
	fmt.Println()

	task3()
	fmt.Println()

	task4()
	fmt.Println()

	task5()
	fmt.Println()

	task6()
	fmt.Println()

	task7()
	fmt.Println()

	task8()
	fmt.Println()

	task9()
	fmt.Println()

	task10()
	fmt.Println()

	task11()
	fmt.Println()

	task12()
	fmt.Println()
}

func task1() {
	fmt.Println("## 課題1")
	fmt.Println("Hello World")
}

func task2() {
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
}

func task3() {
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
}

func task4() {
	b := false
	i := 100
	f := 3.14
	s := "Hello World"
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

	fmt.Printf("%v: ", s)
	fmt.Printf("%T, ", s)
	fmt.Printf("%p, ", &s)
	fmt.Printf("%s, ", s)
	fmt.Println()
}

func task5() {
	fmt.Println("## 課題5")

	a := rand.Int()
	b := rand.Int()

	fmt.Printf("a: %d, b: %d\n", a, b)
	if a > b {
		fmt.Println("a > b")
	} else {
		fmt.Println("a <= b")
	}
}

func task6() {
	fmt.Println("## 課題6")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
func task7() {
	fmt.Println("## 課題7")

	a := rand.Intn(100)

	fmt.Printf("a: %d\n", a)

	// a = 50

	switch a {
	case 50:
		fmt.Println("当選")
	default:
		fmt.Println("落選")
	}

}

func task8() {
	fmt.Println("## 課題8")

	add := func(a int, b int) int {
		return a + b
	}

	a := 1
	b := 3

	sum := add(a, b)
	fmt.Println(sum)
}

func task9() {
	fmt.Println("## 課題9")

	lottery := func() (int, bool) {
		a := rand.Intn(100)
		return a, a == 50
	}

	fmt.Println(lottery())
}

func task10() {
	fmt.Println("## 課題10")

	mustBe50 := func(i int) error {
		if i == 50 {
			return nil
		}
		return fmt.Errorf("not 50")
	}

	a := 50
	b := 51

	fmt.Println(a)
	if err := mustBe50(a); err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Printf("%d is 50!\n", a)
	}
	fmt.Println(b)
	if err := mustBe50(b); err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Printf("%d is 50!\n", b)
	}

	fmt.Println(mustBe50(50))
	fmt.Println(mustBe50(51))
}

func task11() {
	fmt.Println("## 課題11")

	// Array
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])
	fmt.Println(a[3])
	fmt.Println(a[4])
	fmt.Println()

	as := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(as)
	fmt.Println()

	// Slice
	s := a[1:4]
	fmt.Println(s)
	fmt.Println(s[0])
	fmt.Println(s[1])
	fmt.Println(s[2])
	fmt.Println()

	ss := as[1:4]
	fmt.Println(ss)
	fmt.Println(ss[0])
	fmt.Println(ss[1])
	fmt.Println(ss[2])
	fmt.Println()

	// Equivalent Expressions
	ss = as[0:5]
	fmt.Println(ss)
	ss = as[0:]
	fmt.Println(ss)
	ss = as[:5]
	fmt.Println(ss)
	ss = as[:]
	fmt.Println(ss)
	fmt.Println()

	ss[0] = "f"
	fmt.Println(as)
	fmt.Println(ss)
	fmt.Println()

	sb := []bool{true, true, false}
	fmt.Println(sb)
}

func task12() {
	fmt.Println("## 課題12")

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Println()

	s = append(s, 6)
	fmt.Println(s)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Println()

	s = append(s, 7, 8, 9)
	fmt.Println(s)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Println()

	s = append(s, 10, 11, 12)
	fmt.Println(s)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Println()

	s = s[:0]
	fmt.Println(s)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Println()

	s = s[:4]
	fmt.Println(s)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Println()

	s = s[2:]
	fmt.Println(s)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Println()
}
