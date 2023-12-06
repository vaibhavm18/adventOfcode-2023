package day4

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Game struct {
	Num   int
	First []int
	last  []int
	win   int
}

func Answer(scanner *bufio.Scanner) int {

	var gameList []Game
	var card []int
	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		game := parseString(line)
		gameList = append(gameList, game)
		card = append(card, 1)
	}

	for i, game := range gameList {
		res := findMatch(&game.First, &game.last)
		if res != 0 {
			pow1 := math.Pow(2.0, float64(res-1))
			ans += int(pow1)
		}
		gameRef := &gameList[i]
		gameRef.win = res
		for k := 0; k < card[i]; k++ {
			l := i
			for j := 0; j < res; j++ {
				l += 1
				card[l] = card[l] + 1
			}

		}
	}

	part2 := 0

	for _, num := range card {
		part2 += num
	}

	fmt.Println("part 2", part2)
	return ans
}

func findMatch(first *[]int, last *[]int) int {
	res := 0
	for _, i := range *first {
		for _, j := range *last {
			if i == j {
				res += 1
			}
		}
	}

	return res
}

func parseString(s string) Game {
	splitString := strings.Split(s, ": ")
	parts := strings.Split(splitString[1], " | ")

	return Game{
		Num:   gameNum(splitString[0]),
		First: parseRes(parts[0]),
		last:  parseRes(parts[1]),
	}
}

func gameNum(s string) int {
	gameNum := strings.Split(s, " ")
	num, _ := strconv.Atoi(gameNum[1])
	return num
}

func parseRes(s string) []int {
	var arr []int
	nums := strings.Split(s, " ")
	for _, num := range nums {
		strToNum, _ := strconv.Atoi(num)
		if strToNum != 0 {
			arr = append(arr, strToNum)
		}
	}
	return arr
}
