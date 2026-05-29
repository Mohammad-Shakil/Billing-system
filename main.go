package main

import "fmt"

func main() {
	fmt.Println("Welcome to the billing system....")
	fmt.Println()
	price := 10
	quantity := 10
	total := 0
	total = price * quantity
	discount := 0
	if total >= 100 {
		discount = 10
	}

	finalBeforedelivery := total - discount

	deliveryCharge := 0
	if finalBeforedelivery >= 100 {
		deliveryCharge = 0
	} else {
		deliveryCharge = 20
	}
	payableAmount := finalBeforedelivery + deliveryCharge
	fmt.Println("Price:", price)
	fmt.Println("Quantity:", quantity)
	fmt.Println("Total bill:", total)
	fmt.Println("Discount:", discount)
	fmt.Println("final Before delivery:", finalBeforedelivery)
	fmt.Println("Delivery charge:", deliveryCharge)
	fmt.Println("Payable amount:", payableAmount)
}
