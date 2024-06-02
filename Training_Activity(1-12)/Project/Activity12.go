package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func printFileContents(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}
	return nil
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("Usage: lineByLine <file1> [<file>...]")
		return
	}

	for _, filename := range flag.Args() {
		fileInfo, err := os.Stat(filename)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		if fileInfo.Mode().IsRegular() {
			fmt.Println(filename, "is a regular file")
		}

		err = printFileContents(filename)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
