package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/vaibhavm18/adventOfcode-2023/internal/day2"
)

func main() {
	filePath := "input/main.txt"
	scanner, file, err := file(&filePath)
	if err != nil {
		fmt.Println("file error", err)
		return
	}

	defer file.Close()
	ans := day2.Answer(scanner)
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
