package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Summary struct {
	MaxRed   int
	MaxGreen int
	MaxBlue  int
	Game     int
}

func Answer(scanner *bufio.Scanner) int {
	var summaries []Summary
	for scanner.Scan() {
		line := scanner.Text()
		summarie, err := parseString(line)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		summaries = append(summaries, summarie)
	}

	red, green, blue := 12, 13, 14

	return calculate(&summaries, red, green, blue)
}

func calculate(s *[]Summary, red int, green int, blue int) int {
	ans := 0

	for _, item := range *s {
		ans += (item.MaxBlue * item.MaxRed * item.MaxGreen)
	}
	return ans
}

func parseString(s string) (Summary, error) {
	parts := strings.Split(s, ":")
	num := gameNum(parts[0])

	gameDetail := strings.Split(parts[1], ";")

	var green, red, blue int
	for _, detail := range gameDetail {
		fields := strings.Split(detail, ",")
		for _, field := range fields {
			str := strings.Split(field, " ")
			count, _ := strconv.Atoi(str[1])
			color := str[2]

			switch color {
			case "blue":
				if blue < count {
					blue = count
				}
			case "red":
				if red < count {
					red = count
				}
			case "green":
				if green < count {
					green = count
				}
			}
		}

	}

	return Summary{
		MaxRed:   red,
		MaxGreen: green,
		MaxBlue:  blue,
		Game:     num,
	}, nil
}

func gameNum(s string) int {
	gameNum := strings.Split(s, " ")
	num, err := strconv.Atoi(gameNum[1])
	if err != nil {
		fmt.Println("invalid input")
		os.Exit(1)
	}
	return num
}
