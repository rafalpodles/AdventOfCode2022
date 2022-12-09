package day9

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const boardSize = 2000

type point struct {
	x int
	y int
}

func Solve() {

	points := 1
	var sl []point
	sl = append(sl, point{boardSize / 2, boardSize / 2})
	head := point{boardSize / 2, boardSize / 2}
	file, err := os.Open("data9.txt")
	checkError(err)
	defer func(file *os.File) {
		checkError(err)
	}(file)

	scanner := bufio.NewScanner(file)
	tail := [10]point{}
	for i := range tail {
		tail[i] = point{boardSize / 2, boardSize / 2}
	}

	for scanner.Scan() {
		s := scanner.Text()
		direction := strings.Split(s, " ")[0]
		steps, _ := strconv.Atoi(strings.Split(s, " ")[1])

		for i := 0; i < steps; i++ {
			switch direction {
			case "D":
				{
					head.y++
				}
			case "U":
				{
					head.y--
				}
			case "R":
				{
					head.x++

				}
			case "L":
				{
					head.x--
				}

			}

			tail[0] = head
			for j := 0; j < len(tail)-1; j++ {
				if math.Abs(float64(tail[j].x-tail[j+1].x)) > 1 || math.Abs(float64(tail[j].y-tail[j+1].y)) > 1 {
					tail[j+1] = countDirection(tail[j], tail[j+1])
					if j == 8 {
						contains := false
						for _, s := range sl {
							if s.y == tail[9].y && s.x == tail[9].x {
								contains = true
							}
						}
						if !contains {
							sl = append(sl, tail[9])
							points++
						}
					}

				}

			}
		}
	}
	fmt.Println(points)
}

func countDirection(first point, second point) point {
	if first.x > second.x && first.y > second.y {
		//move down right
		return point{second.x + 1, second.y + 1}
	}
	if first.x > second.x && first.y < second.y {
		//move up right
		return point{second.x + 1, second.y - 1}
	}
	if first.x < second.x && first.y < second.y {
		//move up left
		return point{second.x - 1, second.y - 1}
	}
	if first.x < second.x && first.y > second.y {
		//move down left
		return point{second.x - 1, second.y + 1}
	}
	if first.x == second.x {
		if first.y > second.y {
			return point{second.x, second.y + 1}
		}
		if first.y < second.y {
			return point{second.x, second.y - 1}
		}
	}
	if first.y == second.y {
		if first.x > second.x {
			return point{second.x + 1, second.y}
		}
		if first.x < second.x {
			return point{second.x - 1, second.y}
		}
	}
	return second
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
