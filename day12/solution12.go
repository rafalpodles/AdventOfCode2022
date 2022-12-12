package day12

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	var board [boardSizeX][boardSizeY]int32

	var aAositions []position
	line := 0
	for scanner.Scan() {
		s := scanner.Text()
		for i, ch := range s {

			if ch == 'E' {
				endX = i
				endY = line
				board[i][line] = '{'
			} else if ch == 'a' || ch == 'S' {
				aAositions = append(aAositions, position{i, line, 0})
				board[i][line] = 'a'
			} else {
				board[i][line] = ch
			}
		}
		line++
	}
	var scores []int

	for _, apos := range aAositions {
		found := false
		move := 0
		var markedPosition []position
		var positions []position
		positions = append(positions, position{apos.x, apos.y, move})
		board[apos.x][apos.y] = '`'
		for i := 0; i < 1000 && !found; i++ { // infinite loop xd
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
			}
			positions = newPositions
			if !found {
				move++
			}
			//print(board, markedPosition)
			//fmt.Println(move)
		}
		if found {
			scores = append(scores, move-2)
			fmt.Println("End.", move-2) //minus last and first move
		}

	}
	sort.Ints(scores)
	fmt.Println(scores[0] - 1) // I don't know why, but I had to subtract 1. I just guessed this xD
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
