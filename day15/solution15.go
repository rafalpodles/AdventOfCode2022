package day15

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Kilka komentarzy:
// Na początku chciałem wszystko włożyć w tablice, ale była ona zbyt duża
// Później sprawdzałem linia po linii, czy odległość od któregoś punktu jest mniejsza niż jego zasięg. Liczyło się 5000 linii/15min.
// Ostatnie podejście to wzięcie każdego obrysu i sprawdzenie czy obrys +1 nie pokrywa się z innym zasięgiem. Liczy od strzała :)
// Żona jest na mnie zła, bo cały czas po pracy siedzę przed kompem. Życie :)
func Solve() {
	file, err := os.Open("data15.txt")
	checkError(err)
	defer func(file *os.File) {
		checkError(err)
	}(file)

	scanner := bufio.NewScanner(file)

	task(*scanner)

}

type point struct {
	xs int64
	ys int64
	xb int64
	yb int64
}

type smallPoint struct {
	x int64
	y int64
}

var points []point

func task(scanner bufio.Scanner) {
	for scanner.Scan() {
		s := scanner.Text()
		split := strings.Split(s, ":")
		split1 := strings.Split(split[0], ",")
		split2 := strings.Split(split[1], ",")
		xs, _ := strconv.ParseInt(split1[0][strings.Index(split1[0], "=")+1:], 10, 64)
		ys, _ := strconv.ParseInt(split1[1][strings.Index(split1[1], "=")+1:], 10, 64)
		xb, _ := strconv.ParseInt(split2[0][strings.Index(split2[0], "=")+1:], 10, 64)
		yb, _ := strconv.ParseInt(split2[1][strings.Index(split2[1], "=")+1:], 10, 64)
		points = append(points, point{xs, ys, xb, yb})

	}

	var ypoint int64 = 2000000
	var max int64 = 4000000

	m := make(map[int64]int)
	found := false

	for _, p := range points {
		distance := math.Abs(float64(p.xs-p.xb)) + math.Abs(float64(p.ys-p.yb))
		var pointsAround []smallPoint
		pointsAround = append(pointsAround, smallPoint{p.xs, p.ys - int64(distance)}) //top
		pointsAround = append(pointsAround, smallPoint{p.xs, p.ys + int64(distance)}) //bottom
		for j := 0; j < int(distance); j++ {
			left := p.xs - int64(distance) + int64(j) - 1
			pointsAround = append(pointsAround, smallPoint{left, p.ys - int64(j)}) //left top
			right := p.xs + int64(distance) - int64(j) + 1
			pointsAround = append(pointsAround, smallPoint{right, p.ys - int64(j)}) //right top

			if p.ys+int64(j) == ypoint {
				for x := p.xs - int64(distance) + int64(j); x < p.xs+int64(distance)-int64(j); x++ {
					m[x] = 1
				}
			}
		}

		for j := 1; j < int(distance); j++ {
			left := p.xs - int64(distance) + int64(j) - 1
			right := p.xs + int64(distance) - int64(j) + 1
			pointsAround = append(pointsAround, smallPoint{left, p.ys + int64(j)})  // left bottom
			pointsAround = append(pointsAround, smallPoint{right, p.ys + int64(j)}) //right bottom
			if p.ys-int64(j) == ypoint {
				for x := p.xs - int64(distance) + int64(j); x < p.xs+int64(distance)-int64(j); x++ {
					m[x] = 1
				}
			}

		}

		for _, pa := range pointsAround {
			if pa.x >= 0 && pa.x <= max && pa.y >= 0 && pa.y <= max {
				ok := true
				for _, p1 := range points {
					dp1 := math.Abs(float64(p1.xs-p1.xb)) + math.Abs(float64(p1.ys-p1.yb))
					paTop1 := math.Abs(float64(p1.xs-pa.x)) + math.Abs(float64(p1.ys-pa.y))
					if dp1 >= paTop1 {
						ok = false
					}
				}
				if ok {
					fmt.Println("Found this shit:")
					fmt.Println("x:", pa.x, "y:", pa.y)
					fmt.Println(max*pa.x + pa.y)
					found = true
					break
				}
			}
		}
		if found {
			break
		}

	}

	fmt.Println(len(m))

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
