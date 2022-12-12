package day12

import (
	"bufio"
	"fmt"
	"os"
)

func Solve() {
	file, err := os.Open("data12.txt")
	checkError(err)
	defer func(file *os.File) {
		checkError(err)
	}(file)

	scanner := bufio.NewScanner(file)

	task1(*scanner)
	task2(*scanner)

}

const boardSizeX = 114
const boardSizeY = 41

type position struct {
	x    int
	y    int
	move int
}

func task1(scanner bufio.Scanner) {
	endX := 0
	endY := 0
	startX := 0
	startY := 0
	var board [boardSizeX][boardSizeY]int32
	var positions []position
	line := 0
	for scanner.Scan() {
		s := scanner.Text()
		for i, ch := range s {

			if ch == 'E' {
				endX = i
				endY = line
				board[i][line] = '{'
			} else if ch == 'S' {
				startX = i
				startY = line
				board[i][line] = '`'
			} else {
				board[i][line] = ch
			}
		}
		line++
	}

	move := 0
	var markedPosition []position
	positions = append(positions, position{startX, startY, move})
	found := false
	for i := 0; i < 1000 && !found; i++ {
		var newPositions []position
		for _, p := range positions {

			if p.move == move {
				if p.x == endX && p.y == endY {

					found = true
				}
				if p.x < boardSizeX-1 && board[p.x+1][p.y]-board[p.x][p.y] < 2 {
					newPositions = appendIfNotExist(newPositions, markedPosition, position{p.x + 1, p.y, move + 1})
				}
				if p.y < boardSizeY-1 && board[p.x][p.y+1]-board[p.x][p.y] < 2 {
					newPositions = appendIfNotExist(newPositions, markedPosition, position{p.x, p.y + 1, move + 1})
				}
				if p.x > 0 && board[p.x-1][p.y]-board[p.x][p.y] < 2 {
					newPositions = appendIfNotExist(newPositions, markedPosition, position{p.x - 1, p.y, move + 1})
				}
				if p.y > 0 && board[p.x][p.y-1]-board[p.x][p.y] < 2 {
					newPositions = appendIfNotExist(newPositions, markedPosition, position{p.x, p.y - 1, move + 1})
				}
				markedPosition = appendMarkedPosition(markedPosition, position{p.x, p.y, 0})
			}
			fmt.Println(p)
		}
		positions = newPositions
		if !found {
			move++
		}
	}
	if found {
		fmt.Println("End.", move-2) //minus first and last move
	}
}

func print(board [boardSizeX][boardSizeY]int32, marked []position) {
	for _, m := range marked {
		board[m.x][m.y] = board[m.x][m.y] - 32
	}
	for row := 0; row < boardSizeY; row++ {
		for column := 0; column < boardSizeX; column++ {
			fmt.Print(string(board[column][row]))
		}
		fmt.Print("\n")
	}
}

func appendMarkedPosition(pos []position, posToAppend position) []position {
	contains := false
	for _, p := range pos {
		if p == posToAppend {
			contains = true
			break
		}
	}
	if !contains {
		pos = append(pos, posToAppend)

	}
	return pos
}

func appendIfNotExist(pos []position, markedPositions []position, posToAppend position) []position {
	contains := false
	for _, p := range pos {
		for _, mp := range markedPositions {
			if p == posToAppend || (posToAppend.x == mp.x && posToAppend.y == mp.y) {
				contains = true
				break
			}
		}

	}
	if !contains {
		pos = append(pos, posToAppend)

	}
	return pos
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
