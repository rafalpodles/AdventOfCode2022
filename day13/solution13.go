package day13

import "os"

// Jakie to było głupie. Kilka h zmarnowane na parsowanie stringa, a poźniej brak jednego warunku w ifologii xD
// Advent of Parsing String
// 2 zadanie kompletnie bez sensu, nie wiążące się z 1.
import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type element struct {
	num   int
	child []element
}

func Solve() {
	file, err := os.Open("data13.txt")
	checkError(err)
	defer func(file *os.File) {
		checkError(err)
	}(file)

	scanner := bufio.NewScanner(file)

	task1(*scanner)
	//task2(*scanner)

}

func task1(scanner bufio.Scanner) {
	var e1 []element
	var e2 []element
	index := 1
	point := 0
	for scanner.Scan() {
		s := scanner.Text()
		if s != "" {
			if e1 == nil {
				e1 = parseString(s)
				fmt.Println("Parsed", s)
				continue
			}
			if e2 == nil {
				e2 = parseString(s)
				fmt.Println("Parsed", s)
			}
			if findPoints(e1, e2) > 0 {
				fmt.Println("Added", index, "points")
				point = point + index
				fmt.Println("Points:", point)
			}
			fmt.Println("///////")
			index++
			e1 = nil
			e2 = nil
		}

	}

	fmt.Println(point)

}

func parseString(s string) []element {
	s = extractFromBrackets(s)
	var el []element
	if s != "" {
		for i := 0; i < len(s); i++ {
			if s[i] == '[' {
				lastPos := findCloseBracketPosition(s[i:])
				el = append(el, element{-1, parseString(s[i : i+lastPos+1])})
				i = i + lastPos
			} else if s[i] == ']' {
				return el
			} else if s[i] == ',' {
				continue
			} else if s[i] == '1' && len(s) > i+1 && s[i+1] == '0' {
				el = append(el, element{10, nil})
				i++
			} else {
				num, _ := strconv.Atoi(string(s[i]))
				el = append(el, element{num, nil})
			}
		}
	} else {
		el = append(el, element{-1, nil})
	}
	return el

}
func findCloseBracketPosition(s string) int {
	c := 0
	for i, x := range s {
		if x == '[' {
			c++
		}
		if x == ']' {
			c--
		}
		if c == 0 {
			return i
		}

	}
	return 0
}

func findPoints(el1 []element, el2 []element) int { // shame xD
	lastIndex := 0
	for i, e := range el1 {
		if i >= len(el2) {
			fmt.Println("Left run of items for ", el1, el2)
			return -1
		}
		if el2[i].child == nil && e.child != nil {
			e2 := el2[i]
			if e2.num == -1 {
				fmt.Println("Right running out of parameters")
				return -1
			}
			if e2.child != nil {
				x := findPoints(e.child, el2[i].child)
				if x > 0 {
					return 1
				} else if x < 0 {
					return -1
				}
			} else if e2.num != -1 && e2.child == nil {
				e2.child = append(e2.child, element{e2.num, nil})
				e2.num = -1
				x := findPoints(e.child, e2.child)
				if x > 0 {
					return 1
				} else if x < 0 {
					return -1
				}
			}
		} else if e.child == nil && el2[i].child != nil {
			if e.num == -1 {
				fmt.Println("Left run out parameters")
				return 1
			}
			e.child = append(e.child, element{e.num, nil})
			e.num = -1
			x := findPoints(e.child, el2[i].child)
			if x > 0 {
				return 1
			} else if x < 0 {
				return -1
			}
		} else if e.child != nil && el2[i].child != nil {
			x := findPoints(e.child, el2[i].child)
			if x > 0 {
				return 1
			} else if x < 0 {
				return -1
			}
		} else {
			if e.num < el2[i].num {
				fmt.Println("Right higher", el1, el2)
				return 1
			}
			if e.num > el2[i].num {
				fmt.Println("Left higher", el1, el2)
				return -1
			}
		}
		lastIndex++
	}
	if lastIndex < len(el2) {
		fmt.Println("Right run of items for ", el1, el2)
		return 1
	}
	return 0
}

func extractFromBrackets(el string) string {
	if strings.HasPrefix(el, "[") && strings.HasSuffix(el, "]") {
		return el[1 : len(el)-1]
	}
	return el
}

func task2(scanner bufio.Scanner) {
	var e [][]element
	var str []string
	for scanner.Scan() {
		s := scanner.Text()
		if s != "" {
			str = append(str, s)
			e1 := parseString(s)
			e = append(e, e1)

		}
	}
	e = append(e, []element{{-1, []element{{2, nil}}}})
	e = append(e, []element{{-1, []element{{6, nil}}}})
	for i := 0; i < len(str); i++ {
		str[i] = strings.Replace(str[i], "[]", "-1", -1)
		str[i] = strings.Replace(str[i], "[", "", -1)
		str[i] = strings.Replace(str[i], "]", "", -1)
	}
	sort.Strings(str)
	p := 1
	ten := 0
	for i, s := range str {
		if strings.HasPrefix(s, "10") { // xDD dobre sortowanie. odejmuje 10 bo są pomiędzy 1 i 2 xDD
			ten++
		}
		if s == "2" || s == "6" {
			p = p * (i - ten + 1)
			fmt.Println(i - ten + 1)
		}
	}

	fmt.Println(p)

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
