package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CalculateAreaofRectangle(length float64, width float64) float64 {
	return length * width
}

func UserInput(prompt string) (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	input := strings.TrimSpace(text)
	value, err := strconv.ParseFloat(input, 64)
	return value, err
}

func main() {
	length, err := UserInput("Enter the length of a rectangle: ")
	if err != nil || length <= 0 {
		fmt.Println("Invalid input for length. Please enter a valid number greater than zero.")
		return
	}
	width, err := UserInput("Enter the width of a rectangle: ")
	if err != nil || width <= 0 {
		fmt.Println("Invalid input for width. Please enter a valid number greater than zero.")
		return
	}
	area := CalculateAreaofRectangle(length, width)
	fmt.Printf("The area of the rectangle is: %.2f\n", area)
}
