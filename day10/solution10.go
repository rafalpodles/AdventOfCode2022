package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	file, err := os.Open("data10.txt")
	checkError(err)
	defer func(file *os.File) {
		checkError(err)
	}(file)

	scanner := bufio.NewScanner(file)

	//task1(*scanner)
	task2(*scanner)

}

func task1(scanner bufio.Scanner) {
	x := 1
	cycle := 1
	points := 0
	multiply := 0
	for scanner.Scan() {
		s := scanner.Text()
		if strings.Contains(s, "addx") {
			v, _ := strconv.Atoi(strings.Split(s, " ")[1])
			for i := 0; i < 2; i++ {

				if cycle == 20+(multiply*40) {
					points = points + (cycle * x)
					multiply++
				}
				if i == 1 {
					x = x + v
				}
				cycle++
			}
		} else {
			if cycle == 20+(multiply*40) {
				points = points + (cycle * x)
				multiply++
			}
			cycle++
		}
	}
	fmt.Println(multiply)
	fmt.Println(points)
}

func task2(scanner bufio.Scanner) {
	var disp [6][40]int32
	for i := range disp[0] {
		for j := 0; j < 6; j++ {
			disp[j][i] = '.'
		}
	}

	x := 1
	cycle := 1
	multiply := 0
	for scanner.Scan() {
		s := scanner.Text()
		if strings.Contains(s, "addx") {
			v, _ := strconv.Atoi(strings.Split(s, " ")[1])
			for i := 0; i < 2; i++ {
				if cycle == 40+(multiply*40) {
					multiply++
				}
				if i == 1 {
					x = x + v
				}
				if x == cycle%40 || x-1 == cycle%40 || x+1 == cycle%40 {
					disp[multiply][(cycle)%40] = '#'
				}
				cycle++
			}
		} else {
			if cycle == 40+(multiply*40) {
				disp[multiply][(cycle)%40] = '#'
				multiply++
			}
			if x == cycle%40 || x-1 == cycle%40 || x+1 == cycle%40 {
				disp[multiply][(cycle)%40] = '#'
			}
			cycle++
		}

	}
	print(disp)
}

func print(board [6][40]int32) {
	for row := 0; row < 6; row++ {
		for column := 0; column < 40; column++ {
			fmt.Print(string(board[row][column]), " ")
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
