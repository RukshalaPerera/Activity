package main

import (
	"fmt"
	"strconv"
)

func main() {
	var operator string
	fmt.Println("Enter operator: +, -, *, /")
	fmt.Scanln(&operator)

	if operator == "" || (operator != "+" && operator != "-" && operator != "*" && operator != "/") {
		fmt.Println("Please enter a valid operator")
		return
	}

	var num1, num2 float64
	var input1, input2 string

	fmt.Println("Enter first number:")
	fmt.Scanln(&input1)
	fmt.Println("Enter second number:")
	fmt.Scanln(&input2)

	var err1, err2 error
	num1, err1 = strconv.ParseFloat(input1, 64)
	num2, err2 = strconv.ParseFloat(input2, 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid input for numbers")
		return
	}

	switch operator {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Cannot divide by zero")
			return
		}
		fmt.Println(num1 / num2)
	default:
		fmt.Println("Invalid operator")
	}
}
