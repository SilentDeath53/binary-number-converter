package main

import (
	"fmt"
	"strconv"
)

func decimalToBinary(decimal int) string {
	binary := ""
	for decimal > 0 {
		remainder := decimal % 2
		binary = strconv.Itoa(remainder) + binary
		decimal = decimal / 2
	}
	return binary
}

func binaryToDecimal(binary string) int {
	decimal := 0
	power := 0
	for i := len(binary) - 1; i >= 0; i-- {
		bit, _ := strconv.Atoi(string(binary[i]))
		decimal += bit * (1 << power)
		power++
	}
	return decimal
}

func main() {
	var choice int
	var value int

	fmt.Println("Binary Decimal Converter")
	fmt.Println("------------------------")
	fmt.Println("1. Decimal to Binary")
	fmt.Println("2. Binary to Decimal")
	fmt.Print("Enter your choice (1 or 2): ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Print("Enter a decimal number: ")
		fmt.Scanln(&value)
		binary := decimalToBinary(value)
		fmt.Printf("Binary representation: %s\n", binary)
	case 2:
		fmt.Print("Enter a binary number: ")
		fmt.Scanln(&value)
		decimal := binaryToDecimal(strconv.Itoa(value))
		fmt.Printf("Decimal representation: %d\n", decimal)
	default:
		fmt.Println("Invalid choice. Please enter 1 or 2.")
	}
}
