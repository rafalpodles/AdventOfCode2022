package day3

import (
	"bufio"
	"fmt"
	"os"
)

func Solve() {
	file, err := os.Open("data3.txt")
	checkError(err)
	defer func(file *os.File) {
		checkError(err)
	}(file)

	scanner := bufio.NewScanner(file)

	task1(*scanner)

}

func task1(scanner bufio.Scanner) {
	points3 := 0
	iter := 0
	var e [3]string

	for scanner.Scan() {
		s := scanner.Text()

		if iter < 3 {
			e[iter] = s
			iter++
		}
		if iter == 3 {
			for _, c1 := range e[0] {
				found := false
				for _, c2 := range e[1] {
					for _, c3 := range e[2] {

						if c1 == c2 && c2 == c3 {
							points3 += countPoint(c1)
							found = true
							break
						}
					}
					if found {
						break
					}
				}
				if found {
					break
				}

			}

			iter = 0
		}

		//s1 := s[:(len(s))/2]
		//s2 := s[len(s)/2:]
		//for _, c1 := range s1 {
		//	found := false
		//	for _, c2 := range s2 {
		//		if c1 == c2 {
		//			points3 += countPoint(c1)
		//			found = true
		//			break
		//		}
		//	}
		//	if found {
		//		break
		//	}
		//
		//}
	}
	fmt.Println(points3)
}

func countPoint(c int32) int {

	if c > 96 {
		return int(c - 96)
	} else {
		return int(c - 38)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
