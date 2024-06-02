package main

import (
	"fmt"
	"strings"
)

func FahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func CelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func main() {
	fmt.Println("Please enter your temperature.")

	var input string
	var temperature float64

	fmt.Print("Enter the temperature (e.g., 32F for Fahrenheit, 25C for Celsius): ")
	fmt.Scanln(&input)

	if strings.HasSuffix(strings.ToLower(input), "f") {
		_, err := fmt.Sscanf(input, "%fF", &temperature)
		if err != nil {
			fmt.Println("Invalid input format:", err)
			return
		}
		celsius := FahrenheitToCelsius(temperature)
		fmt.Printf("%.2f°F is equal to %.2f°C\n", temperature, celsius)

	} else if strings.HasSuffix(strings.ToLower(input), "c") {
		_, err := fmt.Sscanf(input, "%fC", &temperature)
		if err != nil {
			fmt.Println("Invalid input format:", err)
			return
		}
		fahrenheit := CelsiusToFahrenheit(temperature)
		fmt.Printf("%.2f°C is equal to %.2f°F\n", temperature, fahrenheit)
	} else {
		fmt.Println("Invalid input format. Please enter a temperature ending with 'F' or 'C'.")
	}
}
