package day7

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Solve() {
	file, err := os.Open("data7.txt")
	checkError(err)
	defer func(file *os.File) {
		checkError(err)
	}(file)

	scanner := bufio.NewScanner(file)

	task1(*scanner)
	task2(*scanner)

}

func task1(scanner bufio.Scanner) {
	var map_1 map[string]int
	position := ""
	for scanner.Scan() {
		s := scanner.Text()
		if strings.HasPrefix(s, "$") {
			if strings.Contains(s, "cd") {
				split := strings.Split(s, " ")
				if split[2] == ".." {
					index := strings.LastIndex(split[2], "/")
					position = strings.SplitAfterN(position, "/", index-1)[0]
				} else {
					position = position + "/" + split[2]
					map_1[position] = 0
				}

			}
		}

	}
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
