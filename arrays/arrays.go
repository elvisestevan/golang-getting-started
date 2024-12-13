package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := [3]int{1, 2, 3}

	for ix, num := range nums {
		fmt.Println(strconv.Itoa(ix) + ": " + strconv.Itoa(num*2))
	}

	list := []string{"1", "2", "3"}

	newList := append(list, "4")

	fmt.Printf("%v", list)
	fmt.Printf("%v", newList)

}