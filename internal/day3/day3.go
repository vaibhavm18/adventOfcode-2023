package day3

import (
	"bufio"
	"fmt"
	"strconv"
	"unicode"
)

func stringToRune(s string) []rune {
	var charList []rune

	for _, char := range s {
		charList = append(charList, char)
	}

	return charList
}

func Sol(scanner *bufio.Scanner) int {
	var mat [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		charList := stringToRune(line)
		mat = append(mat, charList)
	}

	return separateSymbols(&mat)
}

type number struct {
	Val   int
	Row   int
	Start int
	End   int
}

type symbol struct {
	Val    rune
	Row    int
	Col    int
	Number []number
}

/*

0 1 2  3
4 5 6  7
8 9 10 11

*/

func separateSymbols(mat *[][]rune) int {
	var numList []number
	var symbols []symbol
	for i, row := range *mat {
		numStr := ""
		for j, char := range row {
			if unicode.IsDigit(char) {
				numStr += string(char)
			} else {

				if len(numStr) != 0 {

					val, _ := strconv.Atoi(numStr)
					num := number{
						Val:   val,
						Row:   i,
						Start: j - len(numStr),
						End:   j - 1,
					}
					numList = append(numList, num)
					numStr = ""
				}

				if char != 46 {
					sym := symbol{
						Val: char,
						Row: i,
						Col: j,
					}
					symbols = append(symbols, sym)
				}
			}
		}
	}

	return selectAllTheEngine(&symbols, &numList)
}

func selectAllTheEngine(syms *[]symbol, nums *[]number) int {
	ans := 0
	res := 0
	i := 0
	var newList []symbol
	for _, s := range *syms {
		newSim := s
		var sameSim []number
		for _, num := range *nums {
			if hasSymbol(s.Row, s.Col, num.Start, num.End, num.Row) {
				ans += num.Val
				i++
				if s.Val == '*' {
					sameSim = append(sameSim, num)
				}
			}

		}
		if len(sameSim) >= 2 {
			newSim.Number = sameSim
			newList = append(newList, newSim)
		}
	}
	for _, n := range newList {
		now := 1
		for _, l := range n.Number {
			now *= l.Val
		}
		res += now
	}
	fmt.Println(res)
	fmt.Println(i)
	return ans
}

func hasSymbol(sR, sC, nS, nE, nR int) bool {
	if sR == nR-1 || sR == nR || sR == nR+1 {
		if sC-1 == nE || sC == nE || sC+1 == nE {
			return true
		}

		if sC-1 == nS || sC == nS || sC+1 == nS {
			return true
		}

	}
	return false
}
