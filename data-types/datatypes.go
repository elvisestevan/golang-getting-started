package main

import "fmt"

func main() {
	var Huge int = 1 << 62
	var Huge2 = Huge >> 63
	fmt.Println(Huge, Huge2)
}
