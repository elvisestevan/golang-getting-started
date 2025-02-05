package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Car struct {
	Brand string
	Model string
	Year  int
}

func main() {
	printDayOfTheWeek()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	grade := r.Float64() * 10
	fmt.Println("Random grade:", grade)

	if grade < 5 {
		fmt.Println("Failed")
	} else if grade < 7 {
		fmt.Println("Make-up exam")
	} else {
		fmt.Println("Passed! Congratulations!")
	}

	var car *Car

	car = &Car{
		Brand: "Toyota",
		Model: "Corolla",
		Year:  2019,
	}

	printCarInfo(car)
	printCarInfo(nil)
}

func printCarInfo(car *Car) {
	if car != nil && car.Brand != "" && car.Model != "" && car.Year > 0 {
		fmt.Printf("Car brand: %s\n", car.Brand)
		fmt.Printf("Car model: %s\n", car.Model)
		fmt.Printf("Car year: %d\n", car.Year)
		return
	}
	fmt.Println("No car information available.")
	return
}

func printDayOfTheWeek() {
	today := time.Now()

	switch today.YearDay() {
	case 5:
		fmt.Println("Today is 5th. Clean your house.")
	case 10:
		fmt.Println("Today is 10th. Buy some wine.")
	case 15:
		fmt.Println("Today is 15th. Visit a doctor.")
	case 25:
		fmt.Println("Today is 25th. Buy some food.")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Printf("No information available for that day: %d\n", today.YearDay())
	}
}
