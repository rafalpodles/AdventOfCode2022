package day6

import (
	"bufio"
	"fmt"
	"os"
)

func Solve() {
	file, err := os.Open("data6.txt")
	checkError(err)
	defer func(file *os.File) {
		checkError(err)
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		task1(s)
		task2(s)
	}

}

func task1(s string) {
	chars := []int32{'*', '*', '*', '*'}
	solution(chars, s)

}

func task2(s string) {
	chars := []int32{'*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*'}
	solution(chars, s)

}

func solution(chars []int32, s string) {
	for i, ch := range s {
		if i < len(chars) {
			chars[i] = ch
		}
		if i >= len(chars) {
			if isArrayUniq(chars) {
				fmt.Println(i)
				break
			} else {
				chars = chars[1:]
				chars = append(chars, ch)

			}
		}
	}
}
func isArrayUniq(arr []int32) bool {
	for i, ch := range arr {
		for j := i + 1; j < len(arr); j++ {
			if ch == arr[j] {
				return false
			}
		}

	}
	return true
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
