package main

import (
	"fmt"
	"math/rand"
	"time"
)

func guessingNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10) + 1
}

func main() {
	K := 3
	number := guessingNumber()
	var guess int
	var input string
	fmt.Println("A number is chosen between 1 to 10. Guess the number within 3 trials.")

	for i := 0; i < K; i++ {
		for {
			fmt.Println("Guess the number:")

			_, err := fmt.Scanln(&input)
			if err != nil {
				fmt.Println("Invalid input. Please enter an number.")
				continue
			}

			_, err = fmt.Sscanf(input, "%d", &guess)
			if err != nil {
				fmt.Println("Invalid input. Please enter an number.")
				continue
			}

			break
		}

		if number == guess {
			fmt.Println("Congratulations! You guessed the number.")
			return
		} else if number > guess && i != K-1 {
			fmt.Println("The number is greater than", guess)
		} else if number < guess && i != K-1 {
			fmt.Println("The number is less than", guess)
		}
	}
	fmt.Println("You have exhausted", K, "trials.")
	fmt.Println("The number was", number)
}
