package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	file, err := os.Open("data4.txt")
	checkError(err)
	defer func(file *os.File) {
		checkError(err)
	}(file)

	scanner := bufio.NewScanner(file)

	task1(*scanner)

}

func task1(scanner bufio.Scanner) {
	points := 0
	points1 := 0
	for scanner.Scan() {
		s := scanner.Text()
		splits := strings.Split(s, ",")
		e1 := splits[0]
		e2 := splits[1]
		e1_s := strings.Split(e1, "-")
		e1_from, _ := strconv.Atoi(e1_s[0])
		e1_to, _ := strconv.Atoi(e1_s[1])
		e2_s := strings.Split(e2, "-")
		e2_from, _ := strconv.Atoi(e2_s[0])
		e2_to, _ := strconv.Atoi(e2_s[1])
		if findOverlapingIncludedAll(e1_from, e1_to, e2_from, e2_to) {
			points++
		}
		if findOverlapingIncluded(e1_from, e1_to, e2_from, e2_to) {
			points1++
		}

	}
	fmt.Println(points)
	fmt.Println(1000 - points1)
}

func findOverlapingIncludedAll(e1_from int, e1_to int, e2_from int, e2_to int) bool {
	if e1_from <= e2_from && e1_to >= e2_to {
		return true
	}
	if e2_from <= e1_from && e2_to >= e1_to {
		return true
	}
	return false
}

func findOverlapingIncluded(e1_from int, e1_to int, e2_from int, e2_to int) bool {
	if e1_from > e2_to || e1_to < e2_from {
		return true
	}
	if e2_from > e1_to || e2_to < e1_from {
		return true
	}
	return false
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
