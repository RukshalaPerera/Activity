package main

import (
	"fmt"
	"math"
)

func CalculateMonthlyLoanPayment(rate, loan, payment float64) float64 {
	monthlyRate := rate / 12 / 100
	MonthlyLoanPayment := (loan + monthlyRate) / (1 - math.Pow((1+monthlyRate), -payment))
	return MonthlyLoanPayment
}
func main() {
	var rate, loan, payment float64

	fmt.Println("Enter rate: ")
	fmt.Scan(&rate)

	fmt.Println("Enter the principal loan amount: ")
	fmt.Scan(&loan)

	fmt.Println("Enter the term in months: ")
	fmt.Scan(&payment)

	result := CalculateMonthlyLoanPayment(rate, loan, payment)
	fmt.Println("Monthly loan payment is:", result)
}
