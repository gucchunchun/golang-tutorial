package main

import "fmt"

type CustomError struct {
	Message string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("CustomError: %s", e.Message)
}

func returnCustomError(i int) error {
	if i == 0 {
		return CustomError{Message: "i must not be 0"}
	}
	return nil
}

func task22() {
	a := 1
	b := 0

	fmt.Println(a)
	err := returnCustomError(a)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("No error")
	}
	fmt.Println()

	fmt.Println(b)
	err = returnCustomError(b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("No error")
	}
}
