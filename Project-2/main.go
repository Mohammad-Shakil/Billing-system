package main

import "fmt"

type car struct {
	Brand string
	Model string
	Year  int
}

func main() {
	cars := []car{
		{Brand: "Toyota", Model: "corolla", Year: 2026},
		{Brand: "Honda", Model: "Civic", Year: 2026},
		{Brand: "BMW", Model: "M5", Year: 2026},
	}
	for i := 0; i < len(cars); i++ {
		fmt.Println(cars[i].Brand, cars[i].Model, cars[i].Year)
	}

}
