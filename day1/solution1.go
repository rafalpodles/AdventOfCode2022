package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Solve() {
	file, err := os.Open("data1.txt")
	checkError(err)
	defer func(file *os.File) {
		checkError(err)
	}(file)

	scanner := bufio.NewScanner(file)

	task1(*scanner)

}

func task1(scanner bufio.Scanner) {
	var x = 0
	var arr []int
	for scanner.Scan() {
		s := scanner.Text()
		if s != "" {
			i, _ := strconv.Atoi(s)
			x += i
		} else {
			arr = append(arr, x)
			x = 0
		}
	}
	max := findMinAndMax(arr)
	fmt.Println("max:", max)
	top3 := findTop3(arr)
	fmt.Println("top3:", top3)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func findMinAndMax(a []int) (max int) {
	max = a[0]
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}
func findTop3(array []int) (top int) {
	length := len(array)
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	sum := array[length-1] + array[length-2] + array[length-3]
	return sum
}
