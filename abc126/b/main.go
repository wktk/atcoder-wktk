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
	sc.Scan()
	return sc.Text()
}

func main() {
	input := strings.Split(line(), "")
	a, _ := strconv.Atoi(input[0])
	b, _ := strconv.Atoi(input[1])
	c, _ := strconv.Atoi(input[2])
	d, _ := strconv.Atoi(input[3])
	if isMonth(a, b) {
		if isMonth(c, d) {
			fmt.Println("AMBIGUOUS")
		} else {
			fmt.Println("MMYY")
		}
	} else {
		if isMonth(c, d) {
			fmt.Println("YYMM")
		} else {
			fmt.Println("NA")
		}
	}

}

func isMonth(a int, b int) bool {
	if a == 1 {
		if b == 0 || b == 1 || b == 2 {
			return true
		}
	} else if a == 0 && b != 0 {
		return true
	}
	return false
}
