package day8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Solve() {
	file, err := os.Open("data8.txt")
	checkError(err)
	defer func(file *os.File) {
		checkError(err)
	}(file)

	scanner := bufio.NewScanner(file)

	task1(*scanner)
	task2(*scanner)

}

const arraySize = 99

func task1(scanner bufio.Scanner) {

	var trees [arraySize][arraySize]int
	row := 0
	for scanner.Scan() {
		s := scanner.Text()
		for col, t := range s {
			trees[row][col], _ = strconv.Atoi(string(t))
		}
		row++
	}
	points := 0
	maxPoints := 0
	for i := 1; i < len(trees[0])-1; i++ {
		for j := 1; j < len(trees[0])-1; j++ {

			t := trees[i][j]
			left := trees[i][:j]
			right := trees[i][j+1:]
			up := boardColumn(trees, j)[:i]
			down := boardColumn(trees, j)[i+1:]

			if containsAllLower(left, t) || containsAllLower(right, t) || containsAllLower(up, t) || containsAllLower(down, t) {
				points++
			}
			leftPoints := countFreeSpaceAround(left, t, true)
			rightPoints := countFreeSpaceAround(right, t, false)
			upPoints := countFreeSpaceAround(up, t, true)
			downPoints := countFreeSpaceAround(down, t, false)
			viewPoints := leftPoints * rightPoints * upPoints * downPoints
			if maxPoints < viewPoints {
				maxPoints = viewPoints
			}

		}
	}
	points = points + 99 + 99 + 97 + 97
	//points = points + 16
	fmt.Println(points)
	fmt.Println(maxPoints)

}

func containsAllLower(slice []int, el int) bool {
	for _, x := range slice {
		if el > x {
			continue
		} else {
			return false
		}
	}
	return true
}

// revert is for left and up
func countFreeSpaceAround(slice []int, el int, revert bool) int {
	points := 0

	if revert {
		for i := len(slice); i > 0; i-- {
			if el > slice[i-1] {
				points++
			} else {
				points++
				return points
			}
		}
	} else {
		for _, x := range slice {
			if el > x {
				points++
			} else {
				points++
				return points
			}
		}
	}
	return points
}

func boardColumn(board [arraySize][arraySize]int, columnIndex int) (column []int) {
	column = make([]int, 0)
	for _, row := range board {
		column = append(column, row[columnIndex])
	}
	return
}

func isHigher(t int, arr []int) bool {
	for _, el := range arr {
		if el > t {
			return true
		}

	}
	return false
}

func task2(scanner bufio.Scanner) {
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)

	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
