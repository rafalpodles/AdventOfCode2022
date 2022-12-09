package day2

import (
	"bufio"
	"fmt"
	"os"
)

func Solve() {
	file, err := os.Open("data2.txt")
	checkError(err)
	defer func(file *os.File) {
		checkError(err)
	}(file)

	scanner := bufio.NewScanner(file)

	task1(*scanner)

}

func task1(scanner bufio.Scanner) {
	points := 0
	for scanner.Scan() {
		s2 := scanner.Text()
		p1 := s2[0]
		p2 := determinateShape(p1, s2[2])
		points += getPointsForSelection(p2)
		points += getPointsForDraw(p1, p2)
		points += getPointsForWinning(p1, p2)
	}
	fmt.Println(points)
}

//A,X - Kamien(1punkt)
//B,Y - Papier(2punkty)
//C,Z - Norzyce(3punkty)

func determinateShape(p1 uint8, p2 uint8) uint8 {
	//Lose
	if p2 == 'X' && p1 == 'A' {
		return 'Z'
	}
	if p2 == 'X' && p1 == 'B' {
		return 'X'
	}
	if p2 == 'X' && p1 == 'C' {
		return 'Y'
	}
	//Draw
	if p2 == 'Y' && p1 == 'A' {
		return 'X'
	}
	if p2 == 'Y' && p1 == 'B' {
		return 'Y'
	}
	if p2 == 'Y' && p1 == 'C' {
		return 'Z'
	}
	//Win
	if p2 == 'Z' && p1 == 'A' {
		return 'Y'
	}
	if p2 == 'Z' && p1 == 'B' {
		return 'Z'
	}
	if p2 == 'Z' && p1 == 'C' {
		return 'X'
	}
	return 'D'

}

func getPointsForWinning(p1 uint8, p2 uint8) int {
	if p2 == 'X' && p1 == 'C' {
		return 6
	}
	if p2 == 'Y' && p1 == 'A' {
		return 6
	}
	if p2 == 'Z' && p1 == 'B' {
		return 6
	}
	return 0
}

func getPointsForSelection(selection uint8) int {
	score := 0
	switch selection {
	case 'X':
		score = 1
	case 'Y':
		score = 2
	case 'Z':
		score = 3

	}
	return score
}

func getPointsForDraw(p1 uint8, p2 uint8) int {
	if p1 == 'A' && p2 == 'X' || p1 == 'B' && p2 == 'Y' || p1 == 'C' && p2 == 'Z' {
		return 3
	}
	return 0
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
