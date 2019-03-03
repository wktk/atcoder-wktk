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
	input := strings.Split(line(), " ")
	a, _ := strconv.Atoi(input[0])
	b, _ := strconv.Atoi(input[1])
	c, _ := strconv.Atoi(input[2])

	var ay []int

	for i := 1; i < 101; i++ {
		if a%i == 0 && b%i == 0 && a >= i && b >= i {
			ay = append(ay, i)
		}
	}

	fmt.Printf("%d\n", ay[len(ay)-c])
}
