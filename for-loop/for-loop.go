package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	j := 1
	for j <= 10 {
		fmt.Println(j)
		j++
	}

	a := "elvis"

	runes := []rune(a)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Println(string(runes))

	for i := range 3 {
		fmt.Println("range", i)
	}

}
