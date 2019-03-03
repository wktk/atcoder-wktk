package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var sc = bufio.NewScanner(os.Stdin)

func line() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	input := line()
	cand := []string{input}

	for q := []string{input}; len(q) > 0; {
		next := q[0]
		q = q[1:]
		for _, index := range avail(next) {
			did := step(next, index)
			if !contains(cand, did) {
				cand = append(cand, did)
				q = append(q, did)
			}
		}
	}

	sort.Slice(cand, func(i, j int) bool { return len(cand[i]) < len(cand[j]) })
	last := cand[0]
	fmt.Printf("%d\n", len(input)-len(last))
}

func step(input string, index int) string {
	var bef string
	var aft string
	if index > 0 {
		bef = input[:index]
	} else {
		bef = ""
	}
	if index > len(input)-3 {
		aft = ""
	} else {
		aft = input[index+2:]
	}
	return bef + aft
}

func avail(input string) []int {
	var res []int

	for i := 0; i < len(input)-1; i++ {
		if input[i] != input[i+1] {
			res = append(res, i)
		}
	}

	return res
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
