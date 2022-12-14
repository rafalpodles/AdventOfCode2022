package day14

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Solve that was surprisingly easy :)
func Solve() {
	file, err := os.Open("data14.txt")
	checkError(err)
	defer func(file *os.File) {
		checkError(err)
	}(file)

	scanner := bufio.NewScanner(file)

	task(*scanner, 1)

}

const boardSizeX = 1000
const boardSizeY = 200

func task(scanner bufio.Scanner, task int) {

	var board [boardSizeX][boardSizeY]int32

	for i := 0; i < boardSizeX; i++ {
		for j := 0; j < boardSizeY; j++ {
			board[i][j] = '.'
		}
	}
	maxY := 0

	for scanner.Scan() {
		fromX := 0
		fromY := 0
		s := scanner.Text()
		slices := strings.Split(s, " -> ")

		for _, x := range slices {
			toXY := strings.Split(x, ",")
			toX, _ := strconv.Atoi(toXY[0])
			toY, _ := strconv.Atoi(toXY[1])
			if toY > maxY {
				maxY = toY
			}
			if fromX != 0 && fromY != 0 {
				if fromY == toY { //move horizontal
					if fromX < toX {
						for i := fromX; i < toX+1; i++ {
							board[i][toY] = '#'
						}
					}
					if fromX > toX {
						for i := toX; i < fromX+1; i++ {
							board[i][toY] = '#'
						}
					}

				}
				if fromX == toX { //move vertical
					if fromY < toY {
						for i := fromY; i < toY+1; i++ {
							board[toX][i] = '#'
						}
					}
					if fromY > toY {
						for i := toY; i < fromY+1; i++ {
							board[toX][i] = '#'
						}
					}
				}

			}
			fromX = toX
			fromY = toY

		}

	}
	if task == 2 {
		maxY = maxY + 2
		for i := 0; i < boardSizeX; i++ {
			board[i][maxY] = '#'
		}
	}

	for x := 0; x < 100000; x++ { // Infinite loop xD
		col := 500
		for i := 1; i < maxY+1; i++ {
			if board[col][i] == '.' {
				continue
			} else if board[col-1][i] == '.' {
				col--
				continue
			} else if board[col+1][i] == '.' {
				col++
				continue
			} else {
				board[col][i-1] = 'o'
				break
			}
		}
	}
	print(board)
	fmt.Println(countSand(board))
}

func countSand(board [boardSizeX][boardSizeY]int32) int {
	points := 0
	for row := 0; row < boardSizeY; row++ {
		for column := 300; column < boardSizeX-300; column++ {
			if board[column][row] == 'o' {
				points++
			}
		}
	}
	return points
}

func print(board [boardSizeX][boardSizeY]int32) {
	for row := 0; row < boardSizeY; row++ {
		for column := 450; column < boardSizeX-450; column++ {
			fmt.Print(string(board[column][row]))
		}
		fmt.Print("\n")
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
