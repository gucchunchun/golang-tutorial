// Namespace
package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Entry point
func main() {
	tasks := []func(){
		task1,
		task2,
		task3,
		task4,
		task5,
		task6,
		task7,
		task8,
		task9,
		task10,
		task11,
		task12,
		task13,
		task14,
		task15,
		task16,
		task17,
		task18,
		task19,
		task20,
		task21,
		task22,
		task23,
		task24,
		task25,
	}

	for i, task := range tasks {
		fmt.Printf("## 課題%d\n", i+1)
		task()
		fmt.Println()
	}
}

func task1() {
	fmt.Println("Hello World")
}

func task2() {
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

	// byte
	var by1 byte = 'A'
	var by2 byte = 65
	fmt.Println(by1)
	fmt.Printf("%T", by1)
	fmt.Println()
	fmt.Println(by2)
	fmt.Printf("%T", by2)
	fmt.Println()

}

func task4() {
	b := false
	i := 100
	f := 3.14
	s := "Hello World"

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
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
func task7() {
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
	add := func(a int, b int) int {
		return a + b
	}

	a := 1
	b := 3

	sum := add(a, b)
	fmt.Println(sum)
}

func task9() {
	lottery := func() (int, bool) {
		a := rand.Intn(100)
		return a, a == 50
	}

	fmt.Println(lottery())
}

func task10() {
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

	s = make([]int, 5)
	fmt.Println(s)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Println()

	s = make([]int, 5, 10)
	fmt.Println(s)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Println()

	var nilSlice []int
	fmt.Println(nilSlice)
	fmt.Printf("len=%d cap=%d %v\n", len(nilSlice), cap(nilSlice), nilSlice)
	fmt.Printf("Is nil? %t\n", nilSlice == nil)
	fmt.Println()
}

func task13() {
	// 切り出し操作は上記で行なってしまっていたためループの実装

	s := []int{1, 2, 3, 4, 5}

	for i, v := range s {
		fmt.Println(i, v)
	}

	for i, _ := range s {
		fmt.Println(i)
	}

	for _, v := range s {
		fmt.Println(v)
	}

	for i := range s {
		fmt.Println(i)
	}

}

func task14() {
	// 作成
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(m)
	fmt.Println()

	// 追加
	m["d"] = 4
	fmt.Println(m)
	fmt.Println()

	// 削除
	delete(m, "d")
	fmt.Println(m)
	fmt.Println()

	//検索
	fmt.Println(m["a"])
	fmt.Println(m["b"])
	fmt.Println(m["c"])
	fmt.Println(m["d"])
	fmt.Println(m["e"])
	fmt.Println()

	// string
	fmt.Println("ms")
	ms := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}
	fmt.Println(ms)
	fmt.Println(ms["d"] == "")
	fmt.Println()

	// bool
	fmt.Println("mb")
	mb := map[string]bool{
		"a": true,
		"b": true,
		"c": false,
	}
	fmt.Println(mb)
	fmt.Println(mb["d"])
	fmt.Println()

	// nil map
	fmt.Println("mNil")
	var mNil map[string]int
	fmt.Println(mNil)
	fmt.Println(mNil == nil)
	// mNil["a"] = 1
	// fmt.Println(mNil)
	fmt.Println()

	// make()
	fmt.Println("mMake")
	var mMake = make(map[string]int)
	fmt.Println(mMake)
	fmt.Println(mMake == nil)
	mMake["a"] = 1
	fmt.Println(mMake)

	// check if the key is present or not
	elem, ok := m["a"]
	fmt.Println(elem, ok)
	elem, ok = m["e"]
	fmt.Println(elem, ok)
}

func task15() {
	// 構造体の定義
	type Person struct {
		Name string
		Age  int
	}

	// 構造体の作成
	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p)
}

func task17() {
	var pi *int

	fmt.Println(pi)

	i := 10
	pi = &i

	fmt.Println(pi)
	fmt.Println(*pi)

	*pi = 20

	fmt.Println(i)
	fmt.Println(*pi)

	pi = nil
	fmt.Println(pi)
}

func task18() {
	input := "I love Go programming!"

	fmt.Println(input)
	fmt.Println(strings.ToLower(input))
	fmt.Println(strings.ToUpper(input))
	fmt.Println(strings.ToTitle(input))
	fmt.Println(strings.Contains(input, "Go"))
	fmt.Println(strings.ReplaceAll(input, "Go", "Golang"))
	fmt.Println()

	input = "apple,banana,orange"
	s := strings.Split(input, ",")
	fmt.Println(s)
	fmt.Println(strings.Join(s, ","))
	fmt.Println()

	input = "   Hello, World!   "
	fmt.Println(strings.TrimSpace(input))

}

func task19() {
	// int -> string
	i := 123
	s := strconv.Itoa(i)
	fmt.Println(s)
	fmt.Printf("%T", s)
	fmt.Println()

	// string -> int
	s = "456"
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
		fmt.Printf("%T", s)
	}
	fmt.Println()

	// float -> int
	f := 3.14
	i = int(f)
	fmt.Println(i)
	fmt.Printf("%T", i)
	fmt.Println()

	// bool -> string
	b := true
	s = strconv.FormatBool(b)
	fmt.Println(s)
	fmt.Printf("%T", s)
	fmt.Println()

	// string -> float
	s = "3.14"
	f, err = strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f)
		fmt.Printf("%T", f)
	}
}

func task20() {
	name := "Alice"
	age := 30

	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Println()

	num := 64.0
	fmt.Printf("Sqrt: %f\n", math.Sqrt(num))
	fmt.Println()

	fmt.Printf("現在の時刻は: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println()

	input := "I love Go. Go is awesome!"
	fmt.Println(strings.ReplaceAll(input, "Go", "Golang"))
	fmt.Println()

	i := 123
	s := strconv.Itoa(i)
	fmt.Println(s)
	fmt.Println()

}

func task23() {
	// 書き込み
	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString("Hello, Go File!")
	if err != nil {
		log.Fatal(err)
	}

	// 読み込み
	data, err := os.ReadFile("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println(string(data))

	// 追記
	file, err = os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString("This is an appended line.\n")
	if err != nil {
		log.Fatal(err)
	}
}

func task25() {
	task25_1()
	task25_2()
}
