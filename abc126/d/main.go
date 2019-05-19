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
	b_ := n
	n = b_

	var grp [][][]int

	for {
		l := line()
		if l == "" {
			break
		}
		fmt.Println(l)
		input := strings.Split(l, " ")

		a, _ := strconv.Atoi(input[0])
		b, _ := strconv.Atoi(input[1])
		c, _ := strconv.Atoi(input[2])

		grp[a][b] = append(grp[a][b], c)
		grp[b][a] = append(grp[b][a], c)
	}

	for rs := range grp {
		fmt.Println(rs)
	}

	/*
		for k := 1; k <= n; k++ {
		}
		for {
			if contains(grp1, k+1) {
				fmt.Println("1")
			} else {
				fmt.Println("0")

			}
		}
	*/
}

func contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
