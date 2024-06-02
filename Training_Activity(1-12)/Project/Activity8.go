package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your word: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)

	if len(word) == 0 {
		fmt.Println("You must enter a word")
	} else {
		words := strings.Fields(word) // Split the input into words
		numWords := len(words)
		fmt.Printf("Number of words: %d\n", numWords)

		numChars := len([]rune(word))
		fmt.Printf("Number of chars: %d\n", numChars)

		letterCounts := make(map[rune]int)
		for _, char := range word {
			letterCounts[char]++
		}
		fmt.Println("Occurrences of specific letters:")
		for letter, count := range letterCounts {
			fmt.Printf("%c: %d\n", letter, count)
		}
	}
}
