package main

import "fmt"

func showWelcome() {
	fmt.Println("Welcome to Shakil Shop")
	fmt.Println("----------------------")
}

func showGoodbye() {
	fmt.Println("----------------------")
	fmt.Println("Thank you for shopping")
	fmt.Println("That's all.....")
}

func showPrice(itemPrice int) {
	fmt.Println("Price:", itemPrice)
}

func showQuantity(itemQuantity int) {
	fmt.Println("Quantity:", itemQuantity)
}

func showTotal(total int) {
	fmt.Println("Total bill:", total)
}

func showDiscount(discount int) {
	fmt.Println("Discount:", discount)
}
func finalBeforeDelivery(finalBeforeDelivery int) {
	fmt.Println("Final Before Delivery", finalBeforeDelivery)
}
func deliveryChargee(deliveryCharge int) {
	fmt.Println("delivery charge:", deliveryCharge)
}
func payableAmountt(payableAmount int) {
	fmt.Println("payable amount:", payableAmount)
}

func main() {
	showWelcome()

	price := 10
	quantity := 10

	total := 0
	total = price * quantity

	discount := 0
	if total >= 100 {
		discount = 10
	}

	finalBeforeDelivery := total - discount

	deliveryCharge := 0
	if finalBeforeDelivery >= 100 {
		deliveryCharge = 0
	} else {
		deliveryCharge = 20
	}

	payableAmount := finalBeforeDelivery + deliveryCharge

	showPrice(price)
	showQuantity(quantity)
	showTotal(total)
	showDiscount(discount)
	deliveryChargee(deliveryCharge)
	payableAmountt(payableAmount)

	showGoodbye()
}
