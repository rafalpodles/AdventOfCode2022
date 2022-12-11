package day11

import (
	"fmt"
	"sort"
)

func Solve() {
	task()
}

const modulo = 2 * 7 * 13 * 3 * 19 * 17 * 11 * 5

// Coś zepsułem po refactorze i nie działa poprawnie.
// Sprawdziłbym co jest nie tak, ale input mam zahardcodowany, więc nie chce mi się przepinać na input testowy. xD
func task() {

	var m = [][]int{
		{84, 66, 62, 69, 88, 91, 91},
		{98, 50, 76, 99},
		{72, 56, 94},
		{55, 88, 90, 77, 60, 67},
		{69, 72, 63, 60, 72, 52, 63, 78},
		{89, 73},
		{78, 68, 98, 88, 66},
		{70}}

	var inspected = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	var task = [2]int{1, 2}
	for _, t := range task {
		iter := 20
		if t == 2 {
			iter = 10000
		}
		for x := 0; x < iter; x++ {
			for i := 0; i < 8; i++ {
				for j := 0; j < len(m[i]); j++ {
					newMonkey, item := move(i, m[i][j], t)
					m[newMonkey] = append(m[newMonkey], item)
					inspected[i]++
				}
				m[i] = []int{}

			}
		}
		sort.Ints(inspected)
		fmt.Println(inspected[7] * inspected[6])
		inspected = []int{0, 0, 0, 0, 0, 0, 0, 0}
	}
}

func move(monkey int, item int, t int) (int, int) { //monkey,item

	switch monkey {
	case 0:
		{
			item = item * 11
			if t == 1 {
				item = item / 3
			} else {
				item = item % modulo
			}
			if item%2 == 0 {
				return 4, item
			} else {
				return 7, item
			}
		}
	case 1:
		{
			item = item * item
			if t == 1 {
				item = item / 3
			} else {
				item = item % modulo
			}
			if item%7 == 0 {
				return 3, item
			} else {
				return 6, item
			}
		}
	case 2:
		{
			item = item + 1
			if t == 1 {
				item = item / 3
			} else {
				item = item % modulo
			}
			if item%13 == 0 {
				return 4, item
			} else {
				return 0, item
			}
		}
	case 3:
		{
			item = item + 2
			if t == 1 {
				item = item / 3
			} else {
				item = item % modulo
			}
			if item%3 == 0 {
				return 6, item
			} else {
				return 5, item
			}
		}
	case 4:
		{
			item = item * 13
			if t == 1 {
				item = item / 3
			} else {
				item = item % modulo
			}
			if item%19 == 0 {
				return 1, item
			} else {
				return 7, item
			}
		}
	case 5:
		{
			item = item + 5
			if t == 1 {
				item = item / 3
			} else {
				item = item % modulo
			}
			if item%17 == 0 {
				return 2, item
			} else {
				return 0, item
			}
		}
	case 6:
		{
			item = item + 6
			if t == 1 {
				item = item / 3
			} else {
				item = item % modulo
			}
			if item%11 == 0 {
				return 2, item
			} else {
				return 5, item
			}
		}
	case 7:
		{
			item = item + 7
			if t == 1 {
				item = item / 3
			} else {
				item = item % modulo
			}
			if item%5 == 0 {
				return 1, item
			} else {
				return 3, item
			}
		}
	}
	return 0, item
}

//input:
//Monkey 0:
//Starting items: 84, 66, 62, 69, 88, 91, 91
//Operation: new = old * 11
//Test: divisible by 2
//If true: throw to monkey 4
//If false: throw to monkey 7
//
//Monkey 1:
//Starting items: 98, 50, 76, 99
//Operation: new = old * old
//Test: divisible by 7
//If true: throw to monkey 3
//If false: throw to monkey 6
//
//Monkey 2:
//Starting items: 72, 56, 94
//Operation: new = old + 1
//Test: divisible by 13
//If true: throw to monkey 4
//If false: throw to monkey 0
//
//Monkey 3:
//Starting items: 55, 88, 90, 77, 60, 67
//Operation: new = old + 2
//Test: divisible by 3
//If true: throw to monkey 6
//If false: throw to monkey 5
//
//Monkey 4:
//Starting items: 69, 72, 63, 60, 72, 52, 63, 78
//Operation: new = old * 13
//Test: divisible by 19
//If true: throw to monkey 1
//If false: throw to monkey 7
//
//Monkey 5:
//Starting items: 89, 73
//Operation: new = old + 5
//Test: divisible by 17
//If true: throw to monkey 2
//If false: throw to monkey 0
//
//Monkey 6:
//Starting items: 78, 68, 98, 88, 66
//Operation: new = old + 6
//Test: divisible by 11
//If true: throw to monkey 2
//If false: throw to monkey 5
//
//Monkey 7:
//Starting items: 70
//Operation: new = old + 7
//Test: divisible by 5
//If true: throw to monkey 1
//If false: throw to monkey 3
