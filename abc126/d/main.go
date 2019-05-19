package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func line() string {
	var text string
	sc.Scan()
	if sc.Err() == nil {
		text = sc.Text()
	}
	return text
}

func main() {
	n, _ := strconv.Atoi(line())

	var grp1 []int
	var grp2 []int

	for {
		l := line()
		if l == "" {
			break
		}
		input := strings.Split(l, " ")

		// fmt.Println(input)

		a, _ := strconv.Atoi(input[0])
		b, _ := strconv.Atoi(input[1])
		c, _ := strconv.Atoi(input[2])

		if contains(grp1, a) || contains(grp2, b) {
			if c&0x1 == 1 {
				grp1 = append(grp1, a)
				grp2 = append(grp2, b)
			} else {
				grp1 = append(grp1, a)
				grp1 = append(grp1, b)
			}
		} else {
			if c&0x1 == 1 {
				grp1 = append(grp1, b)
				grp2 = append(grp2, a)
			} else {
				grp2 = append(grp2, a)
				grp2 = append(grp2, b)
			}
		}
	}
	fmt.Println(grp1)
	fmt.Println(grp2)
	for k := 0; k < n; k++ {
		if contains(grp1, k+1) {
			fmt.Println("1")
		} else {
			fmt.Println("0")
		}
	}
}

func contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
