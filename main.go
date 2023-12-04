package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/vaibhavm18/adventOfcode-2023/internal/day4"
)

func main() {
	filePath := "input/test.txt"
	input := os.Args[1:]

	if len(input) > 0 && input[0] != "" {
		filePath = fmt.Sprintf("input/%s.txt", input[0])
	}
	scanner, file, err := file(&filePath)
	if err != nil {
		fmt.Println("file error", err)
		return
	}

	defer file.Close()
	ans := day4.Answer(scanner)
	fmt.Println(ans)
}

func file(filePath *string) (*bufio.Scanner, *os.File, error) {
	file, err := os.Open(*filePath)

	if err != nil {
		return nil, nil, err
	}

	scanner := bufio.NewScanner(file)
	return scanner, file, nil
}
