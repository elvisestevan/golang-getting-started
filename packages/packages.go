package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	price, _ := decimal.NewFromString("133.33")

	fmt.Println("Price: ", price)
}
