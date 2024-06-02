package main

import (
	"fmt"
	"unicode"
)

func main() {
	var password string
	fmt.Println("Enter your password:")
	fmt.Scanln(&password)

	strength := checkPasswordStrength(password)
	fmt.Printf("Your password strength is: %s\n", strength)
}

func checkPasswordStrength(password string) string {
	var (
		hasUpperCase bool
		hasLowerCase bool
		hasNumber    bool
		passwordLen  = len(password)
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpperCase = true
		case unicode.IsLower(char):
			hasLowerCase = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}

	if passwordLen < 8 {
		return "Weak"
	} else if passwordLen < 12 && (hasUpperCase || hasLowerCase || hasNumber) {
		return "Moderate"
	} else if passwordLen >= 12 && hasUpperCase && hasLowerCase && hasNumber {
		return "Strong"
	} else {
		return "Moderate"
	}
}
