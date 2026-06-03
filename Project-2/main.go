package main

import "fmt"

type car struct {
	brand string
	model string
	year  int
}

func main() {
	cars := []car{
		{brand: "Toyota", model: "corolla", year: 2026},
		{brand: "Honda", model: "Civic", year: 2026},
		{brand: "BMW", model: "M5", year: 2026},
	}
	for i := 0; i < len(cars); i++ {
		fmt.Println("Car", i, cars[i])
	}

}
