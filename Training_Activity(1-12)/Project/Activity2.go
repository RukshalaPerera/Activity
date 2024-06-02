package main

import "fmt"

func main() {
	fmt.Println("Please enter your name:")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello, %s!\n", name)
}
