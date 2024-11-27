package main

import "fmt"

func main() {
	list := []string{"banana", "apple", "orange", "pineapple", "grape"}

	index, word := Find(list, "orange")
	fmt.Printf("Found %s at index %d\n", word, index)

	index, word = Find(list, "kiwi")
	fmt.Printf("Found %s at index %d\n", word, index)

}

func Find(list []string, word string) (index int, wordFound string) {
	for i, w := range list {
		if w == word {
			wordFound = w
			index = i
			break
		}
	}
	return
}
