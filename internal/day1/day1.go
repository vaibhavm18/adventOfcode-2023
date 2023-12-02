package day1

import (
	"bufio"
	"fmt"
	"strconv"
)

func InitializeNumberMap() map[string]string {
	return map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
}

func Answer(scanner *bufio.Scanner) {
	numMap := InitializeNumberMap()
	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		res := numInStr(line, &numMap)
		ans += res
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("scanner error ", err)
	}

	fmt.Println(ans)
}

func numInStr(s string, l *map[string]string) int {
	var substrings []string

	for i, char := range s {
		if char >= '0' && char <= '9' {
			substrings = append(substrings, string(char))
			continue
		}
		for j := i + 1; j < len(s); j++ {
			substring := s[i : j+1]

			num, exist := (*l)[substring]
			if exist {
				substrings = append(substrings, num)
			}
		}
	}

	fmt.Println(substrings)

	first, last := substrings[0], substrings[len(substrings)-1]

	numString := first + last

	res, err := strconv.Atoi(numString)

	if err != nil {
		return 0
	}

	return res

}
