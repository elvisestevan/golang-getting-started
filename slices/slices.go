package main

import "fmt"

func main() {
	fruits := make([]string, 0, 10)

	fruits = append(fruits, "apple", "banana", "grape", "orange", "mango")

	fruits2 := append(fruits, "kiwi", "papaya")

	fmt.Println("Fruits:", fruits)
	fmt.Println(&fruits[0])
	fmt.Println(cap(fruits))
	fmt.Println(len(fruits))
	fmt.Println("Fruits2:", fruits2)
	fmt.Println(&fruits2[0])
	fmt.Println(cap(fruits2))
	fmt.Println(len(fruits2))

}
