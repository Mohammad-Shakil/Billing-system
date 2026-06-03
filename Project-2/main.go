package main

import "fmt"

type Toy struct {
	Name     string
	Price    int
	Avilable bool
}

func isExpensive(toys Toy) bool {
	return toys.Price >= 500
}
func main() {

	toys := []Toy{
		{Name: "Arrma", Price: 400, Avilable: true},
		{Name: "Dhik", Price: 200, Avilable: false},
		{Name: "Wltoys", Price: 600, Avilable: true},
	}
	for i := 0; i < len(toys); i++ {
		if isExpensive(toys[i]) {
			fmt.Println("Toys is expensive", toys[i].Name)
		}
	}
}
