package template

import (
	"bufio"
	"fmt"
	"os"
)

func Solve() {
	file, err := os.Open("datax.txt")
	checkError(err)
	defer func(file *os.File) {
		checkError(err)
	}(file)

	scanner := bufio.NewScanner(file)

	task1(*scanner)
	task2(*scanner)

}

func task1(scanner bufio.Scanner) {
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)

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
