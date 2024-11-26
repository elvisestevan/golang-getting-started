package main

import "fmt"

type CustomError struct {
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func main() {
	_, err := TestCustomError()
	if err != nil {
		fmt.Println(err)
	}
}

func TestCustomError() (int, error) {
	return 0, &CustomError{"Custom error message"}
}

func DeferF() {
	i := 0
	defer println(i)
	i++
	defer println(i)
	defer func() { i++ }()
	println(i)
	defer println(i)
}

func DeferPanicRecover() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("panic %v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
