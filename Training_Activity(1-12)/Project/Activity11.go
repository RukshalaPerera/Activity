package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers          = "0123456789"
	specialChars     = "!@#$%^&*-{};'.<>/?`"
)

func PasswordGenerator(length int, useLetters, useUppercaseLetters, useNumbers, useSpecialChars bool) string {
	var characterSet string

	if useLetters {
		characterSet += lowercaseLetters
	}
	if useUppercaseLetters {
		characterSet += uppercaseLetters
	}
	if useNumbers {
		characterSet += numbers
	}
	if useSpecialChars {
		characterSet += specialChars
	}

	if len(characterSet) == 0 {
		panic("please enter at least one character type")
	}

	password := ""
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	for i := 0; i < length; i++ {
		randomIndex := rng.Intn(len(characterSet))
		password += string(characterSet[randomIndex])
	}

	return password
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var useLetters, useNumbers, useSpecialChars, useUppercaseLetters bool
	var input string

	var length int
	fmt.Println("Enter size of your password:")
	fmt.Scanln(&length)

	fmt.Println("Do you want lowercase letters in your password? (y/n):")
	fmt.Scanln(&input)
	useLetters = input == "y" || input == "Y"

	fmt.Println("Do you want uppercase letters in your password? (y/n):")
	fmt.Scanln(&input)
	useUppercaseLetters = input == "y" || input == "Y"

	fmt.Println("Do you want numbers in your password? (y/n):")
	fmt.Scanln(&input)

	useNumbers = input == "y" || input == "Y"
	fmt.Println("Do you want special characters in your password? (y/n):")
	fmt.Scanln(&input)
	useSpecialChars = input == "y" || input == "Y"

	password := PasswordGenerator(length, useLetters, useUppercaseLetters, useNumbers, useSpecialChars)
	fmt.Println("Generated Password:", password)
}
