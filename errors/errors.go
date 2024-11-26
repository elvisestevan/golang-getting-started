package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func main() {

	sqrt, err := Sqrt(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sqrt(2) = %f\n", sqrt)

	sqrt2, err2 := Sqrt(-1)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("Sqrt(2) = %.2f\n", sqrt2)
}

func Sqrt(value int) (float64, error) {
	if value < 0 {
		//return 0, fmt.Errorf("math: square root of negative number %d", value)
		return 0, errors.New("math: square root of negative number: " + fmt.Sprint(value))
	}

	return math.Sqrt(float64(value)), nil
}
