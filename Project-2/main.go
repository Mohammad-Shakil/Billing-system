package main

import "fmt"

func main() {
	total := 0
	marks := []int{80, 90, 70}
	for i := 0; i < len(marks); i++ {
		total = total + marks[i]
	}
	average := total / len(marks)
	fmt.Println("Average:", average)
	fmt.Println("Subject count:", len(marks))
	fmt.Println("Total marks:", total)
}
